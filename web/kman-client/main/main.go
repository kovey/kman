package main

import (
	"github.com/kovey/cli-go/app"
	"github.com/kovey/debug-go/debug"
	"github.com/kovey/kman/kman-client/cache"
	"github.com/kovey/kman/kman-client/etcd"
)

type serv struct {
	*app.ServBase
}

func (s *serv) Panic(a app.AppInterface) {
	cache.Cached()
	s.ServBase.Panic(a)
}

func (s *serv) Init(app.AppInterface) error {
	if err := etcd.Init(); err != nil {
		return err
	}

	return cache.Cached()
}

func (s *serv) Run(app.AppInterface) error {
	etcd.Watch()
	return cache.Cached()
}

func (s *serv) Shutdown(app.AppInterface) error {
	etcd.Close()
	return nil
}

func main() {
	cli := app.NewApp("kman")
	cli.SetServ(&serv{})
	if err := cli.Run(); err != nil {
		debug.Erro(err.Error())
	}
}
