# 编译proto文件
run-proto:
	@protoc -I  protofile/ protofile/*.proto --go_out=.

# 网关服务
run-gate:
	@go run *.go  --port 3251 --rpcsvport 3435 --type gateService --frontend=true

# db服务
run-db:
	@go run *.go --rpcsvport 3436 --type dbService --frontend=false