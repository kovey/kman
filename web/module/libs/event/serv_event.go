package event

import (
	"fmt"
	"os"

	"github.com/kovey/cli-go/app"
	"github.com/kovey/cli-go/util"
	"github.com/kovey/discovery/algorithm"
	"github.com/kovey/discovery/krpc"
	"github.com/kovey/kow/serv"
)

type ServEvent struct {
	*serv.EventBase
}

func NewServEvent() *ServEvent {
	return &ServEvent{EventBase: &serv.EventBase{}}
}

func (s *ServEvent) OnFlag(app.AppInterface) error {
	return nil
}

func (s *ServEvent) OnBefore(app.AppInterface) error {
	krpc.SetLoadBalance(algorithm.Alg_Random_Weight)
	return nil
}

func (s *ServEvent) OnAfter(app.AppInterface) error {
	return nil
}

func (s *ServEvent) OnRun() error {
	return nil
}

func (s *ServEvent) OnShutdown() {
}

func (s *ServEvent) CreateConfig(path string) error {
	filePath := fmt.Sprintf("%s/.env", path)
	if util.IsFile(filePath) {
		return fmt.Errorf("[%s] is exists", filePath)
	}

	return os.WriteFile(filePath, []byte(env_config), 0644)
}
