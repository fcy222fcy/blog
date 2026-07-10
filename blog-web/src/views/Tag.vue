<template>
  <div class="page-view" id="view-tag">
    <!-- 骨架屏加载 -->
    <ArticleListSkeleton v-if="loading" />

    <!-- 错误状态 -->
    <ErrorState v-else-if="errorMsg" :message="errorMsg" @retry="fetchResults" />

    <template v-else>
      <!-- 标签头部信息 -->
      <div class="tag-header" v-if="!loading && !errorMsg">
        <div class="tag-header-inner">
          <div class="tag-header-content">
            <div class="tag-header-label">TAGS</div>
            <div class="tag-header-count">{{ total }} 个页面</div>
            <h1 class="tag-header-title">#{{ tagName || route.params.id }}</h1>
          </div>
        </div>
      </div>

      <!-- 简洁列表 -->
      <div class="list-card">
        <div class="article-list">
          <router-link
            v-for="article in articles"
            :key="article.id"
            :to="'/post/' + article.slug"
            class="article-list-item"
          >
            <span class="article-list-title">{{ article.title }}</span>
            <span class="article-list-date">
              <svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><rect width="18" height="18" x="3" y="4" rx="2" ry="2"></rect><line x1="16" x2="16" y1="2" y2="6"></line><line x1="8" x2="8" y1="2" y2="6"></line><line x1="3" x2="21" y1="10" y2="10"></line></svg>
              {{ formatDate(article.created_at) }}
            </span>
          </router-link>

          <div v-if="articles.length === 0" class="list-empty">
            <svg xmlns="http://www.w3.org/2000/svg" width="40" height="40" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5"><path d="M20.59 13.41l-7.17 7.17a2 2 0 0 1-2.83 0L2 12V2h10l8.59 8.59a2 2 0 0 1 0 2.82z"></path><line x1="7" y1="7" x2="7.01" y2="7"></line></svg>
            <p class="list-empty-title">该标签下暂无文章</p>
            <router-link to="/" class="list-empty-back">返回首页</router-link>
          </div>
        </div>

        <!-- 分页 -->
        <div class="pagination-inner" v-if="totalPage > 1 && articles.length > 0">
          <div class="pagination-buttons">
            <button
              class="pagination-btn"
              :disabled="currentPage <= 1"
              @click="goToPage(currentPage - 1)"
            >
              <svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="m15 18-6-6 6-6"></path></svg>
              上一页
            </button>
            <button class="pagination-btn active">{{ currentPage }}</button>
            <button
              class="pagination-btn"
              :disabled="currentPage >= totalPage"
              @click="goToPage(currentPage + 1)"
            >
              下一页
              <svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="m9 18 6-6-6-6"></path></svg>
            </button>
          </div>
        </div>
      </div>
    </template>
  </div>
</template>

<script setup>
import { ref, watch, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { getArticleList } from '../api/article'
import ArticleListSkeleton from '../components/common/ArticleListSkeleton.vue'
import ErrorState from '../components/common/ErrorState.vue'
import { formatDate } from '../utils/date'

const route = useRoute()

const tagName = ref('')
const articles = ref([])
const loading = ref(false)
const errorMsg = ref('')
const total = ref(0)
const currentPage = ref(1)
const pageSize = 50
const totalPage = ref(0)

const fetchResults = async () => {
  const tagId = Number(route.params.id)
  if (!tagId) return

  loading.value = true
  errorMsg.value = ''
  try {
    const res = await getArticleList({
      tag: tagId,
      page: currentPage.value,
      page_size: pageSize
    })
    const list = res.data?.list || []
    articles.value = list
    total.value = res.data?.total || 0
    totalPage.value = Math.ceil(total.value / pageSize) || 0

    if (list.length > 0 && list[0].tags && list[0].tags.length) {
      const matched = list[0].tags.find(t => String(t.id) === String(tagId))
      if (matched) tagName.value = matched.name
    }
  } catch (err) {
    errorMsg.value = '获取文章失败，请稍后重试'
    console.error('获取标签文章失败:', err)
  } finally {
    loading.value = false
  }
}

const goToPage = (page) => {
  if (page < 1 || page > totalPage.value) return
  currentPage.value = page
  fetchResults()
  window.scrollTo({ top: 0, behavior: 'smooth' })
}

watch(
  () => route.params.id,
  () => {
    currentPage.value = 1
    tagName.value = ''
    fetchResults()
  }
)

onMounted(() => {
  fetchResults()
})
</script>

<style scoped>
/* 统计卡片 */
.stats-grid {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 16px;
  margin-bottom: 20px;
}
.stat-card {
  padding: 18px 20px;
  background: var(--card-background);
  border: 1px solid var(--card-border);
  border-radius: var(--card-border-radius);
  box-shadow: var(--shadow-l1);
  transition: all 0.2s ease;
}
.stat-card:hover {
  transform: translateY(-2px);
  box-shadow: var(--shadow-l2);
  border-color: rgba(var(--accent-color-rgb), 0.3);
}
.stat-label {
  font-size: 0.82rem;
  color: var(--card-text-color-secondary);
  margin-bottom: 6px;
  font-weight: 500;
}
.stat-value {
  font-size: 1.6rem;
  font-weight: 800;
  color: var(--card-text-color-main);
  line-height: 1.2;
}
.tag-stat-name {
  display: flex;
  align-items: center;
}
.stat-tag-pill {
  display: inline-flex;
  align-items: center;
  min-height: 32px;
  padding: 0 14px;
  border-radius: var(--tag-border-radius);
  font-size: 0.95rem;
  font-weight: 700;
  background: rgba(var(--accent-color-rgb), 0.1);
  color: var(--accent-color);
}

/* 简洁列表大卡片 */
.list-card {
  background: var(--card-background);
  border: 1px solid var(--card-border);
  border-radius: var(--card-border-radius);
  box-shadow: var(--shadow-l1);
  overflow: hidden;
}
.article-list {
  display: flex;
  flex-direction: column;
  padding: 4px 0;
}
.article-list-item {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 20px;
  padding: 16px 28px;
  text-decoration: none;
  color: inherit;
  border-bottom: 1px solid var(--card-separator-color);
  transition: background 0.18s ease, padding-left 0.18s ease;
}
.article-list-item:last-child {
  border-bottom: none;
}
.article-list-item:hover {
  background: rgba(var(--accent-color-rgb), 0.05);
  padding-left: 34px;
}
.article-list-title {
  flex: 1;
  min-width: 0;
  font-size: 1rem;
  font-weight: 600;
  line-height: 1.6;
  color: var(--card-text-color-main);
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  transition: color 0.15s ease;
}
.article-list-item:hover .article-list-title {
  color: var(--accent-color);
}
.article-list-date {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  flex-shrink: 0;
  font-size: 0.85rem;
  color: var(--card-text-color-tertiary);
  font-weight: 500;
  font-variant-numeric: tabular-nums;
}
.article-list-date svg {
  flex-shrink: 0;
}

/* 空状态 */
.list-empty {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 56px 24px;
  color: var(--card-text-color-tertiary);
  text-align: center;
}
.list-empty svg {
  opacity: 0.5;
  margin-bottom: 14px;
}
.list-empty-title {
  margin: 0 0 10px;
  font-size: 0.95rem;
  color: var(--card-text-color-secondary);
  font-weight: 600;
}
.list-empty-back {
  color: var(--accent-color);
  text-decoration: none;
  font-weight: 600;
  font-size: 0.9rem;
}
.list-empty-back:hover {
  text-decoration: underline;
}

/* 分页区（卡片内） */
.pagination-inner {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 16px;
  padding: 14px 28px;
  border-top: 1px solid var(--card-separator-color);
  flex-wrap: wrap;
}
.pagination-info {
  color: var(--card-text-color-secondary);
  font-size: 0.9rem;
  font-weight: 500;
}
.pagination-info span {
  color: var(--accent-color);
  font-weight: 700;
}
.pagination-buttons {
  display: flex;
  align-items: center;
  gap: 8px;
}
.pagination-btn {
  display: inline-flex;
  align-items: center;
  gap: 4px;
  padding: 7px 14px;
  border: 1px solid var(--card-separator-color);
  border-radius: 8px;
  background: var(--card-background);
  color: var(--card-text-color-secondary);
  cursor: pointer;
  font-size: 0.85rem;
  font-weight: 600;
  transition: all 0.18s ease;
}
.pagination-btn:hover:not(:disabled) {
  border-color: var(--accent-color);
  color: var(--accent-color);
  background: rgba(var(--accent-color-rgb), 0.06);
  transform: translateY(-1px);
}
.pagination-btn:disabled {
  opacity: 0.4;
  cursor: not-allowed;
}
.pagination-btn.active {
  background: var(--accent-color);
  border-color: var(--accent-color);
  color: var(--accent-color-text);
}
.pagination-btn.active:hover {
  transform: none;
}

@media (max-width: 900px) {
  .stats-grid {
    grid-template-columns: repeat(2, 1fr);
  }
}
@media (max-width: 768px) {
  .article-list-item {
    padding: 14px 20px;
    flex-direction: column;
    align-items: flex-start;
    gap: 6px;
  }
  .article-list-item:hover {
    padding-left: 24px;
  }
  .article-list-title {
    white-space: normal;
    font-size: 0.96rem;
  }
  .pagination-inner {
    justify-content: center;
    padding: 14px 20px;
  }
}
@media (max-width: 480px) {
  .stats-grid {
    grid-template-columns: 1fr 1fr;
    gap: 10px;
  }
  .stat-card {
    padding: 14px 16px;
  }
  .stat-value {
    font-size: 1.3rem;
  }
}
</style>
