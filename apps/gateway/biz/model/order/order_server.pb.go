// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v5.29.2
// source: order_server.proto

// protobuf的package，根据业务来定，一个文件夹下的package得是一样的

package order

import (
	_ "github.com/lzl-here/bt-shop-backend/apps/gateway/biz/model/api"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_order_server_proto protoreflect.FileDescriptor

var file_order_server_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x1a, 0x09, 0x61, 0x70, 0x69,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x74, 0x72,
	0x61, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x11, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xbb, 0x04, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x72, 0x61, 0x64, 0x65, 0x12, 0x15, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x64, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x64, 0x65,
	0x52, 0x73, 0x70, 0x22, 0x17, 0xd2, 0xc1, 0x18, 0x13, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x64, 0x65, 0x12, 0x44, 0x0a, 0x07,
	0x52, 0x65, 0x54, 0x72, 0x61, 0x64, 0x65, 0x12, 0x11, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e,
	0x52, 0x65, 0x54, 0x72, 0x61, 0x64, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x2e, 0x52, 0x65, 0x54, 0x72, 0x61, 0x64, 0x65, 0x52, 0x73, 0x70, 0x22, 0x13, 0xd2,
	0xc1, 0x18, 0x0f, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x72, 0x65, 0x5f, 0x74, 0x72, 0x61,
	0x64, 0x65, 0x12, 0x54, 0x0a, 0x0b, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x54, 0x72, 0x61, 0x64,
	0x65, 0x12, 0x15, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c,
	0x54, 0x72, 0x61, 0x64, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x54, 0x72, 0x61, 0x64, 0x65, 0x52, 0x73, 0x70, 0x22,
	0x17, 0xd2, 0xc1, 0x18, 0x13, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x63, 0x61, 0x6e, 0x63,
	0x65, 0x6c, 0x5f, 0x74, 0x72, 0x61, 0x64, 0x65, 0x12, 0x5d, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x17, 0x2e, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52,
	0x65, 0x71, 0x1a, 0x17, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x73, 0x70, 0x22, 0x1a, 0xca, 0xc1, 0x18,
	0x16, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x6e, 0x0a, 0x11, 0x50, 0x61, 0x79, 0x53, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x1b, 0x2e, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x2e, 0x50, 0x61, 0x79, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54,
	0x6f, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x2e, 0x50, 0x61, 0x79, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x73, 0x70, 0x22, 0x1f, 0xd2, 0xc1, 0x18, 0x1b, 0x2f, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x2f, 0x70, 0x61, 0x79, 0x5f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74,
	0x6f, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x6a, 0x0a, 0x10, 0x50, 0x61, 0x79, 0x43, 0x61,
	0x6e, 0x63, 0x65, 0x6c, 0x54, 0x6f, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x1a, 0x2e, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x2e, 0x50, 0x61, 0x79, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x54, 0x6f, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e,
	0x50, 0x61, 0x79, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x54, 0x6f, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x52, 0x73, 0x70, 0x22, 0x1e, 0xd2, 0xc1, 0x18, 0x1a, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f,
	0x70, 0x61, 0x79, 0x5f, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x5f, 0x74, 0x6f, 0x5f, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6c, 0x7a, 0x6c, 0x2d, 0x68, 0x65, 0x72, 0x65, 0x2f, 0x62, 0x74, 0x2d, 0x73, 0x68,
	0x6f, 0x70, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x62, 0x69, 0x7a, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_order_server_proto_goTypes = []interface{}{
	(*CreateTradeReq)(nil),       // 0: order.CreateTradeReq
	(*ReTradeReq)(nil),           // 1: order.ReTradeReq
	(*CancelTradeReq)(nil),       // 2: order.CancelTradeReq
	(*GetOrderItemsReq)(nil),     // 3: order.GetOrderItemsReq
	(*PaySuccessToOrderReq)(nil), // 4: order.PaySuccessToOrderReq
	(*PayCancelToOrderReq)(nil),  // 5: order.PayCancelToOrderReq
	(*CreateTradeRsp)(nil),       // 6: order.CreateTradeRsp
	(*ReTradeRsp)(nil),           // 7: order.ReTradeRsp
	(*CancelTradeRsp)(nil),       // 8: order.CancelTradeRsp
	(*GetOrderItemsRsp)(nil),     // 9: order.GetOrderItemsRsp
	(*PaySuccessToOrderRsp)(nil), // 10: order.PaySuccessToOrderRsp
	(*PayCancelToOrderRsp)(nil),  // 11: order.PayCancelToOrderRsp
}
var file_order_server_proto_depIdxs = []int32{
	0,  // 0: order.OrderService.CreateTrade:input_type -> order.CreateTradeReq
	1,  // 1: order.OrderService.ReTrade:input_type -> order.ReTradeReq
	2,  // 2: order.OrderService.CancelTrade:input_type -> order.CancelTradeReq
	3,  // 3: order.OrderService.GetOrderItems:input_type -> order.GetOrderItemsReq
	4,  // 4: order.OrderService.PaySuccessToOrder:input_type -> order.PaySuccessToOrderReq
	5,  // 5: order.OrderService.PayCancelToOrder:input_type -> order.PayCancelToOrderReq
	6,  // 6: order.OrderService.CreateTrade:output_type -> order.CreateTradeRsp
	7,  // 7: order.OrderService.ReTrade:output_type -> order.ReTradeRsp
	8,  // 8: order.OrderService.CancelTrade:output_type -> order.CancelTradeRsp
	9,  // 9: order.OrderService.GetOrderItems:output_type -> order.GetOrderItemsRsp
	10, // 10: order.OrderService.PaySuccessToOrder:output_type -> order.PaySuccessToOrderRsp
	11, // 11: order.OrderService.PayCancelToOrder:output_type -> order.PayCancelToOrderRsp
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_order_server_proto_init() }
func file_order_server_proto_init() {
	if File_order_server_proto != nil {
		return
	}
	file_order_trade_proto_init()
	file_order_order_items_proto_init()
	file_order_order_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_order_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_order_server_proto_goTypes,
		DependencyIndexes: file_order_server_proto_depIdxs,
	}.Build()
	File_order_server_proto = out.File
	file_order_server_proto_rawDesc = nil
	file_order_server_proto_goTypes = nil
	file_order_server_proto_depIdxs = nil
}
