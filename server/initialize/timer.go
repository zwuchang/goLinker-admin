package initialize

import (
	"fmt"
	"goLinker-admin/server/model/system"
	"goLinker-admin/server/task"
	"time"

	"go.uber.org/zap"

	"goLinker-admin/server/global"
)

func Timer() {
	go func() {

		// 等待数据库初始化
		for i := 0; i < 60; i++ { // 最多等待60秒
			if global.GVA_DB != nil {
				break
			}
			time.Sleep(time.Second)
		}

		if global.GVA_DB == nil {
			global.GVA_LOG.Error("数据库未初始化，定时任务启动失败")
			return
		}

		var err error

		// 从数据库获取定时任务配置
		var cronConfigs []system.SysCron
		err = global.GVA_DB.Find(&cronConfigs).Error
		if err != nil {
			global.GVA_LOG.Error("获取定时任务配置失败", zap.Error(err))
		}

		// 配置映射
		configMap := make(map[string]system.SysCron)
		for _, config := range cronConfigs {
			configMap[config.Name] = config
		}

		// 清理DB定时任务
		if config, exists := configMap[global.CronClearDbName]; exists {
			// 数据库中存在该配置，直接使用
			_, err = global.AddCronTask(global.CronClearDbName, config.Spec, func() error {
				fmt.Println("添加", global.CronClearDbName)
				return task.ClearTable(global.GVA_DB) // 定时任务方法定在task文件包中
			}, config.Description)
			if err != nil {
				fmt.Println("add timer error:", err)
			}
		} else {
			// 使用默认配置并保存到数据库
			defaultConfig := system.SysCron{
				Name:        global.CronClearDbName,
				Spec:        "@daily",
				Description: "定时清理数据库【日志，黑名单】内容",
			}
			err = global.GVA_DB.Create(&defaultConfig).Error
			if err != nil {
				fmt.Println("保存定时任务配置失败:", err)
			}
			_, err := global.AddCronTask(defaultConfig.Name, defaultConfig.Spec, func() error {
				return task.ClearTable(global.GVA_DB)
			}, defaultConfig.Description)
			if err != nil {
				fmt.Println("add timer error:", err)
			}
		}

	}()
}
