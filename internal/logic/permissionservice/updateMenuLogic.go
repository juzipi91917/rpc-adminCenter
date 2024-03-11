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

type UpdateMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	MenuDao *dao.MenuDao
}

func NewUpdateMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMenuLogic {
	return &UpdateMenuLogic{
		ctx:     ctx,
		svcCtx:  svcCtx,
		Logger:  logx.WithContext(ctx),
		MenuDao: dao.NewMenuDao(svcCtx),
	}
}

func (l *UpdateMenuLogic) UpdateMenu(in *admin.UpdateMenuRequest) (*admin.UpdateMenuResponse, error) {
	if in.Id == nil {
		l.Errorf("更新菜单失败,输入参数不正确:%+v", in)
		return &admin.UpdateMenuResponse{}, result.NewRpcError("更新菜单失败,输入参数不正确")
	}

	where := &entity.Menu{}

	if in.State != nil {
		where.State = in.State
	}
	if in.Name != nil {
		where.Name = *in.Name
	}
	if in.Route != nil {
		where.Route = *in.Route
	}
	if in.PId != nil {
		where.PId = *in.PId
	}
	if in.LevelIds != nil {
		where.LevelIds = *in.LevelIds
	}
	if in.Remarks != nil {
		where.Remarks = *in.Remarks
	}

	// set update time
	t := time.Now().Unix()
	where.UpdatedAt = &t

	row, err := l.MenuDao.Update(l.ctx, &entity.Menu{Id: in.Id}, where)

	if err != nil {
		l.Errorf("更新信息失败, err:%+v", err)
		return &admin.UpdateMenuResponse{}, result.NewRpcError("更新信息失败, 输入参数不正确")
	}

	return &admin.UpdateMenuResponse{Row: row}, nil

}
