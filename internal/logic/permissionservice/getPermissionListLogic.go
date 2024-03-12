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

type GetPermissionListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	PermissionDao *dao.PermissionDao
}

func NewGetPermissionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPermissionListLogic {
	return &GetPermissionListLogic{
		ctx:           ctx,
		svcCtx:        svcCtx,
		Logger:        logx.WithContext(ctx),
		PermissionDao: dao.NewPermissionDao(svcCtx),
	}
}

func (l *GetPermissionListLogic) GetPermissionList(in *admin.GetPermissionListRequest) (*admin.GetPermissionListResponse, error) {
	list, count, err := l.PermissionDao.GetPageList(l.ctx, in.Page, in.PageSize, &entity.Permission{State: in.State})

	if err != nil {
		l.Errorf("获取权限列表失败,err:%+v", err)
		return &admin.GetPermissionListResponse{}, result.NewRpcError("获取权限列表失败")
	}

	res := &admin.GetPermissionListResponse{
		Total: count,
		Page:  in.Page,
	}

	for _, pItem := range list {
		res.List = append(res.List, &admin.PermissionInfo{
			Id:        *pItem.Id,
			CreatedAt: *pItem.CreatedAt,
			UpdatedAt: *pItem.UpdatedAt,
			State:     *pItem.State,
			MenuId:    pItem.MenuId,
			Name:      pItem.Name,
			Route:     pItem.Route,
			Remarks:   pItem.Remarks,
		})
	}

	return res, nil
}
