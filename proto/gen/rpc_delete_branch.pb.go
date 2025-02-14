// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.3
// source: counterpart/rpc_delete_branch.proto

package gen

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

type DeleteBranchRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteBranchRequest) Reset() {
	*x = DeleteBranchRequest{}
	mi := &file_counterpart_rpc_delete_branch_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteBranchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBranchRequest) ProtoMessage() {}

func (x *DeleteBranchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_counterpart_rpc_delete_branch_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBranchRequest.ProtoReflect.Descriptor instead.
func (*DeleteBranchRequest) Descriptor() ([]byte, []int) {
	return file_counterpart_rpc_delete_branch_proto_rawDescGZIP(), []int{0}
}

func (x *DeleteBranchRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteBranchResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteBranchResponse) Reset() {
	*x = DeleteBranchResponse{}
	mi := &file_counterpart_rpc_delete_branch_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteBranchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBranchResponse) ProtoMessage() {}

func (x *DeleteBranchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_counterpart_rpc_delete_branch_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBranchResponse.ProtoReflect.Descriptor instead.
func (*DeleteBranchResponse) Descriptor() ([]byte, []int) {
	return file_counterpart_rpc_delete_branch_proto_rawDescGZIP(), []int{1}
}

var File_counterpart_rpc_delete_branch_proto protoreflect.FileDescriptor

var file_counterpart_rpc_delete_branch_proto_rawDesc = []byte{
	0x0a, 0x23, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x61, 0x72, 0x74, 0x2f, 0x72, 0x70,
	0x63, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x76, 0x6f, 0x75, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x25, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x16, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x16, 0x5a, 0x14, 0x56, 0x4f, 0x55, 0x2d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_counterpart_rpc_delete_branch_proto_rawDescOnce sync.Once
	file_counterpart_rpc_delete_branch_proto_rawDescData = file_counterpart_rpc_delete_branch_proto_rawDesc
)

func file_counterpart_rpc_delete_branch_proto_rawDescGZIP() []byte {
	file_counterpart_rpc_delete_branch_proto_rawDescOnce.Do(func() {
		file_counterpart_rpc_delete_branch_proto_rawDescData = protoimpl.X.CompressGZIP(file_counterpart_rpc_delete_branch_proto_rawDescData)
	})
	return file_counterpart_rpc_delete_branch_proto_rawDescData
}

var file_counterpart_rpc_delete_branch_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_counterpart_rpc_delete_branch_proto_goTypes = []any{
	(*DeleteBranchRequest)(nil),  // 0: vou.proto.DeleteBranchRequest
	(*DeleteBranchResponse)(nil), // 1: vou.proto.DeleteBranchResponse
}
var file_counterpart_rpc_delete_branch_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_counterpart_rpc_delete_branch_proto_init() }
func file_counterpart_rpc_delete_branch_proto_init() {
	if File_counterpart_rpc_delete_branch_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_counterpart_rpc_delete_branch_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_counterpart_rpc_delete_branch_proto_goTypes,
		DependencyIndexes: file_counterpart_rpc_delete_branch_proto_depIdxs,
		MessageInfos:      file_counterpart_rpc_delete_branch_proto_msgTypes,
	}.Build()
	File_counterpart_rpc_delete_branch_proto = out.File
	file_counterpart_rpc_delete_branch_proto_rawDesc = nil
	file_counterpart_rpc_delete_branch_proto_goTypes = nil
	file_counterpart_rpc_delete_branch_proto_depIdxs = nil
}
