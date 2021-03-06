// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: ids.proto

package ids

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

type Code int32

const (
	Code_SUCC  Code = 0
	Code_FIELD Code = 1
)

// Enum value maps for Code.
var (
	Code_name = map[int32]string{
		0: "SUCC",
		1: "FIELD",
	}
	Code_value = map[string]int32{
		"SUCC":  0,
		"FIELD": 1,
	}
)

func (x Code) Enum() *Code {
	p := new(Code)
	*p = x
	return p
}

func (x Code) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Code) Descriptor() protoreflect.EnumDescriptor {
	return file_ids_proto_enumTypes[0].Descriptor()
}

func (Code) Type() protoreflect.EnumType {
	return &file_ids_proto_enumTypes[0]
}

func (x Code) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Code.Descriptor instead.
func (Code) EnumDescriptor() ([]byte, []int) {
	return file_ids_proto_rawDescGZIP(), []int{0}
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServID *int32 `protobuf:"varint,1,opt,name=servID,proto3,oneof" json:"servID,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ids_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_ids_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_ids_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetServID() int32 {
	if x != nil && x.ServID != nil {
		return *x.ServID
	}
	return 0
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServID int32 `protobuf:"varint,1,opt,name=servID,proto3" json:"servID,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ids_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_ids_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_ids_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetServID() int32 {
	if x != nil {
		return x.ServID
	}
	return 0
}

type ResponseID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID   uint64 `protobuf:"varint,1,opt,name=iD,proto3" json:"iD,omitempty"`
	Code Code   `protobuf:"varint,2,opt,name=code,proto3,enum=Code" json:"code,omitempty"`
}

func (x *ResponseID) Reset() {
	*x = ResponseID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ids_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseID) ProtoMessage() {}

func (x *ResponseID) ProtoReflect() protoreflect.Message {
	mi := &file_ids_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseID.ProtoReflect.Descriptor instead.
func (*ResponseID) Descriptor() ([]byte, []int) {
	return file_ids_proto_rawDescGZIP(), []int{2}
}

func (x *ResponseID) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *ResponseID) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_SUCC
}

var File_ids_proto protoreflect.FileDescriptor

var file_ids_proto_rawDesc = []byte{
	0x0a, 0x09, 0x69, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x31, 0x0a, 0x07, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x49, 0x44,
	0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x49, 0x44, 0x22, 0x22,
	0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65,
	0x72, 0x76, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76,
	0x49, 0x44, 0x22, 0x37, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x44,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x44,
	0x12, 0x19, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x05,
	0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x2a, 0x1b, 0x0a, 0x04, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x55, 0x43, 0x43, 0x10, 0x00, 0x12, 0x09, 0x0a,
	0x05, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x10, 0x01, 0x32, 0x54, 0x0a, 0x0c, 0x49, 0x44, 0x73, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x05, 0x47, 0x65, 0x74, 0x49,
	0x44, 0x12, 0x08, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x44, 0x22, 0x00, 0x12, 0x22, 0x0a, 0x09, 0x47, 0x65,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x49, 0x44, 0x12, 0x08, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x09, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x08,
	0x5a, 0x06, 0x2e, 0x2f, 0x3b, 0x69, 0x64, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ids_proto_rawDescOnce sync.Once
	file_ids_proto_rawDescData = file_ids_proto_rawDesc
)

func file_ids_proto_rawDescGZIP() []byte {
	file_ids_proto_rawDescOnce.Do(func() {
		file_ids_proto_rawDescData = protoimpl.X.CompressGZIP(file_ids_proto_rawDescData)
	})
	return file_ids_proto_rawDescData
}

var file_ids_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ids_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_ids_proto_goTypes = []interface{}{
	(Code)(0),          // 0: Code
	(*Request)(nil),    // 1: Request
	(*Response)(nil),   // 2: Response
	(*ResponseID)(nil), // 3: ResponseID
}
var file_ids_proto_depIdxs = []int32{
	0, // 0: ResponseID.code:type_name -> Code
	1, // 1: IDsInterface.GetID:input_type -> Request
	1, // 2: IDsInterface.GetServID:input_type -> Request
	3, // 3: IDsInterface.GetID:output_type -> ResponseID
	2, // 4: IDsInterface.GetServID:output_type -> Response
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ids_proto_init() }
func file_ids_proto_init() {
	if File_ids_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ids_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_ids_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_ids_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseID); i {
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
	file_ids_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ids_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ids_proto_goTypes,
		DependencyIndexes: file_ids_proto_depIdxs,
		EnumInfos:         file_ids_proto_enumTypes,
		MessageInfos:      file_ids_proto_msgTypes,
	}.Build()
	File_ids_proto = out.File
	file_ids_proto_rawDesc = nil
	file_ids_proto_goTypes = nil
	file_ids_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// IDsInterfaceClient is the client API for IDsInterface service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IDsInterfaceClient interface {
	GetID(ctx context.Context, in *Request, opts ...grpc.CallOption) (*ResponseID, error)
	GetServID(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type iDsInterfaceClient struct {
	cc grpc.ClientConnInterface
}

func NewIDsInterfaceClient(cc grpc.ClientConnInterface) IDsInterfaceClient {
	return &iDsInterfaceClient{cc}
}

func (c *iDsInterfaceClient) GetID(ctx context.Context, in *Request, opts ...grpc.CallOption) (*ResponseID, error) {
	out := new(ResponseID)
	err := c.cc.Invoke(ctx, "/IDsInterface/GetID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iDsInterfaceClient) GetServID(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/IDsInterface/GetServID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IDsInterfaceServer is the server API for IDsInterface service.
type IDsInterfaceServer interface {
	GetID(context.Context, *Request) (*ResponseID, error)
	GetServID(context.Context, *Request) (*Response, error)
}

// UnimplementedIDsInterfaceServer can be embedded to have forward compatible implementations.
type UnimplementedIDsInterfaceServer struct {
}

func (*UnimplementedIDsInterfaceServer) GetID(context.Context, *Request) (*ResponseID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetID not implemented")
}
func (*UnimplementedIDsInterfaceServer) GetServID(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServID not implemented")
}

func RegisterIDsInterfaceServer(s *grpc.Server, srv IDsInterfaceServer) {
	s.RegisterService(&_IDsInterface_serviceDesc, srv)
}

func _IDsInterface_GetID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IDsInterfaceServer).GetID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/IDsInterface/GetID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IDsInterfaceServer).GetID(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _IDsInterface_GetServID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IDsInterfaceServer).GetServID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/IDsInterface/GetServID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IDsInterfaceServer).GetServID(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _IDsInterface_serviceDesc = grpc.ServiceDesc{
	ServiceName: "IDsInterface",
	HandlerType: (*IDsInterfaceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetID",
			Handler:    _IDsInterface_GetID_Handler,
		},
		{
			MethodName: "GetServID",
			Handler:    _IDsInterface_GetServID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ids.proto",
}
