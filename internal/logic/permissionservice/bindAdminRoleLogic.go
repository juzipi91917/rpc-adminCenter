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

type BindAdminRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	AdminRoleDao *dao.AdminRoleDao
	AdminDao     *dao.AdminDao
}

func NewBindAdminRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BindAdminRoleLogic {
	return &BindAdminRoleLogic{
		ctx:          ctx,
		svcCtx:       svcCtx,
		Logger:       logx.WithContext(ctx),
		AdminRoleDao: dao.NewAdminRoleDao(svcCtx),
		AdminDao:     dao.NewAdminDao(svcCtx),
	}
}

func (l *BindAdminRoleLogic) BindAdminRole(in *admin.BindAdminRoleRequest) (*admin.BindAdminRoleResponse, error) {

	if in.AdminId == nil || in.RoleId == nil {
		l.Errorf("绑定角色失败, 缺少相关信息 ,err:%+v", in)
		return &admin.BindAdminRoleResponse{}, result.NewRpcError("绑定角色失败, 缺少相关信息")
	}

	// 检查角色state以及is_deleted
	user, err := l.AdminDao.GetInfo(l.ctx, &entity.Admin{Id: in.AdminId})

	// sql错误或没查到用户或用户已被删除
	if err != nil || user.Id == nil || *user.IsDeleted == 1 {
		l.Errorf("绑定角色失败, 获取用户信息失败,err:%+v", err)
		return &admin.BindAdminRoleResponse{}, result.NewRpcError("绑定角色失败, 获取用户信息失败")
	}
	if *user.State == 0 {
		l.Errorf("绑定角色失败, 用户已被禁用")
		return &admin.BindAdminRoleResponse{}, result.NewRpcError("绑定角色失败, 用户已被禁用")
	}

	t := time.Now().Unix()

	_, err = l.AdminRoleDao.Add(l.ctx, &entity.AdminRole{
		CreatedAt: &t,
		AdminId:   *in.AdminId,
		RoleId:    *in.RoleId,
	})

	if err != nil {
		l.Errorf("绑定角色失败,err:%+v", err)
		return &admin.BindAdminRoleResponse{}, result.NewRpcError("绑定角色失败")

	}

	return &admin.BindAdminRoleResponse{
		Row: 1,
	}, nil
}
