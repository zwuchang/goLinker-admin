<template>
  <div class="profile-container">
    <!-- 顶部个人信息卡片 -->
    <div class="bg-white dark:bg-slate-800 rounded-2xl shadow-sm mb-8">
      <!-- 顶部背景图 -->
      <div class="h-48 bg-blue-50 dark:bg-slate-600 relative">
        <div class="absolute inset-0 bg-pattern opacity-7"></div>
      </div>

      <!-- 个人信息区 -->
      <div class="px-8 -mt-20 pb-8">
        <div class="flex flex-col lg:flex-row items-start gap-8">
          <!-- 左侧头像 -->
          <div class="profile-avatar-wrapper flex-shrink-0 mx-auto lg:mx-0">
            <SelectImage
                v-model="userInfo.headerImg"
                file-type="image"
                rounded
            />
          </div>

          <!-- 右侧信息 -->
          <div class="flex-1 pt-12 lg:pt-20 w-full">
            <div class="flex flex-col lg:flex-row items-start lg:items-start justify-between gap-4">
              <div class="lg:mt-4">
                <div class="flex items-center gap-4 mb-4">
                  <div
                    v-if="!editFlag"
                    class="text-2xl font-bold flex items-center gap-3 text-gray-800 dark:text-gray-100"
                  >
                    {{ userInfo.nickName || '未设置昵称' }}
                    <el-icon
                      class="cursor-pointer text-gray-400 hover:text-gray-600 dark:hover:text-gray-200 transition-colors duration-200"
                      @click="openEdit"
                    >
                      <edit />
                    </el-icon>
                  </div>
                  <div v-else class="flex items-center">
                    <el-input v-model="nickName" class="w-48 mr-4" placeholder="请输入昵称" />
                    <el-button type="primary" plain @click="enterEdit" :loading="nickNameLoading">
                      确认
                    </el-button>
                    <el-button type="danger" plain @click="closeEdit">
                      取消
                    </el-button>
                  </div>
                </div>

                <div class="text-gray-500 dark:text-gray-400 text-sm">
                  <div class="flex items-center gap-2 mb-2">
                    <el-icon><user /></el-icon>
                    <span>用户ID: {{ userInfo.uuid || '未知' }}</span>
                  </div>
                  <div class="flex items-center gap-2">
                    <el-icon><medal /></el-icon>
                    <span>用户角色: {{ userInfo.authority?.authorityName || '普通用户' }}</span>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- 主要内容区 -->
    <div class="grid lg:grid-cols-12 md:grid-cols-1 gap-8">
      <!-- 左侧信息栏 -->
      <div class="lg:col-span-4">
        <div class="bg-white dark:bg-slate-800 rounded-xl p-6 mb-6 profile-card">
          <h2 class="text-lg font-semibold mb-4 flex items-center gap-2">
            <el-icon class="text-blue-500"><info-filled /></el-icon>
            账号信息
          </h2>
          <div class="space-y-4">
            <div class="flex items-center gap-1 lg:gap-3 text-gray-600 dark:text-gray-300">
              <el-icon class="text-blue-500"><calendar /></el-icon>
              <span class="font-medium">注册时间：</span>
              <span>{{ formatDate(userInfo.CreatedAt) }}</span>
            </div>
            <div class="flex items-center gap-1 lg:gap-3 text-gray-600 dark:text-gray-300">
              <el-icon class="text-green-500"><clock /></el-icon>
              <span class="font-medium">最后更新：</span>
              <span>{{ formatDate(userInfo.UpdatedAt) }}</span>
            </div>
            <div class="flex items-center gap-1 lg:gap-3 text-gray-600 dark:text-gray-300">
              <el-icon class="text-purple-500"><timer /></el-icon>
              <span class="font-medium">使用天数：</span>
              <span>{{ formatDaysSince(userInfo.CreatedAt) }} 天</span>
            </div>
            <div class="flex items-center gap-1 lg:gap-3 text-gray-600 dark:text-gray-300">
              <el-icon class="text-yellow-500"><check /></el-icon>
              <span class="font-medium">账号状态：</span>
              <el-tag :type="userInfo.enable === 1 ? 'success' : 'danger'" size="small">
                {{ userInfo.enable === 1 ? '正常' : '已禁用' }}
              </el-tag>
            </div>
          </div>
        </div>

        <!-- 安全设置 -->
        <div class="bg-white dark:bg-slate-800 rounded-xl p-6 profile-card">
          <h2 class="text-lg font-semibold mb-4 flex items-center gap-2">
            <el-icon class="text-blue-500"><lock /></el-icon>
            安全设置
          </h2>
          <div class="space-y-4">
            <div class="flex items-center justify-between">
              <div class="flex items-center gap-3">
                <el-icon class="text-purple-500"><key /></el-icon>
                <span class="font-medium">登录密码</span>
              </div>
              <el-button
                type="primary"
                plain
                @click="showPassword = true"
              >
                修改密码
              </el-button>
            </div>
          </div>
        </div>
      </div>

      <!-- 右侧内容区 -->
      <div class="lg:col-span-8">
        <div class="bg-white dark:bg-slate-800 rounded-xl p-6 profile-card">
          <h2 class="text-lg font-semibold mb-6 flex items-center gap-2">
            <el-icon class="text-blue-500"><data-line /></el-icon>
            账号统计
          </h2>
          <div class="grid grid-cols-2 md:grid-cols-4 gap-4 lg:gap-6">
            <div class="stat-card">
              <div class="text-2xl lg:text-4xl font-bold text-blue-500 mb-2">
                {{ userInfo.ID || 0 }}
              </div>
              <div class="text-gray-500 text-sm">用户ID</div>
            </div>
            <div class="stat-card">
              <div class="text-2xl lg:text-4xl font-bold text-green-500 mb-2">
                {{ formatDaysSince(userInfo.CreatedAt) }}
              </div>
              <div class="text-gray-500 text-sm">使用天数</div>
            </div>
            <div class="stat-card">
              <div class="text-2xl lg:text-4xl font-bold text-purple-500 mb-2">
                {{ userInfo.authority?.authorityName || '用户' }}
              </div>
              <div class="text-gray-500 text-sm">用户角色</div>
            </div>
            <div class="stat-card">
              <div class="text-2xl lg:text-4xl font-bold text-yellow-500 mb-2">
                {{ userInfo.enable === 1 ? '正常' : '禁用' }}
              </div>
              <div class="text-gray-500 text-sm">账号状态</div>
            </div>
          </div>
        </div>


      </div>
    </div>

    <!-- 修改密码弹窗 -->
    <el-dialog
      v-model="showPassword"
      title="修改密码"
      width="400px"
      class="custom-dialog"
      @close="clearPassword"
    >
      <el-form
        ref="modifyPwdForm"
        :model="pwdModify"
        :rules="rules"
        label-width="90px"
        class="py-4"
      >
        <el-form-item :minlength="6" label="原密码" prop="password">
          <el-input v-model="pwdModify.password" show-password placeholder="请输入原密码" />
        </el-form-item>
        <el-form-item :minlength="6" label="新密码" prop="newPassword">
          <el-input v-model="pwdModify.newPassword" show-password placeholder="请输入新密码" />
        </el-form-item>
        <el-form-item :minlength="6" label="确认密码" prop="confirmPassword">
          <el-input v-model="pwdModify.confirmPassword" show-password placeholder="请再次输入新密码" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="showPassword = false">取 消</el-button>
          <el-button type="primary" @click="savePassword" :loading="passwordLoading">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { setSelfInfo, changePassword, getUserInfo } from '@/api/user.js'
import { reactive, ref, watch, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { useUserStore } from '@/pinia/modules/user'
import SelectImage from '@/components/selectImage/selectImage.vue'

defineOptions({
  name: 'Person'
})

const userStore = useUserStore()
const modifyPwdForm = ref(null)
const showPassword = ref(false)
const pwdModify = ref({})
const nickName = ref('')
const editFlag = ref(false)

// 用户信息 - 初始化默认值
const userInfo = ref({
  ID: 0,
  CreatedAt: '',
  UpdatedAt: '',
  uuid: '',
  userName: '',
  nickName: '',
  headerImg: '',
  authorityId: 0,
  authority: null,
  authorities: [],
  phone: '',
  email: '',
  enable: 1
})

// 加载状态
const passwordLoading = ref(false)
const nickNameLoading = ref(false)
const detailLoading = ref(false)

const rules = reactive({
  password: [
    { required: true, message: '请输入密码', trigger: 'blur' },
    { min: 6, message: '最少6个字符', trigger: 'blur' }
  ],
  newPassword: [
    { required: true, message: '请输入新密码', trigger: 'blur' },
    { min: 6, message: '最少6个字符', trigger: 'blur' }
  ],
  confirmPassword: [
    { required: true, message: '请输入确认密码', trigger: 'blur' },
    { min: 6, message: '最少6个字符', trigger: 'blur' },
    {
      validator: (rule, value, callback) => {
        if (value !== pwdModify.value.newPassword) {
          callback(new Error('两次密码不一致'))
        } else {
          callback()
        }
      },
      trigger: 'blur'
    }
  ]
})

// 获取用户详细信息
const fetchUserInfo = async () => {
  detailLoading.value = true
  try {
    const res = await getUserInfo()
    if (res.code === 0) {
      userInfo.value = res.data.userInfo
    }
  } catch (error) {
    ElMessage.error('获取用户信息失败')
  } finally {
    detailLoading.value = false
  }
}

// 修改密码
const savePassword = async () => {
  modifyPwdForm.value.validate(async (valid) => {
    if (valid) {
      passwordLoading.value = true
      try {
        const res = await changePassword({
          password: pwdModify.value.password,
          newPassword: pwdModify.value.newPassword
        })
        if (res.code === 0) {
          ElMessage.success('修改密码成功！')
          showPassword.value = false
        }
      } catch (error) {
        ElMessage.error('修改密码失败')
      } finally {
        passwordLoading.value = false
      }
    }
  })
}

const clearPassword = () => {
  pwdModify.value = {
    password: '',
    newPassword: '',
    confirmPassword: ''
  }
  modifyPwdForm.value?.clearValidate()
}

// 修改昵称
const openEdit = () => {
  nickName.value = userInfo.value.nickName || ''
  editFlag.value = true
}

const closeEdit = () => {
  nickName.value = ''
  editFlag.value = false
}

const enterEdit = async () => {
  nickNameLoading.value = true
  try {
    const res = await setSelfInfo({
      nickName: nickName.value
    })
    if (res.code === 0) {
      userInfo.value.nickName = nickName.value
      userStore.ResetUserInfo({ nickName: nickName.value })
      ElMessage.success('修改成功')
      editFlag.value = false
    }
  } catch (error) {
    ElMessage.error('修改失败')
  } finally {
    nickNameLoading.value = false
  }
}

// 头像修改监听
watch(() => userInfo.value.headerImg, async(val) => {
  const res = await setSelfInfo({ headerImg: val })
  if (res.code === 0) {
    userStore.ResetUserInfo({ headerImg: val })
    ElMessage({
      type: 'success',
      message: '设置成功',
    })
  }
})

// 工具函数
const formatDate = (date) => {
  if (!date) return '未知'
  return new Date(date).toLocaleString('zh-CN')
}

const formatDaysSince = (date) => {
  if (!date) return 0
  const now = new Date()
  const created = new Date(date)
  const diffTime = Math.abs(now - created)
  const diffDays = Math.ceil(diffTime / (1000 * 60 * 60 * 24))
  return diffDays
}

// 页面加载时获取用户详细信息
onMounted(() => {
  fetchUserInfo()
})
</script>

<style lang="scss">
.profile-container {
  @apply p-4 lg:p-6 min-h-screen bg-gray-50 dark:bg-slate-900;

  .bg-pattern {
    background-image: url("data:image/svg+xml,%3Csvg width='60' height='60' viewBox='0 0 60 60' xmlns='http://www.w3.org/2000/svg'%3E%3Cg fill='none' fill-rule='evenodd'%3E%3Cg fill='%23000000' fill-opacity='0.1'%3E%3Cpath d='M36 34v-4h-2v4h-4v2h4v4h2v-4h4v-2h-4zm0-30V0h-2v4h-4v2h4v4h2V6h4V4h-4zM6 34v-4H4v4H0v2h4v4h2v-4h4v-2H6zM6 4V0H4v4H0v2h4v4h2V6h4V4H6z'/%3E%3C/g%3E%3C/g%3E%3C/svg%3E");
  }

  .profile-card {
    @apply shadow-sm hover:shadow-md transition-shadow duration-300;
  }

  .stat-card {
    @apply p-4 lg:p-6 rounded-lg bg-gray-50 dark:bg-slate-700/50 text-center hover:shadow-md transition-all duration-300;
  }

  .custom-dialog {
    :deep(.el-dialog__header) {
      @apply mb-0 pb-4 border-b border-gray-100 dark:border-gray-700;
    }
    :deep(.el-dialog__footer) {
      @apply mt-0 pt-4 border-t border-gray-100 dark:border-gray-700;
    }
    :deep(.el-input__wrapper) {
      @apply shadow-none;
    }
    :deep(.el-input__prefix) {
      @apply mr-2;
    }
  }
}
</style>