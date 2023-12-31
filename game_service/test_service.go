package game_service

import (
	"fmt"
	"context"
	"strings"

	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/component"
	"github.com/topfreegames/pitaya/v2/logger"
)

type (
	TestService struct {
		Base
		app pitaya.Pitaya
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

func NewTestService(app pitaya.Pitaya) (comp component.Component, service_name string) {
	service_name = "test_service"
	comp = &TestService{
		age : 1,
		app : app,
	}

	app.Register(comp,
		component.WithName(service_name),
		component.WithNameFunc(strings.ToLower),
	)

	return comp, service_name
}