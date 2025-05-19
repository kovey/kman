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
	router := kow.POST("/admin/project", newEdit(services.Project()))
	router.Rule("name", "minlen:int:1", "maxlen:int:127").Data(services.ProjectEditData())
	router.Middleware(middlewares.NewAuth(), middlewares.NewCheckAdmin())
}

type edit struct {
	*controller.Base
	service services.ProjectInterface
}

func newEdit(service services.ProjectInterface) *edit {
	return &edit{Base: controller.NewBase(os.Getenv("SERVICE_GROUP"), krpc.ServiceName(os.Getenv("SERVICE_NAME"))), service: service}
}

func (a *edit) Action(ctx *context.Context) error {
	return a.service.Edit(ctx)
}
