package services

import (
	"context"

	"github.com/kovey/kman/service/busi"
	"github.com/kovey/kman/service/module/libs/proto"
	"github.com/kovey/kom/service"
	"github.com/kovey/pool"
)

func init() {
	service.Register(NewConfig())
}

type Config struct {
	*service.Base
	proto.UnimplementedConfigServer
}

func NewConfig() *Config {
	return &Config{Base: service.NewBase(&proto.Config_ServiceDesc)}
}

func (c *Config) Add(ctx context.Context, req *proto.ConfigAddReq) (*proto.ConfigAddResp, error) {
	serv := busi.Config(ctx.(*pool.Context))
	return serv.Add(req)
}
func (c *Config) Edit(ctx context.Context, req *proto.ConfigEditReq) (*proto.ConfigEditResp, error) {
	serv := busi.Config(ctx.(*pool.Context))
	return serv.Edit(req)
}
func (c *Config) Release(ctx context.Context, req *proto.ConfigReleaseReq) (*proto.ConfigReleaseResp, error) {
	serv := busi.Config(ctx.(*pool.Context))
	return serv.Release(req)
}
func (c *Config) List(ctx context.Context, req *proto.ConfigListReq) (*proto.ConfigListResp, error) {
	serv := busi.Config(ctx.(*pool.Context))
	return serv.List(req)
}
