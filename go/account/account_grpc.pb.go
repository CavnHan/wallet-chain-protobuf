// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.0--rc1
// source: proto/account.proto

package account

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	WalletAccountService_GetSupportChains_FullMethodName        = "/proto.WalletAccountService/getSupportChains"
	WalletAccountService_ConvertAddress_FullMethodName          = "/proto.WalletAccountService/convertAddress"
	WalletAccountService_ValidAddress_FullMethodName            = "/proto.WalletAccountService/validAddress"
	WalletAccountService_GetBlockByNumber_FullMethodName        = "/proto.WalletAccountService/getBlockByNumber"
	WalletAccountService_GetBlockByHash_FullMethodName          = "/proto.WalletAccountService/getBlockByHash"
	WalletAccountService_GetBlockHeaderByHash_FullMethodName    = "/proto.WalletAccountService/getBlockHeaderByHash"
	WalletAccountService_GetBlockHeaderByNumber_FullMethodName  = "/proto.WalletAccountService/getBlockHeaderByNumber"
	WalletAccountService_GetAccount_FullMethodName              = "/proto.WalletAccountService/getAccount"
	WalletAccountService_GetFee_FullMethodName                  = "/proto.WalletAccountService/getFee"
	WalletAccountService_SendTx_FullMethodName                  = "/proto.WalletAccountService/SendTx"
	WalletAccountService_GetTxByAddress_FullMethodName          = "/proto.WalletAccountService/getTxByAddress"
	WalletAccountService_GetTxByHash_FullMethodName             = "/proto.WalletAccountService/getTxByHash"
	WalletAccountService_GetBlockByRange_FullMethodName         = "/proto.WalletAccountService/getBlockByRange"
	WalletAccountService_CreateUnSignTransaction_FullMethodName = "/proto.WalletAccountService/createUnSignTransaction"
	WalletAccountService_BuildSignedTransaction_FullMethodName  = "/proto.WalletAccountService/buildSignedTransaction"
	WalletAccountService_DecodeTransaction_FullMethodName       = "/proto.WalletAccountService/decodeTransaction"
	WalletAccountService_VerifySignedTransaction_FullMethodName = "/proto.WalletAccountService/verifySignedTransaction"
	WalletAccountService_GetExtraData_FullMethodName            = "/proto.WalletAccountService/getExtraData"
)

// WalletAccountServiceClient is the client API for WalletAccountService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WalletAccountServiceClient interface {
	GetSupportChains(ctx context.Context, in *SupportChainsRequest, opts ...grpc.CallOption) (*SupportChainsResponse, error)
	ConvertAddress(ctx context.Context, in *ConvertAddressRequest, opts ...grpc.CallOption) (*ConvertAddressResponse, error)
	ValidAddress(ctx context.Context, in *ValidAddressRequest, opts ...grpc.CallOption) (*ValidAddressResponse, error)
	GetBlockByNumber(ctx context.Context, in *BlockNumberRequest, opts ...grpc.CallOption) (*BlockResponse, error)
	GetBlockByHash(ctx context.Context, in *BlockHashRequest, opts ...grpc.CallOption) (*BlockResponse, error)
	GetBlockHeaderByHash(ctx context.Context, in *BlockHeaderHashRequest, opts ...grpc.CallOption) (*BlockHeaderResponse, error)
	GetBlockHeaderByNumber(ctx context.Context, in *BlockHeaderNumberRequest, opts ...grpc.CallOption) (*BlockHeaderResponse, error)
	GetAccount(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*AccountResponse, error)
	GetFee(ctx context.Context, in *FeeRequest, opts ...grpc.CallOption) (*FeeResponse, error)
	SendTx(ctx context.Context, in *SendTxRequest, opts ...grpc.CallOption) (*SendTxResponse, error)
	GetTxByAddress(ctx context.Context, in *TxAddressRequest, opts ...grpc.CallOption) (*TxAddressResponse, error)
	GetTxByHash(ctx context.Context, in *TxHashRequest, opts ...grpc.CallOption) (*TxHashResponse, error)
	GetBlockByRange(ctx context.Context, in *BlockByRangeRequest, opts ...grpc.CallOption) (*BlockByRangeResponse, error)
	CreateUnSignTransaction(ctx context.Context, in *UnSignTransactionRequest, opts ...grpc.CallOption) (*UnSignTransactionResponse, error)
	BuildSignedTransaction(ctx context.Context, in *SignedTransactionRequest, opts ...grpc.CallOption) (*SignedTransactionResponse, error)
	DecodeTransaction(ctx context.Context, in *DecodeTransactionRequest, opts ...grpc.CallOption) (*DecodeTransactionResponse, error)
	VerifySignedTransaction(ctx context.Context, in *VerifyTransactionRequest, opts ...grpc.CallOption) (*VerifyTransactionResponse, error)
	GetExtraData(ctx context.Context, in *ExtraDataRequest, opts ...grpc.CallOption) (*ExtraDataResponse, error)
}

type walletAccountServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWalletAccountServiceClient(cc grpc.ClientConnInterface) WalletAccountServiceClient {
	return &walletAccountServiceClient{cc}
}

func (c *walletAccountServiceClient) GetSupportChains(ctx context.Context, in *SupportChainsRequest, opts ...grpc.CallOption) (*SupportChainsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SupportChainsResponse)
	err := c.cc.Invoke(ctx, WalletAccountService_GetSupportChains_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletAccountServiceClient) ConvertAddress(ctx context.Context, in *ConvertAddressRequest, opts ...grpc.CallOption) (*ConvertAddressResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ConvertAddressResponse)
	err := c.cc.Invoke(ctx, WalletAccountService_ConvertAddress_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletAccountServiceClient) ValidAddress(ctx context.Context, in *ValidAddressRequest, opts ...grpc.CallOption) (*ValidAddressResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ValidAddressResponse)
	err := c.cc.Invoke(ctx, WalletAccountService_ValidAddress_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletAccountServiceClient) GetBlockByNumber(ctx context.Context, in *BlockNumberRequest, opts ...grpc.CallOption) (*BlockResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BlockResponse)
	err := c.cc.Invoke(ctx, WalletAccountService_GetBlockByNumber_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletAccountServiceClient) GetBlockByHash(ctx context.Context, in *BlockHashRequest, opts ...grpc.CallOption) (*BlockResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BlockResponse)
	err := c.cc.Invoke(ctx, WalletAccountService_GetBlockByHash_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletAccountServiceClient) GetBlockHeaderByHash(ctx context.Context, in *BlockHeaderHashRequest, opts ...grpc.CallOption) (*BlockHeaderResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BlockHeaderResponse)
	err := c.cc.Invoke(ctx, WalletAccountService_GetBlockHeaderByHash_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletAccountServiceClient) GetBlockHeaderByNumber(ctx context.Context, in *BlockHeaderNumberRequest, opts ...grpc.CallOption) (*BlockHeaderResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BlockHeaderResponse)
	err := c.cc.Invoke(ctx, WalletAccountService_GetBlockHeaderByNumber_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletAccountServiceClient) GetAccount(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*AccountResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AccountResponse)
	err := c.cc.Invoke(ctx, WalletAccountService_GetAccount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletAccountServiceClient) GetFee(ctx context.Context, in *FeeRequest, opts ...grpc.CallOption) (*FeeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FeeResponse)
	err := c.cc.Invoke(ctx, WalletAccountService_GetFee_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletAccountServiceClient) SendTx(ctx context.Context, in *SendTxRequest, opts ...grpc.CallOption) (*SendTxResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SendTxResponse)
	err := c.cc.Invoke(ctx, WalletAccountService_SendTx_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletAccountServiceClient) GetTxByAddress(ctx context.Context, in *TxAddressRequest, opts ...grpc.CallOption) (*TxAddressResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TxAddressResponse)
	err := c.cc.Invoke(ctx, WalletAccountService_GetTxByAddress_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletAccountServiceClient) GetTxByHash(ctx context.Context, in *TxHashRequest, opts ...grpc.CallOption) (*TxHashResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TxHashResponse)
	err := c.cc.Invoke(ctx, WalletAccountService_GetTxByHash_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletAccountServiceClient) GetBlockByRange(ctx context.Context, in *BlockByRangeRequest, opts ...grpc.CallOption) (*BlockByRangeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BlockByRangeResponse)
	err := c.cc.Invoke(ctx, WalletAccountService_GetBlockByRange_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletAccountServiceClient) CreateUnSignTransaction(ctx context.Context, in *UnSignTransactionRequest, opts ...grpc.CallOption) (*UnSignTransactionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UnSignTransactionResponse)
	err := c.cc.Invoke(ctx, WalletAccountService_CreateUnSignTransaction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletAccountServiceClient) BuildSignedTransaction(ctx context.Context, in *SignedTransactionRequest, opts ...grpc.CallOption) (*SignedTransactionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SignedTransactionResponse)
	err := c.cc.Invoke(ctx, WalletAccountService_BuildSignedTransaction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletAccountServiceClient) DecodeTransaction(ctx context.Context, in *DecodeTransactionRequest, opts ...grpc.CallOption) (*DecodeTransactionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DecodeTransactionResponse)
	err := c.cc.Invoke(ctx, WalletAccountService_DecodeTransaction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletAccountServiceClient) VerifySignedTransaction(ctx context.Context, in *VerifyTransactionRequest, opts ...grpc.CallOption) (*VerifyTransactionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VerifyTransactionResponse)
	err := c.cc.Invoke(ctx, WalletAccountService_VerifySignedTransaction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletAccountServiceClient) GetExtraData(ctx context.Context, in *ExtraDataRequest, opts ...grpc.CallOption) (*ExtraDataResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ExtraDataResponse)
	err := c.cc.Invoke(ctx, WalletAccountService_GetExtraData_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WalletAccountServiceServer is the server API for WalletAccountService service.
// All implementations should embed UnimplementedWalletAccountServiceServer
// for forward compatibility.
type WalletAccountServiceServer interface {
	GetSupportChains(context.Context, *SupportChainsRequest) (*SupportChainsResponse, error)
	ConvertAddress(context.Context, *ConvertAddressRequest) (*ConvertAddressResponse, error)
	ValidAddress(context.Context, *ValidAddressRequest) (*ValidAddressResponse, error)
	GetBlockByNumber(context.Context, *BlockNumberRequest) (*BlockResponse, error)
	GetBlockByHash(context.Context, *BlockHashRequest) (*BlockResponse, error)
	GetBlockHeaderByHash(context.Context, *BlockHeaderHashRequest) (*BlockHeaderResponse, error)
	GetBlockHeaderByNumber(context.Context, *BlockHeaderNumberRequest) (*BlockHeaderResponse, error)
	GetAccount(context.Context, *AccountRequest) (*AccountResponse, error)
	GetFee(context.Context, *FeeRequest) (*FeeResponse, error)
	SendTx(context.Context, *SendTxRequest) (*SendTxResponse, error)
	GetTxByAddress(context.Context, *TxAddressRequest) (*TxAddressResponse, error)
	GetTxByHash(context.Context, *TxHashRequest) (*TxHashResponse, error)
	GetBlockByRange(context.Context, *BlockByRangeRequest) (*BlockByRangeResponse, error)
	CreateUnSignTransaction(context.Context, *UnSignTransactionRequest) (*UnSignTransactionResponse, error)
	BuildSignedTransaction(context.Context, *SignedTransactionRequest) (*SignedTransactionResponse, error)
	DecodeTransaction(context.Context, *DecodeTransactionRequest) (*DecodeTransactionResponse, error)
	VerifySignedTransaction(context.Context, *VerifyTransactionRequest) (*VerifyTransactionResponse, error)
	GetExtraData(context.Context, *ExtraDataRequest) (*ExtraDataResponse, error)
}

// UnimplementedWalletAccountServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedWalletAccountServiceServer struct{}

func (UnimplementedWalletAccountServiceServer) GetSupportChains(context.Context, *SupportChainsRequest) (*SupportChainsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSupportChains not implemented")
}
func (UnimplementedWalletAccountServiceServer) ConvertAddress(context.Context, *ConvertAddressRequest) (*ConvertAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConvertAddress not implemented")
}
func (UnimplementedWalletAccountServiceServer) ValidAddress(context.Context, *ValidAddressRequest) (*ValidAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidAddress not implemented")
}
func (UnimplementedWalletAccountServiceServer) GetBlockByNumber(context.Context, *BlockNumberRequest) (*BlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlockByNumber not implemented")
}
func (UnimplementedWalletAccountServiceServer) GetBlockByHash(context.Context, *BlockHashRequest) (*BlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlockByHash not implemented")
}
func (UnimplementedWalletAccountServiceServer) GetBlockHeaderByHash(context.Context, *BlockHeaderHashRequest) (*BlockHeaderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlockHeaderByHash not implemented")
}
func (UnimplementedWalletAccountServiceServer) GetBlockHeaderByNumber(context.Context, *BlockHeaderNumberRequest) (*BlockHeaderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlockHeaderByNumber not implemented")
}
func (UnimplementedWalletAccountServiceServer) GetAccount(context.Context, *AccountRequest) (*AccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccount not implemented")
}
func (UnimplementedWalletAccountServiceServer) GetFee(context.Context, *FeeRequest) (*FeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFee not implemented")
}
func (UnimplementedWalletAccountServiceServer) SendTx(context.Context, *SendTxRequest) (*SendTxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendTx not implemented")
}
func (UnimplementedWalletAccountServiceServer) GetTxByAddress(context.Context, *TxAddressRequest) (*TxAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTxByAddress not implemented")
}
func (UnimplementedWalletAccountServiceServer) GetTxByHash(context.Context, *TxHashRequest) (*TxHashResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTxByHash not implemented")
}
func (UnimplementedWalletAccountServiceServer) GetBlockByRange(context.Context, *BlockByRangeRequest) (*BlockByRangeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlockByRange not implemented")
}
func (UnimplementedWalletAccountServiceServer) CreateUnSignTransaction(context.Context, *UnSignTransactionRequest) (*UnSignTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUnSignTransaction not implemented")
}
func (UnimplementedWalletAccountServiceServer) BuildSignedTransaction(context.Context, *SignedTransactionRequest) (*SignedTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BuildSignedTransaction not implemented")
}
func (UnimplementedWalletAccountServiceServer) DecodeTransaction(context.Context, *DecodeTransactionRequest) (*DecodeTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DecodeTransaction not implemented")
}
func (UnimplementedWalletAccountServiceServer) VerifySignedTransaction(context.Context, *VerifyTransactionRequest) (*VerifyTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifySignedTransaction not implemented")
}
func (UnimplementedWalletAccountServiceServer) GetExtraData(context.Context, *ExtraDataRequest) (*ExtraDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExtraData not implemented")
}
func (UnimplementedWalletAccountServiceServer) testEmbeddedByValue() {}

// UnsafeWalletAccountServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WalletAccountServiceServer will
// result in compilation errors.
type UnsafeWalletAccountServiceServer interface {
	mustEmbedUnimplementedWalletAccountServiceServer()
}

func RegisterWalletAccountServiceServer(s grpc.ServiceRegistrar, srv WalletAccountServiceServer) {
	// If the following call pancis, it indicates UnimplementedWalletAccountServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&WalletAccountService_ServiceDesc, srv)
}

func _WalletAccountService_GetSupportChains_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SupportChainsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletAccountServiceServer).GetSupportChains(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletAccountService_GetSupportChains_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletAccountServiceServer).GetSupportChains(ctx, req.(*SupportChainsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletAccountService_ConvertAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConvertAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletAccountServiceServer).ConvertAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletAccountService_ConvertAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletAccountServiceServer).ConvertAddress(ctx, req.(*ConvertAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletAccountService_ValidAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletAccountServiceServer).ValidAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletAccountService_ValidAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletAccountServiceServer).ValidAddress(ctx, req.(*ValidAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletAccountService_GetBlockByNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockNumberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletAccountServiceServer).GetBlockByNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletAccountService_GetBlockByNumber_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletAccountServiceServer).GetBlockByNumber(ctx, req.(*BlockNumberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletAccountService_GetBlockByHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockHashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletAccountServiceServer).GetBlockByHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletAccountService_GetBlockByHash_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletAccountServiceServer).GetBlockByHash(ctx, req.(*BlockHashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletAccountService_GetBlockHeaderByHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockHeaderHashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletAccountServiceServer).GetBlockHeaderByHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletAccountService_GetBlockHeaderByHash_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletAccountServiceServer).GetBlockHeaderByHash(ctx, req.(*BlockHeaderHashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletAccountService_GetBlockHeaderByNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockHeaderNumberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletAccountServiceServer).GetBlockHeaderByNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletAccountService_GetBlockHeaderByNumber_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletAccountServiceServer).GetBlockHeaderByNumber(ctx, req.(*BlockHeaderNumberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletAccountService_GetAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletAccountServiceServer).GetAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletAccountService_GetAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletAccountServiceServer).GetAccount(ctx, req.(*AccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletAccountService_GetFee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletAccountServiceServer).GetFee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletAccountService_GetFee_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletAccountServiceServer).GetFee(ctx, req.(*FeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletAccountService_SendTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendTxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletAccountServiceServer).SendTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletAccountService_SendTx_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletAccountServiceServer).SendTx(ctx, req.(*SendTxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletAccountService_GetTxByAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TxAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletAccountServiceServer).GetTxByAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletAccountService_GetTxByAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletAccountServiceServer).GetTxByAddress(ctx, req.(*TxAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletAccountService_GetTxByHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TxHashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletAccountServiceServer).GetTxByHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletAccountService_GetTxByHash_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletAccountServiceServer).GetTxByHash(ctx, req.(*TxHashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletAccountService_GetBlockByRange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockByRangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletAccountServiceServer).GetBlockByRange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletAccountService_GetBlockByRange_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletAccountServiceServer).GetBlockByRange(ctx, req.(*BlockByRangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletAccountService_CreateUnSignTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnSignTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletAccountServiceServer).CreateUnSignTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletAccountService_CreateUnSignTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletAccountServiceServer).CreateUnSignTransaction(ctx, req.(*UnSignTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletAccountService_BuildSignedTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignedTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletAccountServiceServer).BuildSignedTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletAccountService_BuildSignedTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletAccountServiceServer).BuildSignedTransaction(ctx, req.(*SignedTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletAccountService_DecodeTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DecodeTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletAccountServiceServer).DecodeTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletAccountService_DecodeTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletAccountServiceServer).DecodeTransaction(ctx, req.(*DecodeTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletAccountService_VerifySignedTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletAccountServiceServer).VerifySignedTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletAccountService_VerifySignedTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletAccountServiceServer).VerifySignedTransaction(ctx, req.(*VerifyTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletAccountService_GetExtraData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExtraDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletAccountServiceServer).GetExtraData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletAccountService_GetExtraData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletAccountServiceServer).GetExtraData(ctx, req.(*ExtraDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WalletAccountService_ServiceDesc is the grpc.ServiceDesc for WalletAccountService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WalletAccountService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.WalletAccountService",
	HandlerType: (*WalletAccountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getSupportChains",
			Handler:    _WalletAccountService_GetSupportChains_Handler,
		},
		{
			MethodName: "convertAddress",
			Handler:    _WalletAccountService_ConvertAddress_Handler,
		},
		{
			MethodName: "validAddress",
			Handler:    _WalletAccountService_ValidAddress_Handler,
		},
		{
			MethodName: "getBlockByNumber",
			Handler:    _WalletAccountService_GetBlockByNumber_Handler,
		},
		{
			MethodName: "getBlockByHash",
			Handler:    _WalletAccountService_GetBlockByHash_Handler,
		},
		{
			MethodName: "getBlockHeaderByHash",
			Handler:    _WalletAccountService_GetBlockHeaderByHash_Handler,
		},
		{
			MethodName: "getBlockHeaderByNumber",
			Handler:    _WalletAccountService_GetBlockHeaderByNumber_Handler,
		},
		{
			MethodName: "getAccount",
			Handler:    _WalletAccountService_GetAccount_Handler,
		},
		{
			MethodName: "getFee",
			Handler:    _WalletAccountService_GetFee_Handler,
		},
		{
			MethodName: "SendTx",
			Handler:    _WalletAccountService_SendTx_Handler,
		},
		{
			MethodName: "getTxByAddress",
			Handler:    _WalletAccountService_GetTxByAddress_Handler,
		},
		{
			MethodName: "getTxByHash",
			Handler:    _WalletAccountService_GetTxByHash_Handler,
		},
		{
			MethodName: "getBlockByRange",
			Handler:    _WalletAccountService_GetBlockByRange_Handler,
		},
		{
			MethodName: "createUnSignTransaction",
			Handler:    _WalletAccountService_CreateUnSignTransaction_Handler,
		},
		{
			MethodName: "buildSignedTransaction",
			Handler:    _WalletAccountService_BuildSignedTransaction_Handler,
		},
		{
			MethodName: "decodeTransaction",
			Handler:    _WalletAccountService_DecodeTransaction_Handler,
		},
		{
			MethodName: "verifySignedTransaction",
			Handler:    _WalletAccountService_VerifySignedTransaction_Handler,
		},
		{
			MethodName: "getExtraData",
			Handler:    _WalletAccountService_GetExtraData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/account.proto",
}
