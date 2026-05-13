/**
 *  治理策略接口
 */
import request from '@/utils/request'
// 获取负载均衡策略列表
export const getLoadbalanceList = (params) => request.basic.get('/api/v1/policy/policy-loadbalances', params)
// 获取负载均衡策略单条数据
export const getLoadbalance = (id) => request.basic.get(`/api/v1/policy/policy-loadbalances/${id}`)
// 添加负载均衡策略
export const createLoadbalance = (params) => request.basic.post('/api/v1/policy/policy-loadbalances', params)
// 更新负载均衡策略
export const updateLoadbalance = (id, params) => request.basic.put(`/api/v1/policy/policy-loadbalances/${id}`, params)
// 删除负载均衡策略
export const delLoadbalance = (id) => request.basic.delete(`/api/v1/policy/policy-loadbalances/${id}`)
// 获取路由策略列表
export const getRouteList = (params) => request.basic.get('/api/v1/policy/policy-routes', params)
// 获取限流策略列表
export const getLimitList = (params) => request.basic.get('/api/v1/policy/policy-limits', params)
// 获取熔断策略列表
export const getCircuitBreakList = (params) => request.basic.get('/api/v1/policy/policy-circuit-breaks', params)
// 获取权限策略列表
export const getPermissionList = (params) => request.basic.get('/api/v1/policy/policy-permissions', params)
// 获取权限策略单条数据
export const getPermission = (id) => request.basic.get(`/api/v1/policy/policy-permissions/${id}`)
// 添加权限策略
export const createPermission = (params) => request.basic.post('/api/v1/policy/policy-permissions', params)
// 更新权限策略
export const updatePermission = (id, params) => request.basic.put(`/api/v1/policy/policy-permissions/${id}`, params)
// 删除权限策略
export const delPermission = (id) => request.basic.delete(`/api/v1/policy/policy-permissions/${id}`)
// 获取认证策略列表
export const getAuthList = (params) => request.basic.get('/api/v1/policy/policy-auths', params)
// 获取容错策略列表
export const getFaultList = (params) => request.basic.get('/api/v1/policy/policy-faults', params)
// 获取调用策略列表
export const getInvocationList = (params) => request.basic.get('/api/v1/policy/policy-invocations', params)
