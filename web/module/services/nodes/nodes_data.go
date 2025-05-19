package nodes

import "github.com/kovey/kow/validator/rule"

type DeleteReqData struct {
	Node string `json:"node" form:"node"`
}

func (l *DeleteReqData) ValidParams() map[string]any {
	return map[string]any{
		"node": l.Node,
	}
}

func (l *DeleteReqData) Clone() rule.ParamInterface {
	return &DeleteReqData{}
}

type EditReqData struct {
	Node   string `json:"node" form:"node"`
	Weight int64  `json:"weight" form:"weight"`
}

func (l *EditReqData) ValidParams() map[string]any {
	return map[string]any{
		"node":   l.Node,
		"weight": l.Weight,
	}
}

func (l *EditReqData) Clone() rule.ParamInterface {
	return &EditReqData{}
}

type ListReqData struct {
	Node string `json:"node" form:"node"`
}

func (l *ListReqData) ValidParams() map[string]any {
	return map[string]any{
		"node": l.Node,
	}
}

func (l *ListReqData) Clone() rule.ParamInterface {
	return &ListReqData{}
}

type ListInfo struct {
	Node      string `json:"node"`
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Group     string `json:"group"`
	Weight    int64  `json:"weight"`
	Version   string `json:"version"`
	Host      string `json:"host"`
	Port      string `json:"port"`
}

type ListRespData struct {
	Page       int64      `json:"page" form:"page"`
	PageSize   int64      `json:"page_size" form:"page_size"`
	TotalPage  int64      `json:"total_page" form:"total_page"`
	TotalCount int64      `json:"total_count" form:"total_count"`
	List       []ListInfo `json:"list"`
}
