// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: pkg/apis/dfdaemon/v2/dfdaemon.proto

package dfdaemon

import (
	context "context"
	v2 "d7y.io/api/v2/pkg/apis/common/v2"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DfdaemonClient is the client API for Dfdaemon service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DfdaemonClient interface {
	// GetPieceNumbers gets piece numbers from the other peer.
	GetPieceNumbers(ctx context.Context, in *GetPieceNumbersRequest, opts ...grpc.CallOption) (*GetPieceNumbersResponse, error)
	// SyncPieces syncs pieces from the other peer.
	SyncPieces(ctx context.Context, opts ...grpc.CallOption) (Dfdaemon_SyncPiecesClient, error)
	// DownloadTask downloads task back-to-source.
	DownloadTask(ctx context.Context, in *DownloadTaskRequest, opts ...grpc.CallOption) (Dfdaemon_DownloadTaskClient, error)
	// UploadTask uploads task to p2p network.
	UploadTask(ctx context.Context, in *UploadTaskRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// StatTask stats task information.
	StatTask(ctx context.Context, in *StatTaskRequest, opts ...grpc.CallOption) (*v2.Task, error)
	// DeleteTask deletes task from p2p network.
	DeleteTask(ctx context.Context, in *DeleteTaskRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type dfdaemonClient struct {
	cc grpc.ClientConnInterface
}

func NewDfdaemonClient(cc grpc.ClientConnInterface) DfdaemonClient {
	return &dfdaemonClient{cc}
}

func (c *dfdaemonClient) GetPieceNumbers(ctx context.Context, in *GetPieceNumbersRequest, opts ...grpc.CallOption) (*GetPieceNumbersResponse, error) {
	out := new(GetPieceNumbersResponse)
	err := c.cc.Invoke(ctx, "/dfdaemon.v2.Dfdaemon/GetPieceNumbers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dfdaemonClient) SyncPieces(ctx context.Context, opts ...grpc.CallOption) (Dfdaemon_SyncPiecesClient, error) {
	stream, err := c.cc.NewStream(ctx, &Dfdaemon_ServiceDesc.Streams[0], "/dfdaemon.v2.Dfdaemon/SyncPieces", opts...)
	if err != nil {
		return nil, err
	}
	x := &dfdaemonSyncPiecesClient{stream}
	return x, nil
}

type Dfdaemon_SyncPiecesClient interface {
	Send(*SyncPiecesRequest) error
	Recv() (*SyncPiecesResponse, error)
	grpc.ClientStream
}

type dfdaemonSyncPiecesClient struct {
	grpc.ClientStream
}

func (x *dfdaemonSyncPiecesClient) Send(m *SyncPiecesRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dfdaemonSyncPiecesClient) Recv() (*SyncPiecesResponse, error) {
	m := new(SyncPiecesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dfdaemonClient) DownloadTask(ctx context.Context, in *DownloadTaskRequest, opts ...grpc.CallOption) (Dfdaemon_DownloadTaskClient, error) {
	stream, err := c.cc.NewStream(ctx, &Dfdaemon_ServiceDesc.Streams[1], "/dfdaemon.v2.Dfdaemon/DownloadTask", opts...)
	if err != nil {
		return nil, err
	}
	x := &dfdaemonDownloadTaskClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Dfdaemon_DownloadTaskClient interface {
	Recv() (*DownloadTaskResponse, error)
	grpc.ClientStream
}

type dfdaemonDownloadTaskClient struct {
	grpc.ClientStream
}

func (x *dfdaemonDownloadTaskClient) Recv() (*DownloadTaskResponse, error) {
	m := new(DownloadTaskResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dfdaemonClient) UploadTask(ctx context.Context, in *UploadTaskRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/dfdaemon.v2.Dfdaemon/UploadTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dfdaemonClient) StatTask(ctx context.Context, in *StatTaskRequest, opts ...grpc.CallOption) (*v2.Task, error) {
	out := new(v2.Task)
	err := c.cc.Invoke(ctx, "/dfdaemon.v2.Dfdaemon/StatTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dfdaemonClient) DeleteTask(ctx context.Context, in *DeleteTaskRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/dfdaemon.v2.Dfdaemon/DeleteTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DfdaemonServer is the server API for Dfdaemon service.
// All implementations should embed UnimplementedDfdaemonServer
// for forward compatibility
type DfdaemonServer interface {
	// GetPieceNumbers gets piece numbers from the other peer.
	GetPieceNumbers(context.Context, *GetPieceNumbersRequest) (*GetPieceNumbersResponse, error)
	// SyncPieces syncs pieces from the other peer.
	SyncPieces(Dfdaemon_SyncPiecesServer) error
	// DownloadTask downloads task back-to-source.
	DownloadTask(*DownloadTaskRequest, Dfdaemon_DownloadTaskServer) error
	// UploadTask uploads task to p2p network.
	UploadTask(context.Context, *UploadTaskRequest) (*emptypb.Empty, error)
	// StatTask stats task information.
	StatTask(context.Context, *StatTaskRequest) (*v2.Task, error)
	// DeleteTask deletes task from p2p network.
	DeleteTask(context.Context, *DeleteTaskRequest) (*emptypb.Empty, error)
}

// UnimplementedDfdaemonServer should be embedded to have forward compatible implementations.
type UnimplementedDfdaemonServer struct {
}

func (UnimplementedDfdaemonServer) GetPieceNumbers(context.Context, *GetPieceNumbersRequest) (*GetPieceNumbersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPieceNumbers not implemented")
}
func (UnimplementedDfdaemonServer) SyncPieces(Dfdaemon_SyncPiecesServer) error {
	return status.Errorf(codes.Unimplemented, "method SyncPieces not implemented")
}
func (UnimplementedDfdaemonServer) DownloadTask(*DownloadTaskRequest, Dfdaemon_DownloadTaskServer) error {
	return status.Errorf(codes.Unimplemented, "method DownloadTask not implemented")
}
func (UnimplementedDfdaemonServer) UploadTask(context.Context, *UploadTaskRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadTask not implemented")
}
func (UnimplementedDfdaemonServer) StatTask(context.Context, *StatTaskRequest) (*v2.Task, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StatTask not implemented")
}
func (UnimplementedDfdaemonServer) DeleteTask(context.Context, *DeleteTaskRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTask not implemented")
}

// UnsafeDfdaemonServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DfdaemonServer will
// result in compilation errors.
type UnsafeDfdaemonServer interface {
	mustEmbedUnimplementedDfdaemonServer()
}

func RegisterDfdaemonServer(s grpc.ServiceRegistrar, srv DfdaemonServer) {
	s.RegisterService(&Dfdaemon_ServiceDesc, srv)
}

func _Dfdaemon_GetPieceNumbers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPieceNumbersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DfdaemonServer).GetPieceNumbers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dfdaemon.v2.Dfdaemon/GetPieceNumbers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DfdaemonServer).GetPieceNumbers(ctx, req.(*GetPieceNumbersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dfdaemon_SyncPieces_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DfdaemonServer).SyncPieces(&dfdaemonSyncPiecesServer{stream})
}

type Dfdaemon_SyncPiecesServer interface {
	Send(*SyncPiecesResponse) error
	Recv() (*SyncPiecesRequest, error)
	grpc.ServerStream
}

type dfdaemonSyncPiecesServer struct {
	grpc.ServerStream
}

func (x *dfdaemonSyncPiecesServer) Send(m *SyncPiecesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dfdaemonSyncPiecesServer) Recv() (*SyncPiecesRequest, error) {
	m := new(SyncPiecesRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Dfdaemon_DownloadTask_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DownloadTaskRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DfdaemonServer).DownloadTask(m, &dfdaemonDownloadTaskServer{stream})
}

type Dfdaemon_DownloadTaskServer interface {
	Send(*DownloadTaskResponse) error
	grpc.ServerStream
}

type dfdaemonDownloadTaskServer struct {
	grpc.ServerStream
}

func (x *dfdaemonDownloadTaskServer) Send(m *DownloadTaskResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Dfdaemon_UploadTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DfdaemonServer).UploadTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dfdaemon.v2.Dfdaemon/UploadTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DfdaemonServer).UploadTask(ctx, req.(*UploadTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dfdaemon_StatTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DfdaemonServer).StatTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dfdaemon.v2.Dfdaemon/StatTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DfdaemonServer).StatTask(ctx, req.(*StatTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dfdaemon_DeleteTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DfdaemonServer).DeleteTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dfdaemon.v2.Dfdaemon/DeleteTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DfdaemonServer).DeleteTask(ctx, req.(*DeleteTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Dfdaemon_ServiceDesc is the grpc.ServiceDesc for Dfdaemon service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Dfdaemon_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dfdaemon.v2.Dfdaemon",
	HandlerType: (*DfdaemonServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPieceNumbers",
			Handler:    _Dfdaemon_GetPieceNumbers_Handler,
		},
		{
			MethodName: "UploadTask",
			Handler:    _Dfdaemon_UploadTask_Handler,
		},
		{
			MethodName: "StatTask",
			Handler:    _Dfdaemon_StatTask_Handler,
		},
		{
			MethodName: "DeleteTask",
			Handler:    _Dfdaemon_DeleteTask_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SyncPieces",
			Handler:       _Dfdaemon_SyncPieces_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "DownloadTask",
			Handler:       _Dfdaemon_DownloadTask_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pkg/apis/dfdaemon/v2/dfdaemon.proto",
}
