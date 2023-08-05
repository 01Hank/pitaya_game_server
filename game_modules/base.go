package game_modules

import (
	"github.com/topfreegames/pitaya/v2/logger"
)

//对interfaces的默认实现, 这里的module模块指的是系统相关的组件
//在game_server里系统组件不对外提供服务，指对内的一些服务提供支持
//比如db服务 活动服务等等


type Base struct {
	name string
}

//服务初始化执行
func (base *Base) Init() error {
	logger.Log.Info("module is init base")
	return nil
}

//服务初始化过后执行
func (base *Base) AfterInit() {}

//服务关闭前执行
func (base *Base) BeforeShutdown() {}

//服务关闭时执行
func (base *Base) Shutdown() error {
	logger.Log.Info("module is shutdown base")
	return nil
}