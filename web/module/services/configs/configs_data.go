package configs

import "github.com/kovey/kow/validator/rule"

type AddReqData struct {
	Name  string `json:"name" form:"name"`
	Value string `json:"value" form:"value"`
	Key   string `json:"key" form:"key"`
}

func (l *AddReqData) ValidParams() map[string]any {
	return map[string]any{
		"key": l.Key,
	}
}

func (l *AddReqData) Clone() rule.ParamInterface {
	return &AddReqData{}
}

type ReleaseReqData struct {
	Ids []int64 `json:"ids" form:"ids"`
}

func (l *ReleaseReqData) ValidParams() map[string]any {
	return map[string]any{
		"ids": l.Ids,
	}
}

func (l *ReleaseReqData) Clone() rule.ParamInterface {
	return &ReleaseReqData{}
}

type EditReqData struct {
	Value string `json:"value" form:"value"`
	Id    int64  `json:"id" form:"id"`
}

func (l *EditReqData) ValidParams() map[string]any {
	return map[string]any{
		"id":    l.Id,
		"value": l.Value,
	}
}

func (l *EditReqData) Clone() rule.ParamInterface {
	return &EditReqData{}
}

type ListReqData struct {
	Page     int64  `json:"page" form:"page"`
	PageSize int64  `json:"page_size" form:"page_size"`
	Key      string `json:"key" form:"key"`
}

func (l *ListReqData) ValidParams() map[string]any {
	return map[string]any{
		"key":       l.Key,
		"page":      l.Page,
		"page_size": l.PageSize,
	}
}

func (l *ListReqData) Clone() rule.ParamInterface {
	return &ListReqData{}
}

type ListInfo struct {
	Id         int64  `json:"id"`
	Key        string `json:"key"`
	Value      string `json:"value"`
	History1   string `json:"history1"`
	History2   string `json:"history2"`
	Name       string `json:"name"`
	UpdateTime string `json:"update_time"`
	CreateTime string `json:"create_time"`
	Status     string `json:"status"`
}

type ListRespData struct {
	Page       int64      `json:"page" form:"page"`
	PageSize   int64      `json:"page_size" form:"page_size"`
	TotalPage  int64      `json:"total_page" form:"total_page"`
	TotalCount int64      `json:"total_count" form:"total_count"`
	List       []ListInfo `json:"list"`
}
