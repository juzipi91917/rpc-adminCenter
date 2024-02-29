package adminservicelogic

import (
	"context"

	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-adminCenter/internal/svc"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-adminCenter/pb/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogOnLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLogOnLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogOnLogic {
	return &LogOnLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LogOnLogic) LogOn(in *admin.LogOnRequest) (*admin.LogOnResponse, error) {
	// todo: add your logic here and delete this line

	return &admin.LogOnResponse{}, nil
}
