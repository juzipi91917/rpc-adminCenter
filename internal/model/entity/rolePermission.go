package entity

type (
	// RolePermission 角色权限表
	RolePermission struct {
		Id           *int64 `json:"id" gorm:"column:id;type:bigint not null auto_increment;primary_key;comment:自增主键id"`
		CreatedAt    *int64 `gorm:"column:created_at;type:bigint not null;default:0;comment:创建时间"`
		RoleId       int64  `gorm:"column:role_id;type:bigint not null;index;uniqueIndex:role_permission;default:0;comment:角色id"`
		PermissionId int64  `gorm:"column:permission_id;type:bigint not null;index;uniqueIndex:role_permission;default:0;comment:权限id"`
	}
)

func (RolePermission) TableName() string {
	return "role_permission"
}
