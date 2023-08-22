package main

import (
	"fmt"
	"pitaya_server/dbmodule"

	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/interfaces"
)

const (
	MONGO_MODULE = "mongoDB"
	REDIS_MODULE = "redisDB"
)

type BuildModule func(mdName string, serverConf *ServerConfig) (interfaces.Module, error)

// build mongodb
func buildMongo(mdName string, serverConf *ServerConfig) (interfaces.Module, error) {
	conf := serverConf.MongoConf
	mongoMd := dbmodule.BuildMongo(conf.Host, conf.Port, conf.DataName, conf.MaxNum)
	return mongoMd, nil
}

// build redis
func buildRedis(mdName string, serverConf *ServerConfig) (interfaces.Module, error) {
	conf := serverConf.RedisConf
	redisMd := dbmodule.BuildRedis(conf.Host, conf.Port, conf.MaxIdle, conf.MaxActive, conf.IdleTimeout)
	return redisMd, nil
}

// 所有模块
var moduleMap map[string]BuildModule = map[string]BuildModule{
	MONGO_MODULE: buildMongo,
	REDIS_MODULE: buildRedis,
}

// 启动所有模块
func StartModules(app pitaya.Pitaya, conf *ServerConfig, moduleList []string) error {
	for _, mdName := range moduleList {
		newFunc, ok := moduleMap[mdName]
		if !ok {
			return fmt.Errorf("not find module: %s", mdName)
		}

		md, err := newFunc(mdName, conf)
		if err != nil {
			return err
		}

		app.RegisterModule(md, mdName)
	}
	return nil
}
