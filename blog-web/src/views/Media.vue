<template>
  <div class="page-view" id="view-media">
    <h1 class="page-heading">影视游戏</h1>

    <div v-if="loading" class="media-loading">
      <div class="skeleton-bar" style="width: 300px; height: 32px; margin-bottom: 24px;"></div>
      <div class="skeleton-bar" style="width: 100%; height: 48px; margin-bottom: 24px;"></div>
      <div class="skeleton-bar" style="width: 120px; height: 24px; margin-bottom: 16px;"></div>
      <div class="media-grid">
        <div v-for="i in 4" :key="i" class="media-item">
          <div class="skeleton-cover"></div>
          <div class="media-body">
            <div class="skeleton-bar" style="width: 80%; height: 16px; margin-bottom: 8px;"></div>
            <div class="skeleton-bar" style="width: 60%; height: 14px; margin-bottom: 8px;"></div>
            <div class="skeleton-bar" style="width: 90%; height: 36px;"></div>
          </div>
        </div>
      </div>
    </div>

    <template v-else>
      <div class="media-year-bar">
        <button class="year-btn" :class="{ active: activeYear === null }" @click="activeYear = null">全部</button>
        <button v-for="year in years" :key="year" class="year-btn"
                :class="{ active: activeYear === year }" @click="activeYear = year">{{ year }}</button>
      </div>

      <template v-if="filteredGames.length > 0">
        <h3 class="media-section-title">GAMES</h3>
        <div class="media-grid">
          <a v-for="item in filteredGames" :key="item.id" class="media-item"
             :href="item.link || '#'" :target="item.link ? '_blank' : '_self'">
            <div class="media-cover-wrap">
              <img v-if="item.cover" class="media-cover" :src="item.cover" :alt="item.title" loading="lazy">
              <div v-else class="media-cover-placeholder">🎮</div>
            </div>
            <div class="media-body">
              <p v-if="item.comment" class="media-comment">{{ item.comment }}</p>
              <h4>
                {{ item.title_en || item.title }}
                <span class="media-year">{{ item.year }}</span>
              </h4>
              <div class="media-stats">
                <span v-if="item.rating_external > 0">RTG ⭐ {{ item.rating_external }}</span>
                <span v-if="item.platform">{{ item.platform }}</span>
                <span v-if="item.rating > 0">MY ⭐ {{ item.rating }}</span>
                <span v-if="item.playtime">{{ item.playtime }}</span>
              </div>
            </div>
          </a>
        </div>
      </template>

      <template v-if="filteredMovies.length > 0">
        <h3 class="media-section-title">MOVIES</h3>
        <div class="media-grid">
          <a v-for="item in filteredMovies" :key="item.id" class="media-item"
             :href="item.link || '#'" :target="item.link ? '_blank' : '_self'">
            <div class="media-cover-wrap">
              <img v-if="item.cover" class="media-cover" :src="item.cover" :alt="item.title" loading="lazy">
              <div v-else class="media-cover-placeholder">🎬</div>
            </div>
            <div class="media-body">
              <p v-if="item.comment" class="media-comment">{{ item.comment }}</p>
              <h4>
                {{ item.title_en || item.title }}
                <span class="media-year">{{ item.year }}</span>
              </h4>
              <div class="media-stats">
                <span v-if="item.rating_external > 0">RTG ⭐ {{ item.rating_external }}</span>
                <span v-if="item.rating > 0">MY ⭐ {{ item.rating }}</span>
              </div>
            </div>
          </a>
        </div>
      </template>

      <template v-if="filteredTv.length > 0">
        <h3 class="media-section-title">TV SHOWS</h3>
        <div class="media-grid">
          <a v-for="item in filteredTv" :key="item.id" class="media-item"
             :href="item.link || '#'" :target="item.link ? '_blank' : '_self'">
            <div class="media-cover-wrap">
              <img v-if="item.cover" class="media-cover" :src="item.cover" :alt="item.title" loading="lazy">
              <div v-else class="media-cover-placeholder">📺</div>
            </div>
            <div class="media-body">
              <p v-if="item.comment" class="media-comment">{{ item.comment }}</p>
              <h4>
                {{ item.title_en || item.title }}
                <span class="media-year">{{ item.year }}</span>
              </h4>
              <div class="media-stats">
                <span v-if="item.rating_external > 0">RTG ⭐ {{ item.rating_external }}</span>
                <span v-if="item.platform">{{ item.platform }}</span>
                <span v-if="item.rating > 0">MY ⭐ {{ item.rating }}</span>
                <span v-if="item.playtime">{{ item.playtime }}</span>
              </div>
            </div>
          </a>
        </div>
      </template>

      <div v-if="!loading && allFiltered.length === 0" class="card">
        <div class="card-body empty-state-sm">暂无数据</div>
      </div>
    </template>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { getEntertainmentPublic } from '../api/entertainment'

const loading = ref(true)
const list = ref([])
const years = ref([])
const activeYear = ref(null)

const filteredList = computed(() => {
  if (activeYear.value === null) return list.value
  return list.value.filter(i => i.year === activeYear.value)
})

const allFiltered = computed(() => filteredList.value)
const filteredGames = computed(() => filteredList.value.filter(i => i.type === 'game'))
const filteredMovies = computed(() => filteredList.value.filter(i => i.type === 'movie'))
const filteredTv = computed(() => filteredList.value.filter(i => i.type === 'tv'))

const loadData = async () => {
  loading.value = true
  try {
    const res = await getEntertainmentPublic()
    list.value = res.data?.list || []
    years.value = res.data?.years || []
  } catch (e) {
    console.error(e)
  } finally {
    loading.value = false
  }
}

onMounted(loadData)
</script>

<style scoped>
.media-loading {
  padding: 20px 0;
}
.skeleton-bar {
  background: var(--card-background-selected);
  border-radius: 8px;
  animation: skeleton-pulse 1.5s ease-in-out infinite;
}
.skeleton-cover {
  width: 100%;
  aspect-ratio: 3 / 4;
  border-radius: 10px 10px 0 0;
  background: var(--card-background-selected);
  animation: skeleton-pulse 1.5s ease-in-out infinite;
}
@keyframes skeleton-pulse {
  0%, 100% { opacity: 0.5; }
  50% { opacity: 1; }
}
</style>
