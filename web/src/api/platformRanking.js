import service from '@/utils/request'

/**
 * 获取平台排行榜列表
 * @param {Object} data 查询参数
 * @param {number} data.page 页码
 * @param {number} data.pageSize 每页数量
 * @returns {Promise} 平台排行榜列表数据
 */
export const getPlatformRankingList = (data) => {
  return service({
    url: '/navigation/platformRanking/getPlatformRankingList',
    method: 'post',
    data: data
  })
}

/**
 * 创建平台排行榜
 * @param {Object} data 平台排行榜数据
 * @returns {Promise} 创建结果
 */
export const createPlatformRanking = (data) => {
  return service({
    url: '/navigation/platformRanking/createPlatformRanking',
    method: 'post',
    data: data
  })
}

/**
 * 更新平台排行榜
 * @param {Object} data 平台排行榜数据
 * @returns {Promise} 更新结果
 */
export const updatePlatformRanking = (data) => {
  return service({
    url: '/navigation/platformRanking/updatePlatformRanking',
    method: 'post',
    data: data
  })
}

/**
 * 删除平台排行榜
 * @param {Object} data 删除参数
 * @param {number} data.id 平台排行榜ID
 * @returns {Promise} 删除结果
 */
export const deletePlatformRanking = (data) => {
  return service({
    url: '/navigation/platformRanking/deletePlatformRanking',
    method: 'post',
    data: data
  })
}

/**
 * 根据ID获取平台排行榜
 * @param {Object} data 查询参数
 * @param {number} data.id 平台排行榜ID
 * @returns {Promise} 平台排行榜数据
 */
export const getPlatformRankingById = (data) => {
  return service({
    url: '/navigation/platformRanking/findPlatformRanking',
    method: 'get',
    params: data
  })
}
