// Code generated by goctl. DO NOT EDIT.
// Source: admin.proto

package permissionservice

import (
	"context"

	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-adminCenter/pb/admin"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddAdminRequest               = admin.AddAdminRequest
	AddAdminResponse              = admin.AddAdminResponse
	AddMenuRequest                = admin.AddMenuRequest
	AddMenuResponse               = admin.AddMenuResponse
	AddOperationLogRequest        = admin.AddOperationLogRequest
	AddPermissionRequest          = admin.AddPermissionRequest
	AddPermissionResponse         = admin.AddPermissionResponse
	AddRoleRequest                = admin.AddRoleRequest
	AddRoleResponse               = admin.AddRoleResponse
	AdminInfo                     = admin.AdminInfo
	BindAdminRoleRequest          = admin.BindAdminRoleRequest
	BindAdminRoleResponse         = admin.BindAdminRoleResponse
	BindRolePermissionRequest     = admin.BindRolePermissionRequest
	BindRolePermissionResponse    = admin.BindRolePermissionResponse
	DelAdminRoleRequest           = admin.DelAdminRoleRequest
	DelAdminRoleResponse          = admin.DelAdminRoleResponse
	DelRolePermissionRequest      = admin.DelRolePermissionRequest
	DelRolePermissionResponse     = admin.DelRolePermissionResponse
	Empty                         = admin.Empty
	GetAdminListRequest           = admin.GetAdminListRequest
	GetAdminListResponse          = admin.GetAdminListResponse
	GetAdminMenuListRequest       = admin.GetAdminMenuListRequest
	GetAdminMenuListResponse      = admin.GetAdminMenuListResponse
	GetAdminRoleListRequest       = admin.GetAdminRoleListRequest
	GetAdminRoleListResponse      = admin.GetAdminRoleListResponse
	GetMenuListRequest            = admin.GetMenuListRequest
	GetMenuListResponse           = admin.GetMenuListResponse
	GetPermissionListRequest      = admin.GetPermissionListRequest
	GetPermissionListResponse     = admin.GetPermissionListResponse
	GetRoleListRequest            = admin.GetRoleListRequest
	GetRoleListResponse           = admin.GetRoleListResponse
	GetRolePermissionListRequest  = admin.GetRolePermissionListRequest
	GetRolePermissionListResponse = admin.GetRolePermissionListResponse
	LogOnRequest                  = admin.LogOnRequest
	LogOnResponse                 = admin.LogOnResponse
	RefreshTokenRequest           = admin.RefreshTokenRequest
	RefreshTokenResponse          = admin.RefreshTokenResponse
	UpdateAdminRequest            = admin.UpdateAdminRequest
	UpdateAdminResponse           = admin.UpdateAdminResponse
	UpdateMenuRequest             = admin.UpdateMenuRequest
	UpdateMenuResponse            = admin.UpdateMenuResponse
	UpdateOperationLogRequest     = admin.UpdateOperationLogRequest
	UpdatePermissionRequest       = admin.UpdatePermissionRequest
	UpdatePermissionResponse      = admin.UpdatePermissionResponse
	UpdateRoleRequest             = admin.UpdateRoleRequest
	UpdateRoleResponse            = admin.UpdateRoleResponse

	PermissionService interface {
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

	defaultPermissionService struct {
		cli zrpc.Client
	}
)

func NewPermissionService(cli zrpc.Client) PermissionService {
	return &defaultPermissionService{
		cli: cli,
	}
}

func (m *defaultPermissionService) AddMenu(ctx context.Context, in *AddMenuRequest, opts ...grpc.CallOption) (*AddMenuResponse, error) {
	client := admin.NewPermissionServiceClient(m.cli.Conn())
	return client.AddMenu(ctx, in, opts...)
}

func (m *defaultPermissionService) UpdateMenu(ctx context.Context, in *UpdateMenuRequest, opts ...grpc.CallOption) (*UpdateMenuResponse, error) {
	client := admin.NewPermissionServiceClient(m.cli.Conn())
	return client.UpdateMenu(ctx, in, opts...)
}

func (m *defaultPermissionService) GetMenuList(ctx context.Context, in *GetMenuListRequest, opts ...grpc.CallOption) (*GetMenuListResponse, error) {
	client := admin.NewPermissionServiceClient(m.cli.Conn())
	return client.GetMenuList(ctx, in, opts...)
}

func (m *defaultPermissionService) AddPermission(ctx context.Context, in *AddPermissionRequest, opts ...grpc.CallOption) (*AddPermissionResponse, error) {
	client := admin.NewPermissionServiceClient(m.cli.Conn())
	return client.AddPermission(ctx, in, opts...)
}

func (m *defaultPermissionService) UpdatePermission(ctx context.Context, in *UpdatePermissionRequest, opts ...grpc.CallOption) (*UpdatePermissionResponse, error) {
	client := admin.NewPermissionServiceClient(m.cli.Conn())
	return client.UpdatePermission(ctx, in, opts...)
}

func (m *defaultPermissionService) GetPermissionList(ctx context.Context, in *GetPermissionListRequest, opts ...grpc.CallOption) (*GetPermissionListResponse, error) {
	client := admin.NewPermissionServiceClient(m.cli.Conn())
	return client.GetPermissionList(ctx, in, opts...)
}

func (m *defaultPermissionService) AddRole(ctx context.Context, in *AddRoleRequest, opts ...grpc.CallOption) (*AddRoleResponse, error) {
	client := admin.NewPermissionServiceClient(m.cli.Conn())
	return client.AddRole(ctx, in, opts...)
}

func (m *defaultPermissionService) UpdateRole(ctx context.Context, in *UpdateRoleRequest, opts ...grpc.CallOption) (*UpdateRoleResponse, error) {
	client := admin.NewPermissionServiceClient(m.cli.Conn())
	return client.UpdateRole(ctx, in, opts...)
}

func (m *defaultPermissionService) GetRoleList(ctx context.Context, in *GetRoleListRequest, opts ...grpc.CallOption) (*GetRoleListResponse, error) {
	client := admin.NewPermissionServiceClient(m.cli.Conn())
	return client.GetRoleList(ctx, in, opts...)
}

func (m *defaultPermissionService) AddAdmin(ctx context.Context, in *AddAdminRequest, opts ...grpc.CallOption) (*AddAdminResponse, error) {
	client := admin.NewPermissionServiceClient(m.cli.Conn())
	return client.AddAdmin(ctx, in, opts...)
}

func (m *defaultPermissionService) UpdateAdmin(ctx context.Context, in *UpdateAdminRequest, opts ...grpc.CallOption) (*UpdateAdminResponse, error) {
	client := admin.NewPermissionServiceClient(m.cli.Conn())
	return client.UpdateAdmin(ctx, in, opts...)
}

func (m *defaultPermissionService) GetAdminList(ctx context.Context, in *GetAdminListRequest, opts ...grpc.CallOption) (*GetAdminListResponse, error) {
	client := admin.NewPermissionServiceClient(m.cli.Conn())
	return client.GetAdminList(ctx, in, opts...)
}

func (m *defaultPermissionService) GetAdminRoleList(ctx context.Context, in *GetAdminRoleListRequest, opts ...grpc.CallOption) (*GetAdminRoleListResponse, error) {
	client := admin.NewPermissionServiceClient(m.cli.Conn())
	return client.GetAdminRoleList(ctx, in, opts...)
}

func (m *defaultPermissionService) GetRolePermissionList(ctx context.Context, in *GetRolePermissionListRequest, opts ...grpc.CallOption) (*GetRolePermissionListResponse, error) {
	client := admin.NewPermissionServiceClient(m.cli.Conn())
	return client.GetRolePermissionList(ctx, in, opts...)
}

func (m *defaultPermissionService) BindAdminRole(ctx context.Context, in *BindAdminRoleRequest, opts ...grpc.CallOption) (*BindAdminRoleResponse, error) {
	client := admin.NewPermissionServiceClient(m.cli.Conn())
	return client.BindAdminRole(ctx, in, opts...)
}

func (m *defaultPermissionService) BindRolePermission(ctx context.Context, in *BindRolePermissionRequest, opts ...grpc.CallOption) (*BindRolePermissionResponse, error) {
	client := admin.NewPermissionServiceClient(m.cli.Conn())
	return client.BindRolePermission(ctx, in, opts...)
}

func (m *defaultPermissionService) DelAdminRole(ctx context.Context, in *DelAdminRoleRequest, opts ...grpc.CallOption) (*DelAdminRoleResponse, error) {
	client := admin.NewPermissionServiceClient(m.cli.Conn())
	return client.DelAdminRole(ctx, in, opts...)
}

func (m *defaultPermissionService) DelRolePermission(ctx context.Context, in *DelRolePermissionRequest, opts ...grpc.CallOption) (*DelRolePermissionResponse, error) {
	client := admin.NewPermissionServiceClient(m.cli.Conn())
	return client.DelRolePermission(ctx, in, opts...)
}
