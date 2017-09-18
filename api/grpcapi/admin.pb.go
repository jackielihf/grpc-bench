// Code generated by protoc-gen-go.
// source: admin.proto
// DO NOT EDIT!

package grpcapi

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 冲补基金持仓大小的消息
type SupplyOpenFundStockReq struct {
	BaseReq   *BaseReq `protobuf:"bytes,1,opt,name=baseReq" json:"baseReq,omitempty"`
	ClientId  int64    `protobuf:"varint,2,opt,name=clientId" json:"clientId,omitempty"`
	ProjectId int64    `protobuf:"varint,3,opt,name=projectId" json:"projectId,omitempty"`
	StockCode string   `protobuf:"bytes,4,opt,name=stockCode" json:"stockCode,omitempty"`
	Amount    float64  `protobuf:"fixed64,5,opt,name=amount" json:"amount,omitempty"`
}

func (m *SupplyOpenFundStockReq) Reset()                    { *m = SupplyOpenFundStockReq{} }
func (m *SupplyOpenFundStockReq) String() string            { return proto.CompactTextString(m) }
func (*SupplyOpenFundStockReq) ProtoMessage()               {}
func (*SupplyOpenFundStockReq) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *SupplyOpenFundStockReq) GetBaseReq() *BaseReq {
	if m != nil {
		return m.BaseReq
	}
	return nil
}

func (m *SupplyOpenFundStockReq) GetClientId() int64 {
	if m != nil {
		return m.ClientId
	}
	return 0
}

func (m *SupplyOpenFundStockReq) GetProjectId() int64 {
	if m != nil {
		return m.ProjectId
	}
	return 0
}

func (m *SupplyOpenFundStockReq) GetStockCode() string {
	if m != nil {
		return m.StockCode
	}
	return ""
}

func (m *SupplyOpenFundStockReq) GetAmount() float64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

// 冲补基金持仓大小的消息响应
type SupplyOpenFundStockResp struct {
	RespStatus *RespStatus `protobuf:"bytes,1,opt,name=respStatus" json:"respStatus,omitempty"`
}

func (m *SupplyOpenFundStockResp) Reset()                    { *m = SupplyOpenFundStockResp{} }
func (m *SupplyOpenFundStockResp) String() string            { return proto.CompactTextString(m) }
func (*SupplyOpenFundStockResp) ProtoMessage()               {}
func (*SupplyOpenFundStockResp) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *SupplyOpenFundStockResp) GetRespStatus() *RespStatus {
	if m != nil {
		return m.RespStatus
	}
	return nil
}

func init() {
	proto.RegisterType((*SupplyOpenFundStockReq)(nil), "grpcapi.SupplyOpenFundStockReq")
	proto.RegisterType((*SupplyOpenFundStockResp)(nil), "grpcapi.SupplyOpenFundStockResp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for OpenFundAdmin service

type OpenFundAdminClient interface {
	// 红冲蓝补基金持仓
	// 可用于初始化资产单元的基金持仓
	SupplyOpenFundStock(ctx context.Context, in *SupplyOpenFundStockReq, opts ...grpc.CallOption) (*SupplyOpenFundStockResp, error)
}

type openFundAdminClient struct {
	cc *grpc.ClientConn
}

func NewOpenFundAdminClient(cc *grpc.ClientConn) OpenFundAdminClient {
	return &openFundAdminClient{cc}
}

func (c *openFundAdminClient) SupplyOpenFundStock(ctx context.Context, in *SupplyOpenFundStockReq, opts ...grpc.CallOption) (*SupplyOpenFundStockResp, error) {
	out := new(SupplyOpenFundStockResp)
	err := grpc.Invoke(ctx, "/grpcapi.OpenFundAdmin/SupplyOpenFundStock", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for OpenFundAdmin service

type OpenFundAdminServer interface {
	// 红冲蓝补基金持仓
	// 可用于初始化资产单元的基金持仓
	SupplyOpenFundStock(context.Context, *SupplyOpenFundStockReq) (*SupplyOpenFundStockResp, error)
}

func RegisterOpenFundAdminServer(s *grpc.Server, srv OpenFundAdminServer) {
	s.RegisterService(&_OpenFundAdmin_serviceDesc, srv)
}

func _OpenFundAdmin_SupplyOpenFundStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SupplyOpenFundStockReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenFundAdminServer).SupplyOpenFundStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcapi.OpenFundAdmin/SupplyOpenFundStock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenFundAdminServer).SupplyOpenFundStock(ctx, req.(*SupplyOpenFundStockReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _OpenFundAdmin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpcapi.OpenFundAdmin",
	HandlerType: (*OpenFundAdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SupplyOpenFundStock",
			Handler:    _OpenFundAdmin_SupplyOpenFundStock_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admin.proto",
}

func init() { proto.RegisterFile("admin.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 317 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0xc1, 0x4e, 0xea, 0x40,
	0x14, 0xcd, 0xc0, 0x7b, 0x20, 0x17, 0x4d, 0xcc, 0x10, 0xa1, 0xa9, 0x24, 0x56, 0x56, 0x0d, 0x8b,
	0x36, 0x81, 0x9d, 0x3b, 0x31, 0x31, 0x71, 0xa3, 0xc9, 0xf0, 0x05, 0x43, 0x3b, 0x34, 0x55, 0x3a,
	0x77, 0xe8, 0x4c, 0x4d, 0x8c, 0x3b, 0xd7, 0xee, 0xfc, 0x15, 0xff, 0xc4, 0x5f, 0xf0, 0x43, 0x4c,
	0xcb, 0x50, 0x5c, 0x34, 0x2e, 0xef, 0x3d, 0x67, 0xce, 0x39, 0x73, 0x2e, 0xf4, 0x79, 0x9c, 0xa5,
	0x32, 0x50, 0x39, 0x1a, 0xa4, 0xdd, 0x24, 0x57, 0x11, 0x57, 0xa9, 0x3b, 0x4e, 0x10, 0x93, 0x8d,
	0x08, 0xb9, 0x4a, 0x43, 0x2e, 0x25, 0x1a, 0x6e, 0x52, 0x94, 0x7a, 0x47, 0x73, 0x8f, 0x23, 0xcc,
	0x32, 0xb4, 0x8f, 0x26, 0x9f, 0x04, 0x86, 0xcb, 0x42, 0xa9, 0xcd, 0xcb, 0x83, 0x12, 0xf2, 0xb6,
	0x90, 0xf1, 0xd2, 0x60, 0xf4, 0xc4, 0xc4, 0x96, 0x4e, 0xa1, 0xbb, 0xe2, 0x5a, 0x30, 0xb1, 0x75,
	0x88, 0x47, 0xfc, 0xfe, 0xec, 0x34, 0xb0, 0x0e, 0xc1, 0x62, 0xb7, 0x67, 0x7b, 0x02, 0x75, 0xe1,
	0x28, 0xda, 0xa4, 0x42, 0x9a, 0xbb, 0xd8, 0x69, 0x79, 0xc4, 0x6f, 0xb3, 0x7a, 0xa6, 0x63, 0xe8,
	0xa9, 0x1c, 0x1f, 0x45, 0x54, 0x82, 0xed, 0x0a, 0x3c, 0x2c, 0x4a, 0x54, 0x97, 0x8e, 0x37, 0x18,
	0x0b, 0xe7, 0x9f, 0x47, 0xfc, 0x1e, 0x3b, 0x2c, 0xe8, 0x10, 0x3a, 0x3c, 0xc3, 0x42, 0x1a, 0xe7,
	0xbf, 0x47, 0x7c, 0xc2, 0xec, 0x34, 0xb9, 0x87, 0x51, 0x63, 0x6a, 0xad, 0xe8, 0x1c, 0x20, 0x17,
	0x5a, 0x2d, 0x0d, 0x37, 0x85, 0xb6, 0xc9, 0x07, 0x75, 0x72, 0x56, 0x43, 0xec, 0x17, 0x6d, 0xf6,
	0x4e, 0xe0, 0x64, 0x2f, 0x75, 0x5d, 0x76, 0x4a, 0x5f, 0x61, 0xd0, 0xe0, 0x40, 0x2f, 0x6a, 0xa5,
	0xe6, 0xd6, 0x5c, 0xef, 0x6f, 0x82, 0x56, 0x93, 0xcb, 0xb7, 0xaf, 0xef, 0x8f, 0xd6, 0xb9, 0x3b,
	0x0c, 0xab, 0xeb, 0x85, 0xb6, 0x8b, 0x10, 0xd7, 0xd5, 0xbf, 0xaf, 0xc8, 0x74, 0x31, 0x82, 0xb3,
	0x08, 0xb3, 0x20, 0x59, 0x07, 0xcf, 0x3c, 0x28, 0xb5, 0xac, 0xe6, 0xaa, 0x53, 0x5d, 0x6d, 0xfe,
	0x13, 0x00, 0x00, 0xff, 0xff, 0x71, 0x7f, 0x40, 0x39, 0xf9, 0x01, 0x00, 0x00,
}
