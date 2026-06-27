<template>
  <div class="page-view">
    <router-link to="/" class="back-btn">← 返回列表</router-link>

    <article v-if="articleStore.currentArticle" class="post-detail">
      <h1 class="post-title">{{ articleStore.currentArticle.title }}</h1>
      <div class="post-meta">
        <span>📅 {{ formatDate(articleStore.currentArticle.created_at) }}</span>
        <span>👁️ {{ articleStore.currentArticle.view_count }}</span>
        <span>❤️ {{ articleStore.currentArticle.like_count }}</span>
      </div>
      <div class="post-content" v-html="renderedContent"></div>
    </article>

    <!-- 评论区 -->
    <CommentSection v-if="articleStore.currentArticle" />

    <div v-if="articleStore.loading" class="loading">加载中...</div>
  </div>
</template>

<script setup>
import { computed, onMounted, watch } from 'vue'
import { useRoute } from 'vue-router'
import { useArticleStore } from '../stores/article'
import { marked } from 'marked'
import CommentSection from '../components/comment/CommentSection.vue'

const route = useRoute()
const articleStore = useArticleStore()

const renderedContent = computed(() => {
  if (!articleStore.currentArticle?.content) return ''
  return marked(articleStore.currentArticle.content)
})

const formatDate = (dateStr) => {
  if (!dateStr) return ''
  return dateStr.split('T')[0]
}

const loadArticle = () => {
  const slug = route.params.slug
  if (slug) articleStore.fetchArticleDetail(slug)
}

onMounted(loadArticle)
watch(() => route.params.slug, loadArticle)
</script>

<style scoped>
/* 文章详情页样式在全局main.css中定义 */
.loading { text-align: center; padding: 40px; color: var(--card-text-color-secondary); }
</style>
