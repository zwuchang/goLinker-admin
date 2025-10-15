package navigation

import "goLinker-admin/server/global"

// NavAccessStats 访问统计模型
type NavAccessStats struct {
	global.GVA_MODEL
	ApiPath      string `json:"apiPath" form:"apiPath" gorm:"type:varchar(200);comment:API路径"`  // API路径
	Method       string `json:"method" form:"method" gorm:"type:varchar(10);comment:HTTP方法"`    // HTTP方法
	ClientIP     string `json:"clientIp" form:"clientIp" gorm:"type:varchar(50);comment:客户端IP"` // 客户端IP
	UserAgent    string `json:"userAgent" form:"userAgent" gorm:"type:text;comment:用户代理"`       // 用户代理
	Referer      string `json:"referer" form:"referer" gorm:"type:varchar(500);comment:来源页面"`   // 来源页面
	ResponseTime int    `json:"responseTime" form:"responseTime" gorm:"comment:响应时间(毫秒)"`       // 响应时间(毫秒)
	StatusCode   int    `json:"statusCode" form:"statusCode" gorm:"comment:响应状态码"`              // 响应状态码
	RequestSize  int    `json:"requestSize" form:"requestSize" gorm:"comment:请求大小(字节)"`         // 请求大小(字节)
	ResponseSize int    `json:"responseSize" form:"responseSize" gorm:"comment:响应大小(字节)"`       // 响应大小(字节)
	Country      string `json:"country" form:"country" gorm:"type:varchar(50);comment:国家"`      // 国家
	City         string `json:"city" form:"city" gorm:"type:varchar(50);comment:城市"`            // 城市
	Device       string `json:"device" form:"device" gorm:"type:varchar(50);comment:设备类型"`      // 设备类型
	Browser      string `json:"browser" form:"browser" gorm:"type:varchar(50);comment:浏览器"`     // 浏览器
	OS           string `json:"os" form:"os" gorm:"type:varchar(50);comment:操作系统"`              // 操作系统
}

// TableName 表名
func (NavAccessStats) TableName() string {
	return "nav_access_stats"
}
