package rpc

import (
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/commonx/result"
	"context"
)

type UserService struct {
}

// NewUserService -
func NewUserService() *UserService {
	return &UserService{}
}

// CheckUserStrip 调用用户rpc校验用户条数
func (u *UserService) CheckUserStrip(ctx context.Context, uuid string, count int32) (surplusTotal int32, err error) {
	surplusTotal = 10
	if err != nil {
		return 0, result.NewRpcError("获取您的可用条数失败,请联系客服！")
	}
	if surplusTotal < count {
		return 0, result.NewRpcError("您的可用条数不足,请充值！")
	}
	return surplusTotal, nil
}

// UserDeductNumber 调用用户rpc扣除条数
func (u *UserService) UserDeductNumber(ctx context.Context, strip, detailId int64, uuid string) (err error) {

	return nil
}

// UserDeductReturnNumber 调用用户rpc退回条数
func (u *UserService) UserDeductReturnNumber(ctx context.Context, strip, detailId int64, uuid string) (err error) {

	return nil
}
