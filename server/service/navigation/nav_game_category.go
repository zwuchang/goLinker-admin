package navigation

import (
	"errors"
	"goLinker-admin/server/global"
	"goLinker-admin/server/model/navigation"
	navRequest "goLinker-admin/server/model/navigation/request"
	"strconv"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type NavGameCategoryService struct{}

// CreateGameCategory 创建游戏类别
func (s *NavGameCategoryService) CreateGameCategory(category navigation.NavGameCategory) (err error) {
	if global.GVA_DB == nil {
		global.GVA_LOG.Error("数据库连接为空")
		return errors.New("数据库连接为空")
	}

	// 参数验证
	if category.Name == "" {
		return errors.New("类别名称不能为空")
	}

	// 检查名称是否重复
	var existingCategory navigation.NavGameCategory
	err = global.GVA_DB.Where("name = ?", category.Name).First(&existingCategory).Error
	if err == nil {
		return errors.New("类别名称已存在")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		global.GVA_LOG.Error("查询类别失败", zap.Error(err))
		return err
	}

	// 创建类别
	err = global.GVA_DB.Create(&category).Error
	if err != nil {
		global.GVA_LOG.Error("创建游戏类别失败", zap.Error(err))
		return err
	}

	global.GVA_LOG.Info("创建游戏类别成功", zap.String("name", category.Name))
	return nil
}

// UpdateGameCategory 更新游戏类别
func (s *NavGameCategoryService) UpdateGameCategory(category navigation.NavGameCategory) (err error) {
	if global.GVA_DB == nil {
		global.GVA_LOG.Error("数据库连接为空")
		return errors.New("数据库连接为空")
	}

	// 参数验证
	if category.ID == 0 {
		return errors.New("类别ID不能为空")
	}
	if category.Name == "" {
		return errors.New("类别名称不能为空")
	}

	// 检查名称是否重复（排除自己）
	var existingCategory navigation.NavGameCategory
	err = global.GVA_DB.Where("name = ? AND id != ?", category.Name, category.ID).First(&existingCategory).Error
	if err == nil {
		return errors.New("类别名称已存在")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		global.GVA_LOG.Error("查询类别失败", zap.Error(err))
		return err
	}

	// 更新类别
	err = global.GVA_DB.Save(&category).Error
	if err != nil {
		global.GVA_LOG.Error("更新游戏类别失败", zap.Error(err))
		return err
	}

	global.GVA_LOG.Info("更新游戏类别成功", zap.Uint("id", category.ID))
	return nil
}

// DeleteGameCategory 删除游戏类别
func (s *NavGameCategoryService) DeleteGameCategory(id uint) (err error) {
	if global.GVA_DB == nil {
		global.GVA_LOG.Error("数据库连接为空")
		return errors.New("数据库连接为空")
	}

	// 参数验证
	if id == 0 {
		return errors.New("类别ID不能为空")
	}

	// 检查是否有游戏使用此类别
	var count int64
	err = global.GVA_DB.Model(&navigation.NavGame{}).Where("category_ids LIKE ?", "%"+strconv.FormatUint(uint64(id), 10)+"%").Count(&count).Error
	if err != nil {
		global.GVA_LOG.Error("查询游戏数量失败", zap.Error(err))
		return err
	}
	if count > 0 {
		return errors.New("该类别下还有游戏，无法删除")
	}

	// 删除类别
	err = global.GVA_DB.Delete(&navigation.NavGameCategory{}, id).Error
	if err != nil {
		global.GVA_LOG.Error("删除游戏类别失败", zap.Error(err))
		return err
	}

	global.GVA_LOG.Info("删除游戏类别成功", zap.Uint("id", id))
	return nil
}

// GetGameCategoryList 获取游戏类别列表
func (s *NavGameCategoryService) GetGameCategoryList(info navRequest.NavGameCategorySearch) (list []navigation.NavGameCategory, total int64, err error) {
	if global.GVA_DB == nil {
		global.GVA_LOG.Error("数据库连接为空")
		return list, 0, errors.New("数据库连接为空")
	}

	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&navigation.NavGameCategory{})

	// 搜索条件
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.Status != 0 {
		db = db.Where("status = ?", info.Status)
	}

	// 获取总数
	err = db.Count(&total).Error
	if err != nil {
		global.GVA_LOG.Error("获取游戏类别总数失败", zap.Error(err))
		return list, 0, err
	}

	// 获取列表
	err = db.Limit(limit).Offset(offset).Order("sort ASC, created_at DESC").Find(&list).Error
	if err != nil {
		global.GVA_LOG.Error("获取游戏类别列表失败", zap.Error(err))
		return list, 0, err
	}

	return list, total, nil
}

// GetAllGameCategories 获取所有启用的游戏类别
func (s *NavGameCategoryService) GetAllGameCategories() (list []navigation.NavGameCategory, err error) {
	if global.GVA_DB == nil {
		global.GVA_LOG.Error("数据库连接为空")
		return list, errors.New("数据库连接为空")
	}

	err = global.GVA_DB.Where("status = ?", 1).Order("sort ASC, created_at DESC").Find(&list).Error
	if err != nil {
		global.GVA_LOG.Error("获取所有游戏类别失败", zap.Error(err))
		return list, err
	}

	return list, nil
}
