<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="searchForm" :inline="true" :model="searchInfo">
        <el-form-item label="显示名称">
          <el-input v-model="searchInfo.displayName" placeholder="显示名称" />
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
        <el-button icon="delete" :disabled="!multipleSelection.length" @click="onDelete">
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
          label="显示名称"
          min-width="120"
          prop="displayName"
          sortable="custom"
        />
        <el-table-column
          align="left"
          label="跳转地址"
          min-width="200"
          prop="jumpUrl"
          show-overflow-tooltip
        />
        <el-table-column
          align="left"
          label="排序"
          min-width="80"
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
              @click="editContactMethod(scope.row)"
            >
              编辑
            </el-button>
            <el-button
              icon="delete"
              type="primary"
              link
              @click="removeContactMethod(scope.row)"
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
      <ContactMethodForm
        :id="formId"
        :edit-data="editData"
        @close="dialogFormVisible = false"
        @success="getTableData"
      />
    </el-dialog>
  </div>
</template>

<script setup>
import {
  createContactMethod,
  updateContactMethod,
  deleteContactMethod,
  getContactMethod,
  getContactMethodList
} from '@/api/contactMethod'
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { formatDate } from '@/utils/format'
import { useAppStore } from '@/pinia/modules/app'
import { toSQLLine } from '@/utils/stringFun'
import ContactMethodForm from './form.vue'

const appStore = useAppStore()

const searchForm = ref()
const tableData = ref([])
const page = ref(1)
const pageSize = ref(10)
const total = ref(0)
const searchInfo = ref({})
const multipleSelection = ref([])

const form = ref({
  displayName: '',
  jumpUrl: '',
  image: '',
  sort: 0,
  status: 1
})

const rules = ref({
  displayName: [
    { required: true, message: '请输入显示名称', trigger: 'blur' }
  ],
  jumpUrl: [
    { required: true, message: '请输入跳转地址', trigger: 'blur' }
  ]
})

const dialogFormVisible = ref(false)
const dialogTitle = ref('')
const formId = ref(0)
const editData = ref({})

// 搜索
const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  getTableData()
}

const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 排序
const sortChange = ({ prop, order }) => {
  if (prop) {
    searchInfo.value.orderKey = toSQLLine(prop)
    searchInfo.value.desc = order === 'descending'
  }
  getTableData()
}

// 选择
const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

// 批量删除
const onDelete = async () => {
  if (multipleSelection.value.length === 0) {
    ElMessage.warning('请选择要删除的数据')
    return
  }
  
  ElMessageBox.confirm('确定要删除选中的联系方式吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    try {
      const ids = multipleSelection.value.map(item => item.ID)
      for (const id of ids) {
        await deleteContactMethod({ ID: id })
      }
      ElMessage({
        type: 'success',
        message: '删除成功'
      })
      getTableData()
    } catch (error) {
      console.error('删除失败:', error)
      ElMessage.error('删除失败')
    }
  })
}

// 查询
const getTableData = async () => {
  try {
    const table = await getContactMethodList({
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
  } catch (error) {
    console.error('获取列表失败:', error)
    ElMessage.error('获取列表失败')
  }
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

const editContactMethod = async (row) => {
  formId.value = row.ID
  editData.value = { ...row }
  dialogTitle.value = '编辑联系方式'
  dialogFormVisible.value = true
}

const openDialog = () => {
  formId.value = 0
  editData.value = {}
  dialogTitle.value = '新增联系方式'
  dialogFormVisible.value = true
}

const removeContactMethod = async (row) => {
  ElMessageBox.confirm('确定要删除这个联系方式吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    try {
      const res = await deleteContactMethod({ ID: row.ID })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === 1 && page.value > 1) {
          page.value--
        }
        getTableData()
      }
    } catch (error) {
      console.error('删除失败:', error)
      ElMessage.error('删除失败')
    }
  })
}

onMounted(() => {
  getTableData()
})
</script>

<style scoped>
.gva-table-box {
  padding: 20px;
}

.gva-btn-list {
  margin-bottom: 20px;
}

.gva-pagination {
  margin-top: 20px;
  text-align: right;
}
</style>
