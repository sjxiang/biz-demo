// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.1
// source: idl/note.proto

package note

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
	NoteService_CreateNote_FullMethodName = "/note.NoteService/CreateNote"
	NoteService_MGetNote_FullMethodName   = "/note.NoteService/MGetNote"
	NoteService_DeleteNote_FullMethodName = "/note.NoteService/DeleteNote"
	NoteService_QueryNote_FullMethodName  = "/note.NoteService/QueryNote"
	NoteService_UpdateNote_FullMethodName = "/note.NoteService/UpdateNote"
)

// NoteServiceClient is the client API for NoteService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NoteServiceClient interface {
	CreateNote(ctx context.Context, in *CreateNoteRequest, opts ...grpc.CallOption) (*CreateNoteResponse, error)
	MGetNote(ctx context.Context, in *MGetNoteRequest, opts ...grpc.CallOption) (*MGetNoteResponse, error)
	DeleteNote(ctx context.Context, in *DeleteNoteRequest, opts ...grpc.CallOption) (*DeleteNoteResponse, error)
	QueryNote(ctx context.Context, in *QueryNoteRequest, opts ...grpc.CallOption) (*QueryNoteResponse, error)
	UpdateNote(ctx context.Context, in *UpdateNoteRequest, opts ...grpc.CallOption) (*UpdateNoteResponse, error)
}

type noteServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNoteServiceClient(cc grpc.ClientConnInterface) NoteServiceClient {
	return &noteServiceClient{cc}
}

func (c *noteServiceClient) CreateNote(ctx context.Context, in *CreateNoteRequest, opts ...grpc.CallOption) (*CreateNoteResponse, error) {
	out := new(CreateNoteResponse)
	err := c.cc.Invoke(ctx, NoteService_CreateNote_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noteServiceClient) MGetNote(ctx context.Context, in *MGetNoteRequest, opts ...grpc.CallOption) (*MGetNoteResponse, error) {
	out := new(MGetNoteResponse)
	err := c.cc.Invoke(ctx, NoteService_MGetNote_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noteServiceClient) DeleteNote(ctx context.Context, in *DeleteNoteRequest, opts ...grpc.CallOption) (*DeleteNoteResponse, error) {
	out := new(DeleteNoteResponse)
	err := c.cc.Invoke(ctx, NoteService_DeleteNote_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noteServiceClient) QueryNote(ctx context.Context, in *QueryNoteRequest, opts ...grpc.CallOption) (*QueryNoteResponse, error) {
	out := new(QueryNoteResponse)
	err := c.cc.Invoke(ctx, NoteService_QueryNote_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noteServiceClient) UpdateNote(ctx context.Context, in *UpdateNoteRequest, opts ...grpc.CallOption) (*UpdateNoteResponse, error) {
	out := new(UpdateNoteResponse)
	err := c.cc.Invoke(ctx, NoteService_UpdateNote_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NoteServiceServer is the server API for NoteService service.
// All implementations must embed UnimplementedNoteServiceServer
// for forward compatibility
type NoteServiceServer interface {
	CreateNote(context.Context, *CreateNoteRequest) (*CreateNoteResponse, error)
	MGetNote(context.Context, *MGetNoteRequest) (*MGetNoteResponse, error)
	DeleteNote(context.Context, *DeleteNoteRequest) (*DeleteNoteResponse, error)
	QueryNote(context.Context, *QueryNoteRequest) (*QueryNoteResponse, error)
	UpdateNote(context.Context, *UpdateNoteRequest) (*UpdateNoteResponse, error)
	mustEmbedUnimplementedNoteServiceServer()
}

// UnimplementedNoteServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNoteServiceServer struct {
}

func (UnimplementedNoteServiceServer) CreateNote(context.Context, *CreateNoteRequest) (*CreateNoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNote not implemented")
}
func (UnimplementedNoteServiceServer) MGetNote(context.Context, *MGetNoteRequest) (*MGetNoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MGetNote not implemented")
}
func (UnimplementedNoteServiceServer) DeleteNote(context.Context, *DeleteNoteRequest) (*DeleteNoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNote not implemented")
}
func (UnimplementedNoteServiceServer) QueryNote(context.Context, *QueryNoteRequest) (*QueryNoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryNote not implemented")
}
func (UnimplementedNoteServiceServer) UpdateNote(context.Context, *UpdateNoteRequest) (*UpdateNoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNote not implemented")
}
func (UnimplementedNoteServiceServer) mustEmbedUnimplementedNoteServiceServer() {}

// UnsafeNoteServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NoteServiceServer will
// result in compilation errors.
type UnsafeNoteServiceServer interface {
	mustEmbedUnimplementedNoteServiceServer()
}

func RegisterNoteServiceServer(s grpc.ServiceRegistrar, srv NoteServiceServer) {
	s.RegisterService(&NoteService_ServiceDesc, srv)
}

func _NoteService_CreateNote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoteServiceServer).CreateNote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NoteService_CreateNote_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoteServiceServer).CreateNote(ctx, req.(*CreateNoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NoteService_MGetNote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MGetNoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoteServiceServer).MGetNote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NoteService_MGetNote_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoteServiceServer).MGetNote(ctx, req.(*MGetNoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NoteService_DeleteNote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoteServiceServer).DeleteNote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NoteService_DeleteNote_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoteServiceServer).DeleteNote(ctx, req.(*DeleteNoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NoteService_QueryNote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryNoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoteServiceServer).QueryNote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NoteService_QueryNote_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoteServiceServer).QueryNote(ctx, req.(*QueryNoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NoteService_UpdateNote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoteServiceServer).UpdateNote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NoteService_UpdateNote_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoteServiceServer).UpdateNote(ctx, req.(*UpdateNoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NoteService_ServiceDesc is the grpc.ServiceDesc for NoteService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NoteService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "note.NoteService",
	HandlerType: (*NoteServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateNote",
			Handler:    _NoteService_CreateNote_Handler,
		},
		{
			MethodName: "MGetNote",
			Handler:    _NoteService_MGetNote_Handler,
		},
		{
			MethodName: "DeleteNote",
			Handler:    _NoteService_DeleteNote_Handler,
		},
		{
			MethodName: "QueryNote",
			Handler:    _NoteService_QueryNote_Handler,
		},
		{
			MethodName: "UpdateNote",
			Handler:    _NoteService_UpdateNote_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "idl/note.proto",
}
