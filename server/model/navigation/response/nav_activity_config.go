package response

import "goLinker-admin/server/model/navigation"

// NavActivityConfigResponse 活动配置响应
type NavActivityConfigResponse struct {
	ID           uint   `json:"id"`           // ID
	Title        string `json:"title"`        // 显示标题
	Image        string `json:"image"`        // 活动图片
	JumpUrl      string `json:"jumpUrl"`      // 跳转地址
	CategoryName string `json:"categoryName"` // 类别名称
	CategoryIcon string `json:"categoryIcon"` // 类别图标
	Content      string `json:"content"`      // 活动内容
	Status       int    `json:"status"`       // 状态
	IsVisible    int    `json:"isVisible"`    // 是否显示
	Sort         int    `json:"sort"`         // 排序
	CreatedAt    string `json:"created_at"`   // 创建时间
	UpdatedAt    string `json:"updated_at"`   // 更新时间
}

// NavActivityConfigListResponse 活动配置列表响应
type NavActivityConfigListResponse struct {
	List     []NavActivityConfigResponse `json:"list"`      // 活动配置列表
	Total    int64                       `json:"total"`     // 总数
	Page     int                         `json:"page"`      // 页码
	PageSize int                         `json:"page_size"` // 每页数量
}

// NavActivityConfigDetailResponse 活动配置详情响应
type NavActivityConfigDetailResponse struct {
	Config navigation.NavActivityConfig `json:"config"`
}

// PublicActivityItemResponse 公开活动项响应（用于前端显示）
type PublicActivityItemResponse struct {
	ID           uint   `json:"id"`           // ID
	Title        string `json:"title"`        // 显示标题
	Image        string `json:"image"`        // 活动图片
	JumpUrl      string `json:"jumpUrl"`      // 跳转地址
	CategoryName string `json:"categoryName"` // 类别名称
	CategoryIcon string `json:"categoryIcon"` // 类别图标
	Content      string `json:"content"`      // 活动内容
}

// PublicActivityListResponse 公开活动列表响应
type PublicActivityListResponse struct {
	List     []PublicActivityItemResponse `json:"list"`      // 活动配置列表
	Total    int64                        `json:"total"`     // 总数
	Page     int                          `json:"page"`      // 页码
	PageSize int                          `json:"page_size"` // 每页数量
}
