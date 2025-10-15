package navigation

import (
	"goLinker-admin/server/global"
	"goLinker-admin/server/model/common/response"
	navRequest "goLinker-admin/server/model/navigation/request"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type NavAccessStatsApi struct{}

// GetAccessStatsList 获取访问统计列表
// @Tags     AccessStatsModule
// @Summary  获取访问统计列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body navRequest.NavAccessStatsSearch true "访问统计搜索参数"
// @Success  200  {object} response.Response{data=map[string]interface{}} "获取成功"
// @Router   /navigation/accessStats/getAccessStatsList [post]
func (a *NavAccessStatsApi) GetAccessStatsList(c *gin.Context) {
	var req navRequest.NavAccessStatsSearch
	// 尝试绑定JSON，如果失败则使用默认值
	err := c.ShouldBindJSON(&req)
	if err != nil {
		// 如果JSON绑定失败，使用默认分页参数
		req.Page = 1
		req.PageSize = 10
		req.OrderBy = "created_at"
		req.OrderType = "desc"
	}

	list, total, err := navAccessStatsService.GetAccessStatsList(req)
	if err != nil {
		global.GVA_LOG.Error("获取访问统计列表失败!", zap.Error(err))
		response.FailWithMessage("获取访问统计列表失败", c)
		return
	}

	response.OkWithDetailed(map[string]interface{}{
		"list":     list,
		"total":    total,
		"page":     req.Page,
		"pageSize": req.PageSize,
	}, "获取成功", c)
}

// GetAccessStatsSummary 获取访问统计汇总
// @Tags     AccessStatsModule
// @Summary  获取访问统计汇总
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body navRequest.NavAccessStatsSummaryRequest true "访问统计汇总请求"
// @Success  200  {object} response.Response{data=navResponse.NavAccessStatsSummaryResponse} "获取成功"
// @Router   /navigation/accessStats/getAccessStatsSummary [post]
func (a *NavAccessStatsApi) GetAccessStatsSummary(c *gin.Context) {
	var req navRequest.NavAccessStatsSummaryRequest
	// 尝试绑定JSON，如果失败则使用默认值
	err := c.ShouldBindJSON(&req)
	if err != nil {
		// 如果JSON绑定失败，尝试从query参数获取
		startTime := c.Query("startTime")
		endTime := c.Query("endTime")
		req.StartTime = startTime
		req.EndTime = endTime
	}

	startTime := req.StartTime
	endTime := req.EndTime

	summary, err := navAccessStatsService.GetAccessStatsSummary(startTime, endTime)
	if err != nil {
		global.GVA_LOG.Error("获取访问统计汇总失败!", zap.Error(err))
		response.FailWithMessage("获取访问统计汇总失败", c)
		return
	}

	response.OkWithDetailed(summary, "获取成功", c)
}

// CleanOldStats 清理旧数据
// @Tags     AccessStatsModule
// @Summary  清理旧数据
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    days query int true "保留天数"
// @Success  200  {object} response.Response{msg=string} "清理成功"
// @Router   /navigation/accessStats/cleanOldStats [post]
func (a *NavAccessStatsApi) CleanOldStats(c *gin.Context) {
	daysStr := c.Query("days")
	days, err := strconv.Atoi(daysStr)
	if err != nil || days <= 0 {
		response.FailWithMessage("请输入有效的保留天数", c)
		return
	}

	err = navAccessStatsService.CleanOldStats(days)
	if err != nil {
		global.GVA_LOG.Error("清理旧数据失败!", zap.Error(err))
		response.FailWithMessage("清理旧数据失败", c)
		return
	}

	response.OkWithMessage("清理成功", c)
}
