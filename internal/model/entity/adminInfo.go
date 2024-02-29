package entity

type (
	// AdminInfo 后台用户详情表
	AdminInfo struct {
		Id          *int64 `gorm:"column:id;type:bigint not null auto_increment;primary_key;comment:自增主键id"` // 自增主键id
		UpdatedAt   *int64 `gorm:"column:updated_at;type:bigint not null;default:0;comment:最后更新时间"`
		AdminId     *int64 `gorm:"column:admin_id;type:bigint not null;comment:后台用户id"`
		GiveAdminId int64  `gorm:"column:give_admin_id;type:bigint not null;comment:后台用户id"`
		RegisterAt  int64  `gorm:"column:register_at;type:bigint not null;default:0;comment:注册时间"`
		RegisterIp  int64  `gorm:"column:register_at;type:bigint not null;default:0;comment:注册ip"`
		NickName    string `gorm:"column:nick_name;type:varchar(255) not null;default:'';comment:用户昵称"`
	}
)

func (AdminInfo) TableName() string {
	return "admin_info"
}
