package main

import (
	"errors"
	"flag"
	"fmt"
	"strconv"
	"time"

	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/acceptor"
	"github.com/topfreegames/pitaya/v2/cluster"
	"github.com/topfreegames/pitaya/v2/config"
	"github.com/topfreegames/pitaya/v2/constants"
	"github.com/topfreegames/pitaya/v2/groups"
	"github.com/topfreegames/pitaya/v2/logger"
	"github.com/topfreegames/pitaya/v2/modules"
)

type (
	GameServer struct {
		app pitaya.Pitaya
		mgr ServiceMgrIN
	}

	MongoConfig struct {
		Host     string
		Port     string
		DataName string
		MaxNum   int
	}

	RedisConfig struct {
		Host        string
		Port        string
		MaxIdle     int
		MaxActive   int
		IdleTimeout time.Duration
	}

	ServerConfig struct {
		MongoConf MongoConfig
		RedisConf RedisConfig
	}
)

// 启动
func (gs *GameServer) start(conf *ServerConfig) {
	defer gs.app.Shutdown()
	defer gs.mgr.Close()

	gs.mgr.Start(gs)
	gs.app.Start()
}

var app pitaya.Pitaya
var gameServer *GameServer

func main() {
	port := flag.Int("port", 3250, "the port to listen")
	svType := flag.String("type", "connector", "the server type")
	isFrontend := flag.Bool("frontend", true, "if server is frontend")
	rpcServerPort := flag.Int("rpcsvport", 3434, "the port that grpc server will listen")

	flag.Parse()

	meta := map[string]string{
		constants.GRPCHostKey: "127.0.0.1",
		constants.GRPCPortKey: strconv.Itoa(*rpcServerPort),
	}

	var bs *modules.ETCDBindingStorage
	app, bs = createApp(*port, *isFrontend, *svType, meta, *rpcServerPort)
	app.RegisterModule(bs, "bindingsStorage")

	gameServer = &GameServer{}
	gameServer.app = app
	conf, err := createGameServer(*svType, *isFrontend, gameServer)
	if err != nil {
		logger.Log.Errorf(fmt.Sprintf("%v", err))
		return
	}

	logger.Log.Info("server is start and svType is:", *svType)
	gameServer.start(conf)
}

// 创建一个pitaya服务
func createApp(port int, isFrontend bool, svType string, meta map[string]string, rpcServerPort int) (pitaya.Pitaya, *modules.ETCDBindingStorage) {
	builder := pitaya.NewDefaultBuilder(isFrontend, svType, pitaya.Cluster, meta, *config.NewDefaultBuilderConfig())

	grpcServerConfig := config.NewDefaultGRPCServerConfig()
	grpcServerConfig.Port = rpcServerPort
	gs, err := cluster.NewGRPCServer(*grpcServerConfig, builder.Server, builder.MetricsReporters)
	if err != nil {
		panic(err)
	}
	builder.RPCServer = gs
	builder.Groups = groups.NewMemoryGroupService(*config.NewDefaultMemoryGroupConfig())

	bs := modules.NewETCDBindingStorage(builder.Server, builder.SessionPool, *config.NewDefaultETCDBindingConfig())

	gc, err := cluster.NewGRPCClient(
		*config.NewDefaultGRPCClientConfig(),
		builder.Server,
		builder.MetricsReporters,
		bs,
		cluster.NewInfoRetriever(*config.NewDefaultInfoRetrieverConfig()),
	)
	if err != nil {
		panic(err)
	}
	builder.RPCClient = gc

	if isFrontend {
		tcp := acceptor.NewTCPAcceptor(fmt.Sprintf(":%d", port))
		builder.AddAcceptor(tcp)
	}

	return builder.Build(), bs
}

// 所有服务集合
var mgrList map[string]NewBuild = map[string]NewBuild{
	"gateService": NewGateMgr, // 网关服务
	"dbService":   NewDBMgr,   // db服务
}

// 创建服务
func createGameServer(svType string, isFrontend bool, gs *GameServer) (*ServerConfig, error) {
	conf := &ServerConfig{}
	newFunc, ok := mgrList[svType]
	if !ok {
		return conf, errors.New("not find svType")
	}

	mgr, conf := newFunc(gs, svType, isFrontend)
	gs.mgr = mgr
	return conf, nil
}
