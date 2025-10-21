package navigation

import (
	"goLinker-admin/server/global"
)

// NavGameConfig 游戏配置模型
type NavGameConfig struct {
	global.GVA_MODEL
	DownloadUrl    string `json:"download_url" form:"download_url" gorm:"type:varchar(500);comment:下载地址"`        // 下载地址
	AudioUrl       string `json:"audio_url" form:"audio_url" gorm:"type:varchar(500);comment:播放音频地址"`            // 播放音频地址
	WebsiteTitle   string `json:"website_title" form:"website_title" gorm:"type:varchar(200);comment:网站标题"`      // 网站标题
	WebsiteDesc    string `json:"website_desc" form:"website_desc" gorm:"type:text;comment:网站描述"`                // 网站描述
	WebsiteIcon    string `json:"website_icon" form:"website_icon" gorm:"type:varchar(500);comment:网站图标"`        // 网站图标
	WebsiteLogo    string `json:"website_logo" form:"website_logo" gorm:"type:varchar(500);comment:网站Logo"`      // 网站Logo
	MarketLogo     string `json:"market_logo" form:"market_logo" gorm:"type:varchar(500);comment:市场Logo"`        // 市场Logo
	FloatingStatus int    `json:"floating_status" form:"floating_status" gorm:"comment:悬浮图标状态;default:0"`        // 悬浮图标状态 1:开启 0:关闭
	FloatingIcon1  string `json:"floating_icon1" form:"floating_icon1" gorm:"type:varchar(500);comment:悬浮1图标地址"` // 悬浮1图标地址
	FloatingIcon2  string `json:"floating_icon2" form:"floating_icon2" gorm:"type:varchar(500);comment:悬浮2图标地址"` // 悬浮2图标地址
	FloatingIcon3  string `json:"floating_icon3" form:"floating_icon3" gorm:"type:varchar(500);comment:悬浮3图标地址"` // 悬浮3图标地址
	Status         int    `json:"status" form:"status" gorm:"comment:状态;default:1"`                              // 状态 1:启用 0:禁用
}

func (NavGameConfig) TableName() string {
	return "nav_game_config"
}
