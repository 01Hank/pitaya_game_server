package game_modules

import (
	"fmt"
	"reflect"

	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/logger"
	"github.com/topfreegames/pitaya/v2/interfaces"
)

type (
	MdWrapper struct {
		module   interfaces.Module 
		name     string
		m_type   reflect.Type
	}

	ModuleManager struct {
		modules map[string]*MdWrapper //所有的组件
	}
)

// 获取module服务
func (mgr *ModuleManager) GetModule(name string) (interface{}, error) {
	if len(mgr.modules) <= 0 {
		return nil, fmt.Errorf("not find module")
	}

	md, ok := mgr.modules[name]
	if !ok {
		return nil, fmt.Errorf("not find module")
	}

	return md, nil
}

// 注册所有服务
func (mgr *ModuleManager) RegisterToPitaya(app pitaya.Pitaya, exclude_modules []string) error {
	ex_md_map := make(map[string]bool, 0)
	for _, module_name := range exclude_modules {
		ex_md_map[module_name] = true
	}
	
	for name, md := range mgr.modules {
		if _, ok := ex_md_map[name]; ok {
			continue
		}

		err := app.RegisterModule(md.module, name)
		if err != nil {
			panic(err)
		}
		logger.Log.Info("registered game_module suc and name: %s", name)
	}

	return nil
}

// 初始化服务
func InitModules() *ModuleManager {
	maps := make(map[string]*MdWrapper, 0)

	//测试组件
	maps["test_module"] = &MdWrapper{
		module : NewTestModule(),
		name : "test_module",
	}

	return &ModuleManager{
		modules : maps,
	}
}