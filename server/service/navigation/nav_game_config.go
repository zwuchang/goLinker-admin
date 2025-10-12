package navigation

import (
	"errors"
	"goLinker-admin/server/global"
	"goLinker-admin/server/model/navigation"
	navRequest "goLinker-admin/server/model/navigation/request"
)

type NavGameConfigService struct{}

// CreateGameConfig 创建游戏配置
func (navGameConfigService *NavGameConfigService) CreateGameConfig(gameConfig navRequest.CreateNavGameConfigRequest) (err error) {
	if global.GVA_DB == nil {
		global.GVA_LOG.Error("数据库连接为空")
		return errors.New("数据库连接异常")
	}

	// 检查是否已存在配置
	var count int64
	err = global.GVA_DB.Model(&navigation.NavGameConfig{}).Count(&count).Error
	if err != nil {
		global.GVA_LOG.Error("查询游戏配置失败: " + err.Error())
		return err
	}

	if count > 0 {
		return errors.New("游戏配置已存在，只能有一个配置")
	}

	// 创建游戏配置
	newGameConfig := navigation.NavGameConfig{
		DownloadUrl:  gameConfig.DownloadUrl,
		AudioUrl:     gameConfig.AudioUrl,
		WebsiteTitle: gameConfig.WebsiteTitle,
		WebsiteDesc:  gameConfig.WebsiteDesc,
		WebsiteIcon:  gameConfig.WebsiteIcon,
		WebsiteLogo:  gameConfig.WebsiteLogo,
		Status:       gameConfig.Status,
	}

	if newGameConfig.Status == 0 {
		newGameConfig.Status = 1 // 默认启用
	}

	err = global.GVA_DB.Create(&newGameConfig).Error
	if err != nil {
		global.GVA_LOG.Error("创建游戏配置失败: " + err.Error())
		return err
	}

	return nil
}

// UpdateGameConfig 更新游戏配置
func (navGameConfigService *NavGameConfigService) UpdateGameConfig(gameConfig navRequest.UpdateNavGameConfigRequest) (err error) {
	if global.GVA_DB == nil {
		global.GVA_LOG.Error("数据库连接为空")
		return errors.New("数据库连接异常")
	}

	// 检查配置是否存在
	var existingConfig navigation.NavGameConfig
	err = global.GVA_DB.Where("id = ?", gameConfig.ID).First(&existingConfig).Error
	if err != nil {
		global.GVA_LOG.Error("查询游戏配置失败: " + err.Error())
		return err
	}

	// 更新配置
	updateData := map[string]interface{}{
		"download_url":  gameConfig.DownloadUrl,
		"audio_url":     gameConfig.AudioUrl,
		"website_title": gameConfig.WebsiteTitle,
		"website_desc":  gameConfig.WebsiteDesc,
		"website_icon":  gameConfig.WebsiteIcon,
		"website_logo":  gameConfig.WebsiteLogo,
		"status":        gameConfig.Status,
	}

	err = global.GVA_DB.Model(&existingConfig).Updates(updateData).Error
	if err != nil {
		global.GVA_LOG.Error("更新游戏配置失败: " + err.Error())
		return err
	}

	return nil
}

// GetGameConfig 获取游戏配置
func (navGameConfigService *NavGameConfigService) GetGameConfig() (gameConfig navigation.NavGameConfig, err error) {
	if global.GVA_DB == nil {
		global.GVA_LOG.Error("数据库连接为空")
		return gameConfig, errors.New("数据库连接异常")
	}

	err = global.GVA_DB.First(&gameConfig).Error
	if err != nil {
		global.GVA_LOG.Error("获取游戏配置失败: " + err.Error())
		return gameConfig, err
	}

	return gameConfig, nil
}

// GetGameConfigList 获取游戏配置列表
func (navGameConfigService *NavGameConfigService) GetGameConfigList(info navRequest.NavGameConfigSearch) (list []navigation.NavGameConfig, total int64, err error) {
	if global.GVA_DB == nil {
		global.GVA_LOG.Error("数据库连接为空")
		return list, total, errors.New("数据库连接异常")
	}

	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	db := global.GVA_DB.Model(&navigation.NavGameConfig{})

	// 构建查询条件
	if info.WebsiteTitle != "" {
		db = db.Where("website_title LIKE ?", "%"+info.WebsiteTitle+"%")
	}
	if info.Status != 0 {
		db = db.Where("status = ?", info.Status)
	}

	// 排序
	if info.OrderKey != "" {
		order := info.OrderKey
		if info.Desc {
			order = order + " desc"
		}
		db = db.Order(order)
	} else {
		db = db.Order("created_at desc")
	}

	// 获取总数
	err = db.Count(&total).Error
	if err != nil {
		global.GVA_LOG.Error("获取游戏配置总数失败: " + err.Error())
		return list, total, err
	}

	// 分页查询
	err = db.Limit(limit).Offset(offset).Find(&list).Error
	if err != nil {
		global.GVA_LOG.Error("获取游戏配置列表失败: " + err.Error())
		return list, total, err
	}

	return list, total, nil
}

// DeleteGameConfig 删除游戏配置
func (navGameConfigService *NavGameConfigService) DeleteGameConfig(id uint) (err error) {
	if global.GVA_DB == nil {
		global.GVA_LOG.Error("数据库连接为空")
		return errors.New("数据库连接异常")
	}

	err = global.GVA_DB.Where("id = ?", id).Delete(&navigation.NavGameConfig{}).Error
	if err != nil {
		global.GVA_LOG.Error("删除游戏配置失败: " + err.Error())
		return err
	}

	return nil
}
