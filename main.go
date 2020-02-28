package main

import (
	"github.com/trewanek/go-echo-boiler/infrastructure/waf/echo/server"
)

func main() {
	server.NewServer().Run()
}
