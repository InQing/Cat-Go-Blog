package server

import (
	"Go-Blog/router"
	"log"
	"net/http"
)

var App = &CatServer{}

type CatServer struct {
}

func (*CatServer) Start(ip, port string) {
	server := http.Server{
		Addr: ip + ":" + port,
	}
	//路由
	router.Router()
	log.Println("Server start on", ip, ":", port)
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
