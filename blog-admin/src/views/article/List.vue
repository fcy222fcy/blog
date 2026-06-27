<template>
  <div>
    <div class="stats-grid">
      <div class="stat-card">
        <div class="stat-label">文章总数</div>
        <div class="stat-value">{{ total }}</div>
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
        <div style="display: flex; gap: 12px; flex-wrap: wrap; align-items: center;">
          <div class="search-box">
            <span class="search-box-icon">⌕</span>
            <input type="text" v-model="keyword" placeholder="搜索文章..." @input="loadArticles">
          </div>
          <select class="form-select" style="width: auto; padding: 9px 32px 9px 12px;" v-model="categoryFilter" @change="loadArticles">
            <option value="">全部分类</option>
            <option v-for="cat in categories" :key="cat.id" :value="cat.id">{{ cat.name }}</option>
          </select>
          <select class="form-select" style="width: auto; padding: 9px 32px 9px 12px;" v-model="statusFilter" @change="loadArticles">
            <option value="">全部状态</option>
            <option value="published">已发布</option>
            <option value="draft">草稿</option>
          </select>
          <button class="btn btn-primary" @click="$router.push('/articles/edit')">
            <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><line x1="12" y1="5" x2="12" y2="19"></line><line x1="5" y1="12" x2="19" y2="12"></line></svg>
            <span>新建文章</span>
          </button>
        </div>
      </div>

      <div class="article-cards-container">
        <div v-for="article in articles" :key="article.id" class="article-card-item">
          <div class="article-card-header">
            <h3 class="article-card-title" @click="$router.push('/articles/edit/' + article.id)">{{ article.title }}</h3>
            <span class="status-badge" :class="article.status === 'published' ? 'status-published' : 'status-draft'">
              {{ article.status === 'published' ? '已发布' : '草稿' }}
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
            <div class="article-card-actions">
              <button class="action-btn btn-edit btn-sm" @click="$router.push('/articles/edit/' + article.id)">编辑</button>
              <button class="action-btn btn-delete btn-sm" @click="handleDelete(article.id)">删除</button>
            </div>
          </div>
        </div>
        <div v-if="articles.length === 0" style="text-align: center; padding: 40px; color: var(--card-text-color-tertiary);">
          暂无文章
        </div>
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
import { ref, onMounted } from 'vue'
import { getArticleList, deleteArticle } from '../../api/article'
import { getCategoryList } from '../../api/category'
import { ElMessage, ElMessageBox } from 'element-plus'

const articles = ref([])
const categories = ref([])
const loading = ref(false)
const page = ref(1)
const total = ref(0)
const keyword = ref('')
const categoryFilter = ref('')
const statusFilter = ref('')
const stats = ref({ published: 0, draft: 0, totalViews: 0 })

const formatDate = (d) => d ? d.split('T')[0] : ''

const loadArticles = async () => {
  loading.value = true
  try {
    const params = { page: page.value, page_size: 10 }
    if (keyword.value) params.keyword = keyword.value
    if (categoryFilter.value) params.category_id = categoryFilter.value
    if (statusFilter.value) params.status = statusFilter.value
    const res = await getArticleList(params)
    articles.value = res.data?.list || []
    total.value = res.data?.total || 0
    // 统计状态
    stats.value.published = articles.value.filter(a => a.status === 'published').length
    stats.value.draft = articles.value.filter(a => a.status === 'draft').length
    stats.value.totalViews = articles.value.reduce((sum, a) => sum + (a.view_count || 0), 0)
  } catch (e) { console.error(e) }
  loading.value = false
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

onMounted(() => {
  loadArticles()
  loadCategories()
})
</script>

<style scoped>
.article-cards-container {
  padding: 0;
}
.article-card-item {
  padding: 20px 28px;
  border-bottom: 1px solid var(--card-separator-color);
  transition: background 0.15s ease;
}
.article-card-item:hover {
  background: rgba(var(--accent-color-rgb), 0.02);
}
.article-card-item:last-child {
  border-bottom: none;
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
