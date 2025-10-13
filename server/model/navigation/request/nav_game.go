package request

import (
	"goLinker-admin/server/model/common/request"
	"goLinker-admin/server/model/navigation"
)

// GetById 通用ID请求结构
type GetById struct {
	ID uint `json:"id" form:"id" binding:"required"`
}

// NavGameCategorySearch 游戏类别搜索请求
type NavGameCategorySearch struct {
	navigation.NavGameCategory
	request.PageInfo
}

// NavGameSearch 游戏搜索请求
type NavGameSearch struct {
	navigation.NavGame
	request.PageInfo
	CategoryId uint   `json:"category_id" form:"category_id"` // 按类别筛选
	OrderBy    string `json:"order_by" form:"order_by"`       // 排序字段
	OrderType  string `json:"order_type" form:"order_type"`   // 排序类型：asc, desc
}
