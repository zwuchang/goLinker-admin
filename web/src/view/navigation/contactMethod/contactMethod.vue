<template>
  <div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDrawer">
          新增联系方式
        </el-button>
      </div>
      <el-table
        ref="multipleTable"
        :data="tableData"
        style="width: 100%"
        tooltip-effect="dark"
        row-key="ID"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="创建时间" width="180">
          <template #default="scope">
            <span>{{ formatDate(scope.row.CreatedAt) }}</span>
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="显示名称"
          prop="displayName"
          width="120"
        />
        <el-table-column
          align="left"
          label="跳转地址"
          prop="jumpUrl"
          width="200"
          show-overflow-tooltip
        />
        <el-table-column
          align="left"
          label="排序"
          prop="sort"
          width="80"
        />
        <el-table-column
          align="left"
          label="状态"
          prop="status"
          width="80"
        >
          <template #default="scope">
            <el-tag :type="scope.row.status === 1 ? 'success' : 'danger'">
              {{ scope.row.status === 1 ? '启用' : '禁用' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="操作" min-width="160">
          <template #default="scope">
            <el-button
              type="primary"
              link
              icon="edit"
              @click="editContactMethod(scope.row)"
            >
              编辑
            </el-button>
            <el-button
              type="primary"
              link
              icon="delete"
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
    
    <!-- 编辑抽屉 -->
    <el-drawer
      v-model="drawerFormVisible"
      :before-close="closeDrawer"
      :show-close="false"
      size="50%"
    >
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">{{ type === 'create' ? '新增' : '编辑' }}联系方式</span>
          <div>
            <el-button @click="closeDrawer">取 消</el-button>
            <el-button type="primary" @click="enterDrawer" :loading="loading">确 定</el-button>
          </div>
        </div>
      </template>
      
      <el-form :model="form" label-width="120px" :rules="rules" ref="formRef">
        <el-form-item label="显示名称" prop="displayName">
          <el-input v-model="form.displayName" placeholder="如：Twitter、Facebook、微信等" />
        </el-form-item>
        
        <el-form-item label="跳转地址" prop="jumpUrl">
          <el-input v-model="form.jumpUrl" placeholder="请输入跳转地址" />
        </el-form-item>
        
        <el-form-item label="联系方式图标">
          <select-image v-model="form.image" />
        </el-form-item>
        
        <el-form-item label="排序">
          <el-input-number v-model="form.sort" :min="0" :max="999" />
        </el-form-item>
        
        <el-form-item label="状态">
          <el-radio-group v-model="form.status">
            <el-radio :label="1">启用</el-radio>
            <el-radio :label="0">禁用</el-radio>
          </el-radio-group>
        </el-form-item>
      </el-form>
    </el-drawer>
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
import SelectImage from '@/components/selectImage/selectImage.vue'
import { ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { formatDate } from '@/utils/format'

defineOptions({
  name: 'ContactMethod'
})

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

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const drawerFormVisible = ref(false)
const type = ref('')
const formRef = ref()
const loading = ref(false)

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async () => {
  try {
    const table = await getContactMethodList({
      page: page.value,
      pageSize: pageSize.value
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

getTableData()

const editContactMethod = async (row) => {
  try {
    const res = await getContactMethod({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
      form.value = { ...res.data.contactMethod }
      drawerFormVisible.value = true
    }
  } catch (error) {
    console.error('获取详情失败:', error)
    ElMessage.error('获取详情失败')
  }
}

const closeDrawer = () => {
  drawerFormVisible.value = false
  form.value = {
    displayName: '',
    jumpUrl: '',
    image: '',
    sort: 0,
    status: 1
  }
  formRef.value?.resetFields()
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

const enterDrawer = async () => {
  const valid = await formRef.value?.validate()
  if (!valid) return
  
  loading.value = true
  try {
    let res
    switch (type.value) {
      case 'create':
        res = await createContactMethod(form.value)
        break
      case 'update':
        res = await updateContactMethod(form.value)
        break
      default:
        res = await createContactMethod(form.value)
        break
    }

    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: type.value === 'create' ? '创建成功' : '更新成功'
      })
      closeDrawer()
      getTableData()
    }
  } catch (error) {
    console.error('操作失败:', error)
    ElMessage.error('操作失败')
  } finally {
    loading.value = false
  }
}

const openDrawer = () => {
  type.value = 'create'
  drawerFormVisible.value = true
}
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
