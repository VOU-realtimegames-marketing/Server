// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.3
// source: counterpart/rpc_get_branchs.proto

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

type GetBranchsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	StoreId       int64                  `protobuf:"varint,1,opt,name=store_id,json=storeId,proto3" json:"store_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetBranchsRequest) Reset() {
	*x = GetBranchsRequest{}
	mi := &file_counterpart_rpc_get_branchs_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBranchsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBranchsRequest) ProtoMessage() {}

func (x *GetBranchsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_counterpart_rpc_get_branchs_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBranchsRequest.ProtoReflect.Descriptor instead.
func (*GetBranchsRequest) Descriptor() ([]byte, []int) {
	return file_counterpart_rpc_get_branchs_proto_rawDescGZIP(), []int{0}
}

func (x *GetBranchsRequest) GetStoreId() int64 {
	if x != nil {
		return x.StoreId
	}
	return 0
}

type GetBranchsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Branchs       []*Branch              `protobuf:"bytes,1,rep,name=branchs,proto3" json:"branchs,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetBranchsResponse) Reset() {
	*x = GetBranchsResponse{}
	mi := &file_counterpart_rpc_get_branchs_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBranchsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBranchsResponse) ProtoMessage() {}

func (x *GetBranchsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_counterpart_rpc_get_branchs_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBranchsResponse.ProtoReflect.Descriptor instead.
func (*GetBranchsResponse) Descriptor() ([]byte, []int) {
	return file_counterpart_rpc_get_branchs_proto_rawDescGZIP(), []int{1}
}

func (x *GetBranchsResponse) GetBranchs() []*Branch {
	if x != nil {
		return x.Branchs
	}
	return nil
}

var File_counterpart_rpc_get_branchs_proto protoreflect.FileDescriptor

var file_counterpart_rpc_get_branchs_proto_rawDesc = []byte{
	0x0a, 0x21, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x61, 0x72, 0x74, 0x2f, 0x72, 0x70,
	0x63, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x09, 0x76, 0x6f, 0x75, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x61, 0x72, 0x74, 0x2f, 0x62, 0x72, 0x61, 0x6e,
	0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2e, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x42,
	0x72, 0x61, 0x6e, 0x63, 0x68, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a,
	0x08, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x22, 0x41, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x42,
	0x72, 0x61, 0x6e, 0x63, 0x68, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b,
	0x0a, 0x07, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x76, 0x6f, 0x75, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x72, 0x61, 0x6e,
	0x63, 0x68, 0x52, 0x07, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x73, 0x42, 0x16, 0x5a, 0x14, 0x56,
	0x4f, 0x55, 0x2d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_counterpart_rpc_get_branchs_proto_rawDescOnce sync.Once
	file_counterpart_rpc_get_branchs_proto_rawDescData = file_counterpart_rpc_get_branchs_proto_rawDesc
)

func file_counterpart_rpc_get_branchs_proto_rawDescGZIP() []byte {
	file_counterpart_rpc_get_branchs_proto_rawDescOnce.Do(func() {
		file_counterpart_rpc_get_branchs_proto_rawDescData = protoimpl.X.CompressGZIP(file_counterpart_rpc_get_branchs_proto_rawDescData)
	})
	return file_counterpart_rpc_get_branchs_proto_rawDescData
}

var file_counterpart_rpc_get_branchs_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_counterpart_rpc_get_branchs_proto_goTypes = []any{
	(*GetBranchsRequest)(nil),  // 0: vou.proto.GetBranchsRequest
	(*GetBranchsResponse)(nil), // 1: vou.proto.GetBranchsResponse
	(*Branch)(nil),             // 2: vou.proto.Branch
}
var file_counterpart_rpc_get_branchs_proto_depIdxs = []int32{
	2, // 0: vou.proto.GetBranchsResponse.branchs:type_name -> vou.proto.Branch
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_counterpart_rpc_get_branchs_proto_init() }
func file_counterpart_rpc_get_branchs_proto_init() {
	if File_counterpart_rpc_get_branchs_proto != nil {
		return
	}
	file_counterpart_branch_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_counterpart_rpc_get_branchs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_counterpart_rpc_get_branchs_proto_goTypes,
		DependencyIndexes: file_counterpart_rpc_get_branchs_proto_depIdxs,
		MessageInfos:      file_counterpart_rpc_get_branchs_proto_msgTypes,
	}.Build()
	File_counterpart_rpc_get_branchs_proto = out.File
	file_counterpart_rpc_get_branchs_proto_rawDesc = nil
	file_counterpart_rpc_get_branchs_proto_goTypes = nil
	file_counterpart_rpc_get_branchs_proto_depIdxs = nil
}
