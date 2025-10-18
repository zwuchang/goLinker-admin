<template>
  <div>
    <div class="gva-table-box">
      <el-card header="联系配置">
        <el-form :model="form" label-width="120px" :rules="rules" ref="formRef">
          <el-form-item label="客服邮箱" prop="email">
            <el-input v-model="form.email" placeholder="请输入客服邮箱" />
          </el-form-item>
          
          <el-form-item label="横幅图片">
            <el-input v-model="form.bannerImage" placeholder="请输入联系客服横幅图片" />
          </el-form-item>
          
          <el-form-item>
            <el-button type="primary" @click="saveConfig" :loading="loading">保存配置</el-button>
          </el-form-item>
        </el-form>
      </el-card>
    </div>
  </div>
</template>

<script setup>
import { getContactConfig, updateContactConfig } from '@/api/contactConfig'
import SelectImage from '@/components/selectImage/selectImage.vue'
import { ref, onMounted } from 'vue'
import { ElMessage } from 'element-plus'

defineOptions({
  name: 'ContactConfig'
})

const form = ref({
  email: '',
  bannerImage: ''
})

const rules = ref({
  email: [
    { required: true, message: '请输入客服邮箱', trigger: 'blur' },
    { type: 'email', message: '请输入正确的邮箱格式', trigger: 'blur' }
  ]
})

const formRef = ref()
const loading = ref(false)

const getConfig = async () => {
  try {
    const res = await getContactConfig()
    if (res.code === 0) {
      form.value = { ...res.data.contactConfig }
    }
  } catch (error) {
    console.error('获取配置失败:', error)
    ElMessage.error('获取配置失败')
  }
}

const saveConfig = async () => {
  const valid = await formRef.value?.validate()
  if (!valid) return
  
  loading.value = true
  try {
    const res = await updateContactConfig(form.value)
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '保存成功'
      })
    }
  } catch (error) {
    console.error('保存失败:', error)
    ElMessage.error('保存失败')
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  getConfig()
})
</script>

<style scoped>
.gva-table-box {
  padding: 20px;
}

.el-card {
  max-width: 800px;
}
</style>
