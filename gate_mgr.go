package main

import (
	"pitaya_server/gateservice"

	"github.com/topfreegames/pitaya/v2/component"
	"github.com/topfreegames/pitaya/v2/logger"
)

type (
	HallServiceMgr struct {
		base   BaseMgr
		server *GameServer
	}
)

// 开启服务
func (mgr *HallServiceMgr) Start(server *GameServer) error {
	mgr.server = server

	var comp component.Component
	var serviceName string

	//hall大厅服务
	comp, serviceName = gateservice.NewHallService(mgr.server.app)
	AppendCp(mgr.base.serviceList, comp, serviceName)

	return nil
}

// 关闭服务
func (mgr *HallServiceMgr) Close() {
	logger.Log.Info("hallserver close server")
}
