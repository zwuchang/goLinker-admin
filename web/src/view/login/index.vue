<template>
  <div id="userLayout" class="min-h-screen bg-gradient-to-br from-slate-900 via-purple-900 to-slate-900 relative overflow-hidden">
    <!-- 背景装饰元素 -->
    <div class="absolute inset-0 overflow-hidden">
      <!-- 几何图形装饰 -->
      <div class="absolute top-20 left-20 w-72 h-72 bg-gradient-to-r from-blue-500/20 to-purple-500/20 rounded-full blur-3xl animate-pulse"></div>
      <div class="absolute bottom-20 right-20 w-96 h-96 bg-gradient-to-r from-pink-500/20 to-blue-500/20 rounded-full blur-3xl animate-pulse delay-1000"></div>
      <div class="absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 w-64 h-64 bg-gradient-to-r from-cyan-500/10 to-purple-500/10 rounded-full blur-2xl animate-pulse delay-500"></div>
      
      <!-- 网格背景 -->
      <div class="absolute inset-0 opacity-30">
        <div class="absolute inset-0" style="background-image: radial-gradient(circle at 1px 1px, rgba(156, 146, 172, 0.15) 1px, transparent 0); background-size: 20px 20px;"></div>
      </div>
    </div>

    <!-- 主要内容区域 -->
    <div class="relative z-10 min-h-screen flex items-center justify-center px-4 sm:px-6 lg:px-8">
      <div class="max-w-md w-full space-y-8">
        <!-- Logo和标题区域 -->
        <div class="text-center">
          <div class="flex justify-center mb-6">
            <div class="relative">
              <div class="absolute inset-0 bg-gradient-to-r from-blue-500 to-purple-600 rounded-2xl blur opacity-75"></div>
              <div class="relative bg-white/10 backdrop-blur-sm rounded-2xl p-6 border border-white/20">
                <img class="w-16 h-16 mx-auto" :src="$GIN_VUE_ADMIN.appLogo" alt="Logo" />
              </div>
            </div>
          </div>
          <h1 class="text-4xl font-bold text-white mb-2 bg-gradient-to-r from-white to-gray-300 bg-clip-text text-transparent">
            {{ $GIN_VUE_ADMIN.appName }}
          </h1>
          <p class="text-gray-300 text-lg">
            A management platform using Golang and Vue
          </p>
        </div>

        <!-- 登录表单 -->
        <div class="bg-white/10 backdrop-blur-lg rounded-3xl p-8 border border-white/20 shadow-2xl">
          <el-form
            ref="loginForm"
            :model="loginFormData"
            :rules="rules"
            :validate-on-rule-change="false"
            @keyup.enter="submitForm"
            class="space-y-6"
          >
            <el-form-item prop="username" class="mb-6">
              <el-input
                v-model="loginFormData.username"
                size="large"
                placeholder="请输入用户名"
                class="custom-input"
                :prefix-icon="User"
              />
            </el-form-item>
            
            <el-form-item prop="password" class="mb-6">
              <el-input
                v-model="loginFormData.password"
                show-password
                size="large"
                type="password"
                placeholder="请输入密码"
                class="custom-input"
                :prefix-icon="Lock"
              />
            </el-form-item>
            
            <el-form-item
              v-if="loginFormData.openCaptcha"
              prop="captcha"
              class="mb-6"
            >
              <div class="flex gap-3">
                <el-input
                  v-model="loginFormData.captcha"
                  placeholder="请输入验证码"
                  size="large"
                  class="custom-input flex-1"
                />
                <div class="w-32 h-12 bg-white/20 rounded-lg overflow-hidden cursor-pointer hover:bg-white/30 transition-all duration-200">
                  <img
                    v-if="picPath"
                    class="w-full h-full object-cover"
                    :src="picPath"
                    alt="验证码"
                    @click="loginVerify()"
                  />
                </div>
              </div>
            </el-form-item>
            
            <el-form-item class="mb-4">
              <el-button
                class="w-full h-12 text-lg font-semibold bg-gradient-to-r from-blue-600 to-purple-600 hover:from-blue-700 hover:to-purple-700 border-0 rounded-xl shadow-lg hover:shadow-xl transform hover:scale-[1.02] transition-all duration-200"
                type="primary"
                size="large"
                @click="submitForm"
                :loading="isLoading"
              >
                登 录
              </el-button>
            </el-form-item>
            
            <el-form-item class="mb-0" v-if="showInitDb">
              <el-button
                class="w-full h-12 text-lg font-semibold bg-white/20 hover:bg-white/30 text-white border border-white/30 rounded-xl transition-all duration-200"
                size="large"
                @click="checkInit"
              >
                前往初始化
              </el-button>
            </el-form-item>
          </el-form>
        </div>

        <!-- 底部信息 -->
        <div class="text-center text-gray-400 text-sm">
          <p>© 2025 导航后台系统. All rights reserved.</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
  import { captcha } from '@/api/user'
  import { checkDB } from '@/api/initdb'
  import { reactive, ref } from 'vue'
  import { ElMessage } from 'element-plus'
  import { useRouter } from 'vue-router'
  import { useUserStore } from '@/pinia/modules/user'
  import { User, Lock } from '@element-plus/icons-vue'

  defineOptions({
    name: 'Login'
  })

  const router = useRouter()
  const isLoading = ref(false)

  const showInitDb = ref(false)

  // 如果是dev环境，则显示初始化数据库按钮
  if (import.meta.env.MODE === 'development') {
    showInitDb.value = true
  }
  
  // 验证函数
  const checkUsername = (rule, value, callback) => {
    if (value.length < 5) {
      return callback(new Error('请输入正确的用户名'))
    } else {
      callback()
    }
  }
  const checkPassword = (rule, value, callback) => {
    if (value.length < 6) {
      return callback(new Error('请输入正确的密码'))
    } else {
      callback()
    }
  }

  // 获取验证码
  const loginVerify = async () => {
    const ele = await captcha()
    rules.captcha.push({
      max: ele.data.captchaLength,
      min: ele.data.captchaLength,
      message: `请输入${ele.data.captchaLength}位验证码`,
      trigger: 'blur'
    })
    picPath.value = ele.data.picPath
    loginFormData.captchaId = ele.data.captchaId
    loginFormData.openCaptcha = ele.data.openCaptcha
  }
  loginVerify()

  // 登录相关操作
  const loginForm = ref(null)
  const picPath = ref('')
  const loginFormData = reactive({
    username: '',
    password: '',
    captcha: '',
    captchaId: '',
    openCaptcha: false
  })

  if (import.meta.env.MODE === 'development') {
    loginFormData.username = 'admin'
  }

  const rules = reactive({
    username: [{ validator: checkUsername, trigger: 'blur' }],
    password: [{ validator: checkPassword, trigger: 'blur' }],
    captcha: [
      {
        message: '验证码格式不正确',
        trigger: 'blur'
      }
    ]
  })

  const userStore = useUserStore()
  const login = async () => {
    return await userStore.LoginIn(loginFormData)
  }
  const submitForm = () => {
    loginForm.value.validate(async (v) => {
      if (!v) {
        // 未通过前端静态验证
        ElMessage({
          type: 'error',
          message: '请正确填写登录信息',
          showClose: true
        })
        await loginVerify()
        return false
      }

      // 通过验证，请求登陆
      isLoading.value = true
      try {
        const flag = await login()

        // 登陆失败，刷新验证码
        if (!flag) {
          await loginVerify()
          return false
        }

        // 登陆成功
        return true
      } finally {
        isLoading.value = false
      }
    })
  }

  // 跳转初始化
  const checkInit = async () => {
    const res = await checkDB()
    if (res.code === 0) {
      if (res.data?.needInit) {
        userStore.NeedInit()
        await router.push({ name: 'Init' })
      } else {
        ElMessage({
          type: 'info',
          message: '已配置数据库信息，无法初始化'
        })
      }
    }
  }
</script>

<style scoped>
  /* 自定义输入框样式 */
  :deep(.custom-input .el-input__wrapper) {
    background: rgba(255, 255, 255, 0.1) !important;
    border: 1px solid rgba(255, 255, 255, 0.2) !important;
    border-radius: 12px !important;
    box-shadow: none !important;
    backdrop-filter: blur(10px);
    transition: all 0.3s ease;
  }
  
  :deep(.custom-input .el-input__wrapper:hover) {
    border-color: rgba(255, 255, 255, 0.4) !important;
    background: rgba(255, 255, 255, 0.15) !important;
  }
  
  :deep(.custom-input .el-input__wrapper.is-focus) {
    border-color: rgba(59, 130, 246, 0.6) !important;
    background: rgba(255, 255, 255, 0.2) !important;
    box-shadow: 0 0 0 2px rgba(59, 130, 246, 0.2) !important;
  }
  
  :deep(.custom-input .el-input__inner) {
    color: white !important;
    background: transparent !important;
  }
  
  :deep(.custom-input .el-input__inner::placeholder) {
    color: rgba(255, 255, 255, 0.6) !important;
  }
  
  :deep(.custom-input .el-input__prefix) {
    color: rgba(255, 255, 255, 0.7) !important;
  }
  
  :deep(.custom-input .el-input__suffix) {
    color: rgba(255, 255, 255, 0.7) !important;
  }
  
  /* 表单验证错误样式 */
  :deep(.el-form-item__error) {
    color: #fca5a5 !important;
    font-size: 12px;
  }
  
  /* 按钮悬停效果 */
  .el-button:hover {
    transform: translateY(-1px);
  }
  
  /* 响应式设计 */
  @media (max-width: 640px) {
    .min-h-screen {
      padding: 1rem;
    }
  }
</style>
