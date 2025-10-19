package navigation

import (
	"errors"
	"goLinker-admin/server/global"
	"goLinker-admin/server/model/navigation"
	navRequest "goLinker-admin/server/model/navigation/request"

	"go.uber.org/zap"
)

type NavPlatformConfigService struct{}

// GetPlatformConfigList 获取平台游戏配置列表
func (s *NavPlatformConfigService) GetPlatformConfigList(info navRequest.NavPlatformConfigSearch) (list []navigation.NavPlatformConfig, total int64, err error) {
	if global.GVA_DB == nil {
		global.GVA_LOG.Error("数据库连接为空")
		return list, 0, errors.New("数据库连接为空")
	}

	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&navigation.NavPlatformConfig{})

	// 构建查询条件
	if info.PlatformName != "" {
		db = db.Where("platform_name LIKE ?", "%"+info.PlatformName+"%")
	}
	if info.Status != nil {
		db = db.Where("status = ?", *info.Status)
	}
	if info.IsVisible != nil {
		db = db.Where("is_visible = ?", *info.IsVisible)
	}

	// 获取总数
	err = db.Count(&total).Error
	if err != nil {
		global.GVA_LOG.Error("获取平台游戏配置总数失败", zap.Error(err))
		return list, 0, err
	}

	// 按排序字段和创建时间排序
	err = db.Limit(limit).Offset(offset).Order("sort ASC, created_at DESC").Find(&list).Error
	if err != nil {
		global.GVA_LOG.Error("获取平台游戏配置列表失败", zap.Error(err))
		return list, 0, err
	}

	return list, total, nil
}

// CreatePlatformConfig 创建平台游戏配置
func (s *NavPlatformConfigService) CreatePlatformConfig(info navRequest.CreateNavPlatformConfigRequest) (err error) {
	if global.GVA_DB == nil {
		global.GVA_LOG.Error("数据库连接为空")
		return errors.New("数据库连接为空")
	}

	platformConfig := navigation.NavPlatformConfig{
		PlatformName: info.PlatformName,
		PlatformIcon: info.PlatformIcon,
		PlatformApi:  info.PlatformApi,
		Sort:         info.Sort,
		Status:       info.Status,
		Description:  info.Description,
		IsVisible:    info.IsVisible,
	}

	err = global.GVA_DB.Create(&platformConfig).Error
	if err != nil {
		global.GVA_LOG.Error("创建平台游戏配置失败", zap.Error(err))
		return err
	}

	return nil
}

// UpdatePlatformConfig 更新平台游戏配置
func (s *NavPlatformConfigService) UpdatePlatformConfig(info navRequest.UpdateNavPlatformConfigRequest) (err error) {
	if global.GVA_DB == nil {
		global.GVA_LOG.Error("数据库连接为空")
		return errors.New("数据库连接为空")
	}

	var platformConfig navigation.NavPlatformConfig
	err = global.GVA_DB.Where("id = ?", info.ID).First(&platformConfig).Error
	if err != nil {
		global.GVA_LOG.Error("平台游戏配置不存在", zap.Error(err))
		return err
	}

	// 更新字段
	platformConfig.PlatformName = info.PlatformName
	platformConfig.PlatformIcon = info.PlatformIcon
	platformConfig.PlatformApi = info.PlatformApi
	platformConfig.Sort = info.Sort
	platformConfig.Status = info.Status
	platformConfig.Description = info.Description
	platformConfig.IsVisible = info.IsVisible

	err = global.GVA_DB.Save(&platformConfig).Error
	if err != nil {
		global.GVA_LOG.Error("更新平台游戏配置失败", zap.Error(err))
		return err
	}

	return nil
}

// GetPlatformConfigById 根据ID获取平台游戏配置
func (s *NavPlatformConfigService) GetPlatformConfigById(id uint) (platformConfig navigation.NavPlatformConfig, err error) {
	if global.GVA_DB == nil {
		global.GVA_LOG.Error("数据库连接为空")
		return platformConfig, errors.New("数据库连接为空")
	}

	err = global.GVA_DB.Where("id = ?", id).First(&platformConfig).Error
	if err != nil {
		global.GVA_LOG.Error("获取平台游戏配置失败", zap.Error(err))
		return platformConfig, err
	}

	return platformConfig, nil
}

// DeletePlatformConfig 删除平台游戏配置
func (s *NavPlatformConfigService) DeletePlatformConfig(id uint) (err error) {
	if global.GVA_DB == nil {
		global.GVA_LOG.Error("数据库连接为空")
		return errors.New("数据库连接为空")
	}

	err = global.GVA_DB.Where("id = ?", id).Delete(&navigation.NavPlatformConfig{}).Error
	if err != nil {
		global.GVA_LOG.Error("删除平台游戏配置失败", zap.Error(err))
		return err
	}

	return nil
}

// GetVisiblePlatformConfigs 获取可见的平台游戏配置列表（用于公开接口）
func (s *NavPlatformConfigService) GetVisiblePlatformConfigs() (list []navigation.NavPlatformConfig, err error) {
	if global.GVA_DB == nil {
		global.GVA_LOG.Error("数据库连接为空")
		return list, errors.New("数据库连接为空")
	}

	err = global.GVA_DB.Where("status = ? AND is_visible = ?", 1, 1).
		Order("sort ASC, created_at DESC").
		Find(&list).Error
	if err != nil {
		global.GVA_LOG.Error("获取可见平台游戏配置列表失败", zap.Error(err))
		return list, err
	}

	return list, nil
}
