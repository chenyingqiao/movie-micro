build:
	@echo "> make rpc_server"
	@go build -o ./bin/server ./rpc_server.go
	@echo "> make rpc_gateway"
	@go build -o ./bin/gateway ./gateway/server.go
	@echo "> make cronjob day"
	@go build -o ./bin/day_job ./daemon/cmd/day/main.go
	@echo "> make cronjob minute"
	@go build -o ./bin/minute_job ./daemon/cmd/minute/main.go