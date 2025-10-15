package response

// PublicSiteInfoResponse 公开网站信息响应
type PublicSiteInfoResponse struct {
	WebsiteTitle string `json:"websiteTitle"` // 网站标题
	WebsiteDesc  string `json:"websiteDesc"`  // 网站描述
	DownloadUrl  string `json:"downloadUrl"`  // 下载地址
	AudioUrl     string `json:"audioUrl"`     // 音频地址
	WebsiteIcon  string `json:"websiteIcon"`  // 网站图标
	WebsiteLogo  string `json:"websiteLogo"`  // 网站Logo
	UpdateTime   string `json:"updateTime"`   // 更新时间
}
