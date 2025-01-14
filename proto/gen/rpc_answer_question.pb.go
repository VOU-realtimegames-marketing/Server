// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.15.8
// source: game/rpc_answer_question.proto

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

type AnswerQuestionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QuestionNum int32  `protobuf:"varint,1,opt,name=question_num,json=questionNum,proto3" json:"question_num,omitempty"`
	EventId     int64  `protobuf:"varint,2,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	Username    string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Answer      string `protobuf:"bytes,4,opt,name=answer,proto3" json:"answer,omitempty"`
}

func (x *AnswerQuestionRequest) Reset() {
	*x = AnswerQuestionRequest{}
	mi := &file_game_rpc_answer_question_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AnswerQuestionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnswerQuestionRequest) ProtoMessage() {}

func (x *AnswerQuestionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_game_rpc_answer_question_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnswerQuestionRequest.ProtoReflect.Descriptor instead.
func (*AnswerQuestionRequest) Descriptor() ([]byte, []int) {
	return file_game_rpc_answer_question_proto_rawDescGZIP(), []int{0}
}

func (x *AnswerQuestionRequest) GetQuestionNum() int32 {
	if x != nil {
		return x.QuestionNum
	}
	return 0
}

func (x *AnswerQuestionRequest) GetEventId() int64 {
	if x != nil {
		return x.EventId
	}
	return 0
}

func (x *AnswerQuestionRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *AnswerQuestionRequest) GetAnswer() string {
	if x != nil {
		return x.Answer
	}
	return ""
}

type AnswerQuestionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsCorrect     bool   `protobuf:"varint,1,opt,name=is_correct,json=isCorrect,proto3" json:"is_correct,omitempty"`
	CorrectAnswer string `protobuf:"bytes,2,opt,name=correct_answer,json=correctAnswer,proto3" json:"correct_answer,omitempty"`
}

func (x *AnswerQuestionResponse) Reset() {
	*x = AnswerQuestionResponse{}
	mi := &file_game_rpc_answer_question_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AnswerQuestionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnswerQuestionResponse) ProtoMessage() {}

func (x *AnswerQuestionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_game_rpc_answer_question_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnswerQuestionResponse.ProtoReflect.Descriptor instead.
func (*AnswerQuestionResponse) Descriptor() ([]byte, []int) {
	return file_game_rpc_answer_question_proto_rawDescGZIP(), []int{1}
}

func (x *AnswerQuestionResponse) GetIsCorrect() bool {
	if x != nil {
		return x.IsCorrect
	}
	return false
}

func (x *AnswerQuestionResponse) GetCorrectAnswer() string {
	if x != nil {
		return x.CorrectAnswer
	}
	return ""
}

var File_game_rpc_answer_question_proto protoreflect.FileDescriptor

var file_game_rpc_answer_question_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x61, 0x6e, 0x73, 0x77, 0x65,
	0x72, 0x5f, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x09, 0x76, 0x6f, 0x75, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x89, 0x01, 0x0a, 0x15,
	0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x22, 0x5e, 0x0a, 0x16, 0x41, 0x6e, 0x73, 0x77, 0x65,
	0x72, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x43, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74,
	0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x61, 0x6e, 0x73, 0x77,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x63,
	0x74, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x42, 0x16, 0x5a, 0x14, 0x56, 0x4f, 0x55, 0x2d, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_game_rpc_answer_question_proto_rawDescOnce sync.Once
	file_game_rpc_answer_question_proto_rawDescData = file_game_rpc_answer_question_proto_rawDesc
)

func file_game_rpc_answer_question_proto_rawDescGZIP() []byte {
	file_game_rpc_answer_question_proto_rawDescOnce.Do(func() {
		file_game_rpc_answer_question_proto_rawDescData = protoimpl.X.CompressGZIP(file_game_rpc_answer_question_proto_rawDescData)
	})
	return file_game_rpc_answer_question_proto_rawDescData
}

var file_game_rpc_answer_question_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_game_rpc_answer_question_proto_goTypes = []any{
	(*AnswerQuestionRequest)(nil),  // 0: vou.proto.AnswerQuestionRequest
	(*AnswerQuestionResponse)(nil), // 1: vou.proto.AnswerQuestionResponse
}
var file_game_rpc_answer_question_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_game_rpc_answer_question_proto_init() }
func file_game_rpc_answer_question_proto_init() {
	if File_game_rpc_answer_question_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_game_rpc_answer_question_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_game_rpc_answer_question_proto_goTypes,
		DependencyIndexes: file_game_rpc_answer_question_proto_depIdxs,
		MessageInfos:      file_game_rpc_answer_question_proto_msgTypes,
	}.Build()
	File_game_rpc_answer_question_proto = out.File
	file_game_rpc_answer_question_proto_rawDesc = nil
	file_game_rpc_answer_question_proto_goTypes = nil
	file_game_rpc_answer_question_proto_depIdxs = nil
}
