package dao

import (
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/mysqlx"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-adminCenter/internal/model/entity"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-adminCenter/internal/svc"
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
)

type AdminDao struct {
	DB    mysqlx.DaoHandler[entity.Admin]
	Redis *redis.Client
}

// 用于保存查用户角色信息的结构体
type AdminRole struct {
	ID      int64  `json:"id" gorm:"column:id"`
	Name    string `json:"name" gorm:"column:name"`
	Remarks string `json:"remarks" gorm:"column:remarks"`
	State   int32  `json:"state" gorm:"column:state"`
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

// Update -
func (f *AdminDao) Update(ctx context.Context, where, date *entity.Admin) (row int64, err error) {
	tx := f.DB.GetClient().Where(where).Updates(date)
	return tx.RowsAffected, tx.Error
}

// ChangeState 切换用户状态
func (b *AdminDao) ChangeState(ctx context.Context, userID int64) (row int64, err error) {
	// 查询用户当前状态
	user, err := b.GetInfo(ctx, &entity.Admin{Id: &userID})
	if err != nil {
		return 0, err
	}
	if user.Id == nil {
		return 0, errors.New("修改失败, 该用户不存在")
	}

	// 切换状态
	var state int32

	if *user.State == 1 {
		state = 0
	} else {
		state = 1
	}
	user.State = &state

	// 更新用户状态
	row, err = b.Update(ctx, &entity.Admin{Id: &userID}, user)

	if err != nil {
		return 0, err
	}

	return row, nil
}

// GetAdminRoleList 根据用户id查询用户的所有角色信息 返回值为role表的id,name,remarks,state字段
func (f *AdminDao) GetAdminRoleList(ctx context.Context, adminId *int64) ([]*AdminRole, int64, error) {
	// 定义一个切片用于装载查询结果
	var adminRoleList []*AdminRole

	// 执行查询，并将结果装载到切片中
	err := f.DB.GetClient().Debug().Model(&entity.Admin{}).
		Select("r.id, r.name, r.remarks, r.state").
		Joins("JOIN admin_role ar ON admin.id = ar.a_id").
		Joins("JOIN role r ON ar.p_id = r.id").
		Where("admin.id = ?", adminId).
		Find(&adminRoleList).Error

	if err != nil {
		return nil, 0, err
	}

	// 获取查询结果的数量
	count := int64(len(adminRoleList))

	return adminRoleList, count, nil
}
