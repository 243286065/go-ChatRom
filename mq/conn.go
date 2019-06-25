package mq

import (
	"fmt"
	"go-ChatRom/config"
	"log"

	"github.com/streadway/amqp"
)

var conn *amqp.Connection
var channelChat *amqp.Channel = nil
var channelNotify *amqp.Channel = nil
var queueNotify amqp.Queue
var err error

// 如果异常关闭，会接收通知
var notifyClose chan *amqp.Error

func initChannel() bool {
	fmt.Println("----------1-------------------")
	if channelChat != nil && channelNotify != nil {
		return true
	}

	conn, err = amqp.Dial(config.RabbitURL)
	if err != nil {
		log.Println(err.Error())
		return false
	}

	if channelChat == nil {
		channelChat, err = conn.Channel()
		if err != nil {
			log.Println(err.Error())
			return false
		}
		err = channelChat.ExchangeDeclare(
			config.ChatExchangeName, // name
			"direct",                // type
			true,                    // durable
			false,                   // auto-deleted
			false,                   // internal
			false,                   // no-wait
			nil,                     // arguments
		)
		if err != nil {
			log.Println(err.Error())
			return false
		}
	}
	fmt.Println("----------2-------------------")
	if channelNotify == nil {
		channelNotify, err = conn.Channel()
		if err != nil {
			log.Println(err.Error())
			return false
		}
		err = channelNotify.ExchangeDeclare(
			config.NotifyExchangeName, // name
			"fanout",                  // type
			true,                      // durable
			false,                     // auto-deleted
			false,                     // internal
			false,                     // no-wait
			nil,                       // arguments
		)
		if err != nil {
			log.Println(err.Error())
			return false
		}
	}

	queueNotify, err = channelNotify.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when usused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		log.Println(err.Error())
		return false
	}

	err = channelNotify.QueueBind(
		queueNotify.Name, // queue name
		"",               // routing key
		config.NotifyExchangeName, // exchange
		false,
		nil,
	)
	if err != nil {
		log.Println(err.Error())
		return false
	}

	fmt.Println("----------3-------------------")
	fmt.Println(channelNotify == nil)
	return true
}

func init() {
	if initChannel() {
		channelChat.NotifyClose(notifyClose)
		channelNotify.NotifyClose(notifyClose)
	}
	// 断线自动重连
	go func() {
		for {
			select {
			case msg := <-notifyClose:
				conn = nil
				channelChat = nil
				channelNotify = nil
				log.Printf("onNotifyChannelClosed: %+v\n", msg)
				initChannel()
			}
		}
	}()
}

func GetChatChannel() *amqp.Channel {
	return channelChat
}

func GetNotifyChannel() *amqp.Channel {
	return channelNotify
}

func GetNotifyQueue() *amqp.Queue {
	return &queueNotify
}
