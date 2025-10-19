package navigation

import (
	"goLinker-admin/server/global"
)

// NavGameConfig 游戏配置模型
type NavGameConfig struct {
	global.GVA_MODEL
	DownloadUrl  string `json:"download_url" form:"download_url" gorm:"type:varchar(500);comment:下载地址"`   // 下载地址
	AudioUrl     string `json:"audio_url" form:"audio_url" gorm:"type:varchar(500);comment:播放音频地址"`       // 播放音频地址
	WebsiteTitle string `json:"website_title" form:"website_title" gorm:"type:varchar(200);comment:网站标题"` // 网站标题
	WebsiteDesc  string `json:"website_desc" form:"website_desc" gorm:"type:text;comment:网站描述"`           // 网站描述
	WebsiteIcon  string `json:"website_icon" form:"website_icon" gorm:"type:varchar(500);comment:网站图标"`   // 网站图标
	WebsiteLogo  string `json:"website_logo" form:"website_logo" gorm:"type:varchar(500);comment:网站Logo"` // 网站Logo
	MarketLogo   string `json:"market_logo" form:"market_logo" gorm:"type:varchar(500);comment:市场Logo"`   // 市场Logo
	Status       int    `json:"status" form:"status" gorm:"comment:状态;default:1"`                         // 状态 1:启用 0:禁用
}

func (NavGameConfig) TableName() string {
	return "nav_game_config"
}
