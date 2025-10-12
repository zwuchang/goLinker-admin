import service from '@/utils/request'

/**
 * 创建游戏配置
 * @param {Object} data 游戏配置数据
 * @param {string} data.download_url 下载地址
 * @param {string} data.audio_url 播放音频地址
 * @param {string} data.website_title 网站标题
 * @param {string} data.website_desc 网站描述
 * @param {string} data.website_icon 网站图标
 * @param {string} data.website_logo 网站Logo
 * @param {number} data.status 状态
 * @returns {Promise} 创建结果
 */
export const createGameConfig = (data) => {
  return service({
    url: '/navigation/gameConfig/createGameConfig',
    method: 'post',
    data
  })
}

/**
 * 更新游戏配置
 * @param {Object} data 游戏配置数据
 * @param {number} data.id ID
 * @param {string} data.download_url 下载地址
 * @param {string} data.audio_url 播放音频地址
 * @param {string} data.website_title 网站标题
 * @param {string} data.website_desc 网站描述
 * @param {string} data.website_icon 网站图标
 * @param {string} data.website_logo 网站Logo
 * @param {number} data.status 状态
 * @returns {Promise} 更新结果
 */
export const updateGameConfig = (data) => {
  return service({
    url: '/navigation/gameConfig/updateGameConfig',
    method: 'put',
    data
  })
}

/**
 * 删除游戏配置
 * @param {Object} data 删除参数
 * @param {number} data.id ID
 * @returns {Promise} 删除结果
 */
export const deleteGameConfig = (data) => {
  return service({
    url: '/navigation/gameConfig/deleteGameConfig',
    method: 'delete',
    data
  })
}

/**
 * 获取游戏配置
 * @returns {Promise} 游戏配置数据
 */
export const getGameConfig = () => {
  return service({
    url: '/navigation/gameConfig/getGameConfig',
    method: 'get'
  })
}

/**
 * 获取游戏配置列表
 * @param {Object} data 查询参数
 * @param {number} data.page 页码
 * @param {number} data.pageSize 每页数量
 * @param {string} data.website_title 网站标题
 * @param {number} data.status 状态
 * @returns {Promise} 游戏配置列表
 */
export const getGameConfigList = (data) => {
  return service({
    url: '/navigation/gameConfig/getGameConfigList',
    method: 'post',
    data
  })
}
