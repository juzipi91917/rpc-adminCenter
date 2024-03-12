package dao

import (
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/mysqlx"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-adminCenter/internal/model/entity"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-adminCenter/internal/svc"
	"context"
	"github.com/go-redis/redis/v8"
)

type RolePermissionDao struct {
	DB    mysqlx.DaoHandler[entity.RolePermission]
	Redis *redis.Client
}

func NewRolePermissionDao(svcCtx *svc.ServiceContext) *RolePermissionDao {
	return &RolePermissionDao{
		DB:    mysqlx.NewDao[entity.RolePermission](svcCtx.DBCli),
		Redis: svcCtx.RedisCli,
	}
}

// Add -
func (b *RolePermissionDao) Add(ctx context.Context, add *entity.RolePermission) (id *int64, err error) {
	k, err := b.DB.Insert(ctx, add)
	return k.Id, err
}

// GetInfo -
func (b *RolePermissionDao) GetInfo(ctx context.Context, where *entity.RolePermission) (*entity.RolePermission, error) {
	return b.DB.FindOne(ctx, where)
}

// GetPageList -
func (f *RolePermissionDao) GetPageList(ctx context.Context, page, pageSize int64, where *entity.RolePermission) (list []*entity.RolePermission, count int64, err error) {
	offset, limit := PageParameterConversion(page, pageSize)
	if err = f.DB.GetClient().Model(&entity.RolePermission{}).Where(where).Count(&count).Offset(offset).Limit(limit).Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, count, nil
}

// Update -
func (f *RolePermissionDao) Update(ctx context.Context, where, data *entity.RolePermission) (row int64, err error) {
	tx := f.DB.GetClient().Where(where).Updates(data)
	return tx.RowsAffected, tx.Error
}

// Delete -
func (b *RolePermissionDao) Delete(ctx context.Context, where *entity.RolePermission) (row int64, err error) {
	tx := b.DB.GetClient().Delete(&entity.RolePermission{}, where)
	return tx.RowsAffected, tx.Error
}
