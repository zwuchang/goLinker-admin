package initialize

import (
	"net/http"
	"os"

	"goLinker-admin/server/docs"
	"goLinker-admin/server/global"
	"goLinker-admin/server/middleware"
	"goLinker-admin/server/router"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type justFilesFilesystem struct {
	fs http.FileSystem
}

func (fs justFilesFilesystem) Open(name string) (http.File, error) {
	f, err := fs.fs.Open(name)
	if err != nil {
		return nil, err
	}

	stat, err := f.Stat()
	if stat.IsDir() {
		return nil, os.ErrPermission
	}

	return f, nil
}

// 初始化总路由

func Routers() *gin.Engine {
	Router := gin.New()
	Router.Use(gin.Recovery())
	if gin.Mode() == gin.DebugMode {
		Router.Use(gin.Logger())
	}

	systemRouter := router.RouterGroupApp.System
	exampleRouter := router.RouterGroupApp.Example
	navigationRouter := router.RouterGroupApp.Navigation
	// 如果想要不使用nginx代理前端网页，可以修改 web/.env.production 下的
	// VUE_APP_BASE_API = /
	// VUE_APP_BASE_PATH = http://localhost
	// 然后执行打包命令 npm run build。在打开下面3行注释
	// Router.StaticFile("/favicon.ico", "./dist/favicon.ico")
	// Router.Static("/assets", "./dist/assets")   // dist里面的静态资源
	if global.GVA_CONFIG.System.Env == "dev" {
		Router.StaticFile("/", "./dist/index.html") // 前端网页入口页面
		Router.Static("/assets", "./dist/assets")
		Router.StaticFile("/favicon.ico", "./dist/favicon.ico")
	}

	Router.StaticFS(global.GVA_CONFIG.Local.StorePath, justFilesFilesystem{http.Dir(global.GVA_CONFIG.Local.StorePath)}) // Router.Use(middleware.LoadTls())  // 如果需要使用https 请打开此中间件 然后前往 core/server.go 将启动模式 更变为 Router.RunTLS("端口","你的cre/pem文件","你的key文件")
	// 跨域，如需跨域可以打开下面的注释
	Router.Use(middleware.Cors()) // 直接放行全部跨域请求
	// Router.Use(middleware.CorsByRules()) // 按照配置的规则放行跨域请求
	// global.GVA_LOG.Info("use middleware cors")
	docs.SwaggerInfo.BasePath = global.GVA_CONFIG.System.RouterPrefix
	Router.GET(global.GVA_CONFIG.System.RouterPrefix+"/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.GVA_LOG.Info("register swagger handler")
	// 方便统一添加路由组前缀 多服务器上线使用

	PublicGroup := Router.Group(global.GVA_CONFIG.System.RouterPrefix)
	PrivateGroup := Router.Group(global.GVA_CONFIG.System.RouterPrefix)

	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "ok", "version": "20251009"})
		})
	}
	{
		systemRouter.InitBaseRouter(PublicGroup)      // 注册基础功能路由 不做鉴权
		systemRouter.InitInitRouter(PublicGroup)      // 自动初始化相关
		systemRouter.InitWebSocketRouter(PublicGroup) // websocket
	}

	{
		systemRouter.InitApiRouter(PrivateGroup, PublicGroup)               // 注册功能api路由
		systemRouter.InitJwtRouter(PrivateGroup)                            // jwt相关路由
		systemRouter.InitUserRouter(PrivateGroup)                           // 注册用户路由
		systemRouter.InitMenuRouter(PrivateGroup)                           // 注册menu路由
		systemRouter.InitSystemRouter(PrivateGroup)                         // system相关路由
		systemRouter.InitSysVersionRouter(PrivateGroup)                     // 发版相关路由
		systemRouter.InitCasbinRouter(PrivateGroup)                         // 权限相关路由
		systemRouter.InitAuthorityRouter(PrivateGroup)                      // 注册角色路由
		systemRouter.InitSysDictionaryRouter(PrivateGroup)                  // 字典管理
		systemRouter.InitSysOperationRecordRouter(PrivateGroup)             // 操作记录
		systemRouter.InitSysDictionaryDetailRouter(PrivateGroup)            // 字典详情管理
		systemRouter.InitAuthorityBtnRouterRouter(PrivateGroup)             // 按钮权限管理
		systemRouter.InitSysExportTemplateRouter(PrivateGroup, PublicGroup) // 导出模板
		systemRouter.InitSysParamsRouter(PrivateGroup, PublicGroup)         // 参数管理
		navigationRouter.InitContactConfigRouter(PrivateGroup)              // 导航-联系配置
		navigationRouter.InitContactMethodRouter(PrivateGroup)              // 导航-联系方式
		navigationRouter.InitGameCategoryRouter(PrivateGroup)               // 导航-游戏类别
		navigationRouter.InitGameRouter(PrivateGroup)                       // 导航-游戏管理
		navigationRouter.InitGameConfigRouter(PrivateGroup)                 // 导航-游戏配置
		navigationRouter.InitBannerRouter(PrivateGroup)                     // 导航-Banner配置
		navigationRouter.InitAccessStatsRouter(PrivateGroup)                // 导航-访问统计
		navigationRouter.InitNavPlatformConfigRouter(PrivateGroup)          // 导航-平台游戏配置
		navigationRouter.InitPublicRouter(PublicGroup)                      // 公开接口（无需认证）
		exampleRouter.InitCustomerRouter(PrivateGroup)                      // 客户路由
		exampleRouter.InitFileUploadAndDownloadRouter(PrivateGroup)         // 文件上传下载功能路由
		exampleRouter.InitAttachmentCategoryRouterRouter(PrivateGroup)      // 文件上传下载分类
		systemRouter.InitCronRouter(PrivateGroup)                           // 定时任务管理

	}

	// 注册业务路由
	initBizRouter(PrivateGroup, PublicGroup)

	global.GVA_ROUTERS = Router.Routes()

	global.GVA_LOG.Info("router register success")
	return Router
}
