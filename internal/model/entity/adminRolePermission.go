package entity

type (
	// AdminRolePermission 用户角色权限表
	AdminRolePermission struct {
		Id           *int64 `json:"id" gorm:"column:id;type:bigint not null auto_increment;primary_key;comment:自增主键id"` // 自增主键id
		CreatedAt    *int64 `gorm:"column:created_at;type:bigint not null;default:0;comment:创建时间"`                      // 创建时间
		AdminId      int64  `gorm:"column:name;type:bigint not null;uniqueIndex:admin_role_permission;default:0;comment:用户id"`
		RoleId       int64  `gorm:"column:name;type:bigint not null;uniqueIndex:admin_role_permission;default:0;comment:角色id"`
		PermissionId int64  `gorm:"column:name;type:bigint not null;uniqueIndex:admin_role_permission;default:0;comment:权限id"`
	}
)

func (AdminRolePermission) TableName() string {
	return "admin_role_permission"
}
