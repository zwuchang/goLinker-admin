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

// NavGameUpdateRequest 游戏更新请求
type NavGameUpdateRequest struct {
	ID            uint    `json:"id" form:"id" binding:"required"`      // 游戏ID（必填）
	Title         *string `json:"title" form:"title"`                   // 游戏标题（可选）
	Thumbnail     *string `json:"thumbnail" form:"thumbnail"`           // 缩略图（可选）
	IsVisible     *int    `json:"is_visible" form:"is_visible"`         // 是否可见（可选）
	Type          *string `json:"type" form:"type"`                     // 类型（可选）
	VideoUrl      *string `json:"video_url" form:"video_url"`           // 视频地址（可选）
	VideoDuration *string `json:"video_duration" form:"video_duration"` // 视频时长（可选）
	CategoryIds   *string `json:"category_ids" form:"category_ids"`     // 类别ID列表（可选）
	Sticky        *int    `json:"sticky" form:"sticky"`                 // 置顶（可选）
	Views         *int    `json:"views" form:"views"`                   // 浏览次数（可选）
	Description   *string `json:"description" form:"description"`       // 游戏描述（可选）
	Content       *string `json:"content" form:"content"`               // 游戏内容（可选）
	Status        *int    `json:"status" form:"status"`                 // 状态（可选）
	Sort          *int    `json:"sort" form:"sort"`                     // 排序（可选）
	JumpUrl       *string `json:"jump_url" form:"jump_url"`             // 跳转地址（可选）
	DownloadUrl   *string `json:"download_url" form:"download_url"`     // 下载地址（可选）
	DisplayName   *string `json:"display_name" form:"display_name"`     // 显示名称（可选）
	AdName        *string `json:"ad_name" form:"ad_name"`               // 广告名称（可选）
	Article       *string `json:"article" form:"article"`               // 游戏文章（可选）
	Icon          *string `json:"icon" form:"icon"`                     // 游戏图标（可选）
}
