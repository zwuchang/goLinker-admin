package response

// NavPlatformConfigResponse 平台游戏配置响应
type NavPlatformConfigResponse struct {
	ID           uint   `json:"id"`            // ID
	PlatformName string `json:"platform_name"` // 平台名称
	PlatformIcon string `json:"platform_icon"` // 平台图标地址
	PlatformApi  string `json:"platform_api"`  // 平台接口地址
	Sort         int    `json:"sort"`          // 排序
	Status       int    `json:"status"`        // 状态
	Description  string `json:"description"`   // 平台描述
	IsVisible    int    `json:"is_visible"`    // 是否可见
	CreatedAt    string `json:"created_at"`    // 创建时间
	UpdatedAt    string `json:"updated_at"`    // 更新时间
}

// PublicPlatformMenuResponse 公开平台菜单响应
type PublicPlatformMenuResponse struct {
	ID           uint   `json:"id"`            // ID
	PlatformName string `json:"platform_name"` // 平台名称
	PlatformIcon string `json:"platform_icon"` // 平台图标地址
}



