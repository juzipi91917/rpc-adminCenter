package entity

type (
	// Role 用户角色
	Role struct {
		Id        *int64 `json:"id" gorm:"column:id;type:bigint not null auto_increment;primary_key;comment:自增主键id"` // 自增主键id
		CreatedAt *int64 `gorm:"column:created_at;type:bigint not null;default:0;comment:创建时间"`                      // 创建时间
		UpdatedAt *int64 `gorm:"column:updated_at;type:bigint not null;default:0;comment:更新时间"`                      // 更新时间
		State     *int32 `gorm:"column:state;type:tinyint(4) not null;default:0;comment:状态 0禁用 1启用"`
		Name      string `gorm:"column:name;type:varchar(255) not null;default:'';comment:角色名"`
		Remarks   string `gorm:"column:remarks;type:varchar(255) not null;default:'';comment:备注"`
	}
)

func (Role) TableName() string {
	return "role"
}
