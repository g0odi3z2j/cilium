// Copyright 2016 The gRPC Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Service exported by server reflection

// Warning: this entire file is deprecated. Use this instead:
// https://github.com/grpc/grpc-proto/blob/master/grpc/reflection/v1/reflection.proto

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.14.0
// grpc/reflection/v1alpha/reflection.proto is a deprecated file.

package grpc_reflection_v1alpha

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

// ServerReflectionClient is the client API for ServerReflection service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServerReflectionClient interface {
	// The reflection service is structured as a bidirectional stream, ensuring
	// all related requests go to a single server.
	ServerReflectionInfo(ctx context.Context, opts ...grpc.CallOption) (ServerReflection_ServerReflectionInfoClient, error)
}

type serverReflectionClient struct {
	cc grpc.ClientConnInterface
}

func NewServerReflectionClient(cc grpc.ClientConnInterface) ServerReflectionClient {
	return &serverReflectionClient{cc}
}

func (c *serverReflectionClient) ServerReflectionInfo(ctx context.Context, opts ...grpc.CallOption) (ServerReflection_ServerReflectionInfoClient, error) {
	stream, err := c.cc.NewStream(ctx, &ServerReflection_ServiceDesc.Streams[0], "/grpc.reflection.v1alpha.ServerReflection/ServerReflectionInfo", opts...)
	if err != nil {
		return nil, err
	}
	x := &serverReflectionServerReflectionInfoClient{stream}
	return x, nil
}

type ServerReflection_ServerReflectionInfoClient interface {
	Send(*ServerReflectionRequest) error
	Recv() (*ServerReflectionResponse, error)
	grpc.ClientStream
}

type serverReflectionServerReflectionInfoClient struct {
	grpc.ClientStream
}

func (x *serverReflectionServerReflectionInfoClient) Send(m *ServerReflectionRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *serverReflectionServerReflectionInfoClient) Recv() (*ServerReflectionResponse, error) {
	m := new(ServerReflectionResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ServerReflectionServer is the server API for ServerReflection service.
// All implementations should embed UnimplementedServerReflectionServer
// for forward compatibility
type ServerReflectionServer interface {
	// The reflection service is structured as a bidirectional stream, ensuring
	// all related requests go to a single server.
	ServerReflectionInfo(ServerReflection_ServerReflectionInfoServer) error
}

// UnimplementedServerReflectionServer should be embedded to have forward compatible implementations.
type UnimplementedServerReflectionServer struct {
}

func (UnimplementedServerReflectionServer) ServerReflectionInfo(ServerReflection_ServerReflectionInfoServer) error {
	return status.Errorf(codes.Unimplemented, "method ServerReflectionInfo not implemented")
}

// UnsafeServerReflectionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServerReflectionServer will
// result in compilation errors.
type UnsafeServerReflectionServer interface {
	mustEmbedUnimplementedServerReflectionServer()
}

func RegisterServerReflectionServer(s grpc.ServiceRegistrar, srv ServerReflectionServer) {
	s.RegisterService(&ServerReflection_ServiceDesc, srv)
}

func _ServerReflection_ServerReflectionInfo_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ServerReflectionServer).ServerReflectionInfo(&serverReflectionServerReflectionInfoServer{stream})
}

type ServerReflection_ServerReflectionInfoServer interface {
	Send(*ServerReflectionResponse) error
	Recv() (*ServerReflectionRequest, error)
	grpc.ServerStream
}

type serverReflectionServerReflectionInfoServer struct {
	grpc.ServerStream
}

func (x *serverReflectionServerReflectionInfoServer) Send(m *ServerReflectionResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *serverReflectionServerReflectionInfoServer) Recv() (*ServerReflectionRequest, error) {
	m := new(ServerReflectionRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ServerReflection_ServiceDesc is the grpc.ServiceDesc for ServerReflection service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServerReflection_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.reflection.v1alpha.ServerReflection",
	HandlerType: (*ServerReflectionServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ServerReflectionInfo",
			Handler:       _ServerReflection_ServerReflectionInfo_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "grpc/reflection/v1alpha/reflection.proto",
}
