package initialize

import (
	"goLinker-admin/server/global"
)

func bizModel() error {
	db := global.GVA_DB
	// 这里可以添加业务相关的模型
	err := db.AutoMigrate(
	// 如果有其他业务模型，可以在这里添加
	)
	if err != nil {
		return err
	}
	return nil
}
