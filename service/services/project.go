package services

import (
	"context"

	"github.com/kovey/kman/service/busi"
	"github.com/kovey/kman/service/module/libs/proto"
	"github.com/kovey/kom/service"
	"github.com/kovey/pool"
)

func init() {
	service.Register(NewProject())
}

type Project struct {
	*service.Base
	proto.UnimplementedProjectServer
}

func NewProject() *Project {
	return &Project{Base: service.NewBase(&proto.Project_ServiceDesc)}
}

func (p *Project) Add(ctx context.Context, req *proto.ProjectAddReq) (*proto.ProjectAddResp, error) {
	srv := busi.Project(ctx.(*pool.Context))
	return srv.Add(req)
}

func (p *Project) Edit(ctx context.Context, req *proto.ProjectEditReq) (*proto.ProjectEditResp, error) {
	srv := busi.Project(ctx.(*pool.Context))
	return srv.Edit(req)
}

func (p *Project) List(ctx context.Context, req *proto.ProjectListReq) (*proto.ProjectListResp, error) {
	srv := busi.Project(ctx.(*pool.Context))
	return srv.List(req)
}
