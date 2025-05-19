package nodes

import (
	"fmt"
	"strings"

	"github.com/kovey/discovery/grpc"
	"github.com/kovey/kman/service/module/libs/code"
	"github.com/kovey/kman/service/module/libs/etcd"
	"github.com/kovey/kman/service/module/libs/proto"
	"github.com/kovey/pool"
	"github.com/kovey/pool/object"
	clientv3 "go.etcd.io/etcd/client/v3"
)

const (
	ctx_namespace = "busi.nodes"
	ctx_nodes     = "Nodes"
)

func init() {
	pool.Default(ctx_namespace, ctx_nodes, func() any {
		return &Nodes{Object: object.NewObject(ctx_namespace, ctx_nodes)}
	})
}

type Nodes struct {
	*object.Object
}

func NewNodes(ctx object.CtxInterface) *Nodes {
	return ctx.Get(ctx_namespace, ctx_nodes).(*Nodes)
}

func (n *Nodes) prefix(namespace string) string {
	return fmt.Sprintf("/ko/grpc/%s", namespace)
}

func (n *Nodes) Edit(req *proto.NodeEditReq) (*proto.NodeEditResp, error) {
	resp, err := etcd.Get(n.Context, fmt.Sprintf("%s/%s", n.prefix(req.Namespace), req.Node))
	if err != nil {
		return code.SystemErr[*proto.NodeEditResp](err)
	}

	if resp.Count != 1 {
		return code.Err[*proto.NodeEditResp](code.Node_Not_Found, fmt.Errorf("Node not found"), nil)
	}

	value := string(resp.Kvs[0].Value)
	ins := &grpc.Instance{}
	if err := ins.Decode(value); err != nil {
		return code.SystemErr[*proto.NodeEditResp](err)
	}

	if ins.Namespace != req.Namespace {
		return code.Err[*proto.NodeEditResp](code.Node_Not_Found, fmt.Errorf("Node not found"), nil)
	}

	ins.Weight = req.Weight
	buf, err := ins.Encode()
	_, err = etcd.Put(n.Context, fmt.Sprintf("%s/%s", n.prefix(req.Namespace), req.Node), buf)
	if err != nil {
		return code.SystemErr[*proto.NodeEditResp](err)
	}

	return code.Succ(&proto.NodeEditResp{})
}

func (n *Nodes) Delete(req *proto.NodeDeleteReq) (*proto.NodeDeleteResp, error) {
	resp, err := etcd.Get(n.Context, fmt.Sprintf("%s/%s", n.prefix(req.Namespace), req.Node))
	if err != nil {
		return code.SystemErr[*proto.NodeDeleteResp](err)
	}

	if resp.Count != 1 {
		return code.Err[*proto.NodeDeleteResp](code.Node_Not_Found, fmt.Errorf("Node not found"), nil)
	}

	value := string(resp.Kvs[0].Value)
	ins := &grpc.Instance{}
	if err := ins.Decode(value); err != nil {
		return code.SystemErr[*proto.NodeDeleteResp](err)
	}
	if ins.Namespace != req.Namespace {
		return code.Err[*proto.NodeDeleteResp](code.Node_Not_Found, fmt.Errorf("Node not found"), nil)
	}

	_, err = etcd.Delete(n.Context, fmt.Sprintf("%s/%s", n.prefix(req.Namespace), req.Node))
	if err != nil {
		return code.SystemErr[*proto.NodeDeleteResp](err)
	}

	return code.Succ(&proto.NodeDeleteResp{})
}

func (n *Nodes) List(req *proto.NodeListReq) (*proto.NodeListResp, error) {
	key := n.prefix(req.Namespace)
	if req.Node != "" {
		key = fmt.Sprintf("%s/%s", key, req.Node)
	}

	resp, err := etcd.Get(n.Context, key, clientv3.WithPrefix())
	if err != nil {
		return code.SystemErr[*proto.NodeListResp](err)
	}

	respData := &proto.NodeListResp{Page: 1, PageSize: int64(len(resp.Kvs)), TotalPage: 1, TotalCount: int64(len(resp.Kvs)), List: make([]*proto.NodeInfo, len(resp.Kvs))}
	for index, kv := range resp.Kvs {
		ins := &grpc.Instance{}
		ins.Decode(string(kv.Value))
		addrInfo := strings.Split(ins.Addr, ":")
		respData.List[index] = &proto.NodeInfo{
			Node: strings.ReplaceAll(string(kv.Key), n.prefix(req.Namespace)+"/", ""), Name: ins.Name, Namespace: ins.Namespace, GroupName: ins.Group, Weight: ins.Weight, Version: ins.Version, Host: addrInfo[0], Port: addrInfo[1],
		}
	}

	return code.Succ(respData)
}
