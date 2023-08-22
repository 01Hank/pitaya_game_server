package dbservice

import (
	"context"
	"fmt"
	"strings"

	"pitaya_server/dbmodule"
	"pitaya_server/protofile/protos"

	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/component"
	"github.com/topfreegames/pitaya/v2/logger"
)

type (
	TestDBRemoteService struct {
		Base
		app     pitaya.Pitaya
		mongoDB *dbmodule.MongoDBClient
		redisDB *dbmodule.RedisDBClient
	}

	TestDB struct {
		Test string `bson:"test"`
		Name string `bson:"name"`
		Age  int    `bson:"age"`
	}
)

// 初始化db连接
func (tds *TestDBRemoteService) AfterInit() {
	mModule, err := tds.app.GetModule("mongoDB")
	if err != nil {
		logger.Log.Error("init testdbservice error")
		return
	}

	tds.mongoDB = mModule.(*dbmodule.MongoDBClient)

	rModule, err := tds.app.GetModule("redisDB")
	if err != nil {
		logger.Log.Error("init testdbservice error")
		return
	}
	tds.redisDB = rModule.(*dbmodule.RedisDBClient)
	logger.Log.Info("init testdbservice suc")
}

// 远端请求
func (tds *TestDBRemoteService) RemoteTest(ctx context.Context, msg *protos.TestReq) (*protos.TestResp, error) {
	logger.Log.Info("收到远端请求: %s", msg)
	logger.Log.Info(ctx)

	// td := &TestDB{
	// 	Name: msg.Name,
	// 	Age:  int(msg.Age),
	// }
	// err := tds.mongoDB.SaveOne("test", td)
	return &protos.TestResp{
		Err: fmt.Sprintf("%v", nil),
	}, nil
}

func NewDBService(app pitaya.Pitaya) (*TestDBRemoteService, string) {
	serviceName := "dbservice"
	tds := &TestDBRemoteService{
		app: app,
	}

	app.RegisterRemote(tds,
		component.WithName(serviceName),
		component.WithNameFunc(strings.ToLower),
	)

	return tds, serviceName
}
