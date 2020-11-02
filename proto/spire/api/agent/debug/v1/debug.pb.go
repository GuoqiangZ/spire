// Code generated by protoc-gen-go. DO NOT EDIT.
// source: spire/api/agent/debug/v1/debug.proto

package debug

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	types "github.com/spiffe/spire/proto/spire/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GetInfoRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetInfoRequest) Reset()         { *m = GetInfoRequest{} }
func (m *GetInfoRequest) String() string { return proto.CompactTextString(m) }
func (*GetInfoRequest) ProtoMessage()    {}
func (*GetInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e5721b49b138bf5, []int{0}
}

func (m *GetInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetInfoRequest.Unmarshal(m, b)
}
func (m *GetInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetInfoRequest.Marshal(b, m, deterministic)
}
func (m *GetInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetInfoRequest.Merge(m, src)
}
func (m *GetInfoRequest) XXX_Size() int {
	return xxx_messageInfo_GetInfoRequest.Size(m)
}
func (m *GetInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetInfoRequest proto.InternalMessageInfo

type GetInfoResponse struct {
	// Agent SVID chain
	SvidChain []*GetInfoResponse_Cert `protobuf:"bytes,1,rep,name=svid_chain,json=svidChain,proto3" json:"svid_chain,omitempty"`
	// Agent uptime in seconds
	Uptime int32 `protobuf:"varint,2,opt,name=uptime,proto3" json:"uptime,omitempty"`
	// Number of SVIDs cached in memory
	SvidsCount int32 `protobuf:"varint,3,opt,name=svids_count,json=svidsCount,proto3" json:"svids_count,omitempty"`
	// last successful sync with server (in seconds since unix epoch)
	LastSyncSuccess      int64    `protobuf:"varint,4,opt,name=last_sync_success,json=lastSyncSuccess,proto3" json:"last_sync_success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetInfoResponse) Reset()         { *m = GetInfoResponse{} }
func (m *GetInfoResponse) String() string { return proto.CompactTextString(m) }
func (*GetInfoResponse) ProtoMessage()    {}
func (*GetInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e5721b49b138bf5, []int{1}
}

func (m *GetInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetInfoResponse.Unmarshal(m, b)
}
func (m *GetInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetInfoResponse.Marshal(b, m, deterministic)
}
func (m *GetInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetInfoResponse.Merge(m, src)
}
func (m *GetInfoResponse) XXX_Size() int {
	return xxx_messageInfo_GetInfoResponse.Size(m)
}
func (m *GetInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetInfoResponse proto.InternalMessageInfo

func (m *GetInfoResponse) GetSvidChain() []*GetInfoResponse_Cert {
	if m != nil {
		return m.SvidChain
	}
	return nil
}

func (m *GetInfoResponse) GetUptime() int32 {
	if m != nil {
		return m.Uptime
	}
	return 0
}

func (m *GetInfoResponse) GetSvidsCount() int32 {
	if m != nil {
		return m.SvidsCount
	}
	return 0
}

func (m *GetInfoResponse) GetLastSyncSuccess() int64 {
	if m != nil {
		return m.LastSyncSuccess
	}
	return 0
}

type GetInfoResponse_Cert struct {
	// Cerfificate SPIFFE ID
	Id *types.SPIFFEID `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Expiration time
	ExpiresAt int64 `protobuf:"varint,2,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	// Subject
	Subject              string   `protobuf:"bytes,3,opt,name=subject,proto3" json:"subject,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetInfoResponse_Cert) Reset()         { *m = GetInfoResponse_Cert{} }
func (m *GetInfoResponse_Cert) String() string { return proto.CompactTextString(m) }
func (*GetInfoResponse_Cert) ProtoMessage()    {}
func (*GetInfoResponse_Cert) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e5721b49b138bf5, []int{1, 0}
}

func (m *GetInfoResponse_Cert) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetInfoResponse_Cert.Unmarshal(m, b)
}
func (m *GetInfoResponse_Cert) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetInfoResponse_Cert.Marshal(b, m, deterministic)
}
func (m *GetInfoResponse_Cert) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetInfoResponse_Cert.Merge(m, src)
}
func (m *GetInfoResponse_Cert) XXX_Size() int {
	return xxx_messageInfo_GetInfoResponse_Cert.Size(m)
}
func (m *GetInfoResponse_Cert) XXX_DiscardUnknown() {
	xxx_messageInfo_GetInfoResponse_Cert.DiscardUnknown(m)
}

var xxx_messageInfo_GetInfoResponse_Cert proto.InternalMessageInfo

func (m *GetInfoResponse_Cert) GetId() *types.SPIFFEID {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *GetInfoResponse_Cert) GetExpiresAt() int64 {
	if m != nil {
		return m.ExpiresAt
	}
	return 0
}

func (m *GetInfoResponse_Cert) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func init() {
	proto.RegisterType((*GetInfoRequest)(nil), "spire.agent.debug.v1.GetInfoRequest")
	proto.RegisterType((*GetInfoResponse)(nil), "spire.agent.debug.v1.GetInfoResponse")
	proto.RegisterType((*GetInfoResponse_Cert)(nil), "spire.agent.debug.v1.GetInfoResponse.Cert")
}

func init() {
	proto.RegisterFile("spire/api/agent/debug/v1/debug.proto", fileDescriptor_4e5721b49b138bf5)
}

var fileDescriptor_4e5721b49b138bf5 = []byte{
	// 353 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x4f, 0x4b, 0xeb, 0x40,
	0x10, 0x27, 0x49, 0xff, 0xd0, 0x29, 0xbc, 0xbe, 0xb7, 0x3c, 0x25, 0x04, 0xc4, 0x50, 0x5a, 0x08,
	0x3d, 0x6c, 0x68, 0x3d, 0x2a, 0x82, 0xb6, 0x56, 0x72, 0x93, 0x14, 0x3c, 0x78, 0x09, 0xc9, 0x66,
	0xd3, 0xae, 0xd8, 0x24, 0x76, 0x36, 0xc5, 0x7e, 0x24, 0xbf, 0xa5, 0xec, 0x26, 0x7a, 0x10, 0x45,
	0x6f, 0x3b, 0xbf, 0x3f, 0xcc, 0xcc, 0x6f, 0x16, 0x46, 0x58, 0x8a, 0x1d, 0xf7, 0xe3, 0x52, 0xf8,
	0xf1, 0x9a, 0xe7, 0xd2, 0x4f, 0x79, 0x52, 0xad, 0xfd, 0xfd, 0xb4, 0x7e, 0xd0, 0x72, 0x57, 0xc8,
	0x82, 0xfc, 0xd7, 0x2a, 0xaa, 0x15, 0xb4, 0x26, 0xf6, 0x53, 0xc7, 0xa9, 0xbd, 0xf2, 0x50, 0x72,
	0xf4, 0xb1, 0x14, 0x59, 0xc6, 0x45, 0x5a, 0x3b, 0x86, 0x7f, 0xe1, 0xcf, 0x2d, 0x97, 0x41, 0x9e,
	0x15, 0x21, 0x7f, 0xae, 0x38, 0xca, 0xe1, 0xab, 0x09, 0x83, 0x0f, 0x08, 0xcb, 0x22, 0x47, 0x4e,
	0x02, 0x00, 0xdc, 0x8b, 0x34, 0x62, 0x9b, 0x58, 0xe4, 0xb6, 0xe1, 0x5a, 0x5e, 0x7f, 0x36, 0xa1,
	0x5f, 0x35, 0xa3, 0x9f, 0xac, 0x74, 0xce, 0x77, 0x32, 0xec, 0x29, 0xf7, 0x5c, 0x99, 0xc9, 0x31,
	0x74, 0xaa, 0x52, 0x8a, 0x2d, 0xb7, 0x4d, 0xd7, 0xf0, 0xda, 0x61, 0x53, 0x91, 0x53, 0xe8, 0x2b,
	0x11, 0x46, 0xac, 0xa8, 0x72, 0x69, 0x5b, 0x9a, 0xd4, 0x5d, 0x71, 0xae, 0x10, 0x32, 0x81, 0x7f,
	0x4f, 0x31, 0xca, 0x08, 0x0f, 0x39, 0x8b, 0xb0, 0x62, 0x8c, 0x23, 0xda, 0x2d, 0xd7, 0xf0, 0xac,
	0x70, 0xa0, 0x88, 0xd5, 0x21, 0x67, 0xab, 0x1a, 0x76, 0x32, 0x68, 0xa9, 0xbe, 0x64, 0x0c, 0xa6,
	0x48, 0x6d, 0xc3, 0x35, 0xbc, 0xfe, 0xec, 0xa8, 0x99, 0x57, 0xc7, 0x40, 0x57, 0x77, 0xc1, 0x72,
	0x79, 0x13, 0x2c, 0x42, 0x53, 0xa4, 0xe4, 0x04, 0x80, 0xbf, 0x28, 0x12, 0xa3, 0x58, 0xea, 0xb9,
	0xac, 0xb0, 0xd7, 0x20, 0x57, 0x92, 0xd8, 0xd0, 0xc5, 0x2a, 0x79, 0xe4, 0xac, 0x1e, 0xab, 0x17,
	0xbe, 0x97, 0xb3, 0x08, 0xda, 0x0b, 0xb5, 0x38, 0xb9, 0x87, 0x6e, 0xb3, 0x38, 0x19, 0xfd, 0x90,
	0x8b, 0x4e, 0xd9, 0x19, 0xff, 0x2a, 0xbd, 0xeb, 0xcb, 0x87, 0x8b, 0xb5, 0x90, 0x9b, 0x2a, 0xa1,
	0xac, 0xd8, 0x36, 0xb7, 0xf3, 0xeb, 0x73, 0xea, 0xfb, 0xf9, 0xdf, 0x7d, 0x8b, 0x73, 0xfd, 0x48,
	0x3a, 0x5a, 0x75, 0xf6, 0x16, 0x00, 0x00, 0xff, 0xff, 0x2a, 0x1c, 0xda, 0xd6, 0x3f, 0x02, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DebugClient is the client API for Debug service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DebugClient interface {
	// Get information about SPIRE agent
	GetInfo(ctx context.Context, in *GetInfoRequest, opts ...grpc.CallOption) (*GetInfoResponse, error)
}

type debugClient struct {
	cc grpc.ClientConnInterface
}

func NewDebugClient(cc grpc.ClientConnInterface) DebugClient {
	return &debugClient{cc}
}

func (c *debugClient) GetInfo(ctx context.Context, in *GetInfoRequest, opts ...grpc.CallOption) (*GetInfoResponse, error) {
	out := new(GetInfoResponse)
	err := c.cc.Invoke(ctx, "/spire.agent.debug.v1.Debug/GetInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DebugServer is the server API for Debug service.
type DebugServer interface {
	// Get information about SPIRE agent
	GetInfo(context.Context, *GetInfoRequest) (*GetInfoResponse, error)
}

// UnimplementedDebugServer can be embedded to have forward compatible implementations.
type UnimplementedDebugServer struct {
}

func (*UnimplementedDebugServer) GetInfo(ctx context.Context, req *GetInfoRequest) (*GetInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInfo not implemented")
}

func RegisterDebugServer(s *grpc.Server, srv DebugServer) {
	s.RegisterService(&_Debug_serviceDesc, srv)
}

func _Debug_GetInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DebugServer).GetInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.agent.debug.v1.Debug/GetInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DebugServer).GetInfo(ctx, req.(*GetInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Debug_serviceDesc = grpc.ServiceDesc{
	ServiceName: "spire.agent.debug.v1.Debug",
	HandlerType: (*DebugServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetInfo",
			Handler:    _Debug_GetInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spire/api/agent/debug/v1/debug.proto",
}