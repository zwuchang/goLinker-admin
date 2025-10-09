package system

import (
	"goLinker-admin/server/global"

	"github.com/gin-gonic/gin"
)

type WebSocketApi struct{}

// Upgrade 升级为 websocket 服务
func (w *WebSocketApi) Upgrade(c *gin.Context) {
	global.Websocket.HandleRequestWithKeys(c.Writer, c.Request, nil)
}
