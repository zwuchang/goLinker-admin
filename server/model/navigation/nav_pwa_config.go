package navigation

import (
	"goLinker-admin/server/global"
)

// NavPWAConfig PWA配置表
type NavPWAConfig struct {
	global.GVA_MODEL
	BackgroundColor string `json:"backgroundColor" gorm:"column:background_color;comment:背景色;size:20;default:#164533"` // 背景色
	ThemeColor      string `json:"themeColor" gorm:"column:theme_color;comment:主题色;size:20;default:#164533"`           // 主题色
	Display         string `json:"display" gorm:"column:display;comment:显示模式;size:20;default:standalone"`              // 显示模式
	Name            string `json:"name" gorm:"column:name;comment:应用名称;size:100;not null"`                             // 应用名称
	ShortName       string `json:"shortName" gorm:"column:short_name;comment:应用短名称;size:50;not null"`                  // 应用短名称
	Orientation     string `json:"orientation" gorm:"column:orientation;comment:屏幕方向;size:20;default:portrait"`        // 屏幕方向
	Scope           string `json:"scope" gorm:"column:scope;comment:作用域;size:200;not null"`                            // 作用域
	StartUrl        string `json:"startUrl" gorm:"column:start_url;comment:启动URL;size:500;not null"`                   // 启动URL
	Status          int    `json:"status" gorm:"column:status;comment:状态 0:禁用 1:启用;default:1"`                         // 状态
	Icons           string `json:"icons" gorm:"column:icons;comment:图标配置;type:json"`                                   // 图标配置
}

// PWAIconItem 图标项结构
type PWAIconItem struct {
	Sizes string `json:"sizes"`
	Src   string `json:"src"`
	Type  string `json:"type"`
	Sort  int    `json:"sort"`
}

// TableName 表名
func (NavPWAConfig) TableName() string {
	return "nav_pwa_configs"
}
