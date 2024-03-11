package permissionservicelogic

import (
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/commonx/result"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-adminCenter/internal/dao"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-adminCenter/internal/model/entity"
	"context"
	"time"

	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-adminCenter/internal/svc"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-adminCenter/pb/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddPermissionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	PermissionDao *dao.PermissionDao
}

func NewAddPermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPermissionLogic {
	return &AddPermissionLogic{
		ctx:           ctx,
		svcCtx:        svcCtx,
		Logger:        logx.WithContext(ctx),
		PermissionDao: dao.NewPermissionDao(svcCtx),
	}
}

func (l *AddPermissionLogic) AddPermission(in *admin.AddPermissionRequest) (*admin.AddPermissionResponse, error) {

	// created&update time
	t := time.Now().Unix()

	id, err := l.PermissionDao.Add(l.ctx, &entity.Permission{
		CreatedAt: &t,
		UpdatedAt: &t,
		State:     &in.State,
		MenuId:    in.MenuId,
		Name:      in.Name,
		Remarks:   in.Remarks,
		Route:     in.Route,
	})

	if err != nil {
		return &admin.AddPermissionResponse{}, result.NewRpcError("添加权限失败")
	}

	return &admin.AddPermissionResponse{
		Id: *id,
	}, nil
}
