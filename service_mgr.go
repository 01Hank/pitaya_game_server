package main

import (
	"pitaya_game_server/game_service"

	"github.com/topfreegames/pitaya/v2/component"
)

type (
	CpWrapper struct {
		component component.Component
		name string
	}

	ServiceManager struct {
		server     *GameServer           //server引用
		services   map[string]*CpWrapper //所有的服务
	}
)

// 注册所有服务
func (mgr *ServiceManager) RegisterServices(gs *GameServer, exclude_components []string) error {
	mgr.server = gs

	//test服务
	comp, service_name := game_service.NewTestService(gs.app)
	appendCp(mgr.services, comp, service_name)
	
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
		services : make(map[string]*CpWrapper, 0),
	}
}