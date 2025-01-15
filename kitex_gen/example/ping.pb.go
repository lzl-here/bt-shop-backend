// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.29.2
// source: example/ping.proto

// protobuf的package，根据业务来定，一个文件夹下的package得是一样的

package example

import (
	context "context"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 请求体: XXXReq
type PingReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid uint64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *PingReq) Reset() {
	*x = PingReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_ping_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingReq) ProtoMessage() {}

func (x *PingReq) ProtoReflect() protoreflect.Message {
	mi := &file_example_ping_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingReq.ProtoReflect.Descriptor instead.
func (*PingReq) Descriptor() ([]byte, []int) {
	return file_example_ping_proto_rawDescGZIP(), []int{0}
}

func (x *PingReq) GetUid() uint64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

// 响应体: XXXRsp
// 结构统一长这样: 错误码、错误信息、日志id、数据
type PingRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code  int32             `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`               // 错误码
	Msg   string            `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`                  // 信息
	LogId string            `protobuf:"bytes,3,opt,name=log_id,json=logId,proto3" json:"log_id,omitempty"` // 链路追踪id
	Data  *PingRsp_PingData `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`                // 响应体数据
}

func (x *PingRsp) Reset() {
	*x = PingRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_ping_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingRsp) ProtoMessage() {}

func (x *PingRsp) ProtoReflect() protoreflect.Message {
	mi := &file_example_ping_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingRsp.ProtoReflect.Descriptor instead.
func (*PingRsp) Descriptor() ([]byte, []int) {
	return file_example_ping_proto_rawDescGZIP(), []int{1}
}

func (x *PingRsp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *PingRsp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *PingRsp) GetLogId() string {
	if x != nil {
		return x.LogId
	}
	return ""
}

func (x *PingRsp) GetData() *PingRsp_PingData {
	if x != nil {
		return x.Data
	}
	return nil
}

type PingRsp_PingData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Message   string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	CreatedAt string `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt string `protobuf:"bytes,5,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *PingRsp_PingData) Reset() {
	*x = PingRsp_PingData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_ping_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingRsp_PingData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingRsp_PingData) ProtoMessage() {}

func (x *PingRsp_PingData) ProtoReflect() protoreflect.Message {
	mi := &file_example_ping_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingRsp_PingData.ProtoReflect.Descriptor instead.
func (*PingRsp_PingData) Descriptor() ([]byte, []int) {
	return file_example_ping_proto_rawDescGZIP(), []int{1, 0}
}

func (x *PingRsp_PingData) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PingRsp_PingData) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *PingRsp_PingData) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *PingRsp_PingData) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *PingRsp_PingData) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

var File_example_ping_proto protoreflect.FileDescriptor

var file_example_ping_proto_rawDesc = []byte{
	0x0a, 0x12, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x22, 0x1b, 0x0a,
	0x07, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x89, 0x02, 0x0a, 0x07, 0x50,
	0x69, 0x6e, 0x67, 0x52, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x15, 0x0a, 0x06,
	0x6c, 0x6f, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f,
	0x67, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x50, 0x69, 0x6e, 0x67,
	0x52, 0x73, 0x70, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x1a, 0x91, 0x01, 0x0a, 0x08, 0x50, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x7a, 0x6c, 0x2d, 0x68, 0x65, 0x72, 0x65, 0x2f, 0x62, 0x74,
	0x2d, 0x73, 0x68, 0x6f, 0x70, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x6b, 0x69,
	0x74, 0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_example_ping_proto_rawDescOnce sync.Once
	file_example_ping_proto_rawDescData = file_example_ping_proto_rawDesc
)

func file_example_ping_proto_rawDescGZIP() []byte {
	file_example_ping_proto_rawDescOnce.Do(func() {
		file_example_ping_proto_rawDescData = protoimpl.X.CompressGZIP(file_example_ping_proto_rawDescData)
	})
	return file_example_ping_proto_rawDescData
}

var file_example_ping_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_example_ping_proto_goTypes = []interface{}{
	(*PingReq)(nil),          // 0: example.PingReq
	(*PingRsp)(nil),          // 1: example.PingRsp
	(*PingRsp_PingData)(nil), // 2: example.PingRsp.PingData
}
var file_example_ping_proto_depIdxs = []int32{
	2, // 0: example.PingRsp.data:type_name -> example.PingRsp.PingData
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_example_ping_proto_init() }
func file_example_ping_proto_init() {
	if File_example_ping_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_example_ping_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_example_ping_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingRsp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_example_ping_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingRsp_PingData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_example_ping_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_example_ping_proto_goTypes,
		DependencyIndexes: file_example_ping_proto_depIdxs,
		MessageInfos:      file_example_ping_proto_msgTypes,
	}.Build()
	File_example_ping_proto = out.File
	file_example_ping_proto_rawDesc = nil
	file_example_ping_proto_goTypes = nil
	file_example_ping_proto_depIdxs = nil
}

var _ context.Context