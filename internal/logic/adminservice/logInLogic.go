package adminservicelogic

import (
	"context"

	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-adminCenter/internal/svc"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-adminCenter/pb/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogInLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLogInLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogInLogic {
	return &LogInLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LogInLogic) LogIn(in *admin.LogInRequest) (*admin.LogInResponse, error) {
	// todo: add your logic here and delete this line

	return &admin.LogInResponse{}, nil
}
