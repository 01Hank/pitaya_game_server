package main

import (
	"github.com/topfreegames/pitaya/v2/component"
)

// 服务管理器基类
type BaseMgr struct {
	svType      string
	isFrontend  bool
	serviceList map[string]*BaseService
}

// 服务基类
type BaseService struct {
	serviceName string
	comp        component.Component
}

// 添加服务
func AppendCp(ms map[string]*BaseService, comp component.Component, serviceName string) {
	ms[serviceName] = &BaseService{
		serviceName: serviceName,
		comp:        comp,
	}
}
