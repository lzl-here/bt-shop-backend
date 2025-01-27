// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v5.29.2
// source: pay/prepay.proto

// protobuf的package，根据业务来定，一个文件夹下的package得是一样的

package pay

import (
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
type PrepayReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PrepayReq) Reset() {
	*x = PrepayReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pay_prepay_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrepayReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrepayReq) ProtoMessage() {}

func (x *PrepayReq) ProtoReflect() protoreflect.Message {
	mi := &file_pay_prepay_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrepayReq.ProtoReflect.Descriptor instead.
func (*PrepayReq) Descriptor() ([]byte, []int) {
	return file_pay_prepay_proto_rawDescGZIP(), []int{0}
}

type PrepayRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code  int32                    `protobuf:"varint,1,opt,name=code,proto3" form:"code" json:"code,omitempty" query:"code"`                   // 错误码
	Msg   string                   `protobuf:"bytes,2,opt,name=msg,proto3" form:"msg" json:"msg,omitempty" query:"msg"`                        // 信息
	LogId string                   `protobuf:"bytes,3,opt,name=log_id,json=logId,proto3" form:"log_id" json:"log_id,omitempty" query:"log_id"` // 链路追踪id
	Data  *PrepayRsp_PrepayRspData `protobuf:"bytes,4,opt,name=data,proto3" form:"data" json:"data,omitempty" query:"data"`                    // 响应体数据
}

func (x *PrepayRsp) Reset() {
	*x = PrepayRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pay_prepay_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrepayRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrepayRsp) ProtoMessage() {}

func (x *PrepayRsp) ProtoReflect() protoreflect.Message {
	mi := &file_pay_prepay_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrepayRsp.ProtoReflect.Descriptor instead.
func (*PrepayRsp) Descriptor() ([]byte, []int) {
	return file_pay_prepay_proto_rawDescGZIP(), []int{1}
}

func (x *PrepayRsp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *PrepayRsp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *PrepayRsp) GetLogId() string {
	if x != nil {
		return x.LogId
	}
	return ""
}

func (x *PrepayRsp) GetData() *PrepayRsp_PrepayRspData {
	if x != nil {
		return x.Data
	}
	return nil
}

type PrepayRsp_PrepayRspData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PrepayRsp_PrepayRspData) Reset() {
	*x = PrepayRsp_PrepayRspData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pay_prepay_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrepayRsp_PrepayRspData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrepayRsp_PrepayRspData) ProtoMessage() {}

func (x *PrepayRsp_PrepayRspData) ProtoReflect() protoreflect.Message {
	mi := &file_pay_prepay_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrepayRsp_PrepayRspData.ProtoReflect.Descriptor instead.
func (*PrepayRsp_PrepayRspData) Descriptor() ([]byte, []int) {
	return file_pay_prepay_proto_rawDescGZIP(), []int{1, 0}
}

var File_pay_prepay_proto protoreflect.FileDescriptor

var file_pay_prepay_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x61, 0x79, 0x2f, 0x70, 0x72, 0x65, 0x70, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x03, 0x70, 0x61, 0x79, 0x22, 0x0b, 0x0a, 0x09, 0x50, 0x72, 0x65, 0x70, 0x61,
	0x79, 0x52, 0x65, 0x71, 0x22, 0x8b, 0x01, 0x0a, 0x09, 0x50, 0x72, 0x65, 0x70, 0x61, 0x79, 0x52,
	0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x15, 0x0a, 0x06, 0x6c, 0x6f, 0x67, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x12,
	0x30, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x70, 0x61, 0x79, 0x2e, 0x50, 0x72, 0x65, 0x70, 0x61, 0x79, 0x52, 0x73, 0x70, 0x2e, 0x50, 0x72,
	0x65, 0x70, 0x61, 0x79, 0x52, 0x73, 0x70, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x1a, 0x0f, 0x0a, 0x0d, 0x50, 0x72, 0x65, 0x70, 0x61, 0x79, 0x52, 0x73, 0x70, 0x44, 0x61,
	0x74, 0x61, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6c, 0x7a, 0x6c, 0x2d, 0x68, 0x65, 0x72, 0x65, 0x2f, 0x62, 0x74, 0x2d, 0x73, 0x68, 0x6f,
	0x70, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x62, 0x69, 0x7a, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2f, 0x70, 0x61, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pay_prepay_proto_rawDescOnce sync.Once
	file_pay_prepay_proto_rawDescData = file_pay_prepay_proto_rawDesc
)

func file_pay_prepay_proto_rawDescGZIP() []byte {
	file_pay_prepay_proto_rawDescOnce.Do(func() {
		file_pay_prepay_proto_rawDescData = protoimpl.X.CompressGZIP(file_pay_prepay_proto_rawDescData)
	})
	return file_pay_prepay_proto_rawDescData
}

var file_pay_prepay_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pay_prepay_proto_goTypes = []interface{}{
	(*PrepayReq)(nil),               // 0: pay.PrepayReq
	(*PrepayRsp)(nil),               // 1: pay.PrepayRsp
	(*PrepayRsp_PrepayRspData)(nil), // 2: pay.PrepayRsp.PrepayRspData
}
var file_pay_prepay_proto_depIdxs = []int32{
	2, // 0: pay.PrepayRsp.data:type_name -> pay.PrepayRsp.PrepayRspData
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pay_prepay_proto_init() }
func file_pay_prepay_proto_init() {
	if File_pay_prepay_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pay_prepay_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrepayReq); i {
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
		file_pay_prepay_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrepayRsp); i {
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
		file_pay_prepay_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrepayRsp_PrepayRspData); i {
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
			RawDescriptor: file_pay_prepay_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pay_prepay_proto_goTypes,
		DependencyIndexes: file_pay_prepay_proto_depIdxs,
		MessageInfos:      file_pay_prepay_proto_msgTypes,
	}.Build()
	File_pay_prepay_proto = out.File
	file_pay_prepay_proto_rawDesc = nil
	file_pay_prepay_proto_goTypes = nil
	file_pay_prepay_proto_depIdxs = nil
}
