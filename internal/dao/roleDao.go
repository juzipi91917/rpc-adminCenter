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

// 用于保存角色权限的结构体
type RolePermission struct {
	Id      int64  `json:"id" gorm:"column:id"`
	State   int32  `json:"state" gorm:"column:state"`
	MenuId  int64  `json:"menu_id" gorm:"column:menu_id"`
	Name    string `json:"name" gorm:"column:name"`
	Route   string `json:"route" gorm:"column:route"`
	Remarks string `json:"remarks" gorm:"column:remarks"`
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
func (f *RoleDao) Update(ctx context.Context, where, data *entity.Role) (row int64, err error) {
	tx := f.DB.GetClient().Where(where).Updates(data)
	return tx.RowsAffected, tx.Error
}

// GetRolePermission 根据Role表id查询角色的所有权限信息 返回值为一个装有 permission表的id,state, menu_id, name, route, remarks的结构体列表
func (f *RoleDao) GetRolePermission(ctx context.Context, roleId *int64) ([]*RolePermission, int64, error) {
	// 定义一个切片用于装载查询结果
	var rolePermissionList []*RolePermission

	// 执行查询，并将结果装载到切片中
	err := f.DB.GetClient().Debug().
		Model(&entity.Role{}).
		Select("p.id, p.state, p.menu_id, p.name, p.route, p.remarks").
		Joins("JOIN role_permission rp ON role.id = rp.role_id").
		Joins("JOIN permission p ON rp.permission_id =p.id").
		Where("role.id = ?", roleId).
		Find(&rolePermissionList).Error

	if err != nil {
		return nil, 0, err
	}

	// 获取查询结果的数量
	count := int64(len(rolePermissionList))

	return rolePermissionList, count, nil
}
