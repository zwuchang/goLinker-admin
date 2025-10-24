package request

import (
	"goLinker-admin/server/model/common/request"
	"goLinker-admin/server/model/navigation"
)

// NavPWAConfigSearch PWA配置搜索
type NavPWAConfigSearch struct {
	request.PageInfo
	navigation.NavPWAConfig
}

// CreateNavPWAConfigRequest 创建PWA配置请求
type CreateNavPWAConfigRequest struct {
	BackgroundColor string                   `json:"backgroundColor" form:"backgroundColor"`        // 背景色
	ThemeColor      string                   `json:"themeColor" form:"themeColor"`                  // 主题色
	Display         string                   `json:"display" form:"display"`                        // 显示模式
	Name            string                   `json:"name" form:"name" binding:"required"`           // 应用名称
	ShortName       string                   `json:"shortName" form:"shortName" binding:"required"` // 应用短名称
	Orientation     string                   `json:"orientation" form:"orientation"`                // 屏幕方向
	Scope           string                   `json:"scope" form:"scope" binding:"required"`         // 作用域
	StartUrl        string                   `json:"startUrl" form:"startUrl" binding:"required"`   // 启动URL
	Status          int                      `json:"status" form:"status"`                          // 状态
	Icons           []navigation.PWAIconItem `json:"icons" form:"icons"`                            // 图标列表
}

// UpdateNavPWAConfigRequest 更新PWA配置请求
type UpdateNavPWAConfigRequest struct {
	ID              uint                     `json:"id" form:"id" binding:"required"`        // ID
	BackgroundColor *string                  `json:"backgroundColor" form:"backgroundColor"` // 背景色（可选）
	ThemeColor      *string                  `json:"themeColor" form:"themeColor"`           // 主题色（可选）
	Display         *string                  `json:"display" form:"display"`                 // 显示模式（可选）
	Name            *string                  `json:"name" form:"name"`                       // 应用名称（可选）
	ShortName       *string                  `json:"shortName" form:"shortName"`             // 应用短名称（可选）
	Orientation     *string                  `json:"orientation" form:"orientation"`         // 屏幕方向（可选）
	Scope           *string                  `json:"scope" form:"scope"`                     // 作用域（可选）
	StartUrl        *string                  `json:"startUrl" form:"startUrl"`               // 启动URL（可选）
	Status          *int                     `json:"status" form:"status"`                   // 状态（可选）
	Icons           []navigation.PWAIconItem `json:"icons" form:"icons"`                     // 图标列表（可选）
}
