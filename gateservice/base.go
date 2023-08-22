package gateservice

import "github.com/topfreegames/pitaya/v2/logger"

type Base struct{}

// Init was called to initialize the component.
func (c *Base) Init() {
	logger.Log.Info("gateService init")
}

// AfterInit was called after the component is initialized.
func (c *Base) AfterInit() {
	logger.Log.Info("gateService afterinit")
}

// BeforeShutdown was called before the component to shutdown.
func (c *Base) BeforeShutdown() {
	logger.Log.Info("gateService beforeshutdown")
}

// Shutdown was called to shutdown the component.
func (c *Base) Shutdown() {
	logger.Log.Info("gateService shutdown")
}
