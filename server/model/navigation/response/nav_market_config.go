package response

// NavMarketConfigResponse 市场配置响应
type NavMarketConfigResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Logo        string `json:"logo"`
	JumpUrl     string `json:"jump_url"`
	Type        int    `json:"type"`
	Status      int    `json:"status"`
	IsVisible   int    `json:"is_visible"`
	Sort        int    `json:"sort"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

// NavMarketConfigListResponse 市场配置列表响应
type NavMarketConfigListResponse struct {
	List     []NavMarketConfigResponse `json:"list"`
	Total    int64                     `json:"total"`
	Page     int                       `json:"page"`
	PageSize int                       `json:"pageSize"`
}

// PublicMarketItemResponse 公开市场配置项响应
type PublicMarketItemResponse struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Logo    string `json:"logo"`
	JumpUrl string `json:"jump_url"`
}

// PublicMarketListResponse 公开市场配置列表响应
type PublicMarketListResponse struct {
	List       []PublicMarketItemResponse `json:"list"`
	MarketLogo string                     `json:"market_logo"`
	Icon       string                     `json:"icon"`

	Total    int64 `json:"-"`
	Page     int   `json:"-"`
	PageSize int   `json:"-"`
}
