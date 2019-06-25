package handler

import (
	"context"
	"fmt"
	"go-ChatRom/common"
	"go-ChatRom/util"
	"log"
	"net/http"
	"strings"

	cfg "go-ChatRom/config"

	accountProto "go-ChatRom/service/account/proto"
	dbproxyProto "go-ChatRom/service/dbproxy/proto"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro"
)

var (
	accountCli accountProto.AccountService
	dbproxyCli dbproxyProto.DBProxyService
)

func init() {
	service := micro.NewService()
	// 初始化，解析命令行参数等
	service.Init()

	// 初始化一个account服务的客户端
	accountCli = accountProto.NewAccountService("go.micro.service.account", service.Client())

	// 初始化一个dbproxy服务的客户端
	dbproxyCli = dbproxyProto.NewDBProxyService("go.micro.service.dbproxy", service.Client())
}

// SignupHandler: 响应注册页面
func SignupHandler(c *gin.Context) {
	c.Redirect(http.StatusFound, "/static/view/signup.html")
}

// DoSignupHandler: 处理注册Post请求
func DoSignupHandler(c *gin.Context) {
	// 解析表单参数
	username := c.Request.FormValue("username")
	passwd := c.Request.FormValue("password")

	resp, err := accountCli.Signup(context.TODO(), &accountProto.ReqSignup{
		Username: username,
		Password: passwd,
	})

	if err != nil {
		log.Println(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Code":    resp.Code,
		"Message": resp.Message,
	})
}

// SignInHandler: 响应登录页面
func SignInHandler(c *gin.Context) {
	c.Redirect(http.StatusFound, "/static/view/signin.html")
}

// DoSignInHandler: 处理登录Post请求
func DoSignInHandler(c *gin.Context) {
	// 解析表单参数
	username := c.Request.FormValue("username")
	passwd := c.Request.FormValue("password")

	resp, err := accountCli.Signin(context.TODO(), &accountProto.ReqSignin{
		Username: username,
		Password: passwd,
	})

	if err != nil {
		log.Println(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	if resp.Code == common.StatusOK {

		chatServer := "ws://" + strings.Split(c.Request.Host, ":")[0] + ":" + cfg.ChatServerPort + "/client"
		c.JSON(http.StatusOK, gin.H{
			"Code":       resp.Code,
			"Message":    "更新token成功",
			"Token":      resp.Token,
			"Username":   username,
			"Location":   "/static/view/home.html",
			"ChatServer": chatServer,
		})
		return
	} else {

		c.JSON(http.StatusOK, gin.H{
			"Code":    resp.Code,
			"Message": resp.Message,
		})
	}
}

// 验证token是否合法，应该要放到account service中处理
func IsTokenValid(username string, token string) bool {
	fmt.Println(username + ": " + token)
	return true
}

// 进行账户权限的统一验证
func HTTPInterceptor() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Request.FormValue("username")
		token := c.Request.FormValue("token")

		if len(username) < 5 || len(username) > 12 || !IsTokenValid(username, token) {
			c.Abort()
			c.Abort()
			resp := util.NewRespMsg(
				int(common.StatusTokenInvalid),
				"token无效",
				nil,
			)
			c.JSON(http.StatusOK, resp)
			return
		}
	}
}

func DoCheckUserStatusHandler(c *gin.Context) {
	// 解析表单参数
	username := c.Request.FormValue("username")
	token := c.Request.FormValue("token")

	resp, err := accountCli.GetToken(context.TODO(), &accountProto.ReqGetToken{
		Username: username,
	})

	if err != nil {
		log.Println(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	if resp.Code == common.StatusOK && resp.Token == token {
		c.JSON(http.StatusOK, gin.H{
			"Code":    resp.Code,
			"Message": "",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Code":    common.StatusSessionExpiration,
			"Message": "会话已过期",
		})
	}
}
