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

type UpdateAdminLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	AdminDao *dao.AdminDao
}

func NewUpdateAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAdminLogic {
	return &UpdateAdminLogic{
		ctx:      ctx,
		svcCtx:   svcCtx,
		Logger:   logx.WithContext(ctx),
		AdminDao: dao.NewAdminDao(svcCtx),
	}
}

func (l *UpdateAdminLogic) UpdateAdmin(in *admin.UpdateAdminRequest) (*admin.UpdateAdminResponse, error) {

	if in.Id == nil {
		l.Errorf("更新信息失败,输入参数不正确:%+v", in)
		return &admin.UpdateAdminResponse{}, result.NewRpcError("更新信息失败, 输入参数不正确")
	}

	where := &entity.Admin{}

	if in.State != nil {
		where.State = in.State
	}
	if in.Account != nil {
		where.Account = *in.Account
	}
	if in.Password != nil {
		password := *in.Password

		// bcrypt 加密 如果需要的话
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		password = string(hashedPassword)
		where.Password = password
	}
	if in.Mobile != nil {
		where.Mobile = *in.Mobile
	}
	if in.Name != nil {
		where.Name = *in.Name
	}

	// 逻辑删除
	if in.IsDeleted != nil {
		where.IsDeleted = in.IsDeleted
	}

	// set update time
	t := time.Now().Unix()
	where.UpdatedAt = &t

	row, err := l.AdminDao.Update(l.ctx, &entity.Admin{Id: in.Id}, where)

	if err != nil {
		l.Errorf("更新信息失败, err:%+v", err)
		return &admin.UpdateAdminResponse{}, result.NewRpcError("更新信息失败, 输入参数不正确")
	}

	return &admin.UpdateAdminResponse{Row: row}, nil
}
