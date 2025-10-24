import service from '@/utils/request'

/**
 * 创建市场配置
 * @param {Object} data 市场配置数据
 * @param {string} data.name 市场名称
 * @param {string} data.logo 市场Logo地址
 * @param {string} data.jump_url 跳转地址
 * @param {number} data.type 市场类型（1-3）
 * @param {number} data.status 状态（0禁用 1启用）
 * @param {number} data.is_visible 是否可见（0隐藏 1可见）
 * @param {number} data.sort 排序
 * @param {string} data.description 市场描述
 * @returns {Promise} 创建结果
 */
export const createMarketConfig = (data) => {
  return service({
    url: '/navigation/marketConfig/createMarketConfig',
    method: 'post',
    data
  })
}

/**
 * 删除市场配置
 * @param {Object} data 删除参数
 * @param {number} data.id 市场配置ID
 * @returns {Promise} 删除结果
 */
export const deleteMarketConfig = (data) => {
  return service({
    url: '/navigation/marketConfig/deleteMarketConfig',
    method: 'post',
    data
  })
}

/**
 * 更新市场配置
 * @param {Object} data 市场配置数据
 * @param {number} data.id 市场配置ID
 * @param {string} data.name 市场名称
 * @param {string} data.logo 市场Logo地址
 * @param {string} data.jump_url 跳转地址
 * @param {number} data.type 市场类型（1-3）
 * @param {number} data.status 状态（0禁用 1启用）
 * @param {number} data.is_visible 是否可见（0隐藏 1可见）
 * @param {number} data.sort 排序
 * @param {string} data.description 市场描述
 * @returns {Promise} 更新结果
 */
export const updateMarketConfig = (data) => {
  return service({
    url: '/navigation/marketConfig/updateMarketConfig',
    method: 'post',
    data
  })
}

/**
 * 根据ID获取市场配置
 * @param {Object} params 查询参数
 * @param {number} params.id 市场配置ID
 * @returns {Promise} 市场配置详情
 */
export const getMarketConfigById = (params) => {
  return service({
    url: `/navigation/marketConfig/getMarketConfigById/${params.id}`,
    method: 'get'
  })
}

/**
 * 获取市场配置列表
 * @param {Object} data 查询参数
 * @param {number} data.page 页码
 * @param {number} data.pageSize 每页数量
 * @param {string} data.name 市场名称（模糊搜索）
 * @param {number} data.type 市场类型
 * @param {number} data.status 状态
 * @param {number} data.is_visible 是否可见
 * @returns {Promise} 市场配置列表
 */
export const getMarketConfigList = (data) => {
  return service({
    url: '/navigation/marketConfig/getMarketConfigList',
    method: 'post',
    data
  })
}

/**
 * 根据类型获取市场配置列表（公开接口）
 * @param {Object} data 查询参数
 * @param {number} data.type 市场类型（1-3）
 * @returns {Promise} 市场配置列表
 */
export const getMarketByType = (data) => {
  return service({
    url: '/public/game/getMarket',
    method: 'post',
    data
  })
}




