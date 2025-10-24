import service from '@/utils/request'

export const getPWAConfigList = (data) => {
  return service({
    url: '/navigation/pwaConfig/getPWAConfigList',
    method: 'post',
    data: data
  })
}

export const createPWAConfig = (data) => {
  return service({
    url: '/navigation/pwaConfig/createPWAConfig',
    method: 'post',
    data: data
  })
}

export const updatePWAConfig = (data) => {
  return service({
    url: '/navigation/pwaConfig/updatePWAConfig',
    method: 'post',
    data: data
  })
}

export const deletePWAConfig = (data) => {
  return service({
    url: '/navigation/pwaConfig/deletePWAConfig',
    method: 'post',
    data: data
  })
}

export const getPWAConfigById = (data) => {
  return service({
    url: '/navigation/pwaConfig/findPWAConfig',
    method: 'post',
    data: data
  })
}

export const clearPWACache = () => {
  return service({
    url: '/navigation/pwaConfig/clearPWACache',
    method: 'post'
  })
}
