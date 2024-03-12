package permissionservicelogic

import (
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/commonx/result"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-adminCenter/internal/dao"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-adminCenter/internal/model/entity"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-adminCenter/internal/svc"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-adminCenter/pb/admin"
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRoleListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	RoleDao *dao.RoleDao
}

func NewGetRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoleListLogic {
	return &GetRoleListLogic{
		ctx:     ctx,
		svcCtx:  svcCtx,
		Logger:  logx.WithContext(ctx),
		RoleDao: dao.NewRoleDao(svcCtx),
	}
}

func (l *GetRoleListLogic) GetRoleList(in *admin.GetRoleListRequest) (*admin.GetRoleListResponse, error) {
	fmt.Println(in.State)
	list, count, err := l.RoleDao.GetPageList(l.ctx, in.GetPage(), in.GetPageSize(), &entity.Role{State: in.State})

	if err != nil {
		l.Errorf("获取角色列表失败,err:%+v", err)
		return &admin.GetRoleListResponse{}, result.NewRpcError("获取角色列表失败")
	}

	res := &admin.GetRoleListResponse{
		Total: count,
		Page:  in.Page,
	}

	for _, roleItem := range list {
		res.List = append(res.List, &admin.RoleInfo{
			Id:        *roleItem.Id,
			Name:      roleItem.Name,
			CreatedAt: *roleItem.CreatedAt,
			UpdatedAt: *roleItem.UpdatedAt,
			State:     *roleItem.State,
			Remarks:   roleItem.Remarks,
		})
	}

	return res, nil
}
