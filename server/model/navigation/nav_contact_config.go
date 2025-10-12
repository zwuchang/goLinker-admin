package navigation

import (
	"goLinker-admin/server/global"
)

type NavContactConfig struct {
	global.GVA_MODEL
	BannerImage string `json:"bannerImage" form:"bannerImage" gorm:"comment:联系我们的横幅图片"` // 横幅图片
	Email       string `json:"email" form:"email" gorm:"comment:客服邮箱"`                  // 客服邮箱
}

func (NavContactConfig) TableName() string {
	return "nav_contact_config"
}
