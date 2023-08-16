package main

import (
	"github.com/topfreegames/pitaya/v2/component"
)

type (
	ServiceMgrBase struct {
		server     *GameServer           //server引用
		services   map[string]*ServiceBase //所有的服务
		is_frontend bool // 是否前置服务
	}

	ServiceBase struct {
		component component.Component
		name string
	}
)

// 添加服务
func AppendCp(ms map[string]*ServiceBase, comp component.Component, service_name string){
	ms[service_name] = &ServiceBase{
		name      : service_name,
		component : comp,
	}
}