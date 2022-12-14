// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: internal/proto/bread.proto

package __

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

type SaveBreadReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *SaveBreadReq) Reset() {
	*x = SaveBreadReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_bread_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveBreadReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveBreadReq) ProtoMessage() {}

func (x *SaveBreadReq) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_bread_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveBreadReq.ProtoReflect.Descriptor instead.
func (*SaveBreadReq) Descriptor() ([]byte, []int) {
	return file_internal_proto_bread_proto_rawDescGZIP(), []int{0}
}

func (x *SaveBreadReq) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SaveBreadReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type QueryBreadReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Query:
	//
	//	*QueryBreadReq_Id
	//	*QueryBreadReq_Name
	Query isQueryBreadReq_Query `protobuf_oneof:"query"`
}

func (x *QueryBreadReq) Reset() {
	*x = QueryBreadReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_bread_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryBreadReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryBreadReq) ProtoMessage() {}

func (x *QueryBreadReq) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_bread_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryBreadReq.ProtoReflect.Descriptor instead.
func (*QueryBreadReq) Descriptor() ([]byte, []int) {
	return file_internal_proto_bread_proto_rawDescGZIP(), []int{1}
}

func (m *QueryBreadReq) GetQuery() isQueryBreadReq_Query {
	if m != nil {
		return m.Query
	}
	return nil
}

func (x *QueryBreadReq) GetId() uint64 {
	if x, ok := x.GetQuery().(*QueryBreadReq_Id); ok {
		return x.Id
	}
	return 0
}

func (x *QueryBreadReq) GetName() string {
	if x, ok := x.GetQuery().(*QueryBreadReq_Name); ok {
		return x.Name
	}
	return ""
}

type isQueryBreadReq_Query interface {
	isQueryBreadReq_Query()
}

type QueryBreadReq_Id struct {
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3,oneof"`
}

type QueryBreadReq_Name struct {
	Name string `protobuf:"bytes,2,opt,name=name,proto3,oneof"`
}

func (*QueryBreadReq_Id) isQueryBreadReq_Query() {}

func (*QueryBreadReq_Name) isQueryBreadReq_Query() {}

type QueryAllBreadReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *QueryAllBreadReq) Reset() {
	*x = QueryAllBreadReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_bread_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryAllBreadReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryAllBreadReq) ProtoMessage() {}

func (x *QueryAllBreadReq) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_bread_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryAllBreadReq.ProtoReflect.Descriptor instead.
func (*QueryAllBreadReq) Descriptor() ([]byte, []int) {
	return file_internal_proto_bread_proto_rawDescGZIP(), []int{2}
}

type BreadStore struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *BreadStore) Reset() {
	*x = BreadStore{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_bread_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BreadStore) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BreadStore) ProtoMessage() {}

func (x *BreadStore) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_bread_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BreadStore.ProtoReflect.Descriptor instead.
func (*BreadStore) Descriptor() ([]byte, []int) {
	return file_internal_proto_bread_proto_rawDescGZIP(), []int{3}
}

func (x *BreadStore) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BreadStore) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_internal_proto_bread_proto protoreflect.FileDescriptor

var file_internal_proto_bread_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x62, 0x72, 0x65, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x62, 0x72,
	0x65, 0x61, 0x64, 0x1a, 0x1b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x32, 0x0a, 0x0c, 0x53, 0x61, 0x76, 0x65, 0x42, 0x72, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x40, 0x0a, 0x0d, 0x51, 0x75, 0x65, 0x72, 0x79, 0x42, 0x72, 0x65,
	0x61, 0x64, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x07, 0x0a,
	0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x22, 0x12, 0x0a, 0x10, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41,
	0x6c, 0x6c, 0x42, 0x72, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x22, 0x30, 0x0a, 0x0a, 0x42, 0x72,
	0x65, 0x61, 0x64, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0xae, 0x01, 0x0a,
	0x05, 0x42, 0x72, 0x65, 0x61, 0x64, 0x12, 0x32, 0x0a, 0x04, 0x53, 0x61, 0x76, 0x65, 0x12, 0x13,
	0x2e, 0x62, 0x72, 0x65, 0x61, 0x64, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x42, 0x72, 0x65, 0x61, 0x64,
	0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x3b, 0x0a, 0x08, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x41, 0x6c, 0x6c, 0x12, 0x17, 0x2e, 0x62, 0x72, 0x65, 0x61, 0x64, 0x2e, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x41, 0x6c, 0x6c, 0x42, 0x72, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x1a,
	0x16, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x4d,
	0x61, 0x6e, 0x79, 0x52, 0x65, 0x73, 0x70, 0x12, 0x34, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x12, 0x14, 0x2e, 0x62, 0x72, 0x65, 0x61, 0x64, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x42, 0x72,
	0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x42, 0x03, 0x5a,
	0x01, 0x2e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_proto_bread_proto_rawDescOnce sync.Once
	file_internal_proto_bread_proto_rawDescData = file_internal_proto_bread_proto_rawDesc
)

func file_internal_proto_bread_proto_rawDescGZIP() []byte {
	file_internal_proto_bread_proto_rawDescOnce.Do(func() {
		file_internal_proto_bread_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_proto_bread_proto_rawDescData)
	})
	return file_internal_proto_bread_proto_rawDescData
}

var file_internal_proto_bread_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_internal_proto_bread_proto_goTypes = []interface{}{
	(*SaveBreadReq)(nil),     // 0: bread.SaveBreadReq
	(*QueryBreadReq)(nil),    // 1: bread.QueryBreadReq
	(*QueryAllBreadReq)(nil), // 2: bread.QueryAllBreadReq
	(*BreadStore)(nil),       // 3: bread.BreadStore
	(*CommonOneResp)(nil),    // 4: common.CommonOneResp
	(*CommonManyResp)(nil),   // 5: common.CommonManyResp
}
var file_internal_proto_bread_proto_depIdxs = []int32{
	0, // 0: bread.Bread.Save:input_type -> bread.SaveBreadReq
	2, // 1: bread.Bread.QueryAll:input_type -> bread.QueryAllBreadReq
	1, // 2: bread.Bread.Query:input_type -> bread.QueryBreadReq
	4, // 3: bread.Bread.Save:output_type -> common.CommonOneResp
	5, // 4: bread.Bread.QueryAll:output_type -> common.CommonManyResp
	4, // 5: bread.Bread.Query:output_type -> common.CommonOneResp
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_proto_bread_proto_init() }
func file_internal_proto_bread_proto_init() {
	if File_internal_proto_bread_proto != nil {
		return
	}
	file_internal_proto_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_internal_proto_bread_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveBreadReq); i {
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
		file_internal_proto_bread_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryBreadReq); i {
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
		file_internal_proto_bread_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryAllBreadReq); i {
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
		file_internal_proto_bread_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BreadStore); i {
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
	file_internal_proto_bread_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*QueryBreadReq_Id)(nil),
		(*QueryBreadReq_Name)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_proto_bread_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_proto_bread_proto_goTypes,
		DependencyIndexes: file_internal_proto_bread_proto_depIdxs,
		MessageInfos:      file_internal_proto_bread_proto_msgTypes,
	}.Build()
	File_internal_proto_bread_proto = out.File
	file_internal_proto_bread_proto_rawDesc = nil
	file_internal_proto_bread_proto_goTypes = nil
	file_internal_proto_bread_proto_depIdxs = nil
}
