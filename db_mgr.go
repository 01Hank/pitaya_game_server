package main

import (
	"pitaya_server/dbservice"

	"github.com/topfreegames/pitaya/v2/component"
	"github.com/topfreegames/pitaya/v2/logger"
)

type (
	DBServiceMgr struct {
		base   BaseMgr
		server *GameServer
	}
)

// 开启服务
func (mgr *DBServiceMgr) Start(server *GameServer) error {
	mgr.server = server

	var comp component.Component
	var serviceName string

	//db大厅服务
	comp, serviceName = dbservice.NewDBService(mgr.server.app)
	AppendCp(mgr.base.serviceList, comp, serviceName)

	return nil
}

// 关闭服务
func (mgr *DBServiceMgr) Close() {
	logger.Log.Info("dbservice close server")
}
