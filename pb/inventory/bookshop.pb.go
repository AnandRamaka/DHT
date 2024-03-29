// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: bookshop.proto

package inventory

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

// Get
type UrlRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key int32 `protobuf:"varint,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *UrlRequest) Reset() {
	*x = UrlRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bookshop_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UrlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UrlRequest) ProtoMessage() {}

func (x *UrlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bookshop_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UrlRequest.ProtoReflect.Descriptor instead.
func (*UrlRequest) Descriptor() ([]byte, []int) {
	return file_bookshop_proto_rawDescGZIP(), []int{0}
}

func (x *UrlRequest) GetKey() int32 {
	if x != nil {
		return x.Key
	}
	return 0
}

// Get URL
type UrlResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *UrlResponse) Reset() {
	*x = UrlResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bookshop_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UrlResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UrlResponse) ProtoMessage() {}

func (x *UrlResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bookshop_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UrlResponse.ProtoReflect.Descriptor instead.
func (*UrlResponse) Descriptor() ([]byte, []int) {
	return file_bookshop_proto_rawDescGZIP(), []int{1}
}

func (x *UrlResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type ValueResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ValueResponse) Reset() {
	*x = ValueResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bookshop_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValueResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValueResponse) ProtoMessage() {}

func (x *ValueResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bookshop_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValueResponse.ProtoReflect.Descriptor instead.
func (*ValueResponse) Descriptor() ([]byte, []int) {
	return file_bookshop_proto_rawDescGZIP(), []int{2}
}

func (x *ValueResponse) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// Insert kv
type InsertValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   int32  `protobuf:"varint,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *InsertValue) Reset() {
	*x = InsertValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bookshop_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertValue) ProtoMessage() {}

func (x *InsertValue) ProtoReflect() protoreflect.Message {
	mi := &file_bookshop_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertValue.ProtoReflect.Descriptor instead.
func (*InsertValue) Descriptor() ([]byte, []int) {
	return file_bookshop_proto_rawDescGZIP(), []int{3}
}

func (x *InsertValue) GetKey() int32 {
	if x != nil {
		return x.Key
	}
	return 0
}

func (x *InsertValue) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_bookshop_proto protoreflect.FileDescriptor

var file_bookshop_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x1e, 0x0a, 0x0a, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x22, 0x1f, 0x0a, 0x0b, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72,
	0x6c, 0x22, 0x25, 0x0a, 0x0d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x35, 0x0a, 0x0b, 0x49, 0x6e, 0x73, 0x65,
	0x72, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x32,
	0x87, 0x01, 0x0a, 0x09, 0x48, 0x61, 0x73, 0x68, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x25, 0x0a,
	0x06, 0x47, 0x65, 0x74, 0x55, 0x52, 0x4c, 0x12, 0x0b, 0x2e, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x29, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x0b, 0x2e, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x28, 0x0a, 0x08, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x4b, 0x76, 0x12, 0x0c, 0x2e, 0x49, 0x6e,
	0x73, 0x65, 0x72, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x0c, 0x2e, 0x55, 0x72, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0e, 0x5a, 0x0c, 0x70, 0x62, 0x2f,
	0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_bookshop_proto_rawDescOnce sync.Once
	file_bookshop_proto_rawDescData = file_bookshop_proto_rawDesc
)

func file_bookshop_proto_rawDescGZIP() []byte {
	file_bookshop_proto_rawDescOnce.Do(func() {
		file_bookshop_proto_rawDescData = protoimpl.X.CompressGZIP(file_bookshop_proto_rawDescData)
	})
	return file_bookshop_proto_rawDescData
}

var file_bookshop_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_bookshop_proto_goTypes = []interface{}{
	(*UrlRequest)(nil),    // 0: UrlRequest
	(*UrlResponse)(nil),   // 1: UrlResponse
	(*ValueResponse)(nil), // 2: ValueResponse
	(*InsertValue)(nil),   // 3: InsertValue
}
var file_bookshop_proto_depIdxs = []int32{
	0, // 0: HashTable.GetURL:input_type -> UrlRequest
	0, // 1: HashTable.GetValue:input_type -> UrlRequest
	3, // 2: HashTable.InsertKv:input_type -> InsertValue
	1, // 3: HashTable.GetURL:output_type -> UrlResponse
	2, // 4: HashTable.GetValue:output_type -> ValueResponse
	1, // 5: HashTable.InsertKv:output_type -> UrlResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_bookshop_proto_init() }
func file_bookshop_proto_init() {
	if File_bookshop_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bookshop_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UrlRequest); i {
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
		file_bookshop_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UrlResponse); i {
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
		file_bookshop_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValueResponse); i {
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
		file_bookshop_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsertValue); i {
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
			RawDescriptor: file_bookshop_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bookshop_proto_goTypes,
		DependencyIndexes: file_bookshop_proto_depIdxs,
		MessageInfos:      file_bookshop_proto_msgTypes,
	}.Build()
	File_bookshop_proto = out.File
	file_bookshop_proto_rawDesc = nil
	file_bookshop_proto_goTypes = nil
	file_bookshop_proto_depIdxs = nil
}
