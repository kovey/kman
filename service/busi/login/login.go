package login

import (
	"encoding/json"
	"fmt"
	"strings"

	ksql "github.com/kovey/db-go/v3"
	"github.com/kovey/db-go/v3/db"
	"github.com/kovey/kman/service/module/libs/code"
	"github.com/kovey/kman/service/module/libs/proto"
	"github.com/kovey/kman/service/module/models"
	"github.com/kovey/pool"
	"github.com/kovey/pool/object"
)

const (
	ctx_namespace = "busi.login"
	ctx_login     = "Login"
)

func init() {
	pool.Default(ctx_namespace, ctx_login, func() any {
		return &Login{Object: object.NewObject(ctx_namespace, ctx_login)}
	})
}

type Login struct {
	*object.Object
}

func NewLogin(ctx object.CtxInterface) *Login {
	return ctx.Get(ctx_namespace, ctx_login).(*Login)
}

func (l *Login) Login(req *proto.LoginReq) (*proto.LoginResp, error) {
	info := strings.Split(req.Username, "@")
	var project = models.NewProjects()
	if err := db.Model(project).Where(models.Table_Projects_Namespace, ksql.Eq, info[1]).First(l.Context); err != nil {
		return code.SystemErr[*proto.LoginResp](err)
	}
	if project.Empty() {
		return code.Err[*proto.LoginResp](code.Project_Not_Found, fmt.Errorf("Project not found"), nil)
	}

	var opInfo = models.NewOperator()
	err := db.Model(opInfo).Where(models.Table_Operator_Username, ksql.Eq, info[0]).Where(models.Table_Operator_ProjectId, ksql.Eq, project.Id).First(l.Context)
	if err != nil {
		return code.SystemErr[*proto.LoginResp](err)
	}

	if opInfo.Empty() || opInfo.ProjectId != project.Id {
		return code.Err[*proto.LoginResp](code.Admin_Not_Found, fmt.Errorf("Admin not found"), nil)
	}

	resp := &proto.LoginResp{UserId: opInfo.Id, ProjectId: int32(project.Id), Namespace: project.Namespace, Password: opInfo.Password}
	json.Unmarshal([]byte(opInfo.Permissions), &resp.Permissions)
	return code.Succ(resp)
}
