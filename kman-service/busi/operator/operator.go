package operator

import (
	"encoding/json"
	"fmt"
	"time"

	ksql "github.com/kovey/db-go/v3"
	"github.com/kovey/db-go/v3/db"
	"github.com/kovey/kman/kman-service/module/libs/code"
	"github.com/kovey/kman/kman-service/module/libs/proto"
	"github.com/kovey/kman/kman-service/module/models"
	"github.com/kovey/pool"
	"github.com/kovey/pool/object"
)

const (
	ctx_namespace = "busi.operator"
	ctx_operator  = "Operator"
)

func init() {
	pool.Default(ctx_namespace, ctx_operator, func() any {
		return &Operator{Object: object.NewObject(ctx_namespace, ctx_operator)}
	})
}

type Operator struct {
	*object.Object
}

func NewOperator(ctx object.CtxInterface) *Operator {
	return ctx.Get(ctx_namespace, ctx_operator).(*Operator)
}

func (o *Operator) Add(req *proto.OperatorAddReq) (*proto.OperatorAddResp, error) {
	builder := db.Model(models.NewOperator()).Where("username", ksql.Eq, req.Account)
	if req.ProjectId > 0 {
		builder.Where(models.Table_Operator_ProjectId, ksql.Eq, req.ProjectId)
	}
	ok, err := builder.Exist(o.Context)
	if err != nil {
		return code.SystemErr[*proto.OperatorAddResp](err)
	}

	if ok {
		return code.Err[*proto.OperatorAddResp](code.Admin_Exists, err, nil)
	}

	admin := models.NewOperator()
	admin.Username = req.Account
	admin.Password = req.Password
	admin.CreateTime = time.Now().Unix()
	admin.UpdateTime = admin.CreateTime
	admin.ProjectId = int(req.ProjectId)
	admin.Permissions = "[]"
	if err := admin.Save(o.Context); err != nil {
		return code.SystemErr[*proto.OperatorAddResp](err)
	}

	return code.Succ(&proto.OperatorAddResp{})
}

func (o *Operator) Edit(req *proto.OperatorEditReq) (*proto.OperatorEditResp, error) {
	admin := models.NewOperator()
	err := db.Find(o.Context, admin, req.Id)
	if err != nil {
		return code.SystemErr[*proto.OperatorEditResp](err)
	}

	if admin.Empty() || (req.ProjectId > 0 && admin.ProjectId != int(req.ProjectId)) {
		return code.Err[*proto.OperatorEditResp](code.Admin_Not_Found, fmt.Errorf("Admin Not Found"), nil)
	}
	if admin.Username != req.Account {
		ok, err := db.Model(models.NewOperator()).Where("username", ksql.Eq, req.Account).Exist(o.Context)
		if err != nil {
			return code.SystemErr[*proto.OperatorEditResp](err)
		}

		if ok {
			return code.Err[*proto.OperatorEditResp](code.Admin_Exists, fmt.Errorf("Admin is exists"), nil)
		}
	}

	admin.Username = req.Account
	admin.Password = req.Password
	admin.UpdateTime = time.Now().Unix()
	if err := admin.Save(o.Context); err != nil {
		return code.SystemErr[*proto.OperatorEditResp](err)
	}

	return code.Succ(&proto.OperatorEditResp{})
}

func (o *Operator) List(req *proto.OperatorListReq) (*proto.OperatorListResp, error) {
	builder := db.Models(&[]*models.Operator{}).OrderDesc("id")
	if req.ProjectId > 0 {
		builder.Where(models.Table_Configs_ProjectId, ksql.Eq, req.ProjectId)
	}
	if req.Account != "" {
		builder.Where("username", ksql.Eq, req.Account)
	}

	pageInfo, err := builder.Pagination(o.Context, req.Page, req.PageSize)
	if err != nil {
		return code.SystemErr[*proto.OperatorListResp](err)
	}

	resp := &proto.OperatorListResp{Page: req.Page, PageSize: req.PageSize, TotalPage: int64(pageInfo.TotalPage()), TotalCount: int64(pageInfo.TotalCount())}
	resp.List = make([]*proto.OperatorInfo, len(pageInfo.List()))
	for i, row := range pageInfo.List() {
		var permissions []int64
		json.Unmarshal([]byte(row.Permissions), &permissions)
		resp.List[i] = &proto.OperatorInfo{Id: row.Id, Username: row.Username, Permissions: permissions, CreateTime: time.Unix(row.CreateTime, 0).Format(time.DateTime)}
	}

	return code.Succ(resp)
}
