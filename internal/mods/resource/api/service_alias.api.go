package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jd-opensource/joylive-dashboard/internal/mods/resource/biz"
	"github.com/jd-opensource/joylive-dashboard/internal/mods/resource/schema"
	"github.com/jd-opensource/joylive-dashboard/pkg/util"
)

// Service alias management
type ServiceAlias struct {
	ServiceAliasBIZ *biz.ServiceAlias
}

// @Tags ServiceAliasAPI
// @Security ApiKeyAuth
// @Summary Query service alias list
// @Success 200 {object} util.ResponseResult{data=[]schema.ServiceAlias}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/resource/service-aliases [get]
func (a *ServiceAlias) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.ServiceAliasQueryParam
	if err := util.ParseQuery(c, &params); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.ServiceAliasBIZ.Query(ctx, params)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResPage(c, result.Data, result.PageResult)
}

// @Tags ServiceAliasAPI
// @Security ApiKeyAuth
// @Summary Get service alias record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult{data=schema.ServiceAlias}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/resource/service-aliases/{id} [get]
func (a *ServiceAlias) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.ServiceAliasBIZ.Get(ctx, c.Param("id"))
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, item)
}

// @Tags ServiceAliasAPI
// @Security ApiKeyAuth
// @Summary Create service alias record
// @Param body body schema.ServiceAliasForm true "Request body"
// @Success 200 {object} util.ResponseResult{data=schema.ServiceAlias}
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/resource/service-aliases [post]
func (a *ServiceAlias) Create(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(schema.ServiceAliasForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.ServiceAliasBIZ.Create(ctx, item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, result)
}

// @Tags ServiceAliasAPI
// @Security ApiKeyAuth
// @Summary Update service alias record by ID
// @Param id path string true "unique id"
// @Param body body schema.ServiceAliasForm true "Request body"
// @Success 200 {object} util.ResponseResult
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/resource/service-aliases/{id} [put]
func (a *ServiceAlias) Update(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(schema.ServiceAliasForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.ResError(c, err)
		return
	}

	err := a.ServiceAliasBIZ.Update(ctx, c.Param("id"), item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}

// @Tags ServiceAliasAPI
// @Security ApiKeyAuth
// @Summary Delete service alias record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/resource/service-aliases/{id} [delete]
func (a *ServiceAlias) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.ServiceAliasBIZ.Delete(ctx, c.Param("id"))
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}
