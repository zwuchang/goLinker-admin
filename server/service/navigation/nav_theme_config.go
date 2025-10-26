package navigation

import (
	"context"
	"encoding/json"
	"errors"
	"goLinker-admin/server/global"
	"goLinker-admin/server/model/navigation"
	navRequest "goLinker-admin/server/model/navigation/request"

	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type NavThemeConfigService struct{}

// CreateThemeConfig 创建主题配置
func (s *NavThemeConfigService) CreateThemeConfig(req navRequest.CreateNavThemeConfigRequest) (err error) {
	if global.GVA_DB == nil {
		global.GVA_LOG.Error("数据库连接为空")
		return errors.New("数据库连接为空")
	}

	// 参数验证
	if req.Name == "" {
		return errors.New("主题名称不能为空")
	}
	if req.ConfigJson == "" {
		return errors.New("主题配置不能为空")
	}

	// 验证JSON格式
	var configMap map[string]interface{}
	if err := json.Unmarshal([]byte(req.ConfigJson), &configMap); err != nil {
		return errors.New("主题配置JSON格式错误")
	}

	// 检查主题名称是否已存在
	var existingConfig navigation.NavThemeConfig
	err = global.GVA_DB.Where("name = ?", req.Name).First(&existingConfig).Error
	if err == nil {
		return errors.New("主题名称已存在")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		global.GVA_LOG.Error("查询主题配置失败", zap.Error(err))
		return err
	}

	// 如果设置为默认主题，先取消其他默认主题
	if req.IsDefault == 1 {
		err = global.GVA_DB.Model(&navigation.NavThemeConfig{}).Where("is_default = ?", 1).Update("is_default", 0).Error
		if err != nil {
			global.GVA_LOG.Error("取消默认主题失败", zap.Error(err))
			return err
		}
	}

	// 创建主题配置
	themeConfig := navigation.NavThemeConfig{
		Name:        req.Name,
		Description: req.Description,
		ConfigJson:  req.ConfigJson,
		IsDefault:   req.IsDefault,
	}

	err = global.GVA_DB.Create(&themeConfig).Error
	if err != nil {
		global.GVA_LOG.Error("创建主题配置失败", zap.Error(err))
		return err
	}

	global.GVA_LOG.Info("创建主题配置成功", zap.String("name", req.Name))

	// 清除缓存
	s.ClearThemeConfigCache()

	return nil
}

// UpdateThemeConfig 更新主题配置
func (s *NavThemeConfigService) UpdateThemeConfig(req navRequest.UpdateNavThemeConfigRequest) (err error) {
	if global.GVA_DB == nil {
		global.GVA_LOG.Error("数据库连接为空")
		return errors.New("数据库连接为空")
	}

	// 参数验证
	if req.ID == 0 {
		return errors.New("主题配置ID不能为空")
	}
	if req.Name == "" {
		return errors.New("主题名称不能为空")
	}
	if req.ConfigJson == "" {
		return errors.New("主题配置不能为空")
	}

	// 验证JSON格式
	var configMap map[string]interface{}
	if err := json.Unmarshal([]byte(req.ConfigJson), &configMap); err != nil {
		return errors.New("主题配置JSON格式错误")
	}

	// 检查主题配置是否存在
	var existingConfig navigation.NavThemeConfig
	err = global.GVA_DB.Where("id = ?", req.ID).First(&existingConfig).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("主题配置不存在")
	}
	if err != nil {
		global.GVA_LOG.Error("查询主题配置失败", zap.Error(err))
		return err
	}

	// 检查主题名称是否与其他配置重复
	var duplicateConfig navigation.NavThemeConfig
	err = global.GVA_DB.Where("name = ? AND id != ?", req.Name, req.ID).First(&duplicateConfig).Error
	if err == nil {
		return errors.New("主题名称已存在")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		global.GVA_LOG.Error("查询主题配置失败", zap.Error(err))
		return err
	}

	// 如果设置为默认主题，先取消其他默认主题
	if req.IsDefault == 1 {
		err = global.GVA_DB.Model(&navigation.NavThemeConfig{}).Where("is_default = ? AND id != ?", 1, req.ID).Update("is_default", 0).Error
		if err != nil {
			global.GVA_LOG.Error("取消默认主题失败", zap.Error(err))
			return err
		}
	}

	// 更新主题配置
	updateData := map[string]interface{}{
		"name":        req.Name,
		"description": req.Description,
		"config_json": req.ConfigJson,
		"is_default":  req.IsDefault,
	}

	err = global.GVA_DB.Model(&existingConfig).Updates(updateData).Error
	if err != nil {
		global.GVA_LOG.Error("更新主题配置失败", zap.Error(err))
		return err
	}

	global.GVA_LOG.Info("更新主题配置成功", zap.Uint("id", req.ID), zap.String("name", req.Name))

	// 清除缓存
	s.ClearThemeConfigCache()

	return nil
}

// DeleteThemeConfig 删除主题配置
func (s *NavThemeConfigService) DeleteThemeConfig(id uint) (err error) {
	if global.GVA_DB == nil {
		global.GVA_LOG.Error("数据库连接为空")
		return errors.New("数据库连接为空")
	}

	// 检查主题配置是否存在
	var existingConfig navigation.NavThemeConfig
	err = global.GVA_DB.Where("id = ?", id).First(&existingConfig).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("主题配置不存在")
	}
	if err != nil {
		global.GVA_LOG.Error("查询主题配置失败", zap.Error(err))
		return err
	}

	// 检查是否为默认主题
	if existingConfig.IsDefault == 1 {
		return errors.New("不能删除默认主题")
	}

	// 删除主题配置
	err = global.GVA_DB.Delete(&existingConfig).Error
	if err != nil {
		global.GVA_LOG.Error("删除主题配置失败", zap.Error(err))
		return err
	}

	global.GVA_LOG.Info("删除主题配置成功", zap.Uint("id", id))

	// 清除缓存
	s.ClearThemeConfigCache()

	return nil
}

// GetThemeConfigList 获取主题配置列表
func (s *NavThemeConfigService) GetThemeConfigList(req navRequest.NavThemeConfigSearch) (list []navigation.NavThemeConfig, total int64, err error) {
	if global.GVA_DB == nil {
		global.GVA_LOG.Error("数据库连接为空")
		return nil, 0, errors.New("数据库连接为空")
	}

	db := global.GVA_DB.Model(&navigation.NavThemeConfig{})

	// 搜索条件
	if req.Name != "" {
		db = db.Where("name LIKE ?", "%"+req.Name+"%")
	}
	if req.Description != "" {
		db = db.Where("description LIKE ?", "%"+req.Description+"%")
	}

	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 || req.PageSize > 100 {
		req.PageSize = 10
	}

	// 获取总数
	err = db.Count(&total).Error
	if err != nil {
		global.GVA_LOG.Error("获取主题配置总数失败", zap.Error(err))
		return nil, 0, err
	}

	// 分页查询
	offset := (req.Page - 1) * req.PageSize
	err = db.Limit(req.PageSize).Offset(offset).Order("is_default DESC, created_at DESC").Find(&list).Error
	if err != nil {
		global.GVA_LOG.Error("获取主题配置列表失败", zap.Error(err))
		return nil, 0, err
	}

	return list, total, nil
}

// GetThemeConfigById 根据ID获取主题配置
func (s *NavThemeConfigService) GetThemeConfigById(id uint) (config navigation.NavThemeConfig, err error) {
	if global.GVA_DB == nil {
		global.GVA_LOG.Error("数据库连接为空")
		return config, errors.New("数据库连接为空")
	}

	err = global.GVA_DB.Where("id = ?", id).First(&config).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return config, errors.New("主题配置不存在")
	}
	if err != nil {
		global.GVA_LOG.Error("查询主题配置失败", zap.Error(err))
		return config, err
	}

	return config, nil
}

// SetDefaultTheme 设置默认主题
func (s *NavThemeConfigService) SetDefaultTheme(req navRequest.SetDefaultThemeRequest) (err error) {
	if global.GVA_DB == nil {
		global.GVA_LOG.Error("数据库连接为空")
		return errors.New("数据库连接为空")
	}

	// 检查主题配置是否存在
	var existingConfig navigation.NavThemeConfig
	err = global.GVA_DB.Where("id = ?", req.ID).First(&existingConfig).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("主题配置不存在")
	}
	if err != nil {
		global.GVA_LOG.Error("查询主题配置失败", zap.Error(err))
		return err
	}

	// 开启事务
	tx := global.GVA_DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 取消所有默认主题
	err = tx.Model(&navigation.NavThemeConfig{}).Where("is_default = ?", 1).Update("is_default", 0).Error
	if err != nil {
		tx.Rollback()
		global.GVA_LOG.Error("取消默认主题失败", zap.Error(err))
		return err
	}

	// 设置新的默认主题
	err = tx.Model(&existingConfig).Update("is_default", 1).Error
	if err != nil {
		tx.Rollback()
		global.GVA_LOG.Error("设置默认主题失败", zap.Error(err))
		return err
	}

	// 提交事务
	err = tx.Commit().Error
	if err != nil {
		global.GVA_LOG.Error("提交事务失败", zap.Error(err))
		return err
	}

	global.GVA_LOG.Info("设置默认主题成功", zap.Uint("id", req.ID))

	// 清除缓存
	s.ClearThemeConfigCache()

	return nil
}

// GetDefaultTheme 获取默认主题
func (s *NavThemeConfigService) GetDefaultTheme() (config navigation.NavThemeConfig, err error) {
	if global.GVA_DB == nil {
		global.GVA_LOG.Error("数据库连接为空")
		return config, errors.New("数据库连接为空")
	}

	err = global.GVA_DB.Where("is_default = ?", 1).First(&config).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return config, errors.New("未找到默认主题")
	}
	if err != nil {
		global.GVA_LOG.Error("查询默认主题失败", zap.Error(err))
		return config, err
	}

	return config, nil
}

// GetAllThemes 获取所有主题（用于前端主题切换）
func (s *NavThemeConfigService) GetAllThemes() (list []navigation.NavThemeConfig, err error) {
	if global.GVA_DB == nil {
		global.GVA_LOG.Error("数据库连接为空")
		return nil, errors.New("数据库连接为空")
	}

	err = global.GVA_DB.Order("is_default DESC, created_at DESC").Find(&list).Error
	if err != nil {
		global.GVA_LOG.Error("获取所有主题失败", zap.Error(err))
		return nil, err
	}

	return list, nil
}

// GetActiveThemeConfig 获取活跃的主题配置（带缓存）
func (s *NavThemeConfigService) GetActiveThemeConfig() (config navigation.NavThemeConfig, err error) {
	cacheKey := "nav:theme:active"

	// 如果启用了Redis，尝试从缓存获取
	if global.GVA_REDIS != nil {
		cacheData, err := global.GVA_REDIS.Get(context.Background(), cacheKey).Result()
		if err == nil && cacheData != "" {
			var cachedConfig navigation.NavThemeConfig
			if err := json.Unmarshal([]byte(cacheData), &cachedConfig); err == nil {
				global.GVA_LOG.Info("从缓存获取活跃主题配置成功")
				return cachedConfig, nil
			}
		}
	}

	// 从数据库获取活跃主题（优先默认主题）
	err = global.GVA_DB.Where("is_default = ?", 1).First(&config).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 如果没有默认主题，获取最新的主题
		err = global.GVA_DB.Order("created_at DESC").First(&config).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return config, errors.New("未找到可用主题配置")
		}
		if err != nil {
			global.GVA_LOG.Error("查询活跃主题配置失败", zap.Error(err))
			return config, err
		}
	} else if err != nil {
		global.GVA_LOG.Error("查询活跃主题配置失败", zap.Error(err))
		return config, err
	}

	// 如果启用了Redis，将结果缓存（缓存24小时）
	if global.GVA_REDIS != nil {
		if cacheData, err := json.Marshal(config); err == nil {
			if err := global.GVA_REDIS.Set(context.Background(), cacheKey, cacheData, 24*time.Hour).Err(); err != nil {
				global.GVA_LOG.Warn("缓存活跃主题配置失败", zap.Error(err))
			}
		}
	}

	return config, nil
}

// ClearThemeConfigCache 清除主题配置缓存
func (s *NavThemeConfigService) ClearThemeConfigCache() {
	cacheKey := "nav:theme:active"
	if global.GVA_REDIS != nil {
		if err := global.GVA_REDIS.Del(context.Background(), cacheKey).Err(); err != nil {
			global.GVA_LOG.Warn("清除主题配置缓存失败", zap.Error(err))
		} else {
			global.GVA_LOG.Info("清除主题配置缓存成功")
		}
	}
}
