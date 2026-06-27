<template>
  <div>
    <div class="card">
      <div class="card-header">
        <div class="card-title">媒体库</div>
        <div style="display: flex; gap: 12px; flex-wrap: wrap; align-items: center;">
          <div class="search-box">
            <span class="search-box-icon">⌕</span>
            <input type="text" v-model="keyword" placeholder="搜索文件...">
          </div>
          <button class="btn btn-primary" @click="$refs.fileInput.click()">
            <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><line x1="12" y1="5" x2="12" y2="19"></line><line x1="5" y1="12" x2="19" y2="12"></line></svg>
            <span>上传文件</span>
          </button>
        </div>
      </div>
      <div class="card-body">
        <div class="upload-area" @dragover.prevent @drop.prevent="handleDrop" :class="{ dragging: isDragging }">
          <input type="file" ref="fileInput" @change="handleFileSelect" multiple accept="image/*,.pdf" style="display: none;" />
          <div class="upload-icon">
            <svg xmlns="http://www.w3.org/2000/svg" width="40" height="40" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M22 19a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h5l2 3h9a2 2 0 0 1 2 2z"></path></svg>
          </div>
          <div class="upload-text" @click="$refs.fileInput.click()">
            <strong>点击上传</strong> 或拖拽文件到此处<br>
            支持 JPG、PNG、GIF、WebP、PDF，单文件最大 5MB
          </div>
        </div>

        <div style="display: grid; grid-template-columns: repeat(auto-fill, minmax(180px, 1fr)); gap: 16px; margin-top: 20px;">
          <div v-for="item in filteredMedia" :key="item.id" class="media-grid-item">
            <div class="media-grid-preview">
              <img v-if="isImage(item.type)" :src="item.url" :alt="item.name" />
              <div v-else class="media-file-icon">📄</div>
            </div>
            <div class="media-grid-info">
              <div class="media-grid-name" :title="item.name">{{ item.name }}</div>
              <div class="media-grid-size">{{ formatSize(item.size) }}</div>
            </div>
            <div class="media-grid-actions">
              <button class="action-btn btn-edit btn-sm" @click="copyUrl(item.url)">复制链接</button>
              <button class="action-btn btn-delete btn-sm" @click="handleDelete(item.id)">删除</button>
            </div>
          </div>
        </div>

        <div v-if="filteredMedia.length === 0" style="text-align: center; padding: 40px; color: var(--card-text-color-tertiary);">
          暂无媒体文件
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'

const mediaList = ref([])
const keyword = ref('')
const isDragging = ref(false)
const fileInput = ref(null)

const mockMedia = [
  { id: 1, name: 'avatar.jpg', url: 'https://via.placeholder.com/150', type: 'image/jpeg', size: 102400 },
  { id: 2, name: 'banner.png', url: 'https://via.placeholder.com/300x100', type: 'image/png', size: 204800 },
  { id: 3, name: 'document.pdf', url: '#', type: 'application/pdf', size: 1024000 },
]

onMounted(() => { mediaList.value = mockMedia })

const filteredMedia = computed(() => {
  if (!keyword.value) return mediaList.value
  return mediaList.value.filter(m => m.name.includes(keyword.value))
})

const isImage = (type) => type && type.startsWith('image/')

const formatSize = (bytes) => {
  if (bytes === 0) return '0 B'
  const k = 1024
  const sizes = ['B', 'KB', 'MB', 'GB']
  const i = Math.floor(Math.log(bytes) / Math.log(k))
  return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i]
}

const copyUrl = (url) => { navigator.clipboard.writeText(url); ElMessage.success('链接已复制') }

const handleDelete = async (id) => {
  try {
    await ElMessageBox.confirm('确定要删除这个文件吗？', '确认删除', { confirmButtonText: '删除', cancelButtonText: '取消', type: 'warning' })
    mediaList.value = mediaList.value.filter(item => item.id !== id)
    ElMessage.success('删除成功')
  } catch (e) { if (e !== 'cancel') console.error(e) }
}

const handleFileSelect = (event) => {
  const files = Array.from(event.target.files)
  files.forEach(file => {
    mediaList.value.unshift({ id: Date.now(), name: file.name, url: URL.createObjectURL(file), type: file.type, size: file.size })
  })
  ElMessage.success('上传成功')
}

const handleDrop = (event) => {
  isDragging.value = false
  const files = Array.from(event.dataTransfer.files)
  files.forEach(file => {
    mediaList.value.unshift({ id: Date.now(), name: file.name, url: URL.createObjectURL(file), type: file.type, size: file.size })
  })
  ElMessage.success('上传成功')
}
</script>

<style scoped>
.dragging { border-color: var(--accent-color) !important; background: rgba(var(--accent-color-rgb), 0.05) !important; }
.media-grid-item {
  background: var(--card-background);
  border: 1px solid var(--card-border);
  border-radius: var(--card-border-radius);
  overflow: hidden;
  transition: transform 0.2s, box-shadow 0.2s;
}
.media-grid-item:hover { transform: translateY(-2px); box-shadow: var(--shadow-l2); }
.media-grid-preview { height: 140px; display: flex; align-items: center; justify-content: center; background: var(--body-background); }
.media-grid-preview img { width: 100%; height: 100%; object-fit: cover; }
.media-file-icon { font-size: 36px; }
.media-grid-info { padding: 12px; }
.media-grid-name { font-size: 13px; font-weight: 500; white-space: nowrap; overflow: hidden; text-overflow: ellipsis; margin-bottom: 2px; color: var(--card-text-color-main); }
.media-grid-size { font-size: 12px; color: var(--card-text-color-tertiary); }
.media-grid-actions { display: flex; gap: 6px; padding: 0 12px 12px; }
</style>
