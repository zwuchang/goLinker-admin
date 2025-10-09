import service from '@/utils/request'

// 获取定时任务状态
export const getCronStatus = () => {
  return service({
    url: '/cron/list',
    method: 'post'
  })
}

// 更新定时任务规则
export const updateCronSpec = (data) => {
  return service({
    url: '/cron/update',
    method: 'post',
    data
  })
}

// 启动定时任务
export const startCron = (data) => {
  return service({
    url: '/cron/start',
    method: 'post',
    data
  })
}

// 暂停定时任务
export const stopCron = (data) => {
  return service({
    url: '/cron/stop',
    method: 'post',
    data
  })
}