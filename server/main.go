package main

import (
	"goLinker-admin/server/core"
	"goLinker-admin/server/global"
	"goLinker-admin/server/initialize"
	"goLinker-admin/server/service/navigation"

	_ "go.uber.org/automaxprocs"
	"go.uber.org/zap"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// 这部分 @Tag 设置用于排序, 需要排序的接口请按照下面的格式添加
// swag init 对 @Tag 只会从入口文件解析, 默认 main.go
// 也可通过 --generalInfo flag 指定其他文件
// @Tag.Name        Base
// @Tag.Name        SysUser
// @Tag.Description 用户

// @title                       Gin-Vue-Admin Swagger API接口文档
// @version                     v2.8.5
// @description                 使用gin+vue进行极速开发的全栈开发基础平台
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        x-token
// @BasePath                    /
func main() {
	// 初始化系统
	initializeSystem()

	// 程序结束前关闭IP地址库
	defer initialize.CloseIPSearcher()

	// 运行服务器
	core.RunServer()
}

// initializeSystem 初始化系统所有组件
// 提取为单独函数以便于系统重载时调用
func initializeSystem() {
	global.GVA_VP = core.Viper() // 初始化Viper
	initialize.OtherInit()
	global.GVA_LOG = core.Zap() // 初始化zap日志库
	zap.ReplaceGlobals(global.GVA_LOG)
	global.GVA_DB = initialize.Gorm() // gorm连接数据库
	initialize.InitIpSearcher()       // 初始化IP地址库

	initialize.Timer()
	initialize.DBList()
	initialize.WebSocket()     // websocket
	initialize.SetupHandlers() // 注册全局函数

	if global.GVA_DB != nil {
		initialize.RegisterTables() // 初始化表

		// 检查并创建导航管理菜单
		err := navigation.ServiceGroupApp.NavMenuCheckService.CheckAndCreateNavigationMenus()
		if err != nil {
			global.GVA_LOG.Error("导航菜单检查失败", zap.Error(err))
		} else {
			global.GVA_LOG.Info("导航菜单检查完成")
		}

		// 检查并创建导航管理API
		err = navigation.ServiceGroupApp.NavApiCheckService.CheckAndCreateNavigationApis()
		if err != nil {
			global.GVA_LOG.Error("导航API检查失败", zap.Error(err))
		} else {
			global.GVA_LOG.Info("导航API检查完成")
		}
	}
}
