package game_service

import (
	"fmt"
	"context"

	"pitaya_game_server/test_module"

	"github.com/topfreegames/pitaya/v2/component"
	"github.com/topfreegames/pitaya/v2/logger"
)

type (
	TestService struct {
		Base
		age int
		tm *test_module.TestModule
	}

	TestResponse struct {
		Code  int       `json:"code"`
		Desc  string    `json:"desc"`
	}
)

func (ts *TestService) Test(ctx context.Context, msg []byte) (*TestResponse, error) {
	logger.Log.Info("---请求数据:", string(msg))
	ts.tm.PrintTest()
	return &TestResponse{
		Code : 200,
		Desc : "请求成功",
	}, nil
}

func (ts *TestService) Init() {
	fmt.Println("测试服务初始化:" + string(ts.age))
}

func NewTestService(tmd *test_module.TestModule) (comp component.Component, service_name string) {
	comp = &TestService{
		age : 1,
		tm : tmd,
	}

	service_name = "test_service"
	return
}