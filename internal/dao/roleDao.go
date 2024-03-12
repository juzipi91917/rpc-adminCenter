package dao

import (
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/mysqlx"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-adminCenter/internal/model/entity"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-adminCenter/internal/svc"
	"context"
	"github.com/go-redis/redis/v8"
)

type RoleDao struct {
	DB    mysqlx.DaoHandler[entity.Role]
	Redis *redis.Client
}

func NewRoleDao(svcCtx *svc.ServiceContext) *RoleDao {
	return &RoleDao{
		DB:    mysqlx.NewDao[entity.Role](svcCtx.DBCli),
		Redis: svcCtx.RedisCli,
	}
}

// Add -
func (b *RoleDao) Add(ctx context.Context, add *entity.Role) (id *int64, err error) {
	k, err := b.DB.Insert(ctx, add)
	return k.Id, err
}

// GetInfo -
func (b *RoleDao) GetInfo(ctx context.Context, where *entity.Role) (*entity.Role, error) {
	return b.DB.FindOne(ctx, where)
}

// GetPageList -
func (f *RoleDao) GetPageList(ctx context.Context, page, pageSize int64, where *entity.Role) (list []*entity.Role, count int64, err error) {
	offset, limit := PageParameterConversion(page, pageSize)
	if err = f.DB.GetClient().Model(&entity.Role{}).Where(where).Count(&count).Offset(offset).Limit(limit).Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, count, nil
}

// Update -
func (f *RoleDao) Update(ctx context.Context, where, date *entity.Role) (row int64, err error) {
	tx := f.DB.GetClient().Where(where).Updates(date)
	return tx.RowsAffected, tx.Error
}
