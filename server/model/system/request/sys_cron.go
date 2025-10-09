package request

import "time"

// CronInfo 定时任务信息
type CronInfo struct {
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	Spec          string    `json:"spec"`
	ExecuteCount  int       `json:"executeCount"`
	LastRunTime   time.Time `json:"lastRunTime"`
	NextRunTime   time.Time `json:"nextRunTime"`
	IsRunning     bool      `json:"isRunning"`
	TimeSinceLast string    `json:"timeSinceLast"`
	TimeUntilNext string    `json:"timeUntilNext"`
	LastResult    string    `json:"lastResult"`
	Status        string    `json:"status"`
	Duration      string    `json:"duration"`
}

// UpdateCronSpecReq 更新定时任务间隔请求
type UpdateCronSpecReq struct {
	Name        string `json:"name"` // 任务名称
	Spec        string `json:"spec"` // 新的执行规则
	Description string `json:"description"`
}

type CronNameReq struct {
	Name string `json:"name"`
}
