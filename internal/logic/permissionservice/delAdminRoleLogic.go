package permissionservicelogic

import (
	"context"

	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-adminCenter/internal/svc"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-adminCenter/pb/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelAdminRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelAdminRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelAdminRoleLogic {
	return &DelAdminRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelAdminRoleLogic) DelAdminRole(in *admin.DelAdminRoleRequest) (*admin.DelAdminRoleResponse, error) {
	// todo: add your logic here and delete this line

	return &admin.DelAdminRoleResponse{}, nil
}
