package permissionservicelogic

import (
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/commonx/result"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-adminCenter/internal/dao"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-adminCenter/internal/svc"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-adminCenter/pb/admin"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetAdminRoleListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	AdminDao *dao.AdminDao
}

func NewGetAdminRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAdminRoleListLogic {
	return &GetAdminRoleListLogic{
		ctx:      ctx,
		svcCtx:   svcCtx,
		Logger:   logx.WithContext(ctx),
		AdminDao: dao.NewAdminDao(svcCtx),
	}
}

func (l *GetAdminRoleListLogic) GetAdminRoleList(in *admin.GetAdminRoleListRequest) (*admin.GetAdminRoleListResponse, error) {
	if in.Id == nil {
		l.Errorf("获取角色列表失败,err:%+v", in)
		return &admin.GetAdminRoleListResponse{}, result.NewRpcError("获取角色列表失败, 缺少相关信息")
	}

	list, count, err := l.AdminDao.GetAdminRoleList(l.ctx, in.Id)

	if err != nil {
		l.Errorf("获取角色列表失败,err:%+v", err)
		return &admin.GetAdminRoleListResponse{}, result.NewRpcError("获取角色列表失败")
	}

	res := &admin.GetAdminRoleListResponse{
		Total: count,
	}

	// 循环赋值给res的list
	for _, item := range list {
		res.List = append(res.List, &admin.AdminRoleInfo{
			Id:      item.ID,
			Name:    item.Name,
			Remarks: item.Remarks,
			State:   item.State,
		})
	}
	return res, nil
}
