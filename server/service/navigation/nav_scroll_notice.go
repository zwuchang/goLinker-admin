package navigation

import (
	"goLinker-admin/server/global"
	"goLinker-admin/server/model/navigation"
	navRequest "goLinker-admin/server/model/navigation/request"
	"time"
)

type NavScrollNoticeService struct{}

// GetScrollNoticeList 获取滚动通知列表
func (s *NavScrollNoticeService) GetScrollNoticeList(info navRequest.NavScrollNoticeSearch) (list []navigation.NavScrollNotice, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&navigation.NavScrollNotice{})

	// 添加搜索条件
	if info.Title != "" {
		db = db.Where("title LIKE ?", "%"+info.Title+"%")
	}
	if info.Status != 0 {
		db = db.Where("status = ?", info.Status)
	}
	if info.IsVisible != 0 {
		db = db.Where("is_visible = ?", info.IsVisible)
	}

	var scrollNotices []navigation.NavScrollNotice
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Order("sort ASC, created_at DESC").Limit(limit).Offset(offset).Find(&scrollNotices).Error
	return scrollNotices, total, err
}

// CreateScrollNotice 创建滚动通知
func (s *NavScrollNoticeService) CreateScrollNotice(req navRequest.CreateNavScrollNoticeRequest) (err error) {
	scrollNotice := navigation.NavScrollNotice{
		Title:     req.Title,
		Content:   req.Content,
		Status:    req.Status,
		IsVisible: req.IsVisible,
		Sort:      req.Sort,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
	}

	// 处理空的时间字段，设置默认的超长时间范围
	now := time.Now()
	if scrollNotice.StartTime == "" {
		scrollNotice.StartTime = now.Format("2006-01-02 15:04:05")
	}
	if scrollNotice.EndTime == "" {
		// 设置为10年后
		futureTime := now.AddDate(10, 0, 0)
		scrollNotice.EndTime = futureTime.Format("2006-01-02 15:04:05")
	}

	err = global.GVA_DB.Create(&scrollNotice).Error
	return err
}

// UpdateScrollNotice 更新滚动通知
func (s *NavScrollNoticeService) UpdateScrollNotice(req navRequest.UpdateNavScrollNoticeRequest) (err error) {
	var scrollNotice navigation.NavScrollNotice
	err = global.GVA_DB.First(&scrollNotice, req.ID).Error
	if err != nil {
		return err
	}

	updateFields := make(map[string]interface{})
	if req.Title != nil {
		updateFields["title"] = *req.Title
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
	if req.StartTime != nil {
		if *req.StartTime == "" {
			// 设置为当前时间
			now := time.Now()
			updateFields["start_time"] = now.Format("2006-01-02 15:04:05")
		} else {
			updateFields["start_time"] = *req.StartTime
		}
	}
	if req.EndTime != nil {
		if *req.EndTime == "" {
			// 设置为10年后
			now := time.Now()
			futureTime := now.AddDate(10, 0, 0)
			updateFields["end_time"] = futureTime.Format("2006-01-02 15:04:05")
		} else {
			updateFields["end_time"] = *req.EndTime
		}
	}

	err = global.GVA_DB.Model(&scrollNotice).Updates(updateFields).Error
	return err
}

// DeleteScrollNotice 删除滚动通知
func (s *NavScrollNoticeService) DeleteScrollNotice(id uint) (err error) {
	err = global.GVA_DB.Delete(&navigation.NavScrollNotice{}, id).Error
	return err
}

// GetScrollNoticeById 根据ID获取滚动通知
func (s *NavScrollNoticeService) GetScrollNoticeById(id uint) (scrollNotice navigation.NavScrollNotice, err error) {
	err = global.GVA_DB.First(&scrollNotice, id).Error
	return scrollNotice, err
}

// GetVisibleScrollNotices 获取可见的滚动通知列表（公开接口使用）
func (s *NavScrollNoticeService) GetVisibleScrollNotices(limit int) (list []navigation.NavScrollNotice, err error) {
	now := time.Now().Format("2006-01-02 15:04:05")

	err = global.GVA_DB.Where("status = ? AND is_visible = ? AND start_time <= ? AND end_time >= ?",
		1, 1, now, now).
		Order("sort ASC, created_at DESC").
		Limit(limit).
		Find(&list).Error
	return list, err
}
