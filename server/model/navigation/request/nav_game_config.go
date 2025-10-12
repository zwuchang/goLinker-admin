package request

import (
	"goLinker-admin/server/model/common/request"
	"goLinker-admin/server/model/navigation"
)

// NavGameConfigSearch 游戏配置搜索请求
type NavGameConfigSearch struct {
	navigation.NavGameConfig
	request.PageInfo
	OrderKey string `json:"orderKey" form:"orderKey"`
	Desc     bool   `json:"desc" form:"desc"`
}

// CreateNavGameConfigRequest 创建游戏配置请求
type CreateNavGameConfigRequest struct {
	DownloadUrl  string `json:"download_url" form:"download_url" binding:"required"`   // 下载地址
	AudioUrl     string `json:"audio_url" form:"audio_url"`                            // 播放音频地址
	WebsiteTitle string `json:"website_title" form:"website_title" binding:"required"` // 网站标题
	WebsiteDesc  string `json:"website_desc" form:"website_desc"`                      // 网站描述
	WebsiteIcon  string `json:"website_icon" form:"website_icon"`                      // 网站图标
	WebsiteLogo  string `json:"website_logo" form:"website_logo"`                      // 网站Logo
	Status       int    `json:"status" form:"status"`                                  // 状态
}

// UpdateNavGameConfigRequest 更新游戏配置请求
type UpdateNavGameConfigRequest struct {
	ID           uint   `json:"id" form:"id" binding:"required"`                       // ID
	DownloadUrl  string `json:"download_url" form:"download_url" binding:"required"`   // 下载地址
	AudioUrl     string `json:"audio_url" form:"audio_url"`                            // 播放音频地址
	WebsiteTitle string `json:"website_title" form:"website_title" binding:"required"` // 网站标题
	WebsiteDesc  string `json:"website_desc" form:"website_desc"`                      // 网站描述
	WebsiteIcon  string `json:"website_icon" form:"website_icon"`                      // 网站图标
	WebsiteLogo  string `json:"website_logo" form:"website_logo"`                      // 网站Logo
	Status       int    `json:"status" form:"status"`                                  // 状态
}
