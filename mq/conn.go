package mq

import (
	"go-ChatRom/config"
	"log"

	"github.com/streadway/amqp"
)

// 服用一个连接
var conn *amqp.Connection
var err error

// 如果异常关闭，会接收通知
var notifyClose chan *amqp.Error

func init() {

	if initConn() {
		conn.NotifyClose(notifyClose)
	}

	// 断线自动重连
	go func() {
		for {
			select {
			case msg := <-notifyClose:
				conn = nil
				log.Printf("onNotifyConnectClosed: %+v\n", msg)
				initConn()
			}
		}
	}()
}

func initConn() bool {
	conn, err = amqp.Dial(config.RabbitURL)
	if err != nil {
		log.Fatal(err.Error())
		return false
	}

	return conn != nil
}
