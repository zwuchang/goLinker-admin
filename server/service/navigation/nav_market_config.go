package navigation

import (
	"errors"
	"goLinker-admin/server/global"
	"goLinker-admin/server/model/navigation"
	navRequest "goLinker-admin/server/model/navigation/request"
)

type NavMarketConfigService struct{}

// CreateMarketConfig 创建市场配置
func (s *NavMarketConfigService) CreateMarketConfig(req navRequest.CreateMarketConfig) (err error) {
	var marketConfig navigation.NavMarketConfig
	marketConfig.Name = req.Name
	marketConfig.Logo = req.Logo
	marketConfig.JumpUrl = req.JumpUrl
	marketConfig.Type = req.Type
	marketConfig.Status = req.Status
	marketConfig.IsVisible = req.IsVisible
	marketConfig.Sort = req.Sort
	marketConfig.Description = req.Description
	err = global.GVA_DB.Create(&marketConfig).Error
	return err
}

// DeleteMarketConfig 删除市场配置
func (s *NavMarketConfigService) DeleteMarketConfig(id uint) (err error) {
	err = global.GVA_DB.Delete(&navigation.NavMarketConfig{}, "id = ?", id).Error
	return err
}

// UpdateMarketConfig 更新市场配置
func (s *NavMarketConfigService) UpdateMarketConfig(req navRequest.UpdateMarketConfig) (err error) {
	var marketConfig navigation.NavMarketConfig
	err = global.GVA_DB.Where("id = ?", req.ID).First(&marketConfig).Error
	if err != nil {
		return err
	}
	marketConfig.Name = req.Name
	marketConfig.Logo = req.Logo
	marketConfig.JumpUrl = req.JumpUrl
	marketConfig.Type = req.Type
	marketConfig.Status = req.Status
	marketConfig.IsVisible = req.IsVisible
	marketConfig.Sort = req.Sort
	marketConfig.Description = req.Description
	err = global.GVA_DB.Save(&marketConfig).Error
	return err
}

// GetMarketConfigById 根据ID获取市场配置
func (s *NavMarketConfigService) GetMarketConfigById(id uint) (marketConfig navigation.NavMarketConfig, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&marketConfig).Error
	return
}

// GetMarketConfigList 获取市场配置列表
func (s *NavMarketConfigService) GetMarketConfigList(info navRequest.NavMarketConfigSearch) (list []navigation.NavMarketConfig, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&navigation.NavMarketConfig{})
	var marketConfigs []navigation.NavMarketConfig

	// 搜索条件
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.Type != 0 {
		db = db.Where("type = ?", info.Type)
	}
	if info.Status != 0 {
		db = db.Where("status = ?", info.Status)
	}
	if info.IsVisible != 0 {
		db = db.Where("is_visible = ?", info.IsVisible)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Order("sort ASC, created_at DESC").Find(&marketConfigs).Error
	return marketConfigs, total, err
}

// GetMarketConfigsByType 根据类型获取可见的市场配置
func (s *NavMarketConfigService) GetMarketConfigsByType(marketType int, limit int) (list []navigation.NavMarketConfig, err error) {
	if global.GVA_DB == nil {
		global.GVA_LOG.Error("数据库连接为空")
		return list, errors.New("数据库连接为空")
	}

	db := global.GVA_DB.Model(&navigation.NavMarketConfig{})
	err = db.Where("type = ? AND status = ? AND is_visible = ?", marketType, 1, 1).Order("sort ASC").Limit(limit).Find(&list).Error
	return list, err
}
