<template>
  <div class="min-h-screen bg-gradient-to-br from-slate-900 via-purple-900 to-slate-900 relative overflow-hidden">
    <!-- 背景装饰 -->
    <div class="absolute inset-0">
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
      <div class="max-w-4xl w-full space-y-8">
        <!-- 标题区域 -->
        <div class="text-center">
          <h1 class="text-5xl font-bold text-white mb-4 tracking-tight">
             Admin
          </h1>
          <p class="text-xl text-gray-300 mb-8">
            系统初始化配置
          </p>
        </div>

        <!-- 内容卡片 -->
        <div class="bg-white/10 backdrop-blur-lg rounded-2xl shadow-2xl border border-white/20 p-8">
          <!-- 说明页面 -->
          <div
            v-if="!page.showForm"
            :class="[page.showReadme ? 'slide-out-right' : 'slide-in-fwd-top']"
            class="text-center"
          >
            <div class="mb-8">
              <div class="w-16 h-16 bg-gradient-to-r from-blue-500 to-purple-500 rounded-full flex items-center justify-center mx-auto mb-6">
                <svg class="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                </svg>
              </div>
              <h2 class="text-3xl font-bold text-white mb-6">初始化须知</h2>
            </div>

            <div class="space-y-4 text-left max-w-2xl mx-auto">
              <div class="flex items-start space-x-3">
                <div class="w-6 h-6 bg-blue-500 rounded-full flex items-center justify-center flex-shrink-0 mt-0.5">
                  <span class="text-white text-sm font-bold">1</span>
                </div>
                <p class="text-gray-300">您需要有一定的Vue和Golang基础</p>
              </div>
              
              <div class="flex items-start space-x-3">
                <div class="w-6 h-6 bg-blue-500 rounded-full flex items-center justify-center flex-shrink-0 mt-0.5">
                  <span class="text-white text-sm font-bold">2</span>
                </div>
                <p class="text-gray-300">请确保您已了解项目的整体架构和功能模块</p>
              </div>
              
              <div class="flex items-start space-x-3">
                <div class="w-6 h-6 bg-blue-500 rounded-full flex items-center justify-center flex-shrink-0 mt-0.5">
                  <span class="text-white text-sm font-bold">3</span>
                </div>
                <p class="text-gray-300">请您确认是否了解后续的配置流程</p>
              </div>
              
              <div class="flex items-start space-x-3">
                <div class="w-6 h-6 bg-red-500 rounded-full flex items-center justify-center flex-shrink-0 mt-0.5">
                  <span class="text-white text-sm font-bold">4</span>
                </div>
                <p class="text-gray-300">
                  如果您使用MySQL数据库，请确认数据库引擎为
                  <span class="text-red-400 font-bold text-lg ml-1">InnoDB</span>
                </p>
              </div>
              
            </div>

            <div class="flex flex-col sm:flex-row gap-4 justify-center mt-10">
              <el-button 
                type="primary" 
                size="large" 
                @click="showNext"
                class="bg-gradient-to-r from-green-500 to-blue-500 border-0 hover:from-green-600 hover:to-blue-600 px-8 py-3 text-lg font-medium rounded-xl shadow-lg hover:shadow-xl transition-all duration-300"
              >
                我已确认
              </el-button>
            </div>
          </div>

          <!-- 配置表单页面 -->
          <div
            v-if="page.showForm"
            :class="[page.showForm ? 'slide-in-left' : 'slide-out-right']"
            class="max-w-2xl mx-auto"
          >
            <div class="text-center mb-8">
              <div class="w-16 h-16 bg-gradient-to-r from-green-500 to-blue-500 rounded-full flex items-center justify-center mx-auto mb-4">
                <svg class="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z"></path>
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
                </svg>
              </div>
              <h2 class="text-3xl font-bold text-white mb-2">数据库配置</h2>
              <p class="text-gray-300">请填写数据库连接信息</p>
            </div>

            <el-form ref="formRef" :model="form" label-width="120px" size="large" class="space-y-6">
              <el-form-item label="管理员密码" class="mb-6">
                <el-input
                  v-model="form.adminPassword"
                  placeholder="admin账号的默认密码"
                  class="rounded-xl"
                />
              </el-form-item>
              
              <el-form-item label="数据库类型" class="mb-6">
                <el-select
                  v-model="form.dbType"
                  placeholder="请选择"
                  class="w-full rounded-xl"
                  @change="changeDB"
                >
                  <el-option key="mysql" label="MySQL" value="mysql" />
                  <el-option key="pgsql" label="PostgreSQL" value="pgsql" />
                  <el-option key="oracle" label="Oracle" value="oracle" />
                  <el-option key="mssql" label="SQL Server" value="mssql" />
                  <el-option key="sqlite" label="SQLite" value="sqlite" />
                </el-select>
              </el-form-item>

              <el-form-item v-if="form.dbType !== 'sqlite'" label="主机地址" class="mb-6">
                <el-input v-model="form.host" placeholder="请输入数据库链接" class="rounded-xl" />
              </el-form-item>
              
              <el-form-item v-if="form.dbType !== 'sqlite'" label="端口" class="mb-6">
                <el-input v-model="form.port" placeholder="请输入数据库端口" class="rounded-xl" />
              </el-form-item>
              
              <el-form-item v-if="form.dbType !== 'sqlite'" label="用户名" class="mb-6">
                <el-input
                  v-model="form.userName"
                  placeholder="请输入数据库用户名"
                  class="rounded-xl"
                />
              </el-form-item>
              
              <el-form-item v-if="form.dbType !== 'sqlite'" label="密码" class="mb-6">
                <el-input
                  v-model="form.password"
                  placeholder="请输入数据库密码（没有则为空）"
                  type="password"
                  class="rounded-xl"
                />
              </el-form-item>

              <el-form-item label="数据库名" class="mb-6">
                <el-input v-model="form.dbName" placeholder="请输入数据库名称" class="rounded-xl" />
              </el-form-item>
              
              <el-form-item v-if="form.dbType === 'sqlite'" label="数据库路径" class="mb-6">
                <el-input
                  v-model="form.dbPath"
                  placeholder="请输入sqlite数据库文件存放路径"
                  class="rounded-xl"
                />
              </el-form-item>
              
              <el-form-item v-if="form.dbType === 'pgsql'" label="模板" class="mb-6">
                <el-input
                  v-model="form.template"
                  placeholder="请输入postgresql指定template"
                  class="rounded-xl"
                />
              </el-form-item>

              <div class="text-center pt-6">
                <el-button 
                  type="primary" 
                  @click="onSubmit"
                  class="bg-gradient-to-r from-green-500 to-blue-500 border-0 hover:from-green-600 hover:to-blue-600 px-12 py-4 text-lg font-medium rounded-xl shadow-lg hover:shadow-xl transition-all duration-300"
                >
                  立即初始化
                </el-button>
              </div>
            </el-form>
          </div>
        </div>

        <!-- 底部信息 -->
        <div class="text-center text-gray-400 text-sm">
          <p>© 2025  Admin Platform. All rights reserved.</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
  // @ts-ignore
  import { initDB } from '@/api/initdb'
  import { reactive, ref } from 'vue'
  import { ElLoading, ElMessage, ElMessageBox } from 'element-plus'
  import { useRouter } from 'vue-router'

  defineOptions({
    name: 'Init'
  })

  const router = useRouter()

  const page = reactive({
    showReadme: false,
    showForm: false
  })

  const showNext = () => {
    page.showReadme = false
    setTimeout(() => {
      page.showForm = true
    }, 20)
  }

  const out = ref(false)

  const form = reactive({
    adminPassword: '123456',
    dbType: 'mysql',
    host: '127.0.0.1',
    port: '3306',
    userName: 'root',
    password: '',
    dbName: 'goLinker',
    dbPath: ''
  })

  const changeDB = (val) => {
    switch (val) {
      case 'mysql':
        Object.assign(form, {
          adminPassword: '123456',
          reAdminPassword: '',
          dbType: 'mysql',
          host: '127.0.0.1',
          port: '3306',
          userName: 'root',
          password: '',
          dbName: 'gva',
          dbPath: ''
        })
        break
      case 'pgsql':
        Object.assign(form, {
          adminPassword: '123456',
          dbType: 'pgsql',
          host: '127.0.0.1',
          port: '5432',
          userName: 'postgres',
          password: '',
          dbName: 'gva',
          dbPath: '',
          template: 'template0'
        })
        break


      case 'sqlite':
        Object.assign(form, {
          adminPassword: '123456',
          dbType: 'sqlite',
          host: '',
          port: '',
          userName: '',
          password: '',
          dbName: 'gva',
          dbPath: ''
        })
        break
      default:
        Object.assign(form, {
          adminPassword: '123456',
          dbType: 'mysql',
          host: '127.0.0.1',
          port: '3306',
          userName: 'root',
          password: '',
          dbName: 'gva',
          dbPath: ''
        })
    }
  }
  const onSubmit = async () => {
    if (form.adminPassword.length < 6) {
      ElMessage({
        type: 'error',
        message: '密码长度不能小于6位'
      })
      return
    }

    const loading = ElLoading.service({
      lock: true,
      text: '正在初始化数据库，请稍候',
      spinner: 'loading',
      background: 'rgba(0, 0, 0, 0.7)'
    })
    try {
      const res = await initDB(form)
      if (res.code === 0) {
        out.value = true
        ElMessage({
          type: 'success',
          message: res.msg
        })
        
        // 显示AI助手配置提示弹窗
        ElMessageBox.confirm(
          '已经完成基础数据库初始化！',
          '配置完成',
          {
            confirmButtonText: '查看AI配置文档',
            cancelButtonText: '稍后配置',
            type: 'success',
            center: true
          }
        ).then(() => {
          // 点击确认按钮，打开AI配置文档
          router.push({ name: 'Login' })
        }).catch(() => {
          // 点击取消按钮或关闭弹窗，直接跳转到登录页
          router.push({ name: 'Login' })
        })
      }
      loading.close()
    } catch (_) {
      loading.close()
    }
  }
</script>

<style lang="scss" scoped>
  .slide-in-fwd-top {
    -webkit-animation: slide-in-fwd-top 0.6s cubic-bezier(0.25, 0.46, 0.45, 0.94) both;
    animation: slide-in-fwd-top 0.6s cubic-bezier(0.25, 0.46, 0.45, 0.94) both;
  }
  
  .slide-out-right {
    -webkit-animation: slide-out-right 0.5s cubic-bezier(0.55, 0.085, 0.68, 0.53) both;
    animation: slide-out-right 0.5s cubic-bezier(0.55, 0.085, 0.68, 0.53) both;
  }
  
  .slide-in-left {
    -webkit-animation: slide-in-left 0.6s cubic-bezier(0.25, 0.46, 0.45, 0.94) both;
    animation: slide-in-left 0.6s cubic-bezier(0.25, 0.46, 0.45, 0.94) both;
  }

  @keyframes slide-in-fwd-top {
    0% {
      transform: translateY(-50px) scale(0.9);
      opacity: 0;
    }
    100% {
      transform: translateY(0) scale(1);
      opacity: 1;
    }
  }

  @keyframes slide-out-right {
    0% {
      transform: translateX(0) scale(1);
      opacity: 1;
    }
    100% {
      transform: translateX(100px) scale(0.9);
      opacity: 0;
    }
  }

  @keyframes slide-in-left {
    0% {
      transform: translateX(-50px) scale(0.9);
      opacity: 0;
    }
    100% {
      transform: translateX(0) scale(1);
      opacity: 1;
    }
  }

  // 自定义Element Plus样式
  :deep(.el-form-item__label) {
    color: #e2e8f0 !important;
    font-weight: 500;
  }

  :deep(.el-input__wrapper) {
    background-color: rgba(255, 255, 255, 0.1) !important;
    border: 1px solid rgba(255, 255, 255, 0.2) !important;
    box-shadow: none !important;
    border-radius: 12px !important;
    transition: all 0.3s ease !important;
  }

  :deep(.el-input__wrapper:hover) {
    border-color: rgba(59, 130, 246, 0.5) !important;
    background-color: rgba(255, 255, 255, 0.15) !important;
  }

  :deep(.el-input__wrapper.is-focus) {
    border-color: #3b82f6 !important;
    background-color: rgba(255, 255, 255, 0.2) !important;
    box-shadow: 0 0 0 2px rgba(59, 130, 246, 0.2) !important;
  }

  :deep(.el-input__inner) {
    color: #f1f5f9 !important;
    background: transparent !important;
  }

  :deep(.el-input__inner::placeholder) {
    color: #94a3b8 !important;
  }

  :deep(.el-select .el-input__wrapper) {
    background-color: rgba(255, 255, 255, 0.1) !important;
    border: 1px solid rgba(255, 255, 255, 0.2) !important;
  }

  :deep(.el-select-dropdown) {
    background-color: rgba(15, 23, 42, 0.95) !important;
    border: 1px solid rgba(255, 255, 255, 0.2) !important;
    backdrop-filter: blur(10px) !important;
  }

  :deep(.el-select-dropdown__item) {
    color: #e2e8f0 !important;
  }

  :deep(.el-select-dropdown__item:hover) {
    background-color: rgba(59, 130, 246, 0.2) !important;
  }

  :deep(.el-select-dropdown__item.is-selected) {
    background-color: rgba(59, 130, 246, 0.3) !important;
    color: #60a5fa !important;
  }

  // 响应式设计
  @media (max-width: 768px) {
    .max-w-4xl {
      max-width: 95vw !important;
    }
    
    .text-5xl {
      font-size: 2.5rem !important;
    }
    
    .text-3xl {
      font-size: 1.875rem !important;
    }
  }
</style>

