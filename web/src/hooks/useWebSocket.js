import { ref } from 'vue'
import { useUserStore } from '@/pinia/modules/user'

const WS_URL = import.meta.env.VITE_WS || 'ws://127.0.0.1:8881/ws/go'

class WebSocketClient {
    constructor() {
        this.ws = null
        this.isConnected = ref(false)
        this.messageHandlers = new Set()
        this.reconnectTimer = null
        this.heartbeatTimer = null
        this.userStore = useUserStore()
    }

    connect() {
        if (this.ws) return

        try {
            const token = this.userStore.token
            const url = `${WS_URL}?token=${token}`
            this.ws = new WebSocket(url)

            this.ws.onopen = () => {
                this.isConnected.value = true
                this.startHeartbeat()
                console.log('WebSocket connected')
            }

            this.ws.onmessage = (event) => {
                try {
                    const data = JSON.parse(event.data)
                    if (data.type === 'pong') return
                    this.messageHandlers.forEach(handler => handler(data))
                } catch (error) {
                    console.error('Failed to parse WebSocket message:', error)
                }
            }

            this.ws.onclose = () => {
                this.isConnected.value = false
                this.stopHeartbeat()
                this.reconnect()
                console.log('WebSocket disconnected')
            }

            this.ws.onerror = (error) => {
                console.error('WebSocket error:', error)
                this.ws.close()
            }
        } catch (error) {
            console.error('Failed to create WebSocket:', error)
            this.reconnect()
        }
    }

    disconnect() {
        if (this.ws) {
            this.ws.close()
            this.ws = null
        }
        this.stopHeartbeat()
        this.stopReconnect()
    }

    reconnect() {
        this.stopReconnect()
        if (this.ws) {
            // 保存当前状态，避免 onclose 触发时再次调用 reconnect
            const wasConnected = this.ws.readyState === WebSocket.OPEN
            this.ws.onclose = null
            this.ws.onerror = null
            this.ws.close()
            this.ws = null

            // 如果之前是连接状态，手动更新状态
            if (wasConnected) {
                this.isConnected.value = false
                this.stopHeartbeat()
            }
        }
        console.log('立即尝试重连...')
        setTimeout(() => {
            this.connect()
        }, 100)

        // 设置备用重连定时器
        this.reconnectTimer = setTimeout(() => {
            console.log('备用重连尝试...')
            if (!this.isConnected.value && !this.ws) {
                this.connect()
            }
        }, 5000)

    }

    stopReconnect() {
        if (this.reconnectTimer) {
            clearTimeout(this.reconnectTimer)
            this.reconnectTimer = null
        }
    }

    startHeartbeat() {
        this.stopHeartbeat()
        this.heartbeatTimer = setInterval(() => {
            this.sendMessage({ type: 'ping' })
        }, 30000)
    }

    stopHeartbeat() {
        if (this.heartbeatTimer) {
            clearInterval(this.heartbeatTimer)
            this.heartbeatTimer = null
        }
    }

    sendMessage(message) {
        if (!this.ws || this.ws.readyState !== WebSocket.OPEN) {
            console.warn('WebSocket is not connected')
            return
        }

        try {
            this.ws.send(JSON.stringify(message))
        } catch (error) {
            console.error('Failed to send message:', error)
        }
    }

    onMessage(handler) {
        this.messageHandlers.add(handler)
    }

    offMessage(handler) {
        this.messageHandlers.delete(handler)
    }
}

// Singleton instance
let wsClient = null

export function useWebSocket() {
    if (!wsClient) {
        wsClient = new WebSocketClient()
        wsClient.connect()
    }
    return wsClient
}