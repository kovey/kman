package config

import (
	"os"

	"github.com/kovey/discovery/krpc"
	"github.com/kovey/kman/web/module/libs/middlewares"
	"github.com/kovey/kman/web/module/services"
	"github.com/kovey/kow"
	"github.com/kovey/kow/context"
	"github.com/kovey/kow/controller"
)

func init() {
	router := kow.GET("/admin/config", newList(services.Config())).Rule("key", "maxlen:int:255")
	router.Rule("page", "ge:int64:1").Data(services.ConfigListData())
	router.Rule("page_size", "ge:int64:1")
	router.Middleware(middlewares.NewAuth(), middlewares.NewCheckOpt())
}

type list struct {
	*controller.Base
	service services.ConfigInterface
}

func newList(service services.ConfigInterface) *list {
	return &list{Base: controller.NewBase(os.Getenv("SERVICE_GROUP"), krpc.ServiceName(os.Getenv("SERVICE_NAME"))), service: service}
}

func (a *list) Action(ctx *context.Context) error {
	return a.service.List(ctx)
}
