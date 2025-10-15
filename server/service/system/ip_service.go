package system

import (
	"goLinker-admin/server/global"
	"regexp"
	"strings"

	"go.uber.org/zap"
)

type IpService struct{}

var IpServiceApp = new(IpService)

func (s *IpService) GetLocation(ipAddr string) string {
	if ipAddr == "" {
		return ""
	}

	// 添加panic恢复机制
	defer func() {
		if r := recover(); r != nil {
			global.GVA_LOG.Error("IP查询发生panic",
				zap.String("ip", ipAddr),
				zap.Any("panic", r))
		}
	}()

	// 直接使用全局MEMORY模式的搜索器（线程安全）
	if global.IpSearcher == nil {
		global.GVA_LOG.Debug("IP地址库未初始化，跳过查询", zap.String("ip", ipAddr))
		return ""
	}

	result, err := global.IpSearcher.Search(ipAddr)
	if err != nil {
		global.GVA_LOG.Debug("查询IP地址失败", zap.String("ip", ipAddr), zap.Error(err))
		return ""
	}

	// 规范化结果格式
	result = strings.ReplaceAll(result, "\t", " ")
	re := regexp.MustCompile(`\s+`)
	result = re.ReplaceAllString(result, " ")

	return strings.TrimSpace(result)
}
