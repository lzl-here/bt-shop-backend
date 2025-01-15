// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.29.2
// source: person/person_server.proto

// protobuf的package，根据业务来定，一个文件夹下的package得是一样的

package person

import (
	context "context"
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

var File_person_person_server_proto protoreflect.FileDescriptor

var file_person_person_server_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x1a, 0x18, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2f, 0x70, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x4e,
	0x0a, 0x0d, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x3d, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x14, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e,
	0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x73, 0x70, 0x22, 0x00, 0x42, 0x36,
	0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x7a, 0x6c,
	0x2d, 0x68, 0x65, 0x72, 0x65, 0x2f, 0x62, 0x74, 0x2d, 0x73, 0x68, 0x6f, 0x70, 0x2d, 0x62, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f,
	0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_person_person_server_proto_goTypes = []interface{}{
	(*GetPersonReq)(nil), // 0: person.GetPersonReq
	(*GetPersonRsp)(nil), // 1: person.GetPersonRsp
}
var file_person_person_server_proto_depIdxs = []int32{
	0, // 0: person.PersonService.GetPersonInfo:input_type -> person.GetPersonReq
	1, // 1: person.PersonService.GetPersonInfo:output_type -> person.GetPersonRsp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_person_person_server_proto_init() }
func file_person_person_server_proto_init() {
	if File_person_person_server_proto != nil {
		return
	}
	file_person_person_info_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_person_person_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_person_person_server_proto_goTypes,
		DependencyIndexes: file_person_person_server_proto_depIdxs,
	}.Build()
	File_person_person_server_proto = out.File
	file_person_person_server_proto_rawDesc = nil
	file_person_person_server_proto_goTypes = nil
	file_person_person_server_proto_depIdxs = nil
}

var _ context.Context

// Code generated by Kitex v0.12.1. DO NOT EDIT.

type PersonService interface {
	GetPersonInfo(ctx context.Context, req *GetPersonReq) (res *GetPersonRsp, err error)
}