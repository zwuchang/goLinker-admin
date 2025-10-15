package response

// NavAccessStatsResponse 访问统计响应
type NavAccessStatsResponse struct {
	ID           uint   `json:"id"`
	ApiPath      string `json:"apiPath"`      // API路径
	Method       string `json:"method"`       // HTTP方法
	ClientIP     string `json:"clientIp"`     // 客户端IP
	UserAgent    string `json:"userAgent"`    // 用户代理
	Referer      string `json:"referer"`      // 来源页面
	ResponseTime int    `json:"responseTime"` // 响应时间(毫秒)
	StatusCode   int    `json:"statusCode"`   // 响应状态码
	RequestSize  int    `json:"requestSize"`  // 请求大小(字节)
	ResponseSize int    `json:"responseSize"` // 响应大小(字节)
	Country      string `json:"country"`      // 国家
	City         string `json:"city"`         // 城市
	Device       string `json:"device"`       // 设备类型
	Browser      string `json:"browser"`      // 浏览器
	OS           string `json:"os"`           // 操作系统
	CreatedAt    string `json:"createdAt"`    // 创建时间
}

// NavAccessStatsSummaryResponse 访问统计汇总响应
type NavAccessStatsSummaryResponse struct {
	TotalRequests   int64          `json:"totalRequests"`   // 总请求数
	UniqueIPs       int64          `json:"uniqueIPs"`       // 独立IP数
	AvgResponseTime float64        `json:"avgResponseTime"` // 平均响应时间
	SuccessRate     float64        `json:"successRate"`     // 成功率
	TopApiPaths     []ApiPathCount `json:"topApiPaths"`     // 热门API路径
	TopCountries    []CountryCount `json:"topCountries"`    // 热门国家
	TopDevices      []DeviceCount  `json:"topDevices"`      // 热门设备
	TopBrowsers     []BrowserCount `json:"topBrowsers"`     // 热门浏览器
}

// ApiPathCount API路径统计
type ApiPathCount struct {
	ApiPath string `json:"apiPath"`
	Count   int64  `json:"count"`
}

// CountryCount 国家统计
type CountryCount struct {
	Country string `json:"country"`
	Count   int64  `json:"count"`
}

// DeviceCount 设备统计
type DeviceCount struct {
	Device string `json:"device"`
	Count  int64  `json:"count"`
}

// BrowserCount 浏览器统计
type BrowserCount struct {
	Browser string `json:"browser"`
	Count   int64  `json:"count"`
}
