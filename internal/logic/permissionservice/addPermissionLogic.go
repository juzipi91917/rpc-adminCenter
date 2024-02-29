package permissionservicelogic

import (
	"context"

	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-adminCenter/internal/svc"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-adminCenter/pb/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddPermissionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddPermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPermissionLogic {
	return &AddPermissionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddPermissionLogic) AddPermission(in *admin.AddPermissionRequest) (*admin.AddPermissionResponse, error) {
	// todo: add your logic here and delete this line

	return &admin.AddPermissionResponse{}, nil
}
