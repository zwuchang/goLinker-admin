package navigation

import (
	"goLinker-admin/server/global"
)

// NavGameCategoryRelation 游戏分类关联表
type NavGameCategoryRelation struct {
	global.GVA_MODEL
	GameId     uint `json:"game_id" gorm:"comment:游戏ID;index"`     // 游戏ID
	CategoryId uint `json:"category_id" gorm:"comment:分类ID;index"` // 分类ID
	Sort       int  `json:"sort" gorm:"comment:排序;default:0"`      // 排序（可选，用于同一游戏内分类的排序）
}

func (NavGameCategoryRelation) TableName() string {
	return "nav_game_category_relation"
}




