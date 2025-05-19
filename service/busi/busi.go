package busi

import (
	"github.com/kovey/kman/service/busi/config"
	"github.com/kovey/kman/service/busi/login"
	"github.com/kovey/kman/service/busi/nodes"
	"github.com/kovey/kman/service/busi/operator"
	"github.com/kovey/kman/service/busi/project"
	"github.com/kovey/kman/service/module/libs/proto"
	"github.com/kovey/pool/object"
)

type ConfigInterface interface {
	Add(req *proto.ConfigAddReq) (*proto.ConfigAddResp, error)
	Edit(req *proto.ConfigEditReq) (*proto.ConfigEditResp, error)
	Release(req *proto.ConfigReleaseReq) (*proto.ConfigReleaseResp, error)
	List(req *proto.ConfigListReq) (*proto.ConfigListResp, error)
}

type LoginInterface interface {
	Login(*proto.LoginReq) (*proto.LoginResp, error)
}

type NodesInterface interface {
	Edit(req *proto.NodeEditReq) (*proto.NodeEditResp, error)
	Delete(req *proto.NodeDeleteReq) (*proto.NodeDeleteResp, error)
	List(req *proto.NodeListReq) (*proto.NodeListResp, error)
}

type OperatorInterface interface {
	Add(req *proto.OperatorAddReq) (*proto.OperatorAddResp, error)
	Edit(req *proto.OperatorEditReq) (*proto.OperatorEditResp, error)
	List(req *proto.OperatorListReq) (*proto.OperatorListResp, error)
}

type ProjectInterface interface {
	Add(req *proto.ProjectAddReq) (*proto.ProjectAddResp, error)
	Edit(req *proto.ProjectEditReq) (*proto.ProjectEditResp, error)
	List(req *proto.ProjectListReq) (*proto.ProjectListResp, error)
}

func Config(ctx object.CtxInterface) ConfigInterface {
	return config.NewConfig(ctx)
}

func Login(ctx object.CtxInterface) LoginInterface {
	return login.NewLogin(ctx)
}

func Operator(ctx object.CtxInterface) OperatorInterface {
	return operator.NewOperator(ctx)
}

func Nodes(ctx object.CtxInterface) NodesInterface {
	return nodes.NewNodes(ctx)
}

func Project(ctx object.CtxInterface) ProjectInterface {
	return project.NewProject(ctx)
}
