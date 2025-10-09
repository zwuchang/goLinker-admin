package utils

import (
	"sync"

	"goLinker-admin/server/global"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"go.uber.org/zap"
)

var (
	syncedCachedEnforcer *casbin.SyncedCachedEnforcer
	once                 sync.Once
)

// GetCasbin 获取casbin实例
func GetCasbin() *casbin.SyncedCachedEnforcer {
	once.Do(func() {
		// 检查数据库连接
		if global.GVA_DB == nil {
			zap.L().Error("数据库连接未初始化!")
			return
		}

		a, err := gormadapter.NewAdapterByDB(global.GVA_DB)
		if err != nil {
			zap.L().Error("适配数据库失败请检查casbin表是否为InnoDB引擎!", zap.Error(err))
			return
		}
		text := `
		[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act
		
		[role_definition]
		g = _, _
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = r.sub == p.sub && keyMatch2(r.obj,p.obj) && r.act == p.act
		`
		m, err := model.NewModelFromString(text)
		if err != nil {
			zap.L().Error("字符串加载模型失败!", zap.Error(err))
			return
		}
		enforcer, err := casbin.NewSyncedCachedEnforcer(m, a)
		if err != nil {
			zap.L().Error("创建Casbin执行器失败!", zap.Error(err))
			return
		}
		enforcer.SetExpireTime(60 * 60)
		if err := enforcer.LoadPolicy(); err != nil {
			zap.L().Error("加载策略失败!", zap.Error(err))
			return
		}
		syncedCachedEnforcer = enforcer
	})
	return syncedCachedEnforcer
}
