package main

type NewFunc func(gs *GameServer, sv_type string, isFrontend bool)  ServiceMgrIn


// 新建game服务管理器
func NewGameMgr (gs *GameServer, sv_type string, isFrontend bool) ServiceMgrIn {
	base_mgr := ServiceMgrBase{
		server : gs,
		services : make(map[string]*ServiceBase),
	}

	return &GameServiceMgr{
		base:base_mgr,
		app : gs.app,
	}
}