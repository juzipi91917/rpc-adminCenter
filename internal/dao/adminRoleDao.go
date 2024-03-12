package dao

import (
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/mysqlx"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-adminCenter/internal/model/entity"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-adminCenter/internal/svc"
	"context"
	"github.com/go-redis/redis/v8"
)

type AdminRoleDao struct {
	DB    mysqlx.DaoHandler[entity.AdminRole]
	Redis *redis.Client
}

func NewAdminRoleDao(svcCtx *svc.ServiceContext) *AdminRoleDao {
	return &AdminRoleDao{
		DB:    mysqlx.NewDao[entity.AdminRole](svcCtx.DBCli),
		Redis: svcCtx.RedisCli,
	}
}

// Add -
func (b *AdminRoleDao) Add(ctx context.Context, add *entity.AdminRole) (id *int64, err error) {
	k, err := b.DB.Insert(ctx, add)
	return k.Id, err
}

// Delete -
func (b *AdminRoleDao) Delete(ctx context.Context, where *entity.AdminRole) (row int64, err error) {
	tx := b.DB.GetClient().Delete(&entity.AdminRole{}, where)
	return tx.RowsAffected, tx.Error
}

// GetInfo -
func (b *AdminRoleDao) GetInfo(ctx context.Context, where *entity.Admin) (*entity.AdminRole, error) {
	return b.DB.FindOne(ctx, where)
}

// GetPageList -
func (f *AdminRoleDao) GetPageList(ctx context.Context, page, pageSize int64, where *entity.AdminRole) (list []*entity.AdminRole, count int64, err error) {
	offset, limit := PageParameterConversion(page, pageSize)
	if err = f.DB.GetClient().Model(&entity.AdminRole{}).Where(where).Count(&count).Offset(offset).Limit(limit).Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, count, nil
}

// Update -
func (f *AdminRoleDao) Update(ctx context.Context, where, date *entity.AdminRole) (row int64, err error) {
	tx := f.DB.GetClient().Where(where).Updates(date)
	return tx.RowsAffected, tx.Error
}
