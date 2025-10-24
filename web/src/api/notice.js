import service from '@/utils/request'

/**
 * 获取滚动通知列表
 * @param {Object} data 查询参数
 * @param {number} data.page 页码
 * @param {number} data.pageSize 每页数量
 * @param {string} data.title 通知标题（可选）
 * @param {number} data.status 状态（可选）
 * @param {number} data.isVisible 是否显示（可选）
 * @returns {Promise} 滚动通知列表数据
 */
export const getNoticeList = (data) => {
  return service({
    url: '/navigation/notice/getNoticeList',
    method: 'post',
    data: data
  })
}

/**
 * 创建滚动通知
 * @param {Object} data 创建参数
 * @param {string} data.title 通知标题
 * @param {string} data.content 通知内容
 * @param {string} data.linkUrl 跳转链接（可选）
 * @param {number} data.status 状态
 * @param {number} data.isVisible 是否显示
 * @param {number} data.sort 排序
 * @param {string} data.startTime 开始时间（可选）
 * @param {string} data.endTime 结束时间（可选）
 * @returns {Promise} 创建结果
 */
export const createNotice = (data) => {
  return service({
    url: '/navigation/notice/createNotice',
    method: 'post',
    data: data
  })
}

/**
 * 更新滚动通知
 * @param {Object} data 更新参数
 * @param {number} data.id 通知ID
 * @param {string} data.title 通知标题（可选）
 * @param {string} data.content 通知内容（可选）
 * @param {string} data.linkUrl 跳转链接（可选）
 * @param {number} data.status 状态（可选）
 * @param {number} data.isVisible 是否显示（可选）
 * @param {number} data.sort 排序（可选）
 * @param {string} data.startTime 开始时间（可选）
 * @param {string} data.endTime 结束时间（可选）
 * @returns {Promise} 更新结果
 */
export const updateNotice = (data) => {
  return service({
    url: '/navigation/notice/updateNotice',
    method: 'post',
    data: data
  })
}

/**
 * 删除滚动通知
 * @param {Object} data 删除参数
 * @param {number} data.id 通知ID
 * @returns {Promise} 删除结果
 */
export const deleteNotice = (data) => {
  return service({
    url: '/navigation/notice/deleteNotice',
    method: 'post',
    data: data
  })
}

/**
 * 根据ID获取滚动通知
 * @param {Object} data 查询参数
 * @param {number} data.id 通知ID
 * @returns {Promise} 滚动通知详情
 */
export const getNoticeById = (data) => {
  return service({
    url: '/navigation/notice/getNoticeById',
    method: 'post',
    data: data
  })
}
