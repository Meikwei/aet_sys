// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: api/sys/v1/sys_user.proto

package v1

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
	SysUser_Create_FullMethodName         = "/api.sys.v1.sysUser/Create"
	SysUser_DeleteByID_FullMethodName     = "/api.sys.v1.sysUser/DeleteByID"
	SysUser_DeleteByIDs_FullMethodName    = "/api.sys.v1.sysUser/DeleteByIDs"
	SysUser_UpdateByID_FullMethodName     = "/api.sys.v1.sysUser/UpdateByID"
	SysUser_GetByID_FullMethodName        = "/api.sys.v1.sysUser/GetByID"
	SysUser_GetByCondition_FullMethodName = "/api.sys.v1.sysUser/GetByCondition"
	SysUser_ListByIDs_FullMethodName      = "/api.sys.v1.sysUser/ListByIDs"
	SysUser_ListByLastID_FullMethodName   = "/api.sys.v1.sysUser/ListByLastID"
	SysUser_List_FullMethodName           = "/api.sys.v1.sysUser/List"
)

// SysUserClient is the client API for SysUser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SysUserClient interface {
	// create sysUser
	Create(ctx context.Context, in *CreateSysUserRequest, opts ...grpc.CallOption) (*CreateSysUserReply, error)
	// delete sysUser by id
	DeleteByID(ctx context.Context, in *DeleteSysUserByIDRequest, opts ...grpc.CallOption) (*DeleteSysUserByIDReply, error)
	// delete sysUser by batch id
	DeleteByIDs(ctx context.Context, in *DeleteSysUserByIDsRequest, opts ...grpc.CallOption) (*DeleteSysUserByIDsReply, error)
	// update sysUser by id
	UpdateByID(ctx context.Context, in *UpdateSysUserByIDRequest, opts ...grpc.CallOption) (*UpdateSysUserByIDReply, error)
	// get sysUser by id
	GetByID(ctx context.Context, in *GetSysUserByIDRequest, opts ...grpc.CallOption) (*GetSysUserByIDReply, error)
	// get sysUser by condition
	GetByCondition(ctx context.Context, in *GetSysUserByConditionRequest, opts ...grpc.CallOption) (*GetSysUserByConditionReply, error)
	// list of sysUser by batch id
	ListByIDs(ctx context.Context, in *ListSysUserByIDsRequest, opts ...grpc.CallOption) (*ListSysUserByIDsReply, error)
	// list sysUser by last id
	ListByLastID(ctx context.Context, in *ListSysUserByLastIDRequest, opts ...grpc.CallOption) (*ListSysUserByLastIDReply, error)
	// list of sysUser by query parameters
	List(ctx context.Context, in *ListSysUserRequest, opts ...grpc.CallOption) (*ListSysUserReply, error)
}

type sysUserClient struct {
	cc grpc.ClientConnInterface
}

func NewSysUserClient(cc grpc.ClientConnInterface) SysUserClient {
	return &sysUserClient{cc}
}

func (c *sysUserClient) Create(ctx context.Context, in *CreateSysUserRequest, opts ...grpc.CallOption) (*CreateSysUserReply, error) {
	out := new(CreateSysUserReply)
	err := c.cc.Invoke(ctx, SysUser_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysUserClient) DeleteByID(ctx context.Context, in *DeleteSysUserByIDRequest, opts ...grpc.CallOption) (*DeleteSysUserByIDReply, error) {
	out := new(DeleteSysUserByIDReply)
	err := c.cc.Invoke(ctx, SysUser_DeleteByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysUserClient) DeleteByIDs(ctx context.Context, in *DeleteSysUserByIDsRequest, opts ...grpc.CallOption) (*DeleteSysUserByIDsReply, error) {
	out := new(DeleteSysUserByIDsReply)
	err := c.cc.Invoke(ctx, SysUser_DeleteByIDs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysUserClient) UpdateByID(ctx context.Context, in *UpdateSysUserByIDRequest, opts ...grpc.CallOption) (*UpdateSysUserByIDReply, error) {
	out := new(UpdateSysUserByIDReply)
	err := c.cc.Invoke(ctx, SysUser_UpdateByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysUserClient) GetByID(ctx context.Context, in *GetSysUserByIDRequest, opts ...grpc.CallOption) (*GetSysUserByIDReply, error) {
	out := new(GetSysUserByIDReply)
	err := c.cc.Invoke(ctx, SysUser_GetByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysUserClient) GetByCondition(ctx context.Context, in *GetSysUserByConditionRequest, opts ...grpc.CallOption) (*GetSysUserByConditionReply, error) {
	out := new(GetSysUserByConditionReply)
	err := c.cc.Invoke(ctx, SysUser_GetByCondition_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysUserClient) ListByIDs(ctx context.Context, in *ListSysUserByIDsRequest, opts ...grpc.CallOption) (*ListSysUserByIDsReply, error) {
	out := new(ListSysUserByIDsReply)
	err := c.cc.Invoke(ctx, SysUser_ListByIDs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysUserClient) ListByLastID(ctx context.Context, in *ListSysUserByLastIDRequest, opts ...grpc.CallOption) (*ListSysUserByLastIDReply, error) {
	out := new(ListSysUserByLastIDReply)
	err := c.cc.Invoke(ctx, SysUser_ListByLastID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysUserClient) List(ctx context.Context, in *ListSysUserRequest, opts ...grpc.CallOption) (*ListSysUserReply, error) {
	out := new(ListSysUserReply)
	err := c.cc.Invoke(ctx, SysUser_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SysUserServer is the server API for SysUser service.
// All implementations must embed UnimplementedSysUserServer
// for forward compatibility
type SysUserServer interface {
	// create sysUser
	Create(context.Context, *CreateSysUserRequest) (*CreateSysUserReply, error)
	// delete sysUser by id
	DeleteByID(context.Context, *DeleteSysUserByIDRequest) (*DeleteSysUserByIDReply, error)
	// delete sysUser by batch id
	DeleteByIDs(context.Context, *DeleteSysUserByIDsRequest) (*DeleteSysUserByIDsReply, error)
	// update sysUser by id
	UpdateByID(context.Context, *UpdateSysUserByIDRequest) (*UpdateSysUserByIDReply, error)
	// get sysUser by id
	GetByID(context.Context, *GetSysUserByIDRequest) (*GetSysUserByIDReply, error)
	// get sysUser by condition
	GetByCondition(context.Context, *GetSysUserByConditionRequest) (*GetSysUserByConditionReply, error)
	// list of sysUser by batch id
	ListByIDs(context.Context, *ListSysUserByIDsRequest) (*ListSysUserByIDsReply, error)
	// list sysUser by last id
	ListByLastID(context.Context, *ListSysUserByLastIDRequest) (*ListSysUserByLastIDReply, error)
	// list of sysUser by query parameters
	List(context.Context, *ListSysUserRequest) (*ListSysUserReply, error)
	mustEmbedUnimplementedSysUserServer()
}

// UnimplementedSysUserServer must be embedded to have forward compatible implementations.
type UnimplementedSysUserServer struct {
}

func (UnimplementedSysUserServer) Create(context.Context, *CreateSysUserRequest) (*CreateSysUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedSysUserServer) DeleteByID(context.Context, *DeleteSysUserByIDRequest) (*DeleteSysUserByIDReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteByID not implemented")
}
func (UnimplementedSysUserServer) DeleteByIDs(context.Context, *DeleteSysUserByIDsRequest) (*DeleteSysUserByIDsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteByIDs not implemented")
}
func (UnimplementedSysUserServer) UpdateByID(context.Context, *UpdateSysUserByIDRequest) (*UpdateSysUserByIDReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateByID not implemented")
}
func (UnimplementedSysUserServer) GetByID(context.Context, *GetSysUserByIDRequest) (*GetSysUserByIDReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}
func (UnimplementedSysUserServer) GetByCondition(context.Context, *GetSysUserByConditionRequest) (*GetSysUserByConditionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByCondition not implemented")
}
func (UnimplementedSysUserServer) ListByIDs(context.Context, *ListSysUserByIDsRequest) (*ListSysUserByIDsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListByIDs not implemented")
}
func (UnimplementedSysUserServer) ListByLastID(context.Context, *ListSysUserByLastIDRequest) (*ListSysUserByLastIDReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListByLastID not implemented")
}
func (UnimplementedSysUserServer) List(context.Context, *ListSysUserRequest) (*ListSysUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedSysUserServer) mustEmbedUnimplementedSysUserServer() {}

// UnsafeSysUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SysUserServer will
// result in compilation errors.
type UnsafeSysUserServer interface {
	mustEmbedUnimplementedSysUserServer()
}

func RegisterSysUserServer(s grpc.ServiceRegistrar, srv SysUserServer) {
	s.RegisterService(&SysUser_ServiceDesc, srv)
}

func _SysUser_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSysUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysUserServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SysUser_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysUserServer).Create(ctx, req.(*CreateSysUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysUser_DeleteByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSysUserByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysUserServer).DeleteByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SysUser_DeleteByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysUserServer).DeleteByID(ctx, req.(*DeleteSysUserByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysUser_DeleteByIDs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSysUserByIDsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysUserServer).DeleteByIDs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SysUser_DeleteByIDs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysUserServer).DeleteByIDs(ctx, req.(*DeleteSysUserByIDsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysUser_UpdateByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSysUserByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysUserServer).UpdateByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SysUser_UpdateByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysUserServer).UpdateByID(ctx, req.(*UpdateSysUserByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysUser_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSysUserByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysUserServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SysUser_GetByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysUserServer).GetByID(ctx, req.(*GetSysUserByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysUser_GetByCondition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSysUserByConditionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysUserServer).GetByCondition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SysUser_GetByCondition_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysUserServer).GetByCondition(ctx, req.(*GetSysUserByConditionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysUser_ListByIDs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSysUserByIDsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysUserServer).ListByIDs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SysUser_ListByIDs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysUserServer).ListByIDs(ctx, req.(*ListSysUserByIDsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysUser_ListByLastID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSysUserByLastIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysUserServer).ListByLastID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SysUser_ListByLastID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysUserServer).ListByLastID(ctx, req.(*ListSysUserByLastIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysUser_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSysUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysUserServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SysUser_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysUserServer).List(ctx, req.(*ListSysUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SysUser_ServiceDesc is the grpc.ServiceDesc for SysUser service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SysUser_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.sys.v1.sysUser",
	HandlerType: (*SysUserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _SysUser_Create_Handler,
		},
		{
			MethodName: "DeleteByID",
			Handler:    _SysUser_DeleteByID_Handler,
		},
		{
			MethodName: "DeleteByIDs",
			Handler:    _SysUser_DeleteByIDs_Handler,
		},
		{
			MethodName: "UpdateByID",
			Handler:    _SysUser_UpdateByID_Handler,
		},
		{
			MethodName: "GetByID",
			Handler:    _SysUser_GetByID_Handler,
		},
		{
			MethodName: "GetByCondition",
			Handler:    _SysUser_GetByCondition_Handler,
		},
		{
			MethodName: "ListByIDs",
			Handler:    _SysUser_ListByIDs_Handler,
		},
		{
			MethodName: "ListByLastID",
			Handler:    _SysUser_ListByLastID_Handler,
		},
		{
			MethodName: "List",
			Handler:    _SysUser_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/sys/v1/sys_user.proto",
}