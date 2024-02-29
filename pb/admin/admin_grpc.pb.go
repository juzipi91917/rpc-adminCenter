// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.9
// source: admin.proto

package admin

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
	AdminService_LogOn_FullMethodName            = "/adminCenter.AdminService/LogOn"
	AdminService_RefreshToken_FullMethodName     = "/adminCenter.AdminService/RefreshToken"
	AdminService_GetAdminMenuList_FullMethodName = "/adminCenter.AdminService/GetAdminMenuList"
)

// AdminServiceClient is the client API for AdminService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminServiceClient interface {
	LogOn(ctx context.Context, in *LogOnRequest, opts ...grpc.CallOption) (*LogOnResponse, error)
	RefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...grpc.CallOption) (*RefreshTokenResponse, error)
	GetAdminMenuList(ctx context.Context, in *GetAdminMenuListRequest, opts ...grpc.CallOption) (*GetAdminMenuListResponse, error)
}

type adminServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminServiceClient(cc grpc.ClientConnInterface) AdminServiceClient {
	return &adminServiceClient{cc}
}

func (c *adminServiceClient) LogOn(ctx context.Context, in *LogOnRequest, opts ...grpc.CallOption) (*LogOnResponse, error) {
	out := new(LogOnResponse)
	err := c.cc.Invoke(ctx, AdminService_LogOn_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) RefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...grpc.CallOption) (*RefreshTokenResponse, error) {
	out := new(RefreshTokenResponse)
	err := c.cc.Invoke(ctx, AdminService_RefreshToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) GetAdminMenuList(ctx context.Context, in *GetAdminMenuListRequest, opts ...grpc.CallOption) (*GetAdminMenuListResponse, error) {
	out := new(GetAdminMenuListResponse)
	err := c.cc.Invoke(ctx, AdminService_GetAdminMenuList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminServiceServer is the server API for AdminService service.
// All implementations must embed UnimplementedAdminServiceServer
// for forward compatibility
type AdminServiceServer interface {
	LogOn(context.Context, *LogOnRequest) (*LogOnResponse, error)
	RefreshToken(context.Context, *RefreshTokenRequest) (*RefreshTokenResponse, error)
	GetAdminMenuList(context.Context, *GetAdminMenuListRequest) (*GetAdminMenuListResponse, error)
	mustEmbedUnimplementedAdminServiceServer()
}

// UnimplementedAdminServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAdminServiceServer struct {
}

func (UnimplementedAdminServiceServer) LogOn(context.Context, *LogOnRequest) (*LogOnResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogOn not implemented")
}
func (UnimplementedAdminServiceServer) RefreshToken(context.Context, *RefreshTokenRequest) (*RefreshTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshToken not implemented")
}
func (UnimplementedAdminServiceServer) GetAdminMenuList(context.Context, *GetAdminMenuListRequest) (*GetAdminMenuListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAdminMenuList not implemented")
}
func (UnimplementedAdminServiceServer) mustEmbedUnimplementedAdminServiceServer() {}

// UnsafeAdminServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdminServiceServer will
// result in compilation errors.
type UnsafeAdminServiceServer interface {
	mustEmbedUnimplementedAdminServiceServer()
}

func RegisterAdminServiceServer(s grpc.ServiceRegistrar, srv AdminServiceServer) {
	s.RegisterService(&AdminService_ServiceDesc, srv)
}

func _AdminService_LogOn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogOnRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).LogOn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_LogOn_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).LogOn(ctx, req.(*LogOnRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_RefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).RefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_RefreshToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).RefreshToken(ctx, req.(*RefreshTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_GetAdminMenuList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAdminMenuListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).GetAdminMenuList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_GetAdminMenuList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).GetAdminMenuList(ctx, req.(*GetAdminMenuListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AdminService_ServiceDesc is the grpc.ServiceDesc for AdminService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdminService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "adminCenter.AdminService",
	HandlerType: (*AdminServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LogOn",
			Handler:    _AdminService_LogOn_Handler,
		},
		{
			MethodName: "RefreshToken",
			Handler:    _AdminService_RefreshToken_Handler,
		},
		{
			MethodName: "GetAdminMenuList",
			Handler:    _AdminService_GetAdminMenuList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admin.proto",
}

const (
	PermissionService_AddMenu_FullMethodName               = "/adminCenter.PermissionService/AddMenu"
	PermissionService_UpdateMenu_FullMethodName            = "/adminCenter.PermissionService/UpdateMenu"
	PermissionService_GetMenuList_FullMethodName           = "/adminCenter.PermissionService/GetMenuList"
	PermissionService_AddPermission_FullMethodName         = "/adminCenter.PermissionService/AddPermission"
	PermissionService_UpdatePermission_FullMethodName      = "/adminCenter.PermissionService/UpdatePermission"
	PermissionService_GetPermissionList_FullMethodName     = "/adminCenter.PermissionService/GetPermissionList"
	PermissionService_AddRole_FullMethodName               = "/adminCenter.PermissionService/AddRole"
	PermissionService_UpdateRole_FullMethodName            = "/adminCenter.PermissionService/UpdateRole"
	PermissionService_GetRoleList_FullMethodName           = "/adminCenter.PermissionService/GetRoleList"
	PermissionService_AddAdmin_FullMethodName              = "/adminCenter.PermissionService/AddAdmin"
	PermissionService_UpdateAdmin_FullMethodName           = "/adminCenter.PermissionService/UpdateAdmin"
	PermissionService_GetAdminList_FullMethodName          = "/adminCenter.PermissionService/GetAdminList"
	PermissionService_GetAdminRoleList_FullMethodName      = "/adminCenter.PermissionService/GetAdminRoleList"
	PermissionService_GetRolePermissionList_FullMethodName = "/adminCenter.PermissionService/GetRolePermissionList"
	PermissionService_BindAdminRole_FullMethodName         = "/adminCenter.PermissionService/BindAdminRole"
	PermissionService_BindRolePermission_FullMethodName    = "/adminCenter.PermissionService/BindRolePermission"
	PermissionService_DelAdminRole_FullMethodName          = "/adminCenter.PermissionService/DelAdminRole"
	PermissionService_DelRolePermission_FullMethodName     = "/adminCenter.PermissionService/DelRolePermission"
)

// PermissionServiceClient is the client API for PermissionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PermissionServiceClient interface {
	AddMenu(ctx context.Context, in *AddMenuRequest, opts ...grpc.CallOption) (*AddMenuResponse, error)
	UpdateMenu(ctx context.Context, in *UpdateMenuRequest, opts ...grpc.CallOption) (*UpdateMenuResponse, error)
	GetMenuList(ctx context.Context, in *GetMenuListRequest, opts ...grpc.CallOption) (*GetMenuListResponse, error)
	AddPermission(ctx context.Context, in *AddPermissionRequest, opts ...grpc.CallOption) (*AddPermissionResponse, error)
	UpdatePermission(ctx context.Context, in *UpdatePermissionRequest, opts ...grpc.CallOption) (*UpdatePermissionResponse, error)
	GetPermissionList(ctx context.Context, in *GetPermissionListRequest, opts ...grpc.CallOption) (*GetPermissionListResponse, error)
	AddRole(ctx context.Context, in *AddRoleRequest, opts ...grpc.CallOption) (*AddRoleResponse, error)
	UpdateRole(ctx context.Context, in *UpdateRoleRequest, opts ...grpc.CallOption) (*UpdateRoleResponse, error)
	GetRoleList(ctx context.Context, in *GetRoleListRequest, opts ...grpc.CallOption) (*GetRoleListResponse, error)
	AddAdmin(ctx context.Context, in *AddAdminRequest, opts ...grpc.CallOption) (*AddAdminResponse, error)
	UpdateAdmin(ctx context.Context, in *UpdateAdminRequest, opts ...grpc.CallOption) (*UpdateAdminResponse, error)
	GetAdminList(ctx context.Context, in *GetAdminListRequest, opts ...grpc.CallOption) (*GetAdminListResponse, error)
	GetAdminRoleList(ctx context.Context, in *GetAdminRoleListRequest, opts ...grpc.CallOption) (*GetAdminRoleListResponse, error)
	GetRolePermissionList(ctx context.Context, in *GetRolePermissionListRequest, opts ...grpc.CallOption) (*GetRolePermissionListResponse, error)
	BindAdminRole(ctx context.Context, in *BindAdminRoleRequest, opts ...grpc.CallOption) (*BindAdminRoleResponse, error)
	BindRolePermission(ctx context.Context, in *BindRolePermissionRequest, opts ...grpc.CallOption) (*BindRolePermissionResponse, error)
	DelAdminRole(ctx context.Context, in *DelAdminRoleRequest, opts ...grpc.CallOption) (*DelAdminRoleResponse, error)
	DelRolePermission(ctx context.Context, in *DelRolePermissionRequest, opts ...grpc.CallOption) (*DelRolePermissionResponse, error)
}

type permissionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPermissionServiceClient(cc grpc.ClientConnInterface) PermissionServiceClient {
	return &permissionServiceClient{cc}
}

func (c *permissionServiceClient) AddMenu(ctx context.Context, in *AddMenuRequest, opts ...grpc.CallOption) (*AddMenuResponse, error) {
	out := new(AddMenuResponse)
	err := c.cc.Invoke(ctx, PermissionService_AddMenu_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionServiceClient) UpdateMenu(ctx context.Context, in *UpdateMenuRequest, opts ...grpc.CallOption) (*UpdateMenuResponse, error) {
	out := new(UpdateMenuResponse)
	err := c.cc.Invoke(ctx, PermissionService_UpdateMenu_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionServiceClient) GetMenuList(ctx context.Context, in *GetMenuListRequest, opts ...grpc.CallOption) (*GetMenuListResponse, error) {
	out := new(GetMenuListResponse)
	err := c.cc.Invoke(ctx, PermissionService_GetMenuList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionServiceClient) AddPermission(ctx context.Context, in *AddPermissionRequest, opts ...grpc.CallOption) (*AddPermissionResponse, error) {
	out := new(AddPermissionResponse)
	err := c.cc.Invoke(ctx, PermissionService_AddPermission_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionServiceClient) UpdatePermission(ctx context.Context, in *UpdatePermissionRequest, opts ...grpc.CallOption) (*UpdatePermissionResponse, error) {
	out := new(UpdatePermissionResponse)
	err := c.cc.Invoke(ctx, PermissionService_UpdatePermission_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionServiceClient) GetPermissionList(ctx context.Context, in *GetPermissionListRequest, opts ...grpc.CallOption) (*GetPermissionListResponse, error) {
	out := new(GetPermissionListResponse)
	err := c.cc.Invoke(ctx, PermissionService_GetPermissionList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionServiceClient) AddRole(ctx context.Context, in *AddRoleRequest, opts ...grpc.CallOption) (*AddRoleResponse, error) {
	out := new(AddRoleResponse)
	err := c.cc.Invoke(ctx, PermissionService_AddRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionServiceClient) UpdateRole(ctx context.Context, in *UpdateRoleRequest, opts ...grpc.CallOption) (*UpdateRoleResponse, error) {
	out := new(UpdateRoleResponse)
	err := c.cc.Invoke(ctx, PermissionService_UpdateRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionServiceClient) GetRoleList(ctx context.Context, in *GetRoleListRequest, opts ...grpc.CallOption) (*GetRoleListResponse, error) {
	out := new(GetRoleListResponse)
	err := c.cc.Invoke(ctx, PermissionService_GetRoleList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionServiceClient) AddAdmin(ctx context.Context, in *AddAdminRequest, opts ...grpc.CallOption) (*AddAdminResponse, error) {
	out := new(AddAdminResponse)
	err := c.cc.Invoke(ctx, PermissionService_AddAdmin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionServiceClient) UpdateAdmin(ctx context.Context, in *UpdateAdminRequest, opts ...grpc.CallOption) (*UpdateAdminResponse, error) {
	out := new(UpdateAdminResponse)
	err := c.cc.Invoke(ctx, PermissionService_UpdateAdmin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionServiceClient) GetAdminList(ctx context.Context, in *GetAdminListRequest, opts ...grpc.CallOption) (*GetAdminListResponse, error) {
	out := new(GetAdminListResponse)
	err := c.cc.Invoke(ctx, PermissionService_GetAdminList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionServiceClient) GetAdminRoleList(ctx context.Context, in *GetAdminRoleListRequest, opts ...grpc.CallOption) (*GetAdminRoleListResponse, error) {
	out := new(GetAdminRoleListResponse)
	err := c.cc.Invoke(ctx, PermissionService_GetAdminRoleList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionServiceClient) GetRolePermissionList(ctx context.Context, in *GetRolePermissionListRequest, opts ...grpc.CallOption) (*GetRolePermissionListResponse, error) {
	out := new(GetRolePermissionListResponse)
	err := c.cc.Invoke(ctx, PermissionService_GetRolePermissionList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionServiceClient) BindAdminRole(ctx context.Context, in *BindAdminRoleRequest, opts ...grpc.CallOption) (*BindAdminRoleResponse, error) {
	out := new(BindAdminRoleResponse)
	err := c.cc.Invoke(ctx, PermissionService_BindAdminRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionServiceClient) BindRolePermission(ctx context.Context, in *BindRolePermissionRequest, opts ...grpc.CallOption) (*BindRolePermissionResponse, error) {
	out := new(BindRolePermissionResponse)
	err := c.cc.Invoke(ctx, PermissionService_BindRolePermission_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionServiceClient) DelAdminRole(ctx context.Context, in *DelAdminRoleRequest, opts ...grpc.CallOption) (*DelAdminRoleResponse, error) {
	out := new(DelAdminRoleResponse)
	err := c.cc.Invoke(ctx, PermissionService_DelAdminRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionServiceClient) DelRolePermission(ctx context.Context, in *DelRolePermissionRequest, opts ...grpc.CallOption) (*DelRolePermissionResponse, error) {
	out := new(DelRolePermissionResponse)
	err := c.cc.Invoke(ctx, PermissionService_DelRolePermission_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionServiceServer is the server API for PermissionService service.
// All implementations must embed UnimplementedPermissionServiceServer
// for forward compatibility
type PermissionServiceServer interface {
	AddMenu(context.Context, *AddMenuRequest) (*AddMenuResponse, error)
	UpdateMenu(context.Context, *UpdateMenuRequest) (*UpdateMenuResponse, error)
	GetMenuList(context.Context, *GetMenuListRequest) (*GetMenuListResponse, error)
	AddPermission(context.Context, *AddPermissionRequest) (*AddPermissionResponse, error)
	UpdatePermission(context.Context, *UpdatePermissionRequest) (*UpdatePermissionResponse, error)
	GetPermissionList(context.Context, *GetPermissionListRequest) (*GetPermissionListResponse, error)
	AddRole(context.Context, *AddRoleRequest) (*AddRoleResponse, error)
	UpdateRole(context.Context, *UpdateRoleRequest) (*UpdateRoleResponse, error)
	GetRoleList(context.Context, *GetRoleListRequest) (*GetRoleListResponse, error)
	AddAdmin(context.Context, *AddAdminRequest) (*AddAdminResponse, error)
	UpdateAdmin(context.Context, *UpdateAdminRequest) (*UpdateAdminResponse, error)
	GetAdminList(context.Context, *GetAdminListRequest) (*GetAdminListResponse, error)
	GetAdminRoleList(context.Context, *GetAdminRoleListRequest) (*GetAdminRoleListResponse, error)
	GetRolePermissionList(context.Context, *GetRolePermissionListRequest) (*GetRolePermissionListResponse, error)
	BindAdminRole(context.Context, *BindAdminRoleRequest) (*BindAdminRoleResponse, error)
	BindRolePermission(context.Context, *BindRolePermissionRequest) (*BindRolePermissionResponse, error)
	DelAdminRole(context.Context, *DelAdminRoleRequest) (*DelAdminRoleResponse, error)
	DelRolePermission(context.Context, *DelRolePermissionRequest) (*DelRolePermissionResponse, error)
	mustEmbedUnimplementedPermissionServiceServer()
}

// UnimplementedPermissionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPermissionServiceServer struct {
}

func (UnimplementedPermissionServiceServer) AddMenu(context.Context, *AddMenuRequest) (*AddMenuResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMenu not implemented")
}
func (UnimplementedPermissionServiceServer) UpdateMenu(context.Context, *UpdateMenuRequest) (*UpdateMenuResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMenu not implemented")
}
func (UnimplementedPermissionServiceServer) GetMenuList(context.Context, *GetMenuListRequest) (*GetMenuListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMenuList not implemented")
}
func (UnimplementedPermissionServiceServer) AddPermission(context.Context, *AddPermissionRequest) (*AddPermissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPermission not implemented")
}
func (UnimplementedPermissionServiceServer) UpdatePermission(context.Context, *UpdatePermissionRequest) (*UpdatePermissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePermission not implemented")
}
func (UnimplementedPermissionServiceServer) GetPermissionList(context.Context, *GetPermissionListRequest) (*GetPermissionListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPermissionList not implemented")
}
func (UnimplementedPermissionServiceServer) AddRole(context.Context, *AddRoleRequest) (*AddRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddRole not implemented")
}
func (UnimplementedPermissionServiceServer) UpdateRole(context.Context, *UpdateRoleRequest) (*UpdateRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRole not implemented")
}
func (UnimplementedPermissionServiceServer) GetRoleList(context.Context, *GetRoleListRequest) (*GetRoleListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoleList not implemented")
}
func (UnimplementedPermissionServiceServer) AddAdmin(context.Context, *AddAdminRequest) (*AddAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAdmin not implemented")
}
func (UnimplementedPermissionServiceServer) UpdateAdmin(context.Context, *UpdateAdminRequest) (*UpdateAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAdmin not implemented")
}
func (UnimplementedPermissionServiceServer) GetAdminList(context.Context, *GetAdminListRequest) (*GetAdminListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAdminList not implemented")
}
func (UnimplementedPermissionServiceServer) GetAdminRoleList(context.Context, *GetAdminRoleListRequest) (*GetAdminRoleListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAdminRoleList not implemented")
}
func (UnimplementedPermissionServiceServer) GetRolePermissionList(context.Context, *GetRolePermissionListRequest) (*GetRolePermissionListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRolePermissionList not implemented")
}
func (UnimplementedPermissionServiceServer) BindAdminRole(context.Context, *BindAdminRoleRequest) (*BindAdminRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BindAdminRole not implemented")
}
func (UnimplementedPermissionServiceServer) BindRolePermission(context.Context, *BindRolePermissionRequest) (*BindRolePermissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BindRolePermission not implemented")
}
func (UnimplementedPermissionServiceServer) DelAdminRole(context.Context, *DelAdminRoleRequest) (*DelAdminRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelAdminRole not implemented")
}
func (UnimplementedPermissionServiceServer) DelRolePermission(context.Context, *DelRolePermissionRequest) (*DelRolePermissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelRolePermission not implemented")
}
func (UnimplementedPermissionServiceServer) mustEmbedUnimplementedPermissionServiceServer() {}

// UnsafePermissionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PermissionServiceServer will
// result in compilation errors.
type UnsafePermissionServiceServer interface {
	mustEmbedUnimplementedPermissionServiceServer()
}

func RegisterPermissionServiceServer(s grpc.ServiceRegistrar, srv PermissionServiceServer) {
	s.RegisterService(&PermissionService_ServiceDesc, srv)
}

func _PermissionService_AddMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddMenuRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServiceServer).AddMenu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PermissionService_AddMenu_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServiceServer).AddMenu(ctx, req.(*AddMenuRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionService_UpdateMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMenuRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServiceServer).UpdateMenu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PermissionService_UpdateMenu_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServiceServer).UpdateMenu(ctx, req.(*UpdateMenuRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionService_GetMenuList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMenuListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServiceServer).GetMenuList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PermissionService_GetMenuList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServiceServer).GetMenuList(ctx, req.(*GetMenuListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionService_AddPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServiceServer).AddPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PermissionService_AddPermission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServiceServer).AddPermission(ctx, req.(*AddPermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionService_UpdatePermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServiceServer).UpdatePermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PermissionService_UpdatePermission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServiceServer).UpdatePermission(ctx, req.(*UpdatePermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionService_GetPermissionList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPermissionListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServiceServer).GetPermissionList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PermissionService_GetPermissionList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServiceServer).GetPermissionList(ctx, req.(*GetPermissionListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionService_AddRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServiceServer).AddRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PermissionService_AddRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServiceServer).AddRole(ctx, req.(*AddRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionService_UpdateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServiceServer).UpdateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PermissionService_UpdateRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServiceServer).UpdateRole(ctx, req.(*UpdateRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionService_GetRoleList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoleListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServiceServer).GetRoleList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PermissionService_GetRoleList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServiceServer).GetRoleList(ctx, req.(*GetRoleListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionService_AddAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServiceServer).AddAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PermissionService_AddAdmin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServiceServer).AddAdmin(ctx, req.(*AddAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionService_UpdateAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServiceServer).UpdateAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PermissionService_UpdateAdmin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServiceServer).UpdateAdmin(ctx, req.(*UpdateAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionService_GetAdminList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAdminListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServiceServer).GetAdminList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PermissionService_GetAdminList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServiceServer).GetAdminList(ctx, req.(*GetAdminListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionService_GetAdminRoleList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAdminRoleListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServiceServer).GetAdminRoleList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PermissionService_GetAdminRoleList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServiceServer).GetAdminRoleList(ctx, req.(*GetAdminRoleListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionService_GetRolePermissionList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRolePermissionListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServiceServer).GetRolePermissionList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PermissionService_GetRolePermissionList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServiceServer).GetRolePermissionList(ctx, req.(*GetRolePermissionListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionService_BindAdminRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BindAdminRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServiceServer).BindAdminRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PermissionService_BindAdminRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServiceServer).BindAdminRole(ctx, req.(*BindAdminRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionService_BindRolePermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BindRolePermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServiceServer).BindRolePermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PermissionService_BindRolePermission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServiceServer).BindRolePermission(ctx, req.(*BindRolePermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionService_DelAdminRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelAdminRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServiceServer).DelAdminRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PermissionService_DelAdminRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServiceServer).DelAdminRole(ctx, req.(*DelAdminRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionService_DelRolePermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelRolePermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServiceServer).DelRolePermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PermissionService_DelRolePermission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServiceServer).DelRolePermission(ctx, req.(*DelRolePermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PermissionService_ServiceDesc is the grpc.ServiceDesc for PermissionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PermissionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "adminCenter.PermissionService",
	HandlerType: (*PermissionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddMenu",
			Handler:    _PermissionService_AddMenu_Handler,
		},
		{
			MethodName: "UpdateMenu",
			Handler:    _PermissionService_UpdateMenu_Handler,
		},
		{
			MethodName: "GetMenuList",
			Handler:    _PermissionService_GetMenuList_Handler,
		},
		{
			MethodName: "AddPermission",
			Handler:    _PermissionService_AddPermission_Handler,
		},
		{
			MethodName: "UpdatePermission",
			Handler:    _PermissionService_UpdatePermission_Handler,
		},
		{
			MethodName: "GetPermissionList",
			Handler:    _PermissionService_GetPermissionList_Handler,
		},
		{
			MethodName: "AddRole",
			Handler:    _PermissionService_AddRole_Handler,
		},
		{
			MethodName: "UpdateRole",
			Handler:    _PermissionService_UpdateRole_Handler,
		},
		{
			MethodName: "GetRoleList",
			Handler:    _PermissionService_GetRoleList_Handler,
		},
		{
			MethodName: "AddAdmin",
			Handler:    _PermissionService_AddAdmin_Handler,
		},
		{
			MethodName: "UpdateAdmin",
			Handler:    _PermissionService_UpdateAdmin_Handler,
		},
		{
			MethodName: "GetAdminList",
			Handler:    _PermissionService_GetAdminList_Handler,
		},
		{
			MethodName: "GetAdminRoleList",
			Handler:    _PermissionService_GetAdminRoleList_Handler,
		},
		{
			MethodName: "GetRolePermissionList",
			Handler:    _PermissionService_GetRolePermissionList_Handler,
		},
		{
			MethodName: "BindAdminRole",
			Handler:    _PermissionService_BindAdminRole_Handler,
		},
		{
			MethodName: "BindRolePermission",
			Handler:    _PermissionService_BindRolePermission_Handler,
		},
		{
			MethodName: "DelAdminRole",
			Handler:    _PermissionService_DelAdminRole_Handler,
		},
		{
			MethodName: "DelRolePermission",
			Handler:    _PermissionService_DelRolePermission_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admin.proto",
}