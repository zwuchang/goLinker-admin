<template>
  <div class="gva-table-box">
    <div class="gva-form-box">
      <el-form ref="gameCategoryFormRef" :model="form" label-width="80px">
        <el-form-item label="类别名称" prop="name">
          <el-input v-model="form.name" placeholder="请输入类别名称" />
        </el-form-item>
        <el-form-item label="类别描述" prop="description">
          <el-input v-model="form.description" type="textarea" placeholder="请输入类别描述" />
        </el-form-item>
        <el-form-item label="排序" prop="sort">
          <el-input-number v-model="form.sort" :min="0" />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="form.status">
            <el-radio :label="1">启用</el-radio>
            <el-radio :label="0">禁用</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="图标" prop="icon">
          <el-input v-model="form.icon" placeholder="请输入图标名称" />
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
import { createGameCategory, updateGameCategory } from '@/api/gameCategory'
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

const gameCategoryFormRef = ref()

const form = reactive({
  name: '',
  description: '',
  sort: 0,
  status: 1,
  icon: ''
})

// 监听编辑数据变化
watch(() => props.editData, (newData) => {
  if (newData && Object.keys(newData).length > 0) {
    Object.assign(form, newData)
  }
}, { immediate: true, deep: true })

// 提交
const onSubmit = async () => {
  gameCategoryFormRef.value.validate(async (valid) => {
    if (valid) {
      let res
      if (props.id) {
        res = await updateGameCategory(form)
      } else {
        res = await createGameCategory(form)
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
    name: '',
    description: '',
    sort: 0,
    status: 1,
    icon: ''
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
  max-width: 600px;
  margin: 0 auto;
}
</style>
