package navigation

import (
	"goLinker-admin/server/global"
)

// NavScrollNotice 滚动通知表
type NavScrollNotice struct {
	global.GVA_MODEL
	Title     string `json:"title" gorm:"column:title;comment:通知标题;size:200;not null"`            // 通知标题
	Content   string `json:"content" gorm:"column:content;comment:通知内容;type:text"`                // 通知内容
	Status    int    `json:"status" gorm:"column:status;comment:状态 0:禁用 1:启用;default:1"`          // 状态
	IsVisible int    `json:"isVisible" gorm:"column:is_visible;comment:是否显示 0:隐藏 1:显示;default:1"` // 是否显示
	Sort      int    `json:"sort" gorm:"column:sort;comment:排序;default:0"`                        // 排序
	StartTime string `json:"startTime" gorm:"column:start_time;comment:开始时间;type:datetime"`       // 开始时间
	EndTime   string `json:"endTime" gorm:"column:end_time;comment:结束时间;type:datetime"`           // 结束时间
}

// TableName 表名
func (NavScrollNotice) TableName() string {
	return "nav_scroll_notices"
}
