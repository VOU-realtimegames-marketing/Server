// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.3
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
	0x6f, 0x1a, 0x22, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x61, 0x72, 0x74, 0x2f, 0x72,
	0x70, 0x63, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x61,
	0x72, 0x74, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x65, 0x72, 0x70, 0x61, 0x72, 0x74, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x5f, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x61, 0x72, 0x74, 0x2f, 0x72, 0x70, 0x63, 0x5f,
	0x67, 0x65, 0x74, 0x5f, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x23, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x61, 0x72, 0x74, 0x2f, 0x72,
	0x70, 0x63, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x70,
	0x61, 0x72, 0x74, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x32, 0x8e, 0x06, 0x0a, 0x12, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x61, 0x72,
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
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x0b, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x1d, 0x2e, 0x76, 0x6f, 0x75, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x76, 0x6f, 0x75, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x0b, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x1d, 0x2e, 0x76, 0x6f, 0x75, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x76, 0x6f, 0x75, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x0c, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x1e, 0x2e, 0x76, 0x6f, 0x75,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x72, 0x61,
	0x6e, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x76, 0x6f, 0x75,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x72, 0x61,
	0x6e, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4b, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x73, 0x12, 0x1c, 0x2e, 0x76, 0x6f,
	0x75, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x63,
	0x68, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x76, 0x6f, 0x75, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x0c, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x1e, 0x2e, 0x76, 0x6f, 0x75,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x72, 0x61,
	0x6e, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x76, 0x6f, 0x75,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x72, 0x61,
	0x6e, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x57, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x43, 0x6d, 0x73, 0x4f, 0x76, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x12,
	0x20, 0x2e, 0x76, 0x6f, 0x75, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x6d, 0x73, 0x4f, 0x76, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x21, 0x2e, 0x76, 0x6f, 0x75, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65,
	0x74, 0x43, 0x6d, 0x73, 0x4f, 0x76, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5a, 0x0a, 0x0f, 0x46, 0x61, 0x6b, 0x65, 0x43, 0x6d,
	0x73, 0x4f, 0x76, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x12, 0x21, 0x2e, 0x76, 0x6f, 0x75, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x61, 0x6b, 0x65, 0x43, 0x6d, 0x73, 0x4f, 0x76, 0x65,
	0x72, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x76,
	0x6f, 0x75, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x61, 0x6b, 0x65, 0x43, 0x6d, 0x73,
	0x4f, 0x76, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x16, 0x5a, 0x14, 0x56, 0x4f, 0x55, 0x2d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var file_counterpart_service_counterpart_proto_goTypes = []any{
	(*CreateStoreRequest)(nil),       // 0: vou.proto.CreateStoreRequest
	(*GetStoresOfOwnerRequest)(nil),  // 1: vou.proto.GetStoresOfOwnerRequest
	(*UpdateStoreRequest)(nil),       // 2: vou.proto.UpdateStoreRequest
	(*DeleteStoreRequest)(nil),       // 3: vou.proto.DeleteStoreRequest
	(*CreateBranchRequest)(nil),      // 4: vou.proto.CreateBranchRequest
	(*GetBranchsRequest)(nil),        // 5: vou.proto.GetBranchsRequest
	(*DeleteBranchRequest)(nil),      // 6: vou.proto.DeleteBranchRequest
	(*GetCmsOverviewRequest)(nil),    // 7: vou.proto.GetCmsOverviewRequest
	(*FakeCmsOverviewRequest)(nil),   // 8: vou.proto.FakeCmsOverviewRequest
	(*CreateStoreResponse)(nil),      // 9: vou.proto.CreateStoreResponse
	(*GetStoresOfOwnerResponse)(nil), // 10: vou.proto.GetStoresOfOwnerResponse
	(*UpdateStoreResponse)(nil),      // 11: vou.proto.UpdateStoreResponse
	(*DeleteStoreResponse)(nil),      // 12: vou.proto.DeleteStoreResponse
	(*CreateBranchResponse)(nil),     // 13: vou.proto.CreateBranchResponse
	(*GetBranchsResponse)(nil),       // 14: vou.proto.GetBranchsResponse
	(*DeleteBranchResponse)(nil),     // 15: vou.proto.DeleteBranchResponse
	(*GetCmsOverviewResponse)(nil),   // 16: vou.proto.GetCmsOverviewResponse
	(*FakeCmsOverviewResponse)(nil),  // 17: vou.proto.FakeCmsOverviewResponse
}
var file_counterpart_service_counterpart_proto_depIdxs = []int32{
	0,  // 0: vou.proto.CounterpartService.CreateStore:input_type -> vou.proto.CreateStoreRequest
	1,  // 1: vou.proto.CounterpartService.GetAllStoresOfOwner:input_type -> vou.proto.GetStoresOfOwnerRequest
	2,  // 2: vou.proto.CounterpartService.UpdateStore:input_type -> vou.proto.UpdateStoreRequest
	3,  // 3: vou.proto.CounterpartService.DeleteStore:input_type -> vou.proto.DeleteStoreRequest
	4,  // 4: vou.proto.CounterpartService.CreateBranch:input_type -> vou.proto.CreateBranchRequest
	5,  // 5: vou.proto.CounterpartService.GetBranchs:input_type -> vou.proto.GetBranchsRequest
	6,  // 6: vou.proto.CounterpartService.DeleteBranch:input_type -> vou.proto.DeleteBranchRequest
	7,  // 7: vou.proto.CounterpartService.GetCmsOverview:input_type -> vou.proto.GetCmsOverviewRequest
	8,  // 8: vou.proto.CounterpartService.FakeCmsOverview:input_type -> vou.proto.FakeCmsOverviewRequest
	9,  // 9: vou.proto.CounterpartService.CreateStore:output_type -> vou.proto.CreateStoreResponse
	10, // 10: vou.proto.CounterpartService.GetAllStoresOfOwner:output_type -> vou.proto.GetStoresOfOwnerResponse
	11, // 11: vou.proto.CounterpartService.UpdateStore:output_type -> vou.proto.UpdateStoreResponse
	12, // 12: vou.proto.CounterpartService.DeleteStore:output_type -> vou.proto.DeleteStoreResponse
	13, // 13: vou.proto.CounterpartService.CreateBranch:output_type -> vou.proto.CreateBranchResponse
	14, // 14: vou.proto.CounterpartService.GetBranchs:output_type -> vou.proto.GetBranchsResponse
	15, // 15: vou.proto.CounterpartService.DeleteBranch:output_type -> vou.proto.DeleteBranchResponse
	16, // 16: vou.proto.CounterpartService.GetCmsOverview:output_type -> vou.proto.GetCmsOverviewResponse
	17, // 17: vou.proto.CounterpartService.FakeCmsOverview:output_type -> vou.proto.FakeCmsOverviewResponse
	9,  // [9:18] is the sub-list for method output_type
	0,  // [0:9] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_counterpart_service_counterpart_proto_init() }
func file_counterpart_service_counterpart_proto_init() {
	if File_counterpart_service_counterpart_proto != nil {
		return
	}
	file_counterpart_rpc_create_store_proto_init()
	file_counterpart_rpc_get_stores_of_owner_proto_init()
	file_counterpart_rpc_update_store_proto_init()
	file_counterpart_rpc_delete_store_proto_init()
	file_counterpart_rpc_create_branch_proto_init()
	file_counterpart_rpc_get_branchs_proto_init()
	file_counterpart_rpc_delete_branch_proto_init()
	file_counterpart_rpc_cms_proto_init()
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
