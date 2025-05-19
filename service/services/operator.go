package services

import (
	"context"

	"github.com/kovey/kman/service/busi"
	"github.com/kovey/kman/service/module/libs/proto"
	"github.com/kovey/kom/service"
	"github.com/kovey/pool"
)

func init() {
	service.Register(NewOperator())
}

type Operator struct {
	*service.Base
	proto.UnimplementedOperatorServer
}

func NewOperator() *Operator {
	return &Operator{Base: service.NewBase(&proto.Operator_ServiceDesc)}
}

func (o *Operator) Add(ctx context.Context, req *proto.OperatorAddReq) (*proto.OperatorAddResp, error) {
	serv := busi.Operator(ctx.(*pool.Context))
	return serv.Add(req)
}

func (o *Operator) Edit(ctx context.Context, req *proto.OperatorEditReq) (*proto.OperatorEditResp, error) {
	serv := busi.Operator(ctx.(*pool.Context))
	return serv.Edit(req)
}

func (o *Operator) List(ctx context.Context, req *proto.OperatorListReq) (*proto.OperatorListResp, error) {
	serv := busi.Operator(ctx.(*pool.Context))
	return serv.List(req)
}
