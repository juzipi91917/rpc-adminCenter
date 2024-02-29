package adminservicelogic

import (
	"context"

	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-adminCenter/internal/svc"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-adminCenter/pb/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAdminMenuListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAdminMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAdminMenuListLogic {
	return &GetAdminMenuListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAdminMenuListLogic) GetAdminMenuList(in *admin.GetAdminMenuListRequest) (*admin.GetAdminMenuListResponse, error) {
	// todo: add your logic here and delete this line

	return &admin.GetAdminMenuListResponse{}, nil
}
