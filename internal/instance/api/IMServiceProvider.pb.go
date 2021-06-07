//example

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: input.proto

package dubbo3

import (
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

type PublishMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromID string `protobuf:"bytes,1,opt,name=fromID,proto3" json:"fromID,omitempty"`
	ToID   string `protobuf:"bytes,2,opt,name=toID,proto3" json:"toID,omitempty"`
	Data   string `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *PublishMessageRequest) Reset() {
	*x = PublishMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_input_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishMessageRequest) ProtoMessage() {}

func (x *PublishMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_input_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishMessageRequest.ProtoReflect.Descriptor instead.
func (*PublishMessageRequest) Descriptor() ([]byte, []int) {
	return file_input_proto_rawDescGZIP(), []int{0}
}

func (x *PublishMessageRequest) GetFromID() string {
	if x != nil {
		return x.FromID
	}
	return ""
}

func (x *PublishMessageRequest) GetToID() string {
	if x != nil {
		return x.ToID
	}
	return ""
}

func (x *PublishMessageRequest) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type PublishMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    uint32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PublishMessageResponse) Reset() {
	*x = PublishMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_input_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishMessageResponse) ProtoMessage() {}

func (x *PublishMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_input_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishMessageResponse.ProtoReflect.Descriptor instead.
func (*PublishMessageResponse) Descriptor() ([]byte, []int) {
	return file_input_proto_rawDescGZIP(), []int{1}
}

func (x *PublishMessageResponse) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *PublishMessageResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_input_proto protoreflect.FileDescriptor

var file_input_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x22, 0x57, 0x0a, 0x15, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x66, 0x72, 0x6f, 0x6d, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x66, 0x72, 0x6f, 0x6d, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x6f, 0x49, 0x44,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x6f, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0x46, 0x0a, 0x16, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x6a, 0x0a, 0x11, 0x49, 0x4d, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x55, 0x0a,
	0x0e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x11, 0x5a, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x64, 0x75, 0x62, 0x62, 0x6f, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_input_proto_rawDescOnce sync.Once
	file_input_proto_rawDescData = file_input_proto_rawDesc
)

func file_input_proto_rawDescGZIP() []byte {
	file_input_proto_rawDescOnce.Do(func() {
		file_input_proto_rawDescData = protoimpl.X.CompressGZIP(file_input_proto_rawDescData)
	})
	return file_input_proto_rawDescData
}

var file_input_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_input_proto_goTypes = []interface{}{
	(*PublishMessageRequest)(nil),  // 0: protobuf.PublishMessageRequest
	(*PublishMessageResponse)(nil), // 1: protobuf.PublishMessageResponse
}
var file_input_proto_depIdxs = []int32{
	0, // 0: protobuf.IMServiceProvider.PublishMessage:input_type -> protobuf.PublishMessageRequest
	1, // 1: protobuf.IMServiceProvider.PublishMessage:output_type -> protobuf.PublishMessageResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_input_proto_init() }
func file_input_proto_init() {
	if File_input_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_input_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishMessageRequest); i {
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
		file_input_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishMessageResponse); i {
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
			RawDescriptor: file_input_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_input_proto_goTypes,
		DependencyIndexes: file_input_proto_depIdxs,
		MessageInfos:      file_input_proto_msgTypes,
	}.Build()
	File_input_proto = out.File
	file_input_proto_rawDesc = nil
	file_input_proto_goTypes = nil
	file_input_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// IMServiceProviderClient is the client API for IMServiceProvider service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IMServiceProviderClient interface {
	PublishMessage(ctx context.Context, in *PublishMessageRequest, opts ...grpc.CallOption) (*PublishMessageResponse, error)
}

type iMServiceProviderClient struct {
	cc grpc.ClientConnInterface
}

func NewIMServiceProviderClient(cc grpc.ClientConnInterface) IMServiceProviderClient {
	return &iMServiceProviderClient{cc}
}

func (c *iMServiceProviderClient) PublishMessage(ctx context.Context, in *PublishMessageRequest, opts ...grpc.CallOption) (*PublishMessageResponse, error) {
	out := new(PublishMessageResponse)
	err := c.cc.Invoke(ctx, "/protobuf.IMServiceProvider/PublishMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IMServiceProviderServer is the server API for IMServiceProvider service.
type IMServiceProviderServer interface {
	PublishMessage(context.Context, *PublishMessageRequest) (*PublishMessageResponse, error)
}

// UnimplementedIMServiceProviderServer can be embedded to have forward compatible implementations.
type UnimplementedIMServiceProviderServer struct {
}

func (*UnimplementedIMServiceProviderServer) PublishMessage(context.Context, *PublishMessageRequest) (*PublishMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishMessage not implemented")
}

func RegisterIMServiceProviderServer(s *grpc.Server, srv IMServiceProviderServer) {
	s.RegisterService(&_IMServiceProvider_serviceDesc, srv)
}

func _IMServiceProvider_PublishMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IMServiceProviderServer).PublishMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.IMServiceProvider/PublishMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IMServiceProviderServer).PublishMessage(ctx, req.(*PublishMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _IMServiceProvider_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.IMServiceProvider",
	HandlerType: (*IMServiceProviderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PublishMessage",
			Handler:    _IMServiceProvider_PublishMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "input.proto",
}
