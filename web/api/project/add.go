package project

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
	router := kow.PUT("/admin/project", newAdd(services.Project())).Rule("namespace", "minlen:int:5", "maxlen:int:20", "regx:string:^[a-zA-Z0-9][a-zA-Z0-9_]+$").Data(services.ProjectAddData())
	router.Rule("name", "minlen:int:1", "maxlen:int:127")
	router.Middleware(middlewares.NewAuth(), middlewares.NewCheckAdmin())
}

type add struct {
	*controller.Base
	service services.ProjectInterface
}

func newAdd(service services.ProjectInterface) *add {
	return &add{Base: controller.NewBase(os.Getenv("SERVICE_GROUP"), krpc.ServiceName(os.Getenv("SERVICE_NAME"))), service: service}
}

func (a *add) Action(ctx *context.Context) error {
	return a.service.Add(ctx)
}
