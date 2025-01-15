// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.3
// source: counterpart/rpc_update_store.proto

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

type UpdateStoreRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Owner         string                 `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Id            int64                  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Name          *string                `protobuf:"bytes,3,opt,name=name,proto3,oneof" json:"name,omitempty"`
	BusinessType  *string                `protobuf:"bytes,4,opt,name=business_type,json=businessType,proto3,oneof" json:"business_type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateStoreRequest) Reset() {
	*x = UpdateStoreRequest{}
	mi := &file_counterpart_rpc_update_store_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateStoreRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStoreRequest) ProtoMessage() {}

func (x *UpdateStoreRequest) ProtoReflect() protoreflect.Message {
	mi := &file_counterpart_rpc_update_store_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStoreRequest.ProtoReflect.Descriptor instead.
func (*UpdateStoreRequest) Descriptor() ([]byte, []int) {
	return file_counterpart_rpc_update_store_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateStoreRequest) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *UpdateStoreRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateStoreRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *UpdateStoreRequest) GetBusinessType() string {
	if x != nil && x.BusinessType != nil {
		return *x.BusinessType
	}
	return ""
}

type UpdateStoreResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Store         *Store                 `protobuf:"bytes,1,opt,name=store,proto3" json:"store,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateStoreResponse) Reset() {
	*x = UpdateStoreResponse{}
	mi := &file_counterpart_rpc_update_store_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateStoreResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStoreResponse) ProtoMessage() {}

func (x *UpdateStoreResponse) ProtoReflect() protoreflect.Message {
	mi := &file_counterpart_rpc_update_store_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStoreResponse.ProtoReflect.Descriptor instead.
func (*UpdateStoreResponse) Descriptor() ([]byte, []int) {
	return file_counterpart_rpc_update_store_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateStoreResponse) GetStore() *Store {
	if x != nil {
		return x.Store
	}
	return nil
}

var File_counterpart_rpc_update_store_proto protoreflect.FileDescriptor

var file_counterpart_rpc_update_store_proto_rawDesc = []byte{
	0x0a, 0x22, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x61, 0x72, 0x74, 0x2f, 0x72, 0x70,
	0x63, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x76, 0x6f, 0x75, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x61, 0x72, 0x74, 0x2f, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x98, 0x01, 0x0a, 0x12, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x28,
	0x0a, 0x0d, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0c, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73,
	0x73, 0x54, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x22, 0x3d, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x05, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x76, 0x6f, 0x75, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x05, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x42, 0x16, 0x5a, 0x14, 0x56, 0x4f, 0x55, 0x2d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_counterpart_rpc_update_store_proto_rawDescOnce sync.Once
	file_counterpart_rpc_update_store_proto_rawDescData = file_counterpart_rpc_update_store_proto_rawDesc
)

func file_counterpart_rpc_update_store_proto_rawDescGZIP() []byte {
	file_counterpart_rpc_update_store_proto_rawDescOnce.Do(func() {
		file_counterpart_rpc_update_store_proto_rawDescData = protoimpl.X.CompressGZIP(file_counterpart_rpc_update_store_proto_rawDescData)
	})
	return file_counterpart_rpc_update_store_proto_rawDescData
}

var file_counterpart_rpc_update_store_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_counterpart_rpc_update_store_proto_goTypes = []any{
	(*UpdateStoreRequest)(nil),  // 0: vou.proto.UpdateStoreRequest
	(*UpdateStoreResponse)(nil), // 1: vou.proto.UpdateStoreResponse
	(*Store)(nil),               // 2: vou.proto.Store
}
var file_counterpart_rpc_update_store_proto_depIdxs = []int32{
	2, // 0: vou.proto.UpdateStoreResponse.store:type_name -> vou.proto.Store
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_counterpart_rpc_update_store_proto_init() }
func file_counterpart_rpc_update_store_proto_init() {
	if File_counterpart_rpc_update_store_proto != nil {
		return
	}
	file_counterpart_store_proto_init()
	file_counterpart_rpc_update_store_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_counterpart_rpc_update_store_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_counterpart_rpc_update_store_proto_goTypes,
		DependencyIndexes: file_counterpart_rpc_update_store_proto_depIdxs,
		MessageInfos:      file_counterpart_rpc_update_store_proto_msgTypes,
	}.Build()
	File_counterpart_rpc_update_store_proto = out.File
	file_counterpart_rpc_update_store_proto_rawDesc = nil
	file_counterpart_rpc_update_store_proto_goTypes = nil
	file_counterpart_rpc_update_store_proto_depIdxs = nil
}
