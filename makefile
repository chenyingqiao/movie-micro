build:
	@echo "> make auth"
	@cd ./auth && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w' -o ../kubernetes/auth/auth ./server.go
	
	@echo "> make rpc_server"
	@cd ./crawler && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -a -tags netgo -ldflags '-w' -o ../kubernetes/server/server ./rpc_server.go
	@echo "> make rpc_gateway"
	@cd ./crawler && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -a -tags netgo -ldflags '-w' -o ../kubernetes/server/gateway ./gateway/server.go
	@echo "> make cronjob day"
	@cd ./crawler && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -a -tags netgo -ldflags '-w' -o ../kubernetes/job/day_job ./daemon/cmd/day/main.go
	@echo "> make cronjob minute"
	@cd ./crawler && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -a -tags netgo -ldflags '-w' -o ../kubernetes/job/minute_job ./daemon/cmd/minute/main.go
	
	@echo "> make movie"
	@cd ./movie && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -a -tags netgo -ldflags '-w' -o ../kubernetes/movie/movie ./main.go

	@echo "> make talk"
	@cd ./talk && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -a -tags netgo -ldflags '-w' -o ../kubernetes/talk/talk ./main.go
