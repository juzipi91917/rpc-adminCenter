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

type BindRolePermissionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	RoleDao           *dao.RoleDao
	RolePermissionDao *dao.RolePermissionDao
}

func NewBindRolePermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BindRolePermissionLogic {
	return &BindRolePermissionLogic{
		ctx:               ctx,
		svcCtx:            svcCtx,
		Logger:            logx.WithContext(ctx),
		RoleDao:           dao.NewRoleDao(svcCtx),
		RolePermissionDao: dao.NewRolePermissionDao(svcCtx),
	}
}

func (l *BindRolePermissionLogic) BindRolePermission(in *admin.BindRolePermissionRequest) (*admin.BindRolePermissionResponse, error) {
	if in.RoleId == nil || in.PermissionId == nil {
		l.Errorf("绑定权限失败, 缺少相关信息 ,err:%+v", in)
		return &admin.BindRolePermissionResponse{}, result.NewRpcError("绑定权限失败, 缺少相关信息")
	}

	// 检查角色state
	role, err := l.RoleDao.GetInfo(l.ctx, &entity.Role{Id: in.RoleId})

	// sql错误或没查到角色或角色已被删除
	if err != nil || role.Id == nil {
		l.Errorf("绑定权限失败, 获取角色信息失败,err:%+v", err)
		return &admin.BindRolePermissionResponse{}, result.NewRpcError("绑定权限失败, 获取角色信息失败")
	}
	if *role.State == 0 {
		l.Errorf("绑定权限失败, 该角色未启用")
		return &admin.BindRolePermissionResponse{}, result.NewRpcError("绑定角色失败, 该权限未启用")
	}

	t := time.Now().Unix()

	_, err = l.RolePermissionDao.Add(l.ctx, &entity.RolePermission{
		CreatedAt:    &t,
		RoleId:       *in.RoleId,
		PermissionId: *in.PermissionId,
	})

	if err != nil {
		l.Errorf("绑定权限失败,err:%+v", err)
		return &admin.BindRolePermissionResponse{}, result.NewRpcError("绑定角色失败")

	}

	return &admin.BindRolePermissionResponse{
		Row: 1,
	}, nil
}
