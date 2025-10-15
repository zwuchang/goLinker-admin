<template>
  <div class="dashboard-container">
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
const refreshTimer = ref(null)

// 加载汇总数据
const loadSummaryData = async () => {
  try {
    const res = await getAccessStatsSummary({})
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
    const res = await getAccessStatsList({
      page: 1,
      pageSize: 10,
      orderBy: 'created_at',
      orderType: 'desc'
    })
    if (res.code === 0) {
      recentAccessList.value = res.data.list || []
    }
  } catch (error) {
    console.error('加载最近访问记录失败:', error)
  } finally {
    recentLoading.value = false
  }
}

// 刷新数据
const refreshData = () => {
  loadSummaryData()
}

// 刷新最近访问
const refreshRecentAccess = () => {
  loadRecentAccess()
}

// 生命周期钩子
onMounted(() => {
  loadSummaryData()
  loadRecentAccess()
  // 每30秒自动刷新数据
  refreshTimer.value = setInterval(() => {
    loadSummaryData()
    loadRecentAccess()
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