package main

import (
	"github.com/kovey/kman/kman-service/module/libs/event"
	_ "github.com/kovey/kman/kman-service/services"
	"github.com/kovey/kom/server"
)

func main() {
	server.Run(event.NewServEvent())
}
