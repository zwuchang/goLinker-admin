package navigation

import (
	"errors"
	"goLinker-admin/server/global"
	"goLinker-admin/server/model/navigation"
	"goLinker-admin/server/model/navigation/request"
	"sort"
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

	// 验证类别ID是否存在并排序
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

	// 对分类ID进行排序
	sortedCategoryIds, err := sortCategoryIds(game.CategoryIds)
	if err != nil {
		return err
	}
	game.CategoryIds = sortedCategoryIds

	// 创建游戏
	err = global.GVA_DB.Create(&game).Error
	if err != nil {
		global.GVA_LOG.Error("创建游戏失败", zap.Error(err))
		return err
	}

	// 创建分类关联
	relationService := &NavGameCategoryRelationService{}
	err = relationService.CreateGameCategoryRelations(game.ID, game.CategoryIds)
	if err != nil {
		global.GVA_LOG.Error("创建游戏分类关联失败", zap.Error(err))
		// 注意：这里不返回错误，因为游戏已经创建成功，只是关联创建失败
	}

	global.GVA_LOG.Info("创建游戏成功", zap.String("title", game.Title))
	return nil
}

// UpdateGame 更新游戏
func (s *NavGameService) UpdateGame(req request.NavGameUpdateRequest) (err error) {
	if global.GVA_DB == nil {
		global.GVA_LOG.Error("数据库连接为空")
		return errors.New("数据库连接为空")
	}

	// 参数验证
	if req.ID == 0 {
		return errors.New("游戏ID不能为空")
	}

	// 检查游戏是否存在
	var existingGame navigation.NavGame
	err = global.GVA_DB.Where("id = ?", req.ID).First(&existingGame).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("游戏不存在")
	}
	if err != nil {
		global.GVA_LOG.Error("查询游戏失败", zap.Error(err))
		return err
	}

	// 构建更新字段映射
	updateFields := make(map[string]interface{})

	// 如果传递了标题，则更新标题
	if req.Title != nil {
		updateFields["title"] = *req.Title
	}

	// 如果传递了类别，则验证并更新类别
	if req.CategoryIds != nil {
		// 验证类别ID是否存在
		categoryIds := strings.Trim(*req.CategoryIds, "[]")
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

		// 对分类ID进行排序
		sortedCategoryIds, err := sortCategoryIds(*req.CategoryIds)
		if err != nil {
			return err
		}
		updateFields["category_ids"] = sortedCategoryIds

		// 更新分类关联
		relationService := &NavGameCategoryRelationService{}
		err = relationService.CreateGameCategoryRelations(req.ID, sortedCategoryIds)
		if err != nil {
			global.GVA_LOG.Error("更新游戏分类关联失败", zap.Error(err))
			// 注意：这里不返回错误，因为游戏更新成功，只是关联更新失败
		}
	}

	// 如果传递了其他字段，则更新对应字段
	if req.Thumbnail != nil {
		updateFields["thumbnail"] = *req.Thumbnail
	}
	if req.IsVisible != nil {
		updateFields["is_visible"] = *req.IsVisible
	}
	if req.Type != nil {
		updateFields["type"] = *req.Type
	}
	if req.VideoUrl != nil {
		updateFields["video_url"] = *req.VideoUrl
	}
	if req.VideoDuration != nil {
		updateFields["video_duration"] = *req.VideoDuration
	}
	if req.Sticky != nil {
		updateFields["sticky"] = *req.Sticky
	}
	if req.Views != nil {
		updateFields["views"] = *req.Views
	}
	if req.Description != nil {
		updateFields["description"] = *req.Description
	}
	if req.Content != nil {
		updateFields["content"] = *req.Content
	}
	if req.Status != nil {
		updateFields["status"] = *req.Status
	}
	if req.Sort != nil {
		updateFields["sort"] = *req.Sort
	}
	if req.JumpUrl != nil {
		updateFields["jump_url"] = *req.JumpUrl
	}
	if req.DownloadUrl != nil {
		updateFields["download_url"] = *req.DownloadUrl
	}
	if req.DisplayName != nil {
		updateFields["display_name"] = *req.DisplayName
	}
	if req.AdName != nil {
		updateFields["ad_name"] = *req.AdName
	}
	if req.Article != nil {
		updateFields["article"] = *req.Article
	}
	if req.Icon != nil {
		updateFields["icon"] = *req.Icon
	}

	// 如果没有传递任何更新字段，返回错误
	if len(updateFields) == 0 {
		return errors.New("没有传递任何需要更新的字段")
	}

	// 更新游戏
	err = global.GVA_DB.Model(&existingGame).Updates(updateFields).Error
	if err != nil {
		global.GVA_LOG.Error("更新游戏失败", zap.Error(err))
		return err
	}

	global.GVA_LOG.Info("更新游戏成功", zap.Uint("id", req.ID), zap.Any("updated_fields", updateFields))
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
func (s *NavGameService) GetGameList(info request.NavGameSearch) (list []navigation.NavGame, total int64, err error) {
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
		// 使用LEFT JOIN进行精确查询，支持分页和总数统计
		db = db.Joins("LEFT JOIN nav_game_category_relation ON nav_game.id = nav_game_category_relation.game_id").
			Where("nav_game_category_relation.category_id = ?", info.CategoryId)
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
	if info.DisplayName != "" {
		db = db.Where("display_name LIKE ?", "%"+info.DisplayName+"%")
	}
	if info.AdName != "" {
		db = db.Where("ad_name LIKE ?", "%"+info.AdName+"%")
	}

	// 获取总数
	err = db.Count(&total).Error
	if err != nil {
		global.GVA_LOG.Error("获取游戏总数失败", zap.Error(err))
		return list, 0, err
	}

	// 安全排序处理
	orderClause := s.buildSafeOrderClause(info.OrderBy, info.OrderType)

	// 获取列表
	err = db.Limit(limit).Offset(offset).Order(orderClause).Find(&list).Error
	if err != nil {
		global.GVA_LOG.Error("获取游戏列表失败", zap.Error(err))
		return list, 0, err
	}

	return list, total, nil
}

// GetGameById 根据ID获取游戏信息
func (s *NavGameService) GetGameById(id uint) (game navigation.NavGame, err error) {
	if global.GVA_DB == nil {
		global.GVA_LOG.Error("数据库连接为空")
		return game, errors.New("数据库连接为空")
	}

	// 参数验证
	if id == 0 {
		return game, errors.New("游戏ID不能为空")
	}

	err = global.GVA_DB.Where("id = ?", id).First(&game).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			global.GVA_LOG.Error("游戏不存在", zap.Uint("id", id))
			return game, errors.New("游戏不存在")
		}
		global.GVA_LOG.Error("获取游戏失败", zap.Error(err), zap.Uint("id", id))
		return game, err
	}

	return game, nil
}

// buildSafeOrderClause 构建安全的排序子句，防止SQL注入
func (s *NavGameService) buildSafeOrderClause(orderBy, orderType string) string {
	// 允许的排序字段白名单
	allowedFields := map[string]string{
		"id":           "id",
		"title":        "title",
		"type":         "type",
		"status":       "status",
		"is_visible":   "is_visible",
		"sticky":       "sticky",
		"views":        "views",
		"sort":         "sort",
		"display_name": "display_name",
		"ad_name":      "ad_name",
		"created_at":   "created_at",
		"updated_at":   "updated_at",
	}

	// 默认排序
	defaultOrder := "sticky DESC, sort ASC, created_at DESC"

	// 验证排序字段
	if orderBy == "" {
		return defaultOrder
	}

	field, exists := allowedFields[orderBy]
	if !exists {
		global.GVA_LOG.Warn("不允许的排序字段", zap.String("field", orderBy))
		return defaultOrder
	}

	// 验证排序类型
	var orderDirection string
	switch strings.ToLower(orderType) {
	case "asc":
		orderDirection = "ASC"
	case "desc":
		orderDirection = "DESC"
	default:
		global.GVA_LOG.Warn("不允许的排序类型", zap.String("type", orderType))
		return defaultOrder
	}

	// 构建安全的排序子句
	return field + " " + orderDirection
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

// sortCategoryIds 对分类ID进行排序，确保小的在前，大的在后
func sortCategoryIds(categoryIds string) (string, error) {
	if categoryIds == "" {
		return "", nil
	}

	// 移除方括号并分割
	idsStr := strings.Trim(categoryIds, "[]")
	if idsStr == "" {
		return "[]", nil
	}

	// 分割并解析ID
	idList := strings.Split(idsStr, ",")
	var ids []uint
	for _, idStr := range idList {
		idStr = strings.TrimSpace(idStr)
		if idStr != "" {
			id, err := strconv.ParseUint(idStr, 10, 32)
			if err != nil {
				return "", errors.New("类别ID格式错误: " + idStr)
			}
			ids = append(ids, uint(id))
		}
	}

	// 排序
	sort.Slice(ids, func(i, j int) bool {
		return ids[i] < ids[j]
	})

	// 重新组装为字符串
	var result []string
	for _, id := range ids {
		result = append(result, strconv.FormatUint(uint64(id), 10))
	}

	return "[" + strings.Join(result, ",") + "]", nil
}
