package main

import (
	"go-ChatRom/service/dbproxy/handler"
	proto "go-ChatRom/service/dbproxy/proto"
	"log"
	"time"

	micro "github.com/micro/go-micro"
)

func main() {
	// 创建一个service
	service := micro.NewService(
		micro.Name("go.micro.service.dbproxy"),
		micro.RegisterTTL(time.Second*10),
		micro.RegisterInterval(time.Second*5),
	)

	service.Init()

	proto.RegisterDBProxyServiceHandler(service.Server(), new(handler.DBProxy))

	if err := service.Run(); err != nil {
		log.Println(err)
	}

}
