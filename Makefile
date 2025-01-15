# 定义变量
PROTO_PATH=protobuf/proto
GEN_PATH=kitex_gen
MODULE_NAME=github.com/lzl-here/bt-shop-backend

# 服务路径
EXAMPLE_PATH=services/example
EXAMPLE_NAME=example



# 生成proto代码
gen:
	@ make clean
	@ kitex -module $(MODULE_NAME) -type=protobuf -I $(PROTO_PATH) $(PROTO_PATH)/person/person_server.proto
	@ go mod tidy

# 清空proto代码
clean:
	@ rm -rf $(GEN_PATH)/*

# 本地构建运行服务
example-dev: 
	@ go build -o $(EXAMPLE_PATH)/bin/$(EXAMPLE_NAME) $(EXAMPLE_PATH)/cmd/main.go
	@ ./$(EXAMPLE_PATH)/bin/$(EXAMPLE_NAME)  -cfgFile=$(EXAMPLE_PATH)/.env.production



.PHONY: gen clean