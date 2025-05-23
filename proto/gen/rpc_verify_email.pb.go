// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.3
// source: auth/rpc_verify_email.proto

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

type VerifyEmailRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	SecretCode    string                 `protobuf:"bytes,2,opt,name=secret_code,json=secretCode,proto3" json:"secret_code,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VerifyEmailRequest) Reset() {
	*x = VerifyEmailRequest{}
	mi := &file_auth_rpc_verify_email_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VerifyEmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyEmailRequest) ProtoMessage() {}

func (x *VerifyEmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_rpc_verify_email_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyEmailRequest.ProtoReflect.Descriptor instead.
func (*VerifyEmailRequest) Descriptor() ([]byte, []int) {
	return file_auth_rpc_verify_email_proto_rawDescGZIP(), []int{0}
}

func (x *VerifyEmailRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *VerifyEmailRequest) GetSecretCode() string {
	if x != nil {
		return x.SecretCode
	}
	return ""
}

type VerifyEmailResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	IsVerified    bool                   `protobuf:"varint,1,opt,name=is_verified,json=isVerified,proto3" json:"is_verified,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VerifyEmailResponse) Reset() {
	*x = VerifyEmailResponse{}
	mi := &file_auth_rpc_verify_email_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VerifyEmailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyEmailResponse) ProtoMessage() {}

func (x *VerifyEmailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_rpc_verify_email_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyEmailResponse.ProtoReflect.Descriptor instead.
func (*VerifyEmailResponse) Descriptor() ([]byte, []int) {
	return file_auth_rpc_verify_email_proto_rawDescGZIP(), []int{1}
}

func (x *VerifyEmailResponse) GetIsVerified() bool {
	if x != nil {
		return x.IsVerified
	}
	return false
}

var File_auth_rpc_verify_email_proto protoreflect.FileDescriptor

var file_auth_rpc_verify_email_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x76,
	0x6f, 0x75, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4b, 0x0a, 0x12, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x36, 0x0a, 0x13, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x69, 0x73, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0a, 0x69, 0x73, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x42, 0x16, 0x5a,
	0x14, 0x56, 0x4f, 0x55, 0x2d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_auth_rpc_verify_email_proto_rawDescOnce sync.Once
	file_auth_rpc_verify_email_proto_rawDescData = file_auth_rpc_verify_email_proto_rawDesc
)

func file_auth_rpc_verify_email_proto_rawDescGZIP() []byte {
	file_auth_rpc_verify_email_proto_rawDescOnce.Do(func() {
		file_auth_rpc_verify_email_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_rpc_verify_email_proto_rawDescData)
	})
	return file_auth_rpc_verify_email_proto_rawDescData
}

var file_auth_rpc_verify_email_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_auth_rpc_verify_email_proto_goTypes = []any{
	(*VerifyEmailRequest)(nil),  // 0: vou.proto.VerifyEmailRequest
	(*VerifyEmailResponse)(nil), // 1: vou.proto.VerifyEmailResponse
}
var file_auth_rpc_verify_email_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_auth_rpc_verify_email_proto_init() }
func file_auth_rpc_verify_email_proto_init() {
	if File_auth_rpc_verify_email_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_auth_rpc_verify_email_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_auth_rpc_verify_email_proto_goTypes,
		DependencyIndexes: file_auth_rpc_verify_email_proto_depIdxs,
		MessageInfos:      file_auth_rpc_verify_email_proto_msgTypes,
	}.Build()
	File_auth_rpc_verify_email_proto = out.File
	file_auth_rpc_verify_email_proto_rawDesc = nil
	file_auth_rpc_verify_email_proto_goTypes = nil
	file_auth_rpc_verify_email_proto_depIdxs = nil
}
