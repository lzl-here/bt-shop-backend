PROTO_PATH=proto
GEN_PATH=kitex_gen
MODULE_NAME=github.com/lzl-here/bt-shop-backend
ROOT_DIR=$(shell cd .. && pwd)
PROTOBUF_DIR=.
DOC_DIR=$(PROTOBUF_DIR)/doc

# 生成 grpc 代码
gen:
	@ make clean
	@ cd $(PROTOBUF_DIR) && \
		kitex -module $(MODULE_NAME) -type=protobuf -I $(PROTO_PATH) -gen-path $(GEN_PATH) $(PROTO_PATH)/pay/pay_server.proto && \
		kitex -module $(MODULE_NAME) -type=protobuf -I $(PROTO_PATH) -gen-path $(GEN_PATH) $(PROTO_PATH)/order/order_server.proto && \
		kitex -module $(MODULE_NAME) -type=protobuf -I $(PROTO_PATH) -gen-path $(GEN_PATH) $(PROTO_PATH)/goods/goods_server.proto && \
		kitex -module $(MODULE_NAME) -type=protobuf -I $(PROTO_PATH) -gen-path $(GEN_PATH) $(PROTO_PATH)/user/user_server.proto

# 生成网关代码
api:
	@ rm -rf $(ROOT_DIR)/apps/gateway/biz/rpc/*
	@ cd $(ROOT_DIR)/apps/gateway && hz update \
		-I ../../protobuf/$(PROTO_PATH) \
		-I ../../protobuf/third_party \
		-idl ../../protobuf/$(PROTO_PATH)/order/order_server.proto \
		--handler_dir=biz/handler \
		--use $(MODULE_NAME)/protobuf/$(GEN_PATH) \
		--client_dir=biz/rpc \
		--no_recurse
		
	@ cd $(ROOT_DIR)/apps/gateway && hz update \
		-I ../../protobuf/$(PROTO_PATH) \
		-I ../../protobuf/third_party \
		-idl ../../protobuf/$(PROTO_PATH)/goods/goods_server.proto \
		--handler_dir=biz/handler \
		--use $(MODULE_NAME)/protobuf/$(GEN_PATH) \
		--client_dir=biz/rpc \
		--no_recurse

	@ cd $(ROOT_DIR)/apps/gateway && hz update \
		-I ../../protobuf/$(PROTO_PATH) \
		-I ../../protobuf/third_party \
		-idl ../../protobuf/$(PROTO_PATH)/pay/pay_server.proto \
		--handler_dir=biz/handler \
		--use $(MODULE_NAME)/protobuf/$(GEN_PATH) \
		--client_dir=biz/rpc \
		--no_recurse
		
	@ cd $(ROOT_DIR)/apps/gateway && hz update \
		-I ../../protobuf/$(PROTO_PATH) \
		-I ../../protobuf/third_party \
		-idl ../../protobuf/$(PROTO_PATH)/user/user_server.proto \
		--handler_dir=biz/handler \
		--use $(MODULE_NAME)/protobuf/$(GEN_PATH) \
		--client_dir=biz/rpc \
		--no_recurse

# TODO 生成 swagger 文档，目前只有数据结构，没有接口
swagger:
	@ mkdir -p $(DOC_DIR)
	@ cd $(PROTOBUF_DIR) && \
		protoc -I $(PROTO_PATH) \
		-I . \
		--swagger_out=logtostderr=true:$(DOC_DIR) \
		$(PROTO_PATH)/goods/goods_server.proto

	@ cd $(PROTOBUF_DIR) && \
		protoc -I $(PROTO_PATH) \
		-I . \
		--swagger_out=logtostderr=true:$(DOC_DIR) \
		$(PROTO_PATH)/order/order_server.proto

	@ cd $(PROTOBUF_DIR) && \
		protoc -I $(PROTO_PATH) \
		-I . \
		--swagger_out=logtostderr=true:$(DOC_DIR) \
		$(PROTO_PATH)/pay/pay_server.proto

	@ cd $(PROTOBUF_DIR) && \
		protoc -I $(PROTO_PATH) \
		-I . \
		--swagger_out=logtostderr=true:$(DOC_DIR) \
		$(PROTO_PATH)/user/user_server.proto

# 清空生成的代码
clean:
	@ rm -rf $(PROTOBUF_DIR)/$(GEN_PATH)/*
	@ mkdir -p $(PROTOBUF_DIR)/$(GEN_PATH)

.PHONY: gen clean gateway swagger