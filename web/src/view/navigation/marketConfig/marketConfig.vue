<template>
  <div class="market-config-container">
    <div class="gva-search-box">
      <el-form ref="searchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
        <el-form-item label="名称" prop="name">
          <el-input v-model="searchInfo.name" placeholder="请输入名称" clearable />
        </el-form-item>
        <el-form-item label="类型" prop="type">
          <el-select v-model="searchInfo.type" placeholder="请选择类型" clearable>
            <el-option label="可关闭" :value="1" />
            <el-option label="类型2" :value="2" />
            <el-option label="类型3" :value="3" />
          </el-select>
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-select v-model="searchInfo.status" placeholder="请选择状态" clearable>
            <el-option label="启用" :value="1" />
            <el-option label="禁用" :value="0" />
          </el-select>
        </el-form-item>
        <el-form-item label="是否可见" prop="is_visible">
          <el-select v-model="searchInfo.is_visible" placeholder="请选择是否可见" clearable>
            <el-option label="可见" :value="1" />
            <el-option label="隐藏" :value="0" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDialog">新增配置</el-button>
        <el-button type="danger" icon="delete" @click="onDelete">批量删除</el-button>
      </div>
      <el-table
        ref="multipleTable"
        :data="tableData"
        style="width: 100%"
        @selection-change="handleSelectionChange"
        @sort-change="sortChange"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="ID" prop="id" width="80" />
        <el-table-column align="left" label="名称" prop="name" width="150" />
        <el-table-column align="left" label="Logo" prop="logo" width="120">
          <template #default="scope">
            <el-image
              v-if="scope.row.logo"
              :src="scope.row.logo"
              :preview-src-list="[scope.row.logo]"
              style="width: 40px; height: 40px; border-radius: 4px;"
              fit="cover"
            />
            <span v-else>-</span>
          </template>
        </el-table-column>
        <el-table-column align="left" label="跳转地址" prop="jump_url" width="200" show-overflow-tooltip />
        <el-table-column align="left" label="类型" prop="type" width="100">
          <template #default="scope">
            <el-tag :type="getTypeTagType(scope.row.type)">
              {{ getTypeText(scope.row.type) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="状态" prop="status" width="80">
          <template #default="scope">
            <el-tag :type="scope.row.status === 1 ? 'success' : 'danger'">
              {{ scope.row.status === 1 ? '启用' : '禁用' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="是否可见" prop="is_visible" width="100">
          <template #default="scope">
            <el-tag :type="scope.row.is_visible === 1 ? 'success' : 'info'">
              {{ scope.row.is_visible === 1 ? '可见' : '隐藏' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="排序" prop="sort" width="80" />
        <el-table-column align="left" label="描述" prop="description" width="200" show-overflow-tooltip />
        <el-table-column align="left" label="创建时间" prop="created_at" width="160" />
        <el-table-column align="left" label="操作" width="160">
          <template #default="scope">
            <el-button type="primary" link icon="edit" @click="editMarketConfigFunc(scope.row)">编辑</el-button>
            <el-button type="primary" link icon="delete" @click="deleteMarketConfigFunc(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          layout="total, sizes, prev, pager, next, jumper"
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>

    <!-- 表单弹窗 -->
    <el-dialog v-model="dialogFormVisible" :title="dialogTitle" width="600px" destroy-on-close>
      <MarketConfigForm
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
import { getMarketConfigList, deleteMarketConfig } from '@/api/marketConfig'
import { ElMessage, ElMessageBox } from 'element-plus'
import { toSQLLine } from '@/utils/stringFun'
import MarketConfigForm from './form.vue'

defineOptions({
  name: 'MarketConfig'
})

const searchFormRef = ref()
const tableData = ref([])
const page = ref(1)
const pageSize = ref(10)
const total = ref(0)
const marketConfigs = ref([])

const searchInfo = reactive({
  name: '',
  type: null,
  status: null,
  is_visible: null
})

const dialogFormVisible = ref(false)
const dialogTitle = ref('')
const formId = ref(0)
const editData = ref({})

// 获取类型标签类型
const getTypeTagType = (type) => {
  const typeMap = {
    1: 'primary',
    2: 'success',
    3: 'warning'
  }
  return typeMap[type] || 'info'
}

// 获取类型文本
const getTypeText = (type) => {
  const typeMap = {
    1: '类型1',
    2: '类型2',
    3: '类型3'
  }
  return typeMap[type] || '未知'
}

// 获取表格数据
const getTableData = async () => {
  const params = {
    page: page.value,
    pageSize: pageSize.value,
    name: searchInfo.name || '',
    type: searchInfo.type,
    status: searchInfo.status,
    is_visible: searchInfo.is_visible
  }
  // 过滤掉空值
  Object.keys(params).forEach(key => {
    if (params[key] === null || params[key] === undefined || params[key] === '') {
      delete params[key]
    }
  })
  const res = await getMarketConfigList(params)
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
  searchInfo.name = ''
  searchInfo.type = null
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
  marketConfigs.value = val
}

// 新增
const openDialog = () => {
  dialogTitle.value = '新增配置'
  formId.value = 0
  editData.value = {}
  dialogFormVisible.value = true
}

// 编辑
const editMarketConfigFunc = (row) => {
  dialogTitle.value = '编辑配置'
  formId.value = row.id
  editData.value = { ...row }
  dialogFormVisible.value = true
}

// 删除
const deleteMarketConfigFunc = async (row) => {
  ElMessageBox.confirm('此操作将永久删除该配置, 是否继续?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    const res = await deleteMarketConfig({ id: row.id })
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
  if (marketConfigs.value.length === 0) {
    ElMessage({
      type: 'warning',
      message: '请选择要删除的数据'
    })
    return
  }
  ElMessageBox.confirm('此操作将永久删除选中的配置, 是否继续?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    const ids = marketConfigs.value.map(item => item.id)
    for (const id of ids) {
      await deleteMarketConfig({ id })
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

<style scoped>
.market-config-container {
  padding: 20px;
}

.gva-search-box {
  margin-bottom: 20px;
}

.gva-table-box {
  background: white;
  border-radius: 8px;
  padding: 20px;
}

.gva-btn-list {
  margin-bottom: 20px;
}

.gva-pagination {
  margin-top: 20px;
  display: flex;
  justify-content: flex-end;
}
</style>
