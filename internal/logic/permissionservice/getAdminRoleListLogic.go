package permissionservicelogic

import (
	"context"

	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-adminCenter/internal/svc"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-adminCenter/pb/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAdminRoleListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAdminRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAdminRoleListLogic {
	return &GetAdminRoleListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAdminRoleListLogic) GetAdminRoleList(in *admin.GetAdminRoleListRequest) (*admin.GetAdminRoleListResponse, error) {
	// todo: add your logic here and delete this line

	return &admin.GetAdminRoleListResponse{}, nil
}
