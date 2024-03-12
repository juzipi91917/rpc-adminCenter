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

type AddRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	RoleDao *dao.RoleDao
}

func NewAddRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddRoleLogic {
	return &AddRoleLogic{
		ctx:     ctx,
		svcCtx:  svcCtx,
		Logger:  logx.WithContext(ctx),
		RoleDao: dao.NewRoleDao(svcCtx),
	}
}

func (l *AddRoleLogic) AddRole(in *admin.AddRoleRequest) (*admin.AddRoleResponse, error) {

	t := time.Now().Unix()

	id, err := l.RoleDao.Add(l.ctx, &entity.Role{
		CreatedAt: &t,
		UpdatedAt: &t,
		State:     &in.State,
		Name:      in.Name,
		Remarks:   in.Remarks,
	})

	if err != nil {
		l.Errorf("添加角色失败,err:%+v", err)
		return &admin.AddRoleResponse{}, result.NewRpcError("添加角色失败")
	}

	return &admin.AddRoleResponse{
		Id: *id,
	}, nil
}
