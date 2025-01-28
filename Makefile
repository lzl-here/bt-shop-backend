# 定义变量
PROTO_PATH=protobuf/proto
API_PATH=protobuf/api
GEN_PATH=kitex_gen
MODULE_NAME=github.com/lzl-here/bt-shop-backend

# 生成proto代码
gen:
	@ make clean
	@ kitex -module $(MODULE_NAME) -type=protobuf -I $(PROTO_PATH) $(PROTO_PATH)/example/example_server.proto
	@ kitex -module $(MODULE_NAME) -type=protobuf -I $(PROTO_PATH) $(PROTO_PATH)/define/base_response.proto
	@ kitex -module $(MODULE_NAME) -type=protobuf -I $(PROTO_PATH) $(PROTO_PATH)/user/user.proto
	@ kitex -module $(MODULE_NAME) -type=protobuf -I $(PROTO_PATH) $(PROTO_PATH)/pay/pay_server.proto
	@ kitex -module $(MODULE_NAME) -type=protobuf -I $(PROTO_PATH) $(PROTO_PATH)/order/order_server.proto
	@ kitex -module $(MODULE_NAME) -type=protobuf -I $(PROTO_PATH) $(PROTO_PATH)/goods/goods_server.proto




	@ cd apps/gateway  && hz update -I ../../$(API_PATH) -idl ../../$(API_PATH)/user/user.proto
	@ cd apps/gateway  && hz update -I ../../$(API_PATH) -idl ../../$(API_PATH)/pay/api_pay_server.proto


	@ go mod tidy

# 清空proto代码
clean:
	@ rm -rf $(GEN_PATH)/*

.PHONY: gen clean