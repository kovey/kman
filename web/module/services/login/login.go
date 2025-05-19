package login

import (
	"os"
	"strings"

	"github.com/kovey/discovery/krpc"
	"github.com/kovey/kman/web/module/libs/code"
	"github.com/kovey/kman/web/module/libs/password"
	"github.com/kovey/kman/web/module/libs/proto"
	"github.com/kovey/kman/web/module/libs/token"
	"github.com/kovey/kman/web/module/libs/token/access"
	"github.com/kovey/kman/web/module/libs/token/refresh"
	"github.com/kovey/kow/context"
	"github.com/kovey/kow/result"
)

type Login struct {
	servName string
	group    string
}

func NewLogin() *Login {
	return &Login{servName: os.Getenv("SERVICE_NAME"), group: os.Getenv("SERVICE_GROUP")}
}

func (l *Login) loginAdmin(req *LoginReqData) *token.Ext {
	account := strings.Split(req.Username, "@")[0]
	pass := password.Password(account, password.Sha256(os.Getenv("ACCT_PASSWORD")))
	if !password.Verify(account, req.Password, pass) {
		return nil
	}

	return &token.Ext{UserId: -1, Permissions: []int64{-1}, ProjectId: 0, Namespace: ""}
}

func (l *Login) Login(ctx *context.Context) error {
	req := ctx.ReqData.(*LoginReqData)
	var ext *token.Ext
	if req.Username == os.Getenv("ACCT_ADMIN") {
		ext = l.loginAdmin(req)
		if ext == nil {
			return result.Err(ctx, code.Password_Err, "Password error")
		}
	} else {
		conn := ctx.Rpcs.Get(krpc.ServiceName(l.servName), l.group)
		if conn == nil {
			return result.Err(ctx, code.System_Err, "conn not found")
		}

		cli := proto.NewLoginClient(conn)
		resp, err := cli.Login(ctx.Context, &proto.LoginReq{Username: req.Username, Namespace: ctx.GetString("namespace"), ProjectId: int32(ctx.GetInt("projectId"))})
		if err != nil {
			return result.Convert(ctx, err)
		}
		info := strings.Split(req.Username, "@")
		if !password.Verify(info[0], req.Password, resp.Password) {
			return result.Err(ctx, code.Password_Err, "Password error")
		}

		ext = &token.Ext{UserId: resp.UserId, Permissions: resp.Permissions, ProjectId: int(resp.ProjectId), Namespace: resp.Namespace}
	}

	t, err := access.Token(ext)
	if err != nil {
		return result.Err(ctx, code.System_Err, err.Error())
	}

	r, err := refresh.Token(ext)
	if err != nil {
		return result.Err(ctx, code.System_Err, err.Error())
	}

	return result.Succ(ctx, LoginRespData{Token: t, RefreshToken: r})
}
