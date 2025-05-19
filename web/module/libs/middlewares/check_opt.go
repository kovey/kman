package middlewares

import (
	"github.com/kovey/kman/web/module/libs/code"
	"github.com/kovey/kow/context"
	"github.com/kovey/kow/result"
)

type CheckOpt struct {
}

func NewCheckOpt() *CheckOpt {
	return &CheckOpt{}
}

func (c *CheckOpt) Handle(ctx *context.Context) {
	if ctx.GetInt("projectId") > 0 {
		ctx.Next()
		return
	}

	result.Err(ctx, code.No_Access, "No Access")
}
