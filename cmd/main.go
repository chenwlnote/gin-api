package main

import (
	"chenwlnote.gin-api/app/provider/app"
	"chenwlnote.gin-api/app/provider/app/config"
	"chenwlnote.gin-api/app/provider/app/console"
	"chenwlnote.gin-api/app/provider/app/log"
	"chenwlnote.gin-api/app/provider/httpserver"
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
