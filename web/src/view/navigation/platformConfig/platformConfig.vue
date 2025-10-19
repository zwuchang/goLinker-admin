<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="searchForm" :inline="true" :model="searchInfo">
        <el-form-item label="平台名称">
          <el-input v-model="searchInfo.platform_name" placeholder="平台名称" />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchInfo.status" clearable placeholder="请选择">
            <el-option label="启用" :value="1" />
            <el-option label="禁用" :value="0" />
          </el-select>
        </el-form-item>
        <el-form-item label="可见性">
          <el-select v-model="searchInfo.is_visible" clearable placeholder="请选择">
            <el-option label="可见" :value="1" />
            <el-option label="隐藏" :value="0" />
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
        <el-button icon="delete" :disabled="!platformConfigs.length" @click="onDelete">
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
          prop="id"
          sortable="custom"
        />
        <el-table-column
          align="left"
          label="平台名称"
          min-width="150"
          prop="platform_name"
          sortable="custom"
        />
        <el-table-column align="left" label="平台图标" prop="platform_icon" width="100">
          <template #default="scope">
            <el-image
              v-if="scope.row.platform_icon"
              :src="scope.row.platform_icon"
              style="width: 40px; height: 40px"
              fit="cover"
            />
            <span v-else>-</span>
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="平台接口"
          min-width="200"
          prop="platform_api"
          show-overflow-tooltip
        />
        <el-table-column
          align="left"
          label="排序"
          min-width="100"
          prop="sort"
          sortable="custom"
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
          label="可见性"
          min-width="100"
          prop="is_visible"
          sortable="custom"
        >
          <template #default="scope">
            <el-tag :type="scope.row.is_visible === 1 ? 'success' : 'info'">
              {{ scope.row.is_visible === 1 ? '可见' : '隐藏' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="创建时间"
          min-width="180"
          prop="created_at"
          sortable="custom"
        >
          <template #default="scope">
            {{ formatDate(scope.row.created_at) }}
          </template>
        </el-table-column>

        <el-table-column align="left" fixed="right" label="操作" :min-width="appStore.operateMinWith">
          <template #default="scope">
            <el-button
              icon="edit"
              type="primary"
              link
              @click="editPlatformConfigFunc(scope.row)"
            >
              编辑
            </el-button>
            <el-button
              icon="delete"
              type="primary"
              link
              @click="deletePlatformConfigFunc(scope.row)"
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
    <el-dialog v-model="dialogFormVisible" :title="dialogTitle" width="600px">
      <PlatformConfigForm
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
import { getPlatformConfigList, deletePlatformConfig } from '@/api/platformConfig'
import { ElMessage, ElMessageBox } from 'element-plus'
import { useAppStore } from '@/pinia/modules/app'
import { toSQLLine } from '@/utils/stringFun'
import PlatformConfigForm from './form.vue'
import { formatDate } from '@/utils/format'

defineOptions({
  name: 'PlatformConfig'
})

const appStore = useAppStore()

const searchForm = ref()
const tableData = ref([])
const page = ref(1)
const pageSize = ref(10)
const total = ref(0)
const platformConfigs = ref([])

const searchInfo = reactive({
  platform_name: '',
  status: null,
  is_visible: null
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
    platform_name: searchInfo.platform_name || '',
    status: searchInfo.status,
    is_visible: searchInfo.is_visible
  }
  // 过滤掉空值
  Object.keys(params).forEach(key => {
    if (params[key] === null || params[key] === undefined || params[key] === '') {
      delete params[key]
    }
  })
  const res = await getPlatformConfigList(params)
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
  searchInfo.platform_name = ''
  searchInfo.status = null
  searchInfo.is_visible = null
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
  page.value = 1
  getTableData()
}

// 排序
const sortChange = ({ prop, order }) => {
  if (prop) {
    searchInfo.orderKey = toSQLLine(prop)
    searchInfo.desc = order === 'descending'
  }
  getTableData()
}

// 选择
const handleSelectionChange = (val) => {
  platformConfigs.value = val
}

// 新增
const openDialog = () => {
  dialogTitle.value = '新增平台配置'
  formId.value = 0
  editData.value = {}
  dialogFormVisible.value = true
}

// 编辑
const editPlatformConfigFunc = (row) => {
  dialogTitle.value = '编辑平台配置'
  formId.value = row.id
  editData.value = { ...row }
  dialogFormVisible.value = true
}

// 删除
const deletePlatformConfigFunc = async (row) => {
  ElMessageBox.confirm('此操作将永久删除该平台配置, 是否继续?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    const res = await deletePlatformConfig({ id: row.id })
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '删除成功'
      })
      getTableData()
    }
  })
}

// 批量删除
const onDelete = async () => {
  if (platformConfigs.value.length === 0) {
    ElMessage({
      type: 'warning',
      message: '请选择要删除的数据'
    })
    return
  }
  ElMessageBox.confirm('此操作将永久删除选中的平台配置, 是否继续?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    const ids = platformConfigs.value.map(item => item.id)
    for (const id of ids) {
      await deletePlatformConfig({ id })
    }
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    getTableData()
  })
}

onMounted(() => {
  getTableData()
})
</script>
