<template>
  <div>
    <div class="stats-grid">
      <div class="stat-card">
        <div class="stat-label">影视总数</div>
        <div class="stat-value">{{ items.length }}</div>
      </div>
      <div class="stat-card">
        <div class="stat-label">已看完</div>
        <div class="stat-value" style="color: var(--success-color);">{{ items.filter(i => i.status === '已看完').length }}</div>
      </div>
    </div>

    <div style="display: flex; gap: 10px; margin-bottom: 24px;">
      <button class="btn" :class="activeTab === 'all' ? 'btn-primary' : 'btn-secondary'" @click="activeTab = 'all'">全部</button>
      <button class="btn" :class="activeTab === 'movie' ? 'btn-primary' : 'btn-secondary'" @click="activeTab = 'movie'">🎬 电影</button>
      <button class="btn" :class="activeTab === 'tv' ? 'btn-primary' : 'btn-secondary'" @click="activeTab = 'tv'">📺 剧集</button>
      <button class="btn" :class="activeTab === 'game' ? 'btn-primary' : 'btn-secondary'" @click="activeTab = 'game'">🎮 游戏</button>
    </div>

    <div style="display: flex; justify-content: space-between; align-items: center; margin-bottom: 20px;">
      <div class="search-box">
        <span class="search-box-icon">⌕</span>
        <input type="text" v-model="keyword" placeholder="搜索影视游戏...">
      </div>
      <button class="btn btn-primary" @click="showModal = true; resetForm()">
        <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><line x1="12" y1="5" x2="12" y2="19"></line><line x1="5" y1="12" x2="19" y2="12"></line></svg>
        <span>添加条目</span>
      </button>
    </div>

    <div class="media-cards">
      <div v-for="item in filteredItems" :key="item.id" class="media-card">
        <div class="media-card-cover" :style="{ background: getGradient(item.id) }">{{ item.icon }}</div>
        <div class="media-card-info">
          <div class="media-card-title">{{ item.title }}</div>
          <div class="media-card-subtitle">{{ item.subtitle }}</div>
          <div class="media-card-meta">
            <span class="media-card-tag" :style="{ background: 'rgba(var(--accent-color-rgb), 0.1)', color: 'var(--accent-color)' }">{{ item.type }}</span>
            <span class="media-card-rating">⭐ {{ item.rating }}</span>
            <span class="media-card-year">{{ item.year }}</span>
          </div>
          <div class="media-card-status" :class="item.status === '已看完' ? 'status-published' : 'status-draft'">{{ item.status }}</div>
        </div>
        <div class="media-card-actions">
          <button class="action-btn btn-edit btn-sm">编辑</button>
          <button class="action-btn btn-delete btn-sm">删除</button>
        </div>
      </div>
    </div>

    <div v-if="filteredItems.length === 0" class="card">
      <div class="card-body" style="text-align: center; color: var(--card-text-color-tertiary);">暂无条目</div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'

const keyword = ref('')
const activeTab = ref('all')
const showModal = ref(false)

const items = ref([
  { id: 1, title: '垫底辣妹', subtitle: 'Japanese Drama', type: '电影', icon: '🎬', rating: 8.5, year: 2015, status: '已看完', category: 'movie' },
  { id: 2, title: '你的名字', subtitle: 'Your Name', type: '电影', icon: '🎬', rating: 8.4, year: 2016, status: '已看完', category: 'movie' },
  { id: 3, title: '三体', subtitle: 'Three-Body', type: '剧集', icon: '📺', rating: 8.0, year: 2023, status: '在看', category: 'tv' },
  { id: 4, title: '塞尔达传说：王国之泪', subtitle: 'Zelda: TotK', type: '游戏', icon: '🎮', rating: 9.5, year: 2023, status: '已通关', category: 'game' },
])

const gradients = [
  'linear-gradient(135deg, #667eea, #764ba2)',
  'linear-gradient(135deg, #f093fb, #f5576c)',
  'linear-gradient(135deg, #4facfe, #00f2fe)',
  'linear-gradient(135deg, #43e97b, #38f9d7)',
]
const getGradient = (id) => gradients[(id - 1) % gradients.length]

const filteredItems = computed(() => {
  let list = items.value
  if (activeTab.value !== 'all') list = list.filter(i => i.category === activeTab.value)
  if (keyword.value) list = list.filter(i => i.title.includes(keyword.value))
  return list
})

const resetForm = () => { showModal.value = false }
</script>

<style scoped>
.media-card-meta { display: flex; align-items: center; gap: 8px; margin-top: 6px; }
.media-card-tag { padding: 2px 8px; border-radius: 4px; font-size: 12px; }
.media-card-rating { font-size: 13px; color: var(--card-text-color-secondary); }
.media-card-year { font-size: 13px; color: var(--card-text-color-tertiary); }
.media-card-status { margin-top: 6px; font-size: 12px; font-weight: 500; }
</style>
