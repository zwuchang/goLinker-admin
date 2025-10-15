package navigation

import (
	"goLinker-admin/server/global"
)

// NavBanner Banner配置表
type NavBanner struct {
	global.GVA_MODEL
	Title       string `json:"title" gorm:"column:title;comment:标题;size:100;"`                       // 标题
	Description string `json:"description" gorm:"column:description;comment:描述;size:500;"`           // 描述
	MediaType   int    `json:"mediaType" gorm:"column:media_type;comment:媒体类型 1:图片 2:视频;default:1;"` // 媒体类型 1:图片 2:视频
	MediaUrl    string `json:"mediaUrl" gorm:"column:media_url;comment:媒体地址;size:500;"`              // 媒体地址
	LinkUrl     string `json:"linkUrl" gorm:"column:link_url;comment:跳转链接;size:500;"`                // 跳转链接
	Sort        int    `json:"sort" gorm:"column:sort;comment:排序;default:0;"`                        // 排序
	IsVisible   int    `json:"isVisible" gorm:"column:is_visible;comment:是否显示 0:隐藏 1:显示;default:1;"` // 是否显示
	Status      int    `json:"status" gorm:"column:status;comment:状态 0:禁用 1:启用;default:1;"`          // 状态
}

// TableName 表名
func (NavBanner) TableName() string {
	return "nav_banners"
}
