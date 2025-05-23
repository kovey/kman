package models

// Code generated by ksql.
// Do'nt Edit!!!
// Do'nt Edit!!!
// Do'nt Edit!!!
//
// from database: kman
// table:         configs
// orm version:   1.0.5
// created time:  2025-05-13 17:26:03
/**
Table DDL:
CREATE TABLE `configs` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `project_id` int NOT NULL DEFAULT '0' COMMENT '项目',
  `name` varchar(127) NOT NULL DEFAULT '' COMMENT '名称',
  `config_key` varchar(255) NOT NULL DEFAULT '' COMMENT '配置key',
  `config_value` text NOT NULL COMMENT '配置值',
  `history_1` text COMMENT '历史记录1',
  `history_2` text COMMENT '历史记录2',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态 0 - 已发布 1 - 待发布',
  `create_time` bigint NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` bigint NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `config_key_UNIQUE` (`project_id`,`config_key`),
  KEY `idx_project_id` (`project_id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
*/

import (
	"context"

	"github.com/kovey/db-go/v3"
	"github.com/kovey/db-go/v3/model"
)

const (
	Table_Configs             = "configs"      //
	Table_Configs_ConfigKey   = "config_key"   // 配置key
	Table_Configs_ConfigValue = "config_value" // 配置值
	Table_Configs_CreateTime  = "create_time"  // 创建时间
	Table_Configs_History1    = "history_1"    // 历史记录1
	Table_Configs_History2    = "history_2"    // 历史记录2
	Table_Configs_Id          = "id"           //
	Table_Configs_Name        = "name"         // 名称
	Table_Configs_ProjectId   = "project_id"   // 项目
	Table_Configs_Status      = "status"       // 状态 0 - 已发布 1 - 待发布
	Table_Configs_UpdateTime  = "update_time"  // 更新时间
)

type Configs struct {
	*model.Model `db:"-" json:"-"` // model
	ConfigKey    string            `db:"config_key" json:"config_key"`     // 配置key
	ConfigValue  string            `db:"config_value" json:"config_value"` // 配置值
	CreateTime   int64             `db:"create_time" json:"create_time"`   // 创建时间
	History1     *string           `db:"history_1" json:"history_1"`       // 历史记录1
	History2     *string           `db:"history_2" json:"history_2"`       // 历史记录2
	Id           int64             `db:"id" json:"id"`                     //
	Name         string            `db:"name" json:"name"`                 // 名称
	ProjectId    int               `db:"project_id" json:"project_id"`     // 项目
	Status       int8              `db:"status" json:"status"`             // 状态 0 - 已发布 1 - 待发布
	UpdateTime   int64             `db:"update_time" json:"update_time"`   // 更新时间
}

func NewConfigs() *Configs {
	return &Configs{Model: model.NewModel(Table_Configs, Table_Configs_Id, model.Type_Int)}
}

func (self *Configs) Save(ctx context.Context) error {
	return self.Model.Save(ctx, self)
}

func (self *Configs) Clone() ksql.RowInterface {
	return NewConfigs()
}

func (self *Configs) Values() []any {
	return []any{&self.ConfigKey, &self.ConfigValue, &self.CreateTime, &self.History1, &self.History2, &self.Id, &self.Name, &self.ProjectId, &self.Status, &self.UpdateTime}
}

func (self *Configs) Columns() []string {
	return []string{Table_Configs_ConfigKey, Table_Configs_ConfigValue, Table_Configs_CreateTime, Table_Configs_History1, Table_Configs_History2, Table_Configs_Id, Table_Configs_Name, Table_Configs_ProjectId, Table_Configs_Status, Table_Configs_UpdateTime}
}

func (self *Configs) Delete(ctx context.Context) error {
	return self.Model.Delete(ctx, self)
}

func (self *Configs) Query() ksql.BuilderInterface[*Configs] {
	return model.Row(self)
}
