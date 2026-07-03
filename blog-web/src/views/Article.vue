<template>
  <div class="page-view">
    <router-link to="/" class="back-btn">← 返回列表</router-link>

    <!-- 骨架屏加载 -->
    <ArticleDetailSkeleton v-if="articleStore.loading" />

    <!-- 错误状态 -->
    <ErrorState v-else-if="articleStore.error" :message="articleStore.error" @retry="loadArticle" />

    <!-- 文章内容 -->
    <template v-else-if="articleStore.currentArticle">
      <article class="post-detail">
        <h1 class="post-title">{{ articleStore.currentArticle.title }}</h1>
        <div class="post-meta">
          <span class="meta-item">
            <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect x="3" y="4" width="18" height="18" rx="2" ry="2"></rect><line x1="16" y1="2" x2="16" y2="6"></line><line x1="8" y1="2" x2="8" y2="6"></line><line x1="3" y1="10" x2="21" y2="10"></line></svg>
            {{ formatDate(articleStore.currentArticle.created_at) }}
          </span>
          <span class="meta-item" v-if="articleStore.currentArticle.updated_at && articleStore.currentArticle.updated_at !== articleStore.currentArticle.created_at">
            <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M12 20h9"></path><path d="M16.5 3.5a2.121 2.121 0 0 1 3 3L7 19l-4 1 1-4L16.5 3.5z"></path></svg>
            更新于 {{ formatDate(articleStore.currentArticle.updated_at) }}
          </span>
          <span class="meta-item">
            <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"></circle><polyline points="12 6 12 12 16 14"></polyline></svg>
            时长 {{ readingTime }} 分钟
          </span>
          <span class="meta-item">
            <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z"></path><circle cx="12" cy="12" r="3"></circle></svg>
            浏览量 {{ articleStore.currentArticle.view_count }}
          </span>
        </div>
        <div class="post-content" v-html="renderedContent"></div>
      </article>

      <!-- 评论区 -->
      <CommentSection />
    </template>

    <!-- 文章不存在 -->
    <div v-else class="empty-state">
      <p>文章不存在</p>
      <router-link to="/" class="back-home-link">返回首页</router-link>
    </div>
  </div>
</template>

<script setup>
import { computed, onMounted, watch, onUnmounted } from 'vue'
import { useRoute } from 'vue-router'
import { useArticleStore } from '../stores/article'
import { marked } from 'marked'
import CommentSection from '../components/comment/CommentSection.vue'
import ArticleDetailSkeleton from '../components/common/ArticleDetailSkeleton.vue'
import ErrorState from '../components/common/ErrorState.vue'
import { updateMetaTags, addStructuredData, resetMetaTags } from '../utils/seo'

const route = useRoute()
const articleStore = useArticleStore()

// 简单 HTML 净化：移除事件属性和危险标签
const sanitizeHtml = (html) => {
  // 移除 on* 事件属性（onerror, onclick, onload 等）
  let clean = html.replace(/\s+on\w+\s*=\s*(?:"[^"]*"|'[^']*'|[^\s>]+)/gi, '')
  // 移除 javascript: 协议
  clean = clean.replace(/href\s*=\s*(?:"javascript:[^"]*"|'javascript:[^']*')/gi, '')
  clean = clean.replace(/src\s*=\s*(?:"javascript:[^"]*"|'javascript:[^']*')/gi, '')
  // 移除 <script> 和 <iframe> 标签
  clean = clean.replace(/<script[\s\S]*?<\/script>/gi, '')
  clean = clean.replace(/<iframe[\s\S]*?<\/iframe>/gi, '')
  return clean
}

const renderedContent = computed(() => {
  if (!articleStore.currentArticle?.content) return ''
  return sanitizeHtml(marked(articleStore.currentArticle.content))
})

// 计算阅读时长（中文约 400 字/分钟）
const readingTime = computed(() => {
  if (!articleStore.currentArticle?.content) return 1
  const content = articleStore.currentArticle.content
  // 去除 markdown 标记和空白字符后计算字数
  const text = content.replace(/[#*`>\[\]()!\-\n\r]/g, '').trim()
  const charCount = text.length || 0
  return Math.max(1, Math.ceil(charCount / 400))
})

const formatDate = (dateStr) => {
  if (!dateStr) return ''
  return dateStr.split('T')[0]
}

const loadArticle = () => {
  const slug = route.params.slug
  if (slug) articleStore.fetchArticleDetail(slug)
}

// SEO: 文章加载后更新 Meta 标签
watch(
  () => articleStore.currentArticle,
  (article) => {
    if (article) {
      updateMetaTags(article)
      addStructuredData(article)
    }
  },
  { immediate: true }
)

onMounted(loadArticle)
watch(() => route.params.slug, loadArticle)

onUnmounted(() => {
  resetMetaTags()
})
</script>

<style scoped>
/* 文章详情页样式在全局main.css中定义 */
.post-meta {
  display: flex;
  align-items: center;
  gap: 20px;
  margin: 12px 0 24px;
  color: #999;
  font-size: 0.9rem;
}

.meta-item {
  display: inline-flex;
  align-items: center;
  gap: 6px;
}

.meta-item svg {
  color: #999;
  flex-shrink: 0;
}

.empty-state {
  text-align: center;
  padding: 60px 20px;
}

.empty-state p {
  color: var(--text-color-secondary, #6b7280);
  margin-bottom: 16px;
}

.back-home-link {
  color: var(--accent-color, #3b82f6);
  text-decoration: none;
}

.back-home-link:hover {
  text-decoration: underline;
}
</style>
