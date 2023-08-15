package game_service

import (
	"fmt"
	"context"

	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/component"
	"github.com/topfreegames/pitaya/v2/logger"
)

type (
	TestService struct {
		Base
		app *pitaya.App
		age int
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

func NewTestService(app *pitaya.App) (comp component.Component, service_name string) {
	comp = &TestService{
		age : 1,
		app : app,
	}

	service_name = "test_service"
	return
}