<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="显示标题">
          <el-input v-model="searchInfo.title" placeholder="请输入显示标题" />
        </el-form-item>
        <el-form-item label="类别名称">
          <el-input v-model="searchInfo.categoryName" placeholder="请输入类别名称" />
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
        <el-button type="primary" icon="plus" @click="openDialog">新增活动配置</el-button>
        <el-button type="danger" icon="delete" @click="onDelete" :disabled="!multipleSelection.length">批量删除</el-button>
      </div>
      <el-table ref="multipleTable" style="width: 100%" tooltip-effect="dark" :data="tableData" row-key="ID"
        @selection-change="handleSelectionChange">
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="显示标题" prop="title" min-width="200" />
        <el-table-column align="left" label="活动图片" prop="image" min-width="120" >
          <template #default="scope">
            <el-image v-if="scope.row.image" :src="scope.row.image" :preview-src-list="[scope.row.image]" 
              style="width: 60px; height: 60px; cursor: pointer;" fit="cover" />
            <span v-else>-</span>
          </template>
        </el-table-column>
        <el-table-column align="left" label="类别名称" prop="categoryName" min-width="150" />
        <el-table-column align="left" label="状态" min-width="80" >
          <template #default="scope">
            <el-tag v-if="scope.row.status === 1" type="success">启用</el-tag>
            <el-tag v-else type="danger">禁用</el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="是否显示" min-width="100" >
          <template #default="scope">
            <el-tag v-if="scope.row.isVisible === 1" type="success">显示</el-tag>
            <el-tag v-else type="info">隐藏</el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="排序" prop="sort" min-width="80" />
        <el-table-column align="left" label="创建时间" prop="CreatedAt" min-width="180" >
          <template #default="scope">
            {{ formatDate(scope.row.created_at) }}
          </template>
        </el-table-column>
        <el-table-column align="left" label="操作" :min-width="appStore.operateMinWith" fixed="right">
          <template #default="scope">
            <el-button type="primary" link icon="edit" @click="editActivityConfig(scope.row)">编辑</el-button>
            <el-button type="warning" link icon="edit" @click="editContentFunc(scope.row)">编辑内容</el-button>
            <el-button type="primary" link icon="delete" @click="deleteActivityConfig(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination layout="total, sizes, prev, pager, next, jumper" :current-page="page" :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]" :total="total" @current-change="handleCurrentChange"
          @size-change="handleSizeChange" />
      </div>

      <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="活动配置" destroy-on-close width="1000px">
        <div class="gva-table-box">
          <div class="gva-form-box">
            <el-form ref="elFormRef" :model="form" :rules="rules" label-width="120px" label-position="left">
              <el-form-item label="显示标题" prop="title">
                <el-input v-model="form.title" placeholder="请输入显示标题" />
              </el-form-item>
              <el-form-item label="活动图片" prop="image">
                <el-input v-model="form.image" placeholder="请输入活动图片URL" />
                <div class="mt-2" v-if="form.image">
                  <el-image :src="form.image" style="max-width: 200px; max-height: 200px;" fit="contain" />
                </div>
              </el-form-item>
              <el-form-item label="跳转地址" prop="jumpUrl">
                <el-input v-model="form.jumpUrl" placeholder="请输入跳转地址" />
              </el-form-item>
              <el-form-item label="类别名称" prop="categoryName">
                <el-input v-model="form.categoryName" placeholder="请输入类别名称" />
              </el-form-item>
              <el-form-item label="类别图标" prop="categoryIcon">
                <el-input v-model="form.categoryIcon" placeholder="请输入类别图标URL" />
                <div class="mt-2" v-if="form.categoryIcon">
                  <el-image :src="form.categoryIcon" style="max-width: 100px; max-height: 100px;" fit="contain" />
                </div>
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
          </div>
        </div>
        <template #footer>
          <div class="dialog-footer">
            <el-button @click="closeDialog">取 消</el-button>
            <el-button type="primary" @click="enterDialog">确 定</el-button>
          </div>
        </template>
      </el-dialog>

      <!-- 内容查看弹窗 -->
      <el-dialog v-model="contentDialogVisible" title="活动内容" width="80%" top="5vh" destroy-on-close>
        <div class="content-container">
          <div v-if="currentContent" v-html="currentContent"></div>
          <div v-else class="no-content">暂无内容</div>
        </div>
      </el-dialog>

      <!-- 内容编辑弹窗 -->
      <el-dialog 
        v-model="contentEditDialogVisible" 
        title="编辑活动内容" 
        width="95%" 
        top="3vh"
        :close-on-click-modal="false"
        class="content-edit-dialog"
        destroy-on-close
      >
        <div class="content-edit-container">
          <div class="content-info">
            <h3>{{ currentActivityTitle }}</h3>
          </div>
          <div class="editor-layout">
            <div class="editor-left">
              <div class="editor-header">
                <span>源码编辑</span>
              </div>
              <el-input
                v-model="contentText"
                type="textarea"
                :rows="20"
                placeholder="请输入HTML源码..."
                class="source-editor"
              />
            </div>
            <div class="editor-right">
              <div class="editor-header">
                <span>预览</span>
              </div>
              <div class="preview-container">
                <div v-if="contentText" v-html="contentText" class="preview-content"></div>
                <div v-else class="preview-placeholder">暂无内容预览</div>
              </div>
            </div>
          </div>
          <div class="content-actions">
            <el-button type="primary" @click="saveContent">保存内容</el-button>
            <el-button @click="cancelContentEdit">取消</el-button>
          </div>
        </div>
      </el-dialog>
    </div>
  </div>
</template>

<script setup>
import {
  getActivityConfigList,
  createActivityConfig,
  updateActivityConfig as updateActivityConfigApi,
  deleteActivityConfig as deleteActivityConfigApi
} from '@/api/activityConfig'
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { formatDate } from '@/utils/format'
import { useAppStore } from '@/pinia/modules/app'

defineOptions({
  name: 'ActivityConfig'
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
  id:0,
  title: '',
  image: '',
  jumpUrl: '',
  categoryName: '',
  categoryIcon: '',
  status: 1,
  isVisible: 1,
  sort: 0
})

// 内容相关状态
const contentDialogVisible = ref(false)
const contentEditDialogVisible = ref(false)
const currentContent = ref('')
const contentText = ref('')
const currentActivityId = ref(0)
const currentActivityTitle = ref('')

const rules = reactive({
  title: [{ required: true, message: '请输入显示标题', trigger: 'blur' }]
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
  const table = await getActivityConfigList({
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
    title: '',
    image: '',
    jumpUrl: '',
    categoryName: '',
    categoryIcon: '',
    status: 1,
    isVisible: 1,
    sort: 0,
    id:0,
  })
}

// 查看内容
const viewContentFunc = (row) => {
  currentContent.value = row.content || ''
  contentDialogVisible.value = true
}

// 编辑内容
const editContentFunc = (row) => {
  currentActivityId.value = row.id
  currentActivityTitle.value = row.title
  contentText.value = row.content || ''
  contentEditDialogVisible.value = true
}

// 保存内容
const saveContent = async () => {
  try {
    const res = await updateActivityConfigApi({
      id: currentActivityId.value,
      content: contentText.value
    })
    if (res.code === 0) {
      ElMessage.success('内容保存成功')
      contentEditDialogVisible.value = false
      // 更新本地数据
      const activity = tableData.value.find(item => item.id === currentActivityId.value)
      if (activity) {
        activity.content = contentText.value
      }
    } 
  } catch (error) {
    ElMessage.error('保存失败')
  }
}

// 取消内容编辑
const cancelContentEdit = () => {
  contentEditDialogVisible.value = false
  contentText.value = ''
  currentActivityTitle.value = ''
  currentActivityId.value = 0
}

const enterDialog = async () => {
  elFormRef.value?.validate(async (valid) => {
    if (!valid) return
    
    let res
    if (form.id) {
      res = await updateActivityConfigApi(form)
    } else {
      res = await createActivityConfig(form)
    }
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: form.id ? '更新成功' : '创建成功'
      })
      closeDialog()
      getTableData()
    }
  })
}

const editActivityConfig = (row) => {
  Object.assign(form, row)
  dialogFormVisible.value = true
}

const deleteActivityConfig = async (row) => {
  ElMessageBox.confirm('此操作将永久删除该活动配置，是否继续？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    const res = await deleteActivityConfigApi({ id: row.id })
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
  ElMessageBox.confirm('此操作将永久删除所选活动配置，是否继续？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    const ids = multipleSelection.value.map(item => item.id)
    for (const id of ids) {
      await deleteActivityConfigApi({ id })
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
.content-container {
  min-height: 200px;
  max-height: 600px;
  overflow-y: auto;
  padding: 20px;
  background-color: #fff;
  border: 1px solid #e4e7ed;
  border-radius: 4px;
}

.no-content {
  text-align: center;
  color: #909399;
  padding: 40px;
}

.content-edit-dialog {
  .content-edit-container {
    .content-info {
      margin-bottom: 20px;
      padding-bottom: 15px;
      border-bottom: 1px solid #e4e7ed;
      
      h3 {
        margin: 0;
        color: #303133;
        font-size: 18px;
        font-weight: 600;
      }
    }
    
    .editor-layout {
      display: flex;
      gap: 20px;
      height: 500px;
      
      .editor-left,
      .editor-right {
        flex: 1;
        display: flex;
        flex-direction: column;
        
        .editor-header {
          padding: 10px 15px;
          background-color: #f5f7fa;
          border: 1px solid #e4e7ed;
          border-bottom: none;
          font-weight: 600;
          color: #606266;
          font-size: 14px;
        }
      }
      
      .editor-left {
        .source-editor {
          flex: 1;
          
          :deep(.el-textarea__inner) {
            height: 100% !important;
            resize: none;
            font-family: 'Monaco', 'Menlo', 'Ubuntu Mono', monospace;
            font-size: 13px;
            line-height: 1.5;
            border-top-left-radius: 0;
            border-top-right-radius: 0;
          }
        }
      }
      
      .editor-right {
        .preview-container {
          flex: 1;
          border: 1px solid #e4e7ed;
          border-top: none;
          overflow-y: auto;
          background-color: #fff;
          
          .preview-content {
            padding: 20px;
            min-height: 100%;
          }
          
          .preview-placeholder {
            display: flex;
            align-items: center;
            justify-content: center;
            height: 100%;
            color: #909399;
            font-size: 14px;
          }
        }
      }
    }
    
    .content-actions {
      margin-top: 20px;
      text-align: right;
      padding-top: 15px;
      border-top: 1px solid #e4e7ed;
    }
  }
}
</style>

