package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/trewanek/go-echo-boiler/infrastructure/waf/gin/handler"
)

const (
	DefaultPort = ":1323"
)

type server struct {
	gin *gin.Engine
}

func NewServer() *server {
	return &server{
		gin.Default(),
	}
}

func (server *server) Start() {
	server.run()
}

func (server *server) run() {
	server.routing()
	server.setMiddleware()
	log.Fatal(server.gin.Run(DefaultPort))
}

func (server *server) routing() {
	api := server.gin.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			users := v1.Group("/users")
			{
				users.GET("", handler.FetchUsers)
			}
		}
	}

}

func (server *server) setMiddleware() {
	server.gin.Use(gin.Logger())
	server.gin.Use(gin.Recovery())
}
