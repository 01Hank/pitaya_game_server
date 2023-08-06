package game_service

import (
	"fmt"
	"context"

	"pitaya_game_server/game_modules"

	"github.com/topfreegames/pitaya/v2/component"
	"github.com/topfreegames/pitaya/v2/logger"
)

type (
	TestService struct {
		Base
		age int
		module_mgr *game_modules.ModuleManager
	}

	TestResponse struct {
		Code  int       `json:"code"`
		Desc  string    `json:"desc"`
	}
)

func (ts *TestService) Test(ctx context.Context, msg []byte) (*TestResponse, error) {
	logger.Log.Info("---请求数据:", string(msg))
	return &TestResponse{
		Code : 200,
		Desc : "请求成功",
	}, nil
}

func (ts *TestService) Init() {
	fmt.Println("测试服务初始化:" + string(ts.age))
}

func NewTestService(mdg *game_modules.ModuleManager) component.Component {
	return &TestService{
		age : 1,
		module_mgr : mdg,
	}
}