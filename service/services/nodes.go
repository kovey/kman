package services

import (
	"context"

	"github.com/kovey/kman/service/busi"
	"github.com/kovey/kman/service/module/libs/proto"
	"github.com/kovey/kom/service"
	"github.com/kovey/pool"
)

func init() {
	service.Register(NewNodes())
}

type Nodes struct {
	*service.Base
	proto.UnimplementedNodeServer
}

func NewNodes() *Nodes {
	return &Nodes{Base: service.NewBase(&proto.Node_ServiceDesc)}
}

func (n *Nodes) Edit(ctx context.Context, req *proto.NodeEditReq) (*proto.NodeEditResp, error) {
	srv := busi.Nodes(ctx.(*pool.Context))
	return srv.Edit(req)
}

func (n *Nodes) Delete(ctx context.Context, req *proto.NodeDeleteReq) (*proto.NodeDeleteResp, error) {
	srv := busi.Nodes(ctx.(*pool.Context))
	return srv.Delete(req)
}

func (n *Nodes) List(ctx context.Context, req *proto.NodeListReq) (*proto.NodeListResp, error) {
	srv := busi.Nodes(ctx.(*pool.Context))
	return srv.List(req)
}
