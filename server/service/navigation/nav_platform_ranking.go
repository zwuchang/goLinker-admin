package navigation

import (
	"goLinker-admin/server/global"
	"goLinker-admin/server/model/navigation"
	navRequest "goLinker-admin/server/model/navigation/request"
)

type NavPlatformRankingService struct{}

// GetPlatformRankingList 获取平台排行榜列表
func (s *NavPlatformRankingService) GetPlatformRankingList(info navRequest.NavPlatformRankingSearch) (list []navigation.NavPlatformRanking, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&navigation.NavPlatformRanking{})

	// 添加搜索条件
	if info.PlatformName != "" {
		db = db.Where("platform_name LIKE ?", "%"+info.PlatformName+"%")
	}
	if info.Status != 0 {
		db = db.Where("status = ?", info.Status)
	}

	var platformRankings []navigation.NavPlatformRanking
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Order("sort ASC, created_at ASC").Limit(limit).Offset(offset).Find(&platformRankings).Error
	return platformRankings, total, err
}

// CreatePlatformRanking 创建平台排行榜
func (s *NavPlatformRankingService) CreatePlatformRanking(req navRequest.CreateNavPlatformRankingRequest) (err error) {
	platformRanking := navigation.NavPlatformRanking{
		Rank:         req.Rank,
		PlatformName: req.PlatformName,
		Logo:         req.Logo,
		Rating:       req.Rating,
		Features:     req.Features,
		FeatureType:  req.FeatureType,
		VisitUrl:     req.VisitUrl,
		IsNew:        req.IsNew,
		Status:       req.Status,
		IsVisible:    req.IsVisible,
		Sort:         req.Sort,
	}
	err = global.GVA_DB.Create(&platformRanking).Error
	return err
}

// UpdatePlatformRanking 更新平台排行榜
func (s *NavPlatformRankingService) UpdatePlatformRanking(req navRequest.UpdateNavPlatformRankingRequest) (err error) {
	var platformRanking navigation.NavPlatformRanking
	err = global.GVA_DB.First(&platformRanking, req.ID).Error
	if err != nil {
		return err
	}

	updateFields := make(map[string]interface{})
	if req.Rank != nil {
		updateFields["rank"] = *req.Rank
	}
	if req.PlatformName != nil {
		updateFields["platform_name"] = *req.PlatformName
	}
	if req.Logo != nil {
		updateFields["logo"] = *req.Logo
	}
	if req.Rating != nil {
		updateFields["rating"] = *req.Rating
	}
	if req.Features != nil {
		updateFields["features"] = *req.Features
	}
	if req.FeatureType != nil {
		updateFields["feature_type"] = *req.FeatureType
	}
	if req.VisitUrl != nil {
		updateFields["visit_url"] = *req.VisitUrl
	}
	if req.IsNew != nil {
		updateFields["is_new"] = *req.IsNew
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

	err = global.GVA_DB.Model(&platformRanking).Updates(updateFields).Error
	return err
}

// DeletePlatformRanking 删除平台排行榜
func (s *NavPlatformRankingService) DeletePlatformRanking(id uint) (err error) {
	err = global.GVA_DB.Delete(&navigation.NavPlatformRanking{}, id).Error
	return err
}

// GetPlatformRankingById 根据ID获取平台排行榜
func (s *NavPlatformRankingService) GetPlatformRankingById(id uint) (platformRanking navigation.NavPlatformRanking, err error) {
	err = global.GVA_DB.First(&platformRanking, id).Error
	return platformRanking, err
}

// GetVisiblePlatformRankings 获取可见的平台排行榜列表（公开接口使用）
func (s *NavPlatformRankingService) GetVisiblePlatformRankings(limit int) (list []navigation.NavPlatformRanking, err error) {
	err = global.GVA_DB.Where("status = ? AND is_visible = ?", 1, 1).
		Order("sort ASC, rank ASC").
		Limit(limit).
		Find(&list).Error
	return list, err
}
