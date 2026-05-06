package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jd-opensource/joylive-dashboard/internal/mods/resource/biz"
	"github.com/jd-opensource/joylive-dashboard/internal/mods/resource/schema"
	"github.com/jd-opensource/joylive-dashboard/pkg/util"
)

// Service group management
type ServiceGroup struct {
	ServiceGroupBIZ *biz.ServiceGroup
}

// @Tags ServiceGroupAPI
// @Security ApiKeyAuth
// @Summary Query service group list
// @Success 200 {object} util.ResponseResult{data=[]schema.ServiceGroup}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/resource/service-groups [get]
func (a *ServiceGroup) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.ServiceGroupQueryParam
	if err := util.ParseQuery(c, &params); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.ServiceGroupBIZ.Query(ctx, params)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResPage(c, result.Data, result.PageResult)
}

// @Tags ServiceGroupAPI
// @Security ApiKeyAuth
// @Summary Get service group record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult{data=schema.ServiceGroup}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/resource/service-groups/{id} [get]
func (a *ServiceGroup) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.ServiceGroupBIZ.Get(ctx, c.Param("id"))
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, item)
}

// @Tags ServiceGroupAPI
// @Security ApiKeyAuth
// @Summary Create service group record
// @Param body body schema.ServiceGroupForm true "Request body"
// @Success 200 {object} util.ResponseResult{data=schema.ServiceGroup}
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/resource/service-groups [post]
func (a *ServiceGroup) Create(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(schema.ServiceGroupForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.ServiceGroupBIZ.Create(ctx, item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, result)
}

// @Tags ServiceGroupAPI
// @Security ApiKeyAuth
// @Summary Update service group record by ID
// @Param id path string true "unique id"
// @Param body body schema.ServiceGroupForm true "Request body"
// @Success 200 {object} util.ResponseResult
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/resource/service-groups/{id} [put]
func (a *ServiceGroup) Update(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(schema.ServiceGroupForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.ResError(c, err)
		return
	}

	err := a.ServiceGroupBIZ.Update(ctx, c.Param("id"), item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}

// @Tags ServiceGroupAPI
// @Security ApiKeyAuth
// @Summary Delete service group record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/resource/service-groups/{id} [delete]
func (a *ServiceGroup) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.ServiceGroupBIZ.Delete(ctx, c.Param("id"))
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}
