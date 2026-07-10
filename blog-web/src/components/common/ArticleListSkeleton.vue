<template>
  <div class="skeleton-list">
    <div v-for="n in count" :key="n" class="skeleton-card">
      <div class="skeleton-category"></div>
      <div class="skeleton-title"></div>
      <div class="skeleton-summary">
        <div class="skeleton-line full"></div>
        <div class="skeleton-line full"></div>
        <div class="skeleton-line half"></div>
      </div>
      <div class="skeleton-meta">
        <div class="skeleton-meta-item"></div>
        <div class="skeleton-meta-item"></div>
        <div class="skeleton-meta-item"></div>
        <div class="skeleton-meta-item"></div>
      </div>
    </div>
  </div>
</template>

<script setup>
defineProps({
  count: {
    type: Number,
    default: 3
  }
})
</script>

<style scoped>
.skeleton-list {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.skeleton-card {
  background: var(--card-background);
  border-radius: var(--card-border-radius);
  padding: 24px;
  box-shadow: var(--shadow-l1);
}

.skeleton-category {
  width: 60px;
  height: 20px;
  border-radius: 10px;
  background: var(--card-separator-color);
  margin-bottom: 12px;
  animation: skeleton-pulse 1.5s ease-in-out infinite;
}

.skeleton-title {
  width: 70%;
  height: 24px;
  border-radius: 4px;
  background: var(--card-separator-color);
  margin-bottom: 16px;
  animation: skeleton-pulse 1.5s ease-in-out infinite 0.1s;
}

.skeleton-summary {
  display: flex;
  flex-direction: column;
  gap: 8px;
  margin-bottom: 16px;
}

.skeleton-line {
  height: 14px;
  border-radius: 4px;
  background: var(--card-separator-color);
  animation: skeleton-pulse 1.5s ease-in-out infinite 0.2s;
}

.skeleton-line.full {
  width: 100%;
}

.skeleton-line.half {
  width: 60%;
}

.skeleton-meta {
  display: flex;
  gap: 20px;
}

.skeleton-meta-item {
  width: 80px;
  height: 14px;
  border-radius: 4px;
  background: var(--card-separator-color);
  animation: skeleton-pulse 1.5s ease-in-out infinite 0.3s;
}

@keyframes skeleton-pulse {
  0%, 100% { opacity: 1; }
  50% { opacity: 0.4; }
}
</style>
