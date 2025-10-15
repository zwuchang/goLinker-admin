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
			Path:        "/navigation/contactMethod/get",
			Description: "获取单个联系方式",
			ApiGroup:    "导航管理",
			Method:      "POST",
		},
		{
			Path:        "/navigation/contactMethod/list",
			Description: "获取联系方式列表",
			ApiGroup:    "导航管理",
			Method:      "POST",
		},
		// 游戏类别API
		{
			Path:        "/navigation/gameCategory/createGameCategory",
			Description: "创建游戏类别",
			ApiGroup:    "导航管理",
			Method:      "POST",
		},
		{
			Path:        "/navigation/gameCategory/updateGameCategory",
			Description: "更新游戏类别",
			ApiGroup:    "导航管理",
			Method:      "PUT",
		},
		{
			Path:        "/navigation/gameCategory/deleteGameCategory",
			Description: "删除游戏类别",
			ApiGroup:    "导航管理",
			Method:      "DELETE",
		},
		{
			Path:        "/navigation/gameCategory/getGameCategoryList",
			Description: "获取游戏类别列表",
			ApiGroup:    "导航管理",
			Method:      "POST",
		},
		{
			Path:        "/navigation/gameCategory/getAllGameCategories",
			Description: "获取所有启用的游戏类别",
			ApiGroup:    "导航管理",
			Method:      "GET",
		},
		// 游戏管理API
		{
			Path:        "/navigation/game/createGame",
			Description: "创建游戏",
			ApiGroup:    "导航管理",
			Method:      "POST",
		},
		{
			Path:        "/navigation/game/updateGame",
			Description: "更新游戏",
			ApiGroup:    "导航管理",
			Method:      "PUT",
		},
		{
			Path:        "/navigation/game/deleteGame",
			Description: "删除游戏",
			ApiGroup:    "导航管理",
			Method:      "DELETE",
		},
		{
			Path:        "/navigation/game/getGameList",
			Description: "获取游戏列表",
			ApiGroup:    "导航管理",
			Method:      "POST",
		},
		{
			Path:        "/navigation/game/updateGameViews",
			Description: "更新游戏浏览次数",
			ApiGroup:    "导航管理",
			Method:      "POST",
		},
		// 游戏配置API
		{
			Path:        "/navigation/gameConfig/createGameConfig",
			Description: "创建游戏配置",
			ApiGroup:    "导航管理",
			Method:      "POST",
		},
		{
			Path:        "/navigation/gameConfig/updateGameConfig",
			Description: "更新游戏配置",
			ApiGroup:    "导航管理",
			Method:      "PUT",
		},
		{
			Path:        "/navigation/gameConfig/deleteGameConfig",
			Description: "删除游戏配置",
			ApiGroup:    "导航管理",
			Method:      "DELETE",
		},
		{
			Path:        "/navigation/gameConfig/getGameConfig",
			Description: "获取游戏配置",
			ApiGroup:    "导航管理",
			Method:      "GET",
		},
		{
			Path:        "/navigation/gameConfig/getGameConfigList",
			Description: "获取游戏配置列表",
			ApiGroup:    "导航管理",
			Method:      "POST",
		},
		// Banner配置API
		{
			Path:        "/navigation/banner/getBannerList",
			Description: "获取Banner列表",
			ApiGroup:    "导航管理",
			Method:      "POST",
		},
		{
			Path:        "/navigation/banner/getBannerById",
			Description: "根据ID获取Banner",
			ApiGroup:    "导航管理",
			Method:      "GET",
		},
		{
			Path:        "/navigation/banner/createBanner",
			Description: "创建Banner",
			ApiGroup:    "导航管理",
			Method:      "POST",
		},
		{
			Path:        "/navigation/banner/updateBanner",
			Description: "更新Banner",
			ApiGroup:    "导航管理",
			Method:      "POST",
		},
		{
			Path:        "/navigation/banner/deleteBanner",
			Description: "删除Banner",
			ApiGroup:    "导航管理",
			Method:      "POST",
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
		{"/navigation/contactMethod/list", "POST"},
		{"/navigation/gameCategory/createGameCategory", "POST"},
		{"/navigation/gameCategory/getGameCategoryList", "POST"},
		{"/navigation/game/createGame", "POST"},
		{"/navigation/game/getGameList", "POST"},
		{"/navigation/gameConfig/createGameConfig", "POST"},
		{"/navigation/gameConfig/getGameConfig", "GET"},
		{"/navigation/banner/getBannerList", "POST"},
		{"/navigation/banner/createBanner", "POST"},
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
		{"/navigation/contactMethod/get", "POST", "获取单个联系方式"},
		{"/navigation/contactMethod/list", "POST", "获取联系方式列表"},
		{"/navigation/gameCategory/createGameCategory", "POST", "创建游戏类别"},
		{"/navigation/gameCategory/updateGameCategory", "PUT", "更新游戏类别"},
		{"/navigation/gameCategory/deleteGameCategory", "DELETE", "删除游戏类别"},
		{"/navigation/gameCategory/getGameCategoryList", "POST", "获取游戏类别列表"},
		{"/navigation/gameCategory/getAllGameCategories", "GET", "获取所有启用的游戏类别"},
		{"/navigation/game/createGame", "POST", "创建游戏"},
		{"/navigation/game/updateGame", "PUT", "更新游戏"},
		{"/navigation/game/deleteGame", "DELETE", "删除游戏"},
		{"/navigation/game/getGameList", "POST", "获取游戏列表"},
		{"/navigation/game/updateGameViews", "POST", "更新游戏浏览次数"},
		{"/navigation/gameConfig/createGameConfig", "POST", "创建游戏配置"},
		{"/navigation/gameConfig/updateGameConfig", "PUT", "更新游戏配置"},
		{"/navigation/gameConfig/deleteGameConfig", "DELETE", "删除游戏配置"},
		{"/navigation/gameConfig/getGameConfig", "GET", "获取游戏配置"},
		{"/navigation/gameConfig/getGameConfigList", "POST", "获取游戏配置列表"},
		{"/navigation/banner/getBannerList", "POST", "获取Banner列表"},
		{"/navigation/banner/getBannerById", "GET", "根据ID获取Banner"},
		{"/navigation/banner/createBanner", "POST", "创建Banner"},
		{"/navigation/banner/updateBanner", "POST", "更新Banner"},
		{"/navigation/banner/deleteBanner", "DELETE", "删除Banner"},
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
