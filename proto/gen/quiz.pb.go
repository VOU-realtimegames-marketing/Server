// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.15.8
// source: event-quiz/quiz.proto

package gen

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Quiz struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	EventId   int64                  `protobuf:"varint,2,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	QuizGenre string                 `protobuf:"bytes,3,opt,name=quiz_genre,json=quizGenre,proto3" json:"quiz_genre,omitempty"`
	Content   *structpb.Struct       `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Quiz) Reset() {
	*x = Quiz{}
	mi := &file_event_quiz_quiz_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Quiz) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Quiz) ProtoMessage() {}

func (x *Quiz) ProtoReflect() protoreflect.Message {
	mi := &file_event_quiz_quiz_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Quiz.ProtoReflect.Descriptor instead.
func (*Quiz) Descriptor() ([]byte, []int) {
	return file_event_quiz_quiz_proto_rawDescGZIP(), []int{0}
}

func (x *Quiz) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Quiz) GetEventId() int64 {
	if x != nil {
		return x.EventId
	}
	return 0
}

func (x *Quiz) GetQuizGenre() string {
	if x != nil {
		return x.QuizGenre
	}
	return ""
}

func (x *Quiz) GetContent() *structpb.Struct {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *Quiz) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

var File_event_quiz_quiz_proto protoreflect.FileDescriptor

var file_event_quiz_quiz_proto_rawDesc = []byte{
	0x0a, 0x15, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2d, 0x71, 0x75, 0x69, 0x7a, 0x2f, 0x71, 0x75, 0x69,
	0x7a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x76, 0x6f, 0x75, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xbe, 0x01, 0x0a, 0x04, 0x51, 0x75, 0x69, 0x7a, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x71, 0x75, 0x69, 0x7a, 0x5f, 0x67, 0x65,
	0x6e, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x71, 0x75, 0x69, 0x7a, 0x47,
	0x65, 0x6e, 0x72, 0x65, 0x12, 0x31, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x42, 0x16, 0x5a, 0x14, 0x56, 0x4f, 0x55, 0x2d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_event_quiz_quiz_proto_rawDescOnce sync.Once
	file_event_quiz_quiz_proto_rawDescData = file_event_quiz_quiz_proto_rawDesc
)

func file_event_quiz_quiz_proto_rawDescGZIP() []byte {
	file_event_quiz_quiz_proto_rawDescOnce.Do(func() {
		file_event_quiz_quiz_proto_rawDescData = protoimpl.X.CompressGZIP(file_event_quiz_quiz_proto_rawDescData)
	})
	return file_event_quiz_quiz_proto_rawDescData
}

var file_event_quiz_quiz_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_event_quiz_quiz_proto_goTypes = []any{
	(*Quiz)(nil),                  // 0: vou.proto.Quiz
	(*structpb.Struct)(nil),       // 1: google.protobuf.Struct
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_event_quiz_quiz_proto_depIdxs = []int32{
	1, // 0: vou.proto.Quiz.content:type_name -> google.protobuf.Struct
	2, // 1: vou.proto.Quiz.created_at:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_event_quiz_quiz_proto_init() }
func file_event_quiz_quiz_proto_init() {
	if File_event_quiz_quiz_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_event_quiz_quiz_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_event_quiz_quiz_proto_goTypes,
		DependencyIndexes: file_event_quiz_quiz_proto_depIdxs,
		MessageInfos:      file_event_quiz_quiz_proto_msgTypes,
	}.Build()
	File_event_quiz_quiz_proto = out.File
	file_event_quiz_quiz_proto_rawDesc = nil
	file_event_quiz_quiz_proto_goTypes = nil
	file_event_quiz_quiz_proto_depIdxs = nil
}
