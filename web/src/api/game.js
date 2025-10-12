import service from '@/utils/request'

/**
 * 创建游戏
 * @param {Object} data 游戏数据
 * @returns {Promise} 创建结果
 */
export const createGame = (data) => {
  return service({
    url: '/navigation/game/createGame',
    method: 'post',
    data: data
  })
}

/**
 * 更新游戏
 * @param {Object} data 游戏数据
 * @returns {Promise} 更新结果
 */
export const updateGame = (data) => {
  return service({
    url: '/navigation/game/updateGame',
    method: 'put',
    data: data
  })
}

/**
 * 删除游戏
 * @param {Object} data 包含ID的对象
 * @returns {Promise} 删除结果
 */
export const deleteGame = (data) => {
  return service({
    url: '/navigation/game/deleteGame',
    method: 'delete',
    data: data
  })
}

/**
 * 获取游戏列表
 * @param {Object} data 分页和搜索参数
 * @returns {Promise} 游戏列表
 */
export const getGameList = (data) => {
  return service({
    url: '/navigation/game/getGameList',
    method: 'post',
    data: data
  })
}

/**
 * 更新游戏浏览次数
 * @param {Object} data 包含ID的对象
 * @returns {Promise} 更新结果
 */
export const updateGameViews = (data) => {
  return service({
    url: '/navigation/game/updateGameViews',
    method: 'post',
    data: data
  })
}
