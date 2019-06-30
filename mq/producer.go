package mq

import (
	"go-ChatRom/config"
	"log"

	"github.com/streadway/amqp"
)

// Publish : 发布Direct消息
func PublishDirectMessage(exchange string, routingKey string, msg []byte) bool {
	channel, err := conn.Channel()
	if err != nil {
		log.Fatal(err.Error())
		return false
	}
	defer channel.Close()

	err = channel.ExchangeDeclare(
		exchange, // name
		"direct", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		log.Fatal(err.Error())
		return false
	}

	if nil == channel.Publish(
		exchange,
		routingKey,
		false, // 如果没有对应的queue, 就会丢弃这条消息
		false, //
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        msg}) {
		return true
	}
	return false
}

// Publish : 发布Fanout消息
func PublishFanoutMessage(exchange string, msg []byte) bool {
	channel, err := conn.Channel()
	if err != nil {
		log.Fatal(err.Error())
		return false
	}
	defer channel.Close()

	err = channel.ExchangeDeclare(
		exchange, // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		log.Fatal(err.Error())
		return false
	}

	if nil == channel.Publish(
		exchange,
		"",
		false, // 如果没有对应的queue, 就会丢弃这条消息
		false, //
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        msg}) {
		return true
	}
	return false
}

func PublishChatMessage(routingKey string, msg []byte) bool {
	return PublishDirectMessage(config.ChatExchangeName, routingKey, msg)
}

func PublishNotifyMessage(msg []byte) bool {
	return PublishFanoutMessage(config.NotifyExchangeName, msg)
}
