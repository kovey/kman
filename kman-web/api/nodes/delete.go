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
	kow.DELETE("/admin/node", newDel(services.Nodes())).Rule("node", "minlen:int:1", "maxlen:int:256").Data(services.NodesDeleteData()).Middleware(middlewares.NewAuth(), middlewares.NewCheckOpt())
}

type del struct {
	*controller.Base
	service services.NodesInterface
}

func newDel(service services.NodesInterface) *del {
	return &del{Base: controller.NewBase(os.Getenv("SERVICE_GROUP"), krpc.ServiceName(os.Getenv("SERVICE_NAME"))), service: service}
}

func (d *del) Action(ctx *context.Context) error {
	return d.service.Delete(ctx)
}
