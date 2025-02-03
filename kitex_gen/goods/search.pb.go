// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.29.2
// source: goods/search.proto

package goods

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

// 获取关键词下拉列表
type GetKeywordDownListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Prefix string `protobuf:"bytes,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
}

func (x *GetKeywordDownListReq) Reset() {
	*x = GetKeywordDownListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_search_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetKeywordDownListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetKeywordDownListReq) ProtoMessage() {}

func (x *GetKeywordDownListReq) ProtoReflect() protoreflect.Message {
	mi := &file_goods_search_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetKeywordDownListReq.ProtoReflect.Descriptor instead.
func (*GetKeywordDownListReq) Descriptor() ([]byte, []int) {
	return file_goods_search_proto_rawDescGZIP(), []int{0}
}

func (x *GetKeywordDownListReq) GetPrefix() string {
	if x != nil {
		return x.Prefix
	}
	return ""
}

type GetKeywordDownListRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code  int32                                        `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg   string                                       `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	LogId string                                       `protobuf:"bytes,3,opt,name=log_id,json=logId,proto3" json:"log_id,omitempty"`
	Data  *GetKeywordDownListRsp_GetKeyDownListRspData `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetKeywordDownListRsp) Reset() {
	*x = GetKeywordDownListRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_search_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetKeywordDownListRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetKeywordDownListRsp) ProtoMessage() {}

func (x *GetKeywordDownListRsp) ProtoReflect() protoreflect.Message {
	mi := &file_goods_search_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetKeywordDownListRsp.ProtoReflect.Descriptor instead.
func (*GetKeywordDownListRsp) Descriptor() ([]byte, []int) {
	return file_goods_search_proto_rawDescGZIP(), []int{1}
}

func (x *GetKeywordDownListRsp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetKeywordDownListRsp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *GetKeywordDownListRsp) GetLogId() string {
	if x != nil {
		return x.LogId
	}
	return ""
}

func (x *GetKeywordDownListRsp) GetData() *GetKeywordDownListRsp_GetKeyDownListRspData {
	if x != nil {
		return x.Data
	}
	return nil
}

// 分页搜索商品列表
type SearchSpuListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageSize    int32    `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageNo      int32    `protobuf:"varint,2,opt,name=page_no,json=pageNo,proto3" json:"page_no,omitempty"`
	Keyword     string   `protobuf:"bytes,3,opt,name=keyword,proto3" json:"keyword,omitempty"`                                    // 关键词
	CategoryIds []uint64 `protobuf:"varint,4,rep,packed,name=category_ids,json=categoryIds,proto3" json:"category_ids,omitempty"` //分类id
	BrandIds    []uint64 `protobuf:"varint,5,rep,packed,name=brand_ids,json=brandIds,proto3" json:"brand_ids,omitempty"`          // 品牌id
}

func (x *SearchSpuListReq) Reset() {
	*x = SearchSpuListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_search_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchSpuListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchSpuListReq) ProtoMessage() {}

func (x *SearchSpuListReq) ProtoReflect() protoreflect.Message {
	mi := &file_goods_search_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchSpuListReq.ProtoReflect.Descriptor instead.
func (*SearchSpuListReq) Descriptor() ([]byte, []int) {
	return file_goods_search_proto_rawDescGZIP(), []int{2}
}

func (x *SearchSpuListReq) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *SearchSpuListReq) GetPageNo() int32 {
	if x != nil {
		return x.PageNo
	}
	return 0
}

func (x *SearchSpuListReq) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

func (x *SearchSpuListReq) GetCategoryIds() []uint64 {
	if x != nil {
		return x.CategoryIds
	}
	return nil
}

func (x *SearchSpuListReq) GetBrandIds() []uint64 {
	if x != nil {
		return x.BrandIds
	}
	return nil
}

type SearchSpuListRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code  int32                                  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg   string                                 `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	LogId string                                 `protobuf:"bytes,3,opt,name=log_id,json=logId,proto3" json:"log_id,omitempty"`
	Data  *SearchSpuListRsp_SearchSpuListRspData `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SearchSpuListRsp) Reset() {
	*x = SearchSpuListRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_search_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchSpuListRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchSpuListRsp) ProtoMessage() {}

func (x *SearchSpuListRsp) ProtoReflect() protoreflect.Message {
	mi := &file_goods_search_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchSpuListRsp.ProtoReflect.Descriptor instead.
func (*SearchSpuListRsp) Descriptor() ([]byte, []int) {
	return file_goods_search_proto_rawDescGZIP(), []int{3}
}

func (x *SearchSpuListRsp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SearchSpuListRsp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *SearchSpuListRsp) GetLogId() string {
	if x != nil {
		return x.LogId
	}
	return ""
}

func (x *SearchSpuListRsp) GetData() *SearchSpuListRsp_SearchSpuListRspData {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetKeywordDownListRsp_GetKeyDownListRspData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeywordList []*BaseKeyword `protobuf:"bytes,1,rep,name=keyword_list,json=keywordList,proto3" json:"keyword_list,omitempty"`
}

func (x *GetKeywordDownListRsp_GetKeyDownListRspData) Reset() {
	*x = GetKeywordDownListRsp_GetKeyDownListRspData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_search_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetKeywordDownListRsp_GetKeyDownListRspData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetKeywordDownListRsp_GetKeyDownListRspData) ProtoMessage() {}

func (x *GetKeywordDownListRsp_GetKeyDownListRspData) ProtoReflect() protoreflect.Message {
	mi := &file_goods_search_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetKeywordDownListRsp_GetKeyDownListRspData.ProtoReflect.Descriptor instead.
func (*GetKeywordDownListRsp_GetKeyDownListRspData) Descriptor() ([]byte, []int) {
	return file_goods_search_proto_rawDescGZIP(), []int{1, 0}
}

func (x *GetKeywordDownListRsp_GetKeyDownListRspData) GetKeywordList() []*BaseKeyword {
	if x != nil {
		return x.KeywordList
	}
	return nil
}

type SearchSpuListRsp_SearchSpuListRspData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SpuList  []*BaseSpu `protobuf:"bytes,1,rep,name=spu_list,json=spuList,proto3" json:"spu_list,omitempty"`
	PageSize int32      `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageNo   int32      `protobuf:"varint,3,opt,name=page_no,json=pageNo,proto3" json:"page_no,omitempty"`
	Count    int32      `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *SearchSpuListRsp_SearchSpuListRspData) Reset() {
	*x = SearchSpuListRsp_SearchSpuListRspData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_search_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchSpuListRsp_SearchSpuListRspData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchSpuListRsp_SearchSpuListRspData) ProtoMessage() {}

func (x *SearchSpuListRsp_SearchSpuListRspData) ProtoReflect() protoreflect.Message {
	mi := &file_goods_search_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchSpuListRsp_SearchSpuListRspData.ProtoReflect.Descriptor instead.
func (*SearchSpuListRsp_SearchSpuListRspData) Descriptor() ([]byte, []int) {
	return file_goods_search_proto_rawDescGZIP(), []int{3, 0}
}

func (x *SearchSpuListRsp_SearchSpuListRspData) GetSpuList() []*BaseSpu {
	if x != nil {
		return x.SpuList
	}
	return nil
}

func (x *SearchSpuListRsp_SearchSpuListRspData) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *SearchSpuListRsp_SearchSpuListRspData) GetPageNo() int32 {
	if x != nil {
		return x.PageNo
	}
	return 0
}

func (x *SearchSpuListRsp_SearchSpuListRspData) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_goods_search_proto protoreflect.FileDescriptor

var file_goods_search_proto_rawDesc = []byte{
	0x0a, 0x12, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x1a, 0x12, 0x67, 0x6f, 0x6f,
	0x64, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x2f, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x44, 0x6f, 0x77,
	0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72, 0x65, 0x66,
	0x69, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78,
	0x22, 0xec, 0x01, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x44,
	0x6f, 0x77, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67,
	0x12, 0x15, 0x0a, 0x06, 0x6c, 0x6f, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x12, 0x46, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x47, 0x65,
	0x74, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x44, 0x6f, 0x77, 0x6e, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x73, 0x70, 0x2e, 0x47, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x44, 0x6f, 0x77, 0x6e, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x73, 0x70, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a,
	0x4e, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x44, 0x6f, 0x77, 0x6e, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x73, 0x70, 0x44, 0x61, 0x74, 0x61, 0x12, 0x35, 0x0a, 0x0c, 0x6b, 0x65, 0x79, 0x77,
	0x6f, 0x72, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x4b, 0x65, 0x79, 0x77, 0x6f,
	0x72, 0x64, 0x52, 0x0b, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x22,
	0xa2, 0x01, 0x0a, 0x10, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x53, 0x70, 0x75, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x6f, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65,
	0x79, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79,
	0x77, 0x6f, 0x72, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x5f, 0x69, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x04, 0x52, 0x0b, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x72, 0x61, 0x6e, 0x64,
	0x5f, 0x69, 0x64, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x04, 0x52, 0x08, 0x62, 0x72, 0x61, 0x6e,
	0x64, 0x49, 0x64, 0x73, 0x22, 0xa1, 0x02, 0x0a, 0x10, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x53,
	0x70, 0x75, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12,
	0x15, 0x0a, 0x06, 0x6c, 0x6f, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x12, 0x40, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x53, 0x70, 0x75, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x73, 0x70, 0x2e, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x53, 0x70, 0x75, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x73, 0x70, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x8d, 0x01, 0x0a, 0x14, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x53, 0x70, 0x75, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x73, 0x70, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x29, 0x0a, 0x08, 0x73, 0x70, 0x75, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x42, 0x61, 0x73, 0x65,
	0x53, 0x70, 0x75, 0x52, 0x07, 0x73, 0x70, 0x75, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x6e, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x61, 0x67, 0x65,
	0x4e, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x7a, 0x6c, 0x2d, 0x68, 0x65, 0x72, 0x65, 0x2f,
	0x62, 0x74, 0x2d, 0x73, 0x68, 0x6f, 0x70, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f,
	0x6b, 0x69, 0x74, 0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_goods_search_proto_rawDescOnce sync.Once
	file_goods_search_proto_rawDescData = file_goods_search_proto_rawDesc
)

func file_goods_search_proto_rawDescGZIP() []byte {
	file_goods_search_proto_rawDescOnce.Do(func() {
		file_goods_search_proto_rawDescData = protoimpl.X.CompressGZIP(file_goods_search_proto_rawDescData)
	})
	return file_goods_search_proto_rawDescData
}

var file_goods_search_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_goods_search_proto_goTypes = []interface{}{
	(*GetKeywordDownListReq)(nil),                       // 0: goods.GetKeywordDownListReq
	(*GetKeywordDownListRsp)(nil),                       // 1: goods.GetKeywordDownListRsp
	(*SearchSpuListReq)(nil),                            // 2: goods.SearchSpuListReq
	(*SearchSpuListRsp)(nil),                            // 3: goods.SearchSpuListRsp
	(*GetKeywordDownListRsp_GetKeyDownListRspData)(nil), // 4: goods.GetKeywordDownListRsp.GetKeyDownListRspData
	(*SearchSpuListRsp_SearchSpuListRspData)(nil),       // 5: goods.SearchSpuListRsp.SearchSpuListRspData
	(*BaseKeyword)(nil),                                 // 6: goods.BaseKeyword
	(*BaseSpu)(nil),                                     // 7: goods.BaseSpu
}
var file_goods_search_proto_depIdxs = []int32{
	4, // 0: goods.GetKeywordDownListRsp.data:type_name -> goods.GetKeywordDownListRsp.GetKeyDownListRspData
	5, // 1: goods.SearchSpuListRsp.data:type_name -> goods.SearchSpuListRsp.SearchSpuListRspData
	6, // 2: goods.GetKeywordDownListRsp.GetKeyDownListRspData.keyword_list:type_name -> goods.BaseKeyword
	7, // 3: goods.SearchSpuListRsp.SearchSpuListRspData.spu_list:type_name -> goods.BaseSpu
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_goods_search_proto_init() }
func file_goods_search_proto_init() {
	if File_goods_search_proto != nil {
		return
	}
	file_goods_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_goods_search_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetKeywordDownListReq); i {
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
		file_goods_search_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetKeywordDownListRsp); i {
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
		file_goods_search_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchSpuListReq); i {
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
		file_goods_search_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchSpuListRsp); i {
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
		file_goods_search_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetKeywordDownListRsp_GetKeyDownListRspData); i {
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
		file_goods_search_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchSpuListRsp_SearchSpuListRspData); i {
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
			RawDescriptor: file_goods_search_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_goods_search_proto_goTypes,
		DependencyIndexes: file_goods_search_proto_depIdxs,
		MessageInfos:      file_goods_search_proto_msgTypes,
	}.Build()
	File_goods_search_proto = out.File
	file_goods_search_proto_rawDesc = nil
	file_goods_search_proto_goTypes = nil
	file_goods_search_proto_depIdxs = nil
}

var _ context.Context
