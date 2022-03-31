package main

import (
	"fun.tvapi/app/provider/app"
	"fun.tvapi/app/provider/app/config"
	"fun.tvapi/app/provider/app/console"
	"fun.tvapi/app/provider/app/log"
	"fun.tvapi/app/provider/httpserver"
)

func init() {
	app.Boot()
	config.Boot()
	log.Boot()
}

func main() {
	server := app.Get().Server()
	if server.GetAction() == "start" {
		httpserver.Boot()
		httpserver.Start()
	} else {
		console.Help()
	}
}
