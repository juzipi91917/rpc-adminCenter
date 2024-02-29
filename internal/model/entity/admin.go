package entity

type (
	// Admin 后台用户表
	Admin struct {
		Id        *int64 `json:"id" gorm:"column:id;type:bigint not null auto_increment;primary_key;comment:自增主键id"`
		CreatedAt *int64 `gorm:"column:created_at;type:bigint not null;default:0;comment:创建时间"`
		UpdatedAt *int64 `gorm:"column:updated_at;type:bigint not null;default:0;comment:更新时间"`
		State     *int32 `gorm:"column:state;type:tinyint(4) not null;default:0;comment:状态 0禁用 1启用"`
		Account   string `gorm:"column:account;type:varchar(255) not null;default:'';comment:账号"`
		Password  string `gorm:"column:password;type:varchar(255) not null;default:'';comment:密码"`
		Mobile    string `gorm:"column:mobile;type:varchar(16) not null;default:'';comment:手机号"`
	}
)

func (Admin) TableName() string {
	return "admin"
}
