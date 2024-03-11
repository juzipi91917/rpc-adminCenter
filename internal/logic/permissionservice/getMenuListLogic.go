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

type GetMenuListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	MenuDao *dao.MenuDao
}

func NewGetMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuListLogic {
	return &GetMenuListLogic{
		ctx:     ctx,
		svcCtx:  svcCtx,
		Logger:  logx.WithContext(ctx),
		MenuDao: dao.NewMenuDao(svcCtx),
	}
}

func (l *GetMenuListLogic) GetMenuList(in *admin.GetMenuListRequest) (*admin.GetMenuListResponse, error) {
	list, count, err := l.MenuDao.GetPageList(l.ctx, in.GetPage(), in.GetPageSize(), &entity.Menu{State: &in.Status})

	if err != nil {
		l.Errorf("获取菜单列表失败,err:%+v", err)
		return &admin.GetMenuListResponse{}, result.NewRpcError("获取菜单列表失败")
	}

	res := &admin.GetMenuListResponse{
		Total: count,
		Page:  in.Page,
	}

	for _, menuItem := range list {
		res.List = append(res.List, &admin.MenuInfo{
			Id:        *menuItem.Id,
			Name:      menuItem.Name,
			CreatedAt: *menuItem.CreatedAt,
			UpdatedAt: *menuItem.UpdatedAt,
			State:     *menuItem.State,
			Route:     menuItem.Route,
			PId:       menuItem.PId,
			LevelIds:  menuItem.LevelIds,
			Remarks:   menuItem.Remarks,
		})
	}

	return res, nil
}
