// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mockibc/mockibc/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// QueryParamsRequest is request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f082052722bf8a69, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params holds all the parameters of this module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f082052722bf8a69, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

type QueryGetReceivedPacketRequest struct {
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *QueryGetReceivedPacketRequest) Reset()         { *m = QueryGetReceivedPacketRequest{} }
func (m *QueryGetReceivedPacketRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetReceivedPacketRequest) ProtoMessage()    {}
func (*QueryGetReceivedPacketRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f082052722bf8a69, []int{2}
}
func (m *QueryGetReceivedPacketRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetReceivedPacketRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetReceivedPacketRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetReceivedPacketRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetReceivedPacketRequest.Merge(m, src)
}
func (m *QueryGetReceivedPacketRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetReceivedPacketRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetReceivedPacketRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetReceivedPacketRequest proto.InternalMessageInfo

func (m *QueryGetReceivedPacketRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type QueryGetReceivedPacketResponse struct {
	ReceivedPacket ReceivedPacket `protobuf:"bytes,1,opt,name=ReceivedPacket,proto3" json:"ReceivedPacket"`
}

func (m *QueryGetReceivedPacketResponse) Reset()         { *m = QueryGetReceivedPacketResponse{} }
func (m *QueryGetReceivedPacketResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetReceivedPacketResponse) ProtoMessage()    {}
func (*QueryGetReceivedPacketResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f082052722bf8a69, []int{3}
}
func (m *QueryGetReceivedPacketResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetReceivedPacketResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetReceivedPacketResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetReceivedPacketResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetReceivedPacketResponse.Merge(m, src)
}
func (m *QueryGetReceivedPacketResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetReceivedPacketResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetReceivedPacketResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetReceivedPacketResponse proto.InternalMessageInfo

func (m *QueryGetReceivedPacketResponse) GetReceivedPacket() ReceivedPacket {
	if m != nil {
		return m.ReceivedPacket
	}
	return ReceivedPacket{}
}

type QueryAllReceivedPacketRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllReceivedPacketRequest) Reset()         { *m = QueryAllReceivedPacketRequest{} }
func (m *QueryAllReceivedPacketRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllReceivedPacketRequest) ProtoMessage()    {}
func (*QueryAllReceivedPacketRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f082052722bf8a69, []int{4}
}
func (m *QueryAllReceivedPacketRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllReceivedPacketRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllReceivedPacketRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllReceivedPacketRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllReceivedPacketRequest.Merge(m, src)
}
func (m *QueryAllReceivedPacketRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllReceivedPacketRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllReceivedPacketRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllReceivedPacketRequest proto.InternalMessageInfo

func (m *QueryAllReceivedPacketRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryAllReceivedPacketResponse struct {
	ReceivedPacket []ReceivedPacket    `protobuf:"bytes,1,rep,name=ReceivedPacket,proto3" json:"ReceivedPacket"`
	Pagination     *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllReceivedPacketResponse) Reset()         { *m = QueryAllReceivedPacketResponse{} }
func (m *QueryAllReceivedPacketResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllReceivedPacketResponse) ProtoMessage()    {}
func (*QueryAllReceivedPacketResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f082052722bf8a69, []int{5}
}
func (m *QueryAllReceivedPacketResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllReceivedPacketResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllReceivedPacketResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllReceivedPacketResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllReceivedPacketResponse.Merge(m, src)
}
func (m *QueryAllReceivedPacketResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllReceivedPacketResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllReceivedPacketResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllReceivedPacketResponse proto.InternalMessageInfo

func (m *QueryAllReceivedPacketResponse) GetReceivedPacket() []ReceivedPacket {
	if m != nil {
		return m.ReceivedPacket
	}
	return nil
}

func (m *QueryAllReceivedPacketResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "mockibc.mockibc.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "mockibc.mockibc.QueryParamsResponse")
	proto.RegisterType((*QueryGetReceivedPacketRequest)(nil), "mockibc.mockibc.QueryGetReceivedPacketRequest")
	proto.RegisterType((*QueryGetReceivedPacketResponse)(nil), "mockibc.mockibc.QueryGetReceivedPacketResponse")
	proto.RegisterType((*QueryAllReceivedPacketRequest)(nil), "mockibc.mockibc.QueryAllReceivedPacketRequest")
	proto.RegisterType((*QueryAllReceivedPacketResponse)(nil), "mockibc.mockibc.QueryAllReceivedPacketResponse")
}

func init() { proto.RegisterFile("mockibc/mockibc/query.proto", fileDescriptor_f082052722bf8a69) }

var fileDescriptor_f082052722bf8a69 = []byte{
	// 482 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0xeb, 0x6c, 0xf4, 0x60, 0xa4, 0x21, 0xcc, 0xa4, 0x42, 0x80, 0x74, 0x32, 0x8c, 0x4d,
	0x48, 0xd8, 0xea, 0x10, 0x1f, 0x60, 0x3b, 0xb0, 0x0b, 0x48, 0x25, 0x47, 0x2e, 0xc8, 0x49, 0xac,
	0xc8, 0x5a, 0x1a, 0x67, 0xb1, 0x37, 0x31, 0x21, 0x2e, 0x7c, 0x02, 0x24, 0xee, 0x88, 0x8f, 0xc1,
	0x95, 0xdb, 0x8e, 0x93, 0xb8, 0x70, 0x42, 0xa8, 0xe5, 0x83, 0xa0, 0xda, 0x0e, 0x2c, 0x69, 0xd2,
	0xa2, 0x9d, 0x5c, 0xf9, 0xbd, 0xff, 0xff, 0xfd, 0x9e, 0xdf, 0x6b, 0xe0, 0xdd, 0x89, 0x8c, 0x8f,
	0x44, 0x14, 0xd3, 0xea, 0x3c, 0x3e, 0xe1, 0xe5, 0x19, 0x29, 0x4a, 0xa9, 0x25, 0xba, 0xe1, 0x2e,
	0x89, 0x3b, 0xfd, 0xcd, 0x54, 0xa6, 0xd2, 0xc4, 0xe8, 0xfc, 0x97, 0x4d, 0xf3, 0xef, 0xa5, 0x52,
	0xa6, 0x19, 0xa7, 0xac, 0x10, 0x94, 0xe5, 0xb9, 0xd4, 0x4c, 0x0b, 0x99, 0x2b, 0x17, 0x7d, 0x1c,
	0x4b, 0x35, 0x91, 0x8a, 0x46, 0x4c, 0x71, 0xeb, 0x4e, 0x4f, 0x47, 0x11, 0xd7, 0x6c, 0x44, 0x0b,
	0x96, 0x8a, 0xdc, 0x24, 0x57, 0x4e, 0x4d, 0x9a, 0x82, 0x95, 0x6c, 0x52, 0x39, 0x6d, 0x37, 0xa3,
	0x25, 0x8f, 0xb9, 0x38, 0xe5, 0xc9, 0x9b, 0x82, 0xc5, 0x47, 0x5c, 0xdb, 0x34, 0xbc, 0x09, 0xd1,
	0xab, 0x79, 0x99, 0xb1, 0xd1, 0x86, 0xfc, 0xf8, 0x84, 0x2b, 0x8d, 0x5f, 0xc0, 0x5b, 0xb5, 0x5b,
	0x55, 0xc8, 0x5c, 0x71, 0xf4, 0x0c, 0xf6, 0x6d, 0x8d, 0xdb, 0x60, 0x0b, 0xec, 0x5e, 0xdf, 0x1b,
	0x90, 0x46, 0xcf, 0xc4, 0x0a, 0x0e, 0xd6, 0xcf, 0x7f, 0x0e, 0x7b, 0xa1, 0x4b, 0xc6, 0x14, 0xde,
	0x37, 0x6e, 0x87, 0x5c, 0x87, 0x0e, 0x62, 0x6c, 0x18, 0x5c, 0x39, 0xb4, 0x01, 0x3d, 0x91, 0x18,
	0xcf, 0xf5, 0xd0, 0x13, 0x09, 0x96, 0x30, 0xe8, 0x12, 0x38, 0x92, 0x97, 0x70, 0xa3, 0x1e, 0x71,
	0x44, 0xc3, 0x05, 0xa2, 0x7a, 0x9a, 0x23, 0x6b, 0x88, 0x71, 0xea, 0x08, 0xf7, 0xb3, 0xac, 0x9d,
	0xf0, 0x39, 0x84, 0xff, 0xde, 0xdf, 0xd5, 0x7a, 0x44, 0xec, 0xb0, 0xc8, 0x7c, 0x58, 0xc4, 0xae,
	0x82, 0x1b, 0x16, 0x19, 0xb3, 0x94, 0x3b, 0x6d, 0x78, 0x49, 0x89, 0xbf, 0x02, 0xd7, 0x5a, 0x4b,
	0xa5, 0x25, 0xad, 0xad, 0x5d, 0xb9, 0x35, 0x74, 0x58, 0x23, 0xf7, 0x0c, 0xf9, 0xce, 0x4a, 0x72,
	0xcb, 0x72, 0x19, 0x7d, 0xef, 0xdb, 0x1a, 0xbc, 0x66, 0xd0, 0x91, 0x86, 0x7d, 0x3b, 0x67, 0xf4,
	0x60, 0x81, 0x69, 0x71, 0x99, 0xfc, 0x87, 0xcb, 0x93, 0x6c, 0x29, 0x3c, 0xfc, 0xf0, 0xfd, 0xf7,
	0x27, 0xef, 0x0e, 0x1a, 0xd0, 0xf6, 0xb5, 0x46, 0x5f, 0x40, 0xf3, 0x61, 0x10, 0x69, 0x77, 0xee,
	0xda, 0x33, 0x9f, 0xfe, 0x77, 0xbe, 0x83, 0x7a, 0x62, 0xa0, 0x76, 0xd0, 0x36, 0x5d, 0xf1, 0x6f,
	0xa2, 0xef, 0x44, 0xf2, 0x1e, 0x7d, 0x06, 0xf0, 0x66, 0xdd, 0x69, 0x3f, 0xcb, 0xba, 0x28, 0xbb,
	0x76, 0xad, 0x8b, 0xb2, 0x73, 0x63, 0xf0, 0xae, 0xa1, 0xc4, 0x68, 0x6b, 0x15, 0xe5, 0xc1, 0xe8,
	0x7c, 0x1a, 0x80, 0x8b, 0x69, 0x00, 0x7e, 0x4d, 0x03, 0xf0, 0x71, 0x16, 0xf4, 0x2e, 0x66, 0x41,
	0xef, 0xc7, 0x2c, 0xe8, 0xbd, 0x1e, 0x54, 0x92, 0xb7, 0x7f, 0xc5, 0xfa, 0xac, 0xe0, 0x2a, 0xea,
	0x9b, 0xef, 0xc4, 0xd3, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xcc, 0x13, 0x5e, 0x14, 0xfc, 0x04,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a ReceivedPacket by id.
	ReceivedPacket(ctx context.Context, in *QueryGetReceivedPacketRequest, opts ...grpc.CallOption) (*QueryGetReceivedPacketResponse, error)
	// Queries a list of ReceivedPacket items.
	ReceivedPacketAll(ctx context.Context, in *QueryAllReceivedPacketRequest, opts ...grpc.CallOption) (*QueryAllReceivedPacketResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/mockibc.mockibc.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ReceivedPacket(ctx context.Context, in *QueryGetReceivedPacketRequest, opts ...grpc.CallOption) (*QueryGetReceivedPacketResponse, error) {
	out := new(QueryGetReceivedPacketResponse)
	err := c.cc.Invoke(ctx, "/mockibc.mockibc.Query/ReceivedPacket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ReceivedPacketAll(ctx context.Context, in *QueryAllReceivedPacketRequest, opts ...grpc.CallOption) (*QueryAllReceivedPacketResponse, error) {
	out := new(QueryAllReceivedPacketResponse)
	err := c.cc.Invoke(ctx, "/mockibc.mockibc.Query/ReceivedPacketAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a ReceivedPacket by id.
	ReceivedPacket(context.Context, *QueryGetReceivedPacketRequest) (*QueryGetReceivedPacketResponse, error)
	// Queries a list of ReceivedPacket items.
	ReceivedPacketAll(context.Context, *QueryAllReceivedPacketRequest) (*QueryAllReceivedPacketResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) ReceivedPacket(ctx context.Context, req *QueryGetReceivedPacketRequest) (*QueryGetReceivedPacketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReceivedPacket not implemented")
}
func (*UnimplementedQueryServer) ReceivedPacketAll(ctx context.Context, req *QueryAllReceivedPacketRequest) (*QueryAllReceivedPacketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReceivedPacketAll not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockibc.mockibc.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ReceivedPacket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetReceivedPacketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ReceivedPacket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockibc.mockibc.Query/ReceivedPacket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ReceivedPacket(ctx, req.(*QueryGetReceivedPacketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ReceivedPacketAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllReceivedPacketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ReceivedPacketAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockibc.mockibc.Query/ReceivedPacketAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ReceivedPacketAll(ctx, req.(*QueryAllReceivedPacketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mockibc.mockibc.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "ReceivedPacket",
			Handler:    _Query_ReceivedPacket_Handler,
		},
		{
			MethodName: "ReceivedPacketAll",
			Handler:    _Query_ReceivedPacketAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mockibc/mockibc/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryGetReceivedPacketRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetReceivedPacketRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetReceivedPacketRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetReceivedPacketResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetReceivedPacketResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetReceivedPacketResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.ReceivedPacket.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryAllReceivedPacketRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllReceivedPacketRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllReceivedPacketRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllReceivedPacketResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllReceivedPacketResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllReceivedPacketResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.ReceivedPacket) > 0 {
		for iNdEx := len(m.ReceivedPacket) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ReceivedPacket[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryGetReceivedPacketRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovQuery(uint64(m.Id))
	}
	return n
}

func (m *QueryGetReceivedPacketResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ReceivedPacket.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryAllReceivedPacketRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllReceivedPacketResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ReceivedPacket) > 0 {
		for _, e := range m.ReceivedPacket {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryGetReceivedPacketRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGetReceivedPacketRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetReceivedPacketRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryGetReceivedPacketResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGetReceivedPacketResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetReceivedPacketResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReceivedPacket", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ReceivedPacket.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryAllReceivedPacketRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryAllReceivedPacketRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllReceivedPacketRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryAllReceivedPacketResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryAllReceivedPacketResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllReceivedPacketResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReceivedPacket", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ReceivedPacket = append(m.ReceivedPacket, ReceivedPacket{})
			if err := m.ReceivedPacket[len(m.ReceivedPacket)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
