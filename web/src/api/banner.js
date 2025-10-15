import service from '@/utils/request'

/**
 * 获取Banner列表
 * @param {Object} data 查询参数
 * @param {number} data.page 页码
 * @param {number} data.pageSize 每页数量
 * @param {string} data.title 标题搜索
 * @param {number} data.mediaType 媒体类型 1:图片 2:视频
 * @param {number} data.isVisible 是否显示 0:隐藏 1:显示
 * @param {number} data.status 状态 0:禁用 1:启用
 * @param {string} data.orderBy 排序字段
 * @param {string} data.orderType 排序类型 asc/desc
 * @returns {Promise} Banner列表数据
 */
export const getBannerList = (data) => {
  return service({
    url: '/navigation/banner/getBannerList',
    method: 'post',
    data: data
  })
}

/**
 * 根据ID获取Banner
 * @param {number} id Banner ID
 * @returns {Promise} Banner详情数据
 */
export const getBannerById = (id) => {
  return service({
    url: `/navigation/banner/getBannerById/${id}`,
    method: 'get'
  })
}

/**
 * 创建Banner
 * @param {Object} data Banner信息
 * @param {string} data.title 标题
 * @param {string} data.description 描述
 * @param {number} data.mediaType 媒体类型 1:图片 2:视频
 * @param {string} data.mediaUrl 媒体地址
 * @param {string} data.linkUrl 跳转链接
 * @param {number} data.sort 排序
 * @param {number} data.isVisible 是否显示
 * @param {number} data.status 状态
 * @returns {Promise} 创建结果
 */
export const createBanner = (data) => {
  return service({
    url: '/navigation/banner/createBanner',
    method: 'post',
    data: data
  })
}

/**
 * 更新Banner
 * @param {Object} data Banner信息
 * @param {number} data.id Banner ID
 * @param {string} data.title 标题
 * @param {string} data.description 描述
 * @param {number} data.mediaType 媒体类型
 * @param {string} data.mediaUrl 媒体地址
 * @param {string} data.linkUrl 跳转链接
 * @param {number} data.sort 排序
 * @param {number} data.isVisible 是否显示
 * @param {number} data.status 状态
 * @returns {Promise} 更新结果
 */
export const updateBanner = (data) => {
  return service({
    url: '/navigation/banner/updateBanner',
    method: 'post',
    data: data
  })
}

/**
 * 删除Banner
 * @param {number} id Banner ID
 * @returns {Promise} 删除结果
 */
export const deleteBanner = (id) => {
  return service({
    url: '/navigation/banner/deleteBanner',
    method: 'post',
    data: { id: id }
  })
}
