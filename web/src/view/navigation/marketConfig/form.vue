<template>
  <div class="gva-form-box">
    <el-form ref="marketConfigFormRef" :model="form" :rules="rules" label-width="100px">
      <el-form-item label="名称" prop="name">
        <el-input v-model="form.name" placeholder="请输入名称" />
      </el-form-item>
      <el-form-item label="Logo" prop="logo">
        <el-input v-model="form.logo" placeholder="请输入Logo地址" />
      </el-form-item>
      <el-form-item label="跳转地址" prop="jump_url">
        <el-input v-model="form.jump_url" placeholder="请输入跳转地址" />
      </el-form-item>
      <el-form-item label="类型" prop="type">
        <el-radio-group v-model="form.type">
          <el-radio :label="1">可关闭</el-radio>
          <el-radio :label="2">类型2</el-radio>
          <el-radio :label="3">类型3</el-radio>
        </el-radio-group>
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
      <el-form-item label="描述" prop="description">
        <el-input
          v-model="form.description"
          type="textarea"
          :rows="3"
          placeholder="请输入描述"
        />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSubmit">确定</el-button>
        <el-button @click="onClose">取消</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script setup>
import { ref, reactive, watch } from 'vue'
import { createMarketConfig, updateMarketConfig } from '@/api/marketConfig'
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

const marketConfigFormRef = ref()

const form = reactive({
  id: undefined,
  name: '',
  logo: '',
  jump_url: '',
  type: 1,
  sort: 0,
  status: 1,
  is_visible: 1,
  description: ''
})

const rules = reactive({
  name: [{ required: true, message: '名称不能为空', trigger: 'blur' }],
  type: [{ required: true, message: '请选择类型', trigger: 'change' }]
})

// 重置表单
const resetForm = () => {
  Object.assign(form, {
    id: undefined,
    name: '',
    logo: '',
    jump_url: '',
    type: 1,
    sort: 0,
    status: 1,
    is_visible: 1,
    description: ''
  })
}

// 监听编辑数据变化
watch(() => props.editData, (newData) => {
  if (newData && Object.keys(newData).length > 0) {
    Object.assign(form, newData)
  } else {
    resetForm()
  }
}, { immediate: true, deep: true })

// 提交
const onSubmit = async () => {
  marketConfigFormRef.value.validate(async (valid) => {
    if (valid) {
      let res
      if (props.id) {
        res = await updateMarketConfig(form)
      } else {
        res = await createMarketConfig(form)
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
</script>

<style scoped>
.gva-form-box {
  padding: 20px;
}
</style>
