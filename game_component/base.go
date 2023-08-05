package game_component

import (
	"github.com/topfreegames/pitaya/v2/logger"
)

//对interfaces的默认实现, 这里的component模块指的是对外提供的服务组件
//在game_server里服务组件负责对外部请求作出响应
//比如房间服务 任务服务等等


type Base struct {
	name string // 服务名称
}

//服务初始化执行
func (base *Base) Init() error {
	logger.Log.Info("component is init and name: %s", base.name)
	return nil
}

//服务初始化过后执行
func (base *Base) AfterInit() {}

//服务关闭前执行
func (base *Base) BeforeShutdown() {}

//服务关闭时执行
func (base *Base) Shutdown() error {
	logger.Log.Info("component is shutdown and name: %s", base.name)
	return nil
}