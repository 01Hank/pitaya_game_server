package main

import (
	"fmt"
	"pitaya_game_server/game_service"

	"github.com/topfreegames/pitaya/v2"
)

type (
	GameServiceMgr struct {
		base ServiceMgrBase
		app pitaya.Pitaya
	}
)

func (game *GameServiceMgr) Start(app pitaya.Pitaya) error {
	game.app = app

	//test服务
	comp, service_name := game_service.NewTestService(app)
	AppendCp(game.base.services, comp, service_name)

	return nil
}

func (game *GameServiceMgr) Close(){
	fmt.Println("关闭")
}