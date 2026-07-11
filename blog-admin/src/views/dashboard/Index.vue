<template>
  <div>
    <!-- 统计卡片（加载态 + 真实态） -->
    <SkeletonLoader v-if="loading" type="stats" :count="4" />
    <div v-else class="stats-grid">
      <div class="stat-card">
        <div class="stat-label">文章总数</div>
        <div class="stat-value">{{ stats.article_count || 0 }}</div>
      </div>
      <div class="stat-card">
        <div class="stat-label">已发布</div>
        <div class="stat-value">{{ stats.published_count || 0 }}</div>
      </div>
      <div class="stat-card">
        <div class="stat-label">总浏览量</div>
        <div class="stat-value">{{ stats.total_views || 0 }}</div>
      </div>
      <div class="stat-card">
        <div class="stat-label">评论总数</div>
        <div class="stat-value">{{ stats.comment_count || 0 }}</div>
      </div>
    </div>

    <div class="dashboard-grid">
      <div class="dashboard-card" @click="$router.push('/articles')">
        <div class="dashboard-card-icon">📝</div>
        <div class="dashboard-card-info">
          <div class="dashboard-card-title">文章管理</div>
          <div class="dashboard-card-desc">管理博客文章，发布、编辑、删除</div>
          <div class="dashboard-card-meta">共 {{ stats.article_count || 0 }} 篇文章</div>
        </div>
        <div class="dashboard-card-arrow">→</div>
      </div>


      <div class="dashboard-card" @click="$router.push('/categories')">
        <div class="dashboard-card-icon">📂</div>
        <div class="dashboard-card-info">
          <div class="dashboard-card-title">分类管理</div>
          <div class="dashboard-card-desc">管理文章分类，建立内容结构</div>
          <div class="dashboard-card-meta">搭建网站、软件开发、生活记录</div>
        </div>
        <div class="dashboard-card-arrow">→</div>
      </div>

      <div class="dashboard-card" @click="$router.push('/tags')">
        <div class="dashboard-card-icon">🏷️</div>
        <div class="dashboard-card-info">
          <div class="dashboard-card-title">标签管理</div>
          <div class="dashboard-card-desc">管理文章标签，方便内容检索</div>
          <div class="dashboard-card-meta">Hugo、Go、AI、开源</div>
        </div>
        <div class="dashboard-card-arrow">→</div>
      </div>

      <div class="dashboard-card" @click="$router.push('/comments')">
        <div class="dashboard-card-icon">💬</div>
        <div class="dashboard-card-info">
          <div class="dashboard-card-title">评论管理</div>
          <div class="dashboard-card-desc">审核和管理读者评论</div>
          <div class="dashboard-card-meta">待审核、已通过、已拒绝</div>
        </div>
        <div class="dashboard-card-arrow">→</div>
      </div>

      <div class="dashboard-card" @click="$router.push('/daily-question')">
        <div class="dashboard-card-icon">💡</div>
        <div class="dashboard-card-info">
          <div class="dashboard-card-title">每日一问</div>
          <div class="dashboard-card-desc">管理每日问答内容</div>
          <div class="dashboard-card-meta">今日问题已设置</div>
        </div>
        <div class="dashboard-card-arrow">→</div>
      </div>

      <div class="dashboard-card" @click="$router.push('/about')">
        <div class="dashboard-card-icon">👤</div>
        <div class="dashboard-card-info">
          <div class="dashboard-card-title">关于我</div>
          <div class="dashboard-card-desc">编辑个人简介和社交信息</div>
          <div class="dashboard-card-meta">头像、简介、社交链接</div>
        </div>
        <div class="dashboard-card-arrow">→</div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import request from '../../api/request'
import SkeletonLoader from '../../components/common/SkeletonLoader.vue'

const stats = ref({})
const loading = ref(true)

onMounted(async () => {
  try {
    const res = await request.get('/admin/dashboard/stats')
    stats.value = res.data || {}
  } catch (e) { console.error(e) }
  loading.value = false
})
</script>
