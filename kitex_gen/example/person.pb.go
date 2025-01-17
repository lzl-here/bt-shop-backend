// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.29.2
// source: example/person.proto

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
type GetPersonReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid uint64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *GetPersonReq) Reset() {
	*x = GetPersonReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_person_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPersonReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPersonReq) ProtoMessage() {}

func (x *GetPersonReq) ProtoReflect() protoreflect.Message {
	mi := &file_example_person_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPersonReq.ProtoReflect.Descriptor instead.
func (*GetPersonReq) Descriptor() ([]byte, []int) {
	return file_example_person_proto_rawDescGZIP(), []int{0}
}

func (x *GetPersonReq) GetUid() uint64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type GetPersonRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code  int32                    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`               // 错误码
	Msg   string                   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`                  // 信息
	LogId string                   `protobuf:"bytes,3,opt,name=log_id,json=logId,proto3" json:"log_id,omitempty"` // 链路追踪id
	Data  *GetPersonRsp_PersonData `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`                // 响应体数据
}

func (x *GetPersonRsp) Reset() {
	*x = GetPersonRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_person_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPersonRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPersonRsp) ProtoMessage() {}

func (x *GetPersonRsp) ProtoReflect() protoreflect.Message {
	mi := &file_example_person_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPersonRsp.ProtoReflect.Descriptor instead.
func (*GetPersonRsp) Descriptor() ([]byte, []int) {
	return file_example_person_proto_rawDescGZIP(), []int{1}
}

func (x *GetPersonRsp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetPersonRsp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *GetPersonRsp) GetLogId() string {
	if x != nil {
		return x.LogId
	}
	return ""
}

func (x *GetPersonRsp) GetData() *GetPersonRsp_PersonData {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetPersonRsp_PersonData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid       uint64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Age       int32  `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
	Phone     string `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	Email     string `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Password  string `protobuf:"bytes,6,opt,name=password,proto3" json:"password,omitempty"`
	CreatedAt string `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt string `protobuf:"bytes,9,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *GetPersonRsp_PersonData) Reset() {
	*x = GetPersonRsp_PersonData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_person_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPersonRsp_PersonData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPersonRsp_PersonData) ProtoMessage() {}

func (x *GetPersonRsp_PersonData) ProtoReflect() protoreflect.Message {
	mi := &file_example_person_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPersonRsp_PersonData.ProtoReflect.Descriptor instead.
func (*GetPersonRsp_PersonData) Descriptor() ([]byte, []int) {
	return file_example_person_proto_rawDescGZIP(), []int{1, 0}
}

func (x *GetPersonRsp_PersonData) GetUid() uint64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *GetPersonRsp_PersonData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetPersonRsp_PersonData) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *GetPersonRsp_PersonData) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *GetPersonRsp_PersonData) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *GetPersonRsp_PersonData) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *GetPersonRsp_PersonData) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *GetPersonRsp_PersonData) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *GetPersonRsp_PersonData) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

var File_example_person_proto protoreflect.FileDescriptor

var file_example_person_proto_rawDesc = []byte{
	0x0a, 0x14, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x22,
	0x20, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x75, 0x69,
	0x64, 0x22, 0xed, 0x02, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52,
	0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x15, 0x0a, 0x06, 0x6c, 0x6f, 0x67, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x12,
	0x34, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x52, 0x73, 0x70, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0xe9, 0x01, 0x0a, 0x0a, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6c, 0x7a, 0x6c, 0x2d, 0x68, 0x65, 0x72, 0x65, 0x2f, 0x62, 0x74, 0x2d, 0x73, 0x68, 0x6f, 0x70,
	0x2d, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78, 0x5f, 0x67,
	0x65, 0x6e, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_example_person_proto_rawDescOnce sync.Once
	file_example_person_proto_rawDescData = file_example_person_proto_rawDesc
)

func file_example_person_proto_rawDescGZIP() []byte {
	file_example_person_proto_rawDescOnce.Do(func() {
		file_example_person_proto_rawDescData = protoimpl.X.CompressGZIP(file_example_person_proto_rawDescData)
	})
	return file_example_person_proto_rawDescData
}

var file_example_person_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_example_person_proto_goTypes = []interface{}{
	(*GetPersonReq)(nil),            // 0: example.GetPersonReq
	(*GetPersonRsp)(nil),            // 1: example.GetPersonRsp
	(*GetPersonRsp_PersonData)(nil), // 2: example.GetPersonRsp.PersonData
}
var file_example_person_proto_depIdxs = []int32{
	2, // 0: example.GetPersonRsp.data:type_name -> example.GetPersonRsp.PersonData
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_example_person_proto_init() }
func file_example_person_proto_init() {
	if File_example_person_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_example_person_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPersonReq); i {
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
		file_example_person_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPersonRsp); i {
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
		file_example_person_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPersonRsp_PersonData); i {
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
			RawDescriptor: file_example_person_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_example_person_proto_goTypes,
		DependencyIndexes: file_example_person_proto_depIdxs,
		MessageInfos:      file_example_person_proto_msgTypes,
	}.Build()
	File_example_person_proto = out.File
	file_example_person_proto_rawDesc = nil
	file_example_person_proto_goTypes = nil
	file_example_person_proto_depIdxs = nil
}

var _ context.Context
