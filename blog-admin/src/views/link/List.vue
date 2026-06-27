<template>
  <div>
    <div class="stats-grid">
      <div class="stat-card">
        <div class="stat-label">友链总数</div>
        <div class="stat-value">{{ links.length }}</div>
      </div>
    </div>

    <div style="display: flex; gap: 12px; margin-bottom: 20px; align-items: center; justify-content: space-between;">
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
      <div v-for="link in filteredLinks" :key="link.id" class="link-card">
        <div class="link-card-header">
          <div class="link-card-avatar" :style="{ background: getGradient(link.id) }">
            {{ link.avatar || '🔗' }}
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
    </div>

    <div v-if="filteredLinks.length === 0" class="card">
      <div class="card-body" style="text-align: center; color: var(--card-text-color-tertiary);">暂无友链</div>
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
          <div class="form-group" style="margin-bottom: 0;">
            <label class="form-label">头像 (emoji)</label>
            <input type="text" class="form-input" v-model="form.avatar" placeholder="输入 emoji 如 🤖">
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

const links = ref([])
const keyword = ref('')
const showModal = ref(false)
const editingId = ref(null)
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
}

const handleSave = async () => {
  if (!form.value.name || !form.value.url) { ElMessage.warning('请填写名称和链接'); return }
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
