package permissionservicelogic

import (
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/commonx/result"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-adminCenter/internal/dao"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-adminCenter/internal/model/entity"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-adminCenter/internal/svc"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-adminCenter/pb/admin"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetAdminListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	AdminDao *dao.AdminDao
}

func NewGetAdminListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAdminListLogic {
	return &GetAdminListLogic{
		ctx:      ctx,
		svcCtx:   svcCtx,
		Logger:   logx.WithContext(ctx),
		AdminDao: dao.NewAdminDao(svcCtx),
	}
}

func (l *GetAdminListLogic) GetAdminList(in *admin.GetAdminListRequest) (*admin.GetAdminListResponse, error) {

	list, count, err := l.AdminDao.GetPageList(l.ctx, in.GetPage(), in.GetPageSize(), &entity.Admin{State: &in.State})

	if err != nil {
		l.Errorf("获取管理员列表失败,err:%+v", err)
		return &admin.GetAdminListResponse{}, result.NewRpcError("获取管理员列表失败")
	}

	res := &admin.GetAdminListResponse{
		Total: count,
		Page:  in.GetPage(),
	}

	for _, a := range list {
		res.List = append(res.List, &admin.AdminInfo{
			Id:       *a.Id,
			Name:     a.Name,
			CreateAt: *a.CreatedAt,
			UpdateAt: *a.UpdatedAt,
			State:    *a.State,
			Account:  a.Account,
			Password: a.Password,
			Mobile:   a.Mobile,
		})
	}

	return res, nil

}
