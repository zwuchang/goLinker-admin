package navigation

import (
	"errors"
	"goLinker-admin/server/global"
	"goLinker-admin/server/model/navigation"
	"goLinker-admin/server/model/navigation/request"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type NavBannerService struct{}

// GetBannerList 获取Banner列表
func (s *NavBannerService) GetBannerList(req request.NavBannerSearch) (list []navigation.NavBanner, total int64, err error) {
	if global.GVA_DB == nil {
		global.GVA_LOG.Error("数据库连接为空")
		return nil, 0, errors.New("数据库连接为空")
	}

	db := global.GVA_DB.Model(&navigation.NavBanner{})

	// 搜索条件
	if req.Title != "" {
		db = db.Where("title LIKE ?", "%"+req.Title+"%")
	}
	if req.MediaType != nil {
		db = db.Where("media_type = ?", *req.MediaType)
	}
	if req.IsVisible != nil {
		db = db.Where("is_visible = ?", *req.IsVisible)
	}
	if req.Status != nil {
		db = db.Where("status = ?", *req.Status)
	}

	// 安全排序
	orderClause := s.buildSafeOrderClause(req.OrderBy, req.OrderType)
	if orderClause != "" {
		db = db.Order(orderClause)
	} else {
		db = db.Order("sort ASC, created_at DESC")
	}

	// 分页
	err = db.Count(&total).Error
	if err != nil {
		global.GVA_LOG.Error("获取Banner总数失败", zap.Error(err))
		return nil, 0, err
	}

	err = db.Limit(req.PageSize).Offset((req.Page - 1) * req.PageSize).Find(&list).Error
	if err != nil {
		global.GVA_LOG.Error("获取Banner列表失败", zap.Error(err))
		return nil, 0, err
	}

	return list, total, nil
}

// GetBannerById 根据ID获取Banner
func (s *NavBannerService) GetBannerById(id uint) (banner navigation.NavBanner, err error) {
	if global.GVA_DB == nil {
		global.GVA_LOG.Error("数据库连接为空")
		return banner, errors.New("数据库连接为空")
	}

	err = global.GVA_DB.Where("id = ?", id).First(&banner).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return banner, errors.New("Banner不存在")
		}
		global.GVA_LOG.Error("获取Banner失败", zap.Error(err))
		return banner, err
	}

	return banner, nil
}

// CreateBanner 创建Banner
func (s *NavBannerService) CreateBanner(req request.NavBannerCreateRequest) (err error) {
	if global.GVA_DB == nil {
		global.GVA_LOG.Error("数据库连接为空")
		return errors.New("数据库连接为空")
	}

	// 参数验证
	if req.MediaType != 1 && req.MediaType != 2 {
		return errors.New("媒体类型只能是1(图片)或2(视频)")
	}

	banner := navigation.NavBanner{
		Title:       req.Title,
		Description: req.Description,
		MediaType:   req.MediaType,
		MediaUrl:    req.MediaUrl,
		LinkUrl:     req.LinkUrl,
		Sort:        req.Sort,
		IsVisible:   req.IsVisible,
		Status:      req.Status,
	}

	err = global.GVA_DB.Create(&banner).Error
	if err != nil {
		global.GVA_LOG.Error("创建Banner失败", zap.Error(err))
		return err
	}

	global.GVA_LOG.Info("创建Banner成功", zap.String("title", req.Title))
	return nil
}

// UpdateBanner 更新Banner
func (s *NavBannerService) UpdateBanner(req request.NavBannerUpdateRequest) (err error) {
	if global.GVA_DB == nil {
		global.GVA_LOG.Error("数据库连接为空")
		return errors.New("数据库连接为空")
	}

	// 参数验证
	if req.ID == 0 {
		return errors.New("Banner ID不能为空")
	}

	// 检查Banner是否存在
	var existingBanner navigation.NavBanner
	err = global.GVA_DB.Where("id = ?", req.ID).First(&existingBanner).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("Banner不存在")
	}
	if err != nil {
		global.GVA_LOG.Error("查询Banner失败", zap.Error(err))
		return err
	}

	// 构建更新字段映射
	updateFields := make(map[string]interface{})

	// 验证媒体类型
	if req.MediaType != 1 && req.MediaType != 2 {
		return errors.New("媒体类型只能是1(图片)或2(视频)")
	}

	updateFields["title"] = req.Title
	updateFields["description"] = req.Description
	updateFields["media_type"] = req.MediaType
	updateFields["media_url"] = req.MediaUrl
	updateFields["link_url"] = req.LinkUrl
	updateFields["sort"] = req.Sort
	updateFields["is_visible"] = req.IsVisible
	updateFields["status"] = req.Status

	err = global.GVA_DB.Model(&existingBanner).Updates(updateFields).Error
	if err != nil {
		global.GVA_LOG.Error("更新Banner失败", zap.Error(err))
		return err
	}

	global.GVA_LOG.Info("更新Banner成功", zap.Uint("id", req.ID), zap.Any("updated_fields", updateFields))
	return nil
}

// DeleteBanner 删除Banner
func (s *NavBannerService) DeleteBanner(id uint) (err error) {
	if global.GVA_DB == nil {
		global.GVA_LOG.Error("数据库连接为空")
		return errors.New("数据库连接为空")
	}

	// 检查Banner是否存在
	var banner navigation.NavBanner
	err = global.GVA_DB.Where("id = ?", id).First(&banner).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("Banner不存在")
	}
	if err != nil {
		global.GVA_LOG.Error("查询Banner失败", zap.Error(err))
		return err
	}

	err = global.GVA_DB.Delete(&banner).Error
	if err != nil {
		global.GVA_LOG.Error("删除Banner失败", zap.Error(err))
		return err
	}

	global.GVA_LOG.Info("删除Banner成功", zap.Uint("id", id))
	return nil
}

// buildSafeOrderClause 构建安全的排序子句
func (s *NavBannerService) buildSafeOrderClause(orderBy, orderType string) string {
	// 允许的排序字段白名单
	allowedFields := map[string]string{
		"id":         "id",
		"title":      "title",
		"media_type": "media_type",
		"sort":       "sort",
		"is_visible": "is_visible",
		"status":     "status",
		"created_at": "created_at",
		"updated_at": "updated_at",
	}

	// 允许的排序类型
	allowedOrderTypes := map[string]string{
		"asc":  "ASC",
		"desc": "DESC",
	}

	field, exists := allowedFields[orderBy]
	if !exists {
		return ""
	}

	order, exists := allowedOrderTypes[orderType]
	if !exists {
		return ""
	}

	return field + " " + order
}
