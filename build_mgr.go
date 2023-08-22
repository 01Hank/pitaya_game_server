package main

// 定义new方法
type NewBuild func(gs *GameServer, svType string, isFrontend bool) (ServiceMgrIN, *ServerConfig)

// 创建网关服务
func NewGateMgr(gs *GameServer, svType string, isFrontend bool) (ServiceMgrIN, *ServerConfig) {
	baseMgr := BaseMgr{
		svType:      svType,
		isFrontend:  isFrontend,
		serviceList: make(map[string]*BaseService),
	}

	GateServiceMgr := &HallServiceMgr{
		base:   baseMgr,
		server: gs,
	}

	needModules := []string{}
	conf := serverConf(svType)
	StartModules(gs.app, conf, needModules)
	return GateServiceMgr, conf
}

// 创建db服务
func NewDBMgr(gs *GameServer, svType string, isFrontend bool) (ServiceMgrIN, *ServerConfig) {
	baseMgr := BaseMgr{
		svType:      svType,
		isFrontend:  isFrontend,
		serviceList: make(map[string]*BaseService),
	}

	DBServiceMgr := &DBServiceMgr{
		base:   baseMgr,
		server: gs,
	}

	needModules := []string{
		MONGO_MODULE,
		REDIS_MODULE,
	}

	conf := serverConf(svType)
	StartModules(gs.app, conf, needModules)
	return DBServiceMgr, conf
}

// 获取服务器配置
func serverConf(svType string) *ServerConfig {
	mongoConfig := MongoConfig{
		Host:     "127.0.0.1",
		Port:     "27017",
		DataName: "test",
		MaxNum:   10,
	}

	redisConfig := RedisConfig{
		Host:        "127.0.0.1",
		Port:        "6379",
		MaxIdle:     10,
		MaxActive:   0,
		IdleTimeout: 10,
	}

	return &ServerConfig{
		MongoConf: mongoConfig,
		RedisConf: redisConfig,
	}
}
