package system

import (
	"fmt"

	"sort"
	"strings"
	"time"

	"goLinker-admin/server/global"
	"goLinker-admin/server/model/system"
	"goLinker-admin/server/model/system/request"
	"goLinker-admin/server/task"
	"goLinker-admin/server/utils"

	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

type CronService struct{}

// GetAllCronStatus 获取所有定时任务状态
func (s *CronService) GetAllCronStatus() []request.CronInfo {
	var infos []request.CronInfo

	// 获取所有定时任务
	cronList := global.GVA_Timer.FindCronList()
	for cronName, cronManager := range cronList {
		if cronManager == nil {
			continue
		}

		entries := cronManager.GetEntries()
		if len(entries) == 0 {
			continue
		}

		// 获取第一个任务的信息
		entry := entries[0]
		findTask := cronManager.GetTask(entry.ID)
		if findTask == nil {
			continue
		}

		// 获取历史记录
		history := global.GetCronHistory(cronName)

		info := request.CronInfo{
			Name:         cronName,
			Description:  cronManager.GetDescription(entry.ID),
			Spec:         findTask.Spec,
			ExecuteCount: history.Count,
			LastRunTime:  history.LastRunTime,
			NextRunTime:  entry.Next,
			IsRunning:    cronManager.IsRunning(),
			LastResult:   history.LastResult,
			Status:       history.Status,
			Duration:     utils.FormatDuration(history.Duration),
		}

		if !history.LastRunTime.IsZero() {
			info.TimeSinceLast = time.Since(history.LastRunTime).Round(time.Second).String()
		}

		if info.IsRunning {
			info.NextRunTime = entry.Next
			if !entry.Next.IsZero() {
				info.TimeUntilNext = time.Until(entry.Next).Round(time.Second).String()
			}
		}

		infos = append(infos, info)
	}

	// 按任务名称排序
	sort.Slice(infos, func(i, j int) bool {
		return infos[i].Name < infos[j].Name
	})

	return infos
}

// UpdateCronSpec 更新定时任务执行规则
func (s *CronService) UpdateCronSpec(req request.UpdateCronSpecReq) error {
	// 先验证 spec 格式是否正确
	if strings.HasPrefix(req.Spec, "@") {
		// 处理预定义描述符
		switch req.Spec {
		case "@yearly", "@annually", "@monthly", "@weekly", "@daily", "@midnight", "@hourly":
			// 这些是有效的预定义描述符
		case "@every":
			return fmt.Errorf("@every 必须指定时间间隔，例如：@every 1h")
		default:
			if strings.HasPrefix(req.Spec, "@every ") {
				// 解析 @every 格式
				_, err := time.ParseDuration(strings.TrimPrefix(req.Spec, "@every "))
				if err != nil {
					return fmt.Errorf("@every 格式错误: %v", err)
				}
			} else {
				return fmt.Errorf("不支持的预定义描述符: %s", req.Spec)
			}
		}
	} else {
		// 解析标准 cron 表达式
		parser := cron.NewParser(
			cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow,
		)
		_, err := parser.Parse(req.Spec)
		if err != nil {
			return fmt.Errorf("定时任务表达式格式错误: %v", err)
		}
	}

	findCron, found := global.GVA_Timer.FindCron(req.Name)
	if !found || findCron == nil {
		return fmt.Errorf("任务 %s 不存在", req.Name)
	}

	entries := findCron.GetEntries()
	if len(entries) == 0 {
		return fmt.Errorf("任务 %s 没有活动的条目", req.Name)
	}

	// 获取原任务信息
	oldEntry := entries[0]
	oldTask := findCron.GetTask(oldEntry.ID)
	if oldTask == nil {
		return fmt.Errorf("无法获取任务信息")
	}

	// 保存原任务的执行历史
	oldHistory := global.GetCronHistory(req.Name)

	// 停止并清理旧任务
	global.GVA_Timer.StopCron(req.Name)
	global.GVA_Timer.Clear(req.Name)

	// 根据任务名称获取对应的任务函数
	var taskFunc func() error
	switch req.Name {
	case global.CronClearDbName:
		taskFunc = func() error {
			return task.ClearTable(global.GVA_DB)
		}

	default:
		return fmt.Errorf("未知的任务类型: %s", req.Name)
	}

	// 添加新任务
	_, err := global.AddCronTask(req.Name, req.Spec, taskFunc, oldTask.TaskName)
	if err != nil {
		// 如果失败，恢复旧任务
		_, _ = global.AddCronTask(req.Name, oldTask.Spec, taskFunc, oldTask.TaskName)
		return fmt.Errorf("更新任务失败: %v", err)
	}

	// 恢复历史记录
	global.UpdateCronHistory(req.Name, oldHistory)

	// 更新数据库中的配置
	err = global.GVA_DB.Where("name = ?", req.Name).
		Assign(system.SysCron{
			Spec: req.Spec,
		}).
		FirstOrCreate(&system.SysCron{
			Name:        req.Name,
			Spec:        req.Spec,
			Description: oldTask.TaskName,
		}).Error
	if err != nil {
		return fmt.Errorf("更新数据库配置失败: %v", err)
	}

	return nil
}

func (s *CronService) StopCron(req request.CronNameReq) error {
	findCron, found := global.GVA_Timer.FindCron(req.Name)
	if !found || findCron == nil {
		return fmt.Errorf("任务 %s 不存在", req.Name)
	}

	global.GVA_LOG.Info("停止定时任务", zap.String("name", req.Name))
	global.GVA_Timer.StopCron(req.Name)

	// 验证状态
	time.Sleep(100 * time.Millisecond) // 等待状态更新
	if findCron.IsRunning() {
		global.GVA_LOG.Warn("任务停止后仍在运行", zap.String("name", req.Name))
	} else {
		global.GVA_LOG.Info("任务已成功停止", zap.String("name", req.Name))
	}
	return nil
}

func (s *CronService) StartCron(req request.CronNameReq) error {
	findCron, found := global.GVA_Timer.FindCron(req.Name)
	if !found || findCron == nil {
		return fmt.Errorf("任务 %s 不存在", req.Name)
	}
	global.GVA_Timer.StartCron(req.Name)
	return nil
}
