package nodes

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
	kow.GET("/admin/node", newList(services.Nodes())).Rule("node", "minlen:int:0", "maxlen:int:256").Data(services.NodesListData()).Middleware(middlewares.NewAuth(), middlewares.NewCheckOpt())
}

type list struct {
	*controller.Base
	service services.NodesInterface
}

func newList(service services.NodesInterface) *list {
	return &list{Base: controller.NewBase(os.Getenv("SERVICE_GROUP"), krpc.ServiceName(os.Getenv("SERVICE_NAME"))), service: service}
}

func (d *list) Action(ctx *context.Context) error {
	return d.service.List(ctx)
}
