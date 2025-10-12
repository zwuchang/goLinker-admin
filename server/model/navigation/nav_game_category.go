package navigation

import (
	"goLinker-admin/server/global"
)

// NavGameCategory 游戏类别模型
type NavGameCategory struct {
	global.GVA_MODEL
	Name        string `json:"name" form:"name" gorm:"comment:类别名称"`               // 类别名称
	Description string `json:"description" form:"description" gorm:"comment:类别描述"` // 类别描述
	Sort        int    `json:"sort" form:"sort" gorm:"comment:排序;default:0"`       // 排序
	Status      int    `json:"status" form:"status" gorm:"comment:状态;default:1"`   // 状态 1:启用 0:禁用
	Icon        string `json:"icon" form:"icon" gorm:"comment:类别图标"`               // 类别图标
}

func (NavGameCategory) TableName() string {
	return "nav_game_category"
}
