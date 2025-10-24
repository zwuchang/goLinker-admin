package response

import (
	"goLinker-admin/server/model/navigation"
)

// NavPWAConfigResponse PWA配置响应
type NavPWAConfigResponse struct {
	Config navigation.NavPWAConfig `json:"config"`
}

// NavPWAConfigListResponse PWA配置列表响应
type NavPWAConfigListResponse struct {
	List     []navigation.NavPWAConfig `json:"list"`
	Total    int64                     `json:"total"`
	Page     int                       `json:"page"`
	PageSize int                       `json:"pageSize"`
}

// PWAIconItem PWA图标配置（用于公开接口）
type PWAIconItem struct {
	Sizes string `json:"sizes"`
	Src   string `json:"src"`
	Type  string `json:"type"`
}

// PWAConfigItemResponse PWA配置响应（用于公开接口）
type PWAConfigItemResponse struct {
	BackgroundColor string        `json:"background_color"`
	ThemeColor      string        `json:"theme_color"`
	Display         string        `json:"display"`
	Icons           []PWAIconItem `json:"icons"`
	Name            string        `json:"name"`
	Orientation     string        `json:"orientation"`
	Scope           string        `json:"scope"`
	ShortName       string        `json:"short_name"`
	StartUrl        string        `json:"start_url"`
}
