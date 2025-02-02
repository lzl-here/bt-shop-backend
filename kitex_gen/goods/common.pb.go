// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.29.2
// source: goods/common.proto

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

// spu信息
type SpuInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SpuId           uint64           `protobuf:"varint,1,opt,name=spu_id,json=spuId,proto3" json:"spu_id,omitempty"`
	SpuName         string           `protobuf:"bytes,2,opt,name=spu_name,json=spuName,proto3" json:"spu_name,omitempty"`
	SpuDesc         string           `protobuf:"bytes,3,opt,name=spu_desc,json=spuDesc,proto3" json:"spu_desc,omitempty"`
	SpuImgUrl       string           `protobuf:"bytes,4,opt,name=spu_img_url,json=spuImgUrl,proto3" json:"spu_img_url,omitempty"`
	SpuPrice        string           `protobuf:"bytes,5,opt,name=spu_price,json=spuPrice,proto3" json:"spu_price,omitempty"`
	SpuCategoryId   uint64           `protobuf:"varint,7,opt,name=spu_category_id,json=spuCategoryId,proto3" json:"spu_category_id,omitempty"`
	SpuCategoryName string           `protobuf:"bytes,8,opt,name=spu_category_name,json=spuCategoryName,proto3" json:"spu_category_name,omitempty"`
	BrandId         uint64           `protobuf:"varint,9,opt,name=brand_id,json=brandId,proto3" json:"brand_id,omitempty"`
	BrandName       string           `protobuf:"bytes,10,opt,name=brand_name,json=brandName,proto3" json:"brand_name,omitempty"`
	SkuList         []*SkuInfo       `protobuf:"bytes,11,rep,name=sku_list,json=skuList,proto3" json:"sku_list,omitempty"`
	SpecList        []*SpecInfo      `protobuf:"bytes,12,rep,name=spec_list,json=specList,proto3" json:"spec_list,omitempty"`                  // 可选规格
	SpecValueList   []*SpecValueInfo `protobuf:"bytes,13,rep,name=spec_value_list,json=specValueList,proto3" json:"spec_value_list,omitempty"` // 可选规格值
}

func (x *SpuInfo) Reset() {
	*x = SpuInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpuInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpuInfo) ProtoMessage() {}

func (x *SpuInfo) ProtoReflect() protoreflect.Message {
	mi := &file_goods_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpuInfo.ProtoReflect.Descriptor instead.
func (*SpuInfo) Descriptor() ([]byte, []int) {
	return file_goods_common_proto_rawDescGZIP(), []int{0}
}

func (x *SpuInfo) GetSpuId() uint64 {
	if x != nil {
		return x.SpuId
	}
	return 0
}

func (x *SpuInfo) GetSpuName() string {
	if x != nil {
		return x.SpuName
	}
	return ""
}

func (x *SpuInfo) GetSpuDesc() string {
	if x != nil {
		return x.SpuDesc
	}
	return ""
}

func (x *SpuInfo) GetSpuImgUrl() string {
	if x != nil {
		return x.SpuImgUrl
	}
	return ""
}

func (x *SpuInfo) GetSpuPrice() string {
	if x != nil {
		return x.SpuPrice
	}
	return ""
}

func (x *SpuInfo) GetSpuCategoryId() uint64 {
	if x != nil {
		return x.SpuCategoryId
	}
	return 0
}

func (x *SpuInfo) GetSpuCategoryName() string {
	if x != nil {
		return x.SpuCategoryName
	}
	return ""
}

func (x *SpuInfo) GetBrandId() uint64 {
	if x != nil {
		return x.BrandId
	}
	return 0
}

func (x *SpuInfo) GetBrandName() string {
	if x != nil {
		return x.BrandName
	}
	return ""
}

func (x *SpuInfo) GetSkuList() []*SkuInfo {
	if x != nil {
		return x.SkuList
	}
	return nil
}

func (x *SpuInfo) GetSpecList() []*SpecInfo {
	if x != nil {
		return x.SpecList
	}
	return nil
}

func (x *SpuInfo) GetSpecValueList() []*SpecValueInfo {
	if x != nil {
		return x.SpecValueList
	}
	return nil
}

// sku信息
type SkuInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SkuId     uint64 `protobuf:"varint,1,opt,name=sku_id,json=skuId,proto3" json:"sku_id,omitempty"`
	SpuId     uint64 `protobuf:"varint,2,opt,name=spu_id,json=spuId,proto3" json:"spu_id,omitempty"`
	StockNum  uint64 `protobuf:"varint,3,opt,name=stock_num,json=stockNum,proto3" json:"stock_num,omitempty"`
	SkuAmount string `protobuf:"bytes,4,opt,name=sku_amount,json=skuAmount,proto3" json:"sku_amount,omitempty"`
	SkuName   string `protobuf:"bytes,5,opt,name=sku_name,json=skuName,proto3" json:"sku_name,omitempty"`
	SpecStr   string `protobuf:"bytes,6,opt,name=spec_str,json=specStr,proto3" json:"spec_str,omitempty"`
}

func (x *SkuInfo) Reset() {
	*x = SkuInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SkuInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SkuInfo) ProtoMessage() {}

func (x *SkuInfo) ProtoReflect() protoreflect.Message {
	mi := &file_goods_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SkuInfo.ProtoReflect.Descriptor instead.
func (*SkuInfo) Descriptor() ([]byte, []int) {
	return file_goods_common_proto_rawDescGZIP(), []int{1}
}

func (x *SkuInfo) GetSkuId() uint64 {
	if x != nil {
		return x.SkuId
	}
	return 0
}

func (x *SkuInfo) GetSpuId() uint64 {
	if x != nil {
		return x.SpuId
	}
	return 0
}

func (x *SkuInfo) GetStockNum() uint64 {
	if x != nil {
		return x.StockNum
	}
	return 0
}

func (x *SkuInfo) GetSkuAmount() string {
	if x != nil {
		return x.SkuAmount
	}
	return ""
}

func (x *SkuInfo) GetSkuName() string {
	if x != nil {
		return x.SkuName
	}
	return ""
}

func (x *SkuInfo) GetSpecStr() string {
	if x != nil {
		return x.SpecStr
	}
	return ""
}

// 规格信息，查询商品信息时使用
type SpecInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SpecId   uint64 `protobuf:"varint,1,opt,name=spec_id,json=specId,proto3" json:"spec_id,omitempty"`
	SpecName string `protobuf:"bytes,2,opt,name=spec_name,json=specName,proto3" json:"spec_name,omitempty"`
	SpuId    uint64 `protobuf:"varint,3,opt,name=spu_id,json=spuId,proto3" json:"spu_id,omitempty"`
}

func (x *SpecInfo) Reset() {
	*x = SpecInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpecInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpecInfo) ProtoMessage() {}

func (x *SpecInfo) ProtoReflect() protoreflect.Message {
	mi := &file_goods_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpecInfo.ProtoReflect.Descriptor instead.
func (*SpecInfo) Descriptor() ([]byte, []int) {
	return file_goods_common_proto_rawDescGZIP(), []int{2}
}

func (x *SpecInfo) GetSpecId() uint64 {
	if x != nil {
		return x.SpecId
	}
	return 0
}

func (x *SpecInfo) GetSpecName() string {
	if x != nil {
		return x.SpecName
	}
	return ""
}

func (x *SpecInfo) GetSpuId() uint64 {
	if x != nil {
		return x.SpuId
	}
	return 0
}

// 规格值信息，查询商品信息时使用
type SpecValueInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SpecValueId uint64 `protobuf:"varint,1,opt,name=spec_value_id,json=specValueId,proto3" json:"spec_value_id,omitempty"`
	SpecName    string `protobuf:"bytes,2,opt,name=spec_name,json=specName,proto3" json:"spec_name,omitempty"`
	SpecValue   string `protobuf:"bytes,3,opt,name=spec_value,json=specValue,proto3" json:"spec_value,omitempty"`
}

func (x *SpecValueInfo) Reset() {
	*x = SpecValueInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpecValueInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpecValueInfo) ProtoMessage() {}

func (x *SpecValueInfo) ProtoReflect() protoreflect.Message {
	mi := &file_goods_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpecValueInfo.ProtoReflect.Descriptor instead.
func (*SpecValueInfo) Descriptor() ([]byte, []int) {
	return file_goods_common_proto_rawDescGZIP(), []int{3}
}

func (x *SpecValueInfo) GetSpecValueId() uint64 {
	if x != nil {
		return x.SpecValueId
	}
	return 0
}

func (x *SpecValueInfo) GetSpecName() string {
	if x != nil {
		return x.SpecName
	}
	return ""
}

func (x *SpecValueInfo) GetSpecValue() string {
	if x != nil {
		return x.SpecValue
	}
	return ""
}

// 规格项列表， 需要新增规格时使用
type SpecKeyValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SpecName  string `protobuf:"bytes,1,opt,name=spec_name,json=specName,proto3" json:"spec_name,omitempty"`
	SpecValue string `protobuf:"bytes,2,opt,name=spec_value,json=specValue,proto3" json:"spec_value,omitempty"`
}

func (x *SpecKeyValue) Reset() {
	*x = SpecKeyValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpecKeyValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpecKeyValue) ProtoMessage() {}

func (x *SpecKeyValue) ProtoReflect() protoreflect.Message {
	mi := &file_goods_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpecKeyValue.ProtoReflect.Descriptor instead.
func (*SpecKeyValue) Descriptor() ([]byte, []int) {
	return file_goods_common_proto_rawDescGZIP(), []int{4}
}

func (x *SpecKeyValue) GetSpecName() string {
	if x != nil {
		return x.SpecName
	}
	return ""
}

func (x *SpecKeyValue) GetSpecValue() string {
	if x != nil {
		return x.SpecValue
	}
	return ""
}

type AttributeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AttributeName  string `protobuf:"bytes,1,opt,name=attribute_name,json=attributeName,proto3" json:"attribute_name,omitempty"`
	AttributeValue string `protobuf:"bytes,2,opt,name=attribute_value,json=attributeValue,proto3" json:"attribute_value,omitempty"`
}

func (x *AttributeInfo) Reset() {
	*x = AttributeInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_common_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttributeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttributeInfo) ProtoMessage() {}

func (x *AttributeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_goods_common_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttributeInfo.ProtoReflect.Descriptor instead.
func (*AttributeInfo) Descriptor() ([]byte, []int) {
	return file_goods_common_proto_rawDescGZIP(), []int{5}
}

func (x *AttributeInfo) GetAttributeName() string {
	if x != nil {
		return x.AttributeName
	}
	return ""
}

func (x *AttributeInfo) GetAttributeValue() string {
	if x != nil {
		return x.AttributeValue
	}
	return ""
}

// 关键词下拉框项
type Keyword struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Keyword) Reset() {
	*x = Keyword{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_common_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Keyword) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Keyword) ProtoMessage() {}

func (x *Keyword) ProtoReflect() protoreflect.Message {
	mi := &file_goods_common_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Keyword.ProtoReflect.Descriptor instead.
func (*Keyword) Descriptor() ([]byte, []int) {
	return file_goods_common_proto_rawDescGZIP(), []int{6}
}

func (x *Keyword) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Keyword) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// 商品分类
type Category struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CategoryId   uint64      `protobuf:"varint,1,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	CategoryName string      `protobuf:"bytes,2,opt,name=category_name,json=categoryName,proto3" json:"category_name,omitempty"`
	ParentId     uint64      `protobuf:"varint,3,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	Children     []*Category `protobuf:"bytes,4,rep,name=children,proto3" json:"children,omitempty"`
}

func (x *Category) Reset() {
	*x = Category{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_common_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Category) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Category) ProtoMessage() {}

func (x *Category) ProtoReflect() protoreflect.Message {
	mi := &file_goods_common_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Category.ProtoReflect.Descriptor instead.
func (*Category) Descriptor() ([]byte, []int) {
	return file_goods_common_proto_rawDescGZIP(), []int{7}
}

func (x *Category) GetCategoryId() uint64 {
	if x != nil {
		return x.CategoryId
	}
	return 0
}

func (x *Category) GetCategoryName() string {
	if x != nil {
		return x.CategoryName
	}
	return ""
}

func (x *Category) GetParentId() uint64 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *Category) GetChildren() []*Category {
	if x != nil {
		return x.Children
	}
	return nil
}

var File_goods_common_proto protoreflect.FileDescriptor

var file_goods_common_proto_rawDesc = []byte{
	0x0a, 0x12, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x22, 0xb8, 0x03, 0x0a, 0x07,
	0x53, 0x70, 0x75, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x70, 0x75, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x73, 0x70, 0x75, 0x49, 0x64, 0x12, 0x19,
	0x0a, 0x08, 0x73, 0x70, 0x75, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x73, 0x70, 0x75, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x70, 0x75,
	0x5f, 0x64, 0x65, 0x73, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x70, 0x75,
	0x44, 0x65, 0x73, 0x63, 0x12, 0x1e, 0x0a, 0x0b, 0x73, 0x70, 0x75, 0x5f, 0x69, 0x6d, 0x67, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x70, 0x75, 0x49, 0x6d,
	0x67, 0x55, 0x72, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x70, 0x75, 0x5f, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x70, 0x75, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x26, 0x0a, 0x0f, 0x73, 0x70, 0x75, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x73, 0x70, 0x75, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x73, 0x70, 0x75,
	0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x70, 0x75, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x5f, 0x69,
	0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x49, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x29, 0x0a, 0x08, 0x73, 0x6b, 0x75, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0b, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x53, 0x6b, 0x75, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x07, 0x73, 0x6b, 0x75, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x09, 0x73, 0x70,
	0x65, 0x63, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x53, 0x70, 0x65, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08,
	0x73, 0x70, 0x65, 0x63, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x0f, 0x73, 0x70, 0x65, 0x63,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0d, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x53, 0x70, 0x65, 0x63, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0d, 0x73, 0x70, 0x65, 0x63, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x22, 0xa9, 0x01, 0x0a, 0x07, 0x53, 0x6b, 0x75, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x6b, 0x75, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x05, 0x73, 0x6b, 0x75, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x70, 0x75,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x73, 0x70, 0x75, 0x49, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x08, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x6b, 0x75, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x6b, 0x75, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x08,
	0x73, 0x6b, 0x75, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x73, 0x6b, 0x75, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x70, 0x65, 0x63, 0x5f,
	0x73, 0x74, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x70, 0x65, 0x63, 0x53,
	0x74, 0x72, 0x22, 0x57, 0x0a, 0x08, 0x53, 0x70, 0x65, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x17,
	0x0a, 0x07, 0x73, 0x70, 0x65, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x06, 0x73, 0x70, 0x65, 0x63, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x70, 0x65, 0x63, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x70, 0x65, 0x63,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x70, 0x75, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x73, 0x70, 0x75, 0x49, 0x64, 0x22, 0x6f, 0x0a, 0x0d, 0x53,
	0x70, 0x65, 0x63, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x22, 0x0a, 0x0d,
	0x73, 0x70, 0x65, 0x63, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0b, 0x73, 0x70, 0x65, 0x63, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x49, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x73, 0x70, 0x65, 0x63, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x70, 0x65, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x70, 0x65, 0x63, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x70, 0x65, 0x63, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x4a, 0x0a, 0x0c,
	0x53, 0x70, 0x65, 0x63, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x73, 0x70, 0x65, 0x63, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x73, 0x70, 0x65, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x70, 0x65,
	0x63, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x70, 0x65, 0x63, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x5f, 0x0a, 0x0d, 0x41, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x27, 0x0a, 0x0f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x2f, 0x0a, 0x07, 0x4b, 0x65, 0x79,
	0x77, 0x6f, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x9a, 0x01, 0x0a, 0x08, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x08, 0x63, 0x68,
	0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x67,
	0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x08, 0x63,
	0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x7a, 0x6c, 0x2d, 0x68, 0x65, 0x72, 0x65, 0x2f, 0x62,
	0x74, 0x2d, 0x73, 0x68, 0x6f, 0x70, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x6b,
	0x69, 0x74, 0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_goods_common_proto_rawDescOnce sync.Once
	file_goods_common_proto_rawDescData = file_goods_common_proto_rawDesc
)

func file_goods_common_proto_rawDescGZIP() []byte {
	file_goods_common_proto_rawDescOnce.Do(func() {
		file_goods_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_goods_common_proto_rawDescData)
	})
	return file_goods_common_proto_rawDescData
}

var file_goods_common_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_goods_common_proto_goTypes = []interface{}{
	(*SpuInfo)(nil),       // 0: goods.SpuInfo
	(*SkuInfo)(nil),       // 1: goods.SkuInfo
	(*SpecInfo)(nil),      // 2: goods.SpecInfo
	(*SpecValueInfo)(nil), // 3: goods.SpecValueInfo
	(*SpecKeyValue)(nil),  // 4: goods.SpecKeyValue
	(*AttributeInfo)(nil), // 5: goods.AttributeInfo
	(*Keyword)(nil),       // 6: goods.Keyword
	(*Category)(nil),      // 7: goods.Category
}
var file_goods_common_proto_depIdxs = []int32{
	1, // 0: goods.SpuInfo.sku_list:type_name -> goods.SkuInfo
	2, // 1: goods.SpuInfo.spec_list:type_name -> goods.SpecInfo
	3, // 2: goods.SpuInfo.spec_value_list:type_name -> goods.SpecValueInfo
	7, // 3: goods.Category.children:type_name -> goods.Category
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_goods_common_proto_init() }
func file_goods_common_proto_init() {
	if File_goods_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_goods_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpuInfo); i {
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
		file_goods_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SkuInfo); i {
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
		file_goods_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpecInfo); i {
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
		file_goods_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpecValueInfo); i {
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
		file_goods_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpecKeyValue); i {
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
		file_goods_common_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttributeInfo); i {
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
		file_goods_common_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Keyword); i {
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
		file_goods_common_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Category); i {
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
			RawDescriptor: file_goods_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_goods_common_proto_goTypes,
		DependencyIndexes: file_goods_common_proto_depIdxs,
		MessageInfos:      file_goods_common_proto_msgTypes,
	}.Build()
	File_goods_common_proto = out.File
	file_goods_common_proto_rawDesc = nil
	file_goods_common_proto_goTypes = nil
	file_goods_common_proto_depIdxs = nil
}

var _ context.Context
