package permissionservicelogic

import (
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/commonx/result"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-adminCenter/internal/dao"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-adminCenter/internal/model/entity"
	"context"

	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-adminCenter/internal/svc"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-adminCenter/pb/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelRolePermissionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	RolePermissionDao *dao.RolePermissionDao
}

func NewDelRolePermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelRolePermissionLogic {
	return &DelRolePermissionLogic{
		ctx:               ctx,
		svcCtx:            svcCtx,
		Logger:            logx.WithContext(ctx),
		RolePermissionDao: dao.NewRolePermissionDao(svcCtx),
	}
}

func (l *DelRolePermissionLogic) DelRolePermission(in *admin.DelRolePermissionRequest) (*admin.DelRolePermissionResponse, error) {

	if in.PermissionId == nil || in.RoleId == nil {
		l.Errorf("删除角色权限失败, 缺少相关信息 ,err:%+v", in)
		return &admin.DelRolePermissionResponse{}, result.NewRpcError("删除角色权限失败, 缺少相关信息")
	}

	row, err := l.RolePermissionDao.Delete(l.ctx, &entity.RolePermission{PermissionId: *in.PermissionId, RoleId: *in.RoleId})

	if err != nil {
		l.Errorf("删除权限失败,err:%+v", err)
		return &admin.DelRolePermissionResponse{}, result.NewRpcError("删除权限失败")
	}

	return &admin.DelRolePermissionResponse{Row: row}, nil
}
