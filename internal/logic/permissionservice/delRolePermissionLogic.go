package permissionservicelogic

import (
	"context"

	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-adminCenter/internal/svc"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-adminCenter/pb/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelRolePermissionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelRolePermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelRolePermissionLogic {
	return &DelRolePermissionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelRolePermissionLogic) DelRolePermission(in *admin.DelRolePermissionRequest) (*admin.DelRolePermissionResponse, error) {
	// todo: add your logic here and delete this line

	return &admin.DelRolePermissionResponse{}, nil
}
