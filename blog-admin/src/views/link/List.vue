<template>
  <div>
    <!-- 统计卡片（加载态 + 真实态） -->
    <SkeletonLoader v-if="loading" type="stats" :count="1" />
    <div v-else class="stats-grid">
      <div class="stat-card">
        <div class="stat-label">友链总数</div>
        <div class="stat-value">{{ links.length }}</div>
      </div>
    </div>

    <div class="page-toolbar">
      <div class="search-box">
        <span class="search-box-icon">⌕</span>
        <input type="text" v-model="keyword" placeholder="搜索友链...">
      </div>
      <button class="btn btn-primary" @click="showModal = true; resetForm()">
        <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><line x1="12" y1="5" x2="12" y2="19"></line><line x1="5" y1="12" x2="19" y2="12"></line></svg>
        <span>添加链接</span>
      </button>
    </div>

    <div class="link-cards">
      <!-- 骨架屏加载态 -->
      <SkeletonLoader v-if="loading" type="media" :count="4" />
      <!-- 真实列表 -->
      <template v-else>
      <div v-for="link in filteredLinks" :key="link.id" class="link-card">
        <div class="link-card-header">
          <div class="link-card-avatar" :style="{ background: getGradient(link.id) }">
            <img v-if="link.avatar && (link.avatar.startsWith('http://') || link.avatar.startsWith('https://') || link.avatar.startsWith('/'))" :src="link.avatar" alt="avatar" style="width:100%;height:100%;object-fit:cover;border-radius:inherit;">
            <span v-else>{{ link.avatar || '🔗' }}</span>
          </div>
          <div class="link-card-info">
            <div class="link-card-name">{{ link.name }}</div>
            <div class="link-card-url">{{ link.url }}</div>
          </div>
        </div>
        <div class="link-card-body">
          <p class="link-card-desc">{{ link.description || '暂无描述' }}</p>
        </div>
        <div class="link-card-footer">
          <button class="action-btn btn-edit btn-sm" @click="editLink(link)">编辑</button>
          <button class="action-btn btn-delete btn-sm" @click="handleDelete(link.id)">删除</button>
        </div>
      </div>
      </template>
    </div>

    <div v-if="filteredLinks.length === 0 && !loading" class="card">
      <div class="card-body empty-state-sm">暂无友链</div>
    </div>

    <div class="modal-overlay" :class="{ active: showModal }" @click.self="showModal = false">
      <div class="modal">
        <div class="modal-header">
          <h3 class="modal-title">{{ editingId ? '编辑友链' : '添加友链' }}</h3>
          <button class="modal-close" @click="showModal = false">×</button>
        </div>
        <div class="modal-body">
          <div class="form-group">
            <label class="form-label">名称 <span class="required">*</span></label>
            <input type="text" class="form-input" v-model="form.name" placeholder="输入名称">
          </div>
          <div class="form-group">
            <label class="form-label">链接 <span class="required">*</span></label>
            <input type="text" class="form-input" v-model="form.url" placeholder="https://example.com">
          </div>
          <div class="form-group">
            <label class="form-label">描述</label>
            <textarea class="form-textarea" v-model="form.description" placeholder="输入描述..." rows="2"></textarea>
          </div>
          <div class="form-group mb-0">
            <label class="form-label">头像</label>
            <input type="text" class="form-input" v-model="form.avatar" placeholder="输入 emoji 如 🤖 或图片 URL">
          </div>
        </div>
        <div class="modal-footer">
          <button class="btn btn-secondary" @click="showModal = false">取消</button>
          <button class="btn btn-primary" @click="handleSave">保存</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { getLinkList, createLink, updateLink, deleteLink } from '../../api/link'
import SkeletonLoader from '../../components/common/SkeletonLoader.vue'

const links = ref([])
const keyword = ref('')
const showModal = ref(false)
const editingId = ref(null)
const loading = ref(true)
const form = ref({ name: '', url: '', description: '', avatar: '' })

const gradients = [
  'linear-gradient(135deg, #667eea, #764ba2)',
  'linear-gradient(135deg, #f093fb, #f5576c)',
  'linear-gradient(135deg, #43e97b, #38f9d7)',
  'linear-gradient(135deg, #4facfe, #00f2fe)',
]
const getGradient = (id) => gradients[(id - 1) % gradients.length]

const filteredLinks = computed(() => {
  if (!keyword.value) return links.value
  return links.value.filter(l => l.name.includes(keyword.value))
})

const resetForm = () => { editingId.value = null; form.value = { name: '', url: '', description: '', avatar: '' } }

const editLink = (link) => {
  editingId.value = link.id
  form.value = { name: link.name, url: link.url, description: link.description || '', avatar: link.avatar || '' }
  showModal.value = true
}

const loadLinks = async () => {
  try {
    const res = await getLinkList({ page: 1, page_size: 100 })
    links.value = res.data?.list || []
  } catch (e) { console.error(e) }
  loading.value = false
}

const isValidUrl = (url) => {
  try {
    new URL(url)
    return true
  } catch {
    return false
  }
}

const handleSave = async () => {
  if (!form.value.name || form.value.name.trim().length < 2) {
    ElMessage.warning('名称至少需要 2 个字符')
    return
  }
  if (!form.value.url) {
    ElMessage.warning('请输入链接地址')
    return
  }
  if (!isValidUrl(form.value.url)) {
    ElMessage.warning('请输入有效的链接地址（如 https://example.com）')
    return
  }
  try {
    if (editingId.value) { await updateLink(editingId.value, form.value) }
    else { await createLink(form.value) }
    ElMessage.success('保存成功')
    showModal.value = false
    loadLinks()
  } catch (e) { console.error(e) }
}

const handleDelete = async (id) => {
  try {
    await ElMessageBox.confirm('确定要删除这个友链吗？', '确认删除', { confirmButtonText: '删除', cancelButtonText: '取消', type: 'warning' })
    await deleteLink(id)
    ElMessage.success('删除成功')
    loadLinks()
  } catch (e) { if (e !== 'cancel') console.error(e) }
}

onMounted(loadLinks)
</script>
