package navigation

import (
	"errors"
	"goLinker-admin/server/global"
	"goLinker-admin/server/model/navigation"
	navRequest "goLinker-admin/server/model/navigation/request"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type NavActivityConfigService struct{}

// CreateActivityConfig 创建活动配置
func (s *NavActivityConfigService) CreateActivityConfig(req navRequest.CreateNavActivityConfigRequest) (err error) {
	if global.GVA_DB == nil {
		global.GVA_LOG.Error("数据库连接为空")
		return errors.New("数据库连接为空")
	}

	// 参数验证
	if req.Title == "" {
		return errors.New("显示标题不能为空")
	}

	// 创建活动配置
	activityConfig := navigation.NavActivityConfig{
		Title:        req.Title,
		Image:        req.Image,
		JumpUrl:      req.JumpUrl,
		CategoryName: req.CategoryName,
		CategoryIcon: req.CategoryIcon,
		Content:      req.Content,
		Status:       req.Status,
		IsVisible:    req.IsVisible,
		Sort:         req.Sort,
	}

	err = global.GVA_DB.Create(&activityConfig).Error
	if err != nil {
		global.GVA_LOG.Error("创建活动配置失败", zap.Error(err))
		return err
	}

	global.GVA_LOG.Info("创建活动配置成功", zap.String("title", req.Title))
	return nil
}

// UpdateActivityConfig 更新活动配置
func (s *NavActivityConfigService) UpdateActivityConfig(req navRequest.UpdateNavActivityConfigRequest) (err error) {
	if global.GVA_DB == nil {
		global.GVA_LOG.Error("数据库连接为空")
		return errors.New("数据库连接为空")
	}

	// 参数验证
	if req.ID == 0 {
		return errors.New("活动配置ID不能为空")
	}

	// 检查活动配置是否存在
	var existingConfig navigation.NavActivityConfig
	err = global.GVA_DB.Where("id = ?", req.ID).First(&existingConfig).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("活动配置不存在")
	}
	if err != nil {
		global.GVA_LOG.Error("查询活动配置失败", zap.Error(err))
		return err
	}

	// 更新活动配置
	updateFields := make(map[string]interface{})
	if req.Title != nil {
		updateFields["title"] = *req.Title
	}
	if req.Image != nil {
		updateFields["image"] = *req.Image
	}
	if req.JumpUrl != nil {
		updateFields["jump_url"] = *req.JumpUrl
	}
	if req.CategoryName != nil {
		updateFields["category_name"] = *req.CategoryName
	}
	if req.CategoryIcon != nil {
		updateFields["category_icon"] = *req.CategoryIcon
	}
	if req.Content != nil {
		updateFields["content"] = *req.Content
	}
	if req.Status != nil {
		updateFields["status"] = *req.Status
	}
	if req.IsVisible != nil {
		updateFields["is_visible"] = *req.IsVisible
	}
	if req.Sort != nil {
		updateFields["sort"] = *req.Sort
	}

	err = global.GVA_DB.Model(&existingConfig).Updates(updateFields).Error
	if err != nil {
		global.GVA_LOG.Error("更新活动配置失败", zap.Error(err))
		return err
	}

	global.GVA_LOG.Info("更新活动配置成功", zap.Uint("id", req.ID))
	return nil
}

// DeleteActivityConfig 删除活动配置
func (s *NavActivityConfigService) DeleteActivityConfig(id uint) (err error) {
	if global.GVA_DB == nil {
		global.GVA_LOG.Error("数据库连接为空")
		return errors.New("数据库连接为空")
	}

	// 检查活动配置是否存在
	var existingConfig navigation.NavActivityConfig
	err = global.GVA_DB.Where("id = ?", id).First(&existingConfig).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("活动配置不存在")
	}
	if err != nil {
		global.GVA_LOG.Error("查询活动配置失败", zap.Error(err))
		return err
	}

	// 删除活动配置
	err = global.GVA_DB.Delete(&existingConfig).Error
	if err != nil {
		global.GVA_LOG.Error("删除活动配置失败", zap.Error(err))
		return err
	}

	global.GVA_LOG.Info("删除活动配置成功", zap.Uint("id", id))
	return nil
}

// GetActivityConfigList 获取活动配置列表
func (s *NavActivityConfigService) GetActivityConfigList(req navRequest.NavActivityConfigSearch) (list []navigation.NavActivityConfig, total int64, err error) {
	if global.GVA_DB == nil {
		global.GVA_LOG.Error("数据库连接为空")
		return nil, 0, errors.New("数据库连接为空")
	}

	db := global.GVA_DB.Model(&navigation.NavActivityConfig{})

	// 搜索条件
	if req.Title != "" {
		db = db.Where("title LIKE ?", "%"+req.Title+"%")
	}
	if req.CategoryName != "" {
		db = db.Where("category_name LIKE ?", "%"+req.CategoryName+"%")
	}
	if req.Status != 0 {
		db = db.Where("status = ?", req.Status)
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
		global.GVA_LOG.Error("获取活动配置总数失败", zap.Error(err))
		return nil, 0, err
	}

	// 分页查询
	offset := (req.Page - 1) * req.PageSize
	err = db.Limit(req.PageSize).Offset(offset).Order("sort ASC, created_at DESC").Find(&list).Error
	if err != nil {
		global.GVA_LOG.Error("获取活动配置列表失败", zap.Error(err))
		return nil, 0, err
	}

	return list, total, nil
}

// GetActivityConfigById 根据ID获取活动配置
func (s *NavActivityConfigService) GetActivityConfigById(id uint) (config navigation.NavActivityConfig, err error) {
	if global.GVA_DB == nil {
		global.GVA_LOG.Error("数据库连接为空")
		return config, errors.New("数据库连接为空")
	}

	err = global.GVA_DB.Where("id = ?", id).First(&config).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return config, errors.New("活动配置不存在")
	}
	if err != nil {
		global.GVA_LOG.Error("查询活动配置失败", zap.Error(err))
		return config, err
	}

	return config, nil
}

// GetVisibleActivityConfigs 获取可见的活动配置列表（公开接口使用）
func (s *NavActivityConfigService) GetVisibleActivityConfigs(limit int) (list []navigation.NavActivityConfig, err error) {
	if global.GVA_DB == nil {
		global.GVA_LOG.Error("数据库连接为空")
		return nil, errors.New("数据库连接为空")
	}

	err = global.GVA_DB.Where("status = ? AND is_visible = ?", 1, 1).
		Order("sort ASC, created_at DESC").
		Limit(limit).
		Find(&list).Error
	if err != nil {
		global.GVA_LOG.Error("获取可见活动配置失败", zap.Error(err))
		return nil, err
	}

	return list, nil
}
