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
