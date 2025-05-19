package login

import (
	"os"

	"github.com/kovey/discovery/krpc"
	"github.com/kovey/kman/web/module/services"
	"github.com/kovey/kow"
	"github.com/kovey/kow/context"
	"github.com/kovey/kow/controller"
)

func init() {
	kow.GET("/admin/login", newLogin(services.Login())).Rule("username", "maxlen:int:20", "minlen:int:5", "regx:string:^[a-zA-Z0-9][a-zA-Z0-9_]+@[a-zA-Z0-9_]+$").Rule("password", "len:int:64").Data(services.LoginData())
}

type loginAction struct {
	*controller.Base
	service services.LoginInterface
}

func newLogin(service services.LoginInterface) *loginAction {
	return &loginAction{Base: controller.NewBase(os.Getenv("SERVICE_GROUP"), krpc.ServiceName(os.Getenv("SERVICE_NAME"))), service: service}
}

func (l *loginAction) Action(ctx *context.Context) error {
	return l.service.Login(ctx)
}
