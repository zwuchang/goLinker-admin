<template>
  <div class="dashboard-container">
    <!-- 时间查询区域 -->
    <el-card class="time-filter-card">
      <div class="time-filter-content">
        <div class="time-filter-left">
          <span class="filter-label">时间范围：</span>
          <el-radio-group v-model="timeRange" @change="onTimeRangeChange">
            <el-radio-button label="today">今天</el-radio-button>
            <el-radio-button label="yesterday">昨天</el-radio-button>
            <el-radio-button label="month">本月</el-radio-button>
            <el-radio-button label="lastMonth">上个月</el-radio-button>
            <el-radio-button label="custom">自定义</el-radio-button>
          </el-radio-group>
        </div>
        <div class="time-filter-right" v-if="timeRange === 'custom'">
          <el-date-picker
            v-model="customDateRange"
            type="datetimerange"
            range-separator="至"
            start-placeholder="开始时间"
            end-placeholder="结束时间"
            format="YYYY-MM-DD HH:mm:ss"
            value-format="YYYY-MM-DD HH:mm:ss"
            @change="onCustomDateChange"
          />
        </div>
        <div class="time-filter-actions">
          <el-button type="primary" @click="refreshData" :loading="loading">
            <i class="el-icon-refresh"></i>
            刷新数据
          </el-button>
        </div>
      </div>
      <!-- 时间区间显示 -->
      <div class="time-range-display" v-if="currentTimeRange.startTime && currentTimeRange.endTime">
        <span class="range-label">当前时间范围：</span>
        <span class="range-text">
          {{ formatDisplayTime(currentTimeRange.startTime) }} 至 {{ formatDisplayTime(currentTimeRange.endTime) }}
        </span>
      </div>
    </el-card>

    <!-- 统计卡片 -->
    <el-row :gutter="20" class="stats-cards">
      <el-col :span="6">
        <el-card class="stat-card">
          <div class="stat-content">
            <div class="stat-icon total-requests">
              <i class="el-icon-connection"></i>
            </div>
            <div class="stat-info">
              <div class="stat-number">{{ summaryData.totalRequests || 0 }}</div>
              <div class="stat-label">总请求数</div>
            </div>
          </div>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card class="stat-card">
          <div class="stat-content">
            <div class="stat-icon unique-ips">
              <i class="el-icon-user"></i>
            </div>
            <div class="stat-info">
              <div class="stat-number">{{ summaryData.uniqueIPs || 0 }}</div>
              <div class="stat-label">独立IP数</div>
            </div>
          </div>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card class="stat-card">
          <div class="stat-content">
            <div class="stat-icon response-time">
              <i class="el-icon-time"></i>
            </div>
            <div class="stat-info">
              <div class="stat-number">{{ Math.round(summaryData.avgResponseTime || 0) }}ms</div>
              <div class="stat-label">平均响应时间</div>
            </div>
          </div>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card class="stat-card">
          <div class="stat-content">
            <div class="stat-icon success-rate">
              <i class="el-icon-success"></i>
            </div>
            <div class="stat-info">
              <div class="stat-number">{{ Math.round(summaryData.successRate || 0) }}%</div>
              <div class="stat-label">成功率</div>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <!-- 图表区域 -->
    <el-row :gutter="20" class="charts-section">
      <!-- 热门API路径 -->
      <el-col :span="12">
        <el-card class="chart-card">
          <div slot="header" class="card-header">
            <span>热门API路径</span>
            <el-button type="text" @click="refreshData">刷新</el-button>
          </div>
          <div class="chart-container">
            <div v-if="summaryData.topApiPaths && summaryData.topApiPaths.length > 0">
              <div 
                v-for="(item, index) in summaryData.topApiPaths" 
                :key="index"
                class="api-path-item"
              >
                <div class="api-path-name">{{ item.apiPath }}</div>
                <div class="api-path-count">{{ item.count }} 次</div>
              </div>
            </div>
            <el-empty v-else description="暂无数据"></el-empty>
          </div>
        </el-card>
      </el-col>

      <!-- 热门国家 -->
      <el-col :span="12">
        <el-card class="chart-card">
          <div slot="header" class="card-header">
            <span>访问地区分布</span>
          </div>
          <div class="chart-container">
            <div v-if="summaryData.topCountries && summaryData.topCountries.length > 0">
              <div 
                v-for="(item, index) in summaryData.topCountries" 
                :key="index"
                class="country-item"
              >
                <div class="country-name">{{ item.country || '未知' }}</div>
                <div class="country-count">{{ item.count }} 次</div>
              </div>
            </div>
            <el-empty v-else description="暂无数据"></el-empty>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <!-- 设备统计 -->
    <el-row :gutter="20" class="charts-section">
      <!-- 设备类型 -->
      <el-col :span="12">
        <el-card class="chart-card">
          <div slot="header" class="card-header">
            <span>设备类型分布</span>
          </div>
          <div class="chart-container">
            <div v-if="summaryData.topDevices && summaryData.topDevices.length > 0">
              <div 
                v-for="(item, index) in summaryData.topDevices" 
                :key="index"
                class="device-item"
              >
                <div class="device-name">{{ item.device }}</div>
                <div class="device-count">{{ item.count }} 次</div>
              </div>
            </div>
            <el-empty v-else description="暂无数据"></el-empty>
          </div>
        </el-card>
      </el-col>

      <!-- 浏览器分布 -->
      <el-col :span="12">
        <el-card class="chart-card">
          <div slot="header" class="card-header">
            <span>浏览器分布</span>
          </div>
          <div class="chart-container">
            <div v-if="summaryData.topBrowsers && summaryData.topBrowsers.length > 0">
              <div 
                v-for="(item, index) in summaryData.topBrowsers" 
                :key="index"
                class="browser-item"
              >
                <div class="browser-name">{{ item.browser }}</div>
                <div class="browser-count">{{ item.count }} 次</div>
              </div>
            </div>
            <el-empty v-else description="暂无数据"></el-empty>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <!-- 最近访问记录 -->
    <el-card class="recent-access-card">
      <div slot="header" class="card-header">
        <span>最近访问记录</span>
        <el-button type="text" @click="refreshRecentAccess">刷新</el-button>
      </div>
      <el-table 
        :data="recentAccessList" 
        v-loading="recentLoading"
        stripe
        style="width: 100%"
      >
        <el-table-column prop="apiPath" label="API路径" min-width="200"></el-table-column>
        <el-table-column prop="method" label="方法" width="80"></el-table-column>
        <el-table-column prop="clientIp" label="客户端IP" width="120"></el-table-column>
        <el-table-column prop="country" label="地区" width="100"></el-table-column>
        <el-table-column prop="device" label="设备" width="100"></el-table-column>
        <el-table-column prop="browser" label="浏览器" width="100"></el-table-column>
        <el-table-column prop="responseTime" label="响应时间" width="100">
          <template #default="scope">
            {{ scope.row.responseTime }}ms
          </template>
        </el-table-column>
        <el-table-column prop="statusCode" label="状态码" width="80">
          <template #default="scope">
            <el-tag 
              :type="scope.row.statusCode >= 200 && scope.row.statusCode < 300 ? 'success' : 'danger'"
              size="small"
            >
              {{ scope.row.statusCode }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="createdAt" label="访问时间" width="160">
          <template #default="scope">
            {{ formatDate(scope.row.CreatedAt) }}
          </template>
        </el-table-column>
      </el-table>
    </el-card>
  </div>
</template>

<script setup>
import { ref, onMounted, onBeforeUnmount } from 'vue'
import { getAccessStatsSummary, getAccessStatsList } from '@/api/accessStats'
import { formatDate } from '@/utils/format'

// 响应式数据
const summaryData = ref({})
const recentAccessList = ref([])
const recentLoading = ref(false)
const loading = ref(false)
const refreshTimer = ref(null)

// 时间查询相关
const timeRange = ref('yesterday') // 默认昨天
const customDateRange = ref([])
const currentTimeRange = ref({
  startTime: '',
  endTime: ''
})

// 格式化日期为本地时间字符串
const formatLocalDateTime = (date) => {
  const year = date.getFullYear()
  const month = String(date.getMonth() + 1).padStart(2, '0')
  const day = String(date.getDate()).padStart(2, '0')
  const hours = String(date.getHours()).padStart(2, '0')
  const minutes = String(date.getMinutes()).padStart(2, '0')
  const seconds = String(date.getSeconds()).padStart(2, '0')
  return `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`
}

// 格式化显示时间（用于界面显示）
const formatDisplayTime = (timeStr) => {
  if (!timeStr) return ''
  const date = new Date(timeStr.replace(' ', 'T'))
  const year = date.getFullYear()
  const month = String(date.getMonth() + 1).padStart(2, '0')
  const day = String(date.getDate()).padStart(2, '0')
  const hours = String(date.getHours()).padStart(2, '0')
  const minutes = String(date.getMinutes()).padStart(2, '0')
  return `${year}-${month}-${day} ${hours}:${minutes}`
}

// 获取时间范围
const getTimeRange = (range) => {
  const now = new Date()
  const today = new Date(now.getFullYear(), now.getMonth(), now.getDate())
  
  switch (range) {
    case 'today':
      return {
        startTime: formatLocalDateTime(today),
        endTime: formatLocalDateTime(now)
      }
    case 'yesterday':
      const yesterday = new Date(today)
      yesterday.setDate(today.getDate() - 1) // 昨天
      const yesterdayEnd = new Date(yesterday)
      yesterdayEnd.setHours(23, 59, 59, 999) // 设置为昨天23:59:59
      return {
        startTime: formatLocalDateTime(yesterday),
        endTime: formatLocalDateTime(yesterdayEnd)
      }
    case 'month':
      const monthStart = new Date(now.getFullYear(), now.getMonth(), 1)
      const monthEnd = new Date(now.getFullYear(), now.getMonth() + 1, 0) // 本月最后一天
      monthEnd.setHours(23, 59, 59, 999) // 设置为23:59:59
      return {
        startTime: formatLocalDateTime(monthStart),
        endTime: formatLocalDateTime(monthEnd)
      }
    case 'lastMonth':
      const lastMonthStart = new Date(now.getFullYear(), now.getMonth() - 1, 1)
      const lastMonthEnd = new Date(now.getFullYear(), now.getMonth(), 0, 23, 59, 59)
      return {
        startTime: formatLocalDateTime(lastMonthStart),
        endTime: formatLocalDateTime(lastMonthEnd)
      }
    default:
      return { startTime: '', endTime: '' }
  }
}

// 时间范围改变
const onTimeRangeChange = (value) => {
  if (value !== 'custom') {
    currentTimeRange.value = getTimeRange(value)
    loadData()
  }
}

// 自定义时间改变
const onCustomDateChange = (value) => {
  if (value && value.length === 2) {
    currentTimeRange.value = {
      startTime: value[0],
      endTime: value[1]
    }
    loadData()
  }
}

// 加载汇总数据
const loadSummaryData = async () => {
  try {
    const params = {}
    if (currentTimeRange.value.startTime) {
      params.startTime = currentTimeRange.value.startTime
    }
    if (currentTimeRange.value.endTime) {
      params.endTime = currentTimeRange.value.endTime
    }
    
    const res = await getAccessStatsSummary(params)
    if (res.code === 0) {
      summaryData.value = res.data
    }
  } catch (error) {
    console.error('加载汇总数据失败:', error)
  }
}

// 加载最近访问记录
const loadRecentAccess = async () => {
  recentLoading.value = true
  try {
    const params = {
      page: 1,
      pageSize: 10,
      orderBy: 'created_at',
      orderType: 'desc'
    }
    
    if (currentTimeRange.value.startTime) {
      params.startTime = currentTimeRange.value.startTime
    }
    if (currentTimeRange.value.endTime) {
      params.endTime = currentTimeRange.value.endTime
    }
    
    const res = await getAccessStatsList(params)
    if (res.code === 0) {
      recentAccessList.value = res.data.list || []
    }
  } catch (error) {
    console.error('加载最近访问记录失败:', error)
  } finally {
    recentLoading.value = false
  }
}

// 统一加载数据
const loadData = async () => {
  loading.value = true
  try {
    await Promise.all([
      loadSummaryData(),
      loadRecentAccess()
    ])
  } finally {
    loading.value = false
  }
}

// 刷新数据
const refreshData = () => {
  loadData()
}

// 刷新最近访问
const refreshRecentAccess = () => {
  loadRecentAccess()
}

// 生命周期钩子
onMounted(() => {
  // 初始化时间范围为昨天
  currentTimeRange.value = getTimeRange('yesterday')
  loadData()
  
  // 每30秒自动刷新数据
  refreshTimer.value = setInterval(() => {
    loadData()
  }, 30000)
})

onBeforeUnmount(() => {
  if (refreshTimer.value) {
    clearInterval(refreshTimer.value)
  }
})
</script>

<style scoped>
.dashboard-container {
  padding: 20px;
}

.time-filter-card {
  margin-bottom: 20px;
  border-radius: 8px;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
}

.time-filter-content {
  display: flex;
  align-items: center;
  justify-content: space-between;
  flex-wrap: wrap;
  gap: 15px;
}

.time-filter-left {
  display: flex;
  align-items: center;
  gap: 15px;
}

.filter-label {
  font-weight: bold;
  color: #303133;
  font-size: 14px;
}

.time-filter-right {
  display: flex;
  align-items: center;
  gap: 10px;
}

.time-filter-actions {
  display: flex;
  align-items: center;
  gap: 10px;
}

.time-range-display {
  margin-top: 15px;
  padding: 10px 15px;
  background-color: #f5f7fa;
  border-radius: 4px;
  border-left: 4px solid #409eff;
}

.range-label {
  font-weight: bold;
  color: #606266;
  margin-right: 8px;
}

.range-text {
  color: #303133;
  font-family: 'Courier New', monospace;
}

@media (max-width: 768px) {
  .time-filter-content {
    flex-direction: column;
    align-items: stretch;
  }
  
  .time-filter-left,
  .time-filter-right,
  .time-filter-actions {
    justify-content: center;
  }
}

.stats-cards {
  margin-bottom: 20px;
}

.stat-card {
  border-radius: 8px;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
}

.stat-content {
  display: flex;
  align-items: center;
  padding: 10px 0;
}

.stat-icon {
  width: 60px;
  height: 60px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-right: 20px;
  font-size: 24px;
  color: white;
}

.stat-icon.total-requests {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
}

.stat-icon.unique-ips {
  background: linear-gradient(135deg, #f093fb 0%, #f5576c 100%);
}

.stat-icon.response-time {
  background: linear-gradient(135deg, #4facfe 0%, #00f2fe 100%);
}

.stat-icon.success-rate {
  background: linear-gradient(135deg, #43e97b 0%, #38f9d7 100%);
}

.stat-info {
  flex: 1;
}

.stat-number {
  font-size: 28px;
  font-weight: bold;
  color: #303133;
  line-height: 1;
  margin-bottom: 5px;
}

.stat-label {
  font-size: 14px;
  color: #909399;
}

.charts-section {
  margin-bottom: 20px;
}

.chart-card {
  border-radius: 8px;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  font-weight: bold;
  color: #303133;
}

.chart-container {
  min-height: 200px;
  max-height: 300px;
  overflow-y: auto;
}

.api-path-item,
.country-item,
.device-item,
.browser-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 8px 0;
  border-bottom: 1px solid #f0f0f0;
}

.api-path-item:last-child,
.country-item:last-child,
.device-item:last-child,
.browser-item:last-child {
  border-bottom: none;
}

.api-path-name,
.country-name,
.device-name,
.browser-name {
  flex: 1;
  color: #606266;
  font-size: 14px;
}

.api-path-count,
.country-count,
.device-count,
.browser-count {
  color: #409eff;
  font-weight: bold;
  font-size: 14px;
}

.recent-access-card {
  border-radius: 8px;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
}

.el-empty {
  padding: 40px 0;
}
</style>