import service from '@/utils/request'

/**
 * 创建主题配置
 * @param {Object} data 主题配置数据
 * @param {string} data.name 主题名称
 * @param {string} data.description 主题描述
 * @param {string} data.config_json 主题配置JSON
 * @param {number} data.is_default 是否默认主题
 * @returns {Promise} 创建结果
 */
export const createThemeConfig = (data) => {
  return service({
    url: '/navigation/themeConfig/createThemeConfig',
    method: 'post',
    data: data
  })
}

/**
 * 更新主题配置
 * @param {Object} data 主题配置数据
 * @param {number} data.id 主题配置ID
 * @param {string} data.name 主题名称
 * @param {string} data.description 主题描述
 * @param {string} data.config_json 主题配置JSON
 * @param {number} data.is_default 是否默认主题
 * @returns {Promise} 更新结果
 */
export const updateThemeConfig = (data) => {
  return service({
    url: '/navigation/themeConfig/updateThemeConfig',
    method: 'post',
    data: data
  })
}

/**
 * 删除主题配置
 * @param {Object} data 删除参数
 * @param {number} data.id 主题配置ID
 * @returns {Promise} 删除结果
 */
export const deleteThemeConfig = (data) => {
  return service({
    url: '/navigation/themeConfig/deleteThemeConfig',
    method: 'post',
    data: data
  })
}

/**
 * 获取主题配置列表
 * @param {Object} data 查询参数
 * @param {number} data.page 页码
 * @param {number} data.pageSize 每页数量
 * @param {string} data.name 主题名称（可选）
 * @param {string} data.description 主题描述（可选）
 * @returns {Promise} 主题配置列表数据
 */
export const getThemeConfigList = (data) => {
  return service({
    url: '/navigation/themeConfig/getThemeConfigList',
    method: 'post',
    data: data
  })
}

/**
 * 根据ID获取主题配置
 * @param {Object} data 查询参数
 * @param {number} data.id 主题配置ID
 * @returns {Promise} 主题配置详情
 */
export const getThemeConfigById = (data) => {
  return service({
    url: '/navigation/themeConfig/getThemeConfigById',
    method: 'post',
    data: data
  })
}

/**
 * 设置默认主题
 * @param {Object} data 设置参数
 * @param {number} data.id 主题配置ID
 * @returns {Promise} 设置结果
 */
export const setDefaultTheme = (data) => {
  return service({
    url: '/navigation/themeConfig/setDefaultTheme',
    method: 'post',
    data: data
  })
}

/**
 * 获取默认主题
 * @returns {Promise} 默认主题数据
 */
export const getDefaultTheme = () => {
  return service({
    url: '/navigation/themeConfig/getDefaultTheme',
    method: 'post'
  })
}

/**
 * 获取所有主题（用于前端主题切换）
 * @returns {Promise} 所有主题数据
 */
export const getAllThemes = () => {
  return service({
    url: '/navigation/themeConfig/getAllThemes',
    method: 'post'
  })
}
