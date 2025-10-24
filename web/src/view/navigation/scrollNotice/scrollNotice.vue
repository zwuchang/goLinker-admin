<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="通知标题">
          <el-input v-model="searchInfo.title" placeholder="请输入通知标题" />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchInfo.status" placeholder="请选择状态" clearable>
            <el-option label="启用" :value="1" />
            <el-option label="禁用" :value="0" />
          </el-select>
        </el-form-item>
        <el-form-item label="是否显示">
          <el-select v-model="searchInfo.isVisible" placeholder="请选择显示状态" clearable>
            <el-option label="显示" :value="1" />
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
        <el-button type="primary" icon="plus" @click="openDialog">新增滚动通知</el-button>
        <el-button type="danger" icon="delete" @click="onDelete" :disabled="!multipleSelection.length">批量删除</el-button>
      </div>
      <el-table ref="multipleTable" style="width: 100%" tooltip-effect="dark" :data="tableData" row-key="ID"
        @selection-change="handleSelectionChange">
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="通知标题" prop="title" width="200" />
        <el-table-column align="left" label="通知内容" prop="content" width="300" show-overflow-tooltip />
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
        <el-table-column align="left" label="开始时间" prop="startTime" width="150">
          <template #default="scope">
            {{ formatDate(scope.row.startTime) }}
          </template>
        </el-table-column>
        <el-table-column align="left" label="结束时间" prop="endTime" width="150">
          <template #default="scope">
            {{ formatDate(scope.row.endTime) }}
          </template>
        </el-table-column>
        <el-table-column align="left" label="创建时间" prop="CreatedAt" width="150">
          <template #default="scope">
            {{ formatDate(scope.row.CreatedAt) }}
          </template>
        </el-table-column>
        <el-table-column align="left" label="操作" :width="appStore.operateMinWith" fixed="right">
          <template #default="scope">
            <el-button type="primary" link icon="edit" @click="editScrollNotice(scope.row)">编辑</el-button>
            <el-button type="primary" link icon="delete" @click="deleteScrollNotice(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination layout="total, sizes, prev, pager, next, jumper" :current-page="page" :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]" :total="total" @current-change="handleCurrentChange"
          @size-change="handleSizeChange" />
      </div>

      <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="滚动通知" destroy-on-close>
      <div class="gva-table-box">
        <div class="gva-form-box">
          <el-form ref="elFormRef" :model="form" :rules="rules" label-width="120px" label-position="left">
            <el-form-item label="通知标题" prop="title">
              <el-input v-model="form.title" placeholder="请输入通知标题" />
            </el-form-item>
            <el-form-item label="通知内容" prop="content">
              <el-input v-model="form.content" type="textarea" :rows="4" placeholder="请输入通知内容" />
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
            <el-form-item label="开始时间" prop="startTime">
              <el-date-picker v-model="form.startTime" type="datetime" placeholder="请选择开始时间（可选）"
                format="YYYY-MM-DD HH:mm:ss" value-format="YYYY-MM-DD HH:mm:ss" style="width: 100%" />
            </el-form-item>
            <el-form-item label="结束时间" prop="endTime">
              <el-date-picker v-model="form.endTime" type="datetime" placeholder="请选择结束时间（可选）"
                format="YYYY-MM-DD HH:mm:ss" value-format="YYYY-MM-DD HH:mm:ss" style="width: 100%" />
            </el-form-item>
          </el-form>
        </div>
      </div>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取 消</el-button>
          <el-button type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
    
    </div>

  </div>

</template>

<script setup>
import {
  getNoticeList,
  createNotice,
  updateNotice as updateNoticeApi,
  deleteNotice as deleteNoticeApi
} from '@/api/notice'
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { formatDate } from '@/utils/format'
import { useAppStore } from '@/pinia/modules/app'

defineOptions({
  name: 'ScrollNotice'
})

const appStore = useAppStore()

const page = ref(1)
const pageSize = ref(10)
const total = ref(0)
const tableData = ref([])
const searchInfo = ref({})
const multipleSelection = ref([])

const dialogFormVisible = ref(false)
const form = reactive({
  title: '',
  content: '',
  status: 1,
  isVisible: 1,
  sort: 0,
  startTime: '',
  endTime: ''
})

const rules = reactive({
  title: [{ required: true, message: '请输入通知标题', trigger: 'blur' }],
  content: [{ required: true, message: '请输入通知内容', trigger: 'blur' }]
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

const getTableData = async () => {
  const table = await getNoticeList({
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
  form.ID = 0
  dialogFormVisible.value = true
}

const closeDialog = () => {
  dialogFormVisible.value = false
  resetForm()
}

const resetForm = () => {
  Object.assign(form, {
    title: '',
    content: '',
    status: 1,
    isVisible: 1,
    sort: 0,
    startTime: '',
    ID: 0,
    endTime: ''
  })
}

const enterDialog = async () => {
  elFormRef.value?.validate(async (valid) => {
    if (!valid) return
    let res
    if (form.ID) {
      res = await updateNoticeApi(form)
    } else {
      res = await createNotice(form)
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

const editScrollNotice = (row) => {
  Object.assign(form, row)
  dialogFormVisible.value = true
}

const deleteScrollNotice = async (row) => {
  ElMessageBox.confirm('此操作将永久删除该滚动通知，是否继续？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    const res = await deleteNoticeApi({ id: row.ID })
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '删除成功'
      })
      getTableData()
    }
  })
}

const onDelete = async () => {
  ElMessageBox.confirm('此操作将永久删除所选滚动通知，是否继续？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    const ids = multipleSelection.value.map(item => item.ID)
    for (const id of ids) {
      await deleteNoticeApi({ id })
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

<style scoped></style>
