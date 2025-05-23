package login

import "github.com/kovey/kow/validator/rule"

type RefreshReqData struct {
	Token string `json:"token" form:"token"`
}

func (l *RefreshReqData) ValidParams() map[string]any {
	return map[string]any{
		"token": l.Token,
	}
}

func (l *RefreshReqData) Clone() rule.ParamInterface {
	return &RefreshReqData{}
}

type refreshTokenRespData struct {
	Token string `json:"token" form:"token"`
}
