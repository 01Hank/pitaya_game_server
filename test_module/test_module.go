package test_module

import (	
	"github.com/topfreegames/pitaya/v2/modules"
	"github.com/topfreegames/pitaya/v2/logger"
)

type (
	TestModule struct {
		modules.Base
		name string
	}
)

func (tm *TestModule) PrintTest(){
	logger.Log.Info("测试组件请求打印")
}

func (tm *TestModule) Init() error{
	logger.Log.Info("测试组件初始化")
	return nil
}


func NewTestModule() *TestModule {
	return &TestModule{
		name : "test",
	}
}