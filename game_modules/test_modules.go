package game_modules

import (
	"github.com/topfreegames/pitaya/v2/logger"
	"github.com/topfreegames/pitaya/v2/interfaces"
)

type TestModule struct {
	Base
	age  int
}

func (tm *TestModule) Init() error {
	logger.Log.Info("TestModule 初始化")
	return nil
}

func NewTestModule() interfaces.Module {
	return &TestModule{
		age : 1,
	}
}