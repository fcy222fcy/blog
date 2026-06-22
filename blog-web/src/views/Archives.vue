<template>
  <div class="page-view">
    <h1 class="page-heading">文章归档</h1>

    <div v-for="group in articleStore.archives" :key="group.year" class="year-group">
      <h2 class="year-heading">{{ group.year }}</h2>
      <ul class="archive-list">
        <li v-for="article in group.articles" :key="article.id">
          <router-link :to="'/post/' + article.slug">
            <span class="archive-title">{{ article.title }}</span>
            <span class="archive-date">{{ formatDate(article.created_at) }}</span>
          </router-link>
        </li>
      </ul>
    </div>

    <div v-if="articleStore.loading" class="loading">加载中...</div>
  </div>
</template>

<script setup>
import { onMounted } from 'vue'
import { useArticleStore } from '../stores/article'

const articleStore = useArticleStore()

const formatDate = (dateStr) => {
  if (!dateStr) return ''
  return dateStr.split('T')[0]
}

onMounted(() => {
  articleStore.fetchArchives()
})
</script>

<style scoped>
/* 归档页面样式在全局main.css中定义 */
.loading { text-align: center; padding: 40px; color: var(--card-text-color-secondary); }
</style>
