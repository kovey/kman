package services

import (
	"github.com/kovey/kman/web/module/services/configs"
	"github.com/kovey/kman/web/module/services/login"
	"github.com/kovey/kman/web/module/services/nodes"
	"github.com/kovey/kman/web/module/services/operator"
	"github.com/kovey/kman/web/module/services/project"
	"github.com/kovey/kow/context"
)

type LoginInterface interface {
	Login(ctx *context.Context) error
}

type RefreshInterface interface {
	Refresh(ctx *context.Context) error
}

type OperatorInterface interface {
	Add(ctx *context.Context) error
	Edit(ctx *context.Context) error
	List(ctx *context.Context) error
}

type NodesInterface interface {
	Delete(ctx *context.Context) error
	List(ctx *context.Context) error
	Edit(ctx *context.Context) error
}

type ProjectInterface interface {
	Add(ctx *context.Context) error
	Edit(ctx *context.Context) error
	List(ctx *context.Context) error
}

type ConfigInterface interface {
	Add(ctx *context.Context) error
	Edit(ctx *context.Context) error
	List(ctx *context.Context) error
	Release(ctx *context.Context) error
}

func Login() LoginInterface {
	return login.NewLogin()
}

func Refresh() RefreshInterface {
	return &login.Refresh{}
}

func Operator() OperatorInterface {
	return &operator.Operator{}
}

func Nodes() NodesInterface {
	return nodes.NewNodes()
}

func Project() ProjectInterface {
	return &project.Project{}
}

func Config() ConfigInterface {
	return configs.NewConfigs()
}
