package config

import (
	"os"

	"github.com/kovey/discovery/krpc"
	"github.com/kovey/kman/web/module/libs/middlewares"
	"github.com/kovey/kman/web/module/services"
	"github.com/kovey/kow"
	"github.com/kovey/kow/context"
	"github.com/kovey/kow/controller"
	"github.com/kovey/kow/result"
)

func init() {
	router := kow.POST("/admin/config/release", newRelease(services.Config())).Data(services.ConfigReleaseData())
	router.Middleware(middlewares.NewAuth(), middlewares.NewCheckOpt())
}

type release struct {
	*controller.Base
	service services.ConfigInterface
}

func newRelease(service services.ConfigInterface) *release {
	return &release{Base: controller.NewBase(os.Getenv("SERVICE_GROUP"), krpc.ServiceName(os.Getenv("SERVICE_NAME"))), service: service}
}

func (a *release) Action(ctx *context.Context) error {
	ids, ok := ctx.ReqData.ValidParams()["ids"].([]int64)
	if !ok || len(ids) == 0 {
		return result.Err(ctx, result.Codes_Invalid_Params, "Invalid Params")
	}

	return a.service.Release(ctx)
}
