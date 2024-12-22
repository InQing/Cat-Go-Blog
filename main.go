package main

import (
	"Go-Blog/common"
	"Go-Blog/server"
	"Go-Blog/config"
	_ "os"
)

func main() {
	server.App.Start(config.Cfg.Server.Ip, config.Cfg.Server.Port)
}

func init(){
	common.LoadTemplate()
}
