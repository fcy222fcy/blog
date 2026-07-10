<template>
  <div>
    <!-- 统计卡片 -->
    <SkeletonLoader v-if="loading" type="stats" :count="2" />
    <div v-else class="stats-grid">
      <div class="stat-card">
        <div class="stat-label">条目总数</div>
        <div class="stat-value">{{ totalCount }}</div>
      </div>
      <div class="stat-card">
        <div class="stat-label">已完成</div>
        <div class="stat-value text-success">{{ completedCount }}</div>
      </div>
    </div>

    <!-- 类型筛选 Tab -->
    <div class="tab-group">
      <button class="btn" :class="activeType === 'all' ? 'btn-primary' : 'btn-secondary'" @click="activeType = 'all'; loadList()">全部</button>
      <button class="btn" :class="activeType === 'movie' ? 'btn-primary' : 'btn-secondary'" @click="activeType = 'movie'; loadList()">🎬 电影</button>
      <button class="btn" :class="activeType === 'tv' ? 'btn-primary' : 'btn-secondary'" @click="activeType = 'tv'; loadList()">📺 剧集</button>
      <button class="btn" :class="activeType === 'game' ? 'btn-primary' : 'btn-secondary'" @click="activeType = 'game'; loadList()">🎮 游戏</button>
    </div>

    <div class="page-toolbar">
      <div class="search-box">
        <span class="search-box-icon">⌕</span>
        <input type="text" v-model="keyword" placeholder="搜索标题/评论..." @keyup.enter="loadList()">
      </div>
      <button class="btn btn-primary" @click="openModal()">
        <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><line x1="12" y1="5" x2="12" y2="19"></line><line x1="5" y1="12" x2="19" y2="12"></line></svg>
        <span>添加条目</span>
      </button>
    </div>

    <!-- 列表区域 -->
    <div class="media-cards">
      <SkeletonLoader v-if="loading" type="card-grid" :count="6" />
      <template v-else>
        <div v-for="item in list" :key="item.id" class="media-card">
          <div class="media-card-cover" :style="getCoverStyle(item)">
            <span v-if="!item.cover" class="media-card-cover-icon">{{ getTypeIcon(item.type) }}</span>
          </div>
          <div class="media-card-info">
            <div class="media-card-title">{{ item.title }}</div>
            <div class="media-card-subtitle">{{ item.title_en || item.platform || getTypeLabel(item.type) }}</div>
            <div class="media-card-meta">
              <span class="media-card-tag badge-accent">{{ getTypeLabel(item.type) }}</span>
              <span v-if="item.rating > 0" class="media-card-rating">⭐ {{ item.rating }}</span>
              <span class="media-card-year">{{ item.year }}</span>
            </div>
            <div v-if="item.comment" class="media-card-comment">{{ item.comment }}</div>
            <div class="media-card-status" :class="getStatusClass(item.status)">{{ getStatusLabel(item.status) }}</div>
          </div>
          <div class="media-card-actions">
            <button class="action-btn btn-edit btn-sm" @click="openModal(item)">编辑</button>
            <button class="action-btn btn-delete btn-sm" @click="handleDelete(item.id)">删除</button>
          </div>
        </div>
      </template>
    </div>

    <div v-if="list.length === 0 && !loading" class="card">
      <div class="card-body empty-state-sm">暂无条目</div>
    </div>

    <!-- 编辑弹窗 -->
    <div class="modal-overlay" :class="{ active: showModal }" @click.self="showModal = false">
      <div class="modal entertainment-modal">
        <div class="modal-header">
          <h3 class="modal-title">{{ editingId ? '编辑条目' : '添加条目' }}</h3>
          <button class="modal-close" @click="showModal = false">×</button>
        </div>
        <div class="modal-body">
          <div class="form-row-2">
            <div class="form-group">
              <label class="form-label">标题 <span class="required">*</span></label>
              <input type="text" class="form-input" v-model="form.title" placeholder="中文标题">
            </div>
            <div class="form-group">
              <label class="form-label">英文标题</label>
              <input type="text" class="form-input" v-model="form.title_en" placeholder="English title">
            </div>
          </div>
          <div class="form-row-3">
            <div class="form-group">
              <label class="form-label">类型 <span class="required">*</span></label>
              <select class="form-input" v-model="form.type">
                <option value="movie">🎬 电影</option>
                <option value="tv">📺 剧集</option>
                <option value="game">🎮 游戏</option>
              </select>
            </div>
            <div class="form-group">
              <label class="form-label">年份 <span class="required">*</span></label>
              <input type="number" class="form-input" v-model.number="form.year" placeholder="2026" min="1900" max="2100">
            </div>
            <div class="form-group">
              <label class="form-label">状态</label>
              <select class="form-input" v-model="form.status">
                <option value="watching">在看/在玩</option>
                <option value="watched">已看完</option>
                <option value="playing">在玩</option>
                <option value="completed">已完成/已通关</option>
              </select>
            </div>
          </div>
          <div class="form-row-3">
            <div class="form-group">
              <label class="form-label">我的评分</label>
              <input type="number" step="0.1" class="form-input" v-model.number="form.rating" placeholder="0-10" min="0" max="10">
            </div>
            <div class="form-group">
              <label class="form-label">外部评分</label>
              <input type="number" step="0.1" class="form-input" v-model.number="form.rating_external" placeholder="豆瓣/IMDB" min="0" max="10">
            </div>
            <div class="form-group">
              <label class="form-label">平台</label>
              <input type="text" class="form-input" v-model="form.platform" placeholder="PC/PS5/Netflix...">
            </div>
          </div>
          <div class="form-row-2">
            <div class="form-group">
              <label class="form-label">游戏时长</label>
              <input type="text" class="form-input" v-model="form.playtime" placeholder="35H / 2季">
            </div>
            <div class="form-group">
              <label class="form-label">外链</label>
              <input type="text" class="form-input" v-model="form.link" placeholder="TMDB/Steam链接">
            </div>
          </div>
          <div class="form-group">
            <label class="form-label">封面图</label>
            <input type="text" class="form-input" v-model="form.cover" placeholder="https://.../cover.jpg">
          </div>
          <div class="form-group">
            <label class="form-label">评论/观后感</label>
            <textarea class="form-textarea" v-model="form.comment" placeholder="写点观后感..." rows="3"></textarea>
          </div>
        </div>
        <div class="modal-footer">
          <button class="btn btn-secondary" @click="showModal = false">取消</button>
          <button class="btn btn-primary" @click="handleSave" :disabled="saving">{{ saving ? '保存中...' : '保存' }}</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { getEntertainmentList, createEntertainment, updateEntertainment, deleteEntertainment } from '../../api/entertainment'
import SkeletonLoader from '../../components/common/SkeletonLoader.vue'

const loading = ref(true)
const saving = ref(false)
const showModal = ref(false)
const editingId = ref(null)
const list = ref([])
const totalCount = ref(0)
const activeType = ref('all')
const keyword = ref('')

const form = ref({
  title: '',
  title_en: '',
  type: 'movie',
  year: new Date().getFullYear(),
  cover: '',
  rating: 0,
  rating_external: 0,
  platform: '',
  playtime: '',
  comment: '',
  status: 'watching',
  link: '',
})

const completedCount = computed(() => {
  return list.value.filter(i => i.status === 'watched' || i.status === 'completed').length
})

const gradients = [
  'linear-gradient(135deg, #667eea, #764ba2)',
  'linear-gradient(135deg, #f093fb, #f5576c)',
  'linear-gradient(135deg, #43e97b, #38f9d7)',
  'linear-gradient(135deg, #4facfe, #00f2fe)',
  'linear-gradient(135deg, #fa709a, #fee140)',
]

const getGradient = (id) => gradients[(id - 1) % gradients.length]

const getTypeIcon = (type) => ({ movie: '🎬', tv: '📺', game: '🎮' }[type] || '🎬')
const getTypeLabel = (type) => ({ movie: '电影', tv: '剧集', game: '游戏' }[type] || type)

const getStatusLabel = (s) => ({
  watching: '在看', watched: '已看完', playing: '在玩', completed: '已完成'
}[s] || s || '未开始')

const getStatusClass = (s) => {
  if (s === 'watched' || s === 'completed') return 'status-published'
  return 'status-draft'
}

const getCoverStyle = (item) => {
  if (item.cover) return { backgroundImage: `url(${item.cover})`, backgroundSize: 'cover', backgroundPosition: 'center' }
  return { background: getGradient(item.id) }
}

const resetForm = () => {
  editingId.value = null
  form.value = {
    title: '', title_en: '', type: 'movie', year: new Date().getFullYear(),
    cover: '', rating: 0, rating_external: 0, platform: '',
    playtime: '', comment: '', status: 'watching', link: '',
  }
}

const openModal = (item = null) => {
  if (item) {
    editingId.value = item.id
    form.value = {
      title: item.title,
      title_en: item.title_en || '',
      type: item.type,
      year: item.year,
      cover: item.cover || '',
      rating: item.rating || 0,
      rating_external: item.rating_external || 0,
      platform: item.platform || '',
      playtime: item.playtime || '',
      comment: item.comment || '',
      status: item.status || 'watching',
      link: item.link || '',
    }
  } else {
    resetForm()
  }
  showModal.value = true
}

const loadList = async () => {
  loading.value = true
  try {
    const params = { page: 1, page_size: 100 }
    if (activeType.value !== 'all') params.type = activeType.value
    if (keyword.value) params.keyword = keyword.value
    const res = await getEntertainmentList(params)
    list.value = res.data?.list || []
    totalCount.value = res.data?.total || 0
  } catch (e) { console.error(e) }
  loading.value = false
}

const handleSave = async () => {
  if (!form.value.title || form.value.title.trim().length < 1) {
    ElMessage.warning('请输入标题')
    return
  }
  if (!form.value.year || form.value.year < 1900 || form.value.year > 2100) {
    ElMessage.warning('请输入正确的年份')
    return
  }
  saving.value = true
  try {
    const payload = { ...form.value }
    if (!payload.rating) payload.rating = 0
    if (!payload.rating_external) payload.rating_external = 0
    if (editingId.value) {
      await updateEntertainment(editingId.value, payload)
    } else {
      await createEntertainment(payload)
    }
    ElMessage.success('保存成功')
    showModal.value = false
    loadList()
  } catch (e) { console.error(e) }
  saving.value = false
}

const handleDelete = async (id) => {
  try {
    await ElMessageBox.confirm('确定要删除这个条目吗？', '确认删除', { confirmButtonText: '删除', cancelButtonText: '取消', type: 'warning' })
    await deleteEntertainment(id)
    ElMessage.success('删除成功')
    loadList()
  } catch (e) { if (e !== 'cancel') console.error(e) }
}

onMounted(loadList)
</script>

<style scoped>
.media-card-meta { display: flex; align-items: center; gap: 8px; margin-top: 6px; flex-wrap: wrap; }
.media-card-tag { padding: 2px 8px; border-radius: 4px; font-size: 12px; }
.media-card-rating { font-size: 13px; color: var(--card-text-color-secondary); }
.media-card-year { font-size: 13px; color: var(--card-text-color-tertiary); }
.media-card-status { margin-top: 6px; font-size: 12px; font-weight: 500; }
.media-card-comment { margin-top: 6px; font-size: 13px; color: var(--card-text-color-secondary); display: -webkit-box; -webkit-line-clamp: 2; -webkit-box-orient: vertical; overflow: hidden; }

.media-card-cover-icon { font-size: 32px; }

.entertainment-modal {
  width: 640px;
  max-width: 94%;
  max-height: 90vh;
  display: flex;
  flex-direction: column;
}
.entertainment-modal .modal-body {
  overflow-y: auto;
  padding-right: 4px;
}

.form-row-2 { display: grid; grid-template-columns: 1fr 1fr; gap: 12px; }
.form-row-3 { display: grid; grid-template-columns: 1fr 1fr 1fr; gap: 12px; }

@media (max-width: 640px) {
  .form-row-2, .form-row-3 { grid-template-columns: 1fr; }
}
</style>
