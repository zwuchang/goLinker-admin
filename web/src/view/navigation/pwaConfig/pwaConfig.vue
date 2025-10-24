<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="应用名称">
          <el-input v-model="searchInfo.name" placeholder="请输入应用名称" />
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
        <el-button type="primary" icon="plus" @click="openDialog" :disabled="hasConfig">新增PWA配置</el-button>
        <el-button type="warning" icon="refresh" @click="clearCache">清除缓存</el-button>
      </div>
      <el-table ref="multipleTable" tooltip-effect="dark" :data="tableData" row-key="ID">
        <el-table-column align="left" label="应用名称" prop="name" />
        <el-table-column align="left" label="短名称" prop="shortName" />
        <el-table-column align="left" label="作用域" prop="scope" show-overflow-tooltip />
        <el-table-column align="left" label="启动URL" prop="startUrl" show-overflow-tooltip />
        <el-table-column align="left" label="状态" width="80">
          <template #default="scope">
            <el-tag v-if="scope.row.status === 1" type="success">启用</el-tag>
            <el-tag v-else type="danger">禁用</el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="创建时间" prop="CreatedAt">
          <template #default="scope">
            {{ formatDate(scope.row.CreatedAt) }}
          </template>
        </el-table-column>
        <el-table-column align="left" label="操作" :width="appStore.operateMinWith" fixed="right">
          <template #default="scope">
            <el-button type="primary" link icon="edit" @click="editPWAConfig(scope.row)">编辑</el-button>
            <el-button type="primary" link icon="delete" @click="deletePWAConfig(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination layout="total, sizes, prev, pager, next, jumper" :current-page="page" :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]" :total="total" @current-change="handleCurrentChange"
          @size-change="handleSizeChange" />
      </div>


      <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="PWA配置" destroy-on-close width="800px">
        <div class="gva-table-box">
          <div class="gva-form-box">
            <el-form ref="elFormRef" :model="form" :rules="rules" label-width="100px" label-position="right">
              <el-row :gutter="20">
                <el-col :span="12">
                  <el-form-item label="状态" prop="status">
                    <el-radio-group v-model="form.status">
                      <el-radio :label="1">启用</el-radio>
                      <el-radio :label="0">禁用</el-radio>
                    </el-radio-group>
                    <div class="field-tip">
                      <el-text type="info" size="small">
                        <el-icon>
                          <InfoFilled />
                        </el-icon>
                        控制PWA配置是否生效，禁用后配置将不会应用到前端
                      </el-text>
                    </div>
                  </el-form-item>
                </el-col>
                <el-col :span="12">
                  <el-form-item label="应用名称" prop="name">
                    <el-input v-model="form.name" placeholder="请输入应用名称" />
                    <div class="field-tip">
                      <el-text type="info" size="small">
                        <el-icon>
                          <InfoFilled />
                        </el-icon>
                        PWA应用的完整名称，会显示在安装提示和应用列表中
                      </el-text>
                    </div>
                  </el-form-item>
                </el-col>
              </el-row>
              <el-row :gutter="20">
                <el-col :span="12">
                  <el-form-item label="短名称" prop="shortName">
                    <el-input v-model="form.shortName" placeholder="请输入短名称" />
                    <div class="field-tip">
                      <el-text type="info" size="small">
                        <el-icon>
                          <InfoFilled />
                        </el-icon>
                        PWA应用的简短名称，用于桌面图标下方和空间受限的地方
                      </el-text>
                    </div>
                  </el-form-item>
                </el-col>
                <el-col :span="12">
                  <el-form-item label="背景色" prop="backgroundColor">
                    <el-color-picker v-model="form.backgroundColor" />
                    <div class="field-tip">
                      <el-text type="info" size="small">
                        <el-icon>
                          <InfoFilled />
                        </el-icon>
                        应用启动时的背景颜色，在图标加载前显示
                      </el-text>
                    </div>
                  </el-form-item>
                </el-col>
              </el-row>
              <el-row :gutter="20">
                <el-col :span="12">
                  <el-form-item label="主题色" prop="themeColor">
                    <el-color-picker v-model="form.themeColor" />
                    <div class="field-tip">
                      <el-text type="info" size="small">
                        <el-icon>
                          <InfoFilled />
                        </el-icon>
                        应用的主题颜色，影响状态栏和浏览器UI的颜色
                      </el-text>
                    </div>
                  </el-form-item>
                </el-col>
                <el-col :span="12">
                  <el-form-item label="显示模式" prop="display">
                    <el-select v-model="form.display" placeholder="请选择显示模式">
                      <el-option label="standalone" value="standalone" />
                      <el-option label="minimal-ui" value="minimal-ui" />
                      <el-option label="fullscreen" value="fullscreen" />
                      <el-option label="browser" value="browser" />
                    </el-select>
                    <div class="field-tip">
                      <el-text type="info" size="small">
                        <el-icon>
                          <InfoFilled />
                        </el-icon>
                        应用启动时的显示方式，standalone为独立应用模式
                      </el-text>
                    </div>
                  </el-form-item>
                </el-col>
              </el-row>
              <el-row :gutter="20">
                <el-col :span="12">
                  <el-form-item label="屏幕方向" prop="orientation">
                    <el-select v-model="form.orientation" placeholder="请选择屏幕方向">
                      <el-option label="portrait" value="portrait" />
                      <el-option label="landscape" value="landscape" />
                      <el-option label="any" value="any" />
                    </el-select>
                    <div class="field-tip">
                      <el-text type="info" size="small">
                        <el-icon>
                          <InfoFilled />
                        </el-icon>
                        应用的首选屏幕方向，portrait为竖屏，landscape为横屏
                      </el-text>
                    </div>
                  </el-form-item>
                </el-col>
                <el-col :span="12">
                  <el-form-item label="作用域" prop="scope">
                    <el-input v-model="form.scope" placeholder="请输入作用域" />
                    <div class="field-tip">
                      <el-text type="info" size="small">
                        <el-icon>
                          <InfoFilled />
                        </el-icon>
                        定义PWA应用控制的URL范围，如：https://yourdomain.com/
                      </el-text>
                    </div>
                  </el-form-item>
                </el-col>
              </el-row>
              <el-form-item label="启动URL" prop="startUrl">
                <el-input v-model="form.startUrl" placeholder="请输入启动URL" />
                <div class="field-tip">
                  <el-text type="info" size="small">
                    <el-icon>
                      <InfoFilled />
                    </el-icon>
                    填写PWA应用启动时的页面地址，如：https://yourdomain.com/
                  </el-text>
                </div>
              </el-form-item>

              <!-- 图标配置 -->
              <el-form-item label="图标配置">
                <div class="field-tip">
                  <el-text type="info" size="small">
                    <el-icon>
                      <InfoFilled />
                    </el-icon>
                    配置PWA应用的各种尺寸图标，支持PNG、JPG等格式，建议提供多种尺寸
                  </el-text>
                </div>
                <div class="icon-config-container">
                  <div v-for="(icon, index) in form.icons" :key="index" class="icon-item">
                    <el-row :gutter="20">
                      <el-col :span="6">
                        <el-input v-model="icon.sizes" placeholder="尺寸" />
                      </el-col>
                      <el-col :span="12">
                        <el-input v-model="icon.src" placeholder="图标地址" />
                      </el-col>
                      <el-col :span="3">
                        <div class="icon-preview-container">
                          <img v-if="icon.src" :src="icon.src" class="icon-preview" @error="handleIconError" />
                          <span v-else class="no-image">无图片</span>
                        </div>
                      </el-col>
                      <el-col :span="2">
                        <el-button type="danger" icon="delete" size="small" @click="removeIcon(index)" />
                      </el-col>
                    </el-row>
                  </div>
                  <el-button type="primary" icon="plus" size="small" @click="addIcon">添加图标</el-button>
                </div>
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
  getPWAConfigList,
  createPWAConfig,
  updatePWAConfig as updatePWAConfigApi,
  deletePWAConfig as deletePWAConfigApi,
  clearPWACache
} from '@/api/pwaConfig'
import { ref, reactive, onMounted, computed } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { InfoFilled } from '@element-plus/icons-vue'
import { formatDate } from '@/utils/format'
import { useAppStore } from '@/pinia/modules/app'

defineOptions({
  name: 'PWAConfig'
})

const appStore = useAppStore()

const page = ref(1)
const pageSize = ref(10)
const total = ref(0)
const tableData = ref([])
const searchInfo = ref({})

const dialogFormVisible = ref(false)
const form = reactive({
  backgroundColor: '#164533',
  themeColor: '#164533',
  display: 'standalone',
  name: '',
  shortName: '',
  orientation: 'portrait',
  scope: '',
  startUrl: '',
  status: 1,
  icons: [
    {
      sizes: '192x192',
      src: '',
      type: 'image/png',
      sort: 0
    }
  ]
})

// 计算属性：是否已有配置
const hasConfig = computed(() => {
  return tableData.value.length > 0
})

const rules = reactive({
  name: [{ required: true, message: '请输入应用名称', trigger: 'blur' }],
  shortName: [{ required: true, message: '请输入短名称', trigger: 'blur' }],
  scope: [{ required: true, message: '请输入作用域', trigger: 'blur' }],
  startUrl: [{ required: true, message: '请输入启动URL', trigger: 'blur' }]
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
  const table = await getPWAConfigList({
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
    backgroundColor: '#164533',
    themeColor: '#164533',
    display: 'standalone',
    name: '',
    shortName: '',
    orientation: 'portrait',
    scope: '',
    startUrl: '',
    status: 1,
    icons: [
      {
        sizes: '192x192',
        src: '',
        type: 'image/png',
        sort: 0
      }
    ]
  })
}

const addIcon = () => {
  form.icons.push({
    sizes: '512x512',
    src: '',
    type: 'image/png',
    sort: form.icons.length
  })
}

const removeIcon = (index) => {
  if (form.icons.length > 1) {
    form.icons.splice(index, 1)
  } else {
    ElMessage({
      type: 'warning',
      message: '至少保留一个图标'
    })
    return
  }
}

const handleIconError = (e) => {
  e.target.style.display = 'none'
}

const enterDialog = async () => {
  elFormRef.value?.validate(async (valid) => {
    if (!valid) return
    let res
    if (form.ID) {
      res = await updatePWAConfigApi(form)
    } else {
      res = await createPWAConfig(form)
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

const editPWAConfig = (row) => {
  Object.assign(form, row)
  // 解析图标JSON字符串
  if (form.icons && typeof form.icons === 'string') {
    try {
      form.icons = JSON.parse(form.icons)
    } catch (e) {
      form.icons = [
        {
          sizes: '192x192',
          src: '',
          type: 'image/png',
          sort: 0
        }
      ]
    }
  }
  // 确保icons数组存在
  if (!form.icons || form.icons.length === 0) {
    form.icons = [
      {
        sizes: '192x192',
        src: '',
        type: 'image/png',
        sort: 0
      }
    ]
  }
  dialogFormVisible.value = true
}

const deletePWAConfig = async (row) => {
  ElMessageBox.confirm('此操作将永久删除该PWA配置，是否继续？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    const res = await deletePWAConfigApi({ id: row.ID })
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '删除成功'
      })
      getTableData()
    }
  })
}

const clearCache = async () => {
  ElMessageBox.confirm('此操作将清除PWA配置缓存，是否继续？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    const res = await clearPWACache()
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '缓存清除成功'
      })
    }
  })
}

onMounted(() => {
  getTableData()
})
</script>

<style scoped>


.icon-config-container {
  border: 1px solid #dcdfe6;
  border-radius: 4px;
  padding: 10px;
  background-color: #fafafa;
  width: 100%;
}

.icon-item {
  margin-bottom: 10px;
  padding: 8px;
  background: white;
  border-radius: 4px;
  border: 1px solid #e4e7ed;
}

.icon-item:last-child {
  margin-bottom: 0;
}

.icon-preview-container {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 32px;
  width: 100%;
  border: 1px solid #dcdfe6;
  border-radius: 4px;
  background-color: #fafafa;
}

.icon-preview {
  width: 24px;
  height: 24px;
  object-fit: cover;
  border-radius: 2px;
}

.no-image {
  color: #909399;
  font-size: 12px;
}

.field-tip {
  margin-top: 4px;
  line-height: 1.4;
}

.field-tip .el-text {
  display: flex;
  align-items: center;
  gap: 4px;
}
</style>
