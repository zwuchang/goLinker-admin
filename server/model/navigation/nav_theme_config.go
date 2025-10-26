package navigation

import (
	"goLinker-admin/server/global"
)

// NavThemeConfig 主题配置模型
type NavThemeConfig struct {
	global.GVA_MODEL
	Name        string `json:"name" form:"name" gorm:"type:varchar(100);comment:主题名称"`               // 主题名称
	Description string `json:"description" form:"description" gorm:"type:varchar(255);comment:主题描述"` // 主题描述
	ConfigJson  string `json:"config_json" form:"config_json" gorm:"type:text;comment:主题配置JSON"`     // 主题配置JSON
	IsDefault   int    `json:"is_default" form:"is_default" gorm:"comment:是否默认主题;default:0"`         // 是否默认主题 1:是 0:否
}

func (NavThemeConfig) TableName() string {
	return "nav_theme_config"
}
