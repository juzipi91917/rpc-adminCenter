package entity

type (
	Bgm struct {
		Id            *int64 `json:"id" gorm:"column:id;type:bigint not null auto_increment;primary_key;comment:自增主键id"` // 自增主键id
		CreatedAt     *int64 `gorm:"column:created_at;type:bigint not null;default:0;comment:创建时间"`                      // 创建时间
		UpdatedAt     *int64 `gorm:"column:updated_at;type:bigint not null;default:0;comment:更新时间"`                      // 更新时间
		Title         string `gorm:"column:title;type:varchar(255) not null;default:'';comment:bgm名称"`
		Author        string `gorm:"column:author;type:varchar(255) not null;default:'';comment:作者"`
		Rank          int32  `gorm:"column:rank;type:int(11) not null;default:0;comment:等级"`
		OtherUseCount int64  `gorm:"column:other_use_count;type:bigint not null;default:0;comment:被使用次数"`
		Bucket        string `gorm:"column:bucket;type:varchar(255) not null;default:'';comment:背景音乐Bucket"`
		Path          string `gorm:"column:path;type:varchar(255) not null;default:'';comment:背景音乐path"`
		CoverBucket   string `gorm:"column:cover_bucket;type:varchar(255) not null;default:'';comment:背景音乐封面图Bucket"`
		CoverPath     string `gorm:"column:cover_path;type:varchar(255) not null;default:'';comment:背景音乐封面图path"`
		Duration      int32  `gorm:"column:duration;type:int(11) not null;default:0;comment:时长"`
		State         *int32 `gorm:"column:state;type:tinyint(4) not null;default:0;comment:状态 0禁用 1启用"`
		Sort          int32  `gorm:"column:sort;sort:int(11) not null;default:0;comment:分类下排行"`
		ClassId       int64  `gorm:"column:class_id;type:bigint not null;default:0;comment:分类id"`
		Type          int32  `gorm:"column:type;type:tinyint(4) not null;default:0;comment:类型 0未知 1平台数据，2抖音数据"`
	}
	TemplateBgm struct {
		Bucket string `gorm:"column:bucket;type:varchar(255) not null;default:'';comment:背景音乐Bucket"`
		Path   string `gorm:"column:path;type:varchar(255) not null;default:'';comment:背景音乐path"`
	}
)

func (Bgm) TableName() string {
	return "bgm"
}
