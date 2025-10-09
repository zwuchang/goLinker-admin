package global

import (
	"sync"
	"time"

	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

type CronHistory struct {
	LastRunTime   time.Time
	LastResult    string
	Count         int
	Status        string
	Duration      time.Duration `json:",string"`
	TotalTasks    int32
	FinishedTasks int32
}

var (
	CronHistoryMap = make(map[string]*CronHistory)
	CronMutex      sync.RWMutex
	RunningTasks   = make(map[string]bool)
	RunningMutex   sync.RWMutex
)

func GetCronHistory(name string) *CronHistory {
	CronMutex.RLock()
	defer CronMutex.RUnlock()
	if history, ok := CronHistoryMap[name]; ok {
		return history
	}
	return &CronHistory{}
}

func UpdateCronHistory(name string, history *CronHistory) {
	CronMutex.Lock()
	defer CronMutex.Unlock()

	if existing, ok := CronHistoryMap[name]; ok {
		if history.Status == "running" {
			existing.FinishedTasks = 0
			existing.TotalTasks = history.TotalTasks
			existing.LastRunTime = time.Now()
			existing.Status = "running"
			existing.LastResult = "任务执行中..."
			return
		}

		if history.Count > 0 {
			existing.Count = history.Count
		}
		if !history.LastRunTime.IsZero() {
			existing.LastRunTime = history.LastRunTime
		}
		if history.LastResult != "" {
			existing.LastResult = history.LastResult
		}
		if history.Status != "" {
			existing.Status = history.Status
		}
		if history.Duration != 0 {
			existing.Duration = history.Duration
		}
		if history.TotalTasks > 0 {
			existing.TotalTasks = history.TotalTasks
		}
		if history.FinishedTasks > 0 {
			existing.FinishedTasks = history.FinishedTasks
		}
	} else {
		CronHistoryMap[name] = history
	}
}

func AddCronTask(name string, spec string, job func() error, description string) (entryID cron.EntryID, err error) {
	var option []cron.Option
	option = append(option, cron.WithSeconds())

	wrappedJob := func() {
		RunningMutex.Lock()
		GVA_LOG.Debug(" --- 开始执行任务 ---", zap.String("name", name))
		if RunningTasks[name] {
			GVA_LOG.Debug(" --- 上一个任务还没执行完 ---", zap.String("name", name))
			UpdateCronHistory(name, &CronHistory{
				LastRunTime: time.Now(),
				Status:      "skipped",
				LastResult:  "上一个任务还在执行中",
			})
			RunningMutex.Unlock()
			return
		}
		RunningTasks[name] = true
		RunningMutex.Unlock()

		startTime := time.Now()
		err := job()
		duration := time.Since(startTime).Round(time.Microsecond)

		status := "success"

		if err != nil {
			status = "failed"
		}

		current := GetCronHistory(name)
		count := current.Count + 1

		UpdateCronHistory(name, &CronHistory{
			LastRunTime: startTime,
			Count:       count,
			Duration:    duration,
			Status:      status,
		})

		// 任务完成后释放
		RunningMutex.Lock()
		delete(RunningTasks, name)
		GVA_LOG.Debug(" --- 任务执行完毕，释放锁 ---", zap.String("name", name))
		RunningMutex.Unlock()
	}

	return GVA_Timer.AddTaskByFunc(name, spec, wrappedJob, description, option...)
}
