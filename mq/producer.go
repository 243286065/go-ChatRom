package mq

import (
	"fmt"
	"go-ChatRom/config"
	"log"

	"github.com/streadway/amqp"
)

var conn_send *amqp.Connection
var ch_chat_send *amqp.Channel

// Publish : 发布Direct消息
func PublishDirectMessage(channel *amqp.Channel, exchange string, routingKey string, msg []byte) bool {
	// 发送
	err = ch_chat_send.Publish(
		"example.direct", // exchange
		routingkey,       // routing key
		false,            // mandatory
		false,            // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)

	log.Printf(" [x] Sent %s", message)
}

func PublishChatMessage(routingKey string, msg []byte) bool {
	return PublishDirectMessage(GetChatChannel(), config.ChatExchangeName, routingKey, msg)
}

func PublishNotifyMessage(message string) bool {
	err := GetNotifyChannel().Publish(
		config.NotifyExchangeName, // exchange name
		"",    // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)

	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	fmt.Println("Send Message: " + message)
	return true
}

func int() {
	conn_send, err := amqp.Dial(config.RabbitURL)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer conn_send.Close()

	ch_chat_send, err := conn_send.Channel()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer ch_chat_send.Close()

	err = ch_chat_send.ExchangeDeclare(
		config.ChatExchangeName, // name
		"direct",                // type
		true,                    // durable
		false,                   // auto-deleted
		false,                   // internal
		false,                   // no-wait
		nil,                     // arguments
	)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
