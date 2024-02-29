package entity

type (
	// Menu 后台显示菜单表
	Menu struct {
		Id        *int64 `json:"id" gorm:"column:id;type:bigint not null auto_increment;primary_key;comment:自增主键id"` // 自增主键id
		CreatedAt *int64 `gorm:"column:created_at;type:bigint not null;default:0;comment:创建时间"`                      // 创建时间
		UpdatedAt *int64 `gorm:"column:updated_at;type:bigint not null;default:0;comment:更新时间"`                      // 更新时间
		DeletedAt *int64 `gorm:"column:deleted_at;type:bigint not null;default:0;comment:删除时间"`
		Name      string `gorm:"column:name;type:varchar(255) not null;default:'';comment:菜单名称"`
		PId       int64  `gorm:"column:p_id;type:bigint not null;default:0;comment:父级菜单id"`
		Remarks   string `gorm:"column:remarks;type:varchar(255) not null;default:'';comment:备注"`
	}
)

func (Menu) TableName() string {
	return "menu"
}
