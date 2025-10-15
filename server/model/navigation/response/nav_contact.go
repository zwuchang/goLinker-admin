package response

import "goLinker-admin/server/model/navigation"

// ContactMethodInfo 联系方式信息
type ContactMethodInfo struct {
	ID          uint   `json:"id"`
	Image       string `json:"image"`       // 联系方式图标
	JumpUrl     string `json:"jumpUrl"`     // 跳转地址
	DisplayName string `json:"displayName"` // 显示名称
	Sort        int    `json:"-"`           // 排序
}

// ContactConfigInfo 联系配置信息
type ContactConfigInfo struct {
	BannerImage string `json:"bannerImage"` // 横幅图片
	Email       string `json:"email"`       // 客服邮箱
}

// PublicContactResponse 公开联系方式响应
type PublicContactResponse struct {
	ContactMethods []ContactMethodInfo `json:"contactMethods"` // 联系方式列表
	ContactConfig  ContactConfigInfo   `json:"contactConfig"`  // 联系配置
	UpdateTime     string              `json:"updateTime"`     // 更新时间
}

// NavContactConfigResponse 联系配置响应
type NavContactConfigResponse struct {
	ContactConfig navigation.NavContactConfig `json:"contactConfig"`
}

// NavContactMethodResponse 联系方式响应
type NavContactMethodResponse struct {
	ContactMethod navigation.NavContactMethod `json:"contactMethod"`
}

// NavGameArticleResponse 游戏文章响应
type NavGameArticleResponse struct {
	GameId     uint   `json:"gameId"`     // 游戏ID
	Title      string `json:"title"`      // 游戏标题
	Article    string `json:"article"`    // 文章内容
	UpdateTime string `json:"updateTime"` // 更新时间
}
