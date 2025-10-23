package response

import (
	"goLinker-admin/server/model/navigation"
)

// NavPlatformRankingResponse 平台排行榜响应
type NavPlatformRankingResponse struct {
	Ranking navigation.NavPlatformRanking `json:"ranking"`
}

// NavPlatformRankingListResponse 平台排行榜列表响应
type NavPlatformRankingListResponse struct {
	List     []navigation.NavPlatformRanking `json:"list"`
	Total    int64                           `json:"total"`
	Page     int                             `json:"page"`
	PageSize int                             `json:"pageSize"`
}

// PublicPlatformRankingItemResponse 公开平台排行榜项响应
type PublicPlatformRankingItemResponse struct {
	ID           uint    `json:"id"`
	Rank         int     `json:"rank"`
	PlatformName string  `json:"platformName"`
	Logo         string  `json:"logo"`
	Rating       float64 `json:"rating"`
	Features     string  `json:"features"`
	FeatureType  string  `json:"featureType"`
	VisitUrl     string  `json:"visitUrl"`
	IsNew        int     `json:"isNew"`
}

// PublicPlatformRankingListResponse 公开平台排行榜列表响应
type PublicPlatformRankingListResponse struct {
	List     []PublicPlatformRankingItemResponse `json:"list"`
	Total    int64                               `json:"total"`
	Page     int                                 `json:"page"`
	PageSize int                                 `json:"pageSize"`
}

