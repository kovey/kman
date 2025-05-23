package config

import (
	"fmt"
	"time"

	ksql "github.com/kovey/db-go/v3"
	"github.com/kovey/db-go/v3/db"
	"github.com/kovey/kman/kman-service/module/libs/code"
	"github.com/kovey/kman/kman-service/module/libs/etcd"
	"github.com/kovey/kman/kman-service/module/libs/proto"
	"github.com/kovey/kman/kman-service/module/models"
	"github.com/kovey/pool"
	"github.com/kovey/pool/object"
)

const (
	ctx_namespace = "busi.config"
	ctx_config    = "Config"
)

func init() {
	pool.Default(ctx_namespace, ctx_config, func() any {
		return &Config{Object: object.NewObject(ctx_namespace, ctx_config)}
	})
}

type Config struct {
	*object.Object
}

func NewConfig(ctx object.CtxInterface) *Config {
	return ctx.Get(ctx_namespace, ctx_config).(*Config)
}

func (cc *Config) Add(req *proto.ConfigAddReq) (*proto.ConfigAddResp, error) {
	ok, err := db.Model(models.NewConfigs()).Where(models.Table_Configs_ProjectId, ksql.Eq, req.ProjectId).Where(models.Table_Configs_ConfigKey, ksql.Eq, req.Key).Exist(cc.Context)
	if err != nil {
		return code.SystemErr[*proto.ConfigAddResp](err)
	}

	if ok {
		return code.Err[*proto.ConfigAddResp](code.Config_Key_Exists, fmt.Errorf("Config key is exists"), nil)
	}

	str := ""
	c := models.NewConfigs()
	c.Name = req.Name
	c.ConfigKey = req.Key
	c.ProjectId = int(req.ProjectId)
	c.ConfigValue = req.Value
	c.History1 = &str
	c.History2 = &str
	c.CreateTime = time.Now().Unix()
	c.UpdateTime = c.CreateTime
	c.Status = 1
	if err := c.Save(cc.Context); err != nil {
		return code.SystemErr[*proto.ConfigAddResp](err)
	}

	return code.Succ(&proto.ConfigAddResp{})
}

func (c *Config) Edit(req *proto.ConfigEditReq) (*proto.ConfigEditResp, error) {
	var row = models.NewConfigs()
	if err := db.Find(c.Context, row, req.Id); err != nil {
		return code.SystemErr[*proto.ConfigEditResp](err)
	}
	if row.Empty() || row.ProjectId != int(req.ProjectId) {
		return code.Err[*proto.ConfigEditResp](code.Config_Not_Found, fmt.Errorf("Config not found"), nil)
	}

	old := row.ConfigValue
	row.ConfigValue = req.Value
	row.UpdateTime = time.Now().Unix()
	row.History2 = row.History1
	row.History1 = &old
	row.Status = 1
	if err := row.Save(c.Context); err != nil {
		return code.SystemErr[*proto.ConfigEditResp](err)
	}

	return code.Succ(&proto.ConfigEditResp{})
}

func (c *Config) Release(req *proto.ConfigReleaseReq) (*proto.ConfigReleaseResp, error) {
	var rows []*models.Configs
	if err := db.Models(&rows).WhereIn(models.Table_Configs_Id, db.ToList(req.Ids)).Where(models.Table_Configs_Status, ksql.Eq, 1).All(c.Context); err != nil {
		return code.SystemErr[*proto.ConfigReleaseResp](err)
	}

	for _, row := range rows {
		_, err := etcd.Put(c.Context, fmt.Sprintf("/ko/configs/%s/%s", req.Namespace, row.ConfigKey), row.ConfigValue)
		if err != nil {
			return code.SystemErr[*proto.ConfigReleaseResp](err)
		}

		row.Status = 0
		row.UpdateTime = time.Now().Unix()
		if err := row.Save(c.Context); err != nil {
			return code.SystemErr[*proto.ConfigReleaseResp](err)
		}
	}

	return code.Succ(&proto.ConfigReleaseResp{})
}

func (c *Config) List(req *proto.ConfigListReq) (*proto.ConfigListResp, error) {
	builder := db.Models(&[]*models.Configs{}).Where(models.Table_Configs_ProjectId, ksql.Eq, req.ProjectId).OrderDesc(models.Table_Configs_Id)
	if req.Key != "" {
		builder.Where(models.Table_Configs_ConfigKey, ksql.Like, fmt.Sprintf("%%%s%%", req.Key))
	}
	pageInfo, err := builder.Pagination(c.Context, req.Page, req.PageSize)
	if err != nil {
		return code.SystemErr[*proto.ConfigListResp](err)
	}

	respData := &proto.ConfigListResp{Page: 1, PageSize: req.PageSize, TotalPage: int64(pageInfo.TotalPage()), TotalCount: int64(pageInfo.TotalCount()), List: make([]*proto.ConfigInfo, len(pageInfo.List()))}
	for index, row := range pageInfo.List() {
		respData.List[index] = &proto.ConfigInfo{
			Id: row.Id, Name: row.Name, Key: row.ConfigKey, Value: row.ConfigValue, History1: *row.History1, History2: *row.History2,
			UpdateTime: time.Unix(row.UpdateTime, 0).Format(time.DateTime), CreateTime: time.Unix(row.CreateTime, 0).Format(time.DateTime),
		}

		if row.Status == 0 {
			respData.List[index].Status = "已发布"
		} else {
			respData.List[index].Status = "待发布"
		}
	}

	return code.Succ(respData)
}
