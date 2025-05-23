package configs

import (
	"os"

	"github.com/kovey/discovery/krpc"
	"github.com/kovey/kman/kman-web/module/libs/code"
	"github.com/kovey/kman/kman-web/module/libs/proto"
	"github.com/kovey/kow/context"
	"github.com/kovey/kow/result"
)

type Configs struct {
	servName string
	group    string
}

func NewConfigs() *Configs {
	return &Configs{servName: os.Getenv("SERVICE_NAME"), group: os.Getenv("SERVICE_GROUP")}
}

func (n *Configs) Add(ctx *context.Context) error {
	req := ctx.ReqData.(*AddReqData)
	conn := ctx.Rpcs.Get(krpc.ServiceName(n.servName), n.group)
	if conn == nil {
		return result.Err(ctx, code.System_Err, "conn not found")
	}

	cli := proto.NewConfigClient(conn)
	resp, err := cli.Add(ctx.Context, &proto.ConfigAddReq{Name: req.Name, Key: req.Key, Value: req.Value, Namespace: ctx.GetString("namespace"), ProjectId: int32(ctx.GetInt("projectId"))})
	if err != nil {
		return result.Convert(ctx, err)
	}

	return result.Succ(ctx, resp)
}

func (n *Configs) Release(ctx *context.Context) error {
	req := ctx.ReqData.(*ReleaseReqData)
	conn := ctx.Rpcs.Get(krpc.ServiceName(n.servName), n.group)
	if conn == nil {
		return result.Err(ctx, code.System_Err, "conn not found")
	}
	cli := proto.NewConfigClient(conn)
	resp, err := cli.Release(ctx.Context, &proto.ConfigReleaseReq{Ids: req.Ids, Namespace: ctx.GetString("namespace"), ProjectId: int32(ctx.GetInt("projectId"))})
	if err != nil {
		return result.Convert(ctx, err)
	}

	return result.Succ(ctx, resp)
}

func (n *Configs) Edit(ctx *context.Context) error {
	req := ctx.ReqData.(*EditReqData)
	conn := ctx.Rpcs.Get(krpc.ServiceName(n.servName), n.group)
	if conn == nil {
		return result.Err(ctx, code.System_Err, "conn not found")
	}
	cli := proto.NewConfigClient(conn)
	resp, err := cli.Edit(ctx.Context, &proto.ConfigEditReq{Id: req.Id, Value: req.Value, Namespace: ctx.GetString("namespace"), ProjectId: int32(ctx.GetInt("projectId"))})
	if err != nil {
		return result.Convert(ctx, err)
	}

	return result.Succ(ctx, resp)
}

func (n *Configs) List(ctx *context.Context) error {
	req := ctx.ReqData.(*ListReqData)
	conn := ctx.Rpcs.Get(krpc.ServiceName(n.servName), n.group)
	if conn == nil {
		return result.Err(ctx, code.System_Err, "conn not found")
	}
	cli := proto.NewConfigClient(conn)
	resp, err := cli.List(ctx.Context, &proto.ConfigListReq{Page: req.Page, PageSize: req.PageSize, Key: req.Key, Namespace: ctx.GetString("namespace"), ProjectId: int32(ctx.GetInt("projectId"))})
	if err != nil {
		return result.Convert(ctx, err)
	}

	respData := ListRespData{Page: 1, PageSize: req.PageSize, TotalPage: resp.TotalPage, TotalCount: resp.TotalCount, List: make([]ListInfo, len(resp.List))}
	for index, row := range resp.List {
		respData.List[index] = ListInfo{
			Id: row.Id, Name: row.Name, Key: row.Key, Value: row.Value, History1: row.History1, History2: row.History2,
			UpdateTime: row.UpdateTime, CreateTime: row.CreateTime, Status: row.Status,
		}
	}

	return result.Succ(ctx, respData)
}
