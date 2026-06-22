<template>
  <div class="search-modal" :class="{ open: appStore.searchOpen }" v-if="appStore.searchOpen">
    <div class="search-backdrop" @click="appStore.closeSearch"></div>
    <section class="search-panel">
      <div class="search-header">
        <h2>搜索文章</h2>
        <button class="close-btn" @click="appStore.closeSearch">×</button>
      </div>
      <input ref="searchInput" v-model="keyword" type="search" placeholder="输入关键词..." @input="handleSearch">
      <div class="search-results">
        <div v-for="article in results" :key="article.id" class="search-result-item">
          <a :href="'#/post/' + article.slug" @click="appStore.closeSearch">
            <h3>{{ article.title }}</h3>
            <p>{{ article.summary }}</p>
          </a>
        </div>
        <div v-if="keyword && results.length === 0" class="search-empty">没有找到相关文章</div>
      </div>
    </section>
  </div>
</template>

<script setup>
import { ref, watch, nextTick } from 'vue'
import { useAppStore } from '../../stores/app'
import { getArticleList } from '../../api/article'

const appStore = useAppStore()
const keyword = ref('')
const results = ref([])
const searchInput = ref(null)

watch(() => appStore.searchOpen, (val) => {
  if (val) nextTick(() => searchInput.value?.focus())
})

const handleSearch = async () => {
  if (!keyword.value.trim()) { results.value = []; return }
  try {
    const res = await getArticleList({ keyword: keyword.value, page: 1, page_size: 10 })
    results.value = res.data?.list || []
  } catch (e) { console.error(e) }
}
</script>

<style scoped>
/* 搜索弹窗样式在全局main.css中定义 */
</style>
