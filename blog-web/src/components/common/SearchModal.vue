<template>
  <div class="search-modal" :class="{ open: appStore.searchOpen }" v-if="appStore.searchOpen" @keydown.esc="appStore.closeSearch">
    <div class="search-backdrop" @click="appStore.closeSearch"></div>
    <section class="search-panel">
      <div class="search-input-container">
        <svg class="search-panel-icon" xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="11" cy="11" r="8"></circle><line x1="21" y1="21" x2="16.65" y2="16.65"></line></svg>
        <input
          ref="searchInput"
          v-model="keyword"
          type="search"
          class="search-panel-input"
          placeholder="输入关键词搜索文章..."
          @input="handleInput"
          @keyup.enter="goToSearchPage"
        >
        <div class="search-panel-actions">
          <span class="search-panel-tip">
            <kbd>Enter</kbd> 搜索
          </span>
          <span class="search-panel-divider"></span>
          <button class="search-panel-close" @click="appStore.closeSearch" title="关闭 (Esc)">
            <kbd>Esc</kbd>
          </button>
        </div>
      </div>
      <div class="search-results">
        <div v-if="searching" class="search-loading">
          <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" class="animate-spin"><path d="M21 12a9 9 0 1 1-6.219-8.56"/></svg>
          搜索中...
        </div>
        <template v-else>
          <div v-for="article in results" :key="article.id" class="search-result-item" @click="goToArticle(article.slug)">
            <div class="search-result-content">
              <h3 class="search-result-title">
                <span class="search-result-category" v-if="article.category">{{ article.category.name }}</span>
                {{ article.title }}
              </h3>
              <p class="search-result-summary">{{ article.summary }}</p>
              <div class="search-result-meta">
                <span class="search-result-date">{{ formatDateSimple(article.created_at) }}</span>
                <span v-if="article.view_count" class="search-result-views">
                  <svg xmlns="http://www.w3.org/2000/svg" width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M2 12s3-7 10-7 10 7 10 7-3 7-10 7-10-7-10-7Z"></path><circle cx="12" cy="12" r="3"></circle></svg>
                  {{ article.view_count }}
                </span>
              </div>
            </div>
            <svg class="search-result-arrow" xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="m9 18 6-6-6-6"/></svg>
          </div>
        </template>
        <div v-if="keyword && results.length === 0 && !searching" class="search-empty">
          <svg xmlns="http://www.w3.org/2000/svg" width="32" height="32" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5"><circle cx="11" cy="11" r="8"></circle><line x1="21" y1="21" x2="16.65" y2="16.65"></line><line x1="8" y1="8" x2="14" y2="14"></line><line x1="14" y1="8" x2="8" y2="14"></line></svg>
          <p>没有找到与「{{ keyword }}」相关的文章</p>
        </div>
        <div v-if="keyword && results.length > 0" class="search-view-all">
          <a href="javascript:void(0)" @click="goToSearchPage">
            查看全部搜索结果
            <svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="m9 18 6-6-6-6"/></svg>
          </a>
        </div>
      </div>
    </section>
  </div>
</template>

<script setup>
import { ref, watch, nextTick, onMounted, onUnmounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAppStore } from '../../stores/app'
import { searchArticles } from '../../api/article'
import { formatDate } from '../../utils/date'

const router = useRouter()
const appStore = useAppStore()
const keyword = ref('')
const results = ref([])
const searching = ref(false)
const searchInput = ref(null)
let searchTimer = null

const formatDateSimple = (date) => {
  return formatDate(date).split(' ')[0]
}

const goToArticle = (slug) => {
  appStore.closeSearch()
  router.push('/post/' + slug)
}

watch(() => appStore.searchOpen, (val) => {
  if (val) {
    keyword.value = ''
    results.value = []
    nextTick(() => searchInput.value?.focus())
  }
})

const handleInput = () => {
  clearTimeout(searchTimer)
  if (!keyword.value.trim()) {
    results.value = []
    return
  }
  searchTimer = setTimeout(async () => {
    searching.value = true
    try {
      const res = await searchArticles({ keyword: keyword.value, page: 1, page_size: 5 })
      results.value = res.data?.list || []
    } catch (e) {
      console.error(e)
    } finally {
      searching.value = false
    }
  }, 300)
}

const goToSearchPage = () => {
  if (!keyword.value.trim()) return
  appStore.closeSearch()
  router.push({ name: 'Search', query: { q: keyword.value } })
}
</script>

<style scoped>
.animate-spin {
  animation: spin 1s linear infinite;
}
@keyframes spin {
  from { transform: rotate(0deg); }
  to { transform: rotate(360deg); }
}
</style>
