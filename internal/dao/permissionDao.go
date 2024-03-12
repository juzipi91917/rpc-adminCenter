package dao

import (
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/mysqlx"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-adminCenter/internal/model/entity"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-adminCenter/internal/svc"
	"context"
	"github.com/go-redis/redis/v8"
)

type PermissionDao struct {
	DB    mysqlx.DaoHandler[entity.Permission]
	Redis *redis.Client
}

func NewPermissionDao(svcCtx *svc.ServiceContext) *PermissionDao {
	return &PermissionDao{
		DB:    mysqlx.NewDao[entity.Permission](svcCtx.DBCli),
		Redis: svcCtx.RedisCli,
	}
}

// Add -
func (b *PermissionDao) Add(ctx context.Context, add *entity.Permission) (id *int64, err error) {
	k, err := b.DB.Insert(ctx, add)
	return k.Id, err
}

// GetInfo -
func (b *PermissionDao) GetInfo(ctx context.Context, where *entity.Permission) (*entity.Permission, error) {
	return b.DB.FindOne(ctx, where)
}

// GetPageList -
func (f *PermissionDao) GetPageList(ctx context.Context, page, pageSize int64, where *entity.Permission) (list []*entity.Permission, count int64, err error) {
	offset, limit := PageParameterConversion(page, pageSize)
	if err = f.DB.GetClient().Model(&entity.Permission{}).Where(where).Count(&count).Offset(offset).Limit(limit).Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, count, nil
}

// Update -
func (f *PermissionDao) Update(ctx context.Context, where, data *entity.Permission) (row int64, err error) {
	tx := f.DB.GetClient().Where(where).Updates(data)
	return tx.RowsAffected, tx.Error
}
