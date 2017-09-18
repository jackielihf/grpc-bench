// Code generated by protoc-gen-go.
// source: fund_real.proto
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

// 划转资金的消息
type TransferReq struct {
	BaseReq   *BaseReq `protobuf:"bytes,1,opt,name=baseReq" json:"baseReq,omitempty"`
	ClientId  int64    `protobuf:"varint,2,opt,name=clientId" json:"clientId,omitempty"`
	ProjectId int64    `protobuf:"varint,3,opt,name=projectId" json:"projectId,omitempty"`
	Balance   float64  `protobuf:"fixed64,4,opt,name=balance" json:"balance,omitempty"`
}

func (m *TransferReq) Reset()                    { *m = TransferReq{} }
func (m *TransferReq) String() string            { return proto.CompactTextString(m) }
func (*TransferReq) ProtoMessage()               {}
func (*TransferReq) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *TransferReq) GetBaseReq() *BaseReq {
	if m != nil {
		return m.BaseReq
	}
	return nil
}

func (m *TransferReq) GetClientId() int64 {
	if m != nil {
		return m.ClientId
	}
	return 0
}

func (m *TransferReq) GetProjectId() int64 {
	if m != nil {
		return m.ProjectId
	}
	return 0
}

func (m *TransferReq) GetBalance() float64 {
	if m != nil {
		return m.Balance
	}
	return 0
}

// 划转资金的响应
type TransferResp struct {
	RespStatus *RespStatus `protobuf:"bytes,1,opt,name=respStatus" json:"respStatus,omitempty"`
}

func (m *TransferResp) Reset()                    { *m = TransferResp{} }
func (m *TransferResp) String() string            { return proto.CompactTextString(m) }
func (*TransferResp) ProtoMessage()               {}
func (*TransferResp) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *TransferResp) GetRespStatus() *RespStatus {
	if m != nil {
		return m.RespStatus
	}
	return nil
}

// 查询可用资金分配情况的消息
type QueryAssigningReq struct {
	BaseReq         *BaseReq `protobuf:"bytes,1,opt,name=baseReq" json:"baseReq,omitempty"`
	ClientId        int64    `protobuf:"varint,2,opt,name=clientId" json:"clientId,omitempty"`
	ProjectId       int64    `protobuf:"varint,3,opt,name=projectId" json:"projectId,omitempty"`
	RealFundAccount string   `protobuf:"bytes,4,opt,name=realFundAccount" json:"realFundAccount,omitempty"`
	RealClientType  string   `protobuf:"bytes,5,opt,name=realClientType" json:"realClientType,omitempty"`
	RealClientId    string   `protobuf:"bytes,6,opt,name=realClientId" json:"realClientId,omitempty"`
}

func (m *QueryAssigningReq) Reset()                    { *m = QueryAssigningReq{} }
func (m *QueryAssigningReq) String() string            { return proto.CompactTextString(m) }
func (*QueryAssigningReq) ProtoMessage()               {}
func (*QueryAssigningReq) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

func (m *QueryAssigningReq) GetBaseReq() *BaseReq {
	if m != nil {
		return m.BaseReq
	}
	return nil
}

func (m *QueryAssigningReq) GetClientId() int64 {
	if m != nil {
		return m.ClientId
	}
	return 0
}

func (m *QueryAssigningReq) GetProjectId() int64 {
	if m != nil {
		return m.ProjectId
	}
	return 0
}

func (m *QueryAssigningReq) GetRealFundAccount() string {
	if m != nil {
		return m.RealFundAccount
	}
	return ""
}

func (m *QueryAssigningReq) GetRealClientType() string {
	if m != nil {
		return m.RealClientType
	}
	return ""
}

func (m *QueryAssigningReq) GetRealClientId() string {
	if m != nil {
		return m.RealClientId
	}
	return ""
}

// 查询可用资金分配情况的响应
type GetAssigningResp struct {
	RespStatus        *RespStatus `protobuf:"bytes,1,opt,name=respStatus" json:"respStatus,omitempty"`
	AssignedBalance   float64     `protobuf:"fixed64,2,opt,name=assignedBalance" json:"assignedBalance,omitempty"`
	UnAssignedBalance float64     `protobuf:"fixed64,3,opt,name=unAssignedBalance" json:"unAssignedBalance,omitempty"`
}

func (m *GetAssigningResp) Reset()                    { *m = GetAssigningResp{} }
func (m *GetAssigningResp) String() string            { return proto.CompactTextString(m) }
func (*GetAssigningResp) ProtoMessage()               {}
func (*GetAssigningResp) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

func (m *GetAssigningResp) GetRespStatus() *RespStatus {
	if m != nil {
		return m.RespStatus
	}
	return nil
}

func (m *GetAssigningResp) GetAssignedBalance() float64 {
	if m != nil {
		return m.AssignedBalance
	}
	return 0
}

func (m *GetAssigningResp) GetUnAssignedBalance() float64 {
	if m != nil {
		return m.UnAssignedBalance
	}
	return 0
}

// 查询资产单元的当前资金明细的消息
type QueryFundRealReq struct {
	BaseReq   *BaseReq `protobuf:"bytes,1,opt,name=baseReq" json:"baseReq,omitempty"`
	ClientId  int64    `protobuf:"varint,2,opt,name=clientId" json:"clientId,omitempty"`
	ProjectId int64    `protobuf:"varint,3,opt,name=projectId" json:"projectId,omitempty"`
}

func (m *QueryFundRealReq) Reset()                    { *m = QueryFundRealReq{} }
func (m *QueryFundRealReq) String() string            { return proto.CompactTextString(m) }
func (*QueryFundRealReq) ProtoMessage()               {}
func (*QueryFundRealReq) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{4} }

func (m *QueryFundRealReq) GetBaseReq() *BaseReq {
	if m != nil {
		return m.BaseReq
	}
	return nil
}

func (m *QueryFundRealReq) GetClientId() int64 {
	if m != nil {
		return m.ClientId
	}
	return 0
}

func (m *QueryFundRealReq) GetProjectId() int64 {
	if m != nil {
		return m.ProjectId
	}
	return 0
}

// 资产单元资金
type FundReal struct {
	ProjectId       int64   `protobuf:"varint,1,opt,name=projectId" json:"projectId,omitempty"`
	ClientId        int64   `protobuf:"varint,2,opt,name=clientId" json:"clientId,omitempty"`
	FundAccountId   int64   `protobuf:"varint,3,opt,name=fundAccountId" json:"fundAccountId,omitempty"`
	InitDate        int64   `protobuf:"varint,4,opt,name=initDate" json:"initDate,omitempty"`
	AssetType       string  `protobuf:"bytes,5,opt,name=assetType" json:"assetType,omitempty"`
	MoneyType       string  `protobuf:"bytes,6,opt,name=moneyType" json:"moneyType,omitempty"`
	CurrentBalance  float64 `protobuf:"fixed64,7,opt,name=currentBalance" json:"currentBalance,omitempty"`
	EnableBalance   float64 `protobuf:"fixed64,8,opt,name=enableBalance" json:"enableBalance,omitempty"`
	FrozenBalance   float64 `protobuf:"fixed64,9,opt,name=frozenBalance" json:"frozenBalance,omitempty"`
	BuyBalance      float64 `protobuf:"fixed64,10,opt,name=buyBalance" json:"buyBalance,omitempty"`
	RealBuyBalance  float64 `protobuf:"fixed64,11,opt,name=realBuyBalance" json:"realBuyBalance,omitempty"`
	RealSellBalance float64 `protobuf:"fixed64,12,opt,name=realSellBalance" json:"realSellBalance,omitempty"`
	CorrectBalance  float64 `protobuf:"fixed64,13,opt,name=correctBalance" json:"correctBalance,omitempty"`
	DebitBalance    float64 `protobuf:"fixed64,14,opt,name=debitBalance" json:"debitBalance,omitempty"`
	Remark          string  `protobuf:"bytes,15,opt,name=remark" json:"remark,omitempty"`
	UpdateAt        int64   `protobuf:"varint,17,opt,name=updateAt" json:"updateAt,omitempty"`
}

func (m *FundReal) Reset()                    { *m = FundReal{} }
func (m *FundReal) String() string            { return proto.CompactTextString(m) }
func (*FundReal) ProtoMessage()               {}
func (*FundReal) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{5} }

func (m *FundReal) GetProjectId() int64 {
	if m != nil {
		return m.ProjectId
	}
	return 0
}

func (m *FundReal) GetClientId() int64 {
	if m != nil {
		return m.ClientId
	}
	return 0
}

func (m *FundReal) GetFundAccountId() int64 {
	if m != nil {
		return m.FundAccountId
	}
	return 0
}

func (m *FundReal) GetInitDate() int64 {
	if m != nil {
		return m.InitDate
	}
	return 0
}

func (m *FundReal) GetAssetType() string {
	if m != nil {
		return m.AssetType
	}
	return ""
}

func (m *FundReal) GetMoneyType() string {
	if m != nil {
		return m.MoneyType
	}
	return ""
}

func (m *FundReal) GetCurrentBalance() float64 {
	if m != nil {
		return m.CurrentBalance
	}
	return 0
}

func (m *FundReal) GetEnableBalance() float64 {
	if m != nil {
		return m.EnableBalance
	}
	return 0
}

func (m *FundReal) GetFrozenBalance() float64 {
	if m != nil {
		return m.FrozenBalance
	}
	return 0
}

func (m *FundReal) GetBuyBalance() float64 {
	if m != nil {
		return m.BuyBalance
	}
	return 0
}

func (m *FundReal) GetRealBuyBalance() float64 {
	if m != nil {
		return m.RealBuyBalance
	}
	return 0
}

func (m *FundReal) GetRealSellBalance() float64 {
	if m != nil {
		return m.RealSellBalance
	}
	return 0
}

func (m *FundReal) GetCorrectBalance() float64 {
	if m != nil {
		return m.CorrectBalance
	}
	return 0
}

func (m *FundReal) GetDebitBalance() float64 {
	if m != nil {
		return m.DebitBalance
	}
	return 0
}

func (m *FundReal) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *FundReal) GetUpdateAt() int64 {
	if m != nil {
		return m.UpdateAt
	}
	return 0
}

// 查询资金的响应消息
type GetFundRealResp struct {
	RespStatus *RespStatus `protobuf:"bytes,1,opt,name=respStatus" json:"respStatus,omitempty"`
	Data       *FundReal   `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
}

func (m *GetFundRealResp) Reset()                    { *m = GetFundRealResp{} }
func (m *GetFundRealResp) String() string            { return proto.CompactTextString(m) }
func (*GetFundRealResp) ProtoMessage()               {}
func (*GetFundRealResp) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{6} }

func (m *GetFundRealResp) GetRespStatus() *RespStatus {
	if m != nil {
		return m.RespStatus
	}
	return nil
}

func (m *GetFundRealResp) GetData() *FundReal {
	if m != nil {
		return m.Data
	}
	return nil
}

// 查询资产单元的当前持仓
type QueryFundJourReq struct {
	BaseReq      *BaseReq  `protobuf:"bytes,1,opt,name=baseReq" json:"baseReq,omitempty"`
	PageInfo     *PageInfo `protobuf:"bytes,2,opt,name=pageInfo" json:"pageInfo,omitempty"`
	ClientId     int64     `protobuf:"varint,3,opt,name=clientId" json:"clientId,omitempty"`
	ProjectId    int64     `protobuf:"varint,4,opt,name=projectId" json:"projectId,omitempty"`
	BusinessFlag string    `protobuf:"bytes,5,opt,name=businessFlag" json:"businessFlag,omitempty"`
}

func (m *QueryFundJourReq) Reset()                    { *m = QueryFundJourReq{} }
func (m *QueryFundJourReq) String() string            { return proto.CompactTextString(m) }
func (*QueryFundJourReq) ProtoMessage()               {}
func (*QueryFundJourReq) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{7} }

func (m *QueryFundJourReq) GetBaseReq() *BaseReq {
	if m != nil {
		return m.BaseReq
	}
	return nil
}

func (m *QueryFundJourReq) GetPageInfo() *PageInfo {
	if m != nil {
		return m.PageInfo
	}
	return nil
}

func (m *QueryFundJourReq) GetClientId() int64 {
	if m != nil {
		return m.ClientId
	}
	return 0
}

func (m *QueryFundJourReq) GetProjectId() int64 {
	if m != nil {
		return m.ProjectId
	}
	return 0
}

func (m *QueryFundJourReq) GetBusinessFlag() string {
	if m != nil {
		return m.BusinessFlag
	}
	return ""
}

// 资产变动的消息体
type FundJour struct {
	ProjectId          int64   `protobuf:"varint,1,opt,name=projectId" json:"projectId,omitempty"`
	ClientId           int64   `protobuf:"varint,2,opt,name=clientId" json:"clientId,omitempty"`
	FundAccountId      int64   `protobuf:"varint,3,opt,name=fundAccountId" json:"fundAccountId,omitempty"`
	InitDate           int64   `protobuf:"varint,4,opt,name=initDate" json:"initDate,omitempty"`
	FinanceType        string  `protobuf:"bytes,5,opt,name=financeType" json:"financeType,omitempty"`
	BusinessFlag       string  `protobuf:"bytes,6,opt,name=businessFlag" json:"businessFlag,omitempty"`
	OccurEnableBalance float64 `protobuf:"fixed64,7,opt,name=occurEnableBalance" json:"occurEnableBalance,omitempty"`
}

func (m *FundJour) Reset()                    { *m = FundJour{} }
func (m *FundJour) String() string            { return proto.CompactTextString(m) }
func (*FundJour) ProtoMessage()               {}
func (*FundJour) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{8} }

func (m *FundJour) GetProjectId() int64 {
	if m != nil {
		return m.ProjectId
	}
	return 0
}

func (m *FundJour) GetClientId() int64 {
	if m != nil {
		return m.ClientId
	}
	return 0
}

func (m *FundJour) GetFundAccountId() int64 {
	if m != nil {
		return m.FundAccountId
	}
	return 0
}

func (m *FundJour) GetInitDate() int64 {
	if m != nil {
		return m.InitDate
	}
	return 0
}

func (m *FundJour) GetFinanceType() string {
	if m != nil {
		return m.FinanceType
	}
	return ""
}

func (m *FundJour) GetBusinessFlag() string {
	if m != nil {
		return m.BusinessFlag
	}
	return ""
}

func (m *FundJour) GetOccurEnableBalance() float64 {
	if m != nil {
		return m.OccurEnableBalance
	}
	return 0
}

// 基金持仓列表的响应消息
type ListFundJourResp struct {
	RespStatus *RespStatus `protobuf:"bytes,1,opt,name=respStatus" json:"respStatus,omitempty"`
	Page       *PageInfo   `protobuf:"bytes,2,opt,name=page" json:"page,omitempty"`
	Data       []*FundJour `protobuf:"bytes,3,rep,name=data" json:"data,omitempty"`
}

func (m *ListFundJourResp) Reset()                    { *m = ListFundJourResp{} }
func (m *ListFundJourResp) String() string            { return proto.CompactTextString(m) }
func (*ListFundJourResp) ProtoMessage()               {}
func (*ListFundJourResp) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{9} }

func (m *ListFundJourResp) GetRespStatus() *RespStatus {
	if m != nil {
		return m.RespStatus
	}
	return nil
}

func (m *ListFundJourResp) GetPage() *PageInfo {
	if m != nil {
		return m.Page
	}
	return nil
}

func (m *ListFundJourResp) GetData() []*FundJour {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*TransferReq)(nil), "grpcapi.TransferReq")
	proto.RegisterType((*TransferResp)(nil), "grpcapi.TransferResp")
	proto.RegisterType((*QueryAssigningReq)(nil), "grpcapi.QueryAssigningReq")
	proto.RegisterType((*GetAssigningResp)(nil), "grpcapi.GetAssigningResp")
	proto.RegisterType((*QueryFundRealReq)(nil), "grpcapi.QueryFundRealReq")
	proto.RegisterType((*FundReal)(nil), "grpcapi.FundReal")
	proto.RegisterType((*GetFundRealResp)(nil), "grpcapi.GetFundRealResp")
	proto.RegisterType((*QueryFundJourReq)(nil), "grpcapi.QueryFundJourReq")
	proto.RegisterType((*FundJour)(nil), "grpcapi.FundJour")
	proto.RegisterType((*ListFundJourResp)(nil), "grpcapi.ListFundJourResp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for FundRealManager service

type FundRealManagerClient interface {
	// 资产单元划入资金
	TransferIn(ctx context.Context, in *TransferReq, opts ...grpc.CallOption) (*TransferResp, error)
	// 资产单元划出资金
	TransferOut(ctx context.Context, in *TransferReq, opts ...grpc.CallOption) (*TransferResp, error)
	// 资产单元借入资金
	Debit(ctx context.Context, in *TransferReq, opts ...grpc.CallOption) (*TransferResp, error)
	// 查询可用资金的分配情况
	GetAssigning(ctx context.Context, in *QueryAssigningReq, opts ...grpc.CallOption) (*GetAssigningResp, error)
	// 查询可用资金的分配情况
	GetAssigningByRealClient(ctx context.Context, in *QueryAssigningReq, opts ...grpc.CallOption) (*GetAssigningResp, error)
	// 查询资产单元的当前资金明细
	GetFundReal(ctx context.Context, in *QueryFundRealReq, opts ...grpc.CallOption) (*GetFundRealResp, error)
	// 查询资产单元的资产变动流水
	ListFundJour(ctx context.Context, in *QueryFundJourReq, opts ...grpc.CallOption) (*ListFundJourResp, error)
}

type fundRealManagerClient struct {
	cc *grpc.ClientConn
}

func NewFundRealManagerClient(cc *grpc.ClientConn) FundRealManagerClient {
	return &fundRealManagerClient{cc}
}

func (c *fundRealManagerClient) TransferIn(ctx context.Context, in *TransferReq, opts ...grpc.CallOption) (*TransferResp, error) {
	out := new(TransferResp)
	err := grpc.Invoke(ctx, "/grpcapi.FundRealManager/TransferIn", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fundRealManagerClient) TransferOut(ctx context.Context, in *TransferReq, opts ...grpc.CallOption) (*TransferResp, error) {
	out := new(TransferResp)
	err := grpc.Invoke(ctx, "/grpcapi.FundRealManager/TransferOut", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fundRealManagerClient) Debit(ctx context.Context, in *TransferReq, opts ...grpc.CallOption) (*TransferResp, error) {
	out := new(TransferResp)
	err := grpc.Invoke(ctx, "/grpcapi.FundRealManager/Debit", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fundRealManagerClient) GetAssigning(ctx context.Context, in *QueryAssigningReq, opts ...grpc.CallOption) (*GetAssigningResp, error) {
	out := new(GetAssigningResp)
	err := grpc.Invoke(ctx, "/grpcapi.FundRealManager/GetAssigning", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fundRealManagerClient) GetAssigningByRealClient(ctx context.Context, in *QueryAssigningReq, opts ...grpc.CallOption) (*GetAssigningResp, error) {
	out := new(GetAssigningResp)
	err := grpc.Invoke(ctx, "/grpcapi.FundRealManager/GetAssigningByRealClient", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fundRealManagerClient) GetFundReal(ctx context.Context, in *QueryFundRealReq, opts ...grpc.CallOption) (*GetFundRealResp, error) {
	out := new(GetFundRealResp)
	err := grpc.Invoke(ctx, "/grpcapi.FundRealManager/GetFundReal", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fundRealManagerClient) ListFundJour(ctx context.Context, in *QueryFundJourReq, opts ...grpc.CallOption) (*ListFundJourResp, error) {
	out := new(ListFundJourResp)
	err := grpc.Invoke(ctx, "/grpcapi.FundRealManager/ListFundJour", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for FundRealManager service

type FundRealManagerServer interface {
	// 资产单元划入资金
	TransferIn(context.Context, *TransferReq) (*TransferResp, error)
	// 资产单元划出资金
	TransferOut(context.Context, *TransferReq) (*TransferResp, error)
	// 资产单元借入资金
	Debit(context.Context, *TransferReq) (*TransferResp, error)
	// 查询可用资金的分配情况
	GetAssigning(context.Context, *QueryAssigningReq) (*GetAssigningResp, error)
	// 查询可用资金的分配情况
	GetAssigningByRealClient(context.Context, *QueryAssigningReq) (*GetAssigningResp, error)
	// 查询资产单元的当前资金明细
	GetFundReal(context.Context, *QueryFundRealReq) (*GetFundRealResp, error)
	// 查询资产单元的资产变动流水
	ListFundJour(context.Context, *QueryFundJourReq) (*ListFundJourResp, error)
}

func RegisterFundRealManagerServer(s *grpc.Server, srv FundRealManagerServer) {
	s.RegisterService(&_FundRealManager_serviceDesc, srv)
}

func _FundRealManager_TransferIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransferReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FundRealManagerServer).TransferIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcapi.FundRealManager/TransferIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FundRealManagerServer).TransferIn(ctx, req.(*TransferReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FundRealManager_TransferOut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransferReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FundRealManagerServer).TransferOut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcapi.FundRealManager/TransferOut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FundRealManagerServer).TransferOut(ctx, req.(*TransferReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FundRealManager_Debit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransferReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FundRealManagerServer).Debit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcapi.FundRealManager/Debit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FundRealManagerServer).Debit(ctx, req.(*TransferReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FundRealManager_GetAssigning_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAssigningReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FundRealManagerServer).GetAssigning(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcapi.FundRealManager/GetAssigning",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FundRealManagerServer).GetAssigning(ctx, req.(*QueryAssigningReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FundRealManager_GetAssigningByRealClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAssigningReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FundRealManagerServer).GetAssigningByRealClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcapi.FundRealManager/GetAssigningByRealClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FundRealManagerServer).GetAssigningByRealClient(ctx, req.(*QueryAssigningReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FundRealManager_GetFundReal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryFundRealReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FundRealManagerServer).GetFundReal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcapi.FundRealManager/GetFundReal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FundRealManagerServer).GetFundReal(ctx, req.(*QueryFundRealReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FundRealManager_ListFundJour_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryFundJourReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FundRealManagerServer).ListFundJour(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcapi.FundRealManager/ListFundJour",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FundRealManagerServer).ListFundJour(ctx, req.(*QueryFundJourReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _FundRealManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpcapi.FundRealManager",
	HandlerType: (*FundRealManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TransferIn",
			Handler:    _FundRealManager_TransferIn_Handler,
		},
		{
			MethodName: "TransferOut",
			Handler:    _FundRealManager_TransferOut_Handler,
		},
		{
			MethodName: "Debit",
			Handler:    _FundRealManager_Debit_Handler,
		},
		{
			MethodName: "GetAssigning",
			Handler:    _FundRealManager_GetAssigning_Handler,
		},
		{
			MethodName: "GetAssigningByRealClient",
			Handler:    _FundRealManager_GetAssigningByRealClient_Handler,
		},
		{
			MethodName: "GetFundReal",
			Handler:    _FundRealManager_GetFundReal_Handler,
		},
		{
			MethodName: "ListFundJour",
			Handler:    _FundRealManager_ListFundJour_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fund_real.proto",
}

func init() { proto.RegisterFile("fund_real.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 902 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x56, 0xcf, 0x8f, 0x1b, 0x35,
	0x14, 0xd6, 0x6c, 0xf6, 0xe7, 0x4b, 0xba, 0x49, 0x0c, 0xdd, 0x9d, 0x46, 0x4b, 0x15, 0x59, 0x05,
	0xad, 0x2a, 0x48, 0xa4, 0xed, 0x8d, 0x5b, 0xd2, 0xd2, 0x2a, 0xa8, 0x08, 0x98, 0x56, 0x42, 0x20,
	0x01, 0x72, 0x66, 0x9c, 0xd1, 0x94, 0x89, 0x3d, 0xb5, 0x3d, 0x40, 0x38, 0x72, 0xe7, 0x84, 0x38,
	0x73, 0x84, 0xff, 0x85, 0x23, 0xff, 0x02, 0x77, 0xfe, 0x01, 0x0e, 0xc8, 0x9e, 0x19, 0x8f, 0x67,
	0xb2, 0x54, 0xda, 0x45, 0xda, 0x5e, 0x22, 0xf9, 0x7b, 0x5f, 0xbe, 0xe7, 0xe7, 0xef, 0x3d, 0x7b,
	0xa0, 0xbf, 0xca, 0x59, 0xf4, 0xb5, 0xa0, 0x24, 0x9d, 0x64, 0x82, 0x2b, 0x8e, 0x0e, 0x62, 0x91,
	0x85, 0x24, 0x4b, 0x46, 0x67, 0x31, 0xe7, 0x71, 0x4a, 0xa7, 0x24, 0x4b, 0xa6, 0x84, 0x31, 0xae,
	0x88, 0x4a, 0x38, 0x93, 0x05, 0x6d, 0xd4, 0x0b, 0xf9, 0x7a, 0xcd, 0x59, 0xb1, 0xc2, 0x3f, 0x79,
	0xd0, 0x7d, 0x2e, 0x08, 0x93, 0x2b, 0x2a, 0x02, 0xfa, 0x12, 0xdd, 0x87, 0x83, 0x25, 0x91, 0x34,
	0xa0, 0x2f, 0x7d, 0x6f, 0xec, 0x9d, 0x77, 0x2f, 0x06, 0x93, 0x52, 0x76, 0x32, 0x2f, 0xf0, 0xa0,
	0x22, 0xa0, 0x11, 0x1c, 0x86, 0x69, 0x42, 0x99, 0x5a, 0x44, 0xfe, 0xce, 0xd8, 0x3b, 0xef, 0x04,
	0x76, 0x8d, 0xce, 0xe0, 0x28, 0x13, 0xfc, 0x05, 0x0d, 0x75, 0xb0, 0x63, 0x82, 0x35, 0x80, 0x7c,
	0x9d, 0x25, 0x25, 0x2c, 0xa4, 0xfe, 0xee, 0xd8, 0x3b, 0xf7, 0x82, 0x6a, 0x89, 0x1f, 0x42, 0xaf,
	0xde, 0x8e, 0xcc, 0xd0, 0x03, 0x00, 0x41, 0x65, 0xf6, 0x4c, 0x11, 0x95, 0xcb, 0x72, 0x4b, 0x6f,
	0xd8, 0x2d, 0x05, 0x36, 0x14, 0x38, 0x34, 0xfc, 0xb7, 0x07, 0xc3, 0x4f, 0x73, 0x2a, 0x36, 0x33,
	0x29, 0x93, 0x98, 0x25, 0x2c, 0xbe, 0xb9, 0xd2, 0xce, 0xa1, 0xaf, 0x3d, 0x79, 0x9c, 0xb3, 0x68,
	0x16, 0x86, 0x3c, 0x67, 0xca, 0x94, 0x78, 0x14, 0xb4, 0x61, 0xf4, 0x0e, 0x1c, 0x6b, 0xe8, 0xa1,
	0xd1, 0x7d, 0xbe, 0xc9, 0xa8, 0xbf, 0x67, 0x88, 0x2d, 0x14, 0x61, 0xe8, 0xd5, 0xc8, 0x22, 0xf2,
	0xf7, 0x0d, 0xab, 0x81, 0xe1, 0x5f, 0x3d, 0x18, 0x3c, 0xa1, 0xca, 0xa9, 0xf7, 0x9a, 0x67, 0xa7,
	0xf7, 0x4f, 0x8c, 0x0a, 0x8d, 0xe6, 0xa5, 0x45, 0x3b, 0xc6, 0xa2, 0x36, 0x8c, 0xde, 0x85, 0x61,
	0xce, 0x66, 0x2d, 0x6e, 0xc7, 0x70, 0xb7, 0x03, 0xf8, 0x7b, 0x18, 0x18, 0x4b, 0xf4, 0x09, 0x04,
	0x94, 0xa4, 0x37, 0xe6, 0x08, 0xfe, 0x7d, 0x17, 0x0e, 0xab, 0xac, 0x4d, 0xaa, 0xd7, 0x36, 0xef,
	0x55, 0x49, 0xee, 0xc1, 0xad, 0x55, 0xed, 0x9e, 0x4d, 0xd4, 0x04, 0xb5, 0x42, 0xc2, 0x12, 0xf5,
	0x88, 0xa8, 0xa2, 0xb5, 0x3b, 0x81, 0x5d, 0xeb, 0xdc, 0x44, 0x4a, 0xea, 0x7a, 0x5d, 0x03, 0x3a,
	0xba, 0xe6, 0x8c, 0x6e, 0x4c, 0xb4, 0xf0, 0xb8, 0x06, 0x74, 0xb3, 0x84, 0xb9, 0x10, 0x94, 0xa9,
	0xea, 0xa4, 0x0f, 0xcc, 0x49, 0xb7, 0x50, 0xbd, 0x4b, 0xca, 0xc8, 0x32, 0xa5, 0x15, 0xed, 0xd0,
	0xd0, 0x9a, 0xa0, 0xa9, 0x45, 0xf0, 0x1f, 0x28, 0xab, 0x58, 0x47, 0x05, 0xab, 0x01, 0xa2, 0xbb,
	0x00, 0xcb, 0x7c, 0x53, 0x51, 0xc0, 0x50, 0x1c, 0xa4, 0x6a, 0xe0, 0x79, 0xcd, 0xe9, 0x16, 0x7b,
	0x6a, 0xa2, 0xd5, 0x48, 0x3c, 0xa3, 0x69, 0x5a, 0x11, 0x7b, 0x45, 0x4b, 0xb5, 0x60, 0x53, 0x25,
	0x17, 0x82, 0x86, 0xb6, 0xca, 0x5b, 0x65, 0x95, 0x0d, 0x54, 0x8f, 0x44, 0x44, 0x97, 0x89, 0x65,
	0x1d, 0x1b, 0x56, 0x03, 0x43, 0x27, 0xb0, 0x2f, 0xe8, 0x9a, 0x88, 0x6f, 0xfc, 0xbe, 0x39, 0xcc,
	0x72, 0xa5, 0x1d, 0xca, 0xb3, 0x88, 0x28, 0x3a, 0x53, 0xfe, 0xb0, 0x70, 0xa8, 0x5a, 0xe3, 0x35,
	0xf4, 0x9f, 0x50, 0x55, 0xb7, 0xe8, 0x75, 0x87, 0xe8, 0x6d, 0xd8, 0x8d, 0x88, 0x22, 0xa6, 0x87,
	0xba, 0x17, 0x43, 0x4b, 0xb7, 0xca, 0x26, 0x8c, 0xff, 0xf0, 0x9c, 0xa1, 0xf8, 0x90, 0xe7, 0x57,
	0xbe, 0x81, 0xdf, 0x83, 0xc3, 0x8c, 0xc4, 0x74, 0xc1, 0x56, 0x7c, 0x2b, 0xd7, 0x27, 0x65, 0x20,
	0xb0, 0x94, 0x46, 0x7b, 0x77, 0x5e, 0x35, 0x43, 0xbb, 0xed, 0xc1, 0xc0, 0xd0, 0x5b, 0xe6, 0x32,
	0x61, 0x54, 0xca, 0xc7, 0x29, 0x89, 0xcb, 0xee, 0x6d, 0x60, 0xf8, 0x1f, 0xaf, 0x98, 0x33, 0x5d,
	0xc8, 0x6b, 0x9d, 0xb3, 0x31, 0x74, 0x57, 0x09, 0xd3, 0x4d, 0xe0, 0x4c, 0x9a, 0x0b, 0x6d, 0x95,
	0xb3, 0xbf, 0x5d, 0x0e, 0x9a, 0x00, 0xe2, 0x61, 0x98, 0x8b, 0x0f, 0x1a, 0xe3, 0x54, 0x4c, 0xdd,
	0x25, 0x11, 0xfc, 0x8b, 0x07, 0x83, 0xa7, 0x89, 0x54, 0xb5, 0x97, 0xff, 0xa3, 0x7b, 0xb4, 0x65,
	0xff, 0xed, 0xa8, 0x09, 0xdb, 0x26, 0xeb, 0x8c, 0x3b, 0x5b, 0x4d, 0x66, 0x36, 0x60, 0xc2, 0x17,
	0xbf, 0xed, 0x41, 0xbf, 0xea, 0xbb, 0x8f, 0x08, 0x23, 0x31, 0x15, 0xe8, 0x33, 0x80, 0xea, 0x95,
	0x5d, 0x30, 0xf4, 0xa6, 0xfd, 0xab, 0xf3, 0x25, 0x30, 0xba, 0x7d, 0x09, 0x2a, 0x33, 0x7c, 0xf6,
	0xe3, 0x9f, 0x7f, 0xfd, 0xbc, 0x73, 0x82, 0x87, 0x53, 0x6d, 0xc8, 0x54, 0x95, 0xb1, 0x69, 0xc2,
	0xde, 0xf7, 0xee, 0xa3, 0xcf, 0xeb, 0xaf, 0x89, 0x8f, 0x73, 0x75, 0x35, 0xe5, 0xb7, 0x8c, 0xf2,
	0x29, 0x46, 0x2d, 0x65, 0x9e, 0x2b, 0x2d, 0xfd, 0x14, 0xf6, 0x1e, 0xe9, 0xf9, 0xbe, 0x9a, 0xe8,
	0x89, 0x11, 0x1d, 0xe0, 0x6e, 0x21, 0x6a, 0x6e, 0x08, 0xad, 0xf6, 0x15, 0xf4, 0xdc, 0xf7, 0x12,
	0x8d, 0xec, 0xdf, 0xb7, 0x3e, 0x1c, 0x46, 0x77, 0x6c, 0xac, 0xfd, 0xc4, 0xe2, 0x53, 0x23, 0x3f,
	0x44, 0xfd, 0x42, 0x9e, 0x58, 0xbd, 0xef, 0xc0, 0x77, 0xc9, 0xf3, 0x4d, 0x60, 0x9f, 0xeb, 0xeb,
	0xe6, 0xba, 0x67, 0x72, 0xdd, 0x45, 0x67, 0xad, 0x5c, 0xd3, 0xe5, 0x46, 0xdf, 0xa3, 0xc5, 0x00,
	0xa1, 0x2f, 0xa0, 0xeb, 0x5c, 0x61, 0xe8, 0x4e, 0x33, 0x97, 0xf3, 0xfa, 0x8e, 0x7c, 0x37, 0x95,
	0x7b, 0xe7, 0x55, 0x87, 0x86, 0x8e, 0x8b, 0x4c, 0xfa, 0x47, 0x67, 0x40, 0x5f, 0x42, 0xcf, 0xed,
	0xf0, 0xcb, 0xc4, 0xcb, 0x5b, 0xcc, 0xa9, 0xa3, 0x3d, 0x13, 0xed, 0x33, 0x7b, 0xc1, 0x73, 0x31,
	0x4d, 0x13, 0xa9, 0xe6, 0xa7, 0x70, 0x3b, 0xe4, 0xeb, 0x49, 0xbc, 0x9a, 0x7c, 0x4b, 0x26, 0xfa,
	0xaf, 0xa5, 0xc4, 0x72, 0xdf, 0x7c, 0xab, 0x3e, 0xf8, 0x37, 0x00, 0x00, 0xff, 0xff, 0x9a, 0x4c,
	0x3d, 0x88, 0xf3, 0x0a, 0x00, 0x00,
}
