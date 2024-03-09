package entity

type (
	// Menu 后台显示菜单表
	Menu struct {
		Id        *int64 `json:"id" gorm:"column:id;type:bigint not null auto_increment;primary_key;comment:自增主键id"` // 自增主键id
		CreatedAt *int64 `gorm:"column:created_at;type:bigint not null;default:0;comment:创建时间"`                      // 创建时间
		UpdatedAt *int64 `gorm:"column:updated_at;type:bigint not null;default:0;comment:更新时间"`                      // 更新时间
		State     *int32 `gorm:"column:state;type:tinyint(4) not null;default:0;comment:状态 0禁用 1启用"`
		Name      string `gorm:"column:name;type:varchar(255) not null;default:'';comment:菜单名称"`
		Route     string `gorm:"column:route;type:varchar(255) not null;default:'';comment:前端菜单路由"`
		PId       int64  `gorm:"column:p_id;type:bigint not null;default:0;comment:父级菜单id"`
		LevelIds  string `gorm:"column:level_ids;type:varchar(1500) not null;default:'';comment:父级菜单层级id"`
		Remarks   string `gorm:"column:remarks;type:varchar(255) not null;default:'';comment:备注"`
	}
)

func (Menu) TableName() string {
	return "menu"
}
