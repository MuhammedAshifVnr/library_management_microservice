// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.23.4
// source: proto/book_management.proto

package proto

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

type CheckOutBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	BookId int64 `protobuf:"varint,2,opt,name=book_id,json=bookId,proto3" json:"book_id,omitempty"`
}

func (x *CheckOutBookRequest) Reset() {
	*x = CheckOutBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_book_management_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckOutBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckOutBookRequest) ProtoMessage() {}

func (x *CheckOutBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_book_management_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckOutBookRequest.ProtoReflect.Descriptor instead.
func (*CheckOutBookRequest) Descriptor() ([]byte, []int) {
	return file_proto_book_management_proto_rawDescGZIP(), []int{0}
}

func (x *CheckOutBookRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CheckOutBookRequest) GetBookId() int64 {
	if x != nil {
		return x.BookId
	}
	return 0
}

type CheckOutBookResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *CheckOutBookResponse) Reset() {
	*x = CheckOutBookResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_book_management_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckOutBookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckOutBookResponse) ProtoMessage() {}

func (x *CheckOutBookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_book_management_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckOutBookResponse.ProtoReflect.Descriptor instead.
func (*CheckOutBookResponse) Descriptor() ([]byte, []int) {
	return file_proto_book_management_proto_rawDescGZIP(), []int{1}
}

func (x *CheckOutBookResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type CheckInBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	BookId int64 `protobuf:"varint,2,opt,name=book_id,json=bookId,proto3" json:"book_id,omitempty"`
}

func (x *CheckInBookRequest) Reset() {
	*x = CheckInBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_book_management_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckInBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckInBookRequest) ProtoMessage() {}

func (x *CheckInBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_book_management_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckInBookRequest.ProtoReflect.Descriptor instead.
func (*CheckInBookRequest) Descriptor() ([]byte, []int) {
	return file_proto_book_management_proto_rawDescGZIP(), []int{2}
}

func (x *CheckInBookRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CheckInBookRequest) GetBookId() int64 {
	if x != nil {
		return x.BookId
	}
	return 0
}

type CheckInBookResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *CheckInBookResponse) Reset() {
	*x = CheckInBookResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_book_management_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckInBookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckInBookResponse) ProtoMessage() {}

func (x *CheckInBookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_book_management_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckInBookResponse.ProtoReflect.Descriptor instead.
func (*CheckInBookResponse) Descriptor() ([]byte, []int) {
	return file_proto_book_management_proto_rawDescGZIP(), []int{3}
}

func (x *CheckInBookResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_proto_book_management_proto protoreflect.FileDescriptor

var file_proto_book_management_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x62,
	0x6f, 0x6f, 0x6b, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x47,
	0x0a, 0x13, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17,
	0x0a, 0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x22, 0x30, 0x0a, 0x14, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x4f, 0x75, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x46, 0x0a, 0x12, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x49, 0x6e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x62, 0x6f, 0x6f, 0x6b,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x49,
	0x64, 0x22, 0x2f, 0x0a, 0x13, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x42, 0x6f, 0x6f, 0x6b,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x32, 0xce, 0x01, 0x0a, 0x15, 0x42, 0x6f, 0x6f, 0x6b, 0x4d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5b, 0x0a, 0x0c,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x24, 0x2e, 0x62,
	0x6f, 0x6f, 0x6b, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x25, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x42, 0x6f, 0x6f,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x58, 0x0a, 0x0b, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x49, 0x6e, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x23, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x5f,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x49, 0x6e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e,
	0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_book_management_proto_rawDescOnce sync.Once
	file_proto_book_management_proto_rawDescData = file_proto_book_management_proto_rawDesc
)

func file_proto_book_management_proto_rawDescGZIP() []byte {
	file_proto_book_management_proto_rawDescOnce.Do(func() {
		file_proto_book_management_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_book_management_proto_rawDescData)
	})
	return file_proto_book_management_proto_rawDescData
}

var file_proto_book_management_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_book_management_proto_goTypes = []any{
	(*CheckOutBookRequest)(nil),  // 0: book_management.CheckOutBookRequest
	(*CheckOutBookResponse)(nil), // 1: book_management.CheckOutBookResponse
	(*CheckInBookRequest)(nil),   // 2: book_management.CheckInBookRequest
	(*CheckInBookResponse)(nil),  // 3: book_management.CheckInBookResponse
}
var file_proto_book_management_proto_depIdxs = []int32{
	0, // 0: book_management.BookManagementService.CheckOutBook:input_type -> book_management.CheckOutBookRequest
	2, // 1: book_management.BookManagementService.CheckInBook:input_type -> book_management.CheckInBookRequest
	1, // 2: book_management.BookManagementService.CheckOutBook:output_type -> book_management.CheckOutBookResponse
	3, // 3: book_management.BookManagementService.CheckInBook:output_type -> book_management.CheckInBookResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_book_management_proto_init() }
func file_proto_book_management_proto_init() {
	if File_proto_book_management_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_book_management_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CheckOutBookRequest); i {
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
		file_proto_book_management_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CheckOutBookResponse); i {
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
		file_proto_book_management_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CheckInBookRequest); i {
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
		file_proto_book_management_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*CheckInBookResponse); i {
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
			RawDescriptor: file_proto_book_management_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_book_management_proto_goTypes,
		DependencyIndexes: file_proto_book_management_proto_depIdxs,
		MessageInfos:      file_proto_book_management_proto_msgTypes,
	}.Build()
	File_proto_book_management_proto = out.File
	file_proto_book_management_proto_rawDesc = nil
	file_proto_book_management_proto_goTypes = nil
	file_proto_book_management_proto_depIdxs = nil
}
