package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jd-opensource/joylive-dashboard/internal/mods/resource/biz"
	"github.com/jd-opensource/joylive-dashboard/internal/mods/resource/schema"
	"github.com/jd-opensource/joylive-dashboard/pkg/util"
)

// Service management for microservices
type Service struct {
	ServiceBIZ *biz.Service
}

// @Tags ServiceAPI
// @Security ApiKeyAuth
// @Summary Query service list
// @Param current query int true "pagination index" default(1)
// @Param pageSize query int true "pagination size" default(10)
// @Param name query string false "Name of service"
// @Param space_code query string false "Space code"
// @Param role query string false "Filter by role of current user's applications: provider, consumer"
// @Success 200 {object} util.ResponseResult{data=[]schema.Service}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/services [get]
func (a *Service) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.ServiceQueryParam
	if err := util.ParseQuery(c, &params); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.ServiceBIZ.Query(ctx, params)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResPage(c, result.Data, result.PageResult)
}

// @Tags ServiceAPI
// @Security ApiKeyAuth
// @Summary Get service record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult{data=schema.Service}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/services/{id} [get]
func (a *Service) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.ServiceBIZ.Get(ctx, c.Param("id"))
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, item)
}

// @Tags ServiceAPI
// @Security ApiKeyAuth
// @Summary Create service record
// @Param body body schema.ServiceForm true "Request body"
// @Success 200 {object} util.ResponseResult{data=schema.Service}
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/services [post]
func (a *Service) Create(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(schema.ServiceForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.ServiceBIZ.Create(ctx, item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, result)
}

// @Tags ServiceAPI
// @Security ApiKeyAuth
// @Summary Update service record by ID
// @Param id path string true "unique id"
// @Param body body schema.ServiceForm true "Request body"
// @Success 200 {object} util.ResponseResult
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/services/{id} [put]
func (a *Service) Update(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(schema.ServiceForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.ResError(c, err)
		return
	}

	err := a.ServiceBIZ.Update(ctx, c.Param("id"), item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}

// @Tags ServiceAPI
// @Security ApiKeyAuth
// @Summary Delete service record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/services/{id} [delete]
func (a *Service) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.ServiceBIZ.Delete(ctx, c.Param("id"))
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}

// @Tags ServiceAPI
// @Security ApiKeyAuth
// @Summary Delete consumer relationship for service
// @Param id path string true "service id"
// @Success 200 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/services/{id}/consumer [delete]
// @Tags ServiceAPI
// @Security ApiKeyAuth
// @Summary Toggle service authorization
// @Param id path string true "service id"
// @Param body body object true "Authorization flag"
// @Success 200 {object} util.ResponseResult
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/services/{id}/auth [put]
func (a *Service) ToggleAuth(c *gin.Context) {
	ctx := c.Request.Context()
	var body struct {
		Authorized int `json:"authorized"`
	}
	if err := util.ParseJSON(c, &body); err != nil {
		util.ResError(c, err)
		return
	}

	err := a.ServiceBIZ.ToggleAuth(ctx, c.Param("id"), body.Authorized)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}

// @Tags ServiceAPI
// @Security ApiKeyAuth
// @Summary Delete consumer relationship for service
// @Param id path string true "service id"
func (a *Service) DeleteConsumer(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.ServiceBIZ.DeleteConsumer(ctx, c.Param("id"))
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}
