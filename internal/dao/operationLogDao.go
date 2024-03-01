package dao

import (
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/mysqlx"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-adminCenter/internal/model/entity"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-adminCenter/internal/svc"
	"context"
	"github.com/go-redis/redis/v8"
)

type OperationLogDao struct {
	DB    mysqlx.DaoHandler[entity.OperationLog]
	Redis *redis.Client
}

func NewOperationLogDao(svcCtx *svc.ServiceContext) *OperationLogDao {
	return &OperationLogDao{
		DB:    mysqlx.NewDao[entity.OperationLog](svcCtx.DBCli),
		Redis: svcCtx.RedisCli,
	}
}

// Add -
func (o *OperationLogDao) Add(ctx context.Context, add *entity.OperationLog) (id *int64, err error) {
	k, err := o.DB.Insert(ctx, add)
	return k.Id, err
}

// Update -
func (o *OperationLogDao) Update(where, data *entity.OperationLog) (row int64, err error) {
	tx := o.DB.GetClient().Model(&entity.OperationLog{}).Where(where).Updates(data)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return tx.RowsAffected, nil
}
