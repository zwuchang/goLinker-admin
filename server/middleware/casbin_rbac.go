package middleware

import (
	"strconv"
	"strings"

	"goLinker-admin/server/global"
	"goLinker-admin/server/model/common/response"
	"goLinker-admin/server/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// CasbinHandler 拦截器
func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		waitUse, _ := utils.GetClaims(c)
		//获取请求的PATH
		path := c.Request.URL.Path
		obj := strings.TrimPrefix(path, global.GVA_CONFIG.System.RouterPrefix)
		// 获取请求方法
		act := c.Request.Method
		// 获取用户的角色
		sub := strconv.Itoa(int(waitUse.AuthorityId))
		e := utils.GetCasbin() // 判断策略中是否存在

		// 检查Casbin实例是否为空
		if e == nil {
			global.GVA_LOG.Error("Casbin实例未初始化，跳过权限验证")
			c.Next()
			return
		}

		success, err := e.Enforce(sub, obj, act)
		if err != nil {
			global.GVA_LOG.Error("权限验证失败", zap.Error(err))
			response.FailWithDetailed(gin.H{}, "权限验证失败", c)
			c.Abort()
			return
		}

		if !success {
			response.FailWithDetailed(gin.H{}, "权限不足", c)
			c.Abort()
			return
		}
		c.Next()
	}
}
