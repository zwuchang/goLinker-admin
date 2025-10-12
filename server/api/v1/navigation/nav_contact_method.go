package navigation

import (
	"goLinker-admin/server/global"
	"goLinker-admin/server/model/common/request"
	"goLinker-admin/server/model/common/response"
	"goLinker-admin/server/model/navigation"
	navRes "goLinker-admin/server/model/navigation/response"
	"goLinker-admin/server/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ContactMethodApi struct{}

// CreateNavContactMethod 创建联系方式
// @Tags      NavContactMethod
// @Summary   创建联系方式
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      navigation.NavContactMethod            true  "联系方式信息"
// @Success   200   {object}  response.Response{msg=string}       "创建成功"
// @Router    /navigation/contactMethod/create [post]
func (s *ContactMethodApi) CreateNavContactMethod(c *gin.Context) {
	var method navigation.NavContactMethod
	err := c.ShouldBindJSON(&method)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = contactMethodService.CreateNavContactMethod(method)
	if err != nil {
		global.GVA_LOG.Error("创建联系方式失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteNavContactMethod 删除联系方式
// @Tags      NavContactMethod
// @Summary   删除联系方式
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      navigation.NavContactMethod            true  "联系方式ID"
// @Success   200   {object}  response.Response{msg=string}       "删除成功"
// @Router    /navigation/contactMethod/delete [delete]
func (s *ContactMethodApi) DeleteNavContactMethod(c *gin.Context) {
	var method navigation.NavContactMethod
	err := c.ShouldBindJSON(&method)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(method.GVA_MODEL, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = contactMethodService.DeleteNavContactMethod(method)
	if err != nil {
		global.GVA_LOG.Error("删除联系方式失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// UpdateNavContactMethod 更新联系方式
// @Tags      NavContactMethod
// @Summary   更新联系方式
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      navigation.NavContactMethod            true  "联系方式信息"
// @Success   200   {object}  response.Response{msg=string}       "更新成功"
// @Router    /navigation/contactMethod/update [put]
func (s *ContactMethodApi) UpdateNavContactMethod(c *gin.Context) {
	var method navigation.NavContactMethod
	err := c.ShouldBindJSON(&method)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(method.GVA_MODEL, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = contactMethodService.UpdateNavContactMethod(&method)
	if err != nil {
		global.GVA_LOG.Error("更新联系方式失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// GetNavContactMethod 获取单个联系方式
// @Tags      NavContactMethod
// @Summary   获取单个联系方式
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     navigation.NavContactMethod                                                true  "联系方式ID"
// @Success   200   {object}  response.Response{data=navRes.NavContactMethodResponse,msg=string}  "获取成功"
// @Router    /navigation/contactMethod/get [get]
func (s *ContactMethodApi) GetNavContactMethod(c *gin.Context) {
	var method navigation.NavContactMethod
	err := c.ShouldBindQuery(&method)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(method.GVA_MODEL, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	data, err := contactMethodService.GetNavContactMethod(method.ID)
	if err != nil {
		global.GVA_LOG.Error("获取联系方式失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(navRes.NavContactMethodResponse{ContactMethod: data}, "获取成功", c)
}

// GetNavContactMethodList 分页获取联系方式列表
// @Tags      NavContactMethod
// @Summary   分页获取联系方式列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.PageInfo                                        true  "页码, 每页大小"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "获取成功"
// @Router    /navigation/contactMethod/list [get]
func (s *ContactMethodApi) GetNavContactMethodList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(pageInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	methodList, total, err := contactMethodService.GetNavContactMethodList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取联系方式列表失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     methodList,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}
