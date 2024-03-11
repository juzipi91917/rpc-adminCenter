// Code generated by goctl. DO NOT EDIT.
// Source: admin.proto

package operationlogservice

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
	MenuInfo                      = admin.MenuInfo
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

	OperationLogService interface {
		AddOperationLog(ctx context.Context, in *AddOperationLogRequest, opts ...grpc.CallOption) (*Empty, error)
		UpdateOperationLog(ctx context.Context, in *UpdateOperationLogRequest, opts ...grpc.CallOption) (*Empty, error)
	}

	defaultOperationLogService struct {
		cli zrpc.Client
	}
)

func NewOperationLogService(cli zrpc.Client) OperationLogService {
	return &defaultOperationLogService{
		cli: cli,
	}
}

func (m *defaultOperationLogService) AddOperationLog(ctx context.Context, in *AddOperationLogRequest, opts ...grpc.CallOption) (*Empty, error) {
	client := admin.NewOperationLogServiceClient(m.cli.Conn())
	return client.AddOperationLog(ctx, in, opts...)
}

func (m *defaultOperationLogService) UpdateOperationLog(ctx context.Context, in *UpdateOperationLogRequest, opts ...grpc.CallOption) (*Empty, error) {
	client := admin.NewOperationLogServiceClient(m.cli.Conn())
	return client.UpdateOperationLog(ctx, in, opts...)
}
