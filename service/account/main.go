package main

import (
	"log"
	"time"

	"go-ChatRom/service/account/handler"
	proto "go-ChatRom/service/account/proto"

	micro "github.com/micro/go-micro"
)

func main() {
	// 创建一个service
	service := micro.NewService(
		micro.Name("go.micro.service.account"),
		micro.RegisterTTL(time.Second*10),
		micro.RegisterInterval(time.Second*5),
	)

	service.Init()

	proto.RegisterAccountServiceHandler(service.Server(), new(handler.User))

	if err := service.Run(); err != nil {
		log.Println(err)
	}

}
