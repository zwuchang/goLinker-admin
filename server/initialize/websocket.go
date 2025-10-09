package initialize

import (
	"goLinker-admin/server/global"
	"goLinker-admin/server/service/system"

	"github.com/olahol/melody"
)

func WebSocket() {
	// 初始化全局WebSocket服务
	global.Websocket = melody.New()

	// 初始化WebSocket服务管理器
	system.GetWebSocketService()
}
