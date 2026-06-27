<template>
  <div>
    <div class="stats-grid">
      <div class="stat-card">
        <div class="stat-label">标签总数</div>
        <div class="stat-value">{{ tags.length }}</div>
      </div>
      <div class="stat-card">
        <div class="stat-label">文章总数</div>
        <div class="stat-value">{{ totalArticles }}</div>
      </div>
    </div>

    <div style="display: flex; gap: 12px; margin-bottom: 20px; align-items: center; justify-content: space-between;">
      <div class="search-box">
        <span class="search-box-icon">⌕</span>
        <input type="text" v-model="keyword" placeholder="搜索标签...">
      </div>
      <button class="btn btn-primary" @click="showModal = true; resetForm()">
        <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><line x1="12" y1="5" x2="12" y2="19"></line><line x1="5" y1="12" x2="19" y2="12"></line></svg>
        <span>新建标签</span>
      </button>
    </div>

    <div class="tags-grid">
      <div v-for="tag in filteredTags" :key="tag.id" class="tag-card">
        <div class="tag-card-name">{{ tag.name }}</div>
        <div class="tag-card-count">{{ tag.article_count || 0 }} 篇文章</div>
        <div class="tag-card-actions">
          <button class="action-btn btn-edit btn-sm" @click="editTag(tag)">编辑</button>
          <button class="action-btn btn-delete btn-sm" @click="handleDelete(tag.id)">删除</button>
        </div>
      </div>
    </div>

    <div v-if="filteredTags.length === 0" class="card">
      <div class="card-body" style="text-align: center; color: var(--card-text-color-tertiary);">暂无标签</div>
    </div>

    <div class="modal-overlay" :class="{ active: showModal }" @click.self="showModal = false">
      <div class="modal" style="max-width: 400px;">
        <div class="modal-header">
          <h3 class="modal-title">{{ editingId ? '编辑标签' : '新建标签' }}</h3>
          <button class="modal-close" @click="showModal = false">×</button>
        </div>
        <div class="modal-body">
          <div class="form-group">
            <label class="form-label">标签名称 <span class="required">*</span></label>
            <input type="text" class="form-input" v-model="form.name" placeholder="输入标签名称">
          </div>
          <div class="form-group" style="margin-bottom: 0;">
            <label class="form-label">别名</label>
            <input type="text" class="form-input" v-model="form.slug" placeholder="英文别名，如 hugo">
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
import { getTagList, createTag, updateTag, deleteTag } from '../../api/tag'

const tags = ref([])
const keyword = ref('')
const showModal = ref(false)
const editingId = ref(null)
const totalArticles = ref(0)
const form = ref({ name: '', slug: '' })

const filteredTags = computed(() => {
  if (!keyword.value) return tags.value
  return tags.value.filter(t => t.name.includes(keyword.value))
})

const resetForm = () => { editingId.value = null; form.value = { name: '', slug: '' } }

const editTag = (tag) => {
  editingId.value = tag.id
  form.value = { name: tag.name, slug: tag.slug }
  showModal.value = true
}

const loadTags = async () => {
  try {
    const res = await getTagList()
    tags.value = res.data || []
    totalArticles.value = tags.value.reduce((sum, t) => sum + (t.article_count || 0), 0)
  } catch (e) { console.error(e) }
}

const handleSave = async () => {
  if (!form.value.name) { ElMessage.warning('请输入标签名称'); return }
  try {
    if (editingId.value) { await updateTag(editingId.value, form.value) }
    else { await createTag(form.value) }
    ElMessage.success('保存成功')
    showModal.value = false
    loadTags()
  } catch (e) { console.error(e) }
}

const handleDelete = async (id) => {
  try {
    await ElMessageBox.confirm('确定要删除这个标签吗？', '确认删除', { confirmButtonText: '删除', cancelButtonText: '取消', type: 'warning' })
    await deleteTag(id)
    ElMessage.success('删除成功')
    loadTags()
  } catch (e) { if (e !== 'cancel') console.error(e) }
}

onMounted(loadTags)
</script>
