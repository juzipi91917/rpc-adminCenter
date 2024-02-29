package permissionservicelogic

import (
	"context"

	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-adminCenter/internal/svc"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-adminCenter/pb/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAdminListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAdminListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAdminListLogic {
	return &GetAdminListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAdminListLogic) GetAdminList(in *admin.GetAdminListRequest) (*admin.GetAdminListResponse, error) {
	// todo: add your logic here and delete this line

	return &admin.GetAdminListResponse{}, nil
}
