package navigation

import (
	"goLinker-admin/server/global"
)

// NavGame 游戏模型
type NavGame struct {
	global.GVA_MODEL
	Title         string `json:"title" form:"title" gorm:"comment:游戏标题"`                     // 游戏标题
	Thumbnail     string `json:"thumbnail" form:"thumbnail" gorm:"comment:缩略图"`              // 缩略图
	IsVisible     int    `json:"is_visible" form:"is_visible" gorm:"comment:是否可见;default:1"` // 是否可见 1:可见 0:隐藏
	Type          string `json:"type" form:"type" gorm:"comment:类型;default:image_text"`      // 类型 image_text:图文 video:视频
	VideoUrl      string `json:"video_url" form:"video_url" gorm:"comment:视频地址"`             // 视频地址
	VideoDuration string `json:"video_duration" form:"video_duration" gorm:"comment:视频时长"`   // 视频时长
	CategoryIds   string `json:"category_ids" form:"category_ids" gorm:"comment:类别ID列表"`     // 类别ID列表，JSON格式
	Sticky        int    `json:"sticky" form:"sticky" gorm:"comment:置顶;default:0"`           // 置顶 1:置顶 0:普通
	Views         int    `json:"views" form:"views" gorm:"comment:浏览次数;default:0"`           // 浏览次数
	Description   string `json:"description" form:"description" gorm:"comment:游戏描述"`         // 游戏描述
	Content       string `json:"content" form:"content" gorm:"comment:游戏内容"`                 // 游戏内容
	Status        int    `json:"status" form:"status" gorm:"comment:状态;default:1"`           // 状态 1:启用 0:禁用
	Sort          int    `json:"sort" form:"sort" gorm:"comment:排序;default:0"`               // 排序
}

func (NavGame) TableName() string {
	return "nav_game"
}
