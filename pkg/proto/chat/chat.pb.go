// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.13.0
// source: chat/chat.proto

package pbChat

import (
	sdk_ws "open-im/pkg/proto/sdk_ws"
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type MsgDataToMQ struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token       string          `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	OperationID string          `protobuf:"bytes,2,opt,name=operationID,proto3" json:"operationID,omitempty"`
	MsgData     *sdk_ws.MsgData `protobuf:"bytes,3,opt,name=msgData,proto3" json:"msgData,omitempty"`
}

func (x *MsgDataToMQ) Reset() {
	*x = MsgDataToMQ{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgDataToMQ) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgDataToMQ) ProtoMessage() {}

func (x *MsgDataToMQ) ProtoReflect() protoreflect.Message {
	mi := &file_chat_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgDataToMQ.ProtoReflect.Descriptor instead.
func (*MsgDataToMQ) Descriptor() ([]byte, []int) {
	return file_chat_chat_proto_rawDescGZIP(), []int{0}
}

func (x *MsgDataToMQ) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *MsgDataToMQ) GetOperationID() string {
	if x != nil {
		return x.OperationID
	}
	return ""
}

func (x *MsgDataToMQ) GetMsgData() *sdk_ws.MsgData {
	if x != nil {
		return x.MsgData
	}
	return nil
}

type SendMsgReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token       string          `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	OperationID string          `protobuf:"bytes,2,opt,name=operationID,proto3" json:"operationID,omitempty"`
	MsgData     *sdk_ws.MsgData `protobuf:"bytes,3,opt,name=msgData,proto3" json:"msgData,omitempty"`
}

func (x *SendMsgReq) Reset() {
	*x = SendMsgReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMsgReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMsgReq) ProtoMessage() {}

func (x *SendMsgReq) ProtoReflect() protoreflect.Message {
	mi := &file_chat_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMsgReq.ProtoReflect.Descriptor instead.
func (*SendMsgReq) Descriptor() ([]byte, []int) {
	return file_chat_chat_proto_rawDescGZIP(), []int{1}
}

func (x *SendMsgReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *SendMsgReq) GetOperationID() string {
	if x != nil {
		return x.OperationID
	}
	return ""
}

func (x *SendMsgReq) GetMsgData() *sdk_ws.MsgData {
	if x != nil {
		return x.MsgData
	}
	return nil
}

type SendMsgResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrCode     int32  `protobuf:"varint,1,opt,name=errCode,proto3" json:"errCode,omitempty"`
	ErrMsg      string `protobuf:"bytes,2,opt,name=errMsg,proto3" json:"errMsg,omitempty"`
	ServerMsgID string `protobuf:"bytes,4,opt,name=serverMsgID,proto3" json:"serverMsgID,omitempty"`
	ClientMsgID string `protobuf:"bytes,5,opt,name=clientMsgID,proto3" json:"clientMsgID,omitempty"`
	SendTime    int64  `protobuf:"varint,6,opt,name=sendTime,proto3" json:"sendTime,omitempty"`
}

func (x *SendMsgResp) Reset() {
	*x = SendMsgResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_chat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMsgResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMsgResp) ProtoMessage() {}

func (x *SendMsgResp) ProtoReflect() protoreflect.Message {
	mi := &file_chat_chat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMsgResp.ProtoReflect.Descriptor instead.
func (*SendMsgResp) Descriptor() ([]byte, []int) {
	return file_chat_chat_proto_rawDescGZIP(), []int{2}
}

func (x *SendMsgResp) GetErrCode() int32 {
	if x != nil {
		return x.ErrCode
	}
	return 0
}

func (x *SendMsgResp) GetErrMsg() string {
	if x != nil {
		return x.ErrMsg
	}
	return ""
}

func (x *SendMsgResp) GetServerMsgID() string {
	if x != nil {
		return x.ServerMsgID
	}
	return ""
}

func (x *SendMsgResp) GetClientMsgID() string {
	if x != nil {
		return x.ClientMsgID
	}
	return ""
}

func (x *SendMsgResp) GetSendTime() int64 {
	if x != nil {
		return x.SendTime
	}
	return 0
}

var File_chat_chat_proto protoreflect.FileDescriptor

var file_chat_chat_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x06, 0x70, 0x62, 0x43, 0x68, 0x61, 0x74, 0x1a, 0x0f, 0x73, 0x64, 0x6b, 0x5f, 0x77,
	0x73, 0x2f, 0x77, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7b, 0x0a, 0x0b, 0x4d, 0x73,
	0x67, 0x44, 0x61, 0x74, 0x61, 0x54, 0x6f, 0x4d, 0x51, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x20, 0x0a, 0x0b, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x44, 0x12, 0x34, 0x0a, 0x07, 0x6d, 0x73, 0x67, 0x44, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x5f,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x2e, 0x4d, 0x73, 0x67, 0x44, 0x61, 0x74, 0x61, 0x52, 0x07,
	0x6d, 0x73, 0x67, 0x44, 0x61, 0x74, 0x61, 0x22, 0x7a, 0x0a, 0x0a, 0x53, 0x65, 0x6e, 0x64, 0x4d,
	0x73, 0x67, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x34, 0x0a,
	0x07, 0x6d, 0x73, 0x67, 0x44, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x2e, 0x4d, 0x73, 0x67, 0x44, 0x61, 0x74, 0x61, 0x52, 0x07, 0x6d, 0x73, 0x67, 0x44,
	0x61, 0x74, 0x61, 0x22, 0x9f, 0x01, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x73, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x65, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x65, 0x72, 0x72, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65,
	0x72, 0x72, 0x4d, 0x73, 0x67, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4d,
	0x73, 0x67, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x4d, 0x73, 0x67, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x4d, 0x73, 0x67, 0x49, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x4d, 0x73, 0x67, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6e,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x65, 0x6e,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x32, 0x3a, 0x0a, 0x04, 0x43, 0x68, 0x61, 0x74, 0x12, 0x32, 0x0a,
	0x07, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x73, 0x67, 0x12, 0x12, 0x2e, 0x70, 0x62, 0x43, 0x68, 0x61,
	0x74, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x70,
	0x62, 0x43, 0x68, 0x61, 0x74, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x3b, 0x70, 0x62, 0x43, 0x68,
	0x61, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chat_chat_proto_rawDescOnce sync.Once
	file_chat_chat_proto_rawDescData = file_chat_chat_proto_rawDesc
)

func file_chat_chat_proto_rawDescGZIP() []byte {
	file_chat_chat_proto_rawDescOnce.Do(func() {
		file_chat_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_chat_chat_proto_rawDescData)
	})
	return file_chat_chat_proto_rawDescData
}

var file_chat_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_chat_chat_proto_goTypes = []interface{}{
	(*MsgDataToMQ)(nil),    // 0: pbChat.MsgDataToMQ
	(*SendMsgReq)(nil),     // 1: pbChat.SendMsgReq
	(*SendMsgResp)(nil),    // 2: pbChat.SendMsgResp
	(*sdk_ws.MsgData)(nil), // 3: server_api_params.MsgData
}
var file_chat_chat_proto_depIdxs = []int32{
	3, // 0: pbChat.MsgDataToMQ.msgData:type_name -> server_api_params.MsgData
	3, // 1: pbChat.SendMsgReq.msgData:type_name -> server_api_params.MsgData
	1, // 2: pbChat.Chat.SendMsg:input_type -> pbChat.SendMsgReq
	2, // 3: pbChat.Chat.SendMsg:output_type -> pbChat.SendMsgResp
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_chat_chat_proto_init() }
func file_chat_chat_proto_init() {
	if File_chat_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chat_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgDataToMQ); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chat_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMsgReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chat_chat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMsgResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_chat_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chat_chat_proto_goTypes,
		DependencyIndexes: file_chat_chat_proto_depIdxs,
		MessageInfos:      file_chat_chat_proto_msgTypes,
	}.Build()
	File_chat_chat_proto = out.File
	file_chat_chat_proto_rawDesc = nil
	file_chat_chat_proto_goTypes = nil
	file_chat_chat_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ChatClient is the client API for Chat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChatClient interface {
	SendMsg(ctx context.Context, in *SendMsgReq, opts ...grpc.CallOption) (*SendMsgResp, error)
}

type chatClient struct {
	cc grpc.ClientConnInterface
}

func NewChatClient(cc grpc.ClientConnInterface) ChatClient {
	return &chatClient{cc}
}

func (c *chatClient) SendMsg(ctx context.Context, in *SendMsgReq, opts ...grpc.CallOption) (*SendMsgResp, error) {
	out := new(SendMsgResp)
	err := c.cc.Invoke(ctx, "/pbChat.Chat/SendMsg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServer is the server API for Chat service.
type ChatServer interface {
	SendMsg(context.Context, *SendMsgReq) (*SendMsgResp, error)
}

// UnimplementedChatServer can be embedded to have forward compatible implementations.
type UnimplementedChatServer struct {
}

func (*UnimplementedChatServer) SendMsg(context.Context, *SendMsgReq) (*SendMsgResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMsg not implemented")
}

func RegisterChatServer(s *grpc.Server, srv ChatServer) {
	s.RegisterService(&_Chat_serviceDesc, srv)
}

func _Chat_SendMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMsgReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).SendMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbChat.Chat/SendMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).SendMsg(ctx, req.(*SendMsgReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Chat_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pbChat.Chat",
	HandlerType: (*ChatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMsg",
			Handler:    _Chat_SendMsg_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chat/chat.proto",
}
