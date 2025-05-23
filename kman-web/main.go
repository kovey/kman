package main

import (
	"time"

	_ "github.com/kovey/kman/kman-web/api"
	"github.com/kovey/kman/kman-web/module/libs/event"
	"github.com/kovey/kow"
)

func main() {
	kow.SetMaxRunTime(10 * time.Second)
	kow.Run(event.NewServEvent())
}
