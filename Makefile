# 定义变量
PROTO_PATH=protobuf/proto
GEN_PATH=kitex_gen
MODULE_NAME=github.com/lzl-here/bt-shop-backend




# 生成proto代码
gen:
	@ make clean
	@ kitex -module $(MODULE_NAME) -type=protobuf -I $(PROTO_PATH) $(PROTO_PATH)/example/example_server.proto
	@ go mod tidy

# 清空proto代码
clean:
	@ rm -rf $(GEN_PATH)/*





.PHONY: gen clean