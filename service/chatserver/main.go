package main

import (
	cfg "go-ChatRom/config"
	"go-ChatRom/service/chatserver/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/client", handler.ClientConnectHandler)
	http.ListenAndServe("0.0.0.0:"+cfg.ChatServerPort, nil)
}
