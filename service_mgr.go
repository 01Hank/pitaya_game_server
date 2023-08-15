package main

import (
	"fmt"
	"strings"

	"pitaya_game_server/game_service"

	"github.com/topfreegames/pitaya/v2/logger"
	"github.com/topfreegames/pitaya/v2/component"
)

type (
	CpWrapper struct {
		component component.Component
		name string
	}

	ServiceManager struct {
		server     *GameServer           //server引用
		components map[string]*CpWrapper //所有的服务
	}
)

func (mgr *ServiceManager) registerService(exclude_components []string) error {
	ex_cp_map := make(map[string]bool, 0)
	for _, component_name := range exclude_components {
		ex_cp_map[component_name] = true
	}
	
	for name, md := range mgr.components {
		if _, ok := ex_cp_map[name]; ok {
			continue
		}

		//注册本地服务
		mgr.server.app.Register(md.component, 
			component.WithName(name),
			component.WithNameFunc(strings.ToLower),
		)

		//注册远端服务
		mgr.server.app.RegisterRemote(md.component,
			component.WithName(name),
			component.WithNameFunc(strings.ToLower),
		)

		logger.Log.Info("registered service suc and name: %s", name)
	}

	return nil
}

// 注册所有服务
func (mgr *ServiceManager) RegisterServices(gs *GameServer, exclude_components []string) error {
	mgr.server = gs

	//test服务
	comp, service_name := game_service.NewTestService(gs.app)
	appendCp(mgr.components, comp, service_name)

	//注册服务
	mgr.registerService(exclude_components)
	
	return nil
}

// 添加服务
func appendCp(ms map[string]*CpWrapper, comp component.Component, service_name string){
	ms[service_name] = &CpWrapper{
		name      : service_name,
		component : comp,
	}
}

// 初始化服务
func InitServices() *ServiceManager {
	return &ServiceManager{
		components : make(map[string]*CpWrapper, 0),
	}
}