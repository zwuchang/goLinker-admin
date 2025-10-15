package request

import (
	"goLinker-admin/server/model/common/request"
)

// NavBannerSearch Banner搜索参数
type NavBannerSearch struct {
	request.PageInfo
	Title     string `json:"title" form:"title"`         // 标题搜索
	MediaType *int   `json:"mediaType" form:"mediaType"` // 媒体类型搜索
	IsVisible *int   `json:"isVisible" form:"isVisible"` // 是否显示搜索
	Status    *int   `json:"status" form:"status"`       // 状态搜索
	OrderBy   string `json:"orderBy" form:"orderBy"`     // 排序字段
	OrderType string `json:"orderType" form:"orderType"` // 排序类型 asc/desc
}

// NavBannerCreateRequest Banner创建请求
type NavBannerCreateRequest struct {
	Title       string `json:"title" binding:"required"`     // 标题
	Description string `json:"description"`                  // 描述
	MediaType   int    `json:"mediaType" binding:"required"` // 媒体类型 1:图片 2:视频
	MediaUrl    string `json:"mediaUrl" binding:"required"`  // 媒体地址
	LinkUrl     string `json:"linkUrl"`                      // 跳转链接
	Sort        int    `json:"sort"`                         // 排序
	IsVisible   int    `json:"isVisible"`                    // 是否显示
	Status      int    `json:"status"`                       // 状态
}

// NavBannerUpdateRequest Banner更新请求
type NavBannerUpdateRequest struct {
	ID          uint   `json:"id" binding:"required"` // ID
	Title       string `json:"title"`                 // 标题
	Description string `json:"description"`           // 描述
	MediaType   int    `json:"mediaType"`             // 媒体类型
	MediaUrl    string `json:"mediaUrl"`              // 媒体地址
	LinkUrl     string `json:"linkUrl"`               // 跳转链接
	Sort        int    `json:"sort"`                  // 排序
	IsVisible   int    `json:"isVisible"`             // 是否显示
	Status      int    `json:"status"`                // 状态
}

// DeleteBannerRequest Banner删除请求
type DeleteBannerRequest struct {
	ID uint `json:"id" binding:"required"` // Banner ID
}
