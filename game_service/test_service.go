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
		module_mgr *game_modules.ModuleManager //组件管理器
	}

	TestResponse struct {
		Code  int       `json:"code"`
		Desc  string    `json:"desc"`
	}
)

func (ts *TestService) Test(ctx context.Context, msg []byte) (*TestResponse, error) {
	logger.Log.Info("---请求数据:", string(msg))
	wd, err := ts.module_mgr.GetModule("test_module")
	if err != nil {
		return &TestResponse{
			Code : 300,
			Desc : "请求失败",
		}, nil
	}

	name := wd.name
	return &TestResponse{
		Code : 200,
		Desc : "请求成功",
	}, nil
}

func (ts *TestService) Init() {
	fmt.Println("测试服务初始化:" + string(ts.age))
}

func NewTestService(mdg *game_modules.ModuleManager) (comp component.Component, service_name string) {
	comp = &TestService{
		age : 1,
		module_mgr : mdg,
	}

	service_name = "test_service"
	return
}