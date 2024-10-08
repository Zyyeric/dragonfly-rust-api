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

// DfdaemonUploadClient is the client API for DfdaemonUpload service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DfdaemonUploadClient interface {
	// DownloadTask downloads task from p2p network.
	DownloadTask(ctx context.Context, in *DownloadTaskRequest, opts ...grpc.CallOption) (DfdaemonUpload_DownloadTaskClient, error)
	// StatTask stats task information.
	StatTask(ctx context.Context, in *StatTaskRequest, opts ...grpc.CallOption) (*v2.Task, error)
	// DeleteTask deletes task from p2p network.
	DeleteTask(ctx context.Context, in *DeleteTaskRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// SyncPieces syncs piece metadatas from remote peer.
	SyncPieces(ctx context.Context, in *SyncPiecesRequest, opts ...grpc.CallOption) (DfdaemonUpload_SyncPiecesClient, error)
	// DownloadPiece downloads piece from the remote peer.
	DownloadPiece(ctx context.Context, in *DownloadPieceRequest, opts ...grpc.CallOption) (*DownloadPieceResponse, error)
	// DownloadCacheTask downloads cache task from p2p network.
	DownloadCacheTask(ctx context.Context, in *DownloadCacheTaskRequest, opts ...grpc.CallOption) (DfdaemonUpload_DownloadCacheTaskClient, error)
	// StatCacheTask stats cache task information.
	StatCacheTask(ctx context.Context, in *StatCacheTaskRequest, opts ...grpc.CallOption) (*v2.CacheTask, error)
	// DeleteCacheTask deletes cache task from p2p network.
	DeleteCacheTask(ctx context.Context, in *DeleteCacheTaskRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type dfdaemonUploadClient struct {
	cc grpc.ClientConnInterface
}

func NewDfdaemonUploadClient(cc grpc.ClientConnInterface) DfdaemonUploadClient {
	return &dfdaemonUploadClient{cc}
}

func (c *dfdaemonUploadClient) DownloadTask(ctx context.Context, in *DownloadTaskRequest, opts ...grpc.CallOption) (DfdaemonUpload_DownloadTaskClient, error) {
	stream, err := c.cc.NewStream(ctx, &DfdaemonUpload_ServiceDesc.Streams[0], "/dfdaemon.v2.DfdaemonUpload/DownloadTask", opts...)
	if err != nil {
		return nil, err
	}
	x := &dfdaemonUploadDownloadTaskClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DfdaemonUpload_DownloadTaskClient interface {
	Recv() (*DownloadTaskResponse, error)
	grpc.ClientStream
}

type dfdaemonUploadDownloadTaskClient struct {
	grpc.ClientStream
}

func (x *dfdaemonUploadDownloadTaskClient) Recv() (*DownloadTaskResponse, error) {
	m := new(DownloadTaskResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dfdaemonUploadClient) StatTask(ctx context.Context, in *StatTaskRequest, opts ...grpc.CallOption) (*v2.Task, error) {
	out := new(v2.Task)
	err := c.cc.Invoke(ctx, "/dfdaemon.v2.DfdaemonUpload/StatTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dfdaemonUploadClient) DeleteTask(ctx context.Context, in *DeleteTaskRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/dfdaemon.v2.DfdaemonUpload/DeleteTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dfdaemonUploadClient) SyncPieces(ctx context.Context, in *SyncPiecesRequest, opts ...grpc.CallOption) (DfdaemonUpload_SyncPiecesClient, error) {
	stream, err := c.cc.NewStream(ctx, &DfdaemonUpload_ServiceDesc.Streams[1], "/dfdaemon.v2.DfdaemonUpload/SyncPieces", opts...)
	if err != nil {
		return nil, err
	}
	x := &dfdaemonUploadSyncPiecesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DfdaemonUpload_SyncPiecesClient interface {
	Recv() (*SyncPiecesResponse, error)
	grpc.ClientStream
}

type dfdaemonUploadSyncPiecesClient struct {
	grpc.ClientStream
}

func (x *dfdaemonUploadSyncPiecesClient) Recv() (*SyncPiecesResponse, error) {
	m := new(SyncPiecesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dfdaemonUploadClient) DownloadPiece(ctx context.Context, in *DownloadPieceRequest, opts ...grpc.CallOption) (*DownloadPieceResponse, error) {
	out := new(DownloadPieceResponse)
	err := c.cc.Invoke(ctx, "/dfdaemon.v2.DfdaemonUpload/DownloadPiece", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dfdaemonUploadClient) DownloadCacheTask(ctx context.Context, in *DownloadCacheTaskRequest, opts ...grpc.CallOption) (DfdaemonUpload_DownloadCacheTaskClient, error) {
	stream, err := c.cc.NewStream(ctx, &DfdaemonUpload_ServiceDesc.Streams[2], "/dfdaemon.v2.DfdaemonUpload/DownloadCacheTask", opts...)
	if err != nil {
		return nil, err
	}
	x := &dfdaemonUploadDownloadCacheTaskClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DfdaemonUpload_DownloadCacheTaskClient interface {
	Recv() (*DownloadCacheTaskResponse, error)
	grpc.ClientStream
}

type dfdaemonUploadDownloadCacheTaskClient struct {
	grpc.ClientStream
}

func (x *dfdaemonUploadDownloadCacheTaskClient) Recv() (*DownloadCacheTaskResponse, error) {
	m := new(DownloadCacheTaskResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dfdaemonUploadClient) StatCacheTask(ctx context.Context, in *StatCacheTaskRequest, opts ...grpc.CallOption) (*v2.CacheTask, error) {
	out := new(v2.CacheTask)
	err := c.cc.Invoke(ctx, "/dfdaemon.v2.DfdaemonUpload/StatCacheTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dfdaemonUploadClient) DeleteCacheTask(ctx context.Context, in *DeleteCacheTaskRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/dfdaemon.v2.DfdaemonUpload/DeleteCacheTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DfdaemonUploadServer is the server API for DfdaemonUpload service.
// All implementations should embed UnimplementedDfdaemonUploadServer
// for forward compatibility
type DfdaemonUploadServer interface {
	// DownloadTask downloads task from p2p network.
	DownloadTask(*DownloadTaskRequest, DfdaemonUpload_DownloadTaskServer) error
	// StatTask stats task information.
	StatTask(context.Context, *StatTaskRequest) (*v2.Task, error)
	// DeleteTask deletes task from p2p network.
	DeleteTask(context.Context, *DeleteTaskRequest) (*emptypb.Empty, error)
	// SyncPieces syncs piece metadatas from remote peer.
	SyncPieces(*SyncPiecesRequest, DfdaemonUpload_SyncPiecesServer) error
	// DownloadPiece downloads piece from the remote peer.
	DownloadPiece(context.Context, *DownloadPieceRequest) (*DownloadPieceResponse, error)
	// DownloadCacheTask downloads cache task from p2p network.
	DownloadCacheTask(*DownloadCacheTaskRequest, DfdaemonUpload_DownloadCacheTaskServer) error
	// StatCacheTask stats cache task information.
	StatCacheTask(context.Context, *StatCacheTaskRequest) (*v2.CacheTask, error)
	// DeleteCacheTask deletes cache task from p2p network.
	DeleteCacheTask(context.Context, *DeleteCacheTaskRequest) (*emptypb.Empty, error)
}

// UnimplementedDfdaemonUploadServer should be embedded to have forward compatible implementations.
type UnimplementedDfdaemonUploadServer struct {
}

func (UnimplementedDfdaemonUploadServer) DownloadTask(*DownloadTaskRequest, DfdaemonUpload_DownloadTaskServer) error {
	return status.Errorf(codes.Unimplemented, "method DownloadTask not implemented")
}
func (UnimplementedDfdaemonUploadServer) StatTask(context.Context, *StatTaskRequest) (*v2.Task, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StatTask not implemented")
}
func (UnimplementedDfdaemonUploadServer) DeleteTask(context.Context, *DeleteTaskRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTask not implemented")
}
func (UnimplementedDfdaemonUploadServer) SyncPieces(*SyncPiecesRequest, DfdaemonUpload_SyncPiecesServer) error {
	return status.Errorf(codes.Unimplemented, "method SyncPieces not implemented")
}
func (UnimplementedDfdaemonUploadServer) DownloadPiece(context.Context, *DownloadPieceRequest) (*DownloadPieceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownloadPiece not implemented")
}
func (UnimplementedDfdaemonUploadServer) DownloadCacheTask(*DownloadCacheTaskRequest, DfdaemonUpload_DownloadCacheTaskServer) error {
	return status.Errorf(codes.Unimplemented, "method DownloadCacheTask not implemented")
}
func (UnimplementedDfdaemonUploadServer) StatCacheTask(context.Context, *StatCacheTaskRequest) (*v2.CacheTask, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StatCacheTask not implemented")
}
func (UnimplementedDfdaemonUploadServer) DeleteCacheTask(context.Context, *DeleteCacheTaskRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCacheTask not implemented")
}

// UnsafeDfdaemonUploadServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DfdaemonUploadServer will
// result in compilation errors.
type UnsafeDfdaemonUploadServer interface {
	mustEmbedUnimplementedDfdaemonUploadServer()
}

func RegisterDfdaemonUploadServer(s grpc.ServiceRegistrar, srv DfdaemonUploadServer) {
	s.RegisterService(&DfdaemonUpload_ServiceDesc, srv)
}

func _DfdaemonUpload_DownloadTask_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DownloadTaskRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DfdaemonUploadServer).DownloadTask(m, &dfdaemonUploadDownloadTaskServer{stream})
}

type DfdaemonUpload_DownloadTaskServer interface {
	Send(*DownloadTaskResponse) error
	grpc.ServerStream
}

type dfdaemonUploadDownloadTaskServer struct {
	grpc.ServerStream
}

func (x *dfdaemonUploadDownloadTaskServer) Send(m *DownloadTaskResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _DfdaemonUpload_StatTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DfdaemonUploadServer).StatTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dfdaemon.v2.DfdaemonUpload/StatTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DfdaemonUploadServer).StatTask(ctx, req.(*StatTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DfdaemonUpload_DeleteTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DfdaemonUploadServer).DeleteTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dfdaemon.v2.DfdaemonUpload/DeleteTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DfdaemonUploadServer).DeleteTask(ctx, req.(*DeleteTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DfdaemonUpload_SyncPieces_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SyncPiecesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DfdaemonUploadServer).SyncPieces(m, &dfdaemonUploadSyncPiecesServer{stream})
}

type DfdaemonUpload_SyncPiecesServer interface {
	Send(*SyncPiecesResponse) error
	grpc.ServerStream
}

type dfdaemonUploadSyncPiecesServer struct {
	grpc.ServerStream
}

func (x *dfdaemonUploadSyncPiecesServer) Send(m *SyncPiecesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _DfdaemonUpload_DownloadPiece_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownloadPieceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DfdaemonUploadServer).DownloadPiece(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dfdaemon.v2.DfdaemonUpload/DownloadPiece",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DfdaemonUploadServer).DownloadPiece(ctx, req.(*DownloadPieceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DfdaemonUpload_DownloadCacheTask_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DownloadCacheTaskRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DfdaemonUploadServer).DownloadCacheTask(m, &dfdaemonUploadDownloadCacheTaskServer{stream})
}

type DfdaemonUpload_DownloadCacheTaskServer interface {
	Send(*DownloadCacheTaskResponse) error
	grpc.ServerStream
}

type dfdaemonUploadDownloadCacheTaskServer struct {
	grpc.ServerStream
}

func (x *dfdaemonUploadDownloadCacheTaskServer) Send(m *DownloadCacheTaskResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _DfdaemonUpload_StatCacheTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatCacheTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DfdaemonUploadServer).StatCacheTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dfdaemon.v2.DfdaemonUpload/StatCacheTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DfdaemonUploadServer).StatCacheTask(ctx, req.(*StatCacheTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DfdaemonUpload_DeleteCacheTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCacheTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DfdaemonUploadServer).DeleteCacheTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dfdaemon.v2.DfdaemonUpload/DeleteCacheTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DfdaemonUploadServer).DeleteCacheTask(ctx, req.(*DeleteCacheTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DfdaemonUpload_ServiceDesc is the grpc.ServiceDesc for DfdaemonUpload service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DfdaemonUpload_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dfdaemon.v2.DfdaemonUpload",
	HandlerType: (*DfdaemonUploadServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StatTask",
			Handler:    _DfdaemonUpload_StatTask_Handler,
		},
		{
			MethodName: "DeleteTask",
			Handler:    _DfdaemonUpload_DeleteTask_Handler,
		},
		{
			MethodName: "DownloadPiece",
			Handler:    _DfdaemonUpload_DownloadPiece_Handler,
		},
		{
			MethodName: "StatCacheTask",
			Handler:    _DfdaemonUpload_StatCacheTask_Handler,
		},
		{
			MethodName: "DeleteCacheTask",
			Handler:    _DfdaemonUpload_DeleteCacheTask_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DownloadTask",
			Handler:       _DfdaemonUpload_DownloadTask_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SyncPieces",
			Handler:       _DfdaemonUpload_SyncPieces_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "DownloadCacheTask",
			Handler:       _DfdaemonUpload_DownloadCacheTask_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pkg/apis/dfdaemon/v2/dfdaemon.proto",
}

// DfdaemonDownloadClient is the client API for DfdaemonDownload service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DfdaemonDownloadClient interface {
	// DownloadTask downloads task from p2p network.
	DownloadTask(ctx context.Context, in *DownloadTaskRequest, opts ...grpc.CallOption) (DfdaemonDownload_DownloadTaskClient, error)
	// UploadTask uploads task to p2p network.
	UploadTask(ctx context.Context, in *UploadTaskRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// StatTask stats task information.
	StatTask(ctx context.Context, in *StatTaskRequest, opts ...grpc.CallOption) (*v2.Task, error)
	// DeleteTask deletes task from p2p network.
	DeleteTask(ctx context.Context, in *DeleteTaskRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// DeleteHost releases host in scheduler.
	DeleteHost(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// DownloadCacheTask downloads cache task from p2p network.
	DownloadCacheTask(ctx context.Context, in *DownloadCacheTaskRequest, opts ...grpc.CallOption) (DfdaemonDownload_DownloadCacheTaskClient, error)
	// UploadCacheTask uploads cache task to p2p network.
	UploadCacheTask(ctx context.Context, in *UploadCacheTaskRequest, opts ...grpc.CallOption) (*v2.CacheTask, error)
	// StatCacheTask stats cache task information.
	StatCacheTask(ctx context.Context, in *StatCacheTaskRequest, opts ...grpc.CallOption) (*v2.CacheTask, error)
	// DeleteCacheTask deletes cache task from p2p network.
	DeleteCacheTask(ctx context.Context, in *DeleteCacheTaskRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type dfdaemonDownloadClient struct {
	cc grpc.ClientConnInterface
}

func NewDfdaemonDownloadClient(cc grpc.ClientConnInterface) DfdaemonDownloadClient {
	return &dfdaemonDownloadClient{cc}
}

func (c *dfdaemonDownloadClient) DownloadTask(ctx context.Context, in *DownloadTaskRequest, opts ...grpc.CallOption) (DfdaemonDownload_DownloadTaskClient, error) {
	stream, err := c.cc.NewStream(ctx, &DfdaemonDownload_ServiceDesc.Streams[0], "/dfdaemon.v2.DfdaemonDownload/DownloadTask", opts...)
	if err != nil {
		return nil, err
	}
	x := &dfdaemonDownloadDownloadTaskClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DfdaemonDownload_DownloadTaskClient interface {
	Recv() (*DownloadTaskResponse, error)
	grpc.ClientStream
}

type dfdaemonDownloadDownloadTaskClient struct {
	grpc.ClientStream
}

func (x *dfdaemonDownloadDownloadTaskClient) Recv() (*DownloadTaskResponse, error) {
	m := new(DownloadTaskResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dfdaemonDownloadClient) UploadTask(ctx context.Context, in *UploadTaskRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/dfdaemon.v2.DfdaemonDownload/UploadTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dfdaemonDownloadClient) StatTask(ctx context.Context, in *StatTaskRequest, opts ...grpc.CallOption) (*v2.Task, error) {
	out := new(v2.Task)
	err := c.cc.Invoke(ctx, "/dfdaemon.v2.DfdaemonDownload/StatTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dfdaemonDownloadClient) DeleteTask(ctx context.Context, in *DeleteTaskRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/dfdaemon.v2.DfdaemonDownload/DeleteTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dfdaemonDownloadClient) DeleteHost(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/dfdaemon.v2.DfdaemonDownload/DeleteHost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dfdaemonDownloadClient) DownloadCacheTask(ctx context.Context, in *DownloadCacheTaskRequest, opts ...grpc.CallOption) (DfdaemonDownload_DownloadCacheTaskClient, error) {
	stream, err := c.cc.NewStream(ctx, &DfdaemonDownload_ServiceDesc.Streams[1], "/dfdaemon.v2.DfdaemonDownload/DownloadCacheTask", opts...)
	if err != nil {
		return nil, err
	}
	x := &dfdaemonDownloadDownloadCacheTaskClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DfdaemonDownload_DownloadCacheTaskClient interface {
	Recv() (*DownloadCacheTaskResponse, error)
	grpc.ClientStream
}

type dfdaemonDownloadDownloadCacheTaskClient struct {
	grpc.ClientStream
}

func (x *dfdaemonDownloadDownloadCacheTaskClient) Recv() (*DownloadCacheTaskResponse, error) {
	m := new(DownloadCacheTaskResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dfdaemonDownloadClient) UploadCacheTask(ctx context.Context, in *UploadCacheTaskRequest, opts ...grpc.CallOption) (*v2.CacheTask, error) {
	out := new(v2.CacheTask)
	err := c.cc.Invoke(ctx, "/dfdaemon.v2.DfdaemonDownload/UploadCacheTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dfdaemonDownloadClient) StatCacheTask(ctx context.Context, in *StatCacheTaskRequest, opts ...grpc.CallOption) (*v2.CacheTask, error) {
	out := new(v2.CacheTask)
	err := c.cc.Invoke(ctx, "/dfdaemon.v2.DfdaemonDownload/StatCacheTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dfdaemonDownloadClient) DeleteCacheTask(ctx context.Context, in *DeleteCacheTaskRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/dfdaemon.v2.DfdaemonDownload/DeleteCacheTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DfdaemonDownloadServer is the server API for DfdaemonDownload service.
// All implementations should embed UnimplementedDfdaemonDownloadServer
// for forward compatibility
type DfdaemonDownloadServer interface {
	// DownloadTask downloads task from p2p network.
	DownloadTask(*DownloadTaskRequest, DfdaemonDownload_DownloadTaskServer) error
	// UploadTask uploads task to p2p network.
	UploadTask(context.Context, *UploadTaskRequest) (*emptypb.Empty, error)
	// StatTask stats task information.
	StatTask(context.Context, *StatTaskRequest) (*v2.Task, error)
	// DeleteTask deletes task from p2p network.
	DeleteTask(context.Context, *DeleteTaskRequest) (*emptypb.Empty, error)
	// DeleteHost releases host in scheduler.
	DeleteHost(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	// DownloadCacheTask downloads cache task from p2p network.
	DownloadCacheTask(*DownloadCacheTaskRequest, DfdaemonDownload_DownloadCacheTaskServer) error
	// UploadCacheTask uploads cache task to p2p network.
	UploadCacheTask(context.Context, *UploadCacheTaskRequest) (*v2.CacheTask, error)
	// StatCacheTask stats cache task information.
	StatCacheTask(context.Context, *StatCacheTaskRequest) (*v2.CacheTask, error)
	// DeleteCacheTask deletes cache task from p2p network.
	DeleteCacheTask(context.Context, *DeleteCacheTaskRequest) (*emptypb.Empty, error)
}

// UnimplementedDfdaemonDownloadServer should be embedded to have forward compatible implementations.
type UnimplementedDfdaemonDownloadServer struct {
}

func (UnimplementedDfdaemonDownloadServer) DownloadTask(*DownloadTaskRequest, DfdaemonDownload_DownloadTaskServer) error {
	return status.Errorf(codes.Unimplemented, "method DownloadTask not implemented")
}
func (UnimplementedDfdaemonDownloadServer) UploadTask(context.Context, *UploadTaskRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadTask not implemented")
}
func (UnimplementedDfdaemonDownloadServer) StatTask(context.Context, *StatTaskRequest) (*v2.Task, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StatTask not implemented")
}
func (UnimplementedDfdaemonDownloadServer) DeleteTask(context.Context, *DeleteTaskRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTask not implemented")
}
func (UnimplementedDfdaemonDownloadServer) DeleteHost(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteHost not implemented")
}
func (UnimplementedDfdaemonDownloadServer) DownloadCacheTask(*DownloadCacheTaskRequest, DfdaemonDownload_DownloadCacheTaskServer) error {
	return status.Errorf(codes.Unimplemented, "method DownloadCacheTask not implemented")
}
func (UnimplementedDfdaemonDownloadServer) UploadCacheTask(context.Context, *UploadCacheTaskRequest) (*v2.CacheTask, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadCacheTask not implemented")
}
func (UnimplementedDfdaemonDownloadServer) StatCacheTask(context.Context, *StatCacheTaskRequest) (*v2.CacheTask, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StatCacheTask not implemented")
}
func (UnimplementedDfdaemonDownloadServer) DeleteCacheTask(context.Context, *DeleteCacheTaskRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCacheTask not implemented")
}

// UnsafeDfdaemonDownloadServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DfdaemonDownloadServer will
// result in compilation errors.
type UnsafeDfdaemonDownloadServer interface {
	mustEmbedUnimplementedDfdaemonDownloadServer()
}

func RegisterDfdaemonDownloadServer(s grpc.ServiceRegistrar, srv DfdaemonDownloadServer) {
	s.RegisterService(&DfdaemonDownload_ServiceDesc, srv)
}

func _DfdaemonDownload_DownloadTask_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DownloadTaskRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DfdaemonDownloadServer).DownloadTask(m, &dfdaemonDownloadDownloadTaskServer{stream})
}

type DfdaemonDownload_DownloadTaskServer interface {
	Send(*DownloadTaskResponse) error
	grpc.ServerStream
}

type dfdaemonDownloadDownloadTaskServer struct {
	grpc.ServerStream
}

func (x *dfdaemonDownloadDownloadTaskServer) Send(m *DownloadTaskResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _DfdaemonDownload_UploadTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DfdaemonDownloadServer).UploadTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dfdaemon.v2.DfdaemonDownload/UploadTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DfdaemonDownloadServer).UploadTask(ctx, req.(*UploadTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DfdaemonDownload_StatTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DfdaemonDownloadServer).StatTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dfdaemon.v2.DfdaemonDownload/StatTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DfdaemonDownloadServer).StatTask(ctx, req.(*StatTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DfdaemonDownload_DeleteTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DfdaemonDownloadServer).DeleteTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dfdaemon.v2.DfdaemonDownload/DeleteTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DfdaemonDownloadServer).DeleteTask(ctx, req.(*DeleteTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DfdaemonDownload_DeleteHost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DfdaemonDownloadServer).DeleteHost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dfdaemon.v2.DfdaemonDownload/DeleteHost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DfdaemonDownloadServer).DeleteHost(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DfdaemonDownload_DownloadCacheTask_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DownloadCacheTaskRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DfdaemonDownloadServer).DownloadCacheTask(m, &dfdaemonDownloadDownloadCacheTaskServer{stream})
}

type DfdaemonDownload_DownloadCacheTaskServer interface {
	Send(*DownloadCacheTaskResponse) error
	grpc.ServerStream
}

type dfdaemonDownloadDownloadCacheTaskServer struct {
	grpc.ServerStream
}

func (x *dfdaemonDownloadDownloadCacheTaskServer) Send(m *DownloadCacheTaskResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _DfdaemonDownload_UploadCacheTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadCacheTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DfdaemonDownloadServer).UploadCacheTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dfdaemon.v2.DfdaemonDownload/UploadCacheTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DfdaemonDownloadServer).UploadCacheTask(ctx, req.(*UploadCacheTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DfdaemonDownload_StatCacheTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatCacheTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DfdaemonDownloadServer).StatCacheTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dfdaemon.v2.DfdaemonDownload/StatCacheTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DfdaemonDownloadServer).StatCacheTask(ctx, req.(*StatCacheTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DfdaemonDownload_DeleteCacheTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCacheTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DfdaemonDownloadServer).DeleteCacheTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dfdaemon.v2.DfdaemonDownload/DeleteCacheTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DfdaemonDownloadServer).DeleteCacheTask(ctx, req.(*DeleteCacheTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DfdaemonDownload_ServiceDesc is the grpc.ServiceDesc for DfdaemonDownload service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DfdaemonDownload_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dfdaemon.v2.DfdaemonDownload",
	HandlerType: (*DfdaemonDownloadServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UploadTask",
			Handler:    _DfdaemonDownload_UploadTask_Handler,
		},
		{
			MethodName: "StatTask",
			Handler:    _DfdaemonDownload_StatTask_Handler,
		},
		{
			MethodName: "DeleteTask",
			Handler:    _DfdaemonDownload_DeleteTask_Handler,
		},
		{
			MethodName: "DeleteHost",
			Handler:    _DfdaemonDownload_DeleteHost_Handler,
		},
		{
			MethodName: "UploadCacheTask",
			Handler:    _DfdaemonDownload_UploadCacheTask_Handler,
		},
		{
			MethodName: "StatCacheTask",
			Handler:    _DfdaemonDownload_StatCacheTask_Handler,
		},
		{
			MethodName: "DeleteCacheTask",
			Handler:    _DfdaemonDownload_DeleteCacheTask_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DownloadTask",
			Handler:       _DfdaemonDownload_DownloadTask_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "DownloadCacheTask",
			Handler:       _DfdaemonDownload_DownloadCacheTask_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pkg/apis/dfdaemon/v2/dfdaemon.proto",
}
