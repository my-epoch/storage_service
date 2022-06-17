// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: api/proto/v1/storage_service.proto

package gen

import (
	context "context"
	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// StorageServiceAPIClient is the client API for StorageServiceAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StorageServiceAPIClient interface {
	UploadFile(ctx context.Context, opts ...grpc.CallOption) (StorageServiceAPI_UploadFileClient, error)
	GetRawFile(ctx context.Context, in *GetRawFileRequest, opts ...grpc.CallOption) (StorageServiceAPI_GetRawFileClient, error)
}

type storageServiceAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewStorageServiceAPIClient(cc grpc.ClientConnInterface) StorageServiceAPIClient {
	return &storageServiceAPIClient{cc}
}

func (c *storageServiceAPIClient) UploadFile(ctx context.Context, opts ...grpc.CallOption) (StorageServiceAPI_UploadFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &StorageServiceAPI_ServiceDesc.Streams[0], "/storage_service_api.StorageServiceAPI/UploadFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &storageServiceAPIUploadFileClient{stream}
	return x, nil
}

type StorageServiceAPI_UploadFileClient interface {
	Send(*httpbody.HttpBody) error
	CloseAndRecv() (*FileUploadResponse, error)
	grpc.ClientStream
}

type storageServiceAPIUploadFileClient struct {
	grpc.ClientStream
}

func (x *storageServiceAPIUploadFileClient) Send(m *httpbody.HttpBody) error {
	return x.ClientStream.SendMsg(m)
}

func (x *storageServiceAPIUploadFileClient) CloseAndRecv() (*FileUploadResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(FileUploadResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *storageServiceAPIClient) GetRawFile(ctx context.Context, in *GetRawFileRequest, opts ...grpc.CallOption) (StorageServiceAPI_GetRawFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &StorageServiceAPI_ServiceDesc.Streams[1], "/storage_service_api.StorageServiceAPI/GetRawFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &storageServiceAPIGetRawFileClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StorageServiceAPI_GetRawFileClient interface {
	Recv() (*httpbody.HttpBody, error)
	grpc.ClientStream
}

type storageServiceAPIGetRawFileClient struct {
	grpc.ClientStream
}

func (x *storageServiceAPIGetRawFileClient) Recv() (*httpbody.HttpBody, error) {
	m := new(httpbody.HttpBody)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StorageServiceAPIServer is the server API for StorageServiceAPI service.
// All implementations must embed UnimplementedStorageServiceAPIServer
// for forward compatibility
type StorageServiceAPIServer interface {
	UploadFile(StorageServiceAPI_UploadFileServer) error
	GetRawFile(*GetRawFileRequest, StorageServiceAPI_GetRawFileServer) error
	mustEmbedUnimplementedStorageServiceAPIServer()
}

// UnimplementedStorageServiceAPIServer must be embedded to have forward compatible implementations.
type UnimplementedStorageServiceAPIServer struct {
}

func (UnimplementedStorageServiceAPIServer) UploadFile(StorageServiceAPI_UploadFileServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
}
func (UnimplementedStorageServiceAPIServer) GetRawFile(*GetRawFileRequest, StorageServiceAPI_GetRawFileServer) error {
	return status.Errorf(codes.Unimplemented, "method GetRawFile not implemented")
}
func (UnimplementedStorageServiceAPIServer) mustEmbedUnimplementedStorageServiceAPIServer() {}

// UnsafeStorageServiceAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StorageServiceAPIServer will
// result in compilation errors.
type UnsafeStorageServiceAPIServer interface {
	mustEmbedUnimplementedStorageServiceAPIServer()
}

func RegisterStorageServiceAPIServer(s grpc.ServiceRegistrar, srv StorageServiceAPIServer) {
	s.RegisterService(&StorageServiceAPI_ServiceDesc, srv)
}

func _StorageServiceAPI_UploadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StorageServiceAPIServer).UploadFile(&storageServiceAPIUploadFileServer{stream})
}

type StorageServiceAPI_UploadFileServer interface {
	SendAndClose(*FileUploadResponse) error
	Recv() (*httpbody.HttpBody, error)
	grpc.ServerStream
}

type storageServiceAPIUploadFileServer struct {
	grpc.ServerStream
}

func (x *storageServiceAPIUploadFileServer) SendAndClose(m *FileUploadResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *storageServiceAPIUploadFileServer) Recv() (*httpbody.HttpBody, error) {
	m := new(httpbody.HttpBody)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _StorageServiceAPI_GetRawFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetRawFileRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StorageServiceAPIServer).GetRawFile(m, &storageServiceAPIGetRawFileServer{stream})
}

type StorageServiceAPI_GetRawFileServer interface {
	Send(*httpbody.HttpBody) error
	grpc.ServerStream
}

type storageServiceAPIGetRawFileServer struct {
	grpc.ServerStream
}

func (x *storageServiceAPIGetRawFileServer) Send(m *httpbody.HttpBody) error {
	return x.ServerStream.SendMsg(m)
}

// StorageServiceAPI_ServiceDesc is the grpc.ServiceDesc for StorageServiceAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StorageServiceAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "storage_service_api.StorageServiceAPI",
	HandlerType: (*StorageServiceAPIServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadFile",
			Handler:       _StorageServiceAPI_UploadFile_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "GetRawFile",
			Handler:       _StorageServiceAPI_GetRawFile_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/proto/v1/storage_service.proto",
}
