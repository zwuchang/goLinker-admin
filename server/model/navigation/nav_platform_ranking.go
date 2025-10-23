package navigation

import (
	"goLinker-admin/server/global"
)

// NavPlatformRanking 平台排行榜表
type NavPlatformRanking struct {
	global.GVA_MODEL
	Rank         int     `json:"rank" gorm:"column:rank;comment:排名;not null"`                             // 排名
	PlatformName string  `json:"platformName" gorm:"column:platform_name;comment:平台名称;size:100;not null"` // 平台名称
	Logo         string  `json:"logo" gorm:"column:logo;comment:平台Logo地址;size:500"`                       // 平台Logo地址
	Rating       float64 `json:"rating" gorm:"column:rating;comment:评分;type:decimal(3,1)"`                // 评分
	Features     string  `json:"features" gorm:"column:features;comment:特色功能描述;size:500"`                 // 特色功能描述
	FeatureType  string  `json:"featureType" gorm:"column:feature_type;comment:功能类型;size:50"`             // 功能类型
	VisitUrl     string  `json:"visitUrl" gorm:"column:visit_url;comment:访问链接;size:500"`                  // 访问链接
	IsNew        int     `json:"isNew" gorm:"column:is_new;comment:是否新平台 0:否 1:是;default:0"`              // 是否新平台
	Status       int     `json:"status" gorm:"column:status;comment:状态 0:禁用 1:启用;default:1"`              // 状态
	IsVisible    int     `json:"isVisible" gorm:"column:is_visible;comment:是否显示 0:隐藏 1:显示;default:1"`     // 是否显示
	Sort         int     `json:"sort" gorm:"column:sort;comment:排序;default:0"`                            // 排序
}

// TableName 表名
func (NavPlatformRanking) TableName() string {
	return "nav_platform_rankings"
}
