package response

type NavGameConfigResponse struct {
	ID             uint   `json:"id"`
	DownloadUrl    string `json:"download_url"`
	AudioUrl       string `json:"audio_url"`
	WebsiteTitle   string `json:"website_title"`
	WebsiteDesc    string `json:"website_desc"`
	WebsiteIcon    string `json:"website_icon"`
	WebsiteLogo    string `json:"website_logo"`
	MarketLogo     string `json:"market_logo"`
	FloatingStatus int    `json:"floating_status"`
	FloatingIcon1  string `json:"floating_icon1"`
	FloatingIcon2  string `json:"floating_icon2"`
	FloatingIcon3  string `json:"floating_icon3"`
	Status         int    `json:"status"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}
