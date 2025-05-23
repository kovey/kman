package event

import (
	"fmt"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/kovey/cli-go/app"
	"github.com/kovey/cli-go/util"
	"github.com/kovey/db-go/v3/db"
	"github.com/kovey/kman/kman-service/module/libs/etcd"
	"github.com/kovey/kom/server"
)

type ServEvent struct {
	*server.EventBase
}

func NewServEvent() *ServEvent {
	return &ServEvent{EventBase: &server.EventBase{}}
}

func (s *ServEvent) OnFlag(a app.AppInterface) error {
	return nil
}

func (s *ServEvent) initMysql() error {
	dns := mysql.Config{
		User:   os.Getenv("DB_USER"),
		Passwd: os.Getenv("DB_PASSWORD"),
		DBName: os.Getenv("DB_NAME"),
		Addr:   fmt.Sprintf("%s:%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT")),
		Params: map[string]string{"charset": os.Getenv("DB_CHARSET"), "parseTime": "true"},
	}
	return db.Init(db.Config{
		DriverName:     os.Getenv("DB_DRIVER"),
		DataSourceName: dns.FormatDSN(),
		MaxIdleTime:    60 * time.Second,
		MaxLifeTime:    120 * time.Second,
		MaxIdleConns:   10,
		MaxOpenConns:   20,
		LogOpened:      true,
		LogMax:         2048,
	})
}

func (s *ServEvent) OnBefore(app.AppInterface) error {
	if err := s.initMysql(); err != nil {
		return err
	}

	return etcd.Init()
}

func (s *ServEvent) OnAfter(app.AppInterface) error {
	return nil
}

func (s *ServEvent) OnRun() error {
	return nil
}

func (s *ServEvent) OnShutdown() {
	db.Close()
	etcd.Close()
}

func (s *ServEvent) CreateConfig(path string) error {
	filePath := fmt.Sprintf("%s/.env", path)
	if util.IsFile(filePath) {
		return fmt.Errorf("[%s] is exists", filePath)
	}

	return os.WriteFile(filePath, []byte(env_config), 0644)
}
