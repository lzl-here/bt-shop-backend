// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v5.29.2
// source: pay_server.proto

// protobuf的package，根据业务来定，一个文件夹下的package得是一样的

package pay

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

var File_pay_server_proto protoreflect.FileDescriptor

var file_pay_server_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x61, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x03, 0x70, 0x61, 0x79, 0x1a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0d, 0x70, 0x61, 0x79, 0x2f, 0x70, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x11, 0x70, 0x61, 0x79, 0x2f, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x32, 0x9a, 0x01, 0x0a, 0x0e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x30, 0x0a, 0x06, 0x50, 0x72, 0x65, 0x70, 0x61,
	0x79, 0x12, 0x0b, 0x2e, 0x70, 0x61, 0x79, 0x2e, 0x50, 0x61, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x0b,
	0x2e, 0x70, 0x61, 0x79, 0x2e, 0x50, 0x61, 0x79, 0x52, 0x73, 0x70, 0x22, 0x0c, 0xd2, 0xc1, 0x18,
	0x08, 0x2f, 0x70, 0x61, 0x79, 0x2f, 0x70, 0x61, 0x79, 0x12, 0x56, 0x0a, 0x0d, 0x41, 0x6c, 0x69,
	0x70, 0x61, 0x79, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x12, 0x15, 0x2e, 0x70, 0x61, 0x79,
	0x2e, 0x41, 0x6c, 0x69, 0x70, 0x61, 0x79, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x52, 0x65,
	0x71, 0x1a, 0x15, 0x2e, 0x70, 0x61, 0x79, 0x2e, 0x41, 0x6c, 0x69, 0x70, 0x61, 0x79, 0x57, 0x65,
	0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x52, 0x73, 0x70, 0x22, 0x17, 0xd2, 0xc1, 0x18, 0x13, 0x2f, 0x70,
	0x61, 0x79, 0x2f, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x2f, 0x61, 0x6c, 0x69, 0x70, 0x61,
	0x79, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6c, 0x7a, 0x6c, 0x2d, 0x68, 0x65, 0x72, 0x65, 0x2f, 0x62, 0x74, 0x2d, 0x73, 0x68, 0x6f, 0x70,
	0x2d, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x62, 0x69, 0x7a, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f,
	0x70, 0x61, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_pay_server_proto_goTypes = []interface{}{
	(*PayReq)(nil),           // 0: pay.PayReq
	(*AlipayWebhookReq)(nil), // 1: pay.AlipayWebhookReq
	(*PayRsp)(nil),           // 2: pay.PayRsp
	(*AlipayWebhookRsp)(nil), // 3: pay.AlipayWebhookRsp
}
var file_pay_server_proto_depIdxs = []int32{
	0, // 0: pay.ExampleService.Prepay:input_type -> pay.PayReq
	1, // 1: pay.ExampleService.AlipayWebhook:input_type -> pay.AlipayWebhookReq
	2, // 2: pay.ExampleService.Prepay:output_type -> pay.PayRsp
	3, // 3: pay.ExampleService.AlipayWebhook:output_type -> pay.AlipayWebhookRsp
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pay_server_proto_init() }
func file_pay_server_proto_init() {
	if File_pay_server_proto != nil {
		return
	}
	file_pay_pay_proto_init()
	file_pay_webhook_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pay_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pay_server_proto_goTypes,
		DependencyIndexes: file_pay_server_proto_depIdxs,
	}.Build()
	File_pay_server_proto = out.File
	file_pay_server_proto_rawDesc = nil
	file_pay_server_proto_goTypes = nil
	file_pay_server_proto_depIdxs = nil
}
