// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.2
// source: api/v1/admin/admin.proto

package admin

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
	MenuService_ListMenu_FullMethodName    = "/api.system.v1.MenuService/ListMenu"
	MenuService_CreateMenu_FullMethodName  = "/api.system.v1.MenuService/CreateMenu"
	MenuService_UpdateMenu_FullMethodName  = "/api.system.v1.MenuService/UpdateMenu"
	MenuService_DeleteMenu_FullMethodName  = "/api.system.v1.MenuService/DeleteMenu"
	MenuService_BatchCreate_FullMethodName = "/api.system.v1.MenuService/BatchCreate"
)

// MenuServiceClient is the client API for MenuService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MenuServiceClient interface {
	ListMenu(ctx context.Context, in *ListMenuReq, opts ...grpc.CallOption) (*ListMenuReply, error)
	CreateMenu(ctx context.Context, in *CreateMenuReq, opts ...grpc.CallOption) (*CreateMenuReply, error)
	UpdateMenu(ctx context.Context, in *UpdateMenuReq, opts ...grpc.CallOption) (*UpdateMenuReply, error)
	DeleteMenu(ctx context.Context, in *DeleteMenuReq, opts ...grpc.CallOption) (*DeleteMenuReply, error)
	BatchCreate(ctx context.Context, in *BatchCreateMenusRequest, opts ...grpc.CallOption) (*BatchCreateMenusResponse, error)
}

type menuServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMenuServiceClient(cc grpc.ClientConnInterface) MenuServiceClient {
	return &menuServiceClient{cc}
}

func (c *menuServiceClient) ListMenu(ctx context.Context, in *ListMenuReq, opts ...grpc.CallOption) (*ListMenuReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListMenuReply)
	err := c.cc.Invoke(ctx, MenuService_ListMenu_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *menuServiceClient) CreateMenu(ctx context.Context, in *CreateMenuReq, opts ...grpc.CallOption) (*CreateMenuReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateMenuReply)
	err := c.cc.Invoke(ctx, MenuService_CreateMenu_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *menuServiceClient) UpdateMenu(ctx context.Context, in *UpdateMenuReq, opts ...grpc.CallOption) (*UpdateMenuReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateMenuReply)
	err := c.cc.Invoke(ctx, MenuService_UpdateMenu_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *menuServiceClient) DeleteMenu(ctx context.Context, in *DeleteMenuReq, opts ...grpc.CallOption) (*DeleteMenuReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteMenuReply)
	err := c.cc.Invoke(ctx, MenuService_DeleteMenu_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *menuServiceClient) BatchCreate(ctx context.Context, in *BatchCreateMenusRequest, opts ...grpc.CallOption) (*BatchCreateMenusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BatchCreateMenusResponse)
	err := c.cc.Invoke(ctx, MenuService_BatchCreate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MenuServiceServer is the server API for MenuService service.
// All implementations must embed UnimplementedMenuServiceServer
// for forward compatibility.
type MenuServiceServer interface {
	ListMenu(context.Context, *ListMenuReq) (*ListMenuReply, error)
	CreateMenu(context.Context, *CreateMenuReq) (*CreateMenuReply, error)
	UpdateMenu(context.Context, *UpdateMenuReq) (*UpdateMenuReply, error)
	DeleteMenu(context.Context, *DeleteMenuReq) (*DeleteMenuReply, error)
	BatchCreate(context.Context, *BatchCreateMenusRequest) (*BatchCreateMenusResponse, error)
	mustEmbedUnimplementedMenuServiceServer()
}

// UnimplementedMenuServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMenuServiceServer struct{}

func (UnimplementedMenuServiceServer) ListMenu(context.Context, *ListMenuReq) (*ListMenuReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMenu not implemented")
}
func (UnimplementedMenuServiceServer) CreateMenu(context.Context, *CreateMenuReq) (*CreateMenuReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMenu not implemented")
}
func (UnimplementedMenuServiceServer) UpdateMenu(context.Context, *UpdateMenuReq) (*UpdateMenuReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMenu not implemented")
}
func (UnimplementedMenuServiceServer) DeleteMenu(context.Context, *DeleteMenuReq) (*DeleteMenuReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMenu not implemented")
}
func (UnimplementedMenuServiceServer) BatchCreate(context.Context, *BatchCreateMenusRequest) (*BatchCreateMenusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchCreate not implemented")
}
func (UnimplementedMenuServiceServer) mustEmbedUnimplementedMenuServiceServer() {}
func (UnimplementedMenuServiceServer) testEmbeddedByValue()                     {}

// UnsafeMenuServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MenuServiceServer will
// result in compilation errors.
type UnsafeMenuServiceServer interface {
	mustEmbedUnimplementedMenuServiceServer()
}

func RegisterMenuServiceServer(s grpc.ServiceRegistrar, srv MenuServiceServer) {
	// If the following call pancis, it indicates UnimplementedMenuServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MenuService_ServiceDesc, srv)
}

func _MenuService_ListMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMenuReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MenuServiceServer).ListMenu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MenuService_ListMenu_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MenuServiceServer).ListMenu(ctx, req.(*ListMenuReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MenuService_CreateMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMenuReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MenuServiceServer).CreateMenu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MenuService_CreateMenu_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MenuServiceServer).CreateMenu(ctx, req.(*CreateMenuReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MenuService_UpdateMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMenuReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MenuServiceServer).UpdateMenu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MenuService_UpdateMenu_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MenuServiceServer).UpdateMenu(ctx, req.(*UpdateMenuReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MenuService_DeleteMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMenuReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MenuServiceServer).DeleteMenu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MenuService_DeleteMenu_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MenuServiceServer).DeleteMenu(ctx, req.(*DeleteMenuReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MenuService_BatchCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchCreateMenusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MenuServiceServer).BatchCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MenuService_BatchCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MenuServiceServer).BatchCreate(ctx, req.(*BatchCreateMenusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MenuService_ServiceDesc is the grpc.ServiceDesc for MenuService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MenuService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.system.v1.MenuService",
	HandlerType: (*MenuServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListMenu",
			Handler:    _MenuService_ListMenu_Handler,
		},
		{
			MethodName: "CreateMenu",
			Handler:    _MenuService_CreateMenu_Handler,
		},
		{
			MethodName: "UpdateMenu",
			Handler:    _MenuService_UpdateMenu_Handler,
		},
		{
			MethodName: "DeleteMenu",
			Handler:    _MenuService_DeleteMenu_Handler,
		},
		{
			MethodName: "BatchCreate",
			Handler:    _MenuService_BatchCreate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/admin/admin.proto",
}
