package dbservice

import (
	"sync"

	"github.com/topfreegames/pitaya/v2/logger"
)

type Base struct {
	bindData sync.Map
}

// Init was called to initialize the component.
func (c *Base) Init() {
	logger.Log.Info("dbService init")
}

// AfterInit was called after the component is initialized.
func (c *Base) AfterInit() {
	logger.Log.Info("dbService afterinit")
}

// BeforeShutdown was called before the component to shutdown.
func (c *Base) BeforeShutdown() {
	logger.Log.Info("dbService beforeshutdown")
}

// Shutdown was called to shutdown the component.
func (c *Base) Shutdown() {
	logger.Log.Info("dbService shutdown")
}

// 绑定
func (c *Base) Bind(key, value string) bool {
	_, ok := c.bindData.LoadOrStore(key, value)
	return ok
}

// 获取绑定的数据
func (c *Base) GetBind(key string) (string, bool) {
	value, ok := c.bindData.Load(key)
	return value.(string), ok
}

// 删除绑定的数据
func (c *Base) DelBind(key string) {
	c.bindData.Delete(key)
}
