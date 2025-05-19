package operator

import "github.com/kovey/kow/validator/rule"

type OperatorAddReqData struct {
	Account   string `json:"account"`
	Password  string `json:"password"`
	ProjectId int    `json:"project_id"`
}

func (l *OperatorAddReqData) ValidParams() map[string]any {
	return map[string]any{
		"account":    l.Account,
		"password":   l.Password,
		"project_id": l.ProjectId,
	}
}

func (l *OperatorAddReqData) Clone() rule.ParamInterface {
	return &OperatorAddReqData{}
}

type OperatorEditReqData struct {
	Id       int64  `json:"id"`
	Account  string `json:"account"`
	Password string `json:"password"`
}

func (l *OperatorEditReqData) ValidParams() map[string]any {
	return map[string]any{
		"id":       l.Id,
		"account":  l.Account,
		"password": l.Password,
	}
}

func (l *OperatorEditReqData) Clone() rule.ParamInterface {
	return &OperatorEditReqData{}
}

type OperatorListReqData struct {
	Account  string `json:"account" form:"account"`
	Page     int64  `json:"page" form:"page"`
	PageSize int64  `json:"page_size" form:"page_size"`
}

func (l *OperatorListReqData) ValidParams() map[string]any {
	return map[string]any{
		"account":   l.Account,
		"page":      l.Page,
		"page_size": l.PageSize,
	}
}

func (l *OperatorListReqData) Clone() rule.ParamInterface {
	return &OperatorListReqData{}
}

type OperatorInfo struct {
	Id          int64   `json:"id"`
	Username    string  `json:"username"`
	Permissions []int64 `json:"permissions"`
	CreateTime  string  `json:"create_time"`
}

type OperatorListRespData struct {
	Page       int64          `json:"page" form:"page"`
	PageSize   int64          `json:"page_size" form:"page_size"`
	TotalPage  int64          `json:"total_page" form:"total_page"`
	TotalCount int64          `json:"total_count" form:"total_count"`
	List       []OperatorInfo `json:"list"`
}
