package response

type NavGameConfigResponse struct {
	ID           uint   `json:"id"`
	DownloadUrl  string `json:"download_url"`
	AudioUrl     string `json:"audio_url"`
	WebsiteTitle string `json:"website_title"`
	WebsiteDesc  string `json:"website_desc"`
	WebsiteIcon  string `json:"website_icon"`
	WebsiteLogo  string `json:"website_logo"`
	MarketLogo   string `json:"market_logo"`
	Status       int    `json:"status"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}
