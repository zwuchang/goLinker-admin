<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="searchForm" :inline="true" :model="searchInfo">
        <el-form-item label="游戏标题">
          <el-input v-model="searchInfo.title" placeholder="游戏标题" />
        </el-form-item>
        <el-form-item label="游戏类别">
          <el-select
            v-model="searchInfo.category_id"
            clearable
            placeholder="请选择"
          >
            <el-option
              v-for="category in categoryList"
              :key="category.ID"
              :label="category.name"
              :value="category.ID"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="类型">
          <el-select v-model="searchInfo.type" clearable placeholder="请选择">
            <el-option label="图文" value="image_text" />
            <el-option label="视频" value="video" />
          </el-select>
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
        <el-form-item label="显示名称">
          <el-input v-model="searchInfo.display_name" placeholder="显示名称" />
        </el-form-item>
        <el-form-item label="广告名称">
          <el-input v-model="searchInfo.ad_name" placeholder="广告名称" />
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
        <el-button icon="delete" :disabled="!games.length" @click="onDelete">
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
        <el-table-column align="left" label="缩略图" prop="thumbnail" width="100">
          <template #default="scope">
            <el-image
              v-if="scope.row.thumbnail"
              :src="scope.row.thumbnail"
              style="width: 60px; height: 40px"
              fit="cover"
            />
            <span v-else>-</span>
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="游戏标题"
          min-width="150"
          prop="title"
          sortable="custom"
        />
        <el-table-column
          align="left"
          label="类型"
          min-width="100"
          prop="type"
          sortable="custom"
        >
          <template #default="scope">
            <el-tag :type="scope.row.type === 'image_text' ? 'primary' : 'success'">
              {{ scope.row.type === 'image_text' ? '图文' : '视频' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="类别"
          min-width="120"
          prop="category_ids"
        >
          <template #default="scope">
            <el-tag v-for="categoryId in getCategoryNames(scope.row.category_ids)" :key="categoryId" size="small">
              {{ categoryId }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="浏览次数"
          min-width="100"
          prop="views"
          sortable="custom"
        >
          <template #default="scope">
            <span 
              class="views-editable" 
              @dblclick="editViews(scope.row)"
              :title="'双击编辑浏览次数'"
            >
              {{ scope.row.views }}
            </span>
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="置顶"
          min-width="80"
          prop="sticky"
        >
          <template #default="scope">
            <el-tag :type="scope.row.sticky === 1 ? 'warning' : 'info'">
              {{ scope.row.sticky === 1 ? '是' : '否' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="状态"
          min-width="80"
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
          min-width="80"
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
          label="显示名称"
          min-width="120"
          prop="display_name"
          sortable="custom"
        />
        <el-table-column
          align="left"
          label="广告名称"
          min-width="120"
          prop="ad_name"
          sortable="custom"
        />
        <el-table-column
          align="left"
          label="跳转地址"
          min-width="200"
          prop="jump_url"
          show-overflow-tooltip
        />
        <el-table-column
          align="left"
          label="下载地址"
          min-width="200"
          prop="download_url"
          show-overflow-tooltip
        />
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
              @click="editGameFunc(scope.row)"
            >
              编辑
            </el-button>
            <el-button
              icon="delete"
              type="primary"
              link
              @click="deleteGameFunc(scope.row)"
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
    <el-dialog v-model="dialogFormVisible" :title="dialogTitle" width="800px">
      <GameForm
        :id="formId"
        :edit-data="editData"
        @close="dialogFormVisible = false"
        @success="getTableData"
      />
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { getGameList, deleteGame, updateGameViews } from '@/api/game'
import { getAllGameCategories } from '@/api/gameCategory'
import { ElMessage, ElMessageBox } from 'element-plus'
import GameForm from './form.vue'
import { useAppStore } from "@/pinia"
import { formatDate } from '@/utils/format'


defineOptions({
  name: 'Game'
})

const appStore = useAppStore()

const games = ref([])
const categoryList = ref([])
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

const dialogTitle = ref('新增游戏')
const dialogFormVisible = ref(false)
const formId = ref(0)
const editData = ref({})

// 获取类别列表
const getCategoryList = async () => {
  const res = await getAllGameCategories()
  if (res.code === 0) {
    categoryList.value = res.data
  }
}

// 查询
const getTableData = async () => {
  const params = {
    page: page.value,
    pageSize: pageSize.value,
    ...searchInfo.value
  }
  // 过滤掉空值
  Object.keys(params).forEach(key => {
    if (params[key] === null || params[key] === undefined || params[key] === '') {
      delete params[key]
    }
  })
  const res = await getGameList(params)
  if (res.code === 0) {
    tableData.value = res.data.list
    total.value = res.data.total
    page.value = res.data.page
    pageSize.value = res.data.pageSize
  }
}

const onReset = () => {
  searchInfo.value = {
    title: '',
    category_id: null,
    type: '',
    status: null,
    is_visible: null,
    display_name: '',
    ad_name: '',
    order_by: '',
    order_type: ''
  }
  getTableData()
}

// 搜索
const onSubmit = () => {
  page.value = 1
  getTableData()
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

// 排序
const sortChange = ({ prop, order }) => {
  if (prop) {
    // 字段名映射
    const fieldMap = {
      'ID': 'id',
      'title': 'title',
      'type': 'type',
      'status': 'status',
      'is_visible': 'is_visible',
      'views': 'views',
      'display_name': 'display_name',
      'ad_name': 'ad_name',
      'CreatedAt': 'created_at'
    }
    
    const orderBy = fieldMap[prop]
    if (orderBy) {
      searchInfo.value.order_by = orderBy
      searchInfo.value.order_type = order === 'descending' ? 'desc' : 'asc'
    } else {
      // 清除排序参数
      searchInfo.value.order_by = ''
      searchInfo.value.order_type = ''
    }
  } else {
    // 清除排序参数
    searchInfo.value.order_by = ''
    searchInfo.value.order_type = ''
  }
  getTableData()
}

// 批量操作
const handleSelectionChange = (val) => {
  games.value = val
}

const onDelete = async () => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    const ids = games.value.map((item) => item.ID)
    // 这里需要实现批量删除API
    for (const id of ids) {
      await deleteGame({ ID: id })
    }
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === ids.length && page.value > 1) {
      page.value--
    }
    getTableData()
  })
}

// 弹窗相关
const openDialog = () => {
  dialogTitle.value = '新增游戏'
  formId.value = 0
  editData.value = {}
  dialogFormVisible.value = true
}

const editGameFunc = (row) => {
  dialogTitle.value = '编辑游戏'
  formId.value = row.ID
  editData.value = { ...row }
  dialogFormVisible.value = true
}

const deleteGameFunc = async (row) => {
  ElMessageBox.confirm('此操作将永久删除该游戏, 是否继续?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    const res = await deleteGame({ ID: row.ID })
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '删除成功!'
      })
      if (tableData.value.length === 1 && page.value > 1) {
        page.value--
      }
      getTableData()
    }
  })
}

// 获取类别名称
const getCategoryNames = (categoryIds) => {
  if (!categoryIds) return []
  const ids = categoryIds.replace(/[\[\]]/g, '').split(',').filter(id => id.trim())
  return ids.map(id => {
    const category = categoryList.value.find(c => c.ID === parseInt(id.trim()))
    return category ? category.name : id.trim()
  })
}

// 编辑浏览次数
const editViews = async (row) => {
  try {
    const { value: newViews } = await ElMessageBox.prompt(
      '请输入新的浏览次数',
      '编辑浏览次数',
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        inputPattern: /^\d+$/,
        inputErrorMessage: '请输入有效的数字',
        inputValue: row.views.toString()
      }
    )
    
    const views = parseInt(newViews)
    if (views < 0) {
      ElMessage.error('浏览次数不能为负数')
      return
    }
    
    const res = await updateGameViews({ ID: row.ID, views: views })
    if (res.code === 0) {
      ElMessage.success('浏览次数更新成功')
      // 更新本地数据
      row.views = views
    } else {
      ElMessage.error(res.msg || '更新失败')
    }
  } catch {
    // 用户取消操作
  }
}

onMounted(() => {
  getCategoryList()
  getTableData()
})
</script>

<style scoped lang="scss">
.warning {
  color: #dc143c;
}

.views-editable {
  cursor: pointer;
  padding: 4px 8px;
  border-radius: 4px;
  transition: background-color 0.2s;
  
  &:hover {
    background-color: #f5f7fa;
  }
}
</style>
