package operator

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
	router := kow.GET("/admin/operator", newList(services.Operator())).Rule("account", "minlen:int:5", "maxlen:int:20", "regx:string:^[a-zA-Z0-9][a-zA-Z0-9_]+$")
	router.Rule("page", "ge:int64:1").Data(services.OperatorListData())
	router.Rule("page_size", "ge:int64:1")
	router.Middleware(middlewares.NewAuth())
}

type list struct {
	*controller.Base
	service services.OperatorInterface
}

func newList(service services.OperatorInterface) *list {
	return &list{Base: controller.NewBase(os.Getenv("SERVICE_GROUP"), krpc.ServiceName(os.Getenv("SERVICE_NAME"))), service: service}
}

func (a *list) Action(ctx *context.Context) error {
	return a.service.List(ctx)
}
