<template>
  <div class="border border-solid border-gray-100 h-full z-10">
    <!-- 可视化编辑器 -->
    <div v-if="!showSourceCode" class="visual-editor">
      <div class="flex justify-between items-center p-2 border-b border-gray-200">
        <Toolbar
          :editor="editorRef"
          :default-config="toolbarConfig"
          mode="default"
        />
        <el-button 
          type="primary" 
          size="small" 
          @click="switchToSourceCode"
          icon="Edit"
        >
          源码编辑
        </el-button>
      </div>
      <Editor
        v-model="valueHtml"
        class="editor-content mt-0.5"
        :default-config="editorConfig"
        mode="default"
        @onCreated="handleCreated"
        @onChange="change"
      />
    </div>
    
    <!-- 源码编辑器 -->
    <div v-else class="source-editor">
      <div class="flex justify-between items-center p-2 border-b border-gray-200">
        <span class="text-sm text-gray-600">HTML源码编辑</span>
        <el-button 
          type="primary" 
          size="small" 
          @click="switchToVisual"
          icon="View"
        >
          可视化编辑
        </el-button>
      </div>
      <div class="source-code-editor">
        <el-input
          v-model="sourceCodeContent"
          type="textarea"
          :rows="20"
          placeholder="请输入HTML源码..."
          @input="onSourceCodeChange"
          class="source-textarea"
        />
      </div>
    </div>
  </div>
</template>

<script setup>
  import '@wangeditor/editor/dist/css/style.css' // 引入 css

  const basePath = import.meta.env.VITE_BASE_API

  import { onBeforeUnmount, ref, shallowRef, watch, nextTick } from 'vue'
  import { Editor, Toolbar } from '@wangeditor/editor-for-vue'

  import { ElMessage } from 'element-plus'
  import { getUrl } from '@/utils/image'
  import { useUserStore } from '@/pinia/modules/user'

  const emits = defineEmits(['change', 'update:modelValue'])

  const change = (editor) => {
    emits('change', editor)
    emits('update:modelValue', valueHtml.value)
  }

  const userStore = useUserStore()
  const props = defineProps({
    modelValue: {
      type: String,
      default: ''
    }
  })

  const editorRef = shallowRef()
  const valueHtml = ref('')
  const showSourceCode = ref(false)
  const sourceCodeContent = ref('')

  // 检查编辑器是否已准备好
  const isEditorReady = () => {
    return editorRef.value && 
           editorRef.value.txt && 
           editorRef.value.txt.html &&
           typeof editorRef.value.txt.html === 'function'
  }

  const toolbarConfig = {}
  const editorConfig = {
    placeholder: '请输入内容...',
    MENU_CONF: {
      // 配置图片上传
      uploadImage: {
        fieldName: 'file',
        server: basePath + '/fileUploadAndDownload/upload?noSave=1',
        headers: {
          'x-token': userStore.token,
        },
        customInsert(res, insertFn) {
          if (res.code === 0) {
            const urlPath = getUrl(res.data.file.url)
            insertFn(urlPath, res.data.file.name)
            return
          }
          ElMessage.error(res.msg)
        }
      }
    }
  }

  // 组件销毁时，也及时销毁编辑器
  onBeforeUnmount(() => {
    const editor = editorRef.value
    if (editor == null) return
    editor.destroy()
  })

  const handleCreated = (editor) => {
    editorRef.value = editor
    valueHtml.value = props.modelValue
    
    // 确保编辑器完全初始化后再设置内容
    if (props.modelValue && editor && editor.txt && editor.txt.html) {
      editor.txt.html(props.modelValue)
    }
  }

  // 切换到源码编辑模式
  const switchToSourceCode = () => {
    try {
      // 获取当前编辑器的内容
      if (isEditorReady()) {
        sourceCodeContent.value = editorRef.value.txt.html()
      } else {
        sourceCodeContent.value = valueHtml.value || props.modelValue || ''
      }
      // 切换到源码模式
      showSourceCode.value = true
    } catch (error) {
      console.error('切换到源码编辑模式时出错:', error)
      showSourceCode.value = true
    }
  }

  // 切换到可视化编辑模式
  const switchToVisual = () => {
    try {
      // 切换到可视化模式
      showSourceCode.value = false
      // 使用nextTick确保DOM更新完成后再设置编辑器内容
      nextTick(() => {
        if (isEditorReady() && sourceCodeContent.value) {
          editorRef.value.txt.html(sourceCodeContent.value)
          change(editorRef.value)
        }
      })
    } catch (error) {
      console.error('切换到可视化编辑模式时出错:', error)
      showSourceCode.value = false
    }
  }

  // XSS防护：过滤危险的HTML标签和属性
  const sanitizeHtml = (html) => {
    if (!html) return ''
    
    // 移除script标签
    html = html.replace(/<script\b[^<]*(?:(?!<\/script>)<[^<]*)*<\/script>/gi, '')
    
    // 移除javascript:协议
    html = html.replace(/javascript:/gi, '')
    
    // 移除on事件属性
    html = html.replace(/\s*on\w+\s*=\s*["'][^"']*["']/gi, '')
    
    // 移除object和embed标签
    html = html.replace(/<(object|embed)\b[^<]*(?:(?!<\/\1>)<[^<]*)*<\/\1>/gi, '')
    
    return html
  }

  // 源码编辑器内容变化处理
  const onSourceCodeChange = (value) => {
    // 应用XSS防护
    const sanitizedValue = sanitizeHtml(value)
    sourceCodeContent.value = sanitizedValue
    // 不需要手动触发事件，watch会自动处理
  }

  watch(
    () => props.modelValue,
    () => {
      valueHtml.value = props.modelValue
    }
  )

  // 监听源码内容变化，在源码编辑模式下同步更新父组件
  watch(sourceCodeContent, (newValue) => {
    if (showSourceCode.value && newValue !== undefined) {
      // 使用setTimeout避免在组件更新过程中触发事件
      setTimeout(() => {
        emits('update:modelValue', newValue)
      }, 0)
    }
  })
</script>

<style scoped lang="scss">
.visual-editor, .source-editor {
  height: 100%;
  min-height: 500px;
  display: flex;
  flex-direction: column;
}

.editor-content {
  flex: 1;
  min-height: 400px;
  overflow-y: auto;
  
  :deep(.w-e-text-container) {
    height: auto !important;
    min-height: 400px;
  }
  
  :deep(.w-e-scroll) {
    height: auto !important;
    min-height: 400px;
  }
}

.source-code-editor {
  padding: 10px;
  flex: 1;
  min-height: 400px;
  overflow-y: auto;
  
  .source-textarea {
    font-family: 'Monaco', 'Menlo', 'Ubuntu Mono', 'Consolas', monospace;
    font-size: 14px;
    line-height: 1.5;
    min-height: 400px;
    
    :deep(.el-textarea__inner) {
      font-family: 'Monaco', 'Menlo', 'Ubuntu Mono', 'Consolas', monospace;
      background-color: #2d3748;
      color: #e2e8f0;
      border: 1px solid #4a5568;
      border-radius: 4px;
      resize: vertical;
      padding: 12px;
      min-height: 400px;
      overflow-y: auto;
      
      &::placeholder {
        color: #a0aec0;
      }
      
      &:focus {
        border-color: #3182ce;
        box-shadow: 0 0 0 2px rgba(49, 130, 206, 0.2);
      }
    }
  }
}

.flex {
  display: flex;
}

.justify-between {
  justify-content: space-between;
}

.items-center {
  align-items: center;
}

.p-2 {
  padding: 8px;
}

.border-b {
  border-bottom: 1px solid #e5e7eb;
}

.text-sm {
  font-size: 14px;
}

.text-gray-600 {
  color: #6b7280;
}
</style>