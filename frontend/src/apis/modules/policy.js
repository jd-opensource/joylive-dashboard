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
// 获取路由策略单条数据
export const getRoute = (id) => request.basic.get(`/api/v1/policy/policy-routes/${id}`)
// 添加路由策略
export const createRoute = (params) => request.basic.post('/api/v1/policy/policy-routes', params)
// 更新路由策略
export const updateRoute = (id, params) => request.basic.put(`/api/v1/policy/policy-routes/${id}`, params)
// 删除路由策略
export const delRoute = (id) => request.basic.delete(`/api/v1/policy/policy-routes/${id}`)
// 获取路由详情列表
export const getRouteDetailList = (params) => request.basic.get('/api/v1/policy/policy-route-details', params)
// 获取路由详情单条数据
export const getRouteDetail = (id) => request.basic.get(`/api/v1/policy/policy-route-details/${id}`)
// 添加路由详情
export const createRouteDetail = (params) => request.basic.post('/api/v1/policy/policy-route-details', params)
// 更新路由详情
export const updateRouteDetail = (id, params) => request.basic.put(`/api/v1/policy/policy-route-details/${id}`, params)
// 删除路由详情
export const delRouteDetail = (id) => request.basic.delete(`/api/v1/policy/policy-route-details/${id}`)
// 获取限流策略列表
export const getLimitList = (params) => request.basic.get('/api/v1/policy/policy-limits', params)
// 获取限流策略单条数据
export const getLimit = (id) => request.basic.get(`/api/v1/policy/policy-limits/${id}`)
// 添加限流策略
export const createLimit = (params) => request.basic.post('/api/v1/policy/policy-limits', params)
// 更新限流策略
export const updateLimit = (id, params) => request.basic.put(`/api/v1/policy/policy-limits/${id}`, params)
// 删除限流策略
export const delLimit = (id) => request.basic.delete(`/api/v1/policy/policy-limits/${id}`)
// 获取熔断策略列表
export const getCircuitBreakList = (params) => request.basic.get('/api/v1/policy/policy-circuit-breaks', params)
// 获取熔断策略单条数据
export const getCircuitBreak = (id) => request.basic.get(`/api/v1/policy/policy-circuit-breaks/${id}`)
// 添加熔断策略
export const createCircuitBreak = (params) => request.basic.post('/api/v1/policy/policy-circuit-breaks', params)
// 更新熔断策略
export const updateCircuitBreak = (id, params) =>
    request.basic.put(`/api/v1/policy/policy-circuit-breaks/${id}`, params)
// 删除熔断策略
export const delCircuitBreak = (id) => request.basic.delete(`/api/v1/policy/policy-circuit-breaks/${id}`)
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
// 获取容错策略单条数据
export const getFault = (id) => request.basic.get(`/api/v1/policy/policy-faults/${id}`)
// 添加容错策略
export const createFault = (params) => request.basic.post('/api/v1/policy/policy-faults', params)
// 更新容错策略
export const updateFault = (id, params) => request.basic.put(`/api/v1/policy/policy-faults/${id}`, params)
// 删除容错策略
export const delFault = (id) => request.basic.delete(`/api/v1/policy/policy-faults/${id}`)
// 获取调用容错策略列表
export const getInvocationList = (params) => request.basic.get('/api/v1/policy/policy-invocations', params)
// 获取调用容错策略单条数据
export const getInvocation = (id) => request.basic.get(`/api/v1/policy/policy-invocations/${id}`)
// 添加调用容错策略
export const createInvocation = (params) => request.basic.post('/api/v1/policy/policy-invocations', params)
// 更新调用容错策略
export const updateInvocation = (id, params) => request.basic.put(`/api/v1/policy/policy-invocations/${id}`, params)
// 删除调用容错策略
export const delInvocation = (id) => request.basic.delete(`/api/v1/policy/policy-invocations/${id}`)
