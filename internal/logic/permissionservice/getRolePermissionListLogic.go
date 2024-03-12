package permissionservicelogic

import (
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/commonx/result"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-adminCenter/internal/dao"
	"context"

	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-adminCenter/internal/svc"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-adminCenter/pb/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRolePermissionListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	RoleDao *dao.RoleDao
}

func NewGetRolePermissionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRolePermissionListLogic {
	return &GetRolePermissionListLogic{
		ctx:     ctx,
		svcCtx:  svcCtx,
		Logger:  logx.WithContext(ctx),
		RoleDao: dao.NewRoleDao(svcCtx),
	}
}

func (l *GetRolePermissionListLogic) GetRolePermissionList(in *admin.GetRolePermissionListRequest) (*admin.GetRolePermissionListResponse, error) {
	if in.Id == nil {
		l.Errorf("获取权限列表失败,err:%+v", in)
		return &admin.GetRolePermissionListResponse{}, result.NewRpcError("获取权限列表失败, 缺少相关信息")
	}

	list, count, err := l.RoleDao.GetRolePermission(l.ctx, in.Id)

	if err != nil {
		l.Errorf("获取权限列表失败,err:%+v", err)
		return &admin.GetRolePermissionListResponse{}, result.NewRpcError("获取权限列表失败")
	}

	res := &admin.GetRolePermissionListResponse{
		Total: count,
	}

	// 循环赋值给res的list
	for _, item := range list {
		res.List = append(res.List, &admin.RolePermissionInfo{
			Id:      item.Id,
			State:   item.State,
			MenuId:  item.MenuId,
			Name:    item.Name,
			Route:   item.Route,
			Remarks: item.Remarks,
		})
	}

	return res, nil
}
