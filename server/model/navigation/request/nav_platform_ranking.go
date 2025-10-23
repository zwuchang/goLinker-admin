package request

import (
	"goLinker-admin/server/model/common/request"
	"goLinker-admin/server/model/navigation"
)

// NavPlatformRankingSearch 平台排行榜搜索
type NavPlatformRankingSearch struct {
	request.PageInfo
	navigation.NavPlatformRanking
}

// CreateNavPlatformRankingRequest 创建平台排行榜请求
type CreateNavPlatformRankingRequest struct {
	Rank         int     `json:"rank" form:"rank" binding:"required"`                 // 排名
	PlatformName string  `json:"platformName" form:"platformName" binding:"required"` // 平台名称
	Logo         string  `json:"logo" form:"logo"`                                    // 平台Logo地址
	Rating       float64 `json:"rating" form:"rating"`                                // 评分
	Features     string  `json:"features" form:"features"`                            // 特色功能描述
	FeatureType  string  `json:"featureType" form:"featureType"`                      // 功能类型
	VisitUrl     string  `json:"visitUrl" form:"visitUrl"`                            // 访问链接
	IsNew        int     `json:"isNew" form:"isNew"`                                  // 是否新平台
	Status       int     `json:"status" form:"status"`                                // 状态
	IsVisible    int     `json:"isVisible" form:"isVisible"`                          // 是否显示
	Sort         int     `json:"sort" form:"sort"`                                    // 排序
}

// UpdateNavPlatformRankingRequest 更新平台排行榜请求
type UpdateNavPlatformRankingRequest struct {
	ID           uint     `json:"id" form:"id" binding:"required"`  // ID
	Rank         *int     `json:"rank" form:"rank"`                 // 排名（可选）
	PlatformName *string  `json:"platformName" form:"platformName"` // 平台名称（可选）
	Logo         *string  `json:"logo" form:"logo"`                 // 平台Logo地址（可选）
	Rating       *float64 `json:"rating" form:"rating"`             // 评分（可选）
	Features     *string  `json:"features" form:"features"`         // 特色功能描述（可选）
	FeatureType  *string  `json:"featureType" form:"featureType"`   // 功能类型（可选）
	VisitUrl     *string  `json:"visitUrl" form:"visitUrl"`         // 访问链接（可选）
	IsNew        *int     `json:"isNew" form:"isNew"`               // 是否新平台（可选）
	Status       *int     `json:"status" form:"status"`             // 状态（可选）
	IsVisible    *int     `json:"isVisible" form:"isVisible"`       // 是否显示（可选）
	Sort         *int     `json:"sort" form:"sort"`                 // 排序（可选）
}
