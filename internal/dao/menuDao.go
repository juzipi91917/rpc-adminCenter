package dao

import (
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/mysqlx"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-adminCenter/internal/model/entity"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-adminCenter/internal/svc"
	"context"
	"github.com/go-redis/redis/v8"
)

type MenuDao struct {
	DB    mysqlx.DaoHandler[entity.Menu]
	Redis *redis.Client
}

func NewMenuDao(svcCtx *svc.ServiceContext) *MenuDao {
	return &MenuDao{
		DB:    mysqlx.NewDao[entity.Menu](svcCtx.DBCli),
		Redis: svcCtx.RedisCli,
	}
}

// Add -
func (b *MenuDao) Add(ctx context.Context, add *entity.Menu) (id *int64, err error) {
	k, err := b.DB.Insert(ctx, add)
	return k.Id, err
}

// GetInfo -
func (b *MenuDao) GetInfo(ctx context.Context, where *entity.Menu) (*entity.Menu, error) {
	return b.DB.FindOne(ctx, where)
}

// GetPageList -
func (f *MenuDao) GetPageList(ctx context.Context, page, pageSize int64, where *entity.Menu) (list []*entity.Menu, count int64, err error) {
	offset, limit := PageParameterConversion(page, pageSize)
	if err = f.DB.GetClient().Model(&entity.Menu{}).Where(where).Count(&count).Offset(offset).Limit(limit).Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, count, nil
}

// Update -
func (f *MenuDao) Update(ctx context.Context, where, data *entity.Menu) (row int64, err error) {
	tx := f.DB.GetClient().Where(where).Updates(data)
	return tx.RowsAffected, tx.Error
}
