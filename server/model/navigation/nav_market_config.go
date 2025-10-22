package navigation

import (
	"goLinker-admin/server/global"
)

// NavMarketConfig 市场配置结构体
type NavMarketConfig struct {
	global.GVA_MODEL
	Name        string `json:"name" form:"name" gorm:"column:name;comment:市场名称;size:100;"`
	Logo        string `json:"logo" form:"logo" gorm:"column:logo;comment:市场Logo地址;size:255;"`
	JumpUrl     string `json:"jump_url" form:"jump_url" gorm:"column:jump_url;comment:跳转地址;size:255;"`
	Type        int    `json:"type" form:"type" gorm:"column:type;comment:市场类型（1-3）;default:1;"`
	Status      int    `json:"status" form:"status" gorm:"column:status;comment:状态（0禁用 1启用）;default:1;"`
	IsVisible   int    `json:"is_visible" form:"is_visible" gorm:"column:is_visible;comment:是否可见（0隐藏 1可见）;default:1;"`
	Sort        int    `json:"sort" form:"sort" gorm:"column:sort;comment:排序;default:0;"`
	Description string `json:"description" form:"description" gorm:"column:description;comment:市场描述;type:text;"`
}

// TableName NavMarketConfig 表名
func (NavMarketConfig) TableName() string {
	return "nav_market_configs"
}


