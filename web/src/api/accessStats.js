import service from '@/utils/request'

/**
 * 获取访问统计汇总
 * @param {Object} data 查询参数
 * @param {string} data.startTime 开始时间
 * @param {string} data.endTime 结束时间
 * @returns {Promise} 访问统计汇总数据
 */
export const getAccessStatsSummary = (data) => {
  return service({
    url: '/navigation/accessStats/getAccessStatsSummary',
    method: 'post',
    data: data
  })
}

/**
 * 获取访问统计列表
 * @param {Object} data 查询参数
 * @param {number} data.page 页码
 * @param {number} data.pageSize 每页数量
 * @param {string} data.startTime 开始时间
 * @param {string} data.endTime 结束时间
 * @param {string} data.orderBy 排序字段
 * @param {string} data.orderType 排序类型
 * @returns {Promise} 访问统计列表数据
 */
export const getAccessStatsList = (data) => {
  return service({
    url: '/navigation/accessStats/getAccessStatsList',
    method: 'post',
    data: data
  })
}