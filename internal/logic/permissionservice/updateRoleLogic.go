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

type UpdateRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	RoleDao *dao.RoleDao
}

func NewUpdateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleLogic {
	return &UpdateRoleLogic{
		ctx:     ctx,
		svcCtx:  svcCtx,
		Logger:  logx.WithContext(ctx),
		RoleDao: dao.NewRoleDao(svcCtx),
	}
}

func (l *UpdateRoleLogic) UpdateRole(in *admin.UpdateRoleRequest) (*admin.UpdateRoleResponse, error) {
	if in.Id == nil {
		l.Errorf("更新角色失败,输入参数不正确:%+v", in)
		return &admin.UpdateRoleResponse{}, result.NewRpcError("更新角色失败, 输入参数不正确")
	}

	where := &entity.Role{}

	if in.State != nil {
		where.State = in.State
	}
	if in.Name != nil {
		where.Name = *in.Name
	}
	if in.Remarks != nil {
		where.Remarks = *in.Remarks
	}

	t := time.Now().Unix()
	where.UpdatedAt = &t

	row, err := l.RoleDao.Update(l.ctx, &entity.Role{Id: in.Id}, where)

	if err != nil {
		l.Errorf("更新角色失败,err:%+v", err)
		return &admin.UpdateRoleResponse{}, result.NewRpcError("更新角色失败")
	}

	return &admin.UpdateRoleResponse{Row: row}, nil
}
