<template>
  <div class="platform-ranking-container">
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="平台名称">
          <el-input v-model="searchInfo.platformName" placeholder="请输入平台名称" />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchInfo.status" placeholder="请选择状态" clearable>
            <el-option label="启用" :value="1" />
            <el-option label="禁用" :value="0" />
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
        <el-button type="primary" icon="plus" @click="openDialog">新增平台排行榜</el-button>
        <el-button type="danger" icon="delete" @click="onDelete" :disabled="!multipleSelection.length">批量删除</el-button>
      </div>
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="排名" prop="rank" width="80" />
        <el-table-column align="left" label="平台名称" prop="platformName" width="150" />
        <el-table-column align="left" label="Logo" width="100">
          <template #default="scope">
            <el-image
              v-if="scope.row.logo"
              :src="scope.row.logo"
              :preview-src-list="[scope.row.logo]"
              style="width: 40px; height: 40px"
              fit="cover"
            />
            <span v-else>无</span>
          </template>
        </el-table-column>
        <el-table-column align="left" label="评分" prop="rating" width="80" />
        <el-table-column align="left" label="特色功能" prop="features" width="150" />
        <el-table-column align="left" label="功能类型" prop="featureType" width="100" />
        <el-table-column align="left" label="是否新平台" width="100">
          <template #default="scope">
            <el-tag v-if="scope.row.isNew === 1" type="success">是</el-tag>
            <el-tag v-else type="info">否</el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="状态" width="80">
          <template #default="scope">
            <el-tag v-if="scope.row.status === 1" type="success">启用</el-tag>
            <el-tag v-else type="danger">禁用</el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="是否显示" width="100">
          <template #default="scope">
            <el-tag v-if="scope.row.isVisible === 1" type="success">显示</el-tag>
            <el-tag v-else type="info">隐藏</el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="排序" prop="sort" width="80" />

        <el-table-column align="left" label="创建时间" prop="CreatedAt" width="150">
          <template #default="scope">
            {{ formatDate(scope.row.CreatedAt) }}
          </template>
        </el-table-column>
        <el-table-column align="left" label="操作" width="200">
          <template #default="scope">
            <el-button type="primary" link icon="edit" size="small" @click="updatePlatformRanking(scope.row)">变更</el-button>
            <el-button type="primary" link icon="delete" size="small" @click="deletePlatformRanking(scope.row)">删除</el-button>
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
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="平台排行榜" destroy-on-close>
      <el-form
        ref="elFormRef"
        :model="form"
        :rules="rules"
        label-width="120px"
        label-position="left"
      >
        <el-form-item label="排名" prop="rank">
          <el-input-number v-model="form.rank" :min="1" placeholder="请输入排名" />
        </el-form-item>
        <el-form-item label="平台名称" prop="platformName">
          <el-input v-model="form.platformName" placeholder="请输入平台名称" />
        </el-form-item>
        <el-form-item label="Logo地址" prop="logo">
          <el-input v-model="form.logo" placeholder="请输入Logo地址" />
          <el-image
            v-if="form.logo"
            :src="form.logo"
            style="width: 100px; height: 100px; margin-top: 10px"
            fit="cover"
          />
        </el-form-item>
        <el-form-item label="评分" prop="rating">
          <el-input-number v-model="form.rating" :min="0" :max="5" :precision="1" placeholder="请输入评分" />
        </el-form-item>
        <el-form-item label="特色功能" prop="features">
          <el-input v-model="form.features" placeholder="请输入特色功能描述" />
        </el-form-item>
        <el-form-item label="功能类型" prop="featureType">
          <el-input v-model="form.featureType" placeholder="请输入功能类型" />
        </el-form-item>
        <el-form-item label="访问链接" prop="visitUrl">
          <el-input v-model="form.visitUrl" placeholder="请输入访问链接" />
        </el-form-item>
        <el-form-item label="是否新平台" prop="isNew">
          <el-radio-group v-model="form.isNew">
            <el-radio :label="1">是</el-radio>
            <el-radio :label="0">否</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="form.status">
            <el-radio :label="1">启用</el-radio>
            <el-radio :label="0">禁用</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="是否显示" prop="isVisible">
          <el-radio-group v-model="form.isVisible">
            <el-radio :label="1">显示</el-radio>
            <el-radio :label="0">隐藏</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="排序" prop="sort">
          <el-input-number v-model="form.sort" :min="0" placeholder="请输入排序值" />
        </el-form-item>

      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取 消</el-button>
          <el-button type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  getPlatformRankingList,
  createPlatformRanking,
  updatePlatformRanking as updatePlatformRankingApi,
  deletePlatformRanking as deletePlatformRankingApi
} from '@/api/platformRanking'
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'

defineOptions({
  name: 'PlatformRanking'
})

const page = ref(1)
const pageSize = ref(10)
const total = ref(0)
const tableData = ref([])
const searchInfo = ref({})
const multipleSelection = ref([])

const dialogFormVisible = ref(false)
const form = reactive({
  rank: 1,
  platformName: '',
  logo: '',
  rating: 0,
  features: '',
  featureType: '',
  visitUrl: '',
  isNew: 0,
  status: 1,
  isVisible: 1,
  sort: 0,
})

const rules = reactive({
  rank: [{ required: true, message: '请输入排名', trigger: 'blur' }],
  platformName: [{ required: true, message: '请输入平台名称', trigger: 'blur' }]
})

const elFormRef = ref()

const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  getTableData()
}

const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

const getTableData = async() => {
  const table = await getPlatformRankingList({
    page: page.value,
    pageSize: pageSize.value,
    ...searchInfo.value
  })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

const openDialog = () => {
  resetForm()
  dialogFormVisible.value = true
}

const closeDialog = () => {
  dialogFormVisible.value = false
  resetForm()
}

const resetForm = () => {
  Object.assign(form, {
    rank: 1,
    platformName: '',
    logo: '',
    rating: 0,
    features: '',
    featureType: '',
    visitUrl: '',
    isNew: 0,
    status: 1,
    isVisible: 1,
    sort: 0,
  })
}

const enterDialog = async() => {
  elFormRef.value?.validate(async(valid) => {
    if (!valid) return
    let res
    if (form.ID) {
      res = await updatePlatformRankingApi(form)
    } else {
      res = await createPlatformRanking(form)
    }
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: form.ID ? '更新成功' : '创建成功'
      })
      closeDialog()
      getTableData()
    }
  })
}

const updatePlatformRanking = (row) => {
  Object.assign(form, row)
  dialogFormVisible.value = true
}

const deletePlatformRanking = async(row) => {
  ElMessageBox.confirm('此操作将永久删除该平台排行榜，是否继续？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
    const res = await deletePlatformRankingApi({ id: row.ID })
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '删除成功'
      })
      getTableData()
    }
  })
}

const onDelete = async() => {
  ElMessageBox.confirm('此操作将永久删除所选平台排行榜，是否继续？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
    const ids = multipleSelection.value.map(item => item.ID)
    for (const id of ids) {
      await deletePlatformRankingApi({ id })
    }
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    getTableData()
  })
}

const formatDate = (date) => {
  if (!date) return ''
  return new Date(date).toLocaleString()
}

onMounted(() => {
  getTableData()
})
</script>

<style scoped>
.platform-ranking-container {
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
  text-align: right;
}

.dialog-footer {
  text-align: right;
}
</style>
