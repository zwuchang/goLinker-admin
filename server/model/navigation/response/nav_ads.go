package response

// NavAdsItemResponse 广告项目响应
type NavAdsItemResponse struct {
	ID          uint   `json:"id"`           // 游戏ID
	AdName      string `json:"ad_name"`      // 广告名称
	Image       string `json:"image"`        // 广告图片
	RedirectUrl string `json:"redirect_url"` // 跳转地址
}

// NavAdsListResponse 广告列表响应
type NavAdsListResponse struct {
	List     []NavAdsItemResponse `json:"list"`     // 广告列表
	Total    int64                `json:"total"`    // 总数
	Page     int                  `json:"page"`     // 当前页
	PageSize int                  `json:"pageSize"` // 每页数量
}





