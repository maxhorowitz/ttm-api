// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.20.3
// source: proto/service.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	AppService_PairwiseAnalysis_FullMethodName = "/app_service.AppService/PairwiseAnalysis"
	AppService_Backtest_FullMethodName         = "/app_service.AppService/Backtest"
)

// AppServiceClient is the client API for AppService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AppServiceClient interface {
	PairwiseAnalysis(ctx context.Context, in *PairwiseAnalysisRequest, opts ...grpc.CallOption) (*PairwiseAnalysisResponse, error)
	Backtest(ctx context.Context, in *BacktestRequest, opts ...grpc.CallOption) (*BacktestResponse, error)
}

type appServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAppServiceClient(cc grpc.ClientConnInterface) AppServiceClient {
	return &appServiceClient{cc}
}

func (c *appServiceClient) PairwiseAnalysis(ctx context.Context, in *PairwiseAnalysisRequest, opts ...grpc.CallOption) (*PairwiseAnalysisResponse, error) {
	out := new(PairwiseAnalysisResponse)
	err := c.cc.Invoke(ctx, AppService_PairwiseAnalysis_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appServiceClient) Backtest(ctx context.Context, in *BacktestRequest, opts ...grpc.CallOption) (*BacktestResponse, error) {
	out := new(BacktestResponse)
	err := c.cc.Invoke(ctx, AppService_Backtest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppServiceServer is the server API for AppService service.
// All implementations must embed UnimplementedAppServiceServer
// for forward compatibility
type AppServiceServer interface {
	PairwiseAnalysis(context.Context, *PairwiseAnalysisRequest) (*PairwiseAnalysisResponse, error)
	Backtest(context.Context, *BacktestRequest) (*BacktestResponse, error)
	mustEmbedUnimplementedAppServiceServer()
}

// UnimplementedAppServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAppServiceServer struct {
}

func (UnimplementedAppServiceServer) PairwiseAnalysis(context.Context, *PairwiseAnalysisRequest) (*PairwiseAnalysisResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PairwiseAnalysis not implemented")
}
func (UnimplementedAppServiceServer) Backtest(context.Context, *BacktestRequest) (*BacktestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Backtest not implemented")
}
func (UnimplementedAppServiceServer) mustEmbedUnimplementedAppServiceServer() {}

// UnsafeAppServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AppServiceServer will
// result in compilation errors.
type UnsafeAppServiceServer interface {
	mustEmbedUnimplementedAppServiceServer()
}

func RegisterAppServiceServer(s grpc.ServiceRegistrar, srv AppServiceServer) {
	s.RegisterService(&AppService_ServiceDesc, srv)
}

func _AppService_PairwiseAnalysis_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PairwiseAnalysisRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).PairwiseAnalysis(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppService_PairwiseAnalysis_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).PairwiseAnalysis(ctx, req.(*PairwiseAnalysisRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppService_Backtest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BacktestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).Backtest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppService_Backtest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).Backtest(ctx, req.(*BacktestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AppService_ServiceDesc is the grpc.ServiceDesc for AppService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AppService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "app_service.AppService",
	HandlerType: (*AppServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PairwiseAnalysis",
			Handler:    _AppService_PairwiseAnalysis_Handler,
		},
		{
			MethodName: "Backtest",
			Handler:    _AppService_Backtest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/service.proto",
}
