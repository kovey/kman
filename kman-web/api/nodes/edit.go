package nodes

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
	kow.POST("/admin/node", newEdit(services.Nodes())).Rule("node", "minlen:int:1", "maxlen:int:256").Rule("weight", "ge:int64:0").Data(services.NodesEditData()).Middleware(middlewares.NewAuth(), middlewares.NewCheckOpt())
}

type edit struct {
	*controller.Base
	service services.NodesInterface
}

func newEdit(service services.NodesInterface) *edit {
	return &edit{Base: controller.NewBase(os.Getenv("SERVICE_GROUP"), krpc.ServiceName(os.Getenv("SERVICE_NAME"))), service: service}
}

func (d *edit) Action(ctx *context.Context) error {
	return d.service.Edit(ctx)
}
