// Code generated by protoc-gen-go.
// source: common.proto
// DO NOT EDIT!

package grpcapi

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 基本请求信息
type BaseReq struct {
	Ver          string `protobuf:"bytes,1,opt,name=ver" json:"ver,omitempty"`
	Uuid         string `protobuf:"bytes,2,opt,name=uuid" json:"uuid,omitempty"`
	Src          string `protobuf:"bytes,3,opt,name=src" json:"src,omitempty"`
	Token        string `protobuf:"bytes,4,opt,name=token" json:"token,omitempty"`
	Station      string `protobuf:"bytes,5,opt,name=station" json:"station,omitempty"`
	OperatorType string `protobuf:"bytes,6,opt,name=operatorType" json:"operatorType,omitempty"`
	Operator     string `protobuf:"bytes,7,opt,name=operator" json:"operator,omitempty"`
	EntrustWay   string `protobuf:"bytes,8,opt,name=entrustWay" json:"entrustWay,omitempty"`
}

func (m *BaseReq) Reset()                    { *m = BaseReq{} }
func (m *BaseReq) String() string            { return proto.CompactTextString(m) }
func (*BaseReq) ProtoMessage()               {}
func (*BaseReq) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *BaseReq) GetVer() string {
	if m != nil {
		return m.Ver
	}
	return ""
}

func (m *BaseReq) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *BaseReq) GetSrc() string {
	if m != nil {
		return m.Src
	}
	return ""
}

func (m *BaseReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *BaseReq) GetStation() string {
	if m != nil {
		return m.Station
	}
	return ""
}

func (m *BaseReq) GetOperatorType() string {
	if m != nil {
		return m.OperatorType
	}
	return ""
}

func (m *BaseReq) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *BaseReq) GetEntrustWay() string {
	if m != nil {
		return m.EntrustWay
	}
	return ""
}

// 响应状态
type RespStatus struct {
	Code int32  `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
}

func (m *RespStatus) Reset()                    { *m = RespStatus{} }
func (m *RespStatus) String() string            { return proto.CompactTextString(m) }
func (*RespStatus) ProtoMessage()               {}
func (*RespStatus) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *RespStatus) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *RespStatus) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

// 分页
type PageInfo struct {
	Total int64  `protobuf:"varint,1,opt,name=total" json:"total,omitempty"`
	Page  int64  `protobuf:"varint,2,opt,name=page" json:"page,omitempty"`
	Limit int64  `protobuf:"varint,3,opt,name=limit" json:"limit,omitempty"`
	Sort  string `protobuf:"bytes,4,opt,name=sort" json:"sort,omitempty"`
	Asc   string `protobuf:"bytes,5,opt,name=asc" json:"asc,omitempty"`
}

func (m *PageInfo) Reset()                    { *m = PageInfo{} }
func (m *PageInfo) String() string            { return proto.CompactTextString(m) }
func (*PageInfo) ProtoMessage()               {}
func (*PageInfo) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *PageInfo) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *PageInfo) GetPage() int64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *PageInfo) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *PageInfo) GetSort() string {
	if m != nil {
		return m.Sort
	}
	return ""
}

func (m *PageInfo) GetAsc() string {
	if m != nil {
		return m.Asc
	}
	return ""
}

// 交易请求头信息
type TradeHeader struct {
	PasswordType string `protobuf:"bytes,1,opt,name=passwordType" json:"passwordType,omitempty"`
	Password     string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
	OpBranchNo   string `protobuf:"bytes,3,opt,name=opBranchNo" json:"opBranchNo,omitempty"`
	BranchNo     string `protobuf:"bytes,4,opt,name=branchNo" json:"branchNo,omitempty"`
}

func (m *TradeHeader) Reset()                    { *m = TradeHeader{} }
func (m *TradeHeader) String() string            { return proto.CompactTextString(m) }
func (*TradeHeader) ProtoMessage()               {}
func (*TradeHeader) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

func (m *TradeHeader) GetPasswordType() string {
	if m != nil {
		return m.PasswordType
	}
	return ""
}

func (m *TradeHeader) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *TradeHeader) GetOpBranchNo() string {
	if m != nil {
		return m.OpBranchNo
	}
	return ""
}

func (m *TradeHeader) GetBranchNo() string {
	if m != nil {
		return m.BranchNo
	}
	return ""
}

func init() {
	proto.RegisterType((*BaseReq)(nil), "grpcapi.BaseReq")
	proto.RegisterType((*RespStatus)(nil), "grpcapi.RespStatus")
	proto.RegisterType((*PageInfo)(nil), "grpcapi.PageInfo")
	proto.RegisterType((*TradeHeader)(nil), "grpcapi.TradeHeader")
}

func init() { proto.RegisterFile("common.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 339 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x92, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x15, 0xd2, 0x36, 0xe5, 0xe8, 0x80, 0x2c, 0x10, 0x16, 0x03, 0x42, 0x99, 0x98, 0x32,
	0xc0, 0x1b, 0x74, 0x82, 0x05, 0xa1, 0x50, 0x89, 0xf9, 0xea, 0xb8, 0x21, 0xa2, 0xc9, 0x19, 0xdb,
	0x29, 0xea, 0x23, 0xf0, 0x78, 0xbc, 0x11, 0xf2, 0x39, 0xa9, 0xda, 0xed, 0xff, 0xfe, 0xbb, 0x4b,
	0x7c, 0xbf, 0x0d, 0x0b, 0x45, 0x6d, 0x4b, 0x5d, 0x61, 0x2c, 0x79, 0x12, 0x59, 0x6d, 0x8d, 0x42,
	0xd3, 0xe4, 0x7f, 0x09, 0x64, 0x4b, 0x74, 0xba, 0xd4, 0xdf, 0xe2, 0x12, 0xd2, 0x9d, 0xb6, 0x32,
	0xb9, 0x4f, 0x1e, 0xce, 0xcb, 0x20, 0x85, 0x80, 0x49, 0xdf, 0x37, 0x95, 0x3c, 0x63, 0x8b, 0x75,
	0xe8, 0x72, 0x56, 0xc9, 0x34, 0x76, 0x39, 0xab, 0xc4, 0x15, 0x4c, 0x3d, 0x7d, 0xe9, 0x4e, 0x4e,
	0xd8, 0x8b, 0x20, 0x24, 0x64, 0xce, 0xa3, 0x6f, 0xa8, 0x93, 0x53, 0xf6, 0x47, 0x14, 0x39, 0x2c,
	0xc8, 0x68, 0x8b, 0x9e, 0xec, 0x6a, 0x6f, 0xb4, 0x9c, 0x71, 0xf9, 0xc4, 0x13, 0xb7, 0x30, 0x1f,
	0x59, 0x66, 0x5c, 0x3f, 0xb0, 0xb8, 0x03, 0xd0, 0x9d, 0xb7, 0xbd, 0xf3, 0x1f, 0xb8, 0x97, 0x73,
	0xae, 0x1e, 0x39, 0xf9, 0x23, 0x40, 0xa9, 0x9d, 0x79, 0xf7, 0xe8, 0x7b, 0x17, 0x76, 0x50, 0x54,
	0x69, 0x5e, 0x6b, 0x5a, 0xb2, 0x0e, 0x3b, 0xb4, 0xae, 0x1e, 0xd6, 0x0a, 0x32, 0x37, 0x30, 0x7f,
	0xc3, 0x5a, 0xbf, 0x74, 0x1b, 0x8a, 0xfb, 0x78, 0xdc, 0xf2, 0x48, 0x5a, 0x46, 0x08, 0xdf, 0x31,
	0x58, 0x6b, 0x1e, 0x4a, 0x4b, 0xd6, 0xa1, 0x73, 0xdb, 0xb4, 0x8d, 0xe7, 0x34, 0xd2, 0x32, 0x42,
	0xe8, 0x74, 0x64, 0xfd, 0x10, 0x07, 0xeb, 0xf0, 0x47, 0x74, 0x6a, 0x48, 0x22, 0xc8, 0xfc, 0x37,
	0x81, 0x8b, 0x95, 0xc5, 0x4a, 0x3f, 0x6b, 0xac, 0xb4, 0x0d, 0xa9, 0x18, 0x74, 0xee, 0x87, 0x6c,
	0xc5, 0xa9, 0xc4, 0x6b, 0x38, 0xf1, 0x42, 0x2a, 0x23, 0x0f, 0x87, 0x3f, 0x70, 0x48, 0x85, 0xcc,
	0xd2, 0x62, 0xa7, 0x3e, 0x5f, 0x69, 0xb8, 0x9e, 0x23, 0x27, 0xcc, 0xae, 0xc7, 0x6a, 0x3c, 0xd9,
	0x81, 0x97, 0x37, 0x70, 0xad, 0xa8, 0x2d, 0xea, 0x4d, 0xb1, 0xc3, 0x02, 0x4d, 0x53, 0x0c, 0xcf,
	0x63, 0x3d, 0xe3, 0xe7, 0xf2, 0xf4, 0x1f, 0x00, 0x00, 0xff, 0xff, 0x5b, 0xa5, 0x87, 0xb6, 0x3e,
	0x02, 0x00, 0x00,
}
