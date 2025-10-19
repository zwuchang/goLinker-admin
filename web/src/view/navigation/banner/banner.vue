<template>
  <div class="app-container">
    <el-card class="box-card">
      <template #header>
        <div class="card-header">
          <span>Banner配置管理</span>
          <el-button type="primary" @click="openDialog">新增Banner</el-button>
        </div>
      </template>

      <!-- 搜索区域 -->
      <el-form :model="searchForm" :inline="true" class="search-form">
        <el-form-item label="标题">
          <el-input v-model="searchForm.title" placeholder="请输入标题" clearable />
        </el-form-item>
        <el-form-item label="媒体类型">
          <el-select v-model="searchForm.mediaType" placeholder="请选择媒体类型" clearable>
            <el-option label="图片" :value="1" />
            <el-option label="视频" :value="2" />
          </el-select>
        </el-form-item>
        <el-form-item label="显示状态">
          <el-select v-model="searchForm.isVisible" placeholder="请选择显示状态" clearable>
            <el-option label="显示" :value="1" />
            <el-option label="隐藏" :value="0" />
          </el-select>
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择状态" clearable>
            <el-option label="启用" :value="1" />
            <el-option label="禁用" :value="0" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="getList">搜索</el-button>
          <el-button @click="resetSearch">重置</el-button>
        </el-form-item>
      </el-form>

      <!-- 表格区域 -->
      <el-table v-loading="listLoading" :data="tableData" border>
        <el-table-column prop="ID" label="ID" width="80" />
        <el-table-column prop="title" label="标题" min-width="150" />
        <el-table-column prop="description" label="描述" min-width="200" show-overflow-tooltip />
        <el-table-column prop="mediaType" label="媒体类型" width="100">
          <template #default="{ row }">
            <el-tag :type="row.mediaType === 1 ? 'success' : 'primary'">
              {{ row.mediaType === 1 ? '图片' : '视频' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="mediaUrl" label="媒体预览" width="120">
          <template #default="{ row }">
            <div v-if="row.mediaType === 1" class="media-preview">
              <el-image
                :src="row.mediaUrl"
                :preview-src-list="[row.mediaUrl]"
                fit="cover"
                style="width: 60px; height: 40px; border-radius: 4px;"
              />
            </div>
            <div v-else-if="row.mediaType === 2" class="media-preview">
              <video
                :src="row.mediaUrl"
                style="width: 60px; height: 40px; border-radius: 4px; object-fit: cover;"
                controls
              />
            </div>
          </template>
        </el-table-column>
        <el-table-column prop="linkUrl" label="跳转链接" min-width="150" show-overflow-tooltip />
        <el-table-column prop="sort" label="排序" width="80" />
        <el-table-column prop="isVisible" label="显示状态" width="100">
          <template #default="{ row }">
            <el-tag :type="row.isVisible === 1 ? 'success' : 'info'">
              {{ row.isVisible === 1 ? '显示' : '隐藏' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="row.status === 1 ? 'success' : 'danger'">
              {{ row.status === 1 ? '启用' : '禁用' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="createdAt" label="创建时间" width="180">
          <template #default="{ row }">
            {{ formatDate(row.CreatedAt) }}
          </template>
        </el-table-column>
        <el-table-column label="操作" width="200" fixed="right">
          <template #default="{ row }">
            <el-button type="primary" size="small" @click="editBanner(row)">编辑</el-button>
            <el-button type="danger" size="small" @click="deleteBanner(row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>

      <!-- 分页 -->
      <el-pagination
        v-model:current-page="pageInfo.page"
        v-model:page-size="pageInfo.pageSize"
        :page-sizes="[10, 20, 50, 100]"
        :total="pageInfo.total"
        layout="total, sizes, prev, pager, next, jumper"
        @size-change="getList"
        @current-change="getList"
        class="pagination"
      />
    </el-card>

    <!-- 新增/编辑对话框 -->
    <el-dialog
      v-model="dialogVisible"
      :title="dialogTitle"
      width="600px"
      :close-on-click-modal="false"
      destroy-on-close
    >
      <el-form
        ref="formRef"
        :model="formData"
        :rules="formRules"
        label-width="100px"
      >
        <el-form-item label="标题" prop="title">
          <el-input v-model="formData.title" placeholder="请输入标题" />
        </el-form-item>
        <el-form-item label="描述" prop="description">
          <el-input
            v-model="formData.description"
            type="textarea"
            :rows="3"
            placeholder="请输入描述"
          />
        </el-form-item>
        <el-form-item label="媒体类型" prop="mediaType">
          <el-radio-group v-model="formData.mediaType">
            <el-radio :label="1">图片</el-radio>
            <el-radio :label="2">视频</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="媒体地址" prop="mediaUrl">
          <el-input v-model="formData.mediaUrl" placeholder="请输入媒体地址" />
          <div v-if="formData.mediaType === 1 && formData.mediaUrl" class="media-preview-form">
            <el-image
              :src="formData.mediaUrl"
              :preview-src-list="[formData.mediaUrl]"
              fit="cover"
              style="width: 200px; height: 120px; border-radius: 4px; margin-top: 10px;"
            />
          </div>
          <div v-else-if="formData.mediaType === 2 && formData.mediaUrl" class="media-preview-form">
            <video
              :src="formData.mediaUrl"
              style="width: 200px; height: 120px; border-radius: 4px; margin-top: 10px; object-fit: cover;"
              controls
            />
          </div>
        </el-form-item>
        <el-form-item label="跳转链接" prop="linkUrl">
          <el-input v-model="formData.linkUrl" placeholder="请输入跳转链接" />
        </el-form-item>
        <el-form-item label="排序" prop="sort">
          <el-input-number v-model="formData.sort" :min="0" :max="9999" />
        </el-form-item>
        <el-form-item label="显示状态" prop="isVisible">
          <el-radio-group v-model="formData.isVisible">
            <el-radio :label="1">显示</el-radio>
            <el-radio :label="0">隐藏</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="formData.status">
            <el-radio :label="1">启用</el-radio>
            <el-radio :label="0">禁用</el-radio>
          </el-radio-group>
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="submitForm" :loading="submitLoading">确定</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { getBannerList, createBanner, updateBanner, deleteBanner as deleteBannerApi } from '@/api/banner'
import { formatDate } from '@/utils/format'

// 响应式数据
const listLoading = ref(false)
const submitLoading = ref(false)
const dialogVisible = ref(false)
const dialogTitle = ref('')
const formRef = ref()

// 表格数据
const tableData = ref([])

// 搜索表单
const searchForm = reactive({
  title: '',
  mediaType: null,
  isVisible: null,
  status: null,
  orderBy: 'sort',
  orderType: 'asc'
})

// 分页信息
const pageInfo = reactive({
  page: 1,
  pageSize: 10,
  total: 0
})

// 表单数据
const formData = reactive({
  id: null,
  title: '',
  description: '',
  mediaType: 1,
  mediaUrl: '',
  linkUrl: '',
  sort: 0,
  isVisible: 1,
  status: 1
})

// 表单验证规则
const formRules = {
  title: [
    { required: true, message: '请输入标题', trigger: 'blur' }
  ],
  mediaType: [
    { required: true, message: '请选择媒体类型', trigger: 'change' }
  ],
  mediaUrl: [
    { required: true, message: '请输入媒体地址', trigger: 'blur' }
  ]
}

// 获取列表数据
const getList = async () => {
  listLoading.value = true
  try {
    const params = {
      ...searchForm,
      page: pageInfo.page,
      pageSize: pageInfo.pageSize
    }
    const res = await getBannerList(params)
    if (res.code === 0) {
      tableData.value = res.data.list
      pageInfo.total = res.data.total
    }
  } catch (error) {
    console.error('获取Banner列表失败:', error)
    ElMessage.error('获取列表失败')
  } finally {
    listLoading.value = false
  }
}

// 重置搜索
const resetSearch = () => {
  Object.assign(searchForm, {
    title: '',
    mediaType: null,
    isVisible: null,
    status: null,
    orderBy: 'sort',
    orderType: 'asc'
  })
  pageInfo.page = 1
  getList()
}

// 打开对话框
const openDialog = () => {
  dialogTitle.value = '新增Banner'
  resetForm()
  dialogVisible.value = true
}

// 编辑Banner
const editBanner = (row) => {
  dialogTitle.value = '编辑Banner'
  Object.assign(formData, {
    id: row.ID,
    title: row.title,
    description: row.description,
    mediaType: row.mediaType,
    mediaUrl: row.mediaUrl,
    linkUrl: row.linkUrl,
    sort: row.sort,
    isVisible: row.isVisible,
    status: row.status
  })
  dialogVisible.value = true
}

// 重置表单
const resetForm = () => {
  Object.assign(formData, {
    id: null,
    title: '',
    description: '',
    mediaType: 1,
    mediaUrl: '',
    linkUrl: '',
    sort: 0,
    isVisible: 1,
    status: 1
  })
  if (formRef.value) {
    formRef.value.resetFields()
  }
}

// 提交表单
const submitForm = async () => {
  if (!formRef.value) return
  
  try {
    await formRef.value.validate()
    submitLoading.value = true
    console.log(formData)

    const isEdit = !!formData.id
    console.log('isEdit', isEdit)
    const res = isEdit ? await updateBanner(formData) : await createBanner(formData)
    
    if (res.code === 0) {
      ElMessage.success(isEdit ? '更新成功' : '创建成功')
      dialogVisible.value = false
      getList()
    }
  } catch (error) {
    console.error('提交表单失败:', error)
    ElMessage.error('操作失败')
  } finally {
    submitLoading.value = false
  }
}

// 删除Banner
const deleteBannerConfirm = async (row) => {
  try {
    const res = await deleteBannerApi(row.ID)
    if (res.code === 0) {
      ElMessage.success('删除成功')
      getList()
    }
  } catch (error) {
    console.error('删除失败:', error)
    ElMessage.error('删除失败')
  }
}

// 删除确认
const deleteBanner = (row) => {
  ElMessageBox.confirm(
    `确定要删除Banner "${row.title}" 吗？`,
    '删除确认',
    {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }
  ).then(() => {
    deleteBannerConfirm(row)
  }).catch(() => {
    // 用户取消删除
  })
}

// 组件挂载时获取数据
onMounted(() => {
  getList()
})
</script>

<style scoped>
.app-container {
  padding: 20px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.search-form {
  margin-bottom: 20px;
}

.media-preview {
  display: flex;
  justify-content: center;
  align-items: center;
}

.media-preview-form {
  display: flex;
  justify-content: center;
  align-items: center;
}

.pagination {
  margin-top: 20px;
  display: flex;
  justify-content: center;
}

:deep(.el-table) {
  margin-top: 20px;
}

:deep(.el-form-item) {
  margin-bottom: 20px;
}
</style>
