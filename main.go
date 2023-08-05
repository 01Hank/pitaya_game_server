package main

import (
	"fmt"
	"flag"
	"strconv"

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
		module_mgr *ModuleManager
	}
)

func (gs *GameServer) start() {
	defer gs.app.Shutdown()

	//注册所有modules
	gs.module_mgr.RegisterToPitaya(gs.app)

	gs.app.Start()
}

var app pitaya.Pitaya
var game_server *GameServer

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

	game_server = &GameServer{
		module_mgr : InitModules(),
		app : app,
	}

	game_server.start()
}

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