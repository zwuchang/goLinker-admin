<template>
  <div class="gva-table-box">
    <div class="gva-form-box">
      <el-form ref="contactMethodFormRef" :model="form" label-width="120px">
        <el-form-item label="显示名称" prop="displayName">
          <el-input v-model="form.displayName" placeholder="如：Twitter、Facebook、微信等" />
        </el-form-item>
        <el-form-item label="跳转地址" prop="jumpUrl">
          <el-input v-model="form.jumpUrl" placeholder="请输入跳转地址" />
        </el-form-item>
        <el-form-item label="联系方式图标" prop="image">
          <div class="icon-input-container">
            <el-input v-model="form.image" placeholder="请输入联系方式图标URL" />
            <img v-if="form.image" :src="form.image" class="icon-preview" @error="handleImageError" />
          </div>
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
      </el-form>
    </div>
    <div class="dialog-footer">
      <el-button @click="closeDialog">取 消</el-button>
      <el-button type="primary" @click="enterDialog">确 定</el-button>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, watch } from 'vue'
import { ElMessage } from 'element-plus'
import {
  createContactMethod,
  updateContactMethod,
  getContactMethod
} from '@/api/contactMethod'

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

const contactMethodFormRef = ref()
const form = reactive({
  displayName: '',
  jumpUrl: '',
  image: '',
  sort: 0,
  status: 1
})

const rules = reactive({
  displayName: [
    { required: true, message: '请输入显示名称', trigger: 'blur' }
  ],
  jumpUrl: [
    { required: true, message: '请输入跳转地址', trigger: 'blur' }
  ]
})

// 监听编辑数据变化
watch(() => props.editData, (newData) => {
  if (newData && Object.keys(newData).length > 0) {
    Object.assign(form, newData)
  }
}, { immediate: true, deep: true })

// 监听ID变化，获取详情
watch(() => props.id, async (newId) => {
  if (newId && newId > 0) {
    try {
      const res = await getContactMethod({ ID: newId })
      if (res.code === 0) {
        Object.assign(form, res.data.contactMethod)
      }
    } catch (error) {
      console.error('获取详情失败:', error)
      ElMessage.error('获取详情失败')
    }
  }
}, { immediate: true })

const handleImageError = (e) => {
  e.target.style.display = 'none'
}

const closeDialog = () => {
  emit('close')
}

const enterDialog = async () => {
  const valid = await contactMethodFormRef.value?.validate()
  if (!valid) return
  
  try {
    let res
    if (props.id && props.id > 0) {
      res = await updateContactMethod({ ...form, ID: props.id })
    } else {
      res = await createContactMethod(form)
    }

    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: props.id && props.id > 0 ? '更新成功' : '创建成功'
      })
      emit('success')
      closeDialog()
    }
  } catch (error) {
    console.error('操作失败:', error)
    ElMessage.error('操作失败')
  }
}
</script>

<style scoped>
.gva-form-box {
  padding: 20px;
}

.icon-input-container {
  display: flex;
  align-items: center;
  gap: 10px;
}

.icon-preview {
  width: 40px;
  height: 40px;
  object-fit: cover;
  border-radius: 4px;
  border: 1px solid #dcdfe6;
}

.dialog-footer {
  text-align: right;
  padding: 20px;
  border-top: 1px solid #ebeef5;
}
</style>
