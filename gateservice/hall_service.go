package gateservice

import (
	"context"
	"pitaya_server/protofile/protos"
	"strings"

	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/component"
	"github.com/topfreegames/pitaya/v2/logger"
)

type (
	HallService struct {
		Base
		app pitaya.Pitaya
	}

	RequestTest struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	ResponseReq struct {
		Code int    `json:"code"`
		Desc string `json:"desc"`
	}

	MongoTestResult struct {
		Test int    `bson:"test"`
		Name string `bson:"name"`
	}
)

func (hs *HallService) Init() {
	logger.Log.Info("初始大厅服务完成")
}

func (hs *HallService) AfterInit() {
}

func (hs *HallService) TestReq(ctx context.Context) (*ResponseReq, error) {
	logger.Log.Info("test service 请求成功")
	return &ResponseReq{
		Code: 200,
		Desc: "请求成功",
	}, nil
}

func (hs *HallService) TestDBReq(ctx context.Context) (*ResponseReq, error) {
	logger.Log.Info("收到远端请求")

	req := &protos.TestReq{
		Name: "test4",
		Age:  20,
	}
	ret := &protos.TestResp{}
	err := hs.app.RPC(ctx, "dbService.dbservice.remotetest", ret, req)
	if err != nil {
		return &ResponseReq{
			Code: 200,
			Desc: "请求失败",
		}, nil
	}

	return &ResponseReq{
		Code: 200,
		Desc: "请求成功",
	}, nil
}

func NewHallService(app pitaya.Pitaya) (*HallService, string) {
	serviceName := "hallservice"
	hs := &HallService{
		app: app,
	}

	app.Register(hs,
		component.WithName(serviceName),
		component.WithNameFunc(strings.ToLower),
	)

	return hs, serviceName
}
