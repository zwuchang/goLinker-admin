package navigation

import (
	"goLinker-admin/server/global"
	"goLinker-admin/server/model/navigation"
	"goLinker-admin/server/model/navigation/request"
	"goLinker-admin/server/model/navigation/response"
	"time"
)

type NavAccessStatsService struct{}

// CreateAccessStats 创建访问统计记录
func (s *NavAccessStatsService) CreateAccessStats(stats navigation.NavAccessStats) error {
	return global.GVA_DB.Create(&stats).Error
}

// GetAccessStatsList 获取访问统计列表
func (s *NavAccessStatsService) GetAccessStatsList(info request.NavAccessStatsSearch) (list []navigation.NavAccessStats, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&navigation.NavAccessStats{})

	// 搜索条件
	if info.ApiPath != "" {
		db = db.Where("api_path LIKE ?", "%"+info.ApiPath+"%")
	}
	if info.Method != "" {
		db = db.Where("method = ?", info.Method)
	}
	if info.ClientIP != "" {
		db = db.Where("client_ip LIKE ?", "%"+info.ClientIP+"%")
	}
	if info.StatusCode != nil {
		db = db.Where("status_code = ?", *info.StatusCode)
	}
	if info.Country != "" {
		db = db.Where("country LIKE ?", "%"+info.Country+"%")
	}
	if info.Device != "" {
		db = db.Where("device LIKE ?", "%"+info.Device+"%")
	}
	if info.Browser != "" {
		db = db.Where("browser LIKE ?", "%"+info.Browser+"%")
	}
	if info.OS != "" {
		db = db.Where("os LIKE ?", "%"+info.OS+"%")
	}
	if info.StartTime != "" {
		db = db.Where("created_at >= ?", info.StartTime)
	}
	if info.EndTime != "" {
		db = db.Where("created_at <= ?", info.EndTime)
	}

	// 排序
	orderBy := "created_at DESC"
	if info.OrderBy != "" {
		orderBy = info.OrderBy
		if info.OrderType != "" {
			orderBy = orderBy + " " + info.OrderType
		}
	}
	db = db.Order(orderBy)

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&list).Error
	return list, total, err
}

// GetAccessStatsSummary 获取访问统计汇总
func (s *NavAccessStatsService) GetAccessStatsSummary(startTime, endTime string) (summary response.NavAccessStatsSummaryResponse, err error) {
	db := global.GVA_DB.Model(&navigation.NavAccessStats{})

	// 时间范围过滤
	if startTime != "" {
		db = db.Where("created_at >= ?", startTime)
	}
	if endTime != "" {
		db = db.Where("created_at <= ?", endTime)
	}

	// 总请求数
	err = db.Count(&summary.TotalRequests).Error
	if err != nil {
		return
	}

	// 独立IP数
	err = db.Distinct("client_ip").Count(&summary.UniqueIPs).Error
	if err != nil {
		return
	}

	// 平均响应时间
	var avgResponseTime float64
	err = db.Select("AVG(response_time)").Scan(&avgResponseTime).Error
	if err != nil {
		return
	}
	summary.AvgResponseTime = avgResponseTime

	// 成功率
	var successCount int64
	err = db.Where("status_code >= 200 AND status_code < 300").Count(&successCount).Error
	if err != nil {
		return
	}
	if summary.TotalRequests > 0 {
		summary.SuccessRate = float64(successCount) / float64(summary.TotalRequests) * 100
	}

	// 热门API路径
	var topApiPaths []response.ApiPathCount
	err = db.Select("api_path, COUNT(*) as count").
		Group("api_path").
		Order("count DESC").
		Limit(10).
		Scan(&topApiPaths).Error
	if err != nil {
		return
	}
	summary.TopApiPaths = topApiPaths

	// 热门国家
	var topCountries []response.CountryCount
	err = db.Select("country, COUNT(*) as count").
		Where("country != ''").
		Group("country").
		Order("count DESC").
		Limit(10).
		Scan(&topCountries).Error
	if err != nil {
		return
	}
	summary.TopCountries = topCountries

	// 热门设备
	var topDevices []response.DeviceCount
	err = db.Select("device, COUNT(*) as count").
		Where("device != ''").
		Group("device").
		Order("count DESC").
		Limit(10).
		Scan(&topDevices).Error
	if err != nil {
		return
	}
	summary.TopDevices = topDevices

	// 热门浏览器
	var topBrowsers []response.BrowserCount
	err = db.Select("browser, COUNT(*) as count").
		Where("browser != ''").
		Group("browser").
		Order("count DESC").
		Limit(10).
		Scan(&topBrowsers).Error
	if err != nil {
		return
	}
	summary.TopBrowsers = topBrowsers

	return summary, nil
}

// CleanOldStats 清理旧数据
func (s *NavAccessStatsService) CleanOldStats(days int) error {
	cutoffTime := time.Now().AddDate(0, 0, -days)
	return global.GVA_DB.Where("created_at < ?", cutoffTime).Delete(&navigation.NavAccessStats{}).Error
}
