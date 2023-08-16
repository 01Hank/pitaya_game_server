package main

import "github.com/topfreegames/pitaya/v2"

type ServiceMgrIn interface {
	Start(app pitaya.Pitaya) error
	Close()
}