// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.29.2
// source: order/trade.proto

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

// 创建交易
type CreateTradeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token     string                    `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	TradeInfo *CreateTradeReq_TradeInfo `protobuf:"bytes,2,opt,name=trade_info,json=tradeInfo,proto3" json:"trade_info,omitempty"`
	PayType   string                    `protobuf:"bytes,3,opt,name=pay_type,json=payType,proto3" json:"pay_type,omitempty"` // 支付类型
}

func (x *CreateTradeReq) Reset() {
	*x = CreateTradeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_trade_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTradeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTradeReq) ProtoMessage() {}

func (x *CreateTradeReq) ProtoReflect() protoreflect.Message {
	mi := &file_order_trade_proto_msgTypes[0]
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
	return file_order_trade_proto_rawDescGZIP(), []int{0}
}

func (x *CreateTradeReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *CreateTradeReq) GetTradeInfo() *CreateTradeReq_TradeInfo {
	if x != nil {
		return x.TradeInfo
	}
	return nil
}

func (x *CreateTradeReq) GetPayType() string {
	if x != nil {
		return x.PayType
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
		mi := &file_order_trade_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTradeRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTradeRsp) ProtoMessage() {}

func (x *CreateTradeRsp) ProtoReflect() protoreflect.Message {
	mi := &file_order_trade_proto_msgTypes[1]
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
	return file_order_trade_proto_rawDescGZIP(), []int{1}
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

// 取消交易
type CancelTradeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TradeNo string `protobuf:"bytes,1,opt,name=trade_no,json=tradeNo,proto3" json:"trade_no,omitempty"`
}

func (x *CancelTradeReq) Reset() {
	*x = CancelTradeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_trade_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelTradeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelTradeReq) ProtoMessage() {}

func (x *CancelTradeReq) ProtoReflect() protoreflect.Message {
	mi := &file_order_trade_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelTradeReq.ProtoReflect.Descriptor instead.
func (*CancelTradeReq) Descriptor() ([]byte, []int) {
	return file_order_trade_proto_rawDescGZIP(), []int{2}
}

func (x *CancelTradeReq) GetTradeNo() string {
	if x != nil {
		return x.TradeNo
	}
	return ""
}

type CancelTradeRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code  int32                              `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg   string                             `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	LogId string                             `protobuf:"bytes,3,opt,name=log_id,json=logId,proto3" json:"log_id,omitempty"`
	Data  *CancelTradeRsp_CancelTradeRspData `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *CancelTradeRsp) Reset() {
	*x = CancelTradeRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_trade_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelTradeRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelTradeRsp) ProtoMessage() {}

func (x *CancelTradeRsp) ProtoReflect() protoreflect.Message {
	mi := &file_order_trade_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelTradeRsp.ProtoReflect.Descriptor instead.
func (*CancelTradeRsp) Descriptor() ([]byte, []int) {
	return file_order_trade_proto_rawDescGZIP(), []int{3}
}

func (x *CancelTradeRsp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CancelTradeRsp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *CancelTradeRsp) GetLogId() string {
	if x != nil {
		return x.LogId
	}
	return ""
}

func (x *CancelTradeRsp) GetData() *CancelTradeRsp_CancelTradeRspData {
	if x != nil {
		return x.Data
	}
	return nil
}

type CreateTradeReq_TradeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderInfoList []*TradeOrderInfo `protobuf:"bytes,1,rep,name=order_info_list,json=orderInfoList,proto3" json:"order_info_list,omitempty"`
	TradeAmount   string            `protobuf:"bytes,2,opt,name=trade_amount,json=tradeAmount,proto3" json:"trade_amount,omitempty"`
}

func (x *CreateTradeReq_TradeInfo) Reset() {
	*x = CreateTradeReq_TradeInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_trade_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTradeReq_TradeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTradeReq_TradeInfo) ProtoMessage() {}

func (x *CreateTradeReq_TradeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_order_trade_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTradeReq_TradeInfo.ProtoReflect.Descriptor instead.
func (*CreateTradeReq_TradeInfo) Descriptor() ([]byte, []int) {
	return file_order_trade_proto_rawDescGZIP(), []int{0, 0}
}

func (x *CreateTradeReq_TradeInfo) GetOrderInfoList() []*TradeOrderInfo {
	if x != nil {
		return x.OrderInfoList
	}
	return nil
}

func (x *CreateTradeReq_TradeInfo) GetTradeAmount() string {
	if x != nil {
		return x.TradeAmount
	}
	return ""
}

type CreateTradeRsp_CreateTradeRspData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TradeNo    string `protobuf:"bytes,1,opt,name=trade_no,json=tradeNo,proto3" json:"trade_no,omitempty"` // 电商交易号
	PayPageUrl string `protobuf:"bytes,2,opt,name=pay_page_url,json=payPageUrl,proto3" json:"pay_page_url,omitempty"`
}

func (x *CreateTradeRsp_CreateTradeRspData) Reset() {
	*x = CreateTradeRsp_CreateTradeRspData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_trade_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTradeRsp_CreateTradeRspData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTradeRsp_CreateTradeRspData) ProtoMessage() {}

func (x *CreateTradeRsp_CreateTradeRspData) ProtoReflect() protoreflect.Message {
	mi := &file_order_trade_proto_msgTypes[5]
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
	return file_order_trade_proto_rawDescGZIP(), []int{1, 0}
}

func (x *CreateTradeRsp_CreateTradeRspData) GetTradeNo() string {
	if x != nil {
		return x.TradeNo
	}
	return ""
}

func (x *CreateTradeRsp_CreateTradeRspData) GetPayPageUrl() string {
	if x != nil {
		return x.PayPageUrl
	}
	return ""
}

type CancelTradeRsp_CancelTradeRspData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TradeNo string `protobuf:"bytes,1,opt,name=trade_no,json=tradeNo,proto3" json:"trade_no,omitempty"` // 电商交易号
}

func (x *CancelTradeRsp_CancelTradeRspData) Reset() {
	*x = CancelTradeRsp_CancelTradeRspData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_trade_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelTradeRsp_CancelTradeRspData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelTradeRsp_CancelTradeRspData) ProtoMessage() {}

func (x *CancelTradeRsp_CancelTradeRspData) ProtoReflect() protoreflect.Message {
	mi := &file_order_trade_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelTradeRsp_CancelTradeRspData.ProtoReflect.Descriptor instead.
func (*CancelTradeRsp_CancelTradeRspData) Descriptor() ([]byte, []int) {
	return file_order_trade_proto_rawDescGZIP(), []int{3, 0}
}

func (x *CancelTradeRsp_CancelTradeRspData) GetTradeNo() string {
	if x != nil {
		return x.TradeNo
	}
	return ""
}

var File_order_trade_proto protoreflect.FileDescriptor

var file_order_trade_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x74, 0x72, 0x61, 0x64, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x1a, 0x12, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf0,
	0x01, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x64, 0x65, 0x52, 0x65,
	0x71, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x3e, 0x0a, 0x0a, 0x74, 0x72, 0x61, 0x64, 0x65,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x64, 0x65, 0x52,
	0x65, 0x71, 0x2e, 0x54, 0x72, 0x61, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x74, 0x72,
	0x61, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x61, 0x79, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x79, 0x54, 0x79,
	0x70, 0x65, 0x1a, 0x6d, 0x0a, 0x09, 0x54, 0x72, 0x61, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x3d, 0x0a, 0x0f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x2e, 0x54, 0x72, 0x61, 0x64, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x0d, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x21,
	0x0a, 0x0c, 0x74, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x64, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x22, 0xde, 0x01, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x64,
	0x65, 0x52, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x15, 0x0a, 0x06, 0x6c, 0x6f,
	0x67, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x49,
	0x64, 0x12, 0x3c, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x28, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72,
	0x61, 0x64, 0x65, 0x52, 0x73, 0x70, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61,
	0x64, 0x65, 0x52, 0x73, 0x70, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a,
	0x51, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x64, 0x65, 0x52, 0x73,
	0x70, 0x44, 0x61, 0x74, 0x61, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x6e,
	0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x72, 0x61, 0x64, 0x65, 0x4e, 0x6f,
	0x12, 0x20, 0x0a, 0x0c, 0x70, 0x61, 0x79, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x61, 0x79, 0x50, 0x61, 0x67, 0x65, 0x55,
	0x72, 0x6c, 0x22, 0x2b, 0x0a, 0x0e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x54, 0x72, 0x61, 0x64,
	0x65, 0x52, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x6e, 0x6f,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x72, 0x61, 0x64, 0x65, 0x4e, 0x6f, 0x22,
	0xbc, 0x01, 0x0a, 0x0e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x54, 0x72, 0x61, 0x64, 0x65, 0x52,
	0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x15, 0x0a, 0x06, 0x6c, 0x6f, 0x67, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x12,
	0x3c, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x54, 0x72, 0x61, 0x64,
	0x65, 0x52, 0x73, 0x70, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x54, 0x72, 0x61, 0x64, 0x65,
	0x52, 0x73, 0x70, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x2f, 0x0a,
	0x12, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x54, 0x72, 0x61, 0x64, 0x65, 0x52, 0x73, 0x70, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x6e, 0x6f, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x72, 0x61, 0x64, 0x65, 0x4e, 0x6f, 0x42, 0x35,
	0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x7a, 0x6c,
	0x2d, 0x68, 0x65, 0x72, 0x65, 0x2f, 0x62, 0x74, 0x2d, 0x73, 0x68, 0x6f, 0x70, 0x2d, 0x62, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_order_trade_proto_rawDescOnce sync.Once
	file_order_trade_proto_rawDescData = file_order_trade_proto_rawDesc
)

func file_order_trade_proto_rawDescGZIP() []byte {
	file_order_trade_proto_rawDescOnce.Do(func() {
		file_order_trade_proto_rawDescData = protoimpl.X.CompressGZIP(file_order_trade_proto_rawDescData)
	})
	return file_order_trade_proto_rawDescData
}

var file_order_trade_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_order_trade_proto_goTypes = []interface{}{
	(*CreateTradeReq)(nil),                    // 0: order.CreateTradeReq
	(*CreateTradeRsp)(nil),                    // 1: order.CreateTradeRsp
	(*CancelTradeReq)(nil),                    // 2: order.CancelTradeReq
	(*CancelTradeRsp)(nil),                    // 3: order.CancelTradeRsp
	(*CreateTradeReq_TradeInfo)(nil),          // 4: order.CreateTradeReq.TradeInfo
	(*CreateTradeRsp_CreateTradeRspData)(nil), // 5: order.CreateTradeRsp.CreateTradeRspData
	(*CancelTradeRsp_CancelTradeRspData)(nil), // 6: order.CancelTradeRsp.CancelTradeRspData
	(*TradeOrderInfo)(nil),                    // 7: order.TradeOrderInfo
}
var file_order_trade_proto_depIdxs = []int32{
	4, // 0: order.CreateTradeReq.trade_info:type_name -> order.CreateTradeReq.TradeInfo
	5, // 1: order.CreateTradeRsp.data:type_name -> order.CreateTradeRsp.CreateTradeRspData
	6, // 2: order.CancelTradeRsp.data:type_name -> order.CancelTradeRsp.CancelTradeRspData
	7, // 3: order.CreateTradeReq.TradeInfo.order_info_list:type_name -> order.TradeOrderInfo
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_order_trade_proto_init() }
func file_order_trade_proto_init() {
	if File_order_trade_proto != nil {
		return
	}
	file_order_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_order_trade_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_order_trade_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_order_trade_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelTradeReq); i {
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
		file_order_trade_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelTradeRsp); i {
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
		file_order_trade_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTradeReq_TradeInfo); i {
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
		file_order_trade_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_order_trade_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelTradeRsp_CancelTradeRspData); i {
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
			RawDescriptor: file_order_trade_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_order_trade_proto_goTypes,
		DependencyIndexes: file_order_trade_proto_depIdxs,
		MessageInfos:      file_order_trade_proto_msgTypes,
	}.Build()
	File_order_trade_proto = out.File
	file_order_trade_proto_rawDesc = nil
	file_order_trade_proto_goTypes = nil
	file_order_trade_proto_depIdxs = nil
}

var _ context.Context
