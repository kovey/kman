package middlewares

import (
	"github.com/kovey/kman/kman-web/module/libs/code"
	"github.com/kovey/kow/context"
	"github.com/kovey/kow/result"
)

type CheckAdmin struct {
}

func NewCheckAdmin() *CheckAdmin {
	return &CheckAdmin{}
}

func (c *CheckAdmin) Handle(ctx *context.Context) {
	if ctx.GetInt("projectId") == 0 {
		ctx.Next()
		return
	}

	result.Err(ctx, code.No_Access, "No Access")
}
