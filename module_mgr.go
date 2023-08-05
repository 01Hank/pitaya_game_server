package main

import (
	"fmt"

	"pitaya_game_server/game_modules"
	"github.com/topfreegames/pitaya/v2/logger"
	"github.com/topfreegames/pitaya/v2/interfaces"
)

type (
	MdWrapper struct {
		module interfaces.Module 
		name string
	}

	Manager struct {
		modules map[string]*MdWrapper //所有的组件
	}
)

// 获取module服务
func (mgr *Manager) GetModule(name string) (*MdWrapper, error) {
	if len(mgr.modules) <= 0 {
		return nil, fmt.Errorf("not find module")
	}

	md, ok := mgr.modules[name]
	if !ok {
		return nil, fmt.Errorf("not find module")
	}

	return md, nil
}

// 初始化服务
func InitModules() *Manager {
	maps := make(map[string]*MdWrapper, 0)

	//测试组件
	maps["test_module"] = &MdWrapper{
		module : game_modules.NewTestModule(),
		name : "test_module",
	}


	logger.Log.Info("--初始成功, len: %d", len(maps))
	return &Manager{
		modules : maps,
	}
}