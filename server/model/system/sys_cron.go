package system

import "goLinker-admin/server/global"

// SysCron 定时任务配置表
type SysCron struct {
	global.GVA_MODEL
	Name        string `json:"name" gorm:"column:name;comment:任务名称"`
	Spec        string `json:"spec" gorm:"column:spec;comment:执行规则"`
	Description string `json:"description" gorm:"column:description;comment:任务描述"`
}

func (SysCron) TableName() string {
	return "sys_cron"
}
