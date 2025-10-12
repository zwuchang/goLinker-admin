package response

import "goLinker-admin/server/model/navigation"

// NavGameCategoryResponse 游戏类别响应
type NavGameCategoryResponse struct {
	Category navigation.NavGameCategory `json:"category"`
}

// NavGameResponse 游戏响应
type NavGameResponse struct {
	Game navigation.NavGame `json:"game"`
}

// NavGameListResponse 游戏列表响应
type NavGameListResponse struct {
	List     []navigation.NavGame `json:"list"`
	Total    int64                `json:"total"`
	Page     int                  `json:"page"`
	PageSize int                  `json:"pageSize"`
}
