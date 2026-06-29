<template>
  <div class="page-view">
    <div class="page-hero">
      <div class="page-hero-overlay">
        <h1 class="page-heading">友链</h1>
        <p class="page-hero-desc">我的朋友们</p>
      </div>
    </div>

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
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { getLinkList } from '../api/link'

const links = ref([])

onMounted(async () => {
  try {
    const res = await getLinkList()
    links.value = res.data || []
  } catch (e) { console.error(e) }
})
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
