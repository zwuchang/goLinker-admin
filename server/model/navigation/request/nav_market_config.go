package request

import (
	"goLinker-admin/server/model/common/request"
	"goLinker-admin/server/model/navigation"
)

// NavMarketConfigSearch 市场配置搜索条件
type NavMarketConfigSearch struct {
	navigation.NavMarketConfig
	request.PageInfo
}

// CreateMarketConfig 创建市场配置请求
type CreateMarketConfig struct {
	Name        string `json:"name" form:"name" binding:"required"`
	Logo        string `json:"logo" form:"logo"`
	JumpUrl     string `json:"jump_url" form:"jump_url"`
	Type        int    `json:"type" form:"type" binding:"required,min=1,max=3"`
	Status      int    `json:"status" form:"status"`
	IsVisible   int    `json:"is_visible" form:"is_visible"`
	Sort        int    `json:"sort" form:"sort"`
	Description string `json:"description" form:"description"`
}

// UpdateMarketConfig 更新市场配置请求
type UpdateMarketConfig struct {
	ID uint `json:"id" form:"id" binding:"required"`
	CreateMarketConfig
}

// DeleteMarketConfig 删除市场配置请求
type DeleteMarketConfig struct {
	ID uint `json:"id" form:"id" binding:"required"`
}

// GetMarketByTypeRequest 根据类型获取市场配置请求
type GetMarketByTypeRequest struct {
	Type int `json:"type" form:"type" binding:"required,min=1,max=3"`
}




