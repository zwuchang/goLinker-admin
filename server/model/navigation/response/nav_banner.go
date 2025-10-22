package response

import (
	"goLinker-admin/server/model/navigation"
)

// NavBannerResponse Banner响应
type NavBannerResponse struct {
	Banner navigation.NavBanner `json:"banner"`
}

// NavBannerListResponse Banner列表响应
type NavBannerListResponse struct {
	List     []navigation.NavBanner `json:"list"`
	Total    int64                  `json:"total"`
	Page     int                    `json:"page"`
	PageSize int                    `json:"pageSize"`
}

// PublicBannerItemResponse 公开Banner项响应
type PublicBannerItemResponse struct {
	ID        uint   `json:"id"`
	MediaType int    `json:"mediaType"`
	MediaUrl  string `json:"mediaUrl"`
	LinkUrl   string `json:"linkUrl"`
}

// PublicBannerListResponse 公开Banner列表响应
type PublicBannerListResponse struct {
	List     []PublicBannerItemResponse `json:"list"`
	Total    int64                      `json:"total"`
	Page     int                        `json:"page"`
	PageSize int                        `json:"pageSize"`
}
