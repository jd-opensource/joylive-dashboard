/**
 *  服务接口
 */
import request from '@/utils/request'
// 获取service列表
export const getServiceList = (params) => request.basic.get('/api/v1/services', params)
// 获取service单条数据
export const getService = (id) => request.basic.get(`/api/v1/services/${id}`)
// 添加service
export const createService = (params) => request.basic.post('/api/v1/services', params)
// 更新service
export const updateService = (id, params) => request.basic.put(`/api/v1/services/${id}`, params)
// 删除service
export const delService = (id) => request.basic.delete(`/api/v1/services/${id}`)
// 删除service的consumer关系
export const delServiceConsumer = (id) => request.basic.delete(`/api/v1/services/${id}/consumer`)
