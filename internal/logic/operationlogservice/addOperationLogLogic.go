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

type AddOperationLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	OperationLogDao *dao.OperationLogDao
}

func NewAddOperationLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddOperationLogLogic {
	return &AddOperationLogLogic{
		ctx:             ctx,
		svcCtx:          svcCtx,
		Logger:          logx.WithContext(ctx),
		OperationLogDao: dao.NewOperationLogDao(svcCtx),
	}
}

func (l *AddOperationLogLogic) AddOperationLog(in *admin.AddOperationLogRequest) (*admin.Empty, error) {
	_, err := l.OperationLogDao.Add(l.ctx, &entity.OperationLog{
		CreatedAt:     tool.GetTimeStamp(),
		UpdatedAt:     tool.GetTimeStamp(),
		OperationType: in.OperationType,
		AdminId:       in.AdminId,
		OperatorName:  in.OperatorName,
		OperationAt:   tool.GetTimeStamp(),
		OperationDesc: in.OperationDesc,
		ModuleName:    in.ModuleName,
		Route:         in.Route,
		OperatorIp:    in.OperatorIp,
		Content:       in.Content,
		State:         in.State,
		ErrMsg:        in.ErrMsg,
	})
	if err != nil {
		l.Error("写入操作记录失败,err:%+v", err)
	}
	return &admin.Empty{}, nil
}
