package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jd-opensource/joylive-dashboard/internal/mods/policy/biz"
	"github.com/jd-opensource/joylive-dashboard/internal/mods/policy/schema"
	"github.com/jd-opensource/joylive-dashboard/pkg/util"
)

// Permission policy management
type PolicyPermission struct {
	PolicyPermissionBIZ *biz.PolicyPermission
}

// @Tags PolicyPermissionAPI
// @Security ApiKeyAuth
// @Summary Query policy permission list
// @Success 200 {object} util.ResponseResult{data=[]schema.PolicyPermission}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/policy/policy-permissions [get]
func (a *PolicyPermission) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.PolicyPermissionQueryParam
	if err := util.ParseQuery(c, &params); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.PolicyPermissionBIZ.Query(ctx, params)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResPage(c, result.Data, result.PageResult)
}

// @Tags PolicyPermissionAPI
// @Security ApiKeyAuth
// @Summary Get policy permission record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult{data=schema.PolicyPermissionForm}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/policy/policy-permissions/{id} [get]
func (a *PolicyPermission) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.PolicyPermissionBIZ.Get(ctx, c.Param("id"))
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, item)
}

// @Tags PolicyPermissionAPI
// @Security ApiKeyAuth
// @Summary Create policy permission record
// @Param body body schema.PolicyPermissionForm true "Request body"
// @Success 200 {object} util.ResponseResult{data=schema.PolicyPermission}
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/policy/policy-permissions [post]
func (a *PolicyPermission) Create(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(schema.PolicyPermissionForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.PolicyPermissionBIZ.Create(ctx, item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, result)
}

// @Tags PolicyPermissionAPI
// @Security ApiKeyAuth
// @Summary Update policy permission record by ID
// @Param id path string true "unique id"
// @Param body body schema.PolicyPermissionForm true "Request body"
// @Success 200 {object} util.ResponseResult
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/policy/policy-permissions/{id} [put]
func (a *PolicyPermission) Update(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(schema.PolicyPermissionForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.ResError(c, err)
		return
	}

	err := a.PolicyPermissionBIZ.Update(ctx, c.Param("id"), item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}

// @Tags PolicyPermissionAPI
// @Security ApiKeyAuth
// @Summary Delete policy permission record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/policy/policy-permissions/{id} [delete]
func (a *PolicyPermission) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.PolicyPermissionBIZ.Delete(ctx, c.Param("id"))
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}
