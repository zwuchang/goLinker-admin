package navigation

import (
	"goLinker-admin/server/global"
)

// NavActivityConfig 活动配置模型
type NavActivityConfig struct {
	global.GVA_MODEL
	Title        string `json:"title" gorm:"column:title;type:varchar(200);comment:显示标题;not null"`       // 显示标题
	Image        string `json:"image" gorm:"column:image;type:varchar(500);comment:活动图片"`                // 活动图片
	JumpUrl      string `json:"jumpUrl" gorm:"column:jump_url;type:varchar(500);comment:跳转地址"`           // 跳转地址
	CategoryName string `json:"categoryName" gorm:"column:category_name;type:varchar(100);comment:类别名称"` // 类别名称
	CategoryIcon string `json:"categoryIcon" gorm:"column:category_icon;type:varchar(500);comment:类别图标"` // 类别图标
	Content      string `json:"content" gorm:"column:content;type:text;comment:活动内容"`                    // 活动内容
	Status       int    `json:"status" gorm:"column:status;comment:状态 0:禁用 1:启用;default:1"`              // 状态
	IsVisible    int    `json:"isVisible" gorm:"column:is_visible;comment:是否显示 0:隐藏 1:显示;default:1"`     // 是否显示
	Sort         int    `json:"sort" gorm:"column:sort;comment:排序;default:0"`                            // 排序
}

func (NavActivityConfig) TableName() string {
	return "nav_activity_configs"
}
