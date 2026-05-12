/**
 *  应用服务关系接口
 */
import request from '@/utils/request'
// 创建application_service关系
export const createApplicationService = (params) => request.basic.post('/api/v1/application-services', params)
