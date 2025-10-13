<template>
  <div class="gva-table-box">
    <div class="gva-form-box">
      <el-form ref="gameFormRef" :model="form" label-width="80px">
        <el-form-item label="游戏标题" prop="title">
          <el-input v-model="form.title" placeholder="请输入游戏标题" />
        </el-form-item>
        <el-form-item label="缩略图" prop="thumbnail">
          <el-input v-model="form.thumbnail" placeholder="请输入缩略图URL" />
        </el-form-item>
        <el-form-item label="游戏类别" prop="category_ids">
          <el-select v-model="selectedCategories" multiple placeholder="请选择游戏类别" style="width: 100%">
            <el-option
              v-for="category in categoryList"
              :key="category.ID"
              :label="category.name"
              :value="category.ID"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="类型" prop="type">
          <el-radio-group v-model="form.type">
            <el-radio label="image_text">图文</el-radio>
            <el-radio label="video">视频</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item v-if="form.type === 'video'" label="视频地址" prop="video_url">
          <el-input v-model="form.video_url" placeholder="请输入视频地址" />
        </el-form-item>
        <el-form-item v-if="form.type === 'video'" label="视频时长" prop="video_duration">
          <el-input v-model="form.video_duration" placeholder="请输入视频时长" />
        </el-form-item>
        <el-form-item label="游戏描述" prop="description">
          <el-input v-model="form.description" type="textarea" placeholder="请输入游戏描述" />
        </el-form-item>
        <el-form-item label="游戏内容" prop="content">
          <el-input v-model="form.content" type="textarea" :rows="4" placeholder="请输入游戏内容" />
        </el-form-item>
        <el-form-item label="跳转地址" prop="jump_url">
          <el-input v-model="form.jump_url" placeholder="请输入跳转地址" />
        </el-form-item>
        <el-form-item label="下载地址" prop="download_url">
          <el-input v-model="form.download_url" placeholder="请输入下载地址" />
        </el-form-item>
        <el-form-item label="显示名称" prop="display_name">
          <el-input v-model="form.display_name" placeholder="请输入显示名称" />
        </el-form-item>
        <el-form-item label="广告名称" prop="ad_name">
          <el-input v-model="form.ad_name" placeholder="请输入广告名称" />
        </el-form-item>
        <el-form-item label="排序" prop="sort">
          <el-input-number v-model="form.sort" :min="0" />
        </el-form-item>
        <el-form-item label="浏览次数" prop="views">
          <el-input-number v-model="form.views" :min="0" placeholder="请输入浏览次数" />
        </el-form-item>
        <el-form-item label="置顶" prop="sticky">
          <el-radio-group v-model="form.sticky">
            <el-radio :label="1">是</el-radio>
            <el-radio :label="0">否</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="可见性" prop="is_visible">
          <el-radio-group v-model="form.is_visible">
            <el-radio :label="1">可见</el-radio>
            <el-radio :label="0">隐藏</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="form.status">
            <el-radio :label="1">启用</el-radio>
            <el-radio :label="0">禁用</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="onSubmit">确定</el-button>
          <el-button @click="onClose">取消</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, watch } from 'vue'
import { createGame, updateGame } from '@/api/game'
import { getAllGameCategories } from '@/api/gameCategory'
import { ElMessage } from 'element-plus'

const props = defineProps({
  id: {
    type: Number,
    default: 0
  },
  editData: {
    type: Object,
    default: () => ({})
  }
})

const emit = defineEmits(['close', 'success'])

const gameFormRef = ref()
const categoryList = ref([])

const form = reactive({
  title: '',
  thumbnail: '',
  category_ids: '',
  type: 'image_text',
  video_url: '',
  video_duration: '',
  description: '',
  content: '',
  jump_url: '',
  download_url: '',
  display_name: '',
  ad_name: '',
  sort: 0,
  sticky: 0,
  is_visible: 1,
  status: 1,
  views: 0
})

// 选中的类别
const selectedCategories = ref([])

// 监听选中类别变化，更新category_ids
watch(selectedCategories, (newVal) => {
  form.category_ids = `[${newVal.join(',')}]`
}, { deep: true })

// 获取类别列表
const getCategoryList = async () => {
  const res = await getAllGameCategories()
  if (res.code === 0) {
    categoryList.value = res.data
  }
}

// 监听编辑数据变化
watch(() => props.editData, (newData) => {
  if (newData && Object.keys(newData).length > 0) {
    Object.assign(form, newData)
    // 确保数值字段为整数类型
    form.views = parseInt(form.views) || 0
    form.sort = parseInt(form.sort) || 0
    // 解析category_ids
    if (form.category_ids) {
      const categoryIds = form.category_ids.replace(/[\[\]]/g, '').split(',').filter(id => id.trim())
      selectedCategories.value = categoryIds.map(id => parseInt(id.trim()))
    }
  }
}, { immediate: true, deep: true })

// 提交
const onSubmit = async () => {
  gameFormRef.value.validate(async (valid) => {
    if (valid) {
      // 确保数值字段为整数类型
      const submitData = {
        ...form,
        views: parseInt(form.views) || 0,
        sort: parseInt(form.sort) || 0
      }
      
      let res
      if (props.id) {
        res = await updateGame(submitData)
      } else {
        res = await createGame(submitData)
      }
      if (res.code === 0) {
        ElMessage.success(props.id ? '更新成功' : '创建成功')
        emit('success')
        onClose()
      }
    }
  })
}

// 关闭
const onClose = () => {
  emit('close')
}

// 重置表单
const resetForm = () => {
  Object.assign(form, {
    title: '',
    thumbnail: '',
    category_ids: '[]',
    type: 'image_text',
    video_url: '',
    video_duration: '',
    description: '',
    content: '',
    jump_url: '',
    download_url: '',
    display_name: '',
    ad_name: '',
    sort: 0,
    sticky: 0,
    is_visible: 1,
    status: 1,
    views: 0
  })
  selectedCategories.value = []
}

// 监听ID变化，重置表单
watch(() => props.id, (newId) => {
  if (!newId) {
    resetForm()
  }
})

// 初始化
getCategoryList()
</script>

<style scoped>
.gva-table-box {
  padding: 20px;
}

.gva-form-box {
  max-width: 800px;
  margin: 0 auto;
}
</style>
