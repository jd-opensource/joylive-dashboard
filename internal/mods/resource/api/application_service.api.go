package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jd-opensource/joylive-dashboard/internal/mods/resource/biz"
	"github.com/jd-opensource/joylive-dashboard/internal/mods/resource/schema"
	"github.com/jd-opensource/joylive-dashboard/pkg/util"
)

// Application and service relationship management
type ApplicationService struct {
	ApplicationServiceBIZ *biz.ApplicationService
}

// @Tags ApplicationServiceAPI
// @Security ApiKeyAuth
// @Summary Query application service list
// @Success 200 {object} util.ResponseResult{data=[]schema.ApplicationService}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/resource/application-services [get]
func (a *ApplicationService) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.ApplicationServiceQueryParam
	if err := util.ParseQuery(c, &params); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.ApplicationServiceBIZ.Query(ctx, params)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResPage(c, result.Data, result.PageResult)
}

// @Tags ApplicationServiceAPI
// @Security ApiKeyAuth
// @Summary Get application service record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult{data=schema.ApplicationService}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/resource/application-services/{id} [get]
func (a *ApplicationService) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.ApplicationServiceBIZ.Get(ctx, c.Param("id"))
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, item)
}

// @Tags ApplicationServiceAPI
// @Security ApiKeyAuth
// @Summary Create application service record
// @Param body body schema.ApplicationServiceForm true "Request body"
// @Success 200 {object} util.ResponseResult{data=schema.ApplicationService}
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/resource/application-services [post]
func (a *ApplicationService) Create(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(schema.ApplicationServiceForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.ApplicationServiceBIZ.Create(ctx, item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, result)
}

// @Tags ApplicationServiceAPI
// @Security ApiKeyAuth
// @Summary Update application service record by ID
// @Param id path string true "unique id"
// @Param body body schema.ApplicationServiceForm true "Request body"
// @Success 200 {object} util.ResponseResult
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/resource/application-services/{id} [put]
func (a *ApplicationService) Update(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(schema.ApplicationServiceForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.ResError(c, err)
		return
	}

	err := a.ApplicationServiceBIZ.Update(ctx, c.Param("id"), item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}

// @Tags ApplicationServiceAPI
// @Security ApiKeyAuth
// @Summary Delete application service record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/resource/application-services/{id} [delete]
func (a *ApplicationService) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.ApplicationServiceBIZ.Delete(ctx, c.Param("id"))
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}
