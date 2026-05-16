package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jd-opensource/joylive-dashboard/internal/mods/policy/biz"
	"github.com/jd-opensource/joylive-dashboard/internal/mods/policy/schema"
	"github.com/jd-opensource/joylive-dashboard/pkg/util"
)

// Fault injection policy management
type PolicyFault struct {
	PolicyFaultBIZ *biz.PolicyFault
}

// @Tags PolicyFaultAPI
// @Security ApiKeyAuth
// @Summary Query policy fault list
// @Success 200 {object} util.ResponseResult{data=[]schema.PolicyFault}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/policy/policy-faults [get]
func (a *PolicyFault) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.PolicyFaultQueryParam
	if err := util.ParseQuery(c, &params); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.PolicyFaultBIZ.Query(ctx, params)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResPage(c, result.Data, result.PageResult)
}

// @Tags PolicyFaultAPI
// @Security ApiKeyAuth
// @Summary Get policy fault record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult{data=schema.PolicyFaultForm}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/policy/policy-faults/{id} [get]
func (a *PolicyFault) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.PolicyFaultBIZ.Get(ctx, c.Param("id"))
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, item)
}

// @Tags PolicyFaultAPI
// @Security ApiKeyAuth
// @Summary Create policy fault record
// @Param body body schema.PolicyFaultForm true "Request body"
// @Success 200 {object} util.ResponseResult{data=schema.PolicyFault}
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/policy/policy-faults [post]
func (a *PolicyFault) Create(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(schema.PolicyFaultForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.PolicyFaultBIZ.Create(ctx, item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, result)
}

// @Tags PolicyFaultAPI
// @Security ApiKeyAuth
// @Summary Update policy fault record by ID
// @Param id path string true "unique id"
// @Param body body schema.PolicyFaultForm true "Request body"
// @Success 200 {object} util.ResponseResult
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/policy/policy-faults/{id} [put]
func (a *PolicyFault) Update(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(schema.PolicyFaultForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.ResError(c, err)
		return
	}

	err := a.PolicyFaultBIZ.Update(ctx, c.Param("id"), item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}

// @Tags PolicyFaultAPI
// @Security ApiKeyAuth
// @Summary Delete policy fault record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/policy/policy-faults/{id} [delete]
func (a *PolicyFault) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.PolicyFaultBIZ.Delete(ctx, c.Param("id"))
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}
