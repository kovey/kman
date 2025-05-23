package services

import (
	"github.com/kovey/kman/kman-web/module/services/configs"
	"github.com/kovey/kman/kman-web/module/services/login"
	"github.com/kovey/kman/kman-web/module/services/nodes"
	"github.com/kovey/kman/kman-web/module/services/operator"
	"github.com/kovey/kman/kman-web/module/services/project"
	"github.com/kovey/kow/validator/rule"
)

func LoginData() rule.ParamInterface {
	return &login.LoginReqData{}
}

func RefreshData() rule.ParamInterface {
	return &login.RefreshReqData{}
}

func OperatorAddData() rule.ParamInterface {
	return &operator.OperatorAddReqData{}
}

func OperatorEditData() rule.ParamInterface {
	return &operator.OperatorEditReqData{}
}

func OperatorListData() rule.ParamInterface {
	return &operator.OperatorListReqData{}
}

func NodesDeleteData() rule.ParamInterface {
	return &nodes.DeleteReqData{}
}

func NodesEditData() rule.ParamInterface {
	return &nodes.EditReqData{}
}

func NodesListData() rule.ParamInterface {
	return &nodes.ListReqData{}
}

func ProjectAddData() rule.ParamInterface {
	return &project.ProjectAddReqData{}
}

func ProjectEditData() rule.ParamInterface {
	return &project.ProjectEditReqData{}
}

func ProjectListData() rule.ParamInterface {
	return &project.ProjectListReqData{}
}

func ConfigAddData() rule.ParamInterface {
	return &configs.AddReqData{}
}

func ConfigEditData() rule.ParamInterface {
	return &configs.EditReqData{}
}

func ConfigListData() rule.ParamInterface {
	return &configs.ListReqData{}
}

func ConfigReleaseData() rule.ParamInterface {
	return &configs.ReleaseReqData{}
}
