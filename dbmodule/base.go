package dbmodule

import (
	"github.com/topfreegames/pitaya/v2/logger"
)

// db模块基类
type BaseDB struct {
	DBHost    string
	DBPort    string
	IsConnect bool
}

// Init was called to initialize the component.
func (b *BaseDB) Init() error {
	logger.Log.Warn("init base db host: %s, port: %s", b.DBHost, b.DBPort)
	return nil
}

// AfterInit was called after the component is initialized.
func (b *BaseDB) AfterInit() {}

// BeforeShutdown was called before the component to shutdown.
func (b *BaseDB) BeforeShutdown() {}

// Shutdown was called to shutdown the component.
func (b *BaseDB) Shutdown() error {
	return nil
}
