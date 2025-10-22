import service from '@/utils/request'

/**
 * 获取平台游戏配置列表
 * @param {Object} data 查询参数
 * @param {number} data.page 页码
 * @param {number} data.pageSize 每页数量
 * @param {string} data.platform_name 平台名称
 * @param {number} data.status 状态
 * @param {number} data.is_visible 是否可见
 * @returns {Promise} 平台游戏配置列表数据
 */
export const getPlatformConfigList = (data) => {
  return service({
    url: '/navigation/platformConfig/getPlatformConfigList',
    method: 'post',
    data: data
  })
}

/**
 * 创建平台游戏配置
 * @param {Object} data 创建参数
 * @param {string} data.platform_name 平台名称
 * @param {string} data.platform_icon 平台图标地址
 * @param {string} data.platform_api 平台接口地址
 * @param {number} data.sort 排序
 * @param {number} data.status 状态
 * @param {string} data.description 平台描述
 * @param {number} data.is_visible 是否可见
 * @returns {Promise} 创建结果
 */
export const createPlatformConfig = (data) => {
  return service({
    url: '/navigation/platformConfig/createPlatformConfig',
    method: 'post',
    data: data
  })
}

/**
 * 更新平台游戏配置
 * @param {Object} data 更新参数
 * @param {number} data.id ID
 * @param {string} data.platform_name 平台名称
 * @param {string} data.platform_icon 平台图标地址
 * @param {string} data.platform_api 平台接口地址
 * @param {number} data.sort 排序
 * @param {number} data.status 状态
 * @param {string} data.description 平台描述
 * @param {number} data.is_visible 是否可见
 * @returns {Promise} 更新结果
 */
export const updatePlatformConfig = (data) => {
  return service({
    url: '/navigation/platformConfig/updatePlatformConfig',
    method: 'post',
    data: data
  })
}

/**
 * 根据ID获取平台游戏配置
 * @param {number} id 平台游戏配置ID
 * @returns {Promise} 平台游戏配置详情
 */
export const getPlatformConfigById = (id) => {
  return service({
    url: `/navigation/platformConfig/getPlatformConfigById/${id}`,
    method: 'get'
  })
}

/**
 * 删除平台游戏配置
 * @param {Object} data 删除参数
 * @param {number} data.id ID
 * @returns {Promise} 删除结果
 */
export const deletePlatformConfig = (data) => {
  return service({
    url: '/navigation/platformConfig/deletePlatformConfig',
    method: 'post',
    data: data
  })
}


