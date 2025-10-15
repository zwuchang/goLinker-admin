package request

import "goLinker-admin/server/model/common/request"

// NavAccessStatsSearch 访问统计搜索参数
type NavAccessStatsSearch struct {
	request.PageInfo
	ApiPath    string `json:"apiPath" form:"apiPath"`       // API路径搜索
	Method     string `json:"method" form:"method"`         // HTTP方法搜索
	ClientIP   string `json:"clientIp" form:"clientIp"`     // 客户端IP搜索
	StatusCode *int   `json:"statusCode" form:"statusCode"` // 状态码搜索
	Country    string `json:"country" form:"country"`       // 国家搜索
	Device     string `json:"device" form:"device"`         // 设备类型搜索
	Browser    string `json:"browser" form:"browser"`       // 浏览器搜索
	OS         string `json:"os" form:"os"`                 // 操作系统搜索
	StartTime  string `json:"startTime" form:"startTime"`   // 开始时间
	EndTime    string `json:"endTime" form:"endTime"`       // 结束时间
	OrderBy    string `json:"orderBy" form:"orderBy"`       // 排序字段
	OrderType  string `json:"orderType" form:"orderType"`   // 排序类型 asc/desc
}

// NavAccessStatsSummaryRequest 访问统计汇总请求
type NavAccessStatsSummaryRequest struct {
	StartTime string `json:"startTime"` // 开始时间
	EndTime   string `json:"endTime"`   // 结束时间
}
