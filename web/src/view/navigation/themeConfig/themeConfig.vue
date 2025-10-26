<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="主题名称">
          <el-input v-model="searchInfo.name" placeholder="请输入主题名称" />
        </el-form-item>
        <el-form-item label="主题描述">
          <el-input v-model="searchInfo.description" placeholder="请输入主题描述" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>

    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDialog">新增主题配置</el-button>
        <el-button type="danger" icon="delete" @click="onDelete" :disabled="!multipleSelection.length">批量删除</el-button>
      </div>
      <el-table ref="multipleTable" style="width: 100%" tooltip-effect="dark" :data="tableData" row-key="ID"
        @selection-change="handleSelectionChange">
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="主题名称" prop="name" />
        <el-table-column align="left" label="主题描述" prop="description"  show-overflow-tooltip />
        <el-table-column align="left" label="是否默认" width="100">
          <template #default="scope">
            <el-tag v-if="scope.row.is_default === 1" type="success">默认</el-tag>
            <el-tag v-else type="info">否</el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="创建时间" prop="created_at" >
          <template #default="scope">
            {{ formatDate(scope.row.created_at) }}
          </template>
        </el-table-column>
        <el-table-column align="left" label="操作" :width="appStore.operateMinWith" fixed="right">
          <template #default="scope">
            <el-button type="primary" link icon="edit" @click="editThemeConfig(scope.row)">编辑</el-button>
            <el-button type="success" link icon="check" @click="setDefault(scope.row)" v-if="scope.row.is_default !== 1">设为默认</el-button>
            <el-button type="primary" link icon="delete" @click="deleteThemeConfig(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination layout="total, sizes, prev, pager, next, jumper" :current-page="page" :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]" :total="total" @current-change="handleCurrentChange"
          @size-change="handleSizeChange" />
      </div>

      <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="主题配置" destroy-on-close width="1000px">
        <div class="gva-table-box">
          <div class="gva-form-box">
            <el-form ref="elFormRef" :model="form" :rules="rules" label-width="120px" label-position="left">
              <el-form-item label="主题名称" prop="name">
                <el-input v-model="form.name" placeholder="请输入主题名称" />
              </el-form-item>
              <el-form-item label="主题描述" prop="description">
                <el-input v-model="form.description" placeholder="请输入主题描述" />
              </el-form-item>
              <el-form-item label="是否默认" prop="is_default">
                <el-radio-group v-model="form.is_default">
                  <el-radio :label="1">是</el-radio>
                  <el-radio :label="0">否</el-radio>
                </el-radio-group>
              </el-form-item>
              
              <!-- 主题配置区域 -->
              <el-form-item label="主题配置">
                <div class="theme-config-layout">
                  <!-- 左侧配置区域 -->
                  <div class="config-section">
                    <div class="theme-config-section">
                      <div class="config-row">
                        <div class="config-item">
                          <label>大背景色</label>
                          <el-color-picker v-model="themeConfig['--bg-color']" color-format="hex" />
                          <el-input v-model="themeConfig['--background-color']" placeholder="#1a895f" style="width: 120px; margin-left: 10px;" @blur="validateColorInput('--background-color', themeConfig['--background-color'])" />
                        </div>
                      </div>
                      
                      <div class="config-row">
                        <div class="config-item">
                          <label>大背景图片</label>
                          <el-input v-model="themeConfig['--bg-image']" placeholder="请输入背景图片URL" style="width: 300px;" />
                        </div>
                      </div>
                      
                      <div class="config-row">
                        <div class="config-item">
                          <label>主要背景色</label>
                          <el-color-picker v-model="themeConfig['--main-background-color']" color-format="hex" />
                          <el-input v-model="themeConfig['--main-background-color']" placeholder="#164533" style="width: 120px; margin-left: 10px;" @blur="validateColorInput('--main-background-color', themeConfig['--main-background-color'])" />
                        </div>
                      </div>
                      
                      <div class="config-row">
                        <div class="config-item">
                          <label>次级背景色</label>
                          <el-color-picker v-model="themeConfig['--secondary-background-color']" color-format="hex" />
                          <el-input v-model="themeConfig['--secondary-background-color']" placeholder="#346452" style="width: 120px; margin-left: 10px;" @blur="validateColorInput('--content-background-color', themeConfig['--content-background-color'])" />
                        </div>
                      </div>

                      <div class="config-row">
                        <div class="config-item">
                          <label>底部菜单背景色</label>
                          <el-color-picker v-model="themeConfig['--footer-background-color']" color-format="hex" />
                          <el-input v-model="themeConfig['--footer-background-color']" placeholder="#0a3424" style="width: 120px; margin-left: 10px;" @blur="validateColorInput('--content-background-color', themeConfig['--content-background-color'])" />
                        </div>
                      </div>

                      <div class="config-row">
                        <div class="config-item">
                          <label>底部菜单选中颜色</label>
                          <el-color-picker v-model="themeConfig['--footer-selected-color']" color-format="hex" />
                          <el-input v-model="themeConfig['--footer-selected-color']" placeholder="#eec96f" style="width: 120px; margin-left: 10px;" @blur="validateColorInput('--content-background-color', themeConfig['--content-background-color'])" />
                        </div>
                      </div>

                      
                      <div class="config-row">
                        <div class="config-item">
                          <label>按钮背景色</label>
                          <el-color-picker v-model="themeConfig['--button-background-color']" color-format="hex" />
                          <el-input v-model="themeConfig['--button-background-color']" placeholder="#2cee88" style="width: 120px; margin-left: 10px;" @blur="validateColorInput('--button-background-color', themeConfig['--button-background-color'])" />
                        </div>
                      </div>

                      <!-- 按钮渐变色配置 - 一行显示 -->
                      <div class="config-row config-row-split">
                        <div class="config-item-half">
                          <label>按钮渐变色起始</label>
                          <el-color-picker v-model="themeConfig['--button-gradient-start']" color-format="hex" />
                          <el-input v-model="themeConfig['--button-gradient-start']" placeholder="#2cee88" style="width: 100px; margin-left: 5px;" @blur="validateColorInput('--button-gradient-start', themeConfig['--button-gradient-start'])" />
                        </div>
                        <div class="config-item-half">
                          <label>按钮渐变色结束</label>
                          <el-color-picker v-model="themeConfig['--button-gradient-end']" color-format="hex" />
                          <el-input v-model="themeConfig['--button-gradient-end']" placeholder="#9ae872" style="width: 100px; margin-left: 5px;" @blur="validateColorInput('--button-gradient-end', themeConfig['--button-gradient-end'])" />
                        </div>
                      </div>

                      <!-- 按钮预览 -->
                      <div class="config-row">
                        <div class="config-preview-button" :style="getButtonPreviewStyle()">
                          <span>按钮预览</span>
                        </div>
                      </div>

                      <div class="config-row">
                        <div class="config-item">
                          <label>按钮边框颜色</label>
                          <el-color-picker v-model="themeConfig['--button-border-color']" color-format="hex" />
                          <el-input v-model="themeConfig['--button-border-color']" placeholder="#2cee88" style="width: 120px; margin-left: 10px;" @blur="validateColorInput('--button-border-color', themeConfig['--button-border-color'])" />
                        </div>
                      </div>

                      <div class="config-row">
                        <div class="config-item">
                          <label>按钮文字颜色</label>
                          <el-color-picker v-model="themeConfig['--button-font-color']" color-format="hex" />
                          <el-input v-model="themeConfig['--button-font-color']" placeholder="#000" style="width: 120px; margin-left: 10px;" @blur="validateColorInput('--button-background-color', themeConfig['--button-background-color'])" />
                        </div>
                      </div>

                      <div class="config-row">
                        <div class="config-item">
                          <label>tab 选中颜色</label>
                          <el-color-picker v-model="themeConfig['--tab-selected-color']" color-format="hex" />
                          <el-input v-model="themeConfig['--tab-selected-color']" placeholder="#000" style="width: 120px; margin-left: 10px;" @blur="validateColorInput('--content-background-color', themeConfig['--content-background-color'])" />
                        </div>
                      </div>
                      
                      <div class="config-row">
                        <div class="config-item">
                          <label>字体颜色</label>
                          <el-color-picker v-model="themeConfig['--text-color']" color-format="hex" />
                          <el-input v-model="themeConfig['--text-color']" placeholder="#fff" style="width: 120px; margin-left: 10px;" @blur="validateColorInput('--text-color', themeConfig['--text-color'])" />
                        </div>
                      </div>

                      <div class="config-row">
                        <div class="config-item">
                          <label>PWA下载背景色</label>
                          <el-color-picker v-model="themeConfig['--pwa-download-bg-color']" color-format="hex" />
                          <el-input v-model="themeConfig['--pwa-download-bg-color']" placeholder="#0a3424" style="width: 120px; margin-left: 10px;" @blur="validateColorInput('--pwa-download-bg-color', themeConfig['--pwa-download-bg-color'])" />
                        </div>
                      </div>

                      <div class="config-row">
                        <div class="config-item">
                          <label>PWA下载背景图</label>
                          <el-input v-model="themeConfig['--pwa-download-bg-image']" placeholder="请输入PWA下载背景图片URL" style="width: 300px;" />
                        </div>
                      </div>
                    </div>
                  </div>
                  
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
  getThemeConfigList,
  createThemeConfig,
  updateThemeConfig as updateThemeConfigApi,
  deleteThemeConfig as deleteThemeConfigApi,
  setDefaultTheme
} from '@/api/themeConfig'
import { ref, reactive, onMounted, watch } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { formatDate } from '@/utils/format'
import { useAppStore } from '@/pinia/modules/app'

defineOptions({
  name: 'ThemeConfig'
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
  name: '',
  description: '',
  config_json: '',
  is_default: 0,
  id: 0,
})

// 主题配置对象
const themeConfig = reactive({
  '--bg-color': '#1a895f',
  '--bg-image': 'https://dydybet9.com/static/skins/br/images/skins/bg_home_green.png',
  '--main-background-color': '#164533',
  '--secondary-background-color': '#346452',
  '--button-background-color': '#2cee88',
  '--button-gradient-start': '#2cee88',
  '--button-gradient-end': '#9ae872',
  '--button-border-color': '#2cee88',
  '--footer-background-color': '#0a3424',
  '--footer-selected-color': '#eec96f',
  '--button-font-color': '#000',
  '--text-color': '#fff',
  '--tab-selected-color': '#000',
  '--pwa-download-bg-color': '#003f32',
  '--pwa-download-bg-image': 'https://dydybet9.com/static/skins/br/images/bg_download_app.png'
})

const rules = reactive({
  name: [{ required: true, message: '请输入主题名称', trigger: 'blur' }]
})

const elFormRef = ref()

// 验证16进制颜色格式
const validateHexColor = (color) => {
  const hexPattern = /^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$/;
  return hexPattern.test(color);
}

// 监听颜色输入，确保格式正确
const validateColorInput = (key, value) => {
  if (value && !validateHexColor(value)) {
    ElMessage.warning(`颜色格式不正确，请输入16进制格式（如：#1a895f）`);
    return false;
  }
  return true;
}

// 监听主题配置变化，同步到JSON
watch(themeConfig, (newVal) => {
  // 验证所有颜色值
  const colorKeys = ['--bg-color', '--main-background-color', '--secondary-background-color',
   '--footer-background-color', '--footer-selected-color', '--button-background-color', '--button-gradient-start', 
   '--button-gradient-end', '--button-border-color', '--button-font-color', 
   '--text-color', '--tab-selected-color', '--pwa-download-bg-color'];
  let isValid = true;
  
  colorKeys.forEach(key => {
    if (newVal[key] && !validateHexColor(newVal[key])) {
      isValid = false;
    }
  });
  
  if (isValid) {
    form.config_json = JSON.stringify(newVal, null, 2);
  }
}, { deep: true })

// 监听JSON变化，同步到主题配置
watch(() => form.config_json, (newVal) => {
  try {
    const config = JSON.parse(newVal)
    Object.assign(themeConfig, config)
  } catch (e) {
    // JSON格式错误时忽略
  }
})


// 获取按钮预览样式
const getButtonPreviewStyle = () => {
  return {
    background: `linear-gradient(150deg, ${themeConfig['--button-gradient-start']}, ${themeConfig['--button-gradient-end']})`,
    border: `1px solid ${themeConfig['--button-border-color']}`,
    color: themeConfig['--button-font-color'],
    borderRadius: '6px',
    display: 'inline-block',
    cursor: 'default'
  }
}

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
  const table = await getThemeConfigList({
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
    name: '',
    description: '',
    config_json: JSON.stringify(themeConfig, null, 2),
    is_default: 0,
    id: 0,
  })
}

const enterDialog = async () => {
  elFormRef.value?.validate(async (valid) => {
    if (!valid) return
    
    // 验证JSON格式
    try {
      JSON.parse(form.config_json)
    } catch (e) {
      ElMessage.error('主题配置JSON格式错误')
      return
    }
    
    let res
    if (form.id) {
      res = await updateThemeConfigApi(form)
    } else {
      res = await createThemeConfig(form)
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

const editThemeConfig = (row) => {
  Object.assign(form, row)
  try {
    const config = JSON.parse(row.config_json)
    Object.assign(themeConfig, config)
  } catch (e) {
    ElMessage.error('主题配置JSON格式错误')
  }
  dialogFormVisible.value = true
}

const deleteThemeConfig = async (row) => {
  ElMessageBox.confirm('此操作将永久删除该主题配置，是否继续？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    const res = await deleteThemeConfigApi({ id: row.id })
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '删除成功'
      })
      getTableData()
    }
  })
}

const setDefault = async (row) => {
  ElMessageBox.confirm('确定要将此主题设为默认主题吗？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    const res = await setDefaultTheme({ id: row.id })
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '设置成功'
      })
      getTableData()
    }
  })
}

const onDelete = async () => {
  ElMessageBox.confirm('此操作将永久删除所选主题配置，是否继续？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    const ids = multipleSelection.value.map(item => item.id)
    for (const id of ids) {
      await deleteThemeConfigApi({ id })
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

.theme-config-layout {
  display: flex;
  gap: 20px;
  align-items: flex-start;
  width: 100%;
}

.config-section {
  flex: 1;
}

.theme-config-section {
  border: 1px solid #dcdfe6;
  border-radius: 6px;
  padding: 20px;
  background-color: #fafafa;
}

.config-row {
  margin-bottom: 15px;
}

.config-item {
  display: flex;
  align-items: center;
  gap: 10px;
}

.config-item label {
  min-width: 120px;
  font-weight: 500;
  color: #606266;
}

.config-row-split {
  display: flex;
  gap: 20px;
}

.config-item-half {
  flex: 1;
  display: flex;
  align-items: center;
  gap: 5px;
}

.config-item-half label {
  min-width: 100px;
  font-weight: 500;
  color: #606266;
}

.config-preview-button {
  width: 100%;
  text-align: center;
  padding: 10px 0;
}
</style>
