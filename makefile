GOCMD=go

define reload_server
	@echo "====================server===================="
	@cd kubernetes/server && kubectl delete -f deployment.yaml -n movie
	@cd kubernetes/server && kubectl create -f deployment.yaml -n movie
endef

define reload_auth
	@echo "====================auth===================="
	@cd kubernetes/auth && kubectl delete -f deployment.yaml -n movie
	@cd kubernetes/auth && kubectl create -f deployment.yaml -n movie
endef

define reload_talk
	@echo "====================talk===================="
	@cd kubernetes/talk && kubectl delete -f deployment.yaml -n movie
	@cd kubernetes/talk && kubectl create -f deployment.yaml -n movie
endef

define reload_movie
	@echo "====================movie===================="
	@cd kubernetes/movie && kubectl delete -f deployment.yaml -n movie
	@cd kubernetes/movie && kubectl create -f deployment.yaml -n movie
endef

define reload_job
	@echo "====================job===================="
	@cd kubernetes/job && kubectl delete -f cron-job-minute.yaml -n movie
	@cd kubernetes/job && kubectl delete -f cron-job.yaml -n movie
	@cd kubernetes/job && kubectl create -f cron-job-minute.yaml -n movie
	@cd kubernetes/job && kubectl create -f cron-job.yaml -n movie
endef


build:
	@echo "====================交叉编译===================="
	@echo "> make auth"
	@cd ./auth && ./generator.sh
	@cd ./auth && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOCMD) build -a -tags netgo -ldflags '-w' -o ../kubernetes/bin/auth ./server.go
	
	@echo "> make rpc_server"
	@cd ./crawler && ./generator.sh
	@cd ./crawler && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOCMD) build  -a -tags netgo -ldflags '-w' -o ../kubernetes/bin/server ./rpc_server.go

	@echo "> make rpc_gateway"
	@cd ./crawler && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOCMD) build  -a -tags netgo -ldflags '-w' -o ../kubernetes/bin/gateway ./gateway/server.go

	@echo "> make cronjob day"
	@cd ./crawler && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOCMD) build  -a -tags netgo -ldflags '-w' -o ../kubernetes/bin/day_job ./daemon/cmd/day/main.go

	@echo "> make cronjob minute"
	@cd ./crawler && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOCMD) build  -a -tags netgo -ldflags '-w' -o ../kubernetes/bin/minute_job ./daemon/cmd/minute/main.go
	
	@echo "> make movie"
	@cd ./movie && ./generator.sh
	@cd ./movie && cp -r ./static ../kubernetes/
	@cd ./movie && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOCMD) build  -a -tags netgo -ldflags '-w' -o ../kubernetes/bin/movie ./main.go

	@echo "> make talk"
	@cd ./talk && ./generator.sh
	@cd ./talk && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOCMD) build  -a -tags netgo -ldflags '-w' -o ../kubernetes/bin/talk ./main.go


remove-all:
	@kubectl delete ns movie

load-base:
	@echo "====================部署基础设施===================="
	@kubectl create ns movie

	# 设置自动注入
	@kubectl label ns movie istio-injection=enabled

	@cd kubernetes/base-facilities/mongo && kubectl create -f secret.yaml -n movie
	@cd kubernetes/base-facilities/mongo && kubectl create -f deployment.yaml -n movie
	@cd kubernetes/base-facilities/mongo && kubectl create -f service.yaml -n movie

	@cd kubernetes/ && kubectl create -f config.yaml -n movie

load:
	@echo "====================部署===================="

	@cd kubernetes/server && kubectl create -f deployment.yaml -n movie
	@cd kubernetes/server && kubectl create -f service.yaml -n movie
	@cd kubernetes/server && kubectl create -f vs.yaml -n movie

	@cd kubernetes/auth && kubectl create -f deployment.yaml -n movie
	@cd kubernetes/auth && kubectl create -f service.yaml -n movie
	@cd kubernetes/auth && kubectl create -f vs.yaml -n movie


	@cd kubernetes/talk && kubectl create -f deployment.yaml -n movie
	@cd kubernetes/talk && kubectl create -f service.yaml -n movie
	@cd kubernetes/talk && kubectl create -f vs.yaml -n movie
	@cd kubernetes/talk && kubectl create -f gateway.yaml -n movie


	@cd kubernetes/movie && kubectl create -f deployment.yaml -n movie
	@cd kubernetes/movie && kubectl create -f service.yaml -n movie
	@cd kubernetes/movie && kubectl create -f vs.yaml -n movie
	@cd kubernetes/movie && kubectl create -f gateway.yaml -n movie

	@cd kubernetes/ && kubectl create -f ingress.yaml -n movie


load-job-minute:
	@cd kubernetes/job && kubectl create -f cron-job-minute.yaml -n movie

load-job-day:
	@cd kubernetes/job && kubectl create -f cron-job.yaml -n movie

reload:
	$(reload_server)
	$(reload_auth)
	$(reload_talk)
	$(reload_movie)
	$(reload_job)

reload-server:
	$(reload_server)
reload-auth:
	$(reload_auth)
reload-talk:
	$(reload_talk)
reload-movie:
	$(reload_movie)
reload-job:
	$(reload_job)

status:
	@kubectl get pods -n movie


dashboard:
	@cd kubernetes/base-facilities && kubectl apply -f k8s-dashboard.yaml
	@cd kubernetes/base-facilities && kubectl apply -f kiali-dashboard.yaml
	@cd kubernetes/base-facilities && kubectl apply -f ingress.yaml
