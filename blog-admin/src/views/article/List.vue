<template>
  <div>
    <!-- 统计卡片（加载态 + 真实态） -->
    <SkeletonLoader v-if="loading" type="stats" :count="4" />
    <div v-else class="stats-grid">
      <div class="stat-card">
        <div class="stat-label">文章总数</div>
        <div class="stat-value">{{ allTotal }}</div>
      </div>
      <div class="stat-card">
        <div class="stat-label">已发布</div>
        <div class="stat-value">{{ stats.published || 0 }}</div>
      </div>
      <div class="stat-card">
        <div class="stat-label">草稿</div>
        <div class="stat-value">{{ stats.draft || 0 }}</div>
      </div>
      <div class="stat-card">
        <div class="stat-label">总浏览量</div>
        <div class="stat-value">{{ stats.totalViews || 0 }}</div>
      </div>
    </div>

    <div class="card">
      <div class="card-header">
        <div class="card-title">文章列表</div>
        <div class="filter-group">
          <div class="search-box">
            <span class="search-box-icon">⌕</span>
            <input type="text" v-model="keyword" placeholder="搜索文章..." @input="loadArticles">
          </div>
          <CustomSelect v-model="categoryFilter" :options="categoryOptions" @change="loadArticles" />
          <button class="btn btn-primary" @click="$router.push('/articles/edit')">
            <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><line x1="12" y1="5" x2="12" y2="19"></line><line x1="5" y1="12" x2="19" y2="12"></line></svg>
            <span>新建文章</span>
          </button>
        </div>
      </div>

      <!-- 状态标签页 -->
      <div class="article-tabs">
        <button
          v-for="tab in statusTabs"
          :key="tab.value"
          class="article-tab"
          :class="{ active: statusFilter === tab.value }"
          @click="switchTab(tab.value)"
        >
          {{ tab.label }}
          <span v-if="tab.count !== undefined" class="article-tab-count">{{ tab.count }}</span>
        </button>
      </div>

      <div class="article-cards-container">
        <!-- 骨架屏加载态 -->
        <SkeletonLoader v-if="loading" type="article" :count="5" />
        <!-- 真实列表 -->
        <template v-else>
        <div v-for="article in articles" :key="article.id" class="article-card-item" @click="$router.push('/articles/edit/' + article.id)">
          <!-- 封面缩略图 -->
          <div class="article-card-cover" :class="{ 'no-cover': !article.cover }">
            <img v-if="article.cover" :src="article.cover" alt="封面" class="article-cover-img">
            <div v-else class="article-cover-placeholder" aria-label="暂无封面">
              <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round" role="img"><title>暂无封面</title><path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"></path><polyline points="14 2 14 8 20 8"></polyline></svg>
            </div>
          </div>
          <!-- 文章信息 -->
          <div class="article-card-info">
            <div class="article-card-header">
              <h3 class="article-card-title">{{ article.title }}</h3>
              <span class="status-badge" :class="'status-' + article.status">
                {{ statusLabel(article.status) }}
              </span>
            </div>
            <p class="article-card-summary">{{ article.summary || '暂无摘要' }}</p>
            <div class="article-card-footer">
              <div class="article-card-meta">
                <span>{{ article.category?.name || '未分类' }}</span>
                <span>·</span>
                <span>{{ formatDate(article.created_at) }}</span>
                <span>·</span>
                <span>{{ article.view_count || 0 }} 次浏览</span>
              </div>
              <div class="article-card-actions" @click.stop>
                <button class="action-btn btn-edit btn-sm" @click="$router.push('/articles/edit/' + article.id)">编辑</button>
                <button class="action-btn btn-delete btn-sm" @click="handleDelete(article.id)">删除</button>
              </div>
            </div>
          </div>
        </div>
        <div v-if="articles.length === 0 && !loading" style="text-align: center; padding: 40px; color: var(--card-text-color-tertiary);">
          暂无文章
        </div>
        </template>
      </div>

      <div class="card-body">
        <div class="pagination">
          <div class="pagination-info">
            共 <span>{{ total }}</span> 篇文章
          </div>
          <div class="pagination-buttons">
            <button class="pagination-btn" :disabled="page <= 1" @click="page--; loadArticles()">上一页</button>
            <button class="pagination-btn active">{{ page }}</button>
            <button class="pagination-btn" :disabled="page * 10 >= total" @click="page++; loadArticles()">下一页</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { getArticleList, deleteArticle } from '../../api/article'
import { getCategoryList } from '../../api/category'
import { ElMessage, ElMessageBox } from 'element-plus'
import SkeletonLoader from '../../components/common/SkeletonLoader.vue'
import CustomSelect from '../../components/common/CustomSelect.vue'

const articles = ref([])
const categories = ref([])
const loading = ref(false)
const page = ref(1)
const total = ref(0)
const allTotal = ref(0)
const keyword = ref('')
const categoryFilter = ref('')
const statusFilter = ref('')
const stats = ref({ published: 0, draft: 0, totalViews: 0 })

const categoryOptions = computed(() => {
  return [{ value: '', label: '全部分类' }, ...categories.value.map(c => ({ value: c.id, label: c.name }))]
})

const statusTabs = computed(() => [
  { label: '全部', value: '', count: total.value },
  { label: '已发布', value: 'published', count: stats.value.published },
  { label: '草稿', value: 'draft', count: stats.value.draft },
  { label: '定时发布', value: 'scheduled', count: stats.value.scheduled || 0 }
])

const formatDate = (d) => d ? d.split('T')[0] : ''
const statusLabel = (s) => ({ published: '已发布', draft: '草稿', scheduled: '定时发布' }[s] || s)

const switchTab = (status) => {
  statusFilter.value = status
  page.value = 1
  loadArticles()
}

const loadStats = async () => {
  try {
    const [allRes, publishedRes, draftRes, scheduledRes] = await Promise.all([
      getArticleList({ page: 1, page_size: 1 }),
      getArticleList({ page: 1, page_size: 1, status: 'published' }),
      getArticleList({ page: 1, page_size: 1, status: 'draft' }),
      getArticleList({ page: 1, page_size: 1, status: 'scheduled' })
    ])
    allTotal.value = allRes.data?.total || 0
    stats.value.published = publishedRes.data?.total || 0
    stats.value.draft = draftRes.data?.total || 0
    stats.value.scheduled = scheduledRes.data?.total || 0
  } catch (e) { console.error(e) }
}

const loadArticles = async () => {
  try {
    const params = { page: page.value, page_size: 10 }
    if (keyword.value) params.keyword = keyword.value
    if (categoryFilter.value) params.category_id = categoryFilter.value
    if (statusFilter.value) params.status = statusFilter.value
    const res = await getArticleList(params)
    articles.value = res.data?.list || []
    total.value = res.data?.total || 0
    stats.value.totalViews = articles.value.reduce((sum, a) => sum + (a.view_count || 0), 0)
  } catch (e) { console.error(e) }
}

const loadCategories = async () => {
  try {
    const res = await getCategoryList()
    categories.value = res.data || []
  } catch (e) { console.error(e) }
}

const handleDelete = async (id) => {
  try {
    await ElMessageBox.confirm('确定要删除这篇文章吗？此操作无法撤销。', '确认删除', {
      confirmButtonText: '删除',
      cancelButtonText: '取消',
      type: 'warning'
    })
    await deleteArticle(id)
    ElMessage.success('删除成功')
    loadArticles()
  } catch (e) {
    if (e !== 'cancel') console.error(e)
  }
}

onMounted(async () => {
  loading.value = true
  loadStats()
  loadCategories()
  await loadArticles()
  loading.value = false
})
</script>

<style scoped>
.filter-group {
  display: flex;
  gap: 12px;
  flex-wrap: wrap;
  align-items: center;
}

.article-tabs {
  display: flex;
  gap: 0;
  border-bottom: 1px solid var(--card-separator-color);
  padding: 0 20px;
}

.article-tab {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  padding: 12px 20px;
  border: none;
  background: none;
  color: var(--card-text-color-secondary);
  font-size: 14px;
  font-weight: 500;
  cursor: pointer;
  border-bottom: 2px solid transparent;
  transition: all 0.2s ease;
  margin-bottom: -1px;
}

.article-tab:hover {
  color: var(--card-text-color-main);
}

.article-tab.active {
  color: var(--accent-color);
  border-bottom-color: var(--accent-color);
}

.article-tab-count {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  min-width: 20px;
  height: 20px;
  padding: 0 6px;
  border-radius: 10px;
  background: var(--input-background);
  color: var(--card-text-color-tertiary);
  font-size: 12px;
  font-weight: 600;
}

.article-tab.active .article-tab-count {
  background: rgba(var(--accent-color-rgb), 0.1);
  color: var(--accent-color);
}

.article-cards-container {
  padding: 0;
  width: 100%;
}
.article-card-item {
  display: flex;
  gap: 16px;
  padding: 20px 28px;
  border-bottom: 1px solid var(--card-separator-color);
  transition: background 0.15s ease;
  cursor: pointer;
  width: 100%;
}
.article-card-item:hover {
  background: rgba(var(--accent-color-rgb), 0.02);
}
.article-card-item:last-child {
  border-bottom: none;
}

/* 封面缩略图 */
.article-card-cover {
  width: 140px;
  height: 90px;
  border-radius: 8px;
  overflow: hidden;
  flex-shrink: 0;
  background: linear-gradient(135deg, #667eea20, #764ba220);
}
.article-card-cover.no-cover {
  background: rgba(var(--accent-color-rgb), 0.03);
  border: 1px dashed var(--card-separator-color);
}
.article-cover-img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}
.article-cover-placeholder {
  width: 100%;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  color: var(--card-text-color-tertiary);
  opacity: 0.5;
}

.article-card-info {
  flex: 1;
  min-width: 0;
}
.article-card-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 12px;
  margin-bottom: 8px;
}
.article-card-title {
  font-size: 16px;
  font-weight: 600;
  color: var(--card-text-color-main);
  cursor: pointer;
  margin: 0;
}
.article-card-title:hover {
  color: var(--accent-color);
}
.article-card-summary {
  font-size: 14px;
  color: var(--card-text-color-secondary);
  margin: 0 0 12px;
  line-height: 1.6;
}
.article-card-footer {
  display: flex;
  align-items: center;
  justify-content: space-between;
}
.article-card-meta {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 13px;
  color: var(--card-text-color-tertiary);
}
.article-card-actions {
  display: flex;
  gap: 8px;
}
</style>
