package api

import (
	"github.com/jd-opensource/joylive-dashboard/internal/mods/resource/biz"
	"github.com/jd-opensource/joylive-dashboard/internal/mods/resource/schema"
	"github.com/jd-opensource/joylive-dashboard/pkg/util"
	"github.com/gin-gonic/gin"
)

// Application management
type Application struct {
	ApplicationBIZ *biz.Application
}

// @Tags ApplicationAPI
// @Security ApiKeyAuth
// @Summary Query application list
// @Param current query int true "pagination index" default(1)
// @Param pageSize query int true "pagination size" default(10)
// @Param name query string false "Name of application"
// @Success 200 {object} util.ResponseResult{data=[]schema.Application}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/applications [get]
func (a *Application) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.ApplicationQueryParam
	if err := util.ParseQuery(c, &params); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.ApplicationBIZ.Query(ctx, params)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResPage(c, result.Data, result.PageResult)
}

// @Tags ApplicationAPI
// @Security ApiKeyAuth
// @Summary Get application record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult{data=schema.Application}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/applications/{id} [get]
func (a *Application) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.ApplicationBIZ.Get(ctx, c.Param("id"))
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, item)
}

// @Tags ApplicationAPI
// @Security ApiKeyAuth
// @Summary Create application record
// @Param body body schema.ApplicationForm true "Request body"
// @Success 200 {object} util.ResponseResult{data=schema.Application}
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/applications [post]
func (a *Application) Create(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(schema.ApplicationForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.ApplicationBIZ.Create(ctx, item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, result)
}

// @Tags ApplicationAPI
// @Security ApiKeyAuth
// @Summary Update application record by ID
// @Param id path string true "unique id"
// @Param body body schema.ApplicationForm true "Request body"
// @Success 200 {object} util.ResponseResult
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/applications/{id} [put]
func (a *Application) Update(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(schema.ApplicationForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.ResError(c, err)
		return
	}

	err := a.ApplicationBIZ.Update(ctx, c.Param("id"), item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}

// @Tags ApplicationAPI
// @Security ApiKeyAuth
// @Summary Delete application record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/applications/{id} [delete]
func (a *Application) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.ApplicationBIZ.Delete(ctx, c.Param("id"))
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}
