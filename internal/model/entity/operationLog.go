package entity

type (
	// OperationLog 操作日志表
	OperationLog struct {
		Id            *int64 `json:"id" gorm:"column:id;type:bigint not null auto_increment;primary_key;comment:自增主键id"` // 自增主键id
		CreatedAt     *int64 `gorm:"column:created_at;type:bigint not null;default:0;comment:创建时间"`
		UpdatedAt     *int64 `gorm:"column:updated_at;type:bigint not null;default:0;comment:更新时间"`
		OperationType string `gorm:"column:name;type:varchar(60) not null;default:'';comment:操作类型"`
		AdminId       int64  `gorm:"column:admin_id;type:bigint not null;default:0;comment:操作人id"`
		OperatorName  string `gorm:"column:operator_name;type:varchar(60) not null;default:'';comment:操作人姓名"`
		OperationAt   *int64 `gorm:"column:operation_at;type:bigint not null;default:0;comment:操作时间"`
		OperationDesc string `gorm:"column:operation_desc;type:varchar(1500) not null;default:'';comment:操作描述"`
		ModuleName    string `gorm:"column:module_name;type:varchar(150) not null;default:'';comment:模块名称"`
		Route         string `gorm:"column:route;type:varchar(150) not null;default:'';comment:路由"`
		OperatorIp    string `gorm:"column:operator_ip;type:varchar(15) not null;default:'';comment:操作ip"`
		Content       string `gorm:"column:content;type:varchar(1500) not null;default:'';comment:操作内容"`
		State         int32  `gorm:"column:state;type:tinyint(4) not null;default:0;comment:操作状态：1发起，2成功 3失败"`
		ErrMsg        string `gorm:"column:err_msg;type:text;comment:操作异常信息"`
	}
)

func (OperationLog) TableName() string {
	return "operation_log"
}

type OperationType int8
