<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="searchForm" :inline="true" :model="searchInfo">
        <el-form-item label="网站标题">
          <el-input v-model="searchInfo.website_title" placeholder="网站标题" />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchInfo.status" clearable placeholder="请选择">
            <el-option label="启用" :value="1" />
            <el-option label="禁用" :value="0" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">
            查询
          </el-button>
          <el-button icon="refresh" @click="onReset"> 重置 </el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDialog">
          新增
        </el-button>
        <el-button icon="delete" :disabled="!gameConfigs.length" @click="onDelete">
          删除
        </el-button>
      </div>
      <el-table
        :data="tableData"
        @sort-change="sortChange"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column
          align="left"
          label="id"
          min-width="60"
          prop="ID"
          sortable="custom"
        />
        <el-table-column
          align="left"
          label="网站标题"
          min-width="150"
          prop="website_title"
          sortable="custom"
        />
        <el-table-column
          align="left"
          label="网站描述"
          min-width="200"
          prop="website_desc"
          show-overflow-tooltip
        />
        <el-table-column
          align="left"
          label="下载地址"
          min-width="200"
          prop="download_url"
          show-overflow-tooltip
        />
        <el-table-column
          align="left"
          label="音频地址"
          min-width="200"
          prop="audio_url"
          show-overflow-tooltip
        />
        <el-table-column
          align="left"
          label="市场Logo"
          min-width="200"
          prop="market_logo"
          show-overflow-tooltip
        />
        <el-table-column
          align="left"
          label="状态"
          min-width="100"
          prop="status"
          sortable="custom"
        >
          <template #default="scope">
            <el-tag :type="scope.row.status === 1 ? 'success' : 'danger'">
              {{ scope.row.status === 1 ? '启用' : '禁用' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="创建时间"
          min-width="180"
          prop="CreatedAt"
          sortable="custom"
        >
          <template #default="scope">
            {{ formatDate(scope.row.CreatedAt) }}
          </template>
        </el-table-column>

        <el-table-column align="left" fixed="right" label="操作" :min-width="appStore.operateMinWith">
          <template #default="scope">
            <el-button
              icon="edit"
              type="primary"
              link
              @click="editGameConfigFunc(scope.row)"
            >
              编辑
            </el-button>
            <el-button
              icon="delete"
              type="primary"
              link
              @click="deleteGameConfigFunc(scope.row)"
            >
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          layout="total, sizes, prev, pager, next, jumper"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>

    <!-- 表单弹窗 -->
    <el-dialog v-model="dialogFormVisible" :title="dialogTitle" width="800px" destroy-on-close>
      <GameConfigForm
        :id="formId"
        :edit-data="editData"
        @close="dialogFormVisible = false"
        @success="getTableData"
      />
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { getGameConfigList, deleteGameConfig } from '@/api/gameConfig'
import { ElMessage, ElMessageBox } from 'element-plus'
import { useAppStore } from '@/pinia/modules/app'
import { toSQLLine } from '@/utils/stringFun'
import { formatDate } from '@/utils/format'
import GameConfigForm from './form.vue'

const appStore = useAppStore()

const searchForm = ref()
const tableData = ref([])
const page = ref(1)
const pageSize = ref(10)
const total = ref(0)
const gameConfigs = ref([])

const searchInfo = reactive({
  website_title: '',
  status: null
})

const dialogFormVisible = ref(false)
const dialogTitle = ref('')
const formId = ref(0)
const editData = ref({})

// 获取表格数据
const getTableData = async () => {
  const params = {
    page: page.value,
    pageSize: pageSize.value,
    website_title: searchInfo.website_title || '',
    status: searchInfo.status
  }
  // 过滤掉空值
  Object.keys(params).forEach(key => {
    if (params[key] === null || params[key] === undefined || params[key] === '') {
      delete params[key]
    }
  })
  const res = await getGameConfigList(params)
  if (res.code === 0) {
    tableData.value = res.data.list
    total.value = res.data.total
  }
}

// 搜索
const onSubmit = () => {
  page.value = 1
  getTableData()
}

// 重置搜索
const onReset = () => {
  searchInfo.website_title = ''
  searchInfo.status = null
  page.value = 1
  getTableData()
}

// 分页
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 排序
const sortChange = ({ prop, order }) => {
  if (prop) {
    if (order === 'ascending') {
      searchInfo.orderKey = toSQLLine(prop)
      searchInfo.desc = false
    } else if (order === 'descending') {
      searchInfo.orderKey = toSQLLine(prop)
      searchInfo.desc = true
    } else {
      searchInfo.orderKey = undefined
      searchInfo.desc = undefined
    }
    getTableData()
  }
}

// 选择变化
const handleSelectionChange = (val) => {
  gameConfigs.value = val
}

// 打开对话框
const openDialog = () => {
  dialogTitle.value = '新增游戏配置'
  formId.value = 0
  editData.value = {}
  dialogFormVisible.value = true
}

// 编辑
const editGameConfigFunc = (row) => {
  dialogTitle.value = '编辑游戏配置'
  formId.value = row.ID
  editData.value = { ...row }
  dialogFormVisible.value = true
}

// 删除
const deleteGameConfigFunc = async (row) => {
  try {
    await ElMessageBox.confirm('此操作将永久删除该配置，是否继续？', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    })
    const res = await deleteGameConfig({ ID: row.ID })
    if (res.code === 0) {
      ElMessage.success('删除成功')
      getTableData()
    }
  } catch {
    ElMessage.info('已取消删除')
  }
}

// 批量删除
const onDelete = async () => {
  if (gameConfigs.value.length === 0) {
    ElMessage.warning('请选择要删除的数据')
    return
  }
  try {
    await ElMessageBox.confirm('此操作将永久删除选中的配置，是否继续？', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    })
    for (const item of gameConfigs.value) {
      await deleteGameConfig({ ID: item.ID })
    }
    ElMessage.success('删除成功')
    getTableData()
  } catch {
    ElMessage.info('已取消删除')
  }
}

onMounted(() => {
  getTableData()
})
</script>

<style lang="scss" scoped>

</style>
