package navigation

import (
	"context"
	"encoding/json"
	"fmt"
	"goLinker-admin/server/global"
	"goLinker-admin/server/model/navigation"
	navRequest "goLinker-admin/server/model/navigation/request"
	navResponse "goLinker-admin/server/model/navigation/response"
	"time"
)

type NavPWAConfigService struct{}

const (
	PWA_CONFIG_CACHE_KEY = "pwa:config:active"
	PWA_CACHE_EXPIRE     = 24 * time.Hour
)

// GetPWAConfigList 获取PWA配置列表
func (s *NavPWAConfigService) GetPWAConfigList(info navRequest.NavPWAConfigSearch) (list []navigation.NavPWAConfig, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&navigation.NavPWAConfig{})

	// 添加搜索条件
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.Status != 0 {
		db = db.Where("status = ?", info.Status)
	}

	var pwaConfigs []navigation.NavPWAConfig
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Order("created_at DESC").Limit(limit).Offset(offset).Find(&pwaConfigs).Error
	return pwaConfigs, total, err
}

// CreatePWAConfig 创建PWA配置（只允许一条记录）
func (s *NavPWAConfigService) CreatePWAConfig(req navRequest.CreateNavPWAConfigRequest) (err error) {
	// 检查是否已存在配置
	var count int64
	err = global.GVA_DB.Model(&navigation.NavPWAConfig{}).Count(&count).Error
	if err != nil {
		return err
	}

	if count > 0 {
		return fmt.Errorf("PWA配置已存在，只能创建一条配置")
	}

	// 将图标数组转换为JSON字符串
	iconsJSON, err := json.Marshal(req.Icons)
	if err != nil {
		return err
	}

	// 创建PWA配置
	pwaConfig := navigation.NavPWAConfig{
		BackgroundColor: req.BackgroundColor,
		ThemeColor:      req.ThemeColor,
		Display:         req.Display,
		Name:            req.Name,
		ShortName:       req.ShortName,
		Orientation:     req.Orientation,
		Scope:           req.Scope,
		StartUrl:        req.StartUrl,
		Status:          req.Status,
		Icons:           string(iconsJSON),
	}

	err = global.GVA_DB.Create(&pwaConfig).Error
	if err != nil {
		return err
	}

	// 清除缓存
	s.clearPWACache()
	return nil
}

// UpdatePWAConfig 更新PWA配置
func (s *NavPWAConfigService) UpdatePWAConfig(req navRequest.UpdateNavPWAConfigRequest) (err error) {
	var pwaConfig navigation.NavPWAConfig
	err = global.GVA_DB.First(&pwaConfig, req.ID).Error
	if err != nil {
		return err
	}

	// 更新配置字段
	updateFields := make(map[string]interface{})
	if req.BackgroundColor != nil {
		updateFields["background_color"] = *req.BackgroundColor
	}
	if req.ThemeColor != nil {
		updateFields["theme_color"] = *req.ThemeColor
	}
	if req.Display != nil {
		updateFields["display"] = *req.Display
	}
	if req.Name != nil {
		updateFields["name"] = *req.Name
	}
	if req.ShortName != nil {
		updateFields["short_name"] = *req.ShortName
	}
	if req.Orientation != nil {
		updateFields["orientation"] = *req.Orientation
	}
	if req.Scope != nil {
		updateFields["scope"] = *req.Scope
	}
	if req.StartUrl != nil {
		updateFields["start_url"] = *req.StartUrl
	}
	if req.Status != nil {
		updateFields["status"] = *req.Status
	}

	// 更新图标
	if len(req.Icons) > 0 {
		iconsJSON, err := json.Marshal(req.Icons)
		if err != nil {
			return err
		}
		updateFields["icons"] = string(iconsJSON)
	}

	if len(updateFields) > 0 {
		err = global.GVA_DB.Model(&pwaConfig).Updates(updateFields).Error
		if err != nil {
			return err
		}
	}

	// 清除缓存
	s.clearPWACache()
	return nil
}

// DeletePWAConfig 删除PWA配置
func (s *NavPWAConfigService) DeletePWAConfig(id uint) (err error) {
	err = global.GVA_DB.Delete(&navigation.NavPWAConfig{}, id).Error
	if err != nil {
		return err
	}

	// 清除缓存
	s.clearPWACache()
	return nil
}

// GetPWAConfigById 根据ID获取PWA配置
func (s *NavPWAConfigService) GetPWAConfigById(id uint) (pwaConfig navigation.NavPWAConfig, err error) {
	err = global.GVA_DB.First(&pwaConfig, id).Error
	return pwaConfig, err
}

// GetActivePWAConfig 获取激活的PWA配置（带缓存）
func (s *NavPWAConfigService) GetActivePWAConfig() (config navResponse.PWAConfigItemResponse, err error) {
	// 如果Redis未启用，直接从数据库获取
	if global.GVA_REDIS == nil {
		return s.getPWAConfigFromDB()
	}

	// Redis已启用，先尝试从缓存获取
	cachedConfig, err := s.getPWACache()
	if err == nil && cachedConfig != nil {
		return *cachedConfig, nil
	}

	// 缓存未命中，从数据库获取
	dbConfig, err := s.getPWAConfigFromDB()
	if err != nil {
		return config, err
	}

	// 写入缓存
	s.setPWACache(&dbConfig)
	return dbConfig, nil
}

// getPWAConfigFromDB 从数据库获取PWA配置
func (s *NavPWAConfigService) getPWAConfigFromDB() (config navResponse.PWAConfigItemResponse, err error) {
	var pwaConfig navigation.NavPWAConfig
	err = global.GVA_DB.Where("status = ?", 1).First(&pwaConfig).Error
	if err != nil {
		// 如果数据库中没有配置，返回默认配置
		return s.getDefaultPWAConfig(), nil
	}

	// 解析图标JSON
	var icons []navigation.PWAIconItem
	if pwaConfig.Icons != "" {
		err = json.Unmarshal([]byte(pwaConfig.Icons), &icons)
		if err != nil {
			return config, err
		}
	}

	// 转换为响应格式
	var pwaIcons []navResponse.PWAIconItem
	for _, icon := range icons {
		pwaIcons = append(pwaIcons, navResponse.PWAIconItem{
			Sizes: icon.Sizes,
			Src:   icon.Src,
			Type:  icon.Type,
		})
	}

	config = navResponse.PWAConfigItemResponse{
		BackgroundColor: pwaConfig.BackgroundColor,
		ThemeColor:      pwaConfig.ThemeColor,
		Display:         pwaConfig.Display,
		Icons:           pwaIcons,
		Name:            pwaConfig.Name,
		Orientation:     pwaConfig.Orientation,
		Scope:           pwaConfig.Scope,
		ShortName:       pwaConfig.ShortName,
		StartUrl:        pwaConfig.StartUrl,
	}
	return config, nil
}

// getDefaultPWAConfig 获取默认PWA配置
func (s *NavPWAConfigService) getDefaultPWAConfig() navResponse.PWAConfigItemResponse {
	return navResponse.PWAConfigItemResponse{
		BackgroundColor: "#164533",
		ThemeColor:      "#164533",
		Display:         "standalone",
		Icons: []navResponse.PWAIconItem{
			{Sizes: "192x192", Src: "https://dydybet9.com/static/logow.png", Type: "image/png"},
			{Sizes: "512x512", Src: "https://dydybet9.com/static/logow.png", Type: "image/png"},
		},
		Name:        "DYDYBET",
		Orientation: "portrait",
		Scope:       "http://localhost:5173/",
		ShortName:   "DYDYBET",
		StartUrl:    "http://localhost:5173/",
	}
}

// getPWACache 获取PWA配置缓存
func (s *NavPWAConfigService) getPWACache() (*navResponse.PWAConfigItemResponse, error) {
	if global.GVA_REDIS == nil {
		return nil, fmt.Errorf("redis not initialized")
	}

	cacheData, err := global.GVA_REDIS.Get(context.Background(), PWA_CONFIG_CACHE_KEY).Result()
	if err != nil {
		return nil, err
	}

	var config navResponse.PWAConfigItemResponse
	err = json.Unmarshal([]byte(cacheData), &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

// setPWACache 设置PWA配置缓存
func (s *NavPWAConfigService) setPWACache(config *navResponse.PWAConfigItemResponse) error {
	if global.GVA_REDIS == nil {
		return fmt.Errorf("redis not initialized")
	}

	configData, err := json.Marshal(config)
	if err != nil {
		return err
	}

	return global.GVA_REDIS.Set(context.Background(), PWA_CONFIG_CACHE_KEY, configData, PWA_CACHE_EXPIRE).Err()
}

// clearPWACache 清除PWA配置缓存
func (s *NavPWAConfigService) clearPWACache() error {
	if global.GVA_REDIS == nil {
		return fmt.Errorf("redis not initialized")
	}

	return global.GVA_REDIS.Del(context.Background(), PWA_CONFIG_CACHE_KEY).Err()
}

// ClearPWACache 清除PWA配置缓存（公开方法）
func (s *NavPWAConfigService) ClearPWACache() error {
	if global.GVA_REDIS == nil {
		return fmt.Errorf("缓存未开启，Redis服务未初始化")
	}

	return global.GVA_REDIS.Del(context.Background(), PWA_CONFIG_CACHE_KEY).Err()
}
