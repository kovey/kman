package login

import (
	"github.com/kovey/kman/web/module/libs/code"
	"github.com/kovey/kman/web/module/libs/token/access"
	"github.com/kovey/kman/web/module/libs/token/refresh"
	"github.com/kovey/kow/context"
	"github.com/kovey/kow/result"
)

type Refresh struct {
}

func (r *Refresh) Refresh(ctx *context.Context) error {
	req := ctx.ReqData.(*RefreshReqData)
	ext, err := refresh.Decode(req.Token)
	if err != nil {
		return result.Err(ctx, code.Token_Err, err.Error())
	}

	token, err := access.Token(ext)
	if err != nil {
		return result.Err(ctx, code.Token_Err, err.Error())
	}

	return result.Succ(ctx, refreshTokenRespData{Token: token})
}
