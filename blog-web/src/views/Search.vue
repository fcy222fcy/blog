<template>
  <div class="page-view">
    <h1 class="page-heading">
      <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="11" cy="11" r="8"></circle><line x1="21" y1="21" x2="16.65" y2="16.65"></line></svg>
      搜索结果：{{ keyword }}
    </h1>

    <!-- 搜索输入框 -->
    <div class="search-input-wrapper">
      <input
        ref="searchInputRef"
        v-model="inputKeyword"
        type="search"
        placeholder="输入关键词搜索文章..."
        class="search-input"
        @keyup.enter="handleSearch"
      >
      <button class="search-btn" @click="handleSearch">
        <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="11" cy="11" r="8"></circle><line x1="21" y1="21" x2="16.65" y2="16.65"></line></svg>
      </button>
    </div>

    <!-- 加载状态 -->
    <Loading v-if="loading" text="搜索中..." />

    <!-- 错误状态 -->
    <div v-else-if="errorMsg" class="error-container">
      <p class="error-message">{{ errorMsg }}</p>
      <button class="retry-btn" @click="fetchResults">
        <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M21.5 2v6h-6M2.5 22v-6h6M2 11.5a10 10 0 0 1 18.8-4.3M22 12.5a10 10 0 0 1-18.8 4.2"/></svg>
        重试
      </button>
    </div>

    <!-- 搜索结果 -->
    <template v-else>
      <div v-if="articles.length > 0" class="search-results-info">
        找到 <strong>{{ total }}</strong> 篇相关文章
      </div>

      <article v-for="article in articles" :key="article.id" class="article-card">
        <router-link :to="'/post/' + article.slug" class="card-link">
          <span class="category-pill" :class="'category-' + (article.category?.slug || 'default')">{{ article.category?.name || '未分类' }}</span>
          <h2>{{ article.title }}</h2>
          <p>{{ article.summary }}</p>
          <dl class="article-meta">
            <div>
              <dt>日期</dt>
              <dd>
                <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><rect width="18" height="18" x="3" y="4" rx="2" ry="2"></rect><line x1="16" x2="16" y1="2" y2="6"></line><line x1="8" x2="8" y1="2" y2="6"></line><line x1="3" x2="21" y1="10" y2="10"></line></svg>
                {{ formatDate(article.created_at) }}
              </dd>
            </div>
            <div>
              <dt>浏览量</dt>
              <dd>
                <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M2 12s3-7 10-7 10 7 10 7-3 7-10 7-10-7-10-7Z"></path><circle cx="12" cy="12" r="3"></circle></svg>
                {{ article.view_count }}
              </dd>
            </div>
          </dl>
        </router-link>
      </article>

      <!-- 分页 -->
      <div v-if="totalPage > 1" class="pagination">
        <button
          class="page-btn"
          :disabled="currentPage <= 1"
          @click="goToPage(currentPage - 1)"
        >
          <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="m15 18-6-6 6-6"></path></svg>
          上一页
        </button>
        <span class="page-info">{{ currentPage }} / {{ totalPage }}</span>
        <button
          class="page-btn"
          :disabled="currentPage >= totalPage"
          @click="goToPage(currentPage + 1)"
        >
          下一页
          <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="m9 18 6-6-6-6"></path></svg>
        </button>
      </div>

      <div v-if="!loading && keyword && articles.length === 0" class="empty">
        <svg xmlns="http://www.w3.org/2000/svg" width="48" height="48" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5"><circle cx="11" cy="11" r="8"></circle><line x1="21" y1="21" x2="16.65" y2="16.65"></line><line x1="8" y1="8" x2="14" y2="14"></line><line x1="14" y1="8" x2="8" y2="14"></line></svg>
        <p>没有找到与「{{ keyword }}」相关的文章</p>
        <p class="empty-hint">换个关键词试试？</p>
      </div>
    </template>
  </div>
</template>

<script setup>
import { ref, watch, onMounted, nextTick } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { searchArticles } from '../api/article'
import Loading from '../components/common/Loading.vue'

const route = useRoute()
const router = useRouter()

const keyword = ref('')
const inputKeyword = ref('')
const articles = ref([])
const loading = ref(false)
const errorMsg = ref('')
const total = ref(0)
const currentPage = ref(1)
const pageSize = 10
const totalPage = ref(0)
const searchInputRef = ref(null)

const formatDate = (dateStr) => {
  if (!dateStr) return ''
  return dateStr.split('T')[0]
}

const fetchResults = async () => {
  if (!keyword.value.trim()) return

  loading.value = true
  errorMsg.value = ''
  try {
    const res = await searchArticles({
      keyword: keyword.value,
      page: currentPage.value,
      page_size: pageSize
    })
    articles.value = res.data?.list || []
    total.value = res.data?.total || 0
    totalPage.value = res.data?.total_page || 0
  } catch (err) {
    errorMsg.value = '搜索失败，请稍后重试'
    console.error('搜索失败:', err)
  } finally {
    loading.value = false
  }
}

const handleSearch = () => {
  const kw = inputKeyword.value.trim()
  if (kw) {
    router.push({ name: 'Search', query: { q: kw } })
  }
}

const goToPage = (page) => {
  if (page < 1 || page > totalPage.value) return
  currentPage.value = page
  fetchResults()
  window.scrollTo({ top: 0, behavior: 'smooth' })
}

// 监听路由 query 变化
watch(
  () => route.query.q,
  (newQ) => {
    if (newQ) {
      keyword.value = newQ
      inputKeyword.value = newQ
      currentPage.value = 1
      fetchResults()
    }
  }
)

onMounted(() => {
  const q = route.query.q
  if (q) {
    keyword.value = q
    inputKeyword.value = q
    fetchResults()
  }
  nextTick(() => searchInputRef.value?.focus())
})
</script>

<style scoped>
.search-input-wrapper {
  display: flex;
  gap: 8px;
  margin-bottom: 24px;
}

.search-input {
  flex: 1;
  padding: 12px 16px;
  border: 1px solid var(--card-border, #e2e8f0);
  border-radius: 8px;
  background: var(--card-background, #fff);
  color: var(--card-text-color, #333);
  font-size: 16px;
  outline: none;
  transition: border-color 0.2s;
}

.search-input:focus {
  border-color: var(--accent-color, #3b82f6);
}

.search-btn {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 48px;
  height: 48px;
  border: none;
  border-radius: 8px;
  background: var(--accent-color, #3b82f6);
  color: white;
  cursor: pointer;
  transition: background-color 0.2s;
}

.search-btn:hover {
  background: var(--accent-color-hover, #2563eb);
}

.search-results-info {
  margin-bottom: 16px;
  color: var(--card-text-color-secondary, #666);
  font-size: 14px;
}

.search-results-info strong {
  color: var(--accent-color, #3b82f6);
}

.empty {
  text-align: center;
  padding: 60px 20px;
  color: var(--card-text-color-secondary, #999);
}

.empty svg {
  margin-bottom: 16px;
  opacity: 0.5;
}

.empty p {
  margin: 8px 0;
}

.empty-hint {
  font-size: 14px;
  opacity: 0.7;
}

.error-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  min-height: 200px;
  gap: 16px;
}

.error-message {
  color: var(--error-color, #ef4444);
  margin: 0;
}

.retry-btn {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  padding: 8px 16px;
  background-color: var(--accent-color, #3b82f6);
  color: white;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  font-size: 14px;
  transition: background-color 0.2s;
}

.retry-btn:hover {
  background-color: var(--accent-color-hover, #2563eb);
}

/* 分页 */
.pagination {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 16px;
  margin-top: 32px;
  padding: 16px 0;
}

.page-btn {
  display: inline-flex;
  align-items: center;
  gap: 4px;
  padding: 8px 16px;
  border: 1px solid var(--card-border, #e2e8f0);
  border-radius: 6px;
  background: var(--card-background, #fff);
  color: var(--card-text-color, #333);
  cursor: pointer;
  font-size: 14px;
  transition: all 0.2s;
}

.page-btn:hover:not(:disabled) {
  border-color: var(--accent-color, #3b82f6);
  color: var(--accent-color, #3b82f6);
}

.page-btn:disabled {
  opacity: 0.4;
  cursor: not-allowed;
}

.page-info {
  color: var(--card-text-color-secondary, #666);
  font-size: 14px;
}
</style>
