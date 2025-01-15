# 服务路径
SERVICE_NAME=example



# 本地构建运行服务
dev: 
	@ go build -o ./bin/$(SERVICE_NAME) ./cmd/main.go
	@ cd ./bin && ./$(SERVICE_NAME)  -cfgFile=../.env.production

