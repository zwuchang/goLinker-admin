<template>
  <div class="gva-table-box">
    <div class="gva-form-box">
      <el-form ref="gameConfigFormRef" :model="form" label-width="100px">
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
          <el-input v-model="form.website_icon" placeholder="请输入网站图标URL" />
        </el-form-item>
        <el-form-item label="网站Logo" prop="website_logo">
          <el-input v-model="form.website_logo" placeholder="请输入网站Logo URL" />
        </el-form-item>
        <el-form-item label="市场Logo" prop="market_logo">
          <el-input v-model="form.market_logo" placeholder="请输入市场Logo URL" />
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
    status: 1
  })
}

// 监听ID变化，重置表单
watch(() => props.id, (newId) => {
  if (!newId) {
    resetForm()
  }
})
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
