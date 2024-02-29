package entity

type (
	// AdminRole 用户角色表
	AdminRole struct {
		Id        *int64 `json:"id" gorm:"column:id;type:bigint not null auto_increment;primary_key;comment:自增主键id"`
		CreatedAt *int64 `gorm:"column:created_at;type:bigint not null;default:0;comment:创建时间"`
		AdminId   int64  `gorm:"column:name;type:bigint not null;index;uniqueIndex:admin_role;default:0;comment:用户id"`
		RoleId    int64  `gorm:"column:p_id;type:bigint not null;index;uniqueIndex:admin_role;default:0;comment:角色id"`
	}
)

func (AdminRole) TableName() string {
	return "admin_role"
}
