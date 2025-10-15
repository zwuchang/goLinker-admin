package initialize

import (
	"goLinker-admin/server/global"
	"os"
	"sync"

	"github.com/zhengjianyang/goCzdb"
	"go.uber.org/zap"
)

var (
	searcherMux sync.Mutex
	initialized bool
)

func InitIpSearcher() error {
	searcherMux.Lock()
	defer searcherMux.Unlock()

	if initialized {
		return nil
	}

	dbPath := global.GVA_CONFIG.System.IpDbPath
	if dbPath == "" {
		global.GVA_LOG.Warn("IP数据库路径未配置，跳过IP地理位置功能")
		initialized = true
		return nil
	}

	// 检查数据库文件是否存在
	fileInfo, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		global.GVA_LOG.Warn("IP数据库文件不存在，跳过IP地理位置功能",
			zap.String("path", dbPath))
		initialized = true
		return nil
	}
	if err != nil {
		global.GVA_LOG.Warn("检查IP数据库文件失败，跳过IP地理位置功能",
			zap.String("path", dbPath),
			zap.Error(err))
		initialized = true
		return nil
	}

	global.GVA_LOG.Info("开始初始化IP地址库",
		zap.String("path", dbPath),
		zap.Int64("size", fileInfo.Size()))

	// 使用MEMORY模式，它是线程安全的
	global.IpSearcher, err = goCzdb.NewDbSearcher(dbPath, "MEMORY", "IcDpY/+as/2WOoAdDe3gvg==")
	if err != nil {
		global.GVA_LOG.Warn("初始化IP地址库失败，跳过IP地理位置功能",
			zap.String("path", dbPath),
			zap.Error(err))
		initialized = true
		return nil
	}

	global.GVA_LOG.Info("IP地址库初始化成功",
		zap.String("path", dbPath),
		zap.Int64("size", fileInfo.Size()))
	initialized = true
	return nil
}

// CloseIPSearcher 关闭IP搜索器，在应用退出时调用
func CloseIPSearcher() {
	searcherMux.Lock()
	defer searcherMux.Unlock()

	if global.IpSearcher != nil {
		global.IpSearcher.Close()
		initialized = false
	}
}
