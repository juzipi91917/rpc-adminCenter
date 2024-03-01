package operationlogservicelogic

import (
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/commonx/tool"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-adminCenter/internal/dao"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-adminCenter/internal/model/entity"
	"context"

	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-adminCenter/internal/svc"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-adminCenter/pb/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOperationLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	OperationLogDao *dao.OperationLogDao
}

func NewUpdateOperationLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOperationLogLogic {
	return &UpdateOperationLogLogic{
		ctx:             ctx,
		svcCtx:          svcCtx,
		Logger:          logx.WithContext(ctx),
		OperationLogDao: dao.NewOperationLogDao(svcCtx),
	}
}

func (l *UpdateOperationLogLogic) UpdateOperationLog(in *admin.UpdateOperationLogRequest) (*admin.Empty, error) {
	if in.Id == 0 {
		l.Errorf("更新操作日志失败,传入数据不正确,data:%+v", in)
		return nil, nil
	}
	_, err := l.OperationLogDao.Update(&entity.OperationLog{Id: &in.Id}, &entity.OperationLog{
		UpdatedAt: tool.GetTimeStamp(),
		State:     in.State,
		ErrMsg:    in.ErrMsg,
	})
	if err != nil {
		l.Errorf("更新操作日志失败,更新数据库失败,err:%+v", err)
		return nil, nil
	}
	return &admin.Empty{}, nil
}
