package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"go-ChatRom/common"
	"go-ChatRom/mq"
	"go-ChatRom/redis"
	"go-ChatRom/service/chatserver/protocol"
	dbProto "go-ChatRom/service/dbproxy/proto"

	rs "github.com/gomodule/redigo/redis"

	"github.com/gorilla/websocket"
	"github.com/micro/go-micro"
)

var (
	dbCli dbProto.DBProxyService
)

func init() {
	service := micro.NewService()
	// 初始化， 解析命令行参数等
	service.Init()
	// 初始化一个dbproxy服务的客户端
	dbCli = dbProto.NewDBProxyService("go.micro.service.dbproxy", service.Client())
}

// 检查当前会话是否已过期
func CheckSessionExpiration(username string, seesionID string) bool {
	rconn := redis.RedisPool().Get()
	defer rconn.Close()
	res, err := rconn.Do("HGET", "ONLINE_"+username, "sessionID")
	if err != nil {
		fmt.Print(err.Error())
	} else {
		if res == nil {
			return true
		}

		if seesionID == string(res.([]byte)) {
			return false
		}
	}

	return true
}

// 获取在线用户列表
func GetUserOnlineList() string {
	rconn := redis.RedisPool().Get()
	defer rconn.Close()

	userList := ""
	res, err := rs.Values(rconn.Do("KEYS", "ONLINE_*"))
	if err != nil {
		fmt.Print(err.Error())
		return ""
	} else {
		for _, value := range res {
			key := string(value.([]byte))[len("ONLINE_"):]
			if len(userList) > 0 {
				userList = userList + "&" + key
			} else {
				userList = key
			}
		}
		return userList
	}
}

// 用于向客户端发送数据
func ClientWriteHandler(conn *websocket.Conn, username string, sessionID string) {
	// 首先发送在线用户列表
	conn.WriteMessage(websocket.TextMessage, []byte(protocol.FormatTextMessage(username, protocol.UpdateOnlineUserListMessage, GetUserOnlineList())))
	exitFlag := false
	var stop = make(chan bool)
	for !exitFlag {
		time.Sleep(1 * time.Second)

		conn.WriteMessage(websocket.PingMessage, []byte{})

		// 检查当前会话是否已经过期
		if CheckSessionExpiration(username, sessionID) {
			fmt.Println("------Session expired!-----")
			exitFlag = true
		}

		// 收取消息队列中的消息
		mq.StartMessageConsumer(stop, username, func(msg []byte) {
			fmt.Println("Recive cht message:" + string(msg))
			conn.WriteMessage(websocket.TextMessage, msg)
		},
			func(msg []byte) {
				fmt.Println("Recive notify message:" + string(msg))
				conn.WriteMessage(websocket.TextMessage, msg)
			})

	}
	stop <- true
	defer conn.Close()
}

// 用于从客户端接收数据
func ClientReadHandler(conn *websocket.Conn, username string, sessionID string) {
	exitFlag := false
	for !exitFlag {
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err.Error())
			exitFlag = true
			break
		}

		fmt.Println("Recv fromct:" + string(message))
		// 收到消息，解析数据
		keysContent := strings.Split(string(message), "#")
		sendMsg := protocol.FormatTextMessage(keysContent[0], protocol.ChatMessage, keysContent[2])

		mq.PublishChatMessage(keysContent[1], []byte(sendMsg))

		time.Sleep(1 * time.Second)
	}
	defer conn.Close()
}

func ClientConnectHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("======" + r.Form.Get("username"))
	username := r.Form.Get("username")
	token := r.Form.Get("token")

	//这里可以先验证下用户是否合法登录
	resp, err := dbCli.ExecuteGetToken(context.TODO(), &dbProto.ReqGetToken{
		Username: username,
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if resp.Code != common.StatusOK || token != resp.Token {
		fmt.Println("Token校验失败")
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Token校验失败"))
		return
	}

	// 将网络请求变为websocket
	var upgrader = websocket.Upgrader{
		// 解决跨域问题
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	go ClientWriteHandler(conn, username, token)
	go ClientReadHandler(conn, username, token)

	conn.SetCloseHandler(func(code int, text string) error {
		// 登记用户离线
		rconn := redis.RedisPool().Get()
		defer rconn.Close()

		if !CheckSessionExpiration(username, token) {
			rconn.Do("HDEL", "ONLINE_"+username, "sessionID")
			message := websocket.FormatCloseMessage(code, "")
			conn.WriteControl(websocket.CloseMessage, message, time.Now().Add(time.Second))
			// 利用消息队列广播离线消息
			mq.PublishNotifyMessage([]byte(protocol.FormatTextMessage(username, protocol.NotifyUserOfflineMessage, "")))
		} else {
			message := websocket.FormatCloseMessage(code, "你已在其它地方登录，本地已断开连接")
			conn.WriteControl(websocket.CloseMessage, message, time.Now().Add(time.Second))
		}

		return nil
	})

	// 登记用户上线
	rconn := redis.RedisPool().Get()
	defer rconn.Close()
	// 用token作为当前会话的ID
	rconn.Do("HSET", "ONLINE_"+username, "sessionID", token)
	fmt.Println("------------------------------------")
	// 利用消息队列广播上线消息
	mq.PublishNotifyMessage([]byte(protocol.FormatTextMessage(username, protocol.NotifyUserOnlineMessage, "")))
}
