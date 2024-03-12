// Code generated by goctl. DO NOT EDIT.
// Source: admin.proto

package adminservice

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
	ChangeAdminStateRequest       = admin.ChangeAdminStateRequest
	ChangeAdminStateResponse      = admin.ChangeAdminStateResponse
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
	MenuInfo                      = admin.MenuInfo
	PermissionInfo                = admin.PermissionInfo
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

	AdminService interface {
		LogOn(ctx context.Context, in *LogOnRequest, opts ...grpc.CallOption) (*LogOnResponse, error)
		RefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...grpc.CallOption) (*RefreshTokenResponse, error)
		GetAdminMenuList(ctx context.Context, in *GetAdminMenuListRequest, opts ...grpc.CallOption) (*GetAdminMenuListResponse, error)
	}

	defaultAdminService struct {
		cli zrpc.Client
	}
)

func NewAdminService(cli zrpc.Client) AdminService {
	return &defaultAdminService{
		cli: cli,
	}
}

func (m *defaultAdminService) LogOn(ctx context.Context, in *LogOnRequest, opts ...grpc.CallOption) (*LogOnResponse, error) {
	client := admin.NewAdminServiceClient(m.cli.Conn())
	return client.LogOn(ctx, in, opts...)
}

func (m *defaultAdminService) RefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...grpc.CallOption) (*RefreshTokenResponse, error) {
	client := admin.NewAdminServiceClient(m.cli.Conn())
	return client.RefreshToken(ctx, in, opts...)
}

func (m *defaultAdminService) GetAdminMenuList(ctx context.Context, in *GetAdminMenuListRequest, opts ...grpc.CallOption) (*GetAdminMenuListResponse, error) {
	client := admin.NewAdminServiceClient(m.cli.Conn())
	return client.GetAdminMenuList(ctx, in, opts...)
}
