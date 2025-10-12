package navigation

import (
	"errors"
	"goLinker-admin/server/global"
	"goLinker-admin/server/model/navigation"
	navRequest "goLinker-admin/server/model/navigation/request"
	"strconv"
	"strings"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type NavGameService struct{}

// CreateGame 创建游戏
func (s *NavGameService) CreateGame(game navigation.NavGame) (err error) {
	if global.GVA_DB == nil {
		global.GVA_LOG.Error("数据库连接为空")
		return errors.New("数据库连接为空")
	}

	// 参数验证
	if game.Title == "" {
		return errors.New("游戏标题不能为空")
	}
	if game.CategoryIds == "" {
		return errors.New("游戏类别不能为空")
	}

	// 验证类别ID是否存在
	categoryIds := strings.Trim(game.CategoryIds, "[]")
	if categoryIds != "" {
		categoryIdList := strings.Split(categoryIds, ",")
		for _, idStr := range categoryIdList {
			idStr = strings.TrimSpace(idStr)
			if idStr != "" {
				categoryId, err := strconv.ParseUint(idStr, 10, 32)
				if err != nil {
					return errors.New("类别ID格式错误")
				}
				var category navigation.NavGameCategory
				err = global.GVA_DB.Where("id = ?", categoryId).First(&category).Error
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return errors.New("类别ID不存在: " + idStr)
				}
				if err != nil {
					global.GVA_LOG.Error("查询类别失败", zap.Error(err))
					return err
				}
			}
		}
	}

	// 创建游戏
	err = global.GVA_DB.Create(&game).Error
	if err != nil {
		global.GVA_LOG.Error("创建游戏失败", zap.Error(err))
		return err
	}

	global.GVA_LOG.Info("创建游戏成功", zap.String("title", game.Title))
	return nil
}

// UpdateGame 更新游戏
func (s *NavGameService) UpdateGame(game navigation.NavGame) (err error) {
	if global.GVA_DB == nil {
		global.GVA_LOG.Error("数据库连接为空")
		return errors.New("数据库连接为空")
	}

	// 参数验证
	if game.ID == 0 {
		return errors.New("游戏ID不能为空")
	}
	if game.Title == "" {
		return errors.New("游戏标题不能为空")
	}
	if game.CategoryIds == "" {
		return errors.New("游戏类别不能为空")
	}

	// 验证类别ID是否存在
	categoryIds := strings.Trim(game.CategoryIds, "[]")
	if categoryIds != "" {
		categoryIdList := strings.Split(categoryIds, ",")
		for _, idStr := range categoryIdList {
			idStr = strings.TrimSpace(idStr)
			if idStr != "" {
				categoryId, err := strconv.ParseUint(idStr, 10, 32)
				if err != nil {
					return errors.New("类别ID格式错误")
				}
				var category navigation.NavGameCategory
				err = global.GVA_DB.Where("id = ?", categoryId).First(&category).Error
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return errors.New("类别ID不存在: " + idStr)
				}
				if err != nil {
					global.GVA_LOG.Error("查询类别失败", zap.Error(err))
					return err
				}
			}
		}
	}

	// 更新游戏
	err = global.GVA_DB.Save(&game).Error
	if err != nil {
		global.GVA_LOG.Error("更新游戏失败", zap.Error(err))
		return err
	}

	global.GVA_LOG.Info("更新游戏成功", zap.Uint("id", game.ID))
	return nil
}

// DeleteGame 删除游戏
func (s *NavGameService) DeleteGame(id uint) (err error) {
	if global.GVA_DB == nil {
		global.GVA_LOG.Error("数据库连接为空")
		return errors.New("数据库连接为空")
	}

	// 参数验证
	if id == 0 {
		return errors.New("游戏ID不能为空")
	}

	// 删除游戏
	err = global.GVA_DB.Delete(&navigation.NavGame{}, id).Error
	if err != nil {
		global.GVA_LOG.Error("删除游戏失败", zap.Error(err))
		return err
	}

	global.GVA_LOG.Info("删除游戏成功", zap.Uint("id", id))
	return nil
}

// GetGameList 获取游戏列表
func (s *NavGameService) GetGameList(info navRequest.NavGameSearch) (list []navigation.NavGame, total int64, err error) {
	if global.GVA_DB == nil {
		global.GVA_LOG.Error("数据库连接为空")
		return list, 0, errors.New("数据库连接为空")
	}

	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&navigation.NavGame{})

	// 搜索条件
	if info.Title != "" {
		db = db.Where("title LIKE ?", "%"+info.Title+"%")
	}
	if info.CategoryId != 0 {
		db = db.Where("category_ids LIKE ?", "%"+strconv.FormatUint(uint64(info.CategoryId), 10)+"%")
	}
	if info.Status != 0 {
		db = db.Where("status = ?", info.Status)
	}
	if info.IsVisible != 0 {
		db = db.Where("is_visible = ?", info.IsVisible)
	}
	if info.Type != "" {
		db = db.Where("type = ?", info.Type)
	}

	// 获取总数
	err = db.Count(&total).Error
	if err != nil {
		global.GVA_LOG.Error("获取游戏总数失败", zap.Error(err))
		return list, 0, err
	}

	// 获取列表
	err = db.Limit(limit).Offset(offset).Order("sticky DESC, sort ASC, created_at DESC").Find(&list).Error
	if err != nil {
		global.GVA_LOG.Error("获取游戏列表失败", zap.Error(err))
		return list, 0, err
	}

	return list, total, nil
}

// UpdateGameViews 更新游戏浏览次数
func (s *NavGameService) UpdateGameViews(id uint) (err error) {
	if global.GVA_DB == nil {
		global.GVA_LOG.Error("数据库连接为空")
		return errors.New("数据库连接为空")
	}

	// 参数验证
	if id == 0 {
		return errors.New("游戏ID不能为空")
	}

	// 增加浏览次数
	err = global.GVA_DB.Model(&navigation.NavGame{}).Where("id = ?", id).Update("views", gorm.Expr("views + 1")).Error
	if err != nil {
		global.GVA_LOG.Error("更新游戏浏览次数失败", zap.Error(err))
		return err
	}

	return nil
}
