// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: MsgService.proto

package __

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sender   int32                  `protobuf:"varint,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Reciever string                 `protobuf:"bytes,2,opt,name=reciever,proto3" json:"reciever,omitempty"`
	Text     string                 `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	SendTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=sendTime,proto3" json:"sendTime,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	mi := &file_MsgService_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_MsgService_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_MsgService_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetSender() int32 {
	if x != nil {
		return x.Sender
	}
	return 0
}

func (x *Message) GetReciever() string {
	if x != nil {
		return x.Reciever
	}
	return ""
}

func (x *Message) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Message) GetSendTime() *timestamppb.Timestamp {
	if x != nil {
		return x.SendTime
	}
	return nil
}

type BeautifiedMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sender   int32  `protobuf:"varint,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Reciever int32  `protobuf:"varint,2,opt,name=reciever,proto3" json:"reciever,omitempty"`
	Text     string `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *BeautifiedMessage) Reset() {
	*x = BeautifiedMessage{}
	mi := &file_MsgService_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BeautifiedMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BeautifiedMessage) ProtoMessage() {}

func (x *BeautifiedMessage) ProtoReflect() protoreflect.Message {
	mi := &file_MsgService_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BeautifiedMessage.ProtoReflect.Descriptor instead.
func (*BeautifiedMessage) Descriptor() ([]byte, []int) {
	return file_MsgService_proto_rawDescGZIP(), []int{1}
}

func (x *BeautifiedMessage) GetSender() int32 {
	if x != nil {
		return x.Sender
	}
	return 0
}

func (x *BeautifiedMessage) GetReciever() int32 {
	if x != nil {
		return x.Reciever
	}
	return 0
}

func (x *BeautifiedMessage) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	mi := &file_MsgService_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_MsgService_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_MsgService_proto_rawDescGZIP(), []int{2}
}

func (x *Status) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type Chat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JsonedChat []byte `protobuf:"bytes,1,opt,name=JsonedChat,proto3" json:"JsonedChat,omitempty"`
}

func (x *Chat) Reset() {
	*x = Chat{}
	mi := &file_MsgService_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Chat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chat) ProtoMessage() {}

func (x *Chat) ProtoReflect() protoreflect.Message {
	mi := &file_MsgService_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chat.ProtoReflect.Descriptor instead.
func (*Chat) Descriptor() ([]byte, []int) {
	return file_MsgService_proto_rawDescGZIP(), []int{3}
}

func (x *Chat) GetJsonedChat() []byte {
	if x != nil {
		return x.JsonedChat
	}
	return nil
}

var File_MsgService_proto protoreflect.FileDescriptor

var file_MsgService_proto_rawDesc = []byte{
	0x0a, 0x10, 0x4d, 0x73, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x4d, 0x73, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x89, 0x01, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x63, 0x69, 0x65, 0x76, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x63, 0x69, 0x65, 0x76, 0x65, 0x72, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x65, 0x78, 0x74, 0x12, 0x36, 0x0a, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x5b, 0x0a, 0x11, 0x42,
	0x65, 0x61, 0x75, 0x74, 0x69, 0x66, 0x69, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x63, 0x69,
	0x65, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x72, 0x65, 0x63, 0x69,
	0x65, 0x76, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x20, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x26, 0x0a, 0x04, 0x43, 0x68,
	0x61, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x4a, 0x73, 0x6f, 0x6e, 0x65, 0x64, 0x43, 0x68, 0x61, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x4a, 0x73, 0x6f, 0x6e, 0x65, 0x64, 0x43, 0x68,
	0x61, 0x74, 0x32, 0x40, 0x0a, 0x08, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x34,
	0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x13, 0x2e,
	0x4d, 0x73, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x1a, 0x10, 0x2e, 0x4d, 0x73, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x43, 0x68, 0x61, 0x74, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_MsgService_proto_rawDescOnce sync.Once
	file_MsgService_proto_rawDescData = file_MsgService_proto_rawDesc
)

func file_MsgService_proto_rawDescGZIP() []byte {
	file_MsgService_proto_rawDescOnce.Do(func() {
		file_MsgService_proto_rawDescData = protoimpl.X.CompressGZIP(file_MsgService_proto_rawDescData)
	})
	return file_MsgService_proto_rawDescData
}

var file_MsgService_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_MsgService_proto_goTypes = []any{
	(*Message)(nil),               // 0: MsgService.Message
	(*BeautifiedMessage)(nil),     // 1: MsgService.BeautifiedMessage
	(*Status)(nil),                // 2: MsgService.Status
	(*Chat)(nil),                  // 3: MsgService.Chat
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_MsgService_proto_depIdxs = []int32{
	4, // 0: MsgService.Message.sendTime:type_name -> google.protobuf.Timestamp
	0, // 1: MsgService.Messages.GetMessages:input_type -> MsgService.Message
	3, // 2: MsgService.Messages.GetMessages:output_type -> MsgService.Chat
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_MsgService_proto_init() }
func file_MsgService_proto_init() {
	if File_MsgService_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_MsgService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_MsgService_proto_goTypes,
		DependencyIndexes: file_MsgService_proto_depIdxs,
		MessageInfos:      file_MsgService_proto_msgTypes,
	}.Build()
	File_MsgService_proto = out.File
	file_MsgService_proto_rawDesc = nil
	file_MsgService_proto_goTypes = nil
	file_MsgService_proto_depIdxs = nil
}
