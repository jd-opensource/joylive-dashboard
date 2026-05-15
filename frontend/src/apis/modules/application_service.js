/**
 *  应用服务关系接口
 */
import request from '@/utils/request'
// 获取application_service列表
export const getApplicationServiceList = (params) => request.basic.get('/api/v1/application-services', params)
// 创建application_service关系
export const createApplicationService = (params) => request.basic.post('/api/v1/application-services', params)
// 更新application_service状态
export const updateApplicationServiceStatus = (id, params) =>
    request.basic.put(`/api/v1/application-services/${id}/status`, params)
