package main

import (
	"fmt"
	"flag"
	"strconv"

	"pitaya_game_server/test_module"

	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/acceptor"
	"github.com/topfreegames/pitaya/v2/config"
	"github.com/topfreegames/pitaya/v2/groups"
	"github.com/topfreegames/pitaya/v2/cluster"
	"github.com/topfreegames/pitaya/v2/modules"
	"github.com/topfreegames/pitaya/v2/constants"
)

type (
	GameServer struct {
		app pitaya.Pitaya
		service_mgr *ServiceManager //服务管理
		tm *test_module.TestModule
	}

	ServerConfig struct {
		exclude_modules  []string //不启动的组件
		exclude_services []string //不启动的服务
	}
)

func (gs *GameServer) start(conf *ServerConfig) {
	defer gs.app.Shutdown()

	gs.app.RegisterModule(gs.tm, "test_mod")

	//注册所有services
	gs.service_mgr.RegisterServices(gs, conf.exclude_services)

	gs.app.Start()
}

var app pitaya.Pitaya
var game_server GameServer

func main()  {
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

	game_server = GameServer{
		app : app,
		service_mgr : InitServices(),
		tm : test_module.NewTestModule(),
	}

	conf := serverConf()
	game_server.start(conf)
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

// 获取服务器配置
func serverConf() *ServerConfig {
	return &ServerConfig{}
}