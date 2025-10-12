import service from '@/utils/request'

/**
 * 获取联系配置
 * @returns {Promise} 联系配置数据
 */
export const getContactConfig = () => {
  return service({
    url: '/navigation/contactConfig/get',
    method: 'get'
  })
}

/**
 * 更新联系配置
 * @param {Object} data 联系配置数据
 * @param {string} data.email 客服邮箱
 * @param {string} data.bannerImage 横幅图片
 * @returns {Promise} 更新结果
 */
export const updateContactConfig = (data) => {
  return service({
    url: '/navigation/contactConfig/update',
    method: 'put',
    data
  })
}
