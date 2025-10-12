package navigation

import (
	"errors"
	"goLinker-admin/server/global"
	"goLinker-admin/server/model/navigation"
	"regexp"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ContactConfigService struct{}

// GetNavContactConfig 获取联系配置
func (s *ContactConfigService) GetNavContactConfig() (config navigation.NavContactConfig, err error) {
	if global.GVA_DB == nil {
		global.GVA_LOG.Error("数据库连接为空")
		return config, errors.New("数据库连接为空")
	}

	err = global.GVA_DB.First(&config).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 如果没有记录，返回空结构体
			global.GVA_LOG.Info("未找到联系配置记录，返回空配置")
			return navigation.NavContactConfig{}, nil
		}
		global.GVA_LOG.Error("获取联系配置失败", zap.Error(err))
		return config, err
	}
	return config, nil
}

// UpdateNavContactConfig 更新联系配置
func (s *ContactConfigService) UpdateNavContactConfig(config *navigation.NavContactConfig) (err error) {
	if global.GVA_DB == nil {
		global.GVA_LOG.Error("数据库连接为空")
		return errors.New("数据库连接为空")
	}

	// 参数验证
	if err := s.validateContactConfig(config); err != nil {
		global.GVA_LOG.Error("联系配置参数验证失败", zap.Error(err))
		return err
	}

	var existingConfig navigation.NavContactConfig
	err = global.GVA_DB.First(&existingConfig).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 如果没有记录，创建新记录
			global.GVA_LOG.Info("创建新的联系配置")
			err = global.GVA_DB.Create(config).Error
			if err != nil {
				global.GVA_LOG.Error("创建联系配置失败", zap.Error(err))
				return err
			}
			return nil
		}
		global.GVA_LOG.Error("查询联系配置失败", zap.Error(err))
		return err
	}

	// 更新现有记录
	config.ID = existingConfig.ID
	err = global.GVA_DB.Save(config).Error
	if err != nil {
		global.GVA_LOG.Error("更新联系配置失败", zap.Error(err))
		return err
	}

	global.GVA_LOG.Info("联系配置更新成功")
	return nil
}

// validateContactConfig 验证联系配置参数
func (s *ContactConfigService) validateContactConfig(config *navigation.NavContactConfig) error {
	if config == nil {
		return errors.New("配置参数不能为空")
	}

	// 验证邮箱格式
	if config.Email != "" {
		emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
		if !emailRegex.MatchString(config.Email) {
			return errors.New("邮箱格式不正确")
		}
	}

	// 验证邮箱长度
	if len(config.Email) > 100 {
		return errors.New("邮箱长度不能超过100个字符")
	}

	// 验证横幅图片URL长度
	if len(config.BannerImage) > 500 {
		return errors.New("横幅图片URL长度不能超过500个字符")
	}

	return nil
}


