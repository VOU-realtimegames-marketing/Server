// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.15.8
// source: counterpart/service_counterpart.proto

package gen

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_counterpart_service_counterpart_proto protoreflect.FileDescriptor

var file_counterpart_service_counterpart_proto_rawDesc = []byte{
	0x0a, 0x25, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x61, 0x72, 0x74, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x61, 0x72,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x76, 0x6f, 0x75, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x22, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x61, 0x72, 0x74, 0x2f,
	0x72, 0x70, 0x63, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x70,
	0x61, 0x72, 0x74, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x73, 0x5f, 0x6f, 0x66, 0x5f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x32, 0xc6, 0x01, 0x0a, 0x12, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x61, 0x72,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4e, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x1d, 0x2e, 0x76, 0x6f, 0x75, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x76, 0x6f, 0x75, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x60, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x4f, 0x66, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12,
	0x22, 0x2e, 0x76, 0x6f, 0x75, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x73, 0x4f, 0x66, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x76, 0x6f, 0x75, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x4f, 0x66, 0x4f, 0x77, 0x6e, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x16, 0x5a, 0x14, 0x56, 0x4f,
	0x55, 0x2d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_counterpart_service_counterpart_proto_goTypes = []any{
	(*CreateStoreRequest)(nil),       // 0: vou.proto.CreateStoreRequest
	(*GetStoresOfOwnerRequest)(nil),  // 1: vou.proto.GetStoresOfOwnerRequest
	(*CreateStoreResponse)(nil),      // 2: vou.proto.CreateStoreResponse
	(*GetStoresOfOwnerResponse)(nil), // 3: vou.proto.GetStoresOfOwnerResponse
}
var file_counterpart_service_counterpart_proto_depIdxs = []int32{
	0, // 0: vou.proto.CounterpartService.CreateStore:input_type -> vou.proto.CreateStoreRequest
	1, // 1: vou.proto.CounterpartService.GetAllStoresOfOwner:input_type -> vou.proto.GetStoresOfOwnerRequest
	2, // 2: vou.proto.CounterpartService.CreateStore:output_type -> vou.proto.CreateStoreResponse
	3, // 3: vou.proto.CounterpartService.GetAllStoresOfOwner:output_type -> vou.proto.GetStoresOfOwnerResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_counterpart_service_counterpart_proto_init() }
func file_counterpart_service_counterpart_proto_init() {
	if File_counterpart_service_counterpart_proto != nil {
		return
	}
	file_counterpart_rpc_create_store_proto_init()
	file_counterpart_rpc_get_stores_of_owner_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_counterpart_service_counterpart_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_counterpart_service_counterpart_proto_goTypes,
		DependencyIndexes: file_counterpart_service_counterpart_proto_depIdxs,
	}.Build()
	File_counterpart_service_counterpart_proto = out.File
	file_counterpart_service_counterpart_proto_rawDesc = nil
	file_counterpart_service_counterpart_proto_goTypes = nil
	file_counterpart_service_counterpart_proto_depIdxs = nil
}
