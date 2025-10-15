package navigation

import (
	"errors"
	"goLinker-admin/server/global"
	"goLinker-admin/server/model/common/request"
	"goLinker-admin/server/model/navigation"
	"net/url"
	"strings"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ContactMethodService struct{}

// CreateNavContactMethod 创建联系方式
func (s *ContactMethodService) CreateNavContactMethod(method navigation.NavContactMethod) (err error) {
	if global.GVA_DB == nil {
		global.GVA_LOG.Error("数据库连接为空")
		return errors.New("数据库连接为空")
	}

	// 参数验证
	if err := s.validateContactMethod(&method); err != nil {
		global.GVA_LOG.Error("联系方式参数验证失败", zap.Error(err))
		return err
	}

	err = global.GVA_DB.Create(&method).Error
	if err != nil {
		global.GVA_LOG.Error("创建联系方式失败", zap.Error(err))
		return err
	}

	global.GVA_LOG.Info("联系方式创建成功", zap.Uint("id", method.ID))
	return nil
}

// DeleteNavContactMethod 删除联系方式
func (s *ContactMethodService) DeleteNavContactMethod(method navigation.NavContactMethod) (err error) {
	if global.GVA_DB == nil {
		global.GVA_LOG.Error("数据库连接为空")
		return errors.New("数据库连接为空")
	}

	// 参数验证
	if method.ID == 0 {
		global.GVA_LOG.Error("联系方式ID不能为空")
		return errors.New("联系方式ID不能为空")
	}

	err = global.GVA_DB.Delete(&method).Error
	if err != nil {
		global.GVA_LOG.Error("删除联系方式失败", zap.Error(err), zap.Uint("id", method.ID))
		return err
	}

	global.GVA_LOG.Info("联系方式删除成功", zap.Uint("id", method.ID))
	return nil
}

// UpdateNavContactMethod 更新联系方式
func (s *ContactMethodService) UpdateNavContactMethod(method *navigation.NavContactMethod) (err error) {
	if global.GVA_DB == nil {
		global.GVA_LOG.Error("数据库连接为空")
		return errors.New("数据库连接为空")
	}

	// 参数验证
	if method.ID == 0 {
		global.GVA_LOG.Error("联系方式ID不能为空")
		return errors.New("联系方式ID不能为空")
	}

	if err := s.validateContactMethod(method); err != nil {
		global.GVA_LOG.Error("联系方式参数验证失败", zap.Error(err))
		return err
	}

	err = global.GVA_DB.Save(method).Error
	if err != nil {
		global.GVA_LOG.Error("更新联系方式失败", zap.Error(err), zap.Uint("id", method.ID))
		return err
	}

	global.GVA_LOG.Info("联系方式更新成功", zap.Uint("id", method.ID))
	return nil
}

// GetNavContactMethod 获取单个联系方式
func (s *ContactMethodService) GetNavContactMethod(id uint) (method navigation.NavContactMethod, err error) {
	if global.GVA_DB == nil {
		global.GVA_LOG.Error("数据库连接为空")
		return method, errors.New("数据库连接为空")
	}

	if id == 0 {
		global.GVA_LOG.Error("联系方式ID不能为空")
		return method, errors.New("联系方式ID不能为空")
	}

	err = global.GVA_DB.Where("id = ?", id).First(&method).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			global.GVA_LOG.Error("联系方式不存在", zap.Uint("id", id))
			return method, errors.New("联系方式不存在")
		}
		global.GVA_LOG.Error("获取联系方式失败", zap.Error(err), zap.Uint("id", id))
		return method, err
	}

	return method, nil
}

// GetNavContactMethodList 分页获取联系方式列表
func (s *ContactMethodService) GetNavContactMethodList(info request.PageInfo) (list interface{}, total int64, err error) {
	if global.GVA_DB == nil {
		global.GVA_LOG.Error("数据库连接为空")
		return nil, 0, errors.New("数据库连接为空")
	}

	// 参数验证
	if info.Page <= 0 {
		info.Page = 1
	}
	if info.PageSize <= 0 {
		info.PageSize = 10
	}
	if info.PageSize > 100 {
		info.PageSize = 100
	}

	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&navigation.NavContactMethod{})

	var methodList []navigation.NavContactMethod
	err = db.Count(&total).Error
	if err != nil {
		global.GVA_LOG.Error("统计联系方式数量失败", zap.Error(err))
		return methodList, total, err
	}

	err = db.Limit(limit).Offset(offset).Order("sort ASC, created_at DESC").Find(&methodList).Error
	if err != nil {
		global.GVA_LOG.Error("获取联系方式列表失败", zap.Error(err))
		return methodList, total, err
	}

	return methodList, total, nil
}

// GetEnabledContactMethods 获取启用的联系方式列表（按排序值排序）
func (s *ContactMethodService) GetEnabledContactMethods() (list []navigation.NavContactMethod, err error) {
	if global.GVA_DB == nil {
		global.GVA_LOG.Error("数据库连接为空")
		return nil, errors.New("数据库连接为空")
	}

	err = global.GVA_DB.Where("status = ?", 1).Order("sort ASC, created_at DESC").Find(&list).Error
	if err != nil {
		global.GVA_LOG.Error("获取启用的联系方式列表失败", zap.Error(err))
		return list, err
	}

	return list, nil
}

// validateContactMethod 验证联系方式参数
func (s *ContactMethodService) validateContactMethod(method *navigation.NavContactMethod) error {
	if method == nil {
		return errors.New("联系方式参数不能为空")
	}

	// 验证显示名称
	if strings.TrimSpace(method.DisplayName) == "" {
		return errors.New("显示名称不能为空")
	}
	if len(method.DisplayName) > 50 {
		return errors.New("显示名称长度不能超过50个字符")
	}

	// 验证跳转地址
	if strings.TrimSpace(method.JumpUrl) == "" {
		return errors.New("跳转地址不能为空")
	}
	if len(method.JumpUrl) > 500 {
		return errors.New("跳转地址长度不能超过500个字符")
	}

	// 验证URL格式
	if _, err := url.Parse(method.JumpUrl); err != nil {
		return errors.New("跳转地址格式不正确")
	}

	// 验证图片URL长度
	if len(method.Image) > 500 {
		return errors.New("图片URL长度不能超过500个字符")
	}

	// 验证排序
	if method.Sort < 0 {
		return errors.New("排序值不能小于0")
	}
	if method.Sort > 9999 {
		return errors.New("排序值不能大于9999")
	}

	// 验证状态
	if method.Status != 0 && method.Status != 1 {
		return errors.New("状态值只能是0或1")
	}

	return nil
}
