package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jd-opensource/joylive-dashboard/internal/mods/policy/biz"
	"github.com/jd-opensource/joylive-dashboard/internal/mods/policy/schema"
	"github.com/jd-opensource/joylive-dashboard/pkg/util"
)

// Auth policy management
type PolicyAuth struct {
	PolicyAuthBIZ *biz.PolicyAuth
}

// @Tags PolicyAuthAPI
// @Security ApiKeyAuth
// @Summary Query policy auth list
// @Success 200 {object} util.ResponseResult{data=[]schema.PolicyAuth}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/policy/policy-auths [get]
func (a *PolicyAuth) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.PolicyAuthQueryParam
	if err := util.ParseQuery(c, &params); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.PolicyAuthBIZ.Query(ctx, params)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResPage(c, result.Data, result.PageResult)
}

// @Tags PolicyAuthAPI
// @Security ApiKeyAuth
// @Summary Get policy auth record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult{data=schema.PolicyAuth}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/policy/policy-auths/{id} [get]
func (a *PolicyAuth) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.PolicyAuthBIZ.Get(ctx, c.Param("id"))
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, item)
}

// @Tags PolicyAuthAPI
// @Security ApiKeyAuth
// @Summary Create policy auth record
// @Param body body schema.PolicyAuthForm true "Request body"
// @Success 200 {object} util.ResponseResult{data=schema.PolicyAuth}
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/policy/policy-auths [post]
func (a *PolicyAuth) Create(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(schema.PolicyAuthForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.PolicyAuthBIZ.Create(ctx, item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, result)
}

// @Tags PolicyAuthAPI
// @Security ApiKeyAuth
// @Summary Update policy auth record by ID
// @Param id path string true "unique id"
// @Param body body schema.PolicyAuthForm true "Request body"
// @Success 200 {object} util.ResponseResult
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/policy/policy-auths/{id} [put]
func (a *PolicyAuth) Update(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(schema.PolicyAuthForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.ResError(c, err)
		return
	}

	err := a.PolicyAuthBIZ.Update(ctx, c.Param("id"), item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}

// @Tags PolicyAuthAPI
// @Security ApiKeyAuth
// @Summary Delete policy auth record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/policy/policy-auths/{id} [delete]
func (a *PolicyAuth) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.PolicyAuthBIZ.Delete(ctx, c.Param("id"))
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}
