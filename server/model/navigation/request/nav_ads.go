package request

// NavAdsSearch 广告搜索请求参数
type NavAdsSearch struct {
	Page     int `json:"page" form:"page"`         // 页码
	PageSize int `json:"pageSize" form:"pageSize"` // 每页数量
}
