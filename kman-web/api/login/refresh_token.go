package login

import (
	"os"

	"github.com/kovey/kman/kman-web/module/services"
	"github.com/kovey/kow"
	"github.com/kovey/kow/context"
	"github.com/kovey/kow/controller"
)

func init() {
	kow.POST("/admin/refresh", newRefresh(services.Refresh())).Rule("token", "jwt").Data(services.RefreshData())
}

type refreshToken struct {
	*controller.Base
	service services.RefreshInterface
}

func newRefresh(service services.RefreshInterface) *refreshToken {
	return &refreshToken{Base: controller.NewBase(os.Getenv("SERVICE_GROUP")), service: service}
}

func (r *refreshToken) Action(ctx *context.Context) error {
	return r.service.Refresh(ctx)
}
