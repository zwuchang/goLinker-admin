package response

import (
	"goLinker-admin/server/model/navigation"
)

// NavScrollNoticeResponse 滚动通知响应
type NavScrollNoticeResponse struct {
	Notice navigation.NavScrollNotice `json:"notice"`
}

// NavScrollNoticeListResponse 滚动通知列表响应
type NavScrollNoticeListResponse struct {
	List     []navigation.NavScrollNotice `json:"list"`
	Total    int64                        `json:"total"`
	Page     int                          `json:"page"`
	PageSize int                          `json:"pageSize"`
}

// PublicScrollNoticeItemResponse 公开滚动通知项响应
type PublicScrollNoticeItemResponse struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// PublicScrollNoticeListResponse 公开滚动通知列表响应
type PublicScrollNoticeListResponse struct {
	List     []PublicScrollNoticeItemResponse `json:"list"`
	Total    int64                            `json:"total"`
	Page     int                              `json:"page"`
	PageSize int                              `json:"pageSize"`
}

// PWAIcon PWA图标配置
type PWAIcon struct {
	Sizes string `json:"sizes"`
	Src   string `json:"src"`
	Type  string `json:"type"`
}

// PWAConfigResponse PWA配置响应
type PWAConfigResponse struct {
	BackgroundColor string    `json:"background_color"`
	ThemeColor      string    `json:"theme_color"`
	Display         string    `json:"display"`
	Icons           []PWAIcon `json:"icons"`
	Name            string    `json:"name"`
	Orientation     string    `json:"orientation"`
	Scope           string    `json:"scope"`
	ShortName       string    `json:"short_name"`
	StartUrl        string    `json:"start_url"`
}
