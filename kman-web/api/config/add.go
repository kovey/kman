package config

import (
	"os"

	"github.com/kovey/discovery/krpc"
	"github.com/kovey/kman/kman-web/module/libs/middlewares"
	"github.com/kovey/kman/kman-web/module/services"
	"github.com/kovey/kow"
	"github.com/kovey/kow/context"
	"github.com/kovey/kow/controller"
)

func init() {
	router := kow.PUT("/admin/config", newAdd(services.Config())).Rule("key", "minlen:int:1", "maxlen:int:255", "regx:string:^[a-zA-Z0-9][a-zA-Z0-9_\\.]+$").Data(services.ConfigAddData())
	router.Rule("name", "minlen:int:1", "maxlen:int:127").Rule("value", "minlen:int:1")
	router.Middleware(middlewares.NewAuth(), middlewares.NewCheckOpt())
}

type add struct {
	*controller.Base
	service services.ConfigInterface
}

func newAdd(service services.ConfigInterface) *add {
	return &add{Base: controller.NewBase(os.Getenv("SERVICE_GROUP"), krpc.ServiceName(os.Getenv("SERVICE_NAME"))), service: service}
}

func (a *add) Action(ctx *context.Context) error {
	return a.service.Add(ctx)
}
