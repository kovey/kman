package nodes

import (
	"os"

	"github.com/kovey/discovery/krpc"
	"github.com/kovey/kman/kman-web/module/libs/code"
	"github.com/kovey/kman/kman-web/module/libs/proto"
	"github.com/kovey/kow/context"
	"github.com/kovey/kow/result"
)

type Nodes struct {
	servName string
	group    string
}

func NewNodes() *Nodes {
	return &Nodes{servName: os.Getenv("SERVICE_NAME"), group: os.Getenv("SERVICE_GROUP")}
}

func (n *Nodes) Delete(ctx *context.Context) error {
	req := ctx.ReqData.(*DeleteReqData)
	conn := ctx.Rpcs.Get(krpc.ServiceName(n.servName), n.group)
	if conn == nil {
		return result.Err(ctx, code.System_Err, "conn not found")
	}

	cli := proto.NewNodeClient(conn)
	resp, err := cli.Delete(ctx.Context, &proto.NodeDeleteReq{Node: req.Node, Namespace: ctx.GetString("namespace"), ProjectId: int32(ctx.GetInt("projectId"))})
	if err != nil {
		return result.Convert(ctx, err)
	}

	return result.Succ(ctx, resp)
}

func (n *Nodes) Edit(ctx *context.Context) error {
	req := ctx.ReqData.(*EditReqData)
	conn := ctx.Rpcs.Get(krpc.ServiceName(n.servName), n.group)
	if conn == nil {
		return result.Err(ctx, code.System_Err, "conn not found")
	}

	cli := proto.NewNodeClient(conn)
	resp, err := cli.Edit(ctx.Context, &proto.NodeEditReq{Node: req.Node, Namespace: ctx.GetString("namespace"), ProjectId: int32(ctx.GetInt("projectId")), Weight: req.Weight})
	if err != nil {
		return result.Convert(ctx, err)
	}

	return result.Succ(ctx, resp)
}

func (n *Nodes) List(ctx *context.Context) error {
	req := ctx.ReqData.(*ListReqData)
	conn := ctx.Rpcs.Get(krpc.ServiceName(n.servName), n.group)
	if conn == nil {
		return result.Err(ctx, code.System_Err, "conn not found")
	}

	cli := proto.NewNodeClient(conn)
	resp, err := cli.List(ctx.Context, &proto.NodeListReq{Page: 1, PageSize: 200, Node: req.Node, Namespace: ctx.GetString("namespace"), ProjectId: int32(ctx.GetInt("projectId"))})
	if err != nil {
		return result.Convert(ctx, err)
	}

	respData := ListRespData{Page: resp.Page, PageSize: resp.PageSize, TotalPage: resp.TotalPage, TotalCount: resp.TotalCount, List: make([]ListInfo, len(resp.List))}
	for index, kv := range resp.List {
		respData.List[index] = ListInfo{
			Node: kv.Node, Name: kv.Name, Namespace: kv.Namespace, Group: kv.GroupName, Weight: kv.Weight, Version: kv.Version, Host: kv.Host, Port: kv.Port,
		}
	}

	return result.Succ(ctx, respData)
}
