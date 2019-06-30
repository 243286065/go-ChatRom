package mq

import (
	"fmt"
	"go-ChatRom/config"
	"log"
)

// 接收通知消息
func StartMessageConsumer(stop chan bool, username string, callbackchat func(msg []byte), callbacknotify func(msg []byte)) {
	fmt.Println("---------StartMessageConsumer---------------")
	channelNotify, err := conn.Channel()
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	defer channelNotify.Close()

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
		log.Fatal(err.Error())
		return
	}

	queueNotify, err := channelNotify.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when usused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	err = channelNotify.QueueBind(
		queueNotify.Name, // queue name
		"",               // routing key
		config.NotifyExchangeName, // exchange
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	msgsNotify, err := channelNotify.Consume(
		queueNotify.Name, // queue
		"",               // consumer
		true,             // auto-ack
		false,            // exclusive
		false,            // no-local
		false,            // no-wait
		nil,              // args
	)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	channelChat, err := conn.Channel()
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	defer channelChat.Close()

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
		log.Fatal(err.Error())
		return
	}

	queueChat, err := channelChat.QueueDeclare(
		config.QueneNameChat, // name
		false,                // durable
		false,                // delete when usused
		true,                 // exclusive
		false,                // no-wait
		nil,                  // arguments
	)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	err = channelChat.QueueBind(
		queueChat.Name,          // queue name
		username,                // routing key
		config.ChatExchangeName, // exchange
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	msgsChat, err := channelChat.Consume(
		queueChat.Name, // queue
		"",             // consumer
		true,           // auto-ack
		false,          // exclusive
		false,          // no-local
		false,          // no-wait
		nil,            // args
	)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	for true {
		select {
		case d := <-msgsNotify:
			fmt.Println(string(d.Body))
			callbacknotify(d.Body)
		case d := <-msgsChat:
			fmt.Println(string(d.Body))
			callbackchat(d.Body)
		case label := <-stop:
			fmt.Println("-----Stop----", label)
			return
		}
	}
}
