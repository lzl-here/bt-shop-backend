# 服务路径
SERVICE_NAME=gateway

build: 
	@ go build -o ./bin/$(SERVICE_NAME) ./.

# 本地构建运行服务
dev: 
	@ make build
	@ cd ./bin && ./$(SERVICE_NAME)  -cfgFile=../.env.production
