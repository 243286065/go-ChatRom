package mq

import (
	"fmt"
	"go-ChatRom/config"
	"log"
)

// 接收通知消息
func StartMessageConsumer(stop chan bool, callbackchat func(msg []byte), callbacknotify func(msg []byte)) {
	fmt.Println("---------StartMessageConsumer---------------")
	msgs_notify, err := GetNotifyChannel().Consume(
		GetNotifyQueue().Name, // queue
		"",    // consumer
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		true,  // no-wait
		nil,   // args
	)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	fmt.Println("-----GetChatChannel----")
	msgs_chat, err := GetChatChannel().Consume(
		config.QueneNameChat,
		config.RoutingkeyChat,
		true,  //自动应答
		false, // 非唯一的消费者
		false, // rabbitMQ只能设置为false
		true,  // noWait, false表示会阻塞直到有消息过来
		nil)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	fmt.Println("-----GetChatChannel--2--")

	// go func() {
	// 	// 循环读取channel的数据
	// 	for d := range msgs_notify {
	// 		processErr := callbacknotify(conn, d.Body)
	// 		if processErr {
	// 			// TODO: 将任务写入错误队列，待后续处理
	// 		}
	// 	}
	// }()

	// go func() {
	// 	// 循环读取channel的数据
	// 	for d := range msgs_chat {
	// 		processErr := callbackchat(conn, d.Body)
	// 		if processErr {
	// 			// TODO: 将任务写入错误队列，待后续处理
	// 		}
	// 	}
	// }()
	fmt.Println("-----select----")
	for true {
		select {
		case d := <-msgs_notify:
			fmt.Println(string(d.Body))
			callbacknotify(d.Body)
		case d := <-msgs_chat:
			fmt.Println(string(d.Body))
			callbackchat(d.Body)
		case <-stop:
			fmt.Println("-----Stop----")
			return
		}
	}
}
