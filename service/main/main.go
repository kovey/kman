package main

import (
	"github.com/kovey/kman/service/module/libs/event"
	_ "github.com/kovey/kman/service/services"
	"github.com/kovey/kom/server"
)

func main() {
	server.Run(event.NewServEvent())
}
