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

type DelAdminRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	AdminRoleDao *dao.AdminRoleDao
}

func NewDelAdminRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelAdminRoleLogic {
	return &DelAdminRoleLogic{
		ctx:          ctx,
		svcCtx:       svcCtx,
		Logger:       logx.WithContext(ctx),
		AdminRoleDao: dao.NewAdminRoleDao(svcCtx),
	}
}

func (l *DelAdminRoleLogic) DelAdminRole(in *admin.DelAdminRoleRequest) (*admin.DelAdminRoleResponse, error) {

	if in.AdminId == nil || in.RoleId == nil {
		l.Errorf("删除角色失败, 缺少相关信息 ,err:%+v", in)
		return &admin.DelAdminRoleResponse{}, result.NewRpcError("删除角色失败, 缺少相关信息")
	}

	row, err := l.AdminRoleDao.Delete(l.ctx, &entity.AdminRole{AdminId: *in.AdminId, RoleId: *in.RoleId})

	if err != nil {
		l.Errorf("删除角色失败,err:%+v", err)
		return &admin.DelAdminRoleResponse{}, result.NewRpcError("删除角色失败")
	}

	return &admin.DelAdminRoleResponse{Row: row}, nil
}
