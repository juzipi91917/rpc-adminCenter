package permissionservicelogic

import (
	"context"

	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-adminCenter/internal/svc"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-adminCenter/pb/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRolePermissionListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRolePermissionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRolePermissionListLogic {
	return &GetRolePermissionListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRolePermissionListLogic) GetRolePermissionList(in *admin.GetRolePermissionListRequest) (*admin.GetRolePermissionListResponse, error) {
	// todo: add your logic here and delete this line

	return &admin.GetRolePermissionListResponse{}, nil
}
