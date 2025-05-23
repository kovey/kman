package operator

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
	router := kow.POST("/admin/operator", newEdit(services.Operator())).Rule("account", "minlen:int:5", "maxlen:int:20", "regx:string:^[a-zA-Z0-9][a-zA-Z0-9_]+$")
	router.Rule("password", "len:int:64").Rule("id", "ge:int64:1").Data(services.OperatorEditData())
	router.Middleware(middlewares.NewAuth())
}

type edit struct {
	*controller.Base
	service services.OperatorInterface
}

func newEdit(service services.OperatorInterface) *edit {
	return &edit{Base: controller.NewBase(os.Getenv("SERVICE_GROUP"), krpc.ServiceName(os.Getenv("SERVICE_NAME"))), service: service}
}

func (a *edit) Action(ctx *context.Context) error {
	return a.service.Edit(ctx)
}
