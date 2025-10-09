package system

import (
	"goLinker-admin/server/global"
	"goLinker-admin/server/model/common/response"
	"goLinker-admin/server/model/system/request"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CronApi struct{}

// GetCronStatus 获取所有定时任务状态
func (s *CronApi) GetCronStatus(c *gin.Context) {
	status := cronService.GetAllCronStatus()
	response.OkWithData(status, c)
}

// UpdateCronSpec 更新定时任务执行规则
func (s *CronApi) UpdateCronSpec(c *gin.Context) {
	var req request.UpdateCronSpecReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := cronService.UpdateCronSpec(req); err != nil {
		global.GVA_LOG.Error("更新定时任务规则失败", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.Ok(c)
}

// StopCron 暂停定时任务
func (s *CronApi) StopCron(c *gin.Context) {
	var req request.CronNameReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := cronService.StopCron(req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithMessage("暂停成功", c)
}

// StartCron 启动定时任务
func (s *CronApi) StartCron(c *gin.Context) {
	var req request.CronNameReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := cronService.StartCron(req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithMessage("启动成功", c)
}
