package middlewares

import (
	"github.com/kovey/kman/web/module/libs/code"
	"github.com/kovey/kman/web/module/libs/token/access"
	"github.com/kovey/kow/context"
	"github.com/kovey/kow/jwt"
	"github.com/kovey/kow/result"
)

type Auth struct {
}

func NewAuth() *Auth {
	return &Auth{}
}

func (a *Auth) Handle(ctx *context.Context) {
	token := ctx.GetHeader("Access-Token")
	if token == "" {
		result.Err(ctx, code.Token_Err, "Token is empty")
		return
	}

	ext, err := access.Decode(token)
	if err != nil {
		if err == jwt.Err_Token_Expired {
			result.Err(ctx, code.Token_Expired, "Token is expired")
			return
		}

		result.Err(ctx, code.Token_Expired, "Token is expired")
		return
	}
	ctx.Set("userId", ext.UserId)
	ctx.Set("permissions", ext.Permissions)
	ctx.Set("projectId", ext.ProjectId)
	ctx.Set("namespace", ext.Namespace)
	ctx.Next()
}
