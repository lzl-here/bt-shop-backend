// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.29.2
// source: order/create_trade.proto

package order

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

// 下单相关
type CreateTradeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TradeNo string `protobuf:"bytes,1,opt,name=trade_no,json=tradeNo,proto3" json:"trade_no,omitempty"`
}

func (x *CreateTradeReq) Reset() {
	*x = CreateTradeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_create_trade_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTradeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTradeReq) ProtoMessage() {}

func (x *CreateTradeReq) ProtoReflect() protoreflect.Message {
	mi := &file_order_create_trade_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTradeReq.ProtoReflect.Descriptor instead.
func (*CreateTradeReq) Descriptor() ([]byte, []int) {
	return file_order_create_trade_proto_rawDescGZIP(), []int{0}
}

func (x *CreateTradeReq) GetTradeNo() string {
	if x != nil {
		return x.TradeNo
	}
	return ""
}

type CreateTradeRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code  int32                              `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg   string                             `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	LogId string                             `protobuf:"bytes,3,opt,name=log_id,json=logId,proto3" json:"log_id,omitempty"`
	Data  *CreateTradeRsp_CreateTradeRspData `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *CreateTradeRsp) Reset() {
	*x = CreateTradeRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_create_trade_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTradeRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTradeRsp) ProtoMessage() {}

func (x *CreateTradeRsp) ProtoReflect() protoreflect.Message {
	mi := &file_order_create_trade_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTradeRsp.ProtoReflect.Descriptor instead.
func (*CreateTradeRsp) Descriptor() ([]byte, []int) {
	return file_order_create_trade_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTradeRsp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CreateTradeRsp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *CreateTradeRsp) GetLogId() string {
	if x != nil {
		return x.LogId
	}
	return ""
}

func (x *CreateTradeRsp) GetData() *CreateTradeRsp_CreateTradeRspData {
	if x != nil {
		return x.Data
	}
	return nil
}

type CreateTradeRsp_CreateTradeRspData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TradeNo       string       `protobuf:"bytes,1,opt,name=trade_no,json=tradeNo,proto3" json:"trade_no,omitempty"`                     // 电商交易号
	OrderInfoList []*OrderInfo `protobuf:"bytes,2,rep,name=order_info_list,json=orderInfoList,proto3" json:"order_info_list,omitempty"` // 订单列表
}

func (x *CreateTradeRsp_CreateTradeRspData) Reset() {
	*x = CreateTradeRsp_CreateTradeRspData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_create_trade_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTradeRsp_CreateTradeRspData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTradeRsp_CreateTradeRspData) ProtoMessage() {}

func (x *CreateTradeRsp_CreateTradeRspData) ProtoReflect() protoreflect.Message {
	mi := &file_order_create_trade_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTradeRsp_CreateTradeRspData.ProtoReflect.Descriptor instead.
func (*CreateTradeRsp_CreateTradeRspData) Descriptor() ([]byte, []int) {
	return file_order_create_trade_proto_rawDescGZIP(), []int{1, 0}
}

func (x *CreateTradeRsp_CreateTradeRspData) GetTradeNo() string {
	if x != nil {
		return x.TradeNo
	}
	return ""
}

func (x *CreateTradeRsp_CreateTradeRspData) GetOrderInfoList() []*OrderInfo {
	if x != nil {
		return x.OrderInfoList
	}
	return nil
}

var File_order_create_trade_proto protoreflect.FileDescriptor

var file_order_create_trade_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74,
	0x72, 0x61, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x1a, 0x12, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2b, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x72, 0x61, 0x64, 0x65, 0x52, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x64, 0x65,
	0x5f, 0x6e, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x72, 0x61, 0x64, 0x65,
	0x4e, 0x6f, 0x22, 0xf6, 0x01, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61,
	0x64, 0x65, 0x52, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x15, 0x0a, 0x06, 0x6c,
	0x6f, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67,
	0x49, 0x64, 0x12, 0x3c, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x28, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x72, 0x61, 0x64, 0x65, 0x52, 0x73, 0x70, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72,
	0x61, 0x64, 0x65, 0x52, 0x73, 0x70, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x1a, 0x69, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x64, 0x65, 0x52,
	0x73, 0x70, 0x44, 0x61, 0x74, 0x61, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x64, 0x65, 0x5f,
	0x6e, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x72, 0x61, 0x64, 0x65, 0x4e,
	0x6f, 0x12, 0x38, 0x0a, 0x0f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0d, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x35, 0x5a, 0x33, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x7a, 0x6c, 0x2d, 0x68, 0x65,
	0x72, 0x65, 0x2f, 0x62, 0x74, 0x2d, 0x73, 0x68, 0x6f, 0x70, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x65,
	0x6e, 0x64, 0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_order_create_trade_proto_rawDescOnce sync.Once
	file_order_create_trade_proto_rawDescData = file_order_create_trade_proto_rawDesc
)

func file_order_create_trade_proto_rawDescGZIP() []byte {
	file_order_create_trade_proto_rawDescOnce.Do(func() {
		file_order_create_trade_proto_rawDescData = protoimpl.X.CompressGZIP(file_order_create_trade_proto_rawDescData)
	})
	return file_order_create_trade_proto_rawDescData
}

var file_order_create_trade_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_order_create_trade_proto_goTypes = []interface{}{
	(*CreateTradeReq)(nil),                    // 0: order.CreateTradeReq
	(*CreateTradeRsp)(nil),                    // 1: order.CreateTradeRsp
	(*CreateTradeRsp_CreateTradeRspData)(nil), // 2: order.CreateTradeRsp.CreateTradeRspData
	(*OrderInfo)(nil),                         // 3: order.OrderInfo
}
var file_order_create_trade_proto_depIdxs = []int32{
	2, // 0: order.CreateTradeRsp.data:type_name -> order.CreateTradeRsp.CreateTradeRspData
	3, // 1: order.CreateTradeRsp.CreateTradeRspData.order_info_list:type_name -> order.OrderInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_order_create_trade_proto_init() }
func file_order_create_trade_proto_init() {
	if File_order_create_trade_proto != nil {
		return
	}
	file_order_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_order_create_trade_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTradeReq); i {
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
		file_order_create_trade_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTradeRsp); i {
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
		file_order_create_trade_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTradeRsp_CreateTradeRspData); i {
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
			RawDescriptor: file_order_create_trade_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_order_create_trade_proto_goTypes,
		DependencyIndexes: file_order_create_trade_proto_depIdxs,
		MessageInfos:      file_order_create_trade_proto_msgTypes,
	}.Build()
	File_order_create_trade_proto = out.File
	file_order_create_trade_proto_rawDesc = nil
	file_order_create_trade_proto_goTypes = nil
	file_order_create_trade_proto_depIdxs = nil
}

var _ context.Context
