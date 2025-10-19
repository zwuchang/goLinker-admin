package navigation

import (
	"goLinker-admin/server/global"
)

// NavPlatformConfig 平台游戏配置
type NavPlatformConfig struct {
	global.GVA_MODEL
	PlatformName string `json:"platform_name" gorm:"column:platform_name;comment:平台名称;size:100;not null"` // 平台名称
	PlatformIcon string `json:"platform_icon" gorm:"column:platform_icon;comment:平台图标地址;size:500"`        // 平台图标地址
	PlatformApi  string `json:"platform_api" gorm:"column:platform_api;comment:平台接口地址;size:500"`          // 平台接口地址
	Sort         int    `json:"sort" gorm:"column:sort;comment:排序;default:0"`                             // 排序
	Status       int    `json:"status" gorm:"column:status;comment:状态 1:启用 0:禁用;default:1"`               // 状态
	Description  string `json:"description" gorm:"column:description;comment:平台描述;size:500"`              // 平台描述
	IsVisible    int    `json:"is_visible" gorm:"column:is_visible;comment:是否可见 1:可见 0:隐藏;default:1"`     // 是否可见
}

// TableName 返回表名
func (NavPlatformConfig) TableName() string {
	return "nav_platform_configs"
}
