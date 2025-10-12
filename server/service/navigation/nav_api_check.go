package navigation

import (
	"errors"
	"goLinker-admin/server/global"
	"goLinker-admin/server/model/system"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type NavApiCheckService struct{}

// CheckAndCreateNavigationApis 检查并创建导航管理相关API
func (s *NavApiCheckService) CheckAndCreateNavigationApis() error {
	global.GVA_LOG.Info("开始检查并创建导航管理API...")

	// 定义导航管理相关API
	navigationApis := []system.SysApi{
		// 联系配置API
		{
			Path:        "/navigation/contactConfig/get",
			Description: "获取联系配置",
			ApiGroup:    "导航管理",
			Method:      "GET",
		},
		{
			Path:        "/navigation/contactConfig/update",
			Description: "更新联系配置",
			ApiGroup:    "导航管理",
			Method:      "PUT",
		},
		// 联系方式API
		{
			Path:        "/navigation/contactMethod/create",
			Description: "创建联系方式",
			ApiGroup:    "导航管理",
			Method:      "POST",
		},
		{
			Path:        "/navigation/contactMethod/update",
			Description: "更新联系方式",
			ApiGroup:    "导航管理",
			Method:      "PUT",
		},
		{
			Path:        "/navigation/contactMethod/delete",
			Description: "删除联系方式",
			ApiGroup:    "导航管理",
			Method:      "DELETE",
		},

		{
			Path:        "/navigation/contactMethod/list",
			Description: "获取联系方式列表",
			ApiGroup:    "导航管理",
			Method:      "GET",
		},
	}

	// 检查并创建每个API
	for _, api := range navigationApis {
		err := s.checkAndCreateApi(api)
		if err != nil {
			global.GVA_LOG.Error("API检查失败", zap.String("path", api.Path), zap.String("method", api.Method), zap.Error(err))
			return err
		}
	}

	global.GVA_LOG.Info("导航管理API检查完成")
	return nil
}

// checkAndCreateApi 检查并创建或更新单个API
func (s *NavApiCheckService) checkAndCreateApi(api system.SysApi) error {
	var existingApi system.SysApi
	err := global.GVA_DB.Where("path = ? AND method = ?", api.Path, api.Method).First(&existingApi).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		// API不存在，创建新API
		if err := global.GVA_DB.Create(&api).Error; err != nil {
			global.GVA_LOG.Error("创建API失败", zap.String("path", api.Path), zap.String("method", api.Method), zap.Error(err))
			return err
		}
		global.GVA_LOG.Info("创建API成功", zap.String("path", api.Path), zap.String("method", api.Method), zap.String("description", api.Description))
		return nil
	} else if err != nil {
		global.GVA_LOG.Error("查询API失败", zap.String("path", api.Path), zap.String("method", api.Method), zap.Error(err))
		return err
	} else {
		// API存在，检查是否需要更新
		needUpdate := false

		// 检查各个字段是否需要更新
		if existingApi.Description != api.Description {
			existingApi.Description = api.Description
			needUpdate = true
		}
		if existingApi.ApiGroup != api.ApiGroup {
			existingApi.ApiGroup = api.ApiGroup
			needUpdate = true
		}

		if needUpdate {
			if err := global.GVA_DB.Save(&existingApi).Error; err != nil {
				global.GVA_LOG.Error("更新API失败", zap.String("path", api.Path), zap.String("method", api.Method), zap.Error(err))
				return err
			}
			global.GVA_LOG.Info("更新API成功", zap.String("path", api.Path), zap.String("method", api.Method), zap.String("description", api.Description))
		} else {
			global.GVA_LOG.Debug("API已存在且无需更新", zap.String("path", api.Path), zap.String("method", api.Method))
		}
		return nil
	}
}

// CheckNavigationApisExist 检查导航管理API是否存在
func (s *NavApiCheckService) CheckNavigationApisExist() bool {
	if global.GVA_DB == nil {
		return false
	}

	// 检查关键API是否存在
	keyApis := []struct {
		Path   string
		Method string
	}{
		{"/navigation/contactConfig/get", "GET"},
		{"/navigation/contactConfig/update", "PUT"},
		{"/navigation/contactMethod/create", "POST"},
		{"/navigation/contactMethod/list", "GET"},
		{"/navigation/menuCheck/status", "GET"},
	}

	for _, api := range keyApis {
		var existingApi system.SysApi
		if errors.Is(global.GVA_DB.Where("path = ? AND method = ?", api.Path, api.Method).First(&existingApi).Error, gorm.ErrRecordNotFound) {
			return false
		}
	}
	return true
}

// GetNavigationApiStatus 获取导航API状态信息
func (s *NavApiCheckService) GetNavigationApiStatus() map[string]interface{} {
	status := make(map[string]interface{})

	if global.GVA_DB == nil {
		status["error"] = "数据库连接为空"
		return status
	}

	// 检查各个API的状态
	apis := []struct {
		Path   string
		Method string
		Name   string
	}{
		{"/navigation/contactConfig/get", "GET", "获取联系配置"},
		{"/navigation/contactConfig/update", "PUT", "更新联系配置"},
		{"/navigation/contactMethod/create", "POST", "创建联系方式"},
		{"/navigation/contactMethod/update", "PUT", "更新联系方式"},
		{"/navigation/contactMethod/delete", "DELETE", "删除联系方式"},
		{"/navigation/contactMethod/get", "GET", "获取单个联系方式"},
		{"/navigation/contactMethod/list", "GET", "获取联系方式列表"},
		{"/navigation/menuCheck/status", "GET", "检查菜单状态"},
		{"/navigation/menuCheck/create", "POST", "创建菜单"},
		{"/navigation/menuCheck/refresh", "PUT", "刷新菜单"},
	}

	apiStatus := make(map[string]bool)

	for _, api := range apis {
		var existingApi system.SysApi
		exists := !errors.Is(global.GVA_DB.Where("path = ? AND method = ?", api.Path, api.Method).First(&existingApi).Error, gorm.ErrRecordNotFound)
		apiStatus[api.Name] = exists
	}

	status["apis"] = apiStatus
	status["all_exist"] = s.CheckNavigationApisExist()

	return status
}

// GetNavigationApiList 获取导航管理API列表
func (s *NavApiCheckService) GetNavigationApiList() ([]system.SysApi, error) {
	if global.GVA_DB == nil {
		return nil, errors.New("数据库连接为空")
	}

	var apis []system.SysApi
	err := global.GVA_DB.Where("api_group = ?", "导航管理").Find(&apis).Error
	if err != nil {
		global.GVA_LOG.Error("获取导航API列表失败", zap.Error(err))
		return nil, err
	}

	return apis, nil
}
