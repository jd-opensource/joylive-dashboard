/**
 *  应用接口
 */
import request from '@/utils/request'
// 获取application列表
export const getApplicationList = (params) => request.basic.get('/api/v1/applications', params)
// 获取application单条数据
export const getApplication = (id) => request.basic.get(`/api/v1/applications/${id}`)
// 添加application
export const createApplication = (params) => request.basic.post('/api/v1/applications', params)
// 更新application
export const updateApplication = (id, params) => request.basic.put(`/api/v1/applications/${id}`, params)
// 删除application
export const delApplication = (id) => request.basic.delete(`/api/v1/applications/${id}`)
