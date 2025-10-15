import service from '@/utils/request'

/**
 * 获取访问统计列表
 * @param {Object} data 搜索参数
 * @returns {Promise} 访问统计列表
 */
export const getAccessStatsList = (data) => {
  return service({
    url: '/navigation/accessStats/getAccessStatsList',
    method: 'post',
    data: data
  })
}

/**
 * 获取访问统计汇总
 * @param {Object} data 查询参数
 * @returns {Promise} 访问统计汇总
 */
export const getAccessStatsSummary = (data) => {
  return service({
    url: '/navigation/accessStats/getAccessStatsSummary',
    method: 'post',
    data: data
  })
}

/**
 * 清理旧数据
 * @param {Object} data 清理参数
 * @returns {Promise} 清理结果
 */
export const cleanOldStats = (data) => {
  return service({
    url: '/navigation/accessStats/cleanOldStats',
    method: 'post',
    data: data
  })
}
