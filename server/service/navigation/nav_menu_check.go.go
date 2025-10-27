package navigation

import (
	"errors"
	"goLinker-admin/server/global"
	"goLinker-admin/server/model/system"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type NavMenuCheckService struct{}

// CheckAndCreateNavigationMenus 检查并创建导航管理相关菜单
func (s *NavMenuCheckService) CheckAndCreateNavigationMenus() error {
	global.GVA_LOG.Info("开始检查并创建导航管理菜单...")

	// 定义导航管理相关菜单
	navigationMenus := []system.SysBaseMenu{
		// 导航管理主菜单
		{
			MenuLevel: 0,
			Hidden:    false,
			ParentId:  0,
			Path:      "navigation",
			Name:      "navigation",
			Component: "view/navigation/index.vue",
			Sort:      6,
			Meta: system.Meta{
				Title: "导航管理",
				Icon:  "guide",
			},
		},
		// 联系配置子菜单
		{
			MenuLevel: 1,
			Hidden:    false,
			ParentId:  0, // 稍后会更新为实际的父菜单ID
			Path:      "contactConfig",
			Name:      "contactConfig",
			Component: "view/navigation/contactConfig/contactConfig.vue",
			Sort:      1,
			Meta: system.Meta{
				Title: "联系配置",
				Icon:  "setting",
			},
		},
		// 联系方式子菜单
		{
			MenuLevel: 1,
			Hidden:    false,
			ParentId:  0, // 稍后会更新为实际的父菜单ID
			Path:      "contactMethod",
			Name:      "contactMethod",
			Component: "view/navigation/contactMethod/contactMethod.vue",
			Sort:      2,
			Meta: system.Meta{
				Title: "联系方式",
				Icon:  "service",
			},
		},
		// 游戏类别子菜单
		{
			MenuLevel: 1,
			Hidden:    false,
			ParentId:  0, // 稍后会更新为实际的父菜单ID
			Path:      "gameCategory",
			Name:      "gameCategory",
			Component: "view/navigation/gameCategory/gameCategory.vue",
			Sort:      3,
			Meta: system.Meta{
				Title: "游戏类别",
				Icon:  "collection",
			},
		},
		// 游戏管理子菜单
		{
			MenuLevel: 1,
			Hidden:    false,
			ParentId:  0, // 稍后会更新为实际的父菜单ID
			Path:      "game",
			Name:      "game",
			Component: "view/navigation/game/game.vue",
			Sort:      4,
			Meta: system.Meta{
				Title: "游戏管理",
				Icon:  "trophy",
			},
		},
		// 游戏配置子菜单
		{
			MenuLevel: 1,
			Hidden:    false,
			ParentId:  0, // 稍后会更新为实际的父菜单ID
			Path:      "gameConfig",
			Name:      "gameConfig",
			Component: "view/navigation/gameConfig/gameConfig.vue",
			Sort:      5,
			Meta: system.Meta{
				Title: "游戏配置",
				Icon:  "setting",
			},
		},
		// Banner配置子菜单
		{
			MenuLevel: 1,
			Hidden:    false,
			ParentId:  0, // 稍后会更新为实际的父菜单ID
			Path:      "banner",
			Name:      "banner",
			Component: "view/navigation/banner/banner.vue",
			Sort:      6,
			Meta: system.Meta{
				Title: "Banner配置",
				Icon:  "picture",
			},
		},
		// 平台配置子菜单
		{
			MenuLevel: 1,
			Hidden:    false,
			ParentId:  0, // 稍后会更新为实际的父菜单ID
			Path:      "platformConfig",
			Name:      "platformConfig",
			Component: "view/navigation/platformConfig/platformConfig.vue",
			Sort:      7,
			Meta: system.Meta{
				Title: "平台配置",
				Icon:  "platform",
			},
		},
		// 市场配置子菜单
		{
			MenuLevel: 1,
			Hidden:    false,
			ParentId:  0, // 稍后会更新为实际的父菜单ID
			Path:      "marketConfig",
			Name:      "marketConfig",
			Component: "view/navigation/marketConfig/marketConfig.vue",
			Sort:      8,
			Meta: system.Meta{
				Title: "市场配置",
				Icon:  "shopping",
			},
		},
		// 平台排行榜子菜单
		{
			MenuLevel: 1,
			Hidden:    false,
			ParentId:  0, // 稍后会更新为实际的父菜单ID
			Path:      "platformRanking",
			Name:      "platformRanking",
			Component: "view/navigation/platformRanking/platformRanking.vue",
			Sort:      9,
			Meta: system.Meta{
				Title: "平台排行榜",
				Icon:  "trophy",
			},
		},
		// 滚动通知子菜单
		{
			MenuLevel: 1,
			Hidden:    false,
			ParentId:  0, // 稍后会更新为实际的父菜单ID
			Path:      "scrollNotice",
			Name:      "scrollNotice",
			Component: "view/navigation/scrollNotice/scrollNotice.vue",
			Sort:      10,
			Meta: system.Meta{
				Title: "滚动通知",
				Icon:  "bell",
			},
		},
		// PWA配置子菜单
		{
			MenuLevel: 1,
			Hidden:    false,
			ParentId:  0, // 稍后会更新为实际的父菜单ID
			Path:      "pwaConfig",
			Name:      "pwaConfig",
			Component: "view/navigation/pwaConfig/pwaConfig.vue",
			Sort:      11,
			Meta: system.Meta{
				Title: "PWA配置",
				Icon:  "mobile",
			},
		},
		// 主题配置子菜单
		{
			MenuLevel: 1,
			Hidden:    false,
			ParentId:  0, // 稍后会更新为实际的父菜单ID
			Path:      "themeConfig",
			Name:      "themeConfig",
			Component: "view/navigation/themeConfig/themeConfig.vue",
			Sort:      12,
			Meta: system.Meta{
				Title: "主题配置",
				Icon:  "brush",
			},
		},
		// 活动配置子菜单
		{
			MenuLevel: 1,
			Hidden:    false,
			ParentId:  0, // 稍后会更新为实际的父菜单ID
			Path:      "activityConfig",
			Name:      "activityConfig",
			Component: "view/navigation/activityConfig/activityConfig.vue",
			Sort:      13,
			Meta: system.Meta{
				Title: "活动配置",
				Icon:  "tickets",
			},
		},
	}

	// 检查并创建主菜单
	var navigationMenuID uint
	for _, menu := range navigationMenus {
		if menu.MenuLevel == 0 {
			// 主菜单
			createdMenu, err := s.checkAndCreateMenu(menu, 0)
			if err != nil {
				return err
			}
			navigationMenuID = createdMenu.ID
			global.GVA_LOG.Info("导航管理主菜单检查完成", zap.Uint("id", navigationMenuID))
		}
	}

	// 检查并创建子菜单
	for _, menu := range navigationMenus {
		if menu.MenuLevel == 1 {
			// 子菜单，设置正确的父菜单ID
			menu.ParentId = navigationMenuID
			_, err := s.checkAndCreateMenu(menu, navigationMenuID)
			if err != nil {
				return err
			}
		}
	}

	global.GVA_LOG.Info("导航管理菜单检查完成")
	return nil
}

// checkAndCreateMenu 检查并创建或更新单个菜单
func (s *NavMenuCheckService) checkAndCreateMenu(menu system.SysBaseMenu, parentId uint) (system.SysBaseMenu, error) {
	var existingMenu system.SysBaseMenu
	err := global.GVA_DB.Where("path = ?", menu.Path).First(&existingMenu).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 菜单不存在，创建新菜单
		menu.ParentId = parentId
		if err := global.GVA_DB.Create(&menu).Error; err != nil {
			global.GVA_LOG.Error("创建菜单失败", zap.String("path", menu.Path), zap.Error(err))
			return menu, err
		}
		global.GVA_LOG.Info("创建菜单成功", zap.String("path", menu.Path), zap.String("title", menu.Meta.Title))
		return menu, nil
	} else if err != nil {
		global.GVA_LOG.Error("查询菜单失败", zap.String("path", menu.Path), zap.Error(err))
		return existingMenu, err
	} else {
		// 菜单存在，检查是否需要更新
		needUpdate := false

		// 检查各个字段是否需要更新
		if existingMenu.Meta.Title != menu.Meta.Title {
			existingMenu.Meta.Title = menu.Meta.Title
			needUpdate = true
		}
		if existingMenu.ParentId != parentId {
			existingMenu.ParentId = parentId
			needUpdate = true
		}
		if existingMenu.Sort != menu.Sort {
			existingMenu.Sort = menu.Sort
			needUpdate = true
		}
		if existingMenu.Component != menu.Component {
			existingMenu.Component = menu.Component
			needUpdate = true
		}
		if existingMenu.Hidden != menu.Hidden {
			existingMenu.Hidden = menu.Hidden
			needUpdate = true
		}
		if existingMenu.Meta.Icon != menu.Meta.Icon {
			existingMenu.Meta.Icon = menu.Meta.Icon
			needUpdate = true
		}

		if needUpdate {
			if err := global.GVA_DB.Save(&existingMenu).Error; err != nil {
				global.GVA_LOG.Error("更新菜单失败", zap.String("path", menu.Path), zap.Error(err))
				return existingMenu, err
			}
			global.GVA_LOG.Info("更新菜单成功", zap.String("path", menu.Path), zap.String("title", menu.Meta.Title))
		} else {
			global.GVA_LOG.Debug("菜单已存在且无需更新", zap.String("path", menu.Path))
		}
		return existingMenu, nil
	}
}

// CheckNavigationMenusExist 检查导航管理菜单是否存在
func (s *NavMenuCheckService) CheckNavigationMenusExist() bool {
	if global.GVA_DB == nil {
		return false
	}

	// 检查关键菜单是否存在
	keyMenus := []string{"navigation", "contactConfig", "contactMethod", "gameCategory", "game", "gameConfig", "banner", "platformConfig", "marketConfig", "platformRanking", "scrollNotice", "pwaConfig", "themeConfig", "activityConfig"}
	for _, menuPath := range keyMenus {
		var menu system.SysBaseMenu
		if errors.Is(global.GVA_DB.Where("path = ?", menuPath).First(&menu).Error, gorm.ErrRecordNotFound) {
			return false
		}
	}
	return true
}

// GetNavigationMenuStatus 获取导航菜单状态信息
func (s *NavMenuCheckService) GetNavigationMenuStatus() map[string]interface{} {
	status := make(map[string]interface{})

	if global.GVA_DB == nil {
		status["error"] = "数据库连接为空"
		return status
	}

	// 检查各个菜单的状态
	menus := []string{"navigation", "contactConfig", "contactMethod", "gameCategory", "game", "gameConfig", "banner", "platformConfig", "marketConfig", "platformRanking", "scrollNotice", "pwaConfig", "themeConfig"}
	menuStatus := make(map[string]bool)

	for _, menuPath := range menus {
		var menu system.SysBaseMenu
		exists := !errors.Is(global.GVA_DB.Where("path = ?", menuPath).First(&menu).Error, gorm.ErrRecordNotFound)
		menuStatus[menuPath] = exists
	}

	status["menus"] = menuStatus
	status["all_exist"] = s.CheckNavigationMenusExist()

	return status
}
