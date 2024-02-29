package permissionservicelogic

import (
	"context"

	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-adminCenter/internal/svc"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-adminCenter/pb/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type BindAdminRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBindAdminRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BindAdminRoleLogic {
	return &BindAdminRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BindAdminRoleLogic) BindAdminRole(in *admin.BindAdminRoleRequest) (*admin.BindAdminRoleResponse, error) {
	// todo: add your logic here and delete this line

	return &admin.BindAdminRoleResponse{}, nil
}
