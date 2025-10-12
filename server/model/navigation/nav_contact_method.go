package navigation

import (
	"goLinker-admin/server/global"
)

type NavContactMethod struct {
	global.GVA_MODEL
	Image       string `json:"image" form:"image" gorm:"comment:联系方式图标"`           // 联系方式图标
	JumpUrl     string `json:"jumpUrl" form:"jumpUrl" gorm:"comment:跳转地址"`         // 跳转地址
	DisplayName string `json:"displayName" form:"displayName" gorm:"comment:显示名称"` // 显示名称
	Sort        int    `json:"sort" form:"sort" gorm:"comment:排序"`                 // 排序
	Status      int    `json:"status" form:"status" gorm:"comment:状态;default:1"`   // 状态 1:启用 0:禁用
}

func (NavContactMethod) TableName() string {
	return "nav_contact_method"
}
