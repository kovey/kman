package project

import (
	"os"

	"github.com/kovey/discovery/krpc"
	"github.com/kovey/kman/kman-web/module/libs/code"
	"github.com/kovey/kman/kman-web/module/libs/proto"
	"github.com/kovey/kow/context"
	"github.com/kovey/kow/result"
)

type Project struct {
	servName string
	group    string
}

func NewProject() *Project {
	return &Project{servName: os.Getenv("SERVICE_NAME"), group: os.Getenv("SERVICE_GROUP")}
}

func (o *Project) Add(ctx *context.Context) error {
	req := ctx.ReqData.(*ProjectAddReqData)
	conn := ctx.Rpcs.Get(krpc.ServiceName(o.servName), o.group)
	if conn == nil {
		return result.Err(ctx, code.System_Err, "conn not found")
	}

	cli := proto.NewProjectClient(conn)
	resp, err := cli.Add(ctx.Context, &proto.ProjectAddReq{Name: req.Name, Namespace: req.Namespace})
	if err != nil {
		return result.Convert(ctx, err)
	}

	return result.Succ(ctx, resp)
}

func (o *Project) Edit(ctx *context.Context) error {
	req := ctx.ReqData.(*ProjectEditReqData)
	conn := ctx.Rpcs.Get(krpc.ServiceName(o.servName), o.group)
	if conn == nil {
		return result.Err(ctx, code.System_Err, "conn not found")
	}

	cli := proto.NewProjectClient(conn)
	resp, err := cli.Edit(ctx.Context, &proto.ProjectEditReq{Id: int64(req.Id), Name: req.Name})
	if err != nil {
		return result.Convert(ctx, err)
	}

	return result.Succ(ctx, resp)
}

func (o *Project) List(ctx *context.Context) error {
	req := ctx.ReqData.(*ProjectListReqData)
	conn := ctx.Rpcs.Get(krpc.ServiceName(o.servName), o.group)
	if conn == nil {
		return result.Err(ctx, code.System_Err, "conn not found")
	}

	cli := proto.NewProjectClient(conn)
	respData, err := cli.List(ctx.Context, &proto.ProjectListReq{Page: req.Page, PageSize: req.PageSize})
	if err != nil {
		return result.Convert(ctx, err)
	}

	resp := ProjectListRespData{Page: req.Page, PageSize: req.PageSize, TotalPage: respData.TotalPage, TotalCount: respData.TotalCount}
	resp.List = make([]ProjectInfo, len(respData.List))
	for i, row := range respData.List {
		resp.List[i] = ProjectInfo{Id: int(row.Id), Name: row.Name, Namespace: row.Namespace, OpenId: row.OpenId, CreateTime: row.CreateTime}
	}

	return result.Succ(ctx, resp)
}
