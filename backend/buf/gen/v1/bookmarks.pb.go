// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        (unknown)
// source: v1/bookmarks.proto

package v1

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

type AddBookmarkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddBookmarkRequest) Reset() {
	*x = AddBookmarkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_bookmarks_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddBookmarkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddBookmarkRequest) ProtoMessage() {}

func (x *AddBookmarkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_bookmarks_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddBookmarkRequest.ProtoReflect.Descriptor instead.
func (*AddBookmarkRequest) Descriptor() ([]byte, []int) {
	return file_v1_bookmarks_proto_rawDescGZIP(), []int{0}
}

type AddBookmarResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddBookmarResponse) Reset() {
	*x = AddBookmarResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_bookmarks_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddBookmarResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddBookmarResponse) ProtoMessage() {}

func (x *AddBookmarResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_bookmarks_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddBookmarResponse.ProtoReflect.Descriptor instead.
func (*AddBookmarResponse) Descriptor() ([]byte, []int) {
	return file_v1_bookmarks_proto_rawDescGZIP(), []int{1}
}

var File_v1_bookmarks_proto protoreflect.FileDescriptor

var file_v1_bookmarks_proto_rawDesc = []byte{
	0x0a, 0x12, 0x76, 0x31, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x22, 0x14, 0x0a, 0x12, 0x41, 0x64, 0x64, 0x42,
	0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x14,
	0x0a, 0x12, 0x41, 0x64, 0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x32, 0x4c, 0x0a, 0x09, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b,
	0x73, 0x12, 0x3f, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b,
	0x12, 0x16, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64,
	0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x73, 0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x31, 0x42, 0x0e, 0x42, 0x6f,
	0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x31,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x75, 0x65, 0x73,
	0x61, 0x61, 0x2f, 0x74, 0x65, 0x61, 0x74, 0x69, 0x6d, 0x65, 0x2d, 0x61, 0x70, 0x70, 0x2f, 0x62,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x62, 0x75, 0x66, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x76,
	0x31, 0xa2, 0x02, 0x03, 0x56, 0x58, 0x58, 0xaa, 0x02, 0x02, 0x56, 0x31, 0xca, 0x02, 0x02, 0x56,
	0x31, 0xe2, 0x02, 0x0e, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x02, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_bookmarks_proto_rawDescOnce sync.Once
	file_v1_bookmarks_proto_rawDescData = file_v1_bookmarks_proto_rawDesc
)

func file_v1_bookmarks_proto_rawDescGZIP() []byte {
	file_v1_bookmarks_proto_rawDescOnce.Do(func() {
		file_v1_bookmarks_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_bookmarks_proto_rawDescData)
	})
	return file_v1_bookmarks_proto_rawDescData
}

var file_v1_bookmarks_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_v1_bookmarks_proto_goTypes = []interface{}{
	(*AddBookmarkRequest)(nil), // 0: v1.AddBookmarkRequest
	(*AddBookmarResponse)(nil), // 1: v1.AddBookmarResponse
}
var file_v1_bookmarks_proto_depIdxs = []int32{
	0, // 0: v1.Bookmarks.AddBookmark:input_type -> v1.AddBookmarkRequest
	1, // 1: v1.Bookmarks.AddBookmark:output_type -> v1.AddBookmarResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_v1_bookmarks_proto_init() }
func file_v1_bookmarks_proto_init() {
	if File_v1_bookmarks_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_bookmarks_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddBookmarkRequest); i {
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
		file_v1_bookmarks_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddBookmarResponse); i {
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
			RawDescriptor: file_v1_bookmarks_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_bookmarks_proto_goTypes,
		DependencyIndexes: file_v1_bookmarks_proto_depIdxs,
		MessageInfos:      file_v1_bookmarks_proto_msgTypes,
	}.Build()
	File_v1_bookmarks_proto = out.File
	file_v1_bookmarks_proto_rawDesc = nil
	file_v1_bookmarks_proto_goTypes = nil
	file_v1_bookmarks_proto_depIdxs = nil
}
