package operator

import (
	"os"

	"github.com/kovey/discovery/krpc"
	"github.com/kovey/kman/web/module/libs/code"
	"github.com/kovey/kman/web/module/libs/password"
	"github.com/kovey/kman/web/module/libs/proto"
	"github.com/kovey/kow/context"
	"github.com/kovey/kow/result"
)

type Operator struct {
	servName string
	group    string
}

func NewOperator() *Operator {
	return &Operator{servName: os.Getenv("SERVICE_NAME"), group: os.Getenv("SERVICE_GROUP")}
}

func (o *Operator) Add(ctx *context.Context) error {
	req := ctx.ReqData.(*OperatorAddReqData)
	pId := ctx.GetInt("projectId")
	if pId == 0 && req.ProjectId == 0 {
		return result.Err(ctx, result.Codes_Invalid_Params, "Invalid Params")
	}
	if pId == 0 {
		pId = req.ProjectId
	}

	conn := ctx.Rpcs.Get(krpc.ServiceName(o.servName), o.group)
	if conn == nil {
		return result.Err(ctx, code.System_Err, "conn not found")
	}

	cli := proto.NewOperatorClient(conn)
	resp, err := cli.Add(ctx.Context, &proto.OperatorAddReq{Account: req.Account, Password: password.Password(req.Account, req.Password), Namespace: ctx.GetString("namespace"), ProjectId: int32(pId)})
	if err != nil {
		return result.Convert(ctx, err)
	}

	return result.Succ(ctx, resp)
}

func (o *Operator) Edit(ctx *context.Context) error {
	req := ctx.ReqData.(*OperatorEditReqData)
	conn := ctx.Rpcs.Get(krpc.ServiceName(o.servName), o.group)
	if conn == nil {
		return result.Err(ctx, code.System_Err, "conn not found")
	}

	cli := proto.NewOperatorClient(conn)
	resp, err := cli.Edit(ctx.Context, &proto.OperatorEditReq{Id: req.Id, Account: req.Account, Password: password.Password(req.Account, req.Password), Namespace: ctx.GetString("namespace"), ProjectId: int32(ctx.GetInt("projectId"))})
	if err != nil {
		return result.Convert(ctx, err)
	}

	return result.Succ(ctx, resp)
}

func (o *Operator) List(ctx *context.Context) error {
	req := ctx.ReqData.(*OperatorListReqData)
	conn := ctx.Rpcs.Get(krpc.ServiceName(o.servName), o.group)
	if conn == nil {
		return result.Err(ctx, code.System_Err, "conn not found")
	}

	cli := proto.NewOperatorClient(conn)
	resp, err := cli.List(ctx.Context, &proto.OperatorListReq{Page: req.Page, PageSize: req.PageSize, Account: req.Account, Namespace: ctx.GetString("namespace"), ProjectId: int32(ctx.GetInt("projectId"))})
	if err != nil {
		return result.Convert(ctx, err)
	}

	respData := OperatorListRespData{Page: req.Page, PageSize: req.PageSize, TotalPage: resp.TotalPage, TotalCount: resp.TotalCount}
	respData.List = make([]OperatorInfo, len(resp.List))
	for i, row := range resp.List {
		respData.List[i] = OperatorInfo{Id: row.Id, Username: row.Username, Permissions: row.Permissions, CreateTime: row.CreateTime}
	}

	return result.Succ(ctx, respData)
}
