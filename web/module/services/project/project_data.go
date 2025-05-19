package project

import "github.com/kovey/kow/validator/rule"

type ProjectAddReqData struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

func (l *ProjectAddReqData) ValidParams() map[string]any {
	return map[string]any{
		"name":      l.Name,
		"namespace": l.Namespace,
	}
}

func (l *ProjectAddReqData) Clone() rule.ParamInterface {
	return &ProjectAddReqData{}
}

type ProjectEditReqData struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (l *ProjectEditReqData) ValidParams() map[string]any {
	return map[string]any{
		"id":   l.Id,
		"name": l.Name,
	}
}

func (l *ProjectEditReqData) Clone() rule.ParamInterface {
	return &ProjectEditReqData{}
}

type ProjectListReqData struct {
	Name     string `json:"name" form:"name"`
	Page     int64  `json:"page" form:"page"`
	PageSize int64  `json:"page_size" form:"page_size"`
}

func (l *ProjectListReqData) ValidParams() map[string]any {
	return map[string]any{
		"name":      l.Name,
		"page":      l.Page,
		"page_size": l.PageSize,
	}
}

func (l *ProjectListReqData) Clone() rule.ParamInterface {
	return &ProjectListReqData{}
}

type ProjectInfo struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Namespace  string `json:"namespace"`
	OpenId     string `json:"open_id"`
	CreateTime string `json:"create_time"`
}

type ProjectListRespData struct {
	Page       int64         `json:"page" form:"page"`
	PageSize   int64         `json:"page_size" form:"page_size"`
	TotalPage  int64         `json:"total_page" form:"total_page"`
	TotalCount int64         `json:"total_count" form:"total_count"`
	List       []ProjectInfo `json:"list"`
}
