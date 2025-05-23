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
	router := kow.PUT("/admin/operator", newAdd(services.Operator())).Rule("account", "minlen:int:5", "maxlen:int:20", "regx:string:^[a-zA-Z0-9][a-zA-Z0-9_]+$").Rule("password", "len:int:64").Data(services.OperatorAddData())
	router.Middleware(middlewares.NewAuth())
}

type add struct {
	*controller.Base
	service services.OperatorInterface
}

func newAdd(service services.OperatorInterface) *add {
	return &add{Base: controller.NewBase(os.Getenv("SERVICE_GROUP"), krpc.ServiceName(os.Getenv("SERVICE_NAME"))), service: service}
}

func (a *add) Action(ctx *context.Context) error {
	return a.service.Add(ctx)
}
