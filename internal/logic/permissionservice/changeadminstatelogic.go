package permissionservicelogic

import (
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/commonx/result"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-adminCenter/internal/dao"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-adminCenter/internal/svc"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-adminCenter/pb/admin"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangeAdminStateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	AdminDao *dao.AdminDao
}

func NewChangeAdminStateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangeAdminStateLogic {
	return &ChangeAdminStateLogic{
		ctx:      ctx,
		svcCtx:   svcCtx,
		Logger:   logx.WithContext(ctx),
		AdminDao: dao.NewAdminDao(svcCtx),
	}
}

func (l *ChangeAdminStateLogic) ChangeAdminState(in *admin.ChangeAdminStateRequest) (*admin.ChangeAdminStateResponse, error) {
	if in.Id == nil {
		l.Errorf("修改用户状态失败,输入参数不正确:%+v", in)
		return &admin.ChangeAdminStateResponse{}, result.NewRpcError("修改用户状态失败, 输入参数不正确")
	}

	row, err := l.AdminDao.ChangeState(l.ctx, *in.Id)
	if err != nil {
		return nil, err
	}

	if err != nil {
		l.Errorf("修改用户状态失败,err:%+v", err)
		return &admin.ChangeAdminStateResponse{}, result.NewRpcError("修改用户状态失败")
	}

	return &admin.ChangeAdminStateResponse{Row: row}, nil
}
