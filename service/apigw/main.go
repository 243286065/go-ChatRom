package main

import (
	cfg "go-ChatRom/config"
	"go-ChatRom/service/apigw/route"
)

func main() {
	r := route.Router()
	r.Run(cfg.WebServerHost)
}
