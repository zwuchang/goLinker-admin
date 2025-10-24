<template>
  <div class="gva-table-box">
    <div class="gva-form-box">
      <el-form ref="gameConfigFormRef" :model="form" label-width="120px">
        <el-form-item label="下载地址" prop="download_url">
          <el-input v-model="form.download_url" placeholder="请输入下载地址" />
        </el-form-item>
        <el-form-item label="播放音频地址" prop="audio_url">
          <el-input v-model="form.audio_url" placeholder="请输入播放音频地址" />
        </el-form-item>
        <el-form-item label="网站标题" prop="website_title">
          <el-input v-model="form.website_title" placeholder="请输入网站标题" />
        </el-form-item>
        <el-form-item label="网站描述" prop="website_desc">
          <el-input v-model="form.website_desc" type="textarea" placeholder="请输入网站描述" />
        </el-form-item>
        <el-form-item label="网站图标" prop="website_icon">
          <div class="icon-input-container">
            <el-input v-model="form.website_icon" placeholder="请输入网站图标URL" />
            <img v-if="form.website_icon" :src="form.website_icon" class="icon-preview" @error="handleImageError" />
          </div>
        </el-form-item>
        <el-form-item label="网站Logo" prop="website_logo">
          <div class="icon-input-container">
            <el-input v-model="form.website_logo" placeholder="请输入网站Logo URL" />
            <img v-if="form.website_logo" :src="form.website_logo" class="icon-preview" @error="handleImageError" />
          </div>
        </el-form-item>
        <el-form-item label="市场Logo" prop="market_logo">
          <div class="icon-input-container">
            <el-input v-model="form.market_logo" placeholder="请输入市场Logo URL" />
            <img v-if="form.market_logo" :src="form.market_logo" class="icon-preview" @error="handleImageError" />
          </div>
        </el-form-item>
        <el-form-item label="悬浮图标状态" prop="floating_status">
          <el-radio-group v-model="form.floating_status">
            <el-radio :label="1">开启</el-radio>
            <el-radio :label="0">关闭</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="悬浮1图标地址" prop="floating_icon1">
          <div class="icon-input-container">
            <el-input v-model="form.floating_icon1" placeholder="请输入悬浮1图标地址" />
            <img v-if="form.floating_icon1" :src="form.floating_icon1" class="icon-preview" @error="handleImageError" />
          </div>
        </el-form-item>
        <el-form-item label="悬浮2图标地址" prop="floating_icon2">
          <div class="icon-input-container">
            <el-input v-model="form.floating_icon2" placeholder="请输入悬浮2图标地址" />
            <img v-if="form.floating_icon2" :src="form.floating_icon2" class="icon-preview" @error="handleImageError" />
          </div>
        </el-form-item>
        <el-form-item label="悬浮3图标地址" prop="floating_icon3">
          <div class="icon-input-container">
            <el-input v-model="form.floating_icon3" placeholder="请输入悬浮3图标地址" />
            <img v-if="form.floating_icon3" :src="form.floating_icon3" class="icon-preview" @error="handleImageError" />
          </div>
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="form.status">
            <el-radio :label="1">启用</el-radio>
            <el-radio :label="0">禁用</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item>
          <el-button @click="onClose">取消</el-button>
          <el-button type="primary" @click="onSubmit">确定</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, watch } from 'vue'
import { createGameConfig, updateGameConfig } from '@/api/gameConfig'
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

const gameConfigFormRef = ref()

const form = reactive({
  download_url: '',
  audio_url: '',
  website_title: '',
  website_desc: '',
  website_icon: '',
  website_logo: '',
  market_logo: '',
  floating_status: 0,
  floating_icon1: '',
  floating_icon2: '',
  floating_icon3: '',
  status: 1
})

// 监听编辑数据变化
watch(() => props.editData, (newData) => {
  if (newData && Object.keys(newData).length > 0) {
    Object.assign(form, newData)
  }
}, { immediate: true, deep: true })

// 提交
const onSubmit = async () => {
  gameConfigFormRef.value.validate(async (valid) => {
    if (valid) {
      let res
      if (props.id) {
        res = await updateGameConfig({ ...form, id: props.id })
      } else {
        res = await createGameConfig(form)
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
    download_url: '',
    audio_url: '',
    website_title: '',
    website_desc: '',
    website_icon: '',
    website_logo: '',
    market_logo: '',
    floating_status: 0,
    floating_icon1: '',
    floating_icon2: '',
    floating_icon3: '',
    status: 1
  })
}

// 监听ID变化，重置表单
watch(() => props.id, (newId) => {
  if (!newId) {
    resetForm()
  }
})

// 处理图片加载错误
const handleImageError = (event) => {
  event.target.style.display = 'none'
}
</script>

<style scoped>
.gva-table-box {
  padding: 20px;
}

.gva-form-box {
  max-width: 800px;
  margin: 0 auto;
}

.icon-input-container {
  width:100%;
  display: flex;
  align-items: center;
  gap: 12px;
}

.icon-input-container .el-input {
  flex: 1;
}

.icon-preview {
  width: 40px;
  height: 40px;
  object-fit: cover;
  border-radius: 4px;
  border: 1px solid #dcdfe6;
  flex-shrink: 0;
}
</style>
