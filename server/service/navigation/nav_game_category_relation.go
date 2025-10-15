package navigation

import (
	"errors"
	"goLinker-admin/server/global"
	"goLinker-admin/server/model/navigation"
	"strconv"
	"strings"

	"go.uber.org/zap"
)

type NavGameCategoryRelationService struct{}

// CreateGameCategoryRelations 创建游戏分类关联
func (s *NavGameCategoryRelationService) CreateGameCategoryRelations(gameId uint, categoryIds string) error {
	if global.GVA_DB == nil {
		return errors.New("数据库连接为空")
	}

	// 解析分类ID字符串
	categoryIdList, err := s.parseCategoryIds(categoryIds)
	if err != nil {
		return err
	}

	// 删除现有关联
	err = s.DeleteGameCategoryRelations(gameId)
	if err != nil {
		return err
	}

	// 创建新关联
	var relations []navigation.NavGameCategoryRelation
	for i, categoryId := range categoryIdList {
		relations = append(relations, navigation.NavGameCategoryRelation{
			GameId:     gameId,
			CategoryId: categoryId,
			Sort:       i + 1, // 按顺序设置排序
		})
	}

	if len(relations) > 0 {
		err = global.GVA_DB.Create(&relations).Error
		if err != nil {
			global.GVA_LOG.Error("创建游戏分类关联失败", zap.Error(err))
			return err
		}
	}

	global.GVA_LOG.Info("创建游戏分类关联成功", zap.Uint("gameId", gameId), zap.Int("count", len(relations)))
	return nil
}

// DeleteGameCategoryRelations 删除游戏的所有分类关联
func (s *NavGameCategoryRelationService) DeleteGameCategoryRelations(gameId uint) error {
	if global.GVA_DB == nil {
		return errors.New("数据库连接为空")
	}

	err := global.GVA_DB.Where("game_id = ?", gameId).Delete(&navigation.NavGameCategoryRelation{}).Error
	if err != nil {
		global.GVA_LOG.Error("删除游戏分类关联失败", zap.Error(err))
		return err
	}

	return nil
}

// GetGameCategoryIds 获取游戏的所有分类ID
func (s *NavGameCategoryRelationService) GetGameCategoryIds(gameId uint) ([]uint, error) {
	if global.GVA_DB == nil {
		return nil, errors.New("数据库连接为空")
	}

	var relations []navigation.NavGameCategoryRelation
	err := global.GVA_DB.Where("game_id = ?", gameId).Order("sort ASC").Find(&relations).Error
	if err != nil {
		global.GVA_LOG.Error("获取游戏分类关联失败", zap.Error(err))
		return nil, err
	}

	var categoryIds []uint
	for _, relation := range relations {
		categoryIds = append(categoryIds, relation.CategoryId)
	}

	return categoryIds, nil
}

// SyncCategoryIds 同步CategoryIds字段（用于兼容性）
func (s *NavGameCategoryRelationService) SyncCategoryIds(gameId uint) error {
	if global.GVA_DB == nil {
		return errors.New("数据库连接为空")
	}

	// 获取分类ID列表
	categoryIds, err := s.GetGameCategoryIds(gameId)
	if err != nil {
		return err
	}

	// 转换为JSON字符串格式
	categoryIdsStr := s.formatCategoryIds(categoryIds)

	// 更新NavGame表的CategoryIds字段
	err = global.GVA_DB.Model(&navigation.NavGame{}).
		Where("id = ?", gameId).
		Update("category_ids", categoryIdsStr).Error
	if err != nil {
		global.GVA_LOG.Error("同步分类ID字段失败", zap.Error(err))
		return err
	}

	return nil
}

// parseCategoryIds 解析分类ID字符串
func (s *NavGameCategoryRelationService) parseCategoryIds(categoryIds string) ([]uint, error) {
	if categoryIds == "" {
		return []uint{}, nil
	}

	// 移除方括号并分割
	idsStr := strings.Trim(categoryIds, "[]")
	if idsStr == "" {
		return []uint{}, nil
	}

	// 分割并解析ID
	idList := strings.Split(idsStr, ",")
	var ids []uint
	for _, idStr := range idList {
		idStr = strings.TrimSpace(idStr)
		if idStr != "" {
			id, err := strconv.ParseUint(idStr, 10, 32)
			if err != nil {
				return nil, errors.New("类别ID格式错误: " + idStr)
			}
			ids = append(ids, uint(id))
		}
	}

	return ids, nil
}

// formatCategoryIds 格式化分类ID为JSON字符串
func (s *NavGameCategoryRelationService) formatCategoryIds(categoryIds []uint) string {
	if len(categoryIds) == 0 {
		return "[]"
	}

	var idStrs []string
	for _, id := range categoryIds {
		idStrs = append(idStrs, strconv.FormatUint(uint64(id), 10))
	}

	return "[" + strings.Join(idStrs, ",") + "]"
}
