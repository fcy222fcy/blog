<template>
  <div>
    <!-- 统计卡片（加载态 + 真实态） -->
    <SkeletonLoader v-if="loading" type="stats" :count="2" />
    <div v-else class="stats-grid">
      <div class="stat-card">
        <div class="stat-label">分类总数</div>
        <div class="stat-value">{{ categories.length }}</div>
      </div>
      <div class="stat-card">
        <div class="stat-label">文章总数</div>
        <div class="stat-value">{{ totalArticles }}</div>
      </div>
    </div>

    <div style="display: flex; gap: 12px; margin-bottom: 20px; align-items: center; justify-content: space-between;">
      <div class="search-box">
        <span class="search-box-icon">⌕</span>
        <input type="text" v-model="keyword" placeholder="搜索分类...">
      </div>
      <button class="btn btn-primary" @click="showModal = true; resetForm()">
        <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><line x1="12" y1="5" x2="12" y2="19"></line><line x1="5" y1="12" x2="19" y2="12"></line></svg>
        <span>新建分类</span>
      </button>
    </div>

    <div class="category-cards">
      <!-- 骨架屏加载态 -->
      <SkeletonLoader v-if="loading" type="card-grid" :count="6" />
      <!-- 真实列表 -->
      <template v-else>
      <div v-for="cat in filteredCategories" :key="cat.id" class="category-card">
        <div class="category-card-cover" :style="{ background: getGradient(cat.id) }">
          <span class="category-card-cover-icon">{{ cat.icon || '📂' }}</span>
        </div>
        <div class="category-card-content">
          <div class="category-card-title">{{ cat.name }}</div>
          <div class="category-card-slug">{{ cat.slug }}</div>
          <div class="category-card-desc">{{ cat.description || '暂无描述' }}</div>
          <div class="category-card-footer">
            <span class="category-card-count">{{ cat.article_count || 0 }} 篇文章</span>
            <div class="category-card-actions">
              <button class="action-btn btn-edit btn-sm" @click="editCategory(cat)">编辑</button>
              <button class="action-btn btn-delete btn-sm" @click="handleDelete(cat.id)">删除</button>
            </div>
          </div>
        </div>
      </div>
      </template>
    </div>

    <div v-if="filteredCategories.length === 0 && !loading" class="card">
      <div class="card-body" style="text-align: center; color: var(--card-text-color-tertiary);">暂无分类</div>
    </div>

    <div class="modal-overlay" :class="{ active: showModal }" @click.self="showModal = false">
      <div class="modal">
        <div class="modal-header">
          <h3 class="modal-title">{{ editingId ? '编辑分类' : '新建分类' }}</h3>
          <button class="modal-close" @click="showModal = false">×</button>
        </div>
        <div class="modal-body">
          <div class="form-group">
            <label class="form-label">分类名称 <span class="required">*</span></label>
            <input type="text" class="form-input" v-model="form.name" placeholder="输入分类名称">
          </div>
          <div class="form-group">
            <label class="form-label">别名</label>
            <input type="text" class="form-input" v-model="form.slug" placeholder="英文别名，如 build">
          </div>
          <div class="form-group">
            <label class="form-label">描述</label>
            <textarea class="form-textarea" v-model="form.description" placeholder="输入分类描述..." rows="3"></textarea>
          </div>
          <div class="form-group">
            <label class="form-label">图标</label>
            <input type="text" class="form-input" v-model="form.icon" placeholder="输入 emoji 图标">
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
import { getCategoryList, createCategory, updateCategory, deleteCategory } from '../../api/category'
import SkeletonLoader from '../../components/common/SkeletonLoader.vue'

const categories = ref([])
const keyword = ref('')
const showModal = ref(false)
const editingId = ref(null)
const totalArticles = ref(0)
const loading = ref(true)
const form = ref({ name: '', slug: '', description: '', icon: '' })

const gradients = [
  'linear-gradient(135deg, #667eea, #764ba2)',
  'linear-gradient(135deg, #f093fb, #f5576c)',
  'linear-gradient(135deg, #43e97b, #38f9d7)',
  'linear-gradient(135deg, #4facfe, #00f2fe)',
  'linear-gradient(135deg, #fa709a, #fee140)',
]
const getGradient = (id) => gradients[(id - 1) % gradients.length]

const filteredCategories = computed(() => {
  if (!keyword.value) return categories.value
  return categories.value.filter(c => c.name.includes(keyword.value))
})

const resetForm = () => { editingId.value = null; form.value = { name: '', slug: '', description: '', icon: '' } }

const editCategory = (cat) => {
  editingId.value = cat.id
  form.value = { name: cat.name, slug: cat.slug, description: cat.description || '', icon: cat.icon || '' }
  showModal.value = true
}

const loadCategories = async () => {
  try {
    const res = await getCategoryList()
    categories.value = res.data || []
    totalArticles.value = categories.value.reduce((sum, c) => sum + (c.article_count || 0), 0)
  } catch (e) { console.error(e) }
  loading.value = false
}

const handleSave = async () => {
  if (!form.value.name || form.value.name.trim().length < 2) {
    ElMessage.warning('分类名称至少需要 2 个字符')
    return
  }
  if (form.value.name.length > 50) {
    ElMessage.warning('分类名称不能超过 50 个字符')
    return
  }
  try {
    if (editingId.value) { await updateCategory(editingId.value, form.value) }
    else { await createCategory(form.value) }
    ElMessage.success('保存成功')
    showModal.value = false
    loadCategories()
  } catch (e) { console.error(e) }
}

const handleDelete = async (id) => {
  try {
    await ElMessageBox.confirm('确定要删除这个分类吗？', '确认删除', { confirmButtonText: '删除', cancelButtonText: '取消', type: 'warning' })
    await deleteCategory(id)
    ElMessage.success('删除成功')
    loadCategories()
  } catch (e) { if (e !== 'cancel') console.error(e) }
}

onMounted(loadCategories)
</script>

<style scoped>
.category-card-slug { font-size: 12px; color: var(--accent-color); margin-bottom: 4px; font-family: monospace; }
.category-card-actions { display: flex; gap: 8px; }
</style>
