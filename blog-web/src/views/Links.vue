<template>
  <div class="page-view">
    <div class="page-hero">
      <div class="page-hero-overlay">
        <h1 class="page-heading">友链</h1>
        <p class="page-hero-desc">我的朋友们</p>
      </div>
    </div>

    <!-- 加载状态 -->
    <Loading v-if="loading" text="加载友链中..." />

    <!-- 错误状态 -->
    <div v-else-if="error" class="error-container">
      <p class="error-message">{{ error }}</p>
      <button class="retry-btn" @click="fetchLinks">
        <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M21.5 2v6h-6M2.5 22v-6h6M2 11.5a10 10 0 0 1 18.8-4.3M22 12.5a10 10 0 0 1-18.8 4.2"/></svg>
        重试
      </button>
    </div>

    <!-- 友链内容 -->
    <template v-else>
      <div class="friend-grid">
        <a v-for="link in links" :key="link.id" :href="link.url" class="friend-card" target="_blank">
          <div class="friend-avatar-img" v-if="link.avatar && (link.avatar.startsWith('http://') || link.avatar.startsWith('https://') || link.avatar.startsWith('/'))">
            <img :src="link.avatar" :alt="link.name">
          </div>
          <div class="friend-avatar-emoji" v-else>
            {{ link.avatar || '🔗' }}
          </div>
          <h3>{{ link.name }}</h3>
          <p>{{ link.description }}</p>
        </a>
      </div>

      <div v-if="links.length === 0" class="empty">暂无友链</div>

      <div class="link-info-box">
        <h2>🤝 交换友链</h2>
        <p>如果你想出现在这个页面，请在下方留言。</p>
        <h2>🚩 要求</h2>
        <ul>
          <li><strong>原创内容</strong>: 定期更新原创内容。</li>
          <li><strong>合规</strong>: 内容符合相关法律法规。</li>
          <li><strong>稳定访问</strong>: 网站可正常访问，且没有过多广告。</li>
        </ul>
      </div>
    </template>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { getLinkList } from '../api/link'
import Loading from '../components/common/Loading.vue'

const links = ref([])
const loading = ref(false)
const error = ref(null)

const fetchLinks = async () => {
  loading.value = true
  error.value = null
  try {
    const res = await getLinkList()
    links.value = res.data || []
  } catch (e) {
    error.value = '加载友链失败，请稍后重试'
    console.error(e)
  } finally {
    loading.value = false
  }
}

onMounted(fetchLinks)
</script>

<style scoped>
/* 友链页面样式在全局main.css中定义 */
.friend-avatar-img {
  width: 80px;
  height: 80px;
  border-radius: 50%;
  overflow: hidden;
  margin-bottom: 12px;
}
.friend-avatar-img img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}
.friend-avatar-emoji {
  width: 80px;
  height: 80px;
  border-radius: 50%;
  background: linear-gradient(135deg, #667eea, #764ba2);
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 36px;
  margin-bottom: 12px;
}
</style>
