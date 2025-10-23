package request

import (
	"goLinker-admin/server/model/common/request"
	"goLinker-admin/server/model/navigation"
)

// NavPlatformConfigSearch 平台游戏配置搜索条件
type NavPlatformConfigSearch struct {
	request.PageInfo
	PlatformName string `json:"platform_name" form:"platform_name"` // 平台名称
	Status       *int   `json:"status" form:"status"`               // 状态
	IsVisible    *int   `json:"is_visible" form:"is_visible"`       // 是否可见
}

// CreateNavPlatformConfigRequest 创建平台游戏配置请求
type CreateNavPlatformConfigRequest struct {
	PlatformName string `json:"platform_name" binding:"required"` // 平台名称
	PlatformIcon string `json:"platform_icon"`                    // 平台图标地址
	PlatformApi  string `json:"platform_api"`                     // 平台接口地址
	Sort         int    `json:"sort"`                             // 排序
	Status       int    `json:"status"`                           // 状态
	Description  string `json:"description"`                      // 平台描述
	IsVisible    int    `json:"is_visible"`                       // 是否可见
}

// UpdateNavPlatformConfigRequest 更新平台游戏配置请求
type UpdateNavPlatformConfigRequest struct {
	ID           uint   `json:"id" binding:"required"`            // ID
	PlatformName string `json:"platform_name" binding:"required"` // 平台名称
	PlatformIcon string `json:"platform_icon"`                    // 平台图标地址
	PlatformApi  string `json:"platform_api"`                     // 平台接口地址
	Sort         int    `json:"sort"`                             // 排序
	Status       int    `json:"status"`                           // 状态
	Description  string `json:"description"`                      // 平台描述
	IsVisible    int    `json:"is_visible"`                       // 是否可见
}

// GetNavPlatformConfigByIdRequest 根据ID获取平台游戏配置请求
type GetNavPlatformConfigByIdRequest struct {
	ID uint `json:"id" binding:"required"` // ID
}

// DeleteNavPlatformConfigRequest 删除平台游戏配置请求
type DeleteNavPlatformConfigRequest struct {
	ID uint `json:"id" binding:"required"` // ID
}

// NavPlatformConfig 平台游戏配置模型（用于转换）
type NavPlatformConfig struct {
	navigation.NavPlatformConfig
}



