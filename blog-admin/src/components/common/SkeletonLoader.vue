<template>
  <!-- 统计卡片骨架屏 -->
  <div v-if="type === 'stats'" class="skeleton-stats-grid">
    <div v-for="i in count" :key="i" class="skeleton-stat-card">
      <div class="skeleton skeleton-label"></div>
      <div class="skeleton skeleton-value"></div>
      <div class="skeleton skeleton-sub"></div>
    </div>
  </div>

  <!-- 文章卡片骨架屏 -->
  <div v-else-if="type === 'article'" class="skeleton-article-list">
    <div v-for="i in count" :key="i" class="skeleton-article-card">
      <div class="skeleton skeleton-cover"></div>
      <div class="skeleton-article-info">
        <div class="skeleton skeleton-title"></div>
        <div class="skeleton skeleton-desc"></div>
        <div class="skeleton skeleton-desc-short"></div>
        <div class="skeleton skeleton-meta"></div>
      </div>
    </div>
  </div>

  <!-- 评论列表骨架屏 -->
  <div v-else-if="type === 'comment'" class="skeleton-comment-list">
    <div v-for="i in count" :key="i" class="skeleton-comment-item">
      <div class="skeleton skeleton-avatar"></div>
      <div class="skeleton-comment-body">
        <div class="skeleton skeleton-line" style="width:30%"></div>
        <div class="skeleton skeleton-line-short"></div>
      </div>
      <div class="skeleton skeleton-action"></div>
    </div>
  </div>

  <!-- 表格行骨架屏 -->
  <div v-else-if="type === 'table'" class="skeleton-table">
    <div v-for="i in count" :key="i" class="skeleton-table-row">
      <div class="skeleton skeleton-cell" :style="{ width: getCellWidth(i, 0) }"></div>
      <div class="skeleton skeleton-cell" :style="{ width: getCellWidth(i, 1) }"></div>
      <div class="skeleton skeleton-cell" :style="{ width: getCellWidth(i, 2) }"></div>
      <div class="skeleton skeleton-cell" :style="{ width: getCellWidth(i, 3) }"></div>
    </div>
  </div>

  <!-- 卡片网格骨架屏 -->
  <div v-else-if="type === 'card-grid'" class="skeleton-card-grid">
    <div v-for="i in count" :key="i" class="skeleton-grid-card">
      <div class="skeleton skeleton-grid-cover"></div>
      <div class="skeleton-grid-info">
        <div class="skeleton skeleton-line" style="width:60%"></div>
        <div class="skeleton skeleton-line" style="width:80%"></div>
        <div class="skeleton skeleton-line" style="width:40%"></div>
      </div>
    </div>
  </div>

  <!-- 媒体卡片网格骨架屏 -->
  <div v-else-if="type === 'media'" class="skeleton-media-grid">
    <div v-for="i in count" :key="i" class="skeleton-media-card">
      <div class="skeleton skeleton-media-cover"></div>
      <div class="skeleton-media-info">
        <div class="skeleton skeleton-line" style="width:70%"></div>
        <div class="skeleton skeleton-line" style="width:40%"></div>
        <div class="skeleton skeleton-line" style="width:50%"></div>
      </div>
    </div>
  </div>
</template>

<script setup>
defineProps({
  type: {
    type: String,
    required: true,
    validator: (v) => ['stats', 'article', 'comment', 'table', 'card-grid', 'media'].includes(v)
  },
  count: {
    type: Number,
    default: 4
  }
})

const cellWidths = [
  ['40px', '30%', '18%', '24%'],
  ['55px', '28%', '22%', '20%'],
  ['48px', '35%', '15%', '26%'],
  ['60px', '25%', '20%', '22%']
]
const getCellWidth = (rowIndex, colIndex) => {
  return cellWidths[(rowIndex - 1) % cellWidths.length][colIndex] || '20%'
}
</script>

<style scoped>
/* 骨架屏核心动画 */
.skeleton {
  background: linear-gradient(90deg,
    var(--skeleton-base) 25%,
    var(--skeleton-shine) 50%,
    var(--skeleton-base) 75%
  );
  background-size: 800px 100%;
  animation: shimmer 1.5s infinite ease-in-out;
  border-radius: 4px;
}

[data-scheme="dark"] .skeleton {
  --skeleton-base: #2d333b;
  --skeleton-shine: #373e47;
}

@keyframes shimmer {
  0% { background-position: -400px 0; }
  100% { background-position: 400px 0; }
}

/* 统计卡片 */
.skeleton-stats-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(220px, 1fr));
  gap: 24px;
  margin-bottom: 28px;
}
.skeleton-stat-card {
  background: var(--card-background);
  border: 1px solid var(--card-border);
  border-radius: var(--card-border-radius);
  padding: 24px 28px;
  box-shadow: var(--shadow-l1);
}
.skeleton-label { width: 60px; height: 14px; margin-bottom: 12px; }
.skeleton-value { width: 80px; height: 32px; margin-bottom: 8px; }
.skeleton-sub { width: 100px; height: 12px; }

/* 文章卡片 */
.skeleton-article-list {
  display: flex;
  flex-direction: column;
}
.skeleton-article-card {
  display: flex;
  gap: 16px;
  padding: 20px 28px;
  border-bottom: 1px solid var(--card-separator-color);
}
.skeleton-cover {
  width: 140px;
  height: 90px;
  border-radius: 8px;
  flex-shrink: 0;
}
.skeleton-article-info {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 10px;
  padding: 4px 0;
}
.skeleton-title { width: 70%; height: 20px; }
.skeleton-desc { width: 85%; height: 14px; }
.skeleton-desc-short { width: 55%; height: 14px; }
.skeleton-meta { width: 35%; height: 12px; margin-top: auto; }

/* 评论列表 */
.skeleton-comment-list {
  display: flex;
  flex-direction: column;
}
.skeleton-comment-item {
  display: flex;
  align-items: center;
  gap: 14px;
  padding: 16px 20px;
  border-bottom: 1px solid var(--card-separator-color);
}
.skeleton-avatar {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  flex-shrink: 0;
}
.skeleton-comment-body {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 8px;
}
.skeleton-line { height: 14px; }
.skeleton-line-short { height: 14px; width: 50%; }
.skeleton-action { width: 80px; height: 28px; border-radius: 6px; flex-shrink: 0; }

/* 表格 */
.skeleton-table {
  display: flex;
  flex-direction: column;
}
.skeleton-table-row {
  display: flex;
  gap: 24px;
  padding: 14px 28px;
  border-bottom: 1px solid var(--card-separator-color);
}
.skeleton-table-row:first-child {
  border-bottom-width: 2px;
}
.skeleton-cell { height: 16px; }

/* 卡片网格 */
.skeleton-card-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
  gap: 16px;
}
.skeleton-grid-card {
  background: var(--card-background);
  border: 1px solid var(--card-border);
  border-radius: var(--card-border-radius);
  box-shadow: var(--shadow-l1);
  overflow: hidden;
}
.skeleton-grid-cover {
  height: 120px;
}
.skeleton-grid-info {
  padding: 16px 18px;
  display: flex;
  flex-direction: column;
  gap: 8px;
}

/* 媒体卡片 */
.skeleton-media-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
  gap: 16px;
}
.skeleton-media-card {
  display: flex;
  background: var(--card-background);
  border: 1px solid var(--card-border);
  border-radius: var(--card-border-radius);
  box-shadow: var(--shadow-l1);
  overflow: hidden;
}
.skeleton-media-cover {
  width: 80px;
  min-height: 120px;
  flex-shrink: 0;
}
.skeleton-media-info {
  flex: 1;
  padding: 14px 16px;
  display: flex;
  flex-direction: column;
  gap: 8px;
}
</style>
