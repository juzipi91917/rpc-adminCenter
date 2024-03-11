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

type UpdatePermissionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	PermissionDao *dao.PermissionDao
}

func NewUpdatePermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePermissionLogic {
	return &UpdatePermissionLogic{
		ctx:           ctx,
		svcCtx:        svcCtx,
		Logger:        logx.WithContext(ctx),
		PermissionDao: dao.NewPermissionDao(svcCtx),
	}
}

func (l *UpdatePermissionLogic) UpdatePermission(in *admin.UpdatePermissionRequest) (*admin.UpdatePermissionResponse, error) {

	if in.Id == nil {
		l.Errorf("更新权限失败,输入参数不正确:%+v", in)
		return &admin.UpdatePermissionResponse{}, result.NewRpcError("更新权限失败,输入参数不正确")
	}

	where := &entity.Permission{}

	if in.State != nil {
		where.State = in.State
	}
	if in.MenuId != nil {
		where.MenuId = *in.MenuId
	}
	if in.Name != nil {
		where.Name = *in.Name
	}
	if in.Route != nil {
		where.Route = *in.Route
	}
	if in.Remarks != nil {
		where.Remarks = *in.Remarks
	}

	// set update time
	t := time.Now().Unix()
	where.UpdatedAt = &t

	row, err := l.PermissionDao.Update(l.ctx, &entity.Permission{Id: in.Id}, where)

	if err != nil {
		l.Errorf("更新权限失败,err:%+v", err)
		return &admin.UpdatePermissionResponse{}, result.NewRpcError("更新权限失败")
	}

	return &admin.UpdatePermissionResponse{Row: row}, nil
}
