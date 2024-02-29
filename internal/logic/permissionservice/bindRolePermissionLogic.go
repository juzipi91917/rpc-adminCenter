package permissionservicelogic

import (
	"context"

	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-adminCenter/internal/svc"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-adminCenter/pb/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type BindRolePermissionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBindRolePermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BindRolePermissionLogic {
	return &BindRolePermissionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BindRolePermissionLogic) BindRolePermission(in *admin.BindRolePermissionRequest) (*admin.BindRolePermissionResponse, error) {
	// todo: add your logic here and delete this line

	return &admin.BindRolePermissionResponse{}, nil
}
