<template>
  <div class="image-upload">
    <div class="upload-area" :class="{ 'drag-over': isDragOver }" v-if="!imageUrl" @click="triggerUpload" @dragover.prevent="isDragOver = true" @dragleave.prevent="isDragOver = false" @drop.prevent="handleDrop">
      <svg xmlns="http://www.w3.org/2000/svg" width="32" height="32" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round">
        <rect x="3" y="3" width="18" height="18" rx="2" ry="2"></rect>
        <circle cx="8.5" cy="8.5" r="1.5"></circle>
        <polyline points="21 15 16 10 5 21"></polyline>
      </svg>
      <span class="upload-text">点击或拖拽上传图片</span>
      <span class="upload-hint">支持 JPG、PNG、GIF、WebP，最大 5MB</span>
      <input
        ref="fileInputRef"
        type="file"
        accept="image/jpeg,image/png,image/gif,image/webp"
        class="file-input"
        @change="handleFileChange"
      />
    </div>

    <div class="preview-area" v-else>
      <img :src="imageUrl" class="preview-image" alt="封面预览" />
      <div class="preview-overlay">
        <button class="preview-btn" @click="triggerUpload" type="button">
          <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <path d="M11 4H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7"></path>
            <path d="M18.5 2.5a2.121 2.121 0 0 1 3 3L12 15l-4 1 1-4 9.5-9.5z"></path>
          </svg>
        </button>
        <button class="preview-btn danger" @click="handleRemove" type="button">
          <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <polyline points="3 6 5 6 21 6"></polyline>
            <path d="M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6m3 0V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2"></path>
          </svg>
        </button>
      </div>
    </div>

    <div class="upload-progress" v-if="uploading">
      <div class="progress-bar">
        <div class="progress-fill"></div>
      </div>
      <span class="progress-text">上传中...</span>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { ElMessage } from 'element-plus'
import { uploadFile, MEDIA_CATEGORIES } from '../../api/media'

const props = defineProps({
  modelValue: {
    type: String,
    default: ''
  },
  category: {
    type: String,
    default: MEDIA_CATEGORIES.COMMON,
    validator: (v) => Object.values(MEDIA_CATEGORIES).includes(v)
  }
})

const emit = defineEmits(['update:modelValue'])

const fileInputRef = ref(null)
const uploading = ref(false)
const isDragOver = ref(false)

const imageUrl = computed({
  get: () => props.modelValue,
  set: (val) => emit('update:modelValue', val)
})

const triggerUpload = () => {
  fileInputRef.value?.click()
}

const handleFileChange = (e) => {
  const file = e.target.files?.[0]
  if (file) {
    uploadImage(file)
  }
  // 重置 input，允许重复选择同一文件
  e.target.value = ''
}

const handleDrop = (e) => {
  isDragOver.value = false
  const file = e.dataTransfer.files?.[0]
  if (file) {
    uploadImage(file)
  }
}

const uploadImage = async (file) => {
  // 前端校验
  const allowedTypes = ['image/jpeg', 'image/png', 'image/gif', 'image/webp']
  if (!allowedTypes.includes(file.type)) {
    ElMessage.error('只支持 JPG、PNG、GIF、WebP 格式的图片')
    return
  }
  if (file.size > 5 * 1024 * 1024) {
    ElMessage.error('图片大小不能超过 5MB')
    return
  }

  uploading.value = true
  try {
    const res = await uploadFile(file, props.category)
    if (res.code === 0) {
      imageUrl.value = res.data.url
      ElMessage.success('上传成功')
    } else {
      ElMessage.error(res.message || '上传失败')
    }
  } catch (e) {
    console.error('上传失败:', e)
    ElMessage.error('上传失败，请重试')
  } finally {
    uploading.value = false
  }
}

const handleRemove = () => {
  imageUrl.value = ''
}
</script>

<style scoped>
.image-upload {
  display: inline-block;
}

.upload-area {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  width: 200px;
  height: 140px;
  border: 2px dashed var(--card-separator-color);
  border-radius: var(--card-border-radius, 8px);
  cursor: pointer;
  transition: border-color 0.2s, background-color 0.2s;
  background: rgba(var(--accent-color-rgb), 0.03);
}

.upload-area:hover {
  border-color: var(--accent-color);
  background: rgba(var(--accent-color-rgb, 64, 158, 255), 0.04);
}

.upload-area.drag-over {
  border-color: var(--accent-color);
  background: rgba(var(--accent-color-rgb, 64, 158, 255), 0.08);
  transform: scale(1.02);
}

.upload-area svg {
  color: var(--card-text-color-tertiary);
  margin-bottom: 8px;
}

.upload-text {
  font-size: 13px;
  color: var(--card-text-color-secondary);
  margin-bottom: 4px;
}

.upload-hint {
  font-size: 11px;
  color: var(--card-text-color-tertiary);
}

.file-input {
  display: none;
}

.preview-area {
  position: relative;
  width: 200px;
  height: 140px;
  border-radius: var(--card-border-radius, 8px);
  overflow: hidden;
}

.preview-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
  display: block;
}

.preview-overlay {
  position: absolute;
  inset: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  background: rgba(0, 0, 0, 0.4);
  opacity: 0;
  transition: opacity 0.2s;
}

.preview-area:hover .preview-overlay {
  opacity: 1;
}

.preview-btn {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 32px;
  height: 32px;
  border: none;
  border-radius: 50%;
  background: rgba(255, 255, 255, 0.9);
  color: var(--card-text-color-secondary);
  cursor: pointer;
  transition: background-color 0.2s, color 0.2s;
}

.preview-btn:hover {
  background: #fff;
  color: var(--accent-color);
}

.preview-btn.danger:hover {
  color: var(--danger-color);
}

.upload-progress {
  display: flex;
  align-items: center;
  gap: 8px;
  margin-top: 8px;
  width: 200px;
}

.progress-bar {
  flex: 1;
  height: 4px;
  background: var(--card-separator-color);
  border-radius: 2px;
  overflow: hidden;
}

.progress-fill {
  height: 100%;
  background: var(--accent-color);
  border-radius: 2px;
  animation: progress-indeterminate 1.4s ease infinite;
}

@keyframes progress-indeterminate {
  0% { width: 0; margin-left: 0; }
  50% { width: 60%; margin-left: 20%; }
  100% { width: 0; margin-left: 100%; }
}

.progress-text {
  font-size: 12px;
  color: var(--card-text-color-tertiary);
  white-space: nowrap;
}
</style>
