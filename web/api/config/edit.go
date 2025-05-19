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
	router := kow.POST("/admin/config", newEdit(services.Config())).Rule("id", "ge:int64:1")
	router.Rule("value", "minlen:int:1").Data(services.ConfigEditData())
	router.Middleware(middlewares.NewAuth(), middlewares.NewCheckOpt())
}

type edit struct {
	*controller.Base
	service services.ConfigInterface
}

func newEdit(service services.ConfigInterface) *edit {
	return &edit{Base: controller.NewBase(os.Getenv("SERVICE_GROUP"), krpc.ServiceName(os.Getenv("SERVICE_NAME"))), service: service}
}

func (a *edit) Action(ctx *context.Context) error {
	return a.service.Edit(ctx)
}
