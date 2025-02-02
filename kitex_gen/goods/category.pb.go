// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.29.2
// source: goods/category.proto

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

type GetCategoryListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetCategoryListReq) Reset() {
	*x = GetCategoryListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_category_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCategoryListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCategoryListReq) ProtoMessage() {}

func (x *GetCategoryListReq) ProtoReflect() protoreflect.Message {
	mi := &file_goods_category_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCategoryListReq.ProtoReflect.Descriptor instead.
func (*GetCategoryListReq) Descriptor() ([]byte, []int) {
	return file_goods_category_proto_rawDescGZIP(), []int{0}
}

type GetCategoryListRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code  int32                                      `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg   string                                     `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	LogId string                                     `protobuf:"bytes,3,opt,name=log_id,json=logId,proto3" json:"log_id,omitempty"`
	Data  *GetCategoryListRsp_GetCategoryListRspData `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetCategoryListRsp) Reset() {
	*x = GetCategoryListRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_category_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCategoryListRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCategoryListRsp) ProtoMessage() {}

func (x *GetCategoryListRsp) ProtoReflect() protoreflect.Message {
	mi := &file_goods_category_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCategoryListRsp.ProtoReflect.Descriptor instead.
func (*GetCategoryListRsp) Descriptor() ([]byte, []int) {
	return file_goods_category_proto_rawDescGZIP(), []int{1}
}

func (x *GetCategoryListRsp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetCategoryListRsp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *GetCategoryListRsp) GetLogId() string {
	if x != nil {
		return x.LogId
	}
	return ""
}

func (x *GetCategoryListRsp) GetData() *GetCategoryListRsp_GetCategoryListRspData {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetCategoryListRsp_GetCategoryListRspData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CategoryList []*GetCategoryListRsp_CategoryInfo `protobuf:"bytes,1,rep,name=category_list,json=categoryList,proto3" json:"category_list,omitempty"`
}

func (x *GetCategoryListRsp_GetCategoryListRspData) Reset() {
	*x = GetCategoryListRsp_GetCategoryListRspData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_category_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCategoryListRsp_GetCategoryListRspData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCategoryListRsp_GetCategoryListRspData) ProtoMessage() {}

func (x *GetCategoryListRsp_GetCategoryListRspData) ProtoReflect() protoreflect.Message {
	mi := &file_goods_category_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCategoryListRsp_GetCategoryListRspData.ProtoReflect.Descriptor instead.
func (*GetCategoryListRsp_GetCategoryListRspData) Descriptor() ([]byte, []int) {
	return file_goods_category_proto_rawDescGZIP(), []int{1, 0}
}

func (x *GetCategoryListRsp_GetCategoryListRspData) GetCategoryList() []*GetCategoryListRsp_CategoryInfo {
	if x != nil {
		return x.CategoryList
	}
	return nil
}

type GetCategoryListRsp_CategoryInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CategoryId   uint64 `protobuf:"varint,1,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	CategoryName string `protobuf:"bytes,2,opt,name=category_name,json=categoryName,proto3" json:"category_name,omitempty"`
	ParentId     uint64 `protobuf:"varint,3,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
}

func (x *GetCategoryListRsp_CategoryInfo) Reset() {
	*x = GetCategoryListRsp_CategoryInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_category_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCategoryListRsp_CategoryInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCategoryListRsp_CategoryInfo) ProtoMessage() {}

func (x *GetCategoryListRsp_CategoryInfo) ProtoReflect() protoreflect.Message {
	mi := &file_goods_category_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCategoryListRsp_CategoryInfo.ProtoReflect.Descriptor instead.
func (*GetCategoryListRsp_CategoryInfo) Descriptor() ([]byte, []int) {
	return file_goods_category_proto_rawDescGZIP(), []int{1, 1}
}

func (x *GetCategoryListRsp_CategoryInfo) GetCategoryId() uint64 {
	if x != nil {
		return x.CategoryId
	}
	return 0
}

func (x *GetCategoryListRsp_CategoryInfo) GetCategoryName() string {
	if x != nil {
		return x.CategoryName
	}
	return ""
}

func (x *GetCategoryListRsp_CategoryInfo) GetParentId() uint64 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

var File_goods_category_proto protoreflect.FileDescriptor

var file_goods_category_proto_rawDesc = []byte{
	0x0a, 0x14, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x22, 0x14, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x22, 0xf1, 0x02, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67,
	0x12, 0x15, 0x0a, 0x06, 0x6c, 0x6f, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x12, 0x44, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x47, 0x65,
	0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x73, 0x70,
	0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x73, 0x70, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x65, 0x0a,
	0x16, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x73, 0x70, 0x44, 0x61, 0x74, 0x61, 0x12, 0x4b, 0x0a, 0x0d, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26,
	0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x73, 0x70, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0c, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x4c, 0x69, 0x73, 0x74, 0x1a, 0x71, 0x0a, 0x0c, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x7a, 0x6c, 0x2d, 0x68, 0x65, 0x72, 0x65, 0x2f, 0x62,
	0x74, 0x2d, 0x73, 0x68, 0x6f, 0x70, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x6b,
	0x69, 0x74, 0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_goods_category_proto_rawDescOnce sync.Once
	file_goods_category_proto_rawDescData = file_goods_category_proto_rawDesc
)

func file_goods_category_proto_rawDescGZIP() []byte {
	file_goods_category_proto_rawDescOnce.Do(func() {
		file_goods_category_proto_rawDescData = protoimpl.X.CompressGZIP(file_goods_category_proto_rawDescData)
	})
	return file_goods_category_proto_rawDescData
}

var file_goods_category_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_goods_category_proto_goTypes = []interface{}{
	(*GetCategoryListReq)(nil),                        // 0: goods.GetCategoryListReq
	(*GetCategoryListRsp)(nil),                        // 1: goods.GetCategoryListRsp
	(*GetCategoryListRsp_GetCategoryListRspData)(nil), // 2: goods.GetCategoryListRsp.GetCategoryListRspData
	(*GetCategoryListRsp_CategoryInfo)(nil),           // 3: goods.GetCategoryListRsp.CategoryInfo
}
var file_goods_category_proto_depIdxs = []int32{
	2, // 0: goods.GetCategoryListRsp.data:type_name -> goods.GetCategoryListRsp.GetCategoryListRspData
	3, // 1: goods.GetCategoryListRsp.GetCategoryListRspData.category_list:type_name -> goods.GetCategoryListRsp.CategoryInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_goods_category_proto_init() }
func file_goods_category_proto_init() {
	if File_goods_category_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_goods_category_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCategoryListReq); i {
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
		file_goods_category_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCategoryListRsp); i {
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
		file_goods_category_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCategoryListRsp_GetCategoryListRspData); i {
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
		file_goods_category_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCategoryListRsp_CategoryInfo); i {
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
			RawDescriptor: file_goods_category_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_goods_category_proto_goTypes,
		DependencyIndexes: file_goods_category_proto_depIdxs,
		MessageInfos:      file_goods_category_proto_msgTypes,
	}.Build()
	File_goods_category_proto = out.File
	file_goods_category_proto_rawDesc = nil
	file_goods_category_proto_goTypes = nil
	file_goods_category_proto_depIdxs = nil
}

var _ context.Context
