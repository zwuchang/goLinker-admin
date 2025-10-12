import service from '@/utils/request'

/**
 * 创建联系方式
 * @param {Object} data 联系方式数据
 * @param {string} data.displayName 显示名称
 * @param {string} data.jumpUrl 跳转地址
 * @param {string} data.image 联系方式图标
 * @param {number} data.sort 排序
 * @param {number} data.status 状态
 * @returns {Promise} 创建结果
 */
export const createContactMethod = (data) => {
  return service({
    url: '/navigation/contactMethod/create',
    method: 'post',
    data
  })
}

/**
 * 更新联系方式
 * @param {Object} data 联系方式数据
 * @param {number} data.ID 联系方式ID
 * @param {string} data.displayName 显示名称
 * @param {string} data.jumpUrl 跳转地址
 * @param {string} data.image 联系方式图标
 * @param {number} data.sort 排序
 * @param {number} data.status 状态
 * @returns {Promise} 更新结果
 */
export const updateContactMethod = (data) => {
  return service({
    url: '/navigation/contactMethod/update',
    method: 'put',
    data
  })
}

/**
 * 删除联系方式
 * @param {Object} data 联系方式数据
 * @param {number} data.ID 联系方式ID
 * @returns {Promise} 删除结果
 */
export const deleteContactMethod = (data) => {
  return service({
    url: '/navigation/contactMethod/delete',
    method: 'delete',
    data
  })
}

/**
 * 获取单个联系方式
 * @param {Object} params 查询参数
 * @param {number} params.ID 联系方式ID
 * @returns {Promise} 联系方式数据
 */
export const getContactMethod = (params) => {
  return service({
    url: '/navigation/contactMethod/get',
    method: 'get',
    params
  })
}

/**
 * 获取联系方式列表
 * @param {Object} params 查询参数
 * @param {number} params.page 页码
 * @param {number} params.pageSize 每页数量
 * @returns {Promise} 联系方式列表
 */
export const getContactMethodList = (params) => {
  return service({
    url: '/navigation/contactMethod/list',
    method: 'get',
    params
  })
}
