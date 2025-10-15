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
