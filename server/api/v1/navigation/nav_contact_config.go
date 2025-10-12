package navigation

import (
	"goLinker-admin/server/global"
	"goLinker-admin/server/model/common/response"
	"goLinker-admin/server/model/navigation"
	navRes "goLinker-admin/server/model/navigation/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ContactConfigApi struct{}

// GetNavContactConfig 获取联系配置
// @Tags      NavContactConfig
// @Summary   获取联系配置
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200   {object}  response.Response{data=navRes.NavContactConfigResponse,msg=string}  "获取成功"
// @Router    /navigation/contactConfig/get [get]
func (s *ContactConfigApi) GetNavContactConfig(c *gin.Context) {
	data, err := contactConfigService.GetNavContactConfig()
	if err != nil {
		global.GVA_LOG.Error("获取联系配置失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(navRes.NavContactConfigResponse{ContactConfig: data}, "获取成功", c)
}

// UpdateNavContactConfig 更新联系配置
// @Tags      NavContactConfig
// @Summary   更新联系配置
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      navigation.NavContactConfig            true  "联系配置"
// @Success   200   {object}  response.Response{msg=string}       "更新成功"
// @Router    /navigation/contactConfig/update [put]
func (s *ContactConfigApi) UpdateNavContactConfig(c *gin.Context) {
	var config navigation.NavContactConfig
	err := c.ShouldBindJSON(&config)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = contactConfigService.UpdateNavContactConfig(&config)
	if err != nil {
		global.GVA_LOG.Error("更新联系配置失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}
