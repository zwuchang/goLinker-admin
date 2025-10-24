package request

import (
	"goLinker-admin/server/model/common/request"
	"goLinker-admin/server/model/navigation"
)

// NavScrollNoticeSearch 滚动通知搜索
type NavScrollNoticeSearch struct {
	request.PageInfo
	navigation.NavScrollNotice
}

// CreateNavScrollNoticeRequest 创建滚动通知请求
type CreateNavScrollNoticeRequest struct {
	Title     string `json:"title" form:"title" binding:"required"`     // 通知标题
	Content   string `json:"content" form:"content" binding:"required"` // 通知内容
	Status    int    `json:"status" form:"status"`                      // 状态
	IsVisible int    `json:"isVisible" form:"isVisible"`                // 是否显示
	Sort      int    `json:"sort" form:"sort"`                          // 排序
	StartTime string `json:"startTime" form:"startTime"`                // 开始时间
	EndTime   string `json:"endTime" form:"endTime"`                    // 结束时间
}

// UpdateNavScrollNoticeRequest 更新滚动通知请求
type UpdateNavScrollNoticeRequest struct {
	ID        uint    `json:"id" form:"id" binding:"required"` // ID
	Title     *string `json:"title" form:"title"`              // 通知标题（可选）
	Content   *string `json:"content" form:"content"`          // 通知内容（可选）
	Status    *int    `json:"status" form:"status"`            // 状态（可选）
	IsVisible *int    `json:"isVisible" form:"isVisible"`      // 是否显示（可选）
	Sort      *int    `json:"sort" form:"sort"`                // 排序（可选）
	StartTime *string `json:"startTime" form:"startTime"`      // 开始时间（可选）
	EndTime   *string `json:"endTime" form:"endTime"`          // 结束时间（可选）
}
