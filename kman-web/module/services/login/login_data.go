package login

import "github.com/kovey/kow/validator/rule"

type LoginReqData struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func (l *LoginReqData) ValidParams() map[string]any {
	return map[string]any{
		"username": l.Username,
		"password": l.Password,
	}
}

func (l *LoginReqData) Clone() rule.ParamInterface {
	return &LoginReqData{}
}

type LoginRespData struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

func (l *LoginRespData) ValidParams() map[string]any {
	return map[string]any{
		"token":         l.Token,
		"refresh_token": l.RefreshToken,
	}
}

func (l *LoginRespData) Clone() rule.ParamInterface {
	return &LoginRespData{}
}
