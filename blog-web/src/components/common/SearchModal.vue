<template>
  <div class="search-modal" :class="{ open: appStore.searchOpen }" v-if="appStore.searchOpen">
    <div class="search-backdrop" @click="appStore.closeSearch"></div>
    <section class="search-panel">
      <div class="search-header">
        <h2>搜索文章</h2>
        <button class="close-btn" @click="appStore.closeSearch">&times;</button>
      </div>
      <input ref="searchInput" v-model="keyword" type="search" placeholder="输入关键词..." @input="handleInput" @keyup.enter="goToSearchPage">
      <div class="search-results">
        <div v-for="article in results" :key="article.id" class="search-result-item">
          <a :href="'#/post/' + article.slug" @click="appStore.closeSearch">
            <h3>{{ article.title }}</h3>
            <p>{{ article.summary }}</p>
          </a>
        </div>
        <div v-if="keyword && results.length === 0 && !searching" class="search-empty">没有找到相关文章</div>
        <div v-if="keyword" class="search-view-all">
          <a href="javascript:void(0)" @click="goToSearchPage">查看全部搜索结果 &rarr;</a>
        </div>
      </div>
    </section>
  </div>
</template>

<script setup>
import { ref, watch, nextTick } from 'vue'
import { useRouter } from 'vue-router'
import { useAppStore } from '../../stores/app'
import { searchArticles } from '../../api/article'

const router = useRouter()
const appStore = useAppStore()
const keyword = ref('')
const results = ref([])
const searching = ref(false)
const searchInput = ref(null)
let searchTimer = null

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
/* 搜索弹窗样式在全局main.css中定义 */
</style>
