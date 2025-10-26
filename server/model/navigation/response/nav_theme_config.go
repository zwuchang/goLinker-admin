package response

// NavThemeConfigResponse 主题配置响应
type NavThemeConfigResponse struct {
	ID          uint   `json:"id"`          // ID
	Name        string `json:"name"`        // 主题名称
	Description string `json:"description"` // 主题描述
	ConfigJson  string `json:"config_json"` // 主题配置JSON
	IsDefault   int    `json:"is_default"`  // 是否默认主题
	CreatedAt   string `json:"created_at"`  // 创建时间
	UpdatedAt   string `json:"updated_at"`  // 更新时间
}

// NavThemeConfigListResponse 主题配置列表响应
type NavThemeConfigListResponse struct {
	List     []NavThemeConfigResponse `json:"list"`      // 主题配置列表
	Total    int64                    `json:"total"`     // 总数
	Page     int                      `json:"page"`      // 页码
	PageSize int                      `json:"page_size"` // 每页数量
}

// NavThemeConfigDetailResponse 主题配置详情响应
type NavThemeConfigDetailResponse struct {
	ID          uint   `json:"id"`          // ID
	Name        string `json:"name"`        // 主题名称
	Description string `json:"description"` // 主题描述
	ConfigJson  string `json:"config_json"` // 主题配置JSON
	IsDefault   int    `json:"is_default"`  // 是否默认主题
	CreatedAt   string `json:"created_at"`  // 创建时间
	UpdatedAt   string `json:"updated_at"`  // 更新时间
}

// PublicThemeConfigResponse 公开主题配置响应（用于前端主题切换）
type PublicThemeConfigResponse struct {
	ID         uint   `json:"id"`          // ID
	Name       string `json:"name"`        // 主题名称
	ConfigJson string `json:"config_json"` // 主题配置JSON
	IsDefault  int    `json:"is_default"`  // 是否默认主题
}