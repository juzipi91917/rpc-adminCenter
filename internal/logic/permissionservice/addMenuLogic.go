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

type AddMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	MenuDao *dao.MenuDao
}

func NewAddMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMenuLogic {
	return &AddMenuLogic{
		ctx:     ctx,
		svcCtx:  svcCtx,
		Logger:  logx.WithContext(ctx),
		MenuDao: dao.NewMenuDao(svcCtx),
	}
}

func (l *AddMenuLogic) AddMenu(in *admin.AddMenuRequest) (*admin.AddMenuResponse, error) {

	t := time.Now().Unix()

	id, err := l.MenuDao.Add(l.ctx, &entity.Menu{
		Name:      in.Name,
		CreatedAt: &t,
		UpdatedAt: &t,
		Route:     in.Route,
		PId:       in.Pid,
		LevelIds:  in.LevelIds,
		Remarks:   in.Remarks,
	})

	if err != nil {
		return &admin.AddMenuResponse{}, result.NewRpcError("添加菜单失败")
	}

	return &admin.AddMenuResponse{
		Id: *id,
	}, nil
}
