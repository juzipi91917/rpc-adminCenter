package dao

import (
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/mysqlx"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-clipFilm/internal/model/entity"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-clipFilm/internal/svc"
	"context"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm/clause"
)

type BgmClassDao struct {
	DB    mysqlx.DaoHandler[entity.BgmClass]
	Redis *redis.Client
}

func NewBgmClassDao(svcCtx *svc.ServiceContext) *BgmClassDao {
	return &BgmClassDao{
		DB:    mysqlx.NewDao[entity.BgmClass](svcCtx.DBCli),
		Redis: svcCtx.RedisCli,
	}
}

// Add -
func (b *BgmClassDao) Add(ctx context.Context, add *entity.BgmClass) (id *int64, err error) {
	k, err := b.DB.Insert(ctx, add)
	return k.Id, err
}

// GetInfo -
func (b *BgmClassDao) GetInfo(ctx context.Context, where *entity.BgmClass) (*entity.BgmClass, error) {
	return b.DB.FindOne(ctx, where)
}

// GetAll -
func (b *BgmClassDao) GetAll(ctx context.Context, state *int32) (all []*entity.BgmClass, err error) {
	order := []clause.OrderByColumn{
		{
			Column: clause.Column{
				Name: "sort",
			},
			Desc:    true,
			Reorder: false,
		},
	}
	return b.DB.FindMany(ctx, &entity.BgmClass{
		State: state,
	}, order)
}

// Update -
func (b *BgmClassDao) Update(ctx context.Context, where, date *entity.BgmClass) (row int64, err error) {
	tx := b.DB.GetClient().Where(where).Updates(date)
	return tx.RowsAffected, tx.Error
}
