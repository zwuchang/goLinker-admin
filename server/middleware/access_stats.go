package middleware

import (
	"goLinker-admin/server/global"
	"goLinker-admin/server/model/navigation"
	"goLinker-admin/server/service"
	"goLinker-admin/server/service/system"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// AccessStatsMiddleware 访问统计中间件
func AccessStatsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		// 获取请求信息
		clientIP := getClientIP(c)
		userAgent := c.GetHeader("User-Agent")
		referer := c.GetHeader("Referer")
		apiPath := c.Request.URL.Path
		method := c.Request.Method

		// 解析用户代理信息
		device, browser, os := parseUserAgent(userAgent)

		// 获取地理位置信息（简化版，实际项目中可以使用IP地理位置服务）
		country, city := getLocationByIP(clientIP)

		// 计算请求大小
		requestSize := int(c.Request.ContentLength)
		if requestSize < 0 {
			requestSize = 0
		}

		// 继续处理请求
		c.Next()

		// 计算响应时间
		responseTime := int(time.Since(startTime).Milliseconds())

		// 获取响应状态码
		statusCode := c.Writer.Status()

		// 计算响应大小
		responseSize := c.Writer.Size()
		if responseSize < 0 {
			responseSize = 0
		}

		// 创建访问统计记录
		stats := navigation.NavAccessStats{
			ApiPath:      apiPath,
			Method:       method,
			ClientIP:     clientIP,
			UserAgent:    userAgent,
			Referer:      referer,
			ResponseTime: responseTime,
			StatusCode:   statusCode,
			RequestSize:  requestSize,
			ResponseSize: responseSize,
			Country:      country,
			City:         city,
			Device:       device,
			Browser:      browser,
			OS:           os,
		}

		// 异步保存统计数据
		go func(statsCopy navigation.NavAccessStats) {
			err := service.ServiceGroupApp.NavigationServiceGroup.NavAccessStatsService.CreateAccessStats(statsCopy)
			if err != nil {
				global.GVA_LOG.Error("保存访问统计失败", zap.Error(err))
			}
		}(stats)

		// 记录访问日志
		global.GVA_LOG.Info("API访问统计",
			zap.String("apiPath", apiPath),
			zap.String("method", method),
			zap.String("clientIP", clientIP),
			zap.Int("statusCode", statusCode),
			zap.Int("responseTime", responseTime),
			zap.String("country", country),
			zap.String("device", device),
		)
	}
}

// getClientIP 获取客户端真实IP
func getClientIP(c *gin.Context) string {
	// 检查 X-Forwarded-For 头
	if xff := c.GetHeader("X-Forwarded-For"); xff != "" {
		ips := strings.Split(xff, ",")
		if len(ips) > 0 {
			return strings.TrimSpace(ips[0])
		}
	}

	// 检查 X-Real-IP 头
	if xri := c.GetHeader("X-Real-IP"); xri != "" {
		return xri
	}

	// 检查 X-Forwarded-Proto 头
	if xfp := c.GetHeader("X-Forwarded-Proto"); xfp != "" {
		return c.ClientIP()
	}

	return c.ClientIP()
}

// parseUserAgent 解析用户代理字符串
func parseUserAgent(userAgent string) (device, browser, os string) {
	ua := strings.ToLower(userAgent)

	// 设备类型检测
	if strings.Contains(ua, "mobile") || strings.Contains(ua, "android") || strings.Contains(ua, "iphone") {
		device = "Mobile"
	} else if strings.Contains(ua, "tablet") || strings.Contains(ua, "ipad") {
		device = "Tablet"
	} else {
		device = "Desktop"
	}

	// 浏览器检测
	if strings.Contains(ua, "chrome") && !strings.Contains(ua, "edge") {
		browser = "Chrome"
	} else if strings.Contains(ua, "firefox") {
		browser = "Firefox"
	} else if strings.Contains(ua, "safari") && !strings.Contains(ua, "chrome") {
		browser = "Safari"
	} else if strings.Contains(ua, "edge") {
		browser = "Edge"
	} else if strings.Contains(ua, "opera") {
		browser = "Opera"
	} else {
		browser = "Other"
	}

	// 操作系统检测
	if strings.Contains(ua, "windows") {
		os = "Windows"
	} else if strings.Contains(ua, "mac") {
		os = "macOS"
	} else if strings.Contains(ua, "linux") {
		os = "Linux"
	} else if strings.Contains(ua, "android") {
		os = "Android"
	} else if strings.Contains(ua, "ios") || strings.Contains(ua, "iphone") || strings.Contains(ua, "ipad") {
		os = "iOS"
	} else {
		os = "Other"
	}

	return device, browser, os
}

// getLocationByIP 根据IP获取地理位置（简化版）
func getLocationByIP(ip string) (country, city string) {
	// 本地IP检测
	global.GVA_LOG.Debug("获取地理位置信息", zap.String("ip", ip))
	if ip == "127.0.0.1" || ip == "::1" || strings.HasPrefix(ip, "192.168.") {
		return "Local", "Local"
	}

	// 检查IP数据库是否可用
	if global.IpSearcher == nil {
		global.GVA_LOG.Debug("IP数据库未初始化，跳过地理位置查询", zap.String("ip", ip))
		return "Unknown", "Unknown"
	}

	// 安全地查询IP地理位置
	defer func() {
		if r := recover(); r != nil {
			global.GVA_LOG.Error("IP地理位置查询panic",
				zap.String("ip", ip),
				zap.Any("panic", r))
		}
	}()

	location := system.IpServiceApp.GetLocation(ip)
	if location == "" {
		return "Unknown", "Unknown"
	}

	global.GVA_LOG.Debug("IP地理位置信息", zap.String("location", location))

	// 解析地理位置信息
	country, city = parseLocation(location)
	return country, city
}

// parseLocation 解析IP库返回的地理位置信息
func parseLocation(location string) (country, city string) {
	if location == "" {
		return "Unknown", "Unknown"
	}

	// 处理IANA特殊地址
	if strings.Contains(location, "IANA") {
		if strings.Contains(location, "本机地址") || strings.Contains(location, "环回地址") {
			return "Local", "Local"
		}
		if strings.Contains(location, "局域网IP") || strings.Contains(location, "Private-Use") {
			return "Local", "Local"
		}
		return "IANA", "Unknown"
	}

	// 处理包含"–"分隔符的格式，如："中国–浙江–丽水 移动"
	if strings.Contains(location, "–") {
		parts := strings.Split(location, "–")
		if len(parts) >= 2 {
			country = strings.TrimSpace(parts[0])
			city = strings.TrimSpace(parts[1])

			// 如果省份后面还有城市信息，只取省份
			if strings.Contains(city, " ") {
				cityParts := strings.Fields(city)
				city = cityParts[0] // 只取省份部分
			}

			return country, city
		}
	}

	// 处理包含"-"分隔符的格式
	if strings.Contains(location, "-") {
		parts := strings.Split(location, "-")
		if len(parts) >= 2 {
			country = strings.TrimSpace(parts[0])
			city = strings.TrimSpace(parts[1])

			// 如果省份后面还有城市信息，只取省份
			if strings.Contains(city, " ") {
				cityParts := strings.Fields(city)
				city = cityParts[0] // 只取省份部分
			}

			return country, city
		}
	}

	// 处理空格分隔的格式，如："美国 加利福尼亚州 圣克拉拉"
	parts := strings.Fields(location)
	if len(parts) >= 2 {
		country = parts[0]
		city = parts[1]

		// 如果第二个部分包含特殊字符（如DNS服务器），只取国家
		if strings.Contains(city, "/") || strings.Contains(city, "DNS") || strings.Contains(city, "服务器") {
			return country, "Unknown"
		}

		return country, city
	} else if len(parts) == 1 {
		return parts[0], "Unknown"
	}

	return "Unknown", "Unknown"
}
