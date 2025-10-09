package system

import (
	"encoding/json"
	"sync"
	"goLinker-admin/server/global"

	"github.com/olahol/melody"
	"go.uber.org/zap"
)

type WebSocketService struct {
	sync.RWMutex
	// 用户ID到会话的映射
	sessions map[uint]*melody.Session
}

type WSMessage struct {
	Type    string      `json:"type"`    // 消息类型
	UserID  uint        `json:"userId"`  // 目标用户ID
	Content interface{} `json:"content"` // 消息内容
}

var (
	wsService     *WebSocketService
	wsServiceOnce sync.Once
)

// GetWebSocketService 获取WebSocket服务单例
func GetWebSocketService() *WebSocketService {
	wsServiceOnce.Do(func() {
		wsService = &WebSocketService{
			sessions: make(map[uint]*melody.Session),
		}
		// 初始化WebSocket事件处理
		wsService.initEventHandlers()
	})
	return wsService
}

// 初始化WebSocket事件处理
func (ws *WebSocketService) initEventHandlers() {
	// 处理新连接
	global.Websocket.HandleConnect(func(s *melody.Session) {
		ws.handleConnect(s)
	})

	// 处理断开连接
	global.Websocket.HandleDisconnect(func(s *melody.Session) {
		ws.handleDisconnect(s)
	})

	// 处理消息
	global.Websocket.HandleMessage(func(s *melody.Session, msg []byte) {
		ws.handleMessage(s, msg)
	})

	// 处理错误
	global.Websocket.HandleError(func(s *melody.Session, err error) {
		global.GVA_LOG.Error("websocket error", zap.Error(err))
	})
}

// 处理新连接
func (ws *WebSocketService) handleConnect(s *melody.Session) {
	userID := ws.getUserIDFromSession(s)
	if userID == 0 {
		return
	}

	ws.Lock()
	ws.sessions[userID] = s
	ws.Unlock()

	global.GVA_LOG.Info("websocket connected", zap.Uint("userId", userID))
}

// 处理断开连接
func (ws *WebSocketService) handleDisconnect(s *melody.Session) {
	userID := ws.getUserIDFromSession(s)
	if userID == 0 {
		return
	}

	ws.Lock()
	delete(ws.sessions, userID)
	ws.Unlock()

	global.GVA_LOG.Info("websocket disconnected", zap.Uint("userId", userID))
}

// 处理消息
func (ws *WebSocketService) handleMessage(s *melody.Session, msg []byte) {
	var message WSMessage
	if err := json.Unmarshal(msg, &message); err != nil {
		global.GVA_LOG.Error("invalid websocket message", zap.Error(err))
		return
	}

	// 根据消息类型处理
	switch message.Type {
	case "cron":
		// 获取所有定时任务信息
		cronService := &CronService{}
		// 获取所有定时任务状态
		cronInfos := cronService.GetAllCronStatus()
		// 构造返回消息
		response := WSMessage{
			Type:    "cron_list",
			Content: cronInfos,
		}

		// 发送消息
		data, _ := json.Marshal(response)
		s.Write(data)
	case "ping":
		s.Write([]byte(`{"type":"pong"}`))
	default:
		global.GVA_LOG.Info("received message",
			zap.String("type", message.Type),
			zap.Any("content", message.Content))
	}
}

// SendToUser 发送消息给指定用户
func (ws *WebSocketService) SendToUser(userID uint, message interface{}) error {
	ws.RLock()
	session, exists := ws.sessions[userID]
	ws.RUnlock()

	if !exists {
		return nil // 用户不在线，静默返回
	}

	data, err := json.Marshal(message)
	if err != nil {
		return err
	}

	return session.Write(data)
}

// Broadcast 广播消息给所有在线用户
func (ws *WebSocketService) Broadcast(message interface{}) error {
	data, err := json.Marshal(message)
	if err != nil {
		return err
	}

	return global.Websocket.Broadcast(data)
}

// 从Session中获取用户ID
func (ws *WebSocketService) getUserIDFromSession(s *melody.Session) uint {
	if userID, exists := s.Get("userId"); exists {
		if id, ok := userID.(uint); ok {
			return id
		}
	}
	return 0
}

// GetOnlineUsers 获取在线用户列表
func (ws *WebSocketService) GetOnlineUsers() []uint {
	ws.RLock()
	defer ws.RUnlock()

	users := make([]uint, 0, len(ws.sessions))
	for userID := range ws.sessions {
		users = append(users, userID)
	}
	return users
}
