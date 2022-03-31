package app

import (
	"os"
)

var app App

type App struct {
	server Server
}

type Server struct {
	env         string
	action      string
	projectPath string
}

func init() {
	println("app boot init")
	initServer()
}

func initServer() {
	if len(os.Args) == 1 {
		app.server.action = os.Args[1]
	}
	if len(os.Args) > 1 {
		app.server.action = os.Args[1]
		app.server.env = os.Args[2]
	}
	if projectPath, err := os.Getwd(); err != nil {
		panic("获取项目路径异常")
	} else {
		app.server.projectPath = projectPath
	}
}

func Boot() {
	println("app boot")
}

func Get() *App {
	return &app
}

func (app *App) Server() *Server {
	return &app.server
}

func (serv *Server) GetEnv() string {
	return serv.env
}

func (serv *Server) GetAction() string {
	return serv.action
}

func (serv *Server) GetProjectPath() string {
	return serv.projectPath
}

func (app *App) IsDevOrTesting() bool {
	return app.server.env == "dev" || app.server.env == "testing"
}
