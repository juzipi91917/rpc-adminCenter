package dao

import (
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/mysqlx"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-adminCenter/internal/model/entity"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-adminCenter/internal/svc"
	"context"
	"github.com/go-redis/redis/v8"
)

type AdminDao struct {
	DB    mysqlx.DaoHandler[entity.Admin]
	Redis *redis.Client
}

func NewAdminDao(svcCtx *svc.ServiceContext) *AdminDao {
	return &AdminDao{
		DB:    mysqlx.NewDao[entity.Admin](svcCtx.DBCli),
		Redis: svcCtx.RedisCli,
	}
}

// Add -
func (b *AdminDao) Add(ctx context.Context, add *entity.Admin) (id *int64, err error) {
	k, err := b.DB.Insert(ctx, add)
	return k.Id, err
}

// GetInfo -
func (b *AdminDao) GetInfo(ctx context.Context, where *entity.Admin) (*entity.Admin, error) {
	return b.DB.FindOne(ctx, where)
}
