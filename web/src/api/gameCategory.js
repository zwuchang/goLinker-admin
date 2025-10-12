import service from '@/utils/request'

/**
 * 创建游戏类别
 * @param {Object} data 游戏类别数据
 * @returns {Promise} 创建结果
 */
export const createGameCategory = (data) => {
  return service({
    url: '/navigation/gameCategory/createGameCategory',
    method: 'post',
    data: data
  })
}

/**
 * 更新游戏类别
 * @param {Object} data 游戏类别数据
 * @returns {Promise} 更新结果
 */
export const updateGameCategory = (data) => {
  return service({
    url: '/navigation/gameCategory/updateGameCategory',
    method: 'put',
    data: data
  })
}

/**
 * 删除游戏类别
 * @param {Object} data 包含ID的对象
 * @returns {Promise} 删除结果
 */
export const deleteGameCategory = (data) => {
  return service({
    url: '/navigation/gameCategory/deleteGameCategory',
    method: 'delete',
    data: data
  })
}

/**
 * 获取游戏类别列表
 * @param {Object} data 分页和搜索参数
 * @returns {Promise} 游戏类别列表
 */
export const getGameCategoryList = (data) => {
  return service({
    url: '/navigation/gameCategory/getGameCategoryList',
    method: 'post',
    data: data
  })
}

/**
 * 获取所有启用的游戏类别
 * @returns {Promise} 游戏类别列表
 */
export const getAllGameCategories = () => {
  return service({
    url: '/navigation/gameCategory/getAllGameCategories',
    method: 'get'
  })
}
