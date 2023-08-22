package main

type ServiceMgrIN interface {
	Start(server *GameServer) error
	Close()
}
