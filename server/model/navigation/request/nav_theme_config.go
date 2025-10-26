package request

import (
	"goLinker-admin/server/model/common/request"
	"goLinker-admin/server/model/navigation"
)

// NavThemeConfigSearch 主题配置搜索请求
type NavThemeConfigSearch struct {
	navigation.NavThemeConfig
	request.PageInfo
}

// CreateNavThemeConfigRequest 创建主题配置请求
type CreateNavThemeConfigRequest struct {
	Name        string `json:"name" form:"name" binding:"required"`               // 主题名称
	Description string `json:"description" form:"description"`                    // 主题描述
	ConfigJson  string `json:"config_json" form:"config_json" binding:"required"` // 主题配置JSON
	IsDefault   int    `json:"is_default" form:"is_default"`                      // 是否默认主题
}

// UpdateNavThemeConfigRequest 更新主题配置请求
type UpdateNavThemeConfigRequest struct {
	ID          uint   `json:"id" form:"id" binding:"required"`                   // ID
	Name        string `json:"name" form:"name" binding:"required"`               // 主题名称
	Description string `json:"description" form:"description"`                    // 主题描述
	ConfigJson  string `json:"config_json" form:"config_json" binding:"required"` // 主题配置JSON
	IsDefault   int    `json:"is_default" form:"is_default"`                      // 是否默认主题
}

// SetDefaultThemeRequest 设置默认主题请求
type SetDefaultThemeRequest struct {
	ID uint `json:"id" form:"id" binding:"required"` // 主题ID
}
