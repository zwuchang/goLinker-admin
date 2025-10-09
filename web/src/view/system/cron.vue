<template>
    <div class="cron-manager">
        <el-card class="box-card">
            <template #header>
                <div class="card-header">
                    <span>定时任务管理</span>
                    <div class="header-right">
                        <el-tooltip :content="wsStatus.tip" placement="top">
                            <el-tag :type="wsStatus.type" class="ws-status" :class="{ 'clickable': wsStatus.canRetry }"
                                @click="handleRetryClick">
                                {{ wsStatus.text }}
                            </el-tag>
                        </el-tooltip>
                        <el-button type="primary" @click="refreshData">
                            <el-icon>
                                <Refresh />
                            </el-icon>
                            刷新
                        </el-button>
                    </div>
                </div>
            </template>

            <el-table :data="cronList" v-loading="loading" border>
                <el-table-column prop="name" label="任务名称" min-width="120" />
                <el-table-column prop="description" label="任务描述" min-width="180" />
                <el-table-column prop="executeCount" label="执行次数" width="100" align="center" />
                <el-table-column label="执行规则" width="180">
                    <template #default="{ row }">
                        <div class="spec-cell" @click="startEdit(row)">
                            {{ row.spec }}
                            <el-icon>
                                <Edit />
                            </el-icon>
                        </div>
                    </template>
                </el-table-column>
                <el-table-column label="上次执行" width="180">
                    <template #default="{ row }">
                        <el-tooltip v-if="row.executeCount > 0" :content="formatTime(row.lastRunTime)" placement="top">
                            <span>{{ showLastRunTime(row) }}</span>
                        </el-tooltip>
                        <span v-else>{{ showLastRunTime(row) }}</span>
                    </template>
                </el-table-column>
                <el-table-column label="执行结果" min-width="200">
                    <template #default="{ row }">
                        <div class="result-cell">
                            <span>{{ formatResult(row.lastResult) }}</span>
                            <el-button v-if="row.lastResult && row.lastResult.length > 8" type="primary" link
                                @click="showDetail(row.lastResult)">
                                详情
                            </el-button>
                        </div>
                    </template>
                </el-table-column>
                <el-table-column label="下次执行" width="180">
                    <template #default="{ row }">
                        <el-tooltip :content="formatTime(row.nextRunTime)" placement="top">
                            <span>{{ row.timeUntilNext || '未安排' }}</span>
                        </el-tooltip>
                    </template>
                </el-table-column>
                <el-table-column label="状态" width="100" align="center">
                    <template #default="{ row }">
                        <el-tag :type="row.isRunning ? 'success' : 'danger'">
                            {{ row.isRunning ? '运行中' : '已停止' }}
                        </el-tag>
                    </template>
                </el-table-column>
                <el-table-column label="操作" width="120" fixed="right">
                    <template #default="{ row }">
                        <el-button :type="row.isRunning ? 'warning' : 'success'" size="small"
                            @click="toggleTaskStatus(row)">
                            {{ row.isRunning ? '暂停' : '启动' }}
                        </el-button>
                    </template>
                </el-table-column>
            </el-table>
        </el-card>

        <!-- 修改规则对话框 -->
        <el-dialog v-model="editDialog.visible" title="修改执行规则" width="500px">
            <el-form :model="editDialog.form" label-width="100px">
                <el-form-item label="执行规则">
                    <el-select v-model="editDialog.form.type" style="width: 100%">
                        <el-option label="每隔一段时间" value="interval" />
                        <el-option label="每天固定时间" value="daily" />
                        <el-option label="自定义规则" value="custom" />
                    </el-select>
                </el-form-item>

                <!-- 间隔执行 -->
                <template v-if="editDialog.form.type === 'interval'">
                    <el-form-item label="间隔时间">
                        <el-input-number v-model="editDialog.form.interval" :min="1" style="width: 150px" />
                        <el-select v-model="editDialog.form.unit" style="width: 100px; margin-left: 10px">
                            <el-option label="秒" value="s" />
                            <el-option label="分钟" value="m" />
                            <el-option label="小时" value="h" />
                        </el-select>
                    </el-form-item>
                </template>

                <!-- 每天固定时间 -->
                <template v-if="editDialog.form.type === 'daily'">
                    <el-form-item label="执行时间">
                        <el-time-picker v-model="editDialog.form.time" format="HH:mm:ss" placeholder="选择时间"
                            style="width: 100%" />
                    </el-form-item>
                </template>

                <!-- 自定义规则 -->
                <template v-if="editDialog.form.type === 'custom'">
                    <el-form-item label="Cron 表达式">
                        <el-input v-model="editDialog.form.customSpec" placeholder="如: */5 * * * *" />
                    </el-form-item>
                </template>
            </el-form>
            <template #footer>
                <el-button @click="editDialog.visible = false">取消</el-button>
                <el-button type="primary" @click="confirmUpdateSpec">确定</el-button>
            </template>
        </el-dialog>
    </div>
</template>

<script setup>
import { ref, onMounted, onBeforeUnmount, watch ,nextTick} from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { getCronStatus, updateCronSpec, startCron, stopCron } from '@/api/cron'
import { Refresh, Edit } from '@element-plus/icons-vue'
import { useWebSocket } from '@/hooks/useWebSocket'

defineOptions({
    name: 'CronManager'
})

const loading = ref(false)
const cronList = ref([])
const ws = useWebSocket()

const wsStatus = ref({
    type: 'info',
    text: '连接中...',
    tip: '正在连接 WebSocket 服务',
    canRetry: false
})

const retryTimer = ref(null)

const requestCronStatus = () => {
    if (ws.isConnected.value) {
        ws.sendMessage({
            type: 'cron',
            content: null
        })
    }
}

const handleRetryClick = () => {
    if (wsStatus.value.canRetry) {
        retryWebSocket()
    }
}

// 手动重连
const retryWebSocket = () => {
    wsStatus.value = {
        type: 'warning',
        text: '连接中...',
        tip: '正在重新连接 WebSocket 服务',
        canRetry: false
    }
    ws.reconnect()
}


// 监听 WebSocket 状态变化
watch(() => ws.isConnected.value, (connected) => {
    // 清除之前的定时器
    if (retryTimer.value) {
        clearTimeout(retryTimer.value)
        retryTimer.value = null
    }
    
    console.log('WebSocket 连接状态变化:', connected)
    
    if (connected) {
        wsStatus.value = {
            type: 'success',
            text: '实时更新',
            tip: 'WebSocket 已连接',
            canRetry: false
        }
        requestCronStatus()
    } else {
        wsStatus.value = {
            type: 'warning',
            text: '连接中...',
            tip: '正在尝试连接 WebSocket 服务',
            canRetry: false
        }
        
        // 使用 nextTick 确保状态更新后再设置定时器
        nextTick(() => {
            retryTimer.value = setTimeout(() => {
                console.log('检查 WebSocket 状态:', ws.isConnected.value)
                if (!ws.isConnected.value) {
                    console.log('设置为可重连状态')
                    wsStatus.value = {
                        type: 'danger',
                        text: '点击重连',
                        tip: 'WebSocket 连接失败，点击重试',
                        canRetry: true
                    }
                }
            }, 5000)
        })
    }
}, { immediate: true }) // 添加 immediate: true 确保初始化时执行



const editDialog = ref({
    visible: false,
    currentRow: null,
    form: {
        type: 'interval',
        interval: 1,
        unit: 'h',
        time: null,
        customSpec: ''
    }
})

const formatDuration = (duration) => {
    if (!duration) return '0ms'
    duration = duration.replace(/"/g, '')
    return duration
}

const formatResult = (text) => {
    if (!text) return '暂无结果'
    return text.length > 8 ? text.slice(0, 8) + '...' : text
}

const showDetail = (content) => {
    ElMessageBox.alert(content, '执行结果详情', {
        confirmButtonText: '确定'
    })
}

const showLastRunTime = (row) => {
    if (!row.timeSinceLast) {
        return '从未执行'
    }
    return `${row.timeSinceLast}前， 耗时:${formatDuration(row.duration)}`
}

const handleWsMessage = (data) => {
    if (data.type === 'cron_list') {
        cronList.value = data.content.map(item => ({
            ...item,
            editing: false,
            newSpec: item.spec
        }))
        loading.value = false
    }
}

const refreshData = async () => {
    loading.value = true
    try {
        if (ws.isConnected.value) {
            requestCronStatus()
        } else {
            const { data } = await getCronStatus()
            cronList.value = data.map(item => ({
                ...item,
                editing: false,
                newSpec: item.spec
            }))
            loading.value = false
        }
    } catch (error) {
        ElMessage.error('获取定时任务状态失败')
        loading.value = false
    }
}

const startEdit = (row) => {
    editDialog.value.visible = true
    editDialog.value.currentRow = row

    const spec = row.spec
    if (spec.startsWith('@every')) {
        const match = spec.match(/@every\s+(\d+)([smh])/)
        if (match) {
            editDialog.value.form = {
                type: 'interval',
                interval: parseInt(match[1]),
                unit: match[2],
                time: null,
                customSpec: ''
            }
            return
        }
    } else if (spec.split(' ').length === 6) {
        const [second, minute, hour, day, month, week] = spec.split(' ')
        if (day === '*' && month === '*' && week === '*') {
            const time = new Date()
            time.setHours(parseInt(hour))
            time.setMinutes(parseInt(minute))
            time.setSeconds(parseInt(second))

            editDialog.value.form = {
                type: 'daily',
                interval: 1,
                unit: 'h',
                time,
                customSpec: ''
            }
            return
        }
    }

    editDialog.value.form = {
        type: 'custom',
        interval: 1,
        unit: 'h',
        time: null,
        customSpec: spec
    }
}

const generateSpec = (form) => {
    switch (form.type) {
        case 'interval':
            return `@every ${form.interval}${form.unit}`
        case 'daily':
            if (!form.time) return ''
            const time = new Date(form.time)
            return `${time.getSeconds()} ${time.getMinutes()} ${time.getHours()} * * *`
        case 'custom':
            return form.customSpec
        default:
            return ''
    }
}

const confirmUpdateSpec = async () => {
    const spec = generateSpec(editDialog.value.form)
    if (!spec) {
        ElMessage.warning('请设置正确的执行规则')
        return
    }

    try {
        const row = editDialog.value.currentRow
        const result = await updateCronSpec({
            name: row.name,
            spec,
            description: row.description
        })
        editDialog.value.visible = false
        if (result.code == 0) {
            ElMessage.success('更新成功')
        }
        await refreshData()
    } catch (error) {
        ElMessage.error('更新失败')
    }
}

const toggleTaskStatus = async (row) => {
    try {
        let result
        let msg
        if (row.isRunning) {
            result = await stopCron({ name: row.name })
            msg = '任务已暂停'
        } else {
            result = await startCron({ name: row.name })
            msg = '任务已启动'
        }
        if (result.code == 0) {
            ElMessage.success(msg)
            await refreshData()
        }
    } catch (error) {
        ElMessage.error(row.isRunning ? '暂停失败' : '启动失败')
    }
}

const formatTime = (time) => {
    if (!time) return ''
    return new Date(time).toLocaleString()
}

let statusTimer = null

onMounted(() => {
    ws.onMessage(handleWsMessage)
    refreshData()
    statusTimer = setInterval(() => {
        requestCronStatus()
    }, 1000)
})

onBeforeUnmount(() => {
    if (statusTimer) {
        clearInterval(statusTimer)
    }
    if (retryTimer.value) {
        clearTimeout(retryTimer.value)
    }
    ws.offMessage(handleWsMessage)
})
</script>

<style scoped>
.cron-manager {
    padding: 20px;
}

.card-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
}

.header-right {
    display: flex;
    align-items: center;
    gap: 10px;
}

.ws-status {
    font-size: 12px;
}

.spec-cell {
    display: flex;
    align-items: center;
    justify-content: space-between;
    cursor: pointer;
}

.spec-cell:hover {
    color: var(--el-color-primary);
}

.result-cell {
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 0 8px;
}

.ws-status {
    font-size: 12px;
    cursor: pointer;
    transition: all 0.3s;
}

.ws-status:hover {
    opacity: 0.8;
}
</style>