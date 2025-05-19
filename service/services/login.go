package services

import (
	"context"

	"github.com/kovey/kman/service/busi"
	"github.com/kovey/kman/service/module/libs/proto"
	"github.com/kovey/kom/service"
	"github.com/kovey/pool"
)

func init() {
	service.Register(NewLogin())
}

func NewLogin() *Login {
	return &Login{Base: service.NewBase(&proto.Login_ServiceDesc)}
}

type Login struct {
	*service.Base
	proto.UnimplementedLoginServer
}

func (l *Login) Login(ctx context.Context, req *proto.LoginReq) (*proto.LoginResp, error) {
	serv := busi.Login(ctx.(*pool.Context))
	return serv.Login(req)
}
