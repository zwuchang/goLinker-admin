<template>
  <div class="gva-table-box">
    <div class="gva-form-box">
      <el-form ref="platformConfigFormRef" :model="form" :rules="rules" label-width="100px">
        <el-form-item label="平台名称" prop="platform_name">
          <el-input v-model="form.platform_name" placeholder="请输入平台名称" />
        </el-form-item>
        <el-form-item label="平台图标" prop="platform_icon">
          <el-input v-model="form.platform_icon" placeholder="请输入平台图标地址" />
        </el-form-item>
        <el-form-item label="平台接口" prop="platform_api">
          <el-input v-model="form.platform_api" placeholder="请输入平台接口地址" />
        </el-form-item>
        <el-form-item label="排序" prop="sort">
          <el-input-number v-model="form.sort" :min="0" :max="999" />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="form.status">
            <el-radio :label="1">启用</el-radio>
            <el-radio :label="0">禁用</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="是否可见" prop="is_visible">
          <el-radio-group v-model="form.is_visible">
            <el-radio :label="1">可见</el-radio>
            <el-radio :label="0">隐藏</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="平台描述" prop="description">
          <el-input
            v-model="form.description"
            type="textarea"
            :rows="3"
            placeholder="请输入平台描述"
          />
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
import { createPlatformConfig, updatePlatformConfig } from '@/api/platformConfig'
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

const platformConfigFormRef = ref()

const form = reactive({
  platform_name: '',
  platform_icon: '',
  platform_api: '',
  sort: 0,
  status: 1,
  is_visible: 1,
  description: ''
})

const rules = {
  platform_name: [
    { required: true, message: '请输入平台名称', trigger: 'blur' }
  ]
}

// 监听编辑数据变化
watch(() => props.editData, (newData) => {
  if (newData && Object.keys(newData).length > 0) {
    Object.assign(form, newData)
  }
}, { immediate: true, deep: true })

// 提交
const onSubmit = async () => {
  platformConfigFormRef.value.validate(async (valid) => {
    if (valid) {
      let res
      if (props.id) {
        res = await updatePlatformConfig(form)
      } else {
        res = await createPlatformConfig(form)
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
    platform_name: '',
    platform_icon: '',
    platform_api: '',
    sort: 0,
    status: 1,
    is_visible: 1,
    description: ''
  })
}

// 监听id变化，重置表单
watch(() => props.id, (newId) => {
  if (newId === 0) {
    resetForm()
  }
})
</script>
