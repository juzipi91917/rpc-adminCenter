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

// PageParameterConversion 分页参数转换
func PageParameterConversion(page, pageSize int64) (offset, limit int) {
	offset = int((page - 1) * pageSize)
	limit = int(pageSize)
	return
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

// GetPageList -
func (f *AdminDao) GetPageList(ctx context.Context, page, pageSize int64, where *entity.Admin) (list []*entity.Admin, count int64, err error) {
	offset, limit := PageParameterConversion(page, pageSize)
	if err = f.DB.GetClient().Model(&entity.Admin{}).Where(where).Count(&count).Offset(offset).Limit(limit).Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, count, nil
}
