// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.15.8
// source: auth/service_auth.proto

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

var File_auth_service_auth_proto protoreflect.FileDescriptor

var file_auth_service_auth_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x76, 0x6f, 0x75, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x1a, 0x1a, 0x61, 0x75, 0x74, 0x68, 0x2f,
	0x72, 0x70, 0x63, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x72, 0x70, 0x63, 0x5f,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1d, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x21, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x72, 0x65, 0x6e, 0x65, 0x77, 0x5f,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x32, 0xf9, 0x02, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x55, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x21, 0x2e, 0x76, 0x6f, 0x75, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x76, 0x6f, 0x75, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x09, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x20, 0x2e, 0x76, 0x6f, 0x75, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x76, 0x6f, 0x75, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x56, 0x0a,
	0x0d, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x20,
	0x2e, 0x76, 0x6f, 0x75, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x21, 0x2e, 0x76, 0x6f, 0x75, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x67, 0x0a, 0x10, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x27, 0x2e, 0x76, 0x6f, 0x75, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x6e, 0x65, 0x77,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x28, 0x2e, 0x76, 0x6f, 0x75, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x16,
	0x5a, 0x14, 0x56, 0x4f, 0x55, 0x2d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_auth_service_auth_proto_goTypes = []any{
	(*CreateUserRequest)(nil),        // 0: vou.proto.auth.CreateUserRequest
	(*LoginUserRequest)(nil),         // 1: vou.proto.auth.LoginUserRequest
	(*AuthorizeRequest)(nil),         // 2: vou.proto.auth.AuthorizeRequest
	(*RenewAccessTokenRequest)(nil),  // 3: vou.proto.auth.RenewAccessTokenRequest
	(*CreateUserResponse)(nil),       // 4: vou.proto.auth.CreateUserResponse
	(*LoginUserResponse)(nil),        // 5: vou.proto.auth.LoginUserResponse
	(*AuthorizeResponse)(nil),        // 6: vou.proto.auth.AuthorizeResponse
	(*RenewAccessTokenResponse)(nil), // 7: vou.proto.auth.RenewAccessTokenResponse
}
var file_auth_service_auth_proto_depIdxs = []int32{
	0, // 0: vou.proto.auth.AuthService.CreateUser:input_type -> vou.proto.auth.CreateUserRequest
	1, // 1: vou.proto.auth.AuthService.LoginUser:input_type -> vou.proto.auth.LoginUserRequest
	2, // 2: vou.proto.auth.AuthService.AuthorizeUser:input_type -> vou.proto.auth.AuthorizeRequest
	3, // 3: vou.proto.auth.AuthService.RenewAccessToken:input_type -> vou.proto.auth.RenewAccessTokenRequest
	4, // 4: vou.proto.auth.AuthService.CreateUser:output_type -> vou.proto.auth.CreateUserResponse
	5, // 5: vou.proto.auth.AuthService.LoginUser:output_type -> vou.proto.auth.LoginUserResponse
	6, // 6: vou.proto.auth.AuthService.AuthorizeUser:output_type -> vou.proto.auth.AuthorizeResponse
	7, // 7: vou.proto.auth.AuthService.RenewAccessToken:output_type -> vou.proto.auth.RenewAccessTokenResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_auth_service_auth_proto_init() }
func file_auth_service_auth_proto_init() {
	if File_auth_service_auth_proto != nil {
		return
	}
	file_auth_rpc_create_user_proto_init()
	file_auth_rpc_login_user_proto_init()
	file_auth_rpc_authorize_user_proto_init()
	file_auth_rpc_renew_access_token_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_auth_service_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auth_service_auth_proto_goTypes,
		DependencyIndexes: file_auth_service_auth_proto_depIdxs,
	}.Build()
	File_auth_service_auth_proto = out.File
	file_auth_service_auth_proto_rawDesc = nil
	file_auth_service_auth_proto_goTypes = nil
	file_auth_service_auth_proto_depIdxs = nil
}
