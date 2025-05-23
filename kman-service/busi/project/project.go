package project

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"time"

	ksql "github.com/kovey/db-go/v3"
	"github.com/kovey/db-go/v3/db"
	"github.com/kovey/kman/kman-service/module/libs/code"
	"github.com/kovey/kman/kman-service/module/libs/proto"
	"github.com/kovey/kman/kman-service/module/models"
	"github.com/kovey/kom"
	"github.com/kovey/pool"
	"github.com/kovey/pool/object"
)

const (
	ctx_namespace = "busi.project"
	ctx_project   = "Project"
)

func init() {
	pool.Default(ctx_namespace, ctx_project, func() any {
		return &Project{Object: object.NewObject(ctx_namespace, ctx_project)}
	})
}

type Project struct {
	*object.Object
}

func NewProject(ctx object.CtxInterface) *Project {
	return ctx.Get(ctx_namespace, ctx_project).(*Project)
}

func Md5(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func (p *Project) Add(req *proto.ProjectAddReq) (*proto.ProjectAddResp, error) {
	if req.Namespace == os.Getenv(kom.ETCD_NAMESPACE) {
		return code.Err[*proto.ProjectAddResp](code.Project_Exists, fmt.Errorf("Project is exists"), nil)
	}

	ok, err := db.Model(models.NewProjects()).Where(models.Table_Projects_Namespace, ksql.Eq, req.Namespace).Exist(p.Context)
	if err != nil {
		return code.SystemErr[*proto.ProjectAddResp](err)
	}

	if ok {
		return code.Err[*proto.ProjectAddResp](code.Project_Exists, fmt.Errorf("Project is exists"), nil)
	}

	admin := models.NewProjects()
	admin.Name = req.Name
	admin.Namespace = req.Namespace
	admin.CreateTime = time.Now().Unix()
	admin.UpdateTime = admin.CreateTime
	admin.OpenId = Md5(fmt.Sprintf("%s-%d", admin.Namespace, time.Now().UnixNano()))
	if err := admin.Save(p.Context); err != nil {
		return code.Err[*proto.ProjectAddResp](code.System_Err, err, nil)
	}

	return code.Succ(&proto.ProjectAddResp{})
}

func (p *Project) Edit(req *proto.ProjectEditReq) (*proto.ProjectEditResp, error) {
	admin := models.NewProjects()
	err := db.Find(p.Context, admin, req.Id)
	if err != nil {
		return code.SystemErr[*proto.ProjectEditResp](err)
	}
	if admin.Empty() {
		return code.Err[*proto.ProjectEditResp](code.Project_Not_Found, fmt.Errorf("Project Not Found"), nil)
	}
	admin.Name = req.Name
	admin.UpdateTime = time.Now().Unix()
	if err := admin.Save(p.Context); err != nil {
		return code.SystemErr[*proto.ProjectEditResp](err)
	}

	return code.Succ(&proto.ProjectEditResp{})
}

func (p *Project) List(req *proto.ProjectListReq) (*proto.ProjectListResp, error) {
	builder := db.Models(&[]*models.Projects{}).OrderDesc("id")
	if req.Name != "" {
		builder.Where("name", ksql.Eq, req.Name)
	}

	pageInfo, err := builder.Pagination(p.Context, req.Page, req.PageSize)
	if err != nil {
		return code.SystemErr[*proto.ProjectListResp](err)
	}

	resp := &proto.ProjectListResp{Page: req.Page, PageSize: req.PageSize, TotalPage: int64(pageInfo.TotalPage()), TotalCount: int64(pageInfo.TotalCount())}
	resp.List = make([]*proto.ProjectInfo, len(pageInfo.List()))
	for i, row := range pageInfo.List() {
		resp.List[i] = &proto.ProjectInfo{Id: int32(row.Id), Name: row.Name, Namespace: row.Namespace, OpenId: row.OpenId, CreateTime: time.Unix(row.CreateTime, 0).Format(time.DateTime)}
	}

	return code.Succ(resp)
}
