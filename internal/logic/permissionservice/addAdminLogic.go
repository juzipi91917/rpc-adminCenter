package permissionservicelogic

import (
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/commonx/result"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-adminCenter/internal/dao"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-adminCenter/internal/model/entity"
	"context"
	"golang.org/x/crypto/bcrypt"
	"time"

	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-adminCenter/internal/svc"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-adminCenter/pb/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddAdminLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	AdminDao *dao.AdminDao
}

func NewAddAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddAdminLogic {
	return &AddAdminLogic{
		ctx:      ctx,
		svcCtx:   svcCtx,
		Logger:   logx.WithContext(ctx),
		AdminDao: dao.NewAdminDao(svcCtx),
	}
}

func (l *AddAdminLogic) AddAdmin(in *admin.AddAdminRequest) (*admin.AddAdminResponse, error) {
	pwdLen := len(in.Password)
	if pwdLen <= 0 || pwdLen > 255 {
		return nil, result.NewRpcError("密码长度不合法")
	}

	password := in.Password

	// 以下两行为 `bcrypt` 加密 如果需要的话
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	password = string(hashedPassword)

	// 时间戳
	t := time.Now().Unix()

	id, err := l.AdminDao.Add(l.ctx, &entity.Admin{
		Id:        &in.Id,
		CreatedAt: &t,
		UpdatedAt: &t,
		State:     &in.State,
		Account:   in.Account,
		Password:  password,
		Mobile:    in.Mobile,
		Name:      in.Name,
	})
	if err != nil {
		return nil, result.NewRpcError("添加用户失败")
	}

	return &admin.AddAdminResponse{
		Id: *id,
	}, nil
}
