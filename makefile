GOCMD=go

define reload_server
	@echo "====================server===================="
	@cd kubernetes/server && kubectl delete -f deployment.yaml -n movie
	@docker rmi -f server
	@cd kubernetes/server && docker build -t server .
	@cd kubernetes/server && kubectl create -f deployment.yaml -n movie
endef

define reload_auth
	@echo "====================auth===================="
	@cd kubernetes/auth && kubectl delete -f deployment.yaml -n movie
	@docker rmi -f auth
	@cd kubernetes/auth && docker build -t auth .
	@cd kubernetes/auth && kubectl create -f deployment.yaml -n movie
endef

define reload_talk
	@echo "====================talk===================="
	@cd kubernetes/talk && kubectl delete -f deployment.yaml -n movie
	@docker rmi -f talk
	@cd kubernetes/talk && docker build -t talk .
	@cd kubernetes/talk && kubectl create -f deployment.yaml -n movie
endef

define reload_movie
	@echo "====================movie===================="
	@cd kubernetes/movie && kubectl delete -f deployment.yaml -n movie
	@docker rmi -f movie
	@cd kubernetes/movie && docker build -t movie .
	@cd kubernetes/movie && kubectl create -f deployment.yaml -n movie
endef

define reload_job
	@echo "====================job===================="
	@cd kubernetes/job && kubectl delete -f cron-job-minute.yaml -n movie
	@cd kubernetes/job && kubectl delete -f cron-job.yaml -n movie
	@docker rmi -f movie-job
	@cd kubernetes/job && docker build -t movie-job .
	@cd kubernetes/job && kubectl create -f cron-job-minute.yaml -n movie
	@cd kubernetes/job && kubectl create -f cron-job.yaml -n movie
endef


build:
	@echo "====================交叉编译===================="
	@echo "> make auth"
	@cd ./auth && ./generator.sh
	@cd ./auth && $(GOCMD) mod tidy
	@cd ./auth && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOCMD) build -a -tags netgo -ldflags '-w' -o ../kubernetes/auth/auth ./server.go
	
	@echo "> make rpc_server"
	@cd ./crawler && ./generator.sh
	@cd ./crawler && $(GOCMD) mod tidy
	@cd ./crawler && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOCMD) build  -a -tags netgo -ldflags '-w' -o ../kubernetes/server/server ./rpc_server.go

	@echo "> make rpc_gateway"
	@cd ./crawler && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOCMD) build  -a -tags netgo -ldflags '-w' -o ../kubernetes/server/gateway ./gateway/server.go

	@echo "> make cronjob day"
	@cd ./crawler && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOCMD) build  -a -tags netgo -ldflags '-w' -o ../kubernetes/job/day_job ./daemon/cmd/day/main.go

	@echo "> make cronjob minute"
	@cd ./crawler && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOCMD) build  -a -tags netgo -ldflags '-w' -o ../kubernetes/job/minute_job ./daemon/cmd/minute/main.go
	
	@echo "> make movie"
	@cd ./movie && ./generator.sh
	@cd ./movie && $(GOCMD) mod tidy
	@cd ./movie && cp -r ./static ../kubernetes/movie/
	@cd ./movie && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOCMD) build  -a -tags netgo -ldflags '-w' -o ../kubernetes/movie/movie ./main.go

	@echo "> make talk"
	@cd ./talk && ./generator.sh
	@cd ./talk && $(GOCMD) mod tidy
	@cd ./talk && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOCMD) build  -a -tags netgo -ldflags '-w' -o ../kubernetes/talk/talk ./main.go

clean:

	@docker rmi -f auth
	@docker rmi -f server
	@docker rmi -f movie
	@docker rmi -f talk
	@docker rmi -f movie-job

load:
	@echo "====================部署===================="

	@kubectl create ns movie

	# 设置自动注入
	@kubectl label ns movie istio-injection=enabled

	@cd kubernetes/base-facilities/mongo && kubectl create -f secret.yaml -n movie
	@cd kubernetes/base-facilities/mongo && kubectl create -f deployment.yaml -n movie
	@cd kubernetes/base-facilities/mongo && kubectl create -f service.yaml -n movie

	@kubectl create -f config.yaml -n movie

	@cd kubernetes/server && docker build -t server .
	@cd kubernetes/server && kubectl create -f deployment.yaml -n movie
	@cd kubernetes/server && kubectl create -f service.yaml -n movie
	@cd kubernetes/server && kubectl create -f vs.yaml -n movie

	@cd kubernetes/auth && docker build -t auth .
	@cd kubernetes/auth && kubectl create -f deployment.yaml -n movie
	@cd kubernetes/auth && kubectl create -f service.yaml -n movie
	@cd kubernetes/auth && kubectl create -f vs.yaml -n movie


	@cd kubernetes/talk && docker build -t talk .
	@cd kubernetes/talk && kubectl create -f deployment.yaml -n movie
	@cd kubernetes/talk && kubectl create -f service.yaml -n movie
	@cd kubernetes/talk && kubectl create -f vs.yaml -n movie
	@cd kubernetes/talk && kubectl create -f gateway.yaml -n movie


	@cd kubernetes/movie && docker build -t movie .
	@cd kubernetes/movie && kubectl create -f deployment.yaml -n movie
	@cd kubernetes/movie && kubectl create -f service.yaml -n movie
	@cd kubernetes/movie && kubectl create -f vs.yaml -n movie
	@cd kubernetes/movie && kubectl create -f gateway.yaml -n movie

	@cd kubernetes/job && docker build -t movie-job .
	@cd kubernetes/job && kubectl create -f cron-job-minute.yaml -n movie
	@cd kubernetes/job && kubectl create -f cron-job.yaml -n movie

reload:
	$(reload_server)
	$(reload_auth)
	$(reload_talk)
	$(reload_movie)
	$(reload_job)

reload_server:
	$(reload_server)
reload_auth:
	$(reload_auth)
reload_talk:
	$(reload_talk)
reload_movie:
	$(reload_movie)
reload_job:
	$(reload_job)
