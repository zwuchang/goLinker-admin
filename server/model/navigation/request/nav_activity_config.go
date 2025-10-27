package request

import (
	"goLinker-admin/server/model/common/request"
	"goLinker-admin/server/model/navigation"
)

// NavActivityConfigSearch 活动配置搜索请求
type NavActivityConfigSearch struct {
	request.PageInfo
	navigation.NavActivityConfig
}

// CreateNavActivityConfigRequest 创建活动配置请求
type CreateNavActivityConfigRequest struct {
	Title        string `json:"title" form:"title" binding:"required"` // 显示标题
	Image        string `json:"image" form:"image"`                    // 活动图片
	JumpUrl      string `json:"jumpUrl" form:"jumpUrl"`                // 跳转地址
	CategoryName string `json:"categoryName" form:"categoryName"`      // 类别名称
	CategoryIcon string `json:"categoryIcon" form:"categoryIcon"`      // 类别图标
	Content      string `json:"content" form:"content"`                // 活动内容
	Status       int    `json:"status" form:"status"`                  // 状态
	IsVisible    int    `json:"isVisible" form:"isVisible"`            // 是否显示
	Sort         int    `json:"sort" form:"sort"`                      // 排序
}

// UpdateNavActivityConfigRequest 更新活动配置请求
type UpdateNavActivityConfigRequest struct {
	ID           uint    `json:"id" form:"id" binding:"required"`  // ID
	Title        *string `json:"title" form:"title"`               // 显示标题（可选）
	Image        *string `json:"image" form:"image"`               // 活动图片（可选）
	JumpUrl      *string `json:"jumpUrl" form:"jumpUrl"`           // 跳转地址（可选）
	CategoryName *string `json:"categoryName" form:"categoryName"` // 类别名称（可选）
	CategoryIcon *string `json:"categoryIcon" form:"categoryIcon"` // 类别图标（可选）
	Content      *string `json:"content" form:"content"`           // 活动内容（可选）
	Status       *int    `json:"status" form:"status"`             // 状态（可选）
	IsVisible    *int    `json:"isVisible" form:"isVisible"`       // 是否显示（可选）
	Sort         *int    `json:"sort" form:"sort"`                 // 排序（可选）
}
