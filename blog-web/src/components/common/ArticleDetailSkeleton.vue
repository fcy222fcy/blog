<template>
  <div class="skeleton-detail">
    <div class="skeleton-back"></div>
    <div class="skeleton-title"></div>
    <div class="skeleton-meta">
      <div class="skeleton-meta-item"></div>
      <div class="skeleton-meta-item"></div>
      <div class="skeleton-meta-item"></div>
    </div>
    <div class="skeleton-content">
      <div class="skeleton-line full"></div>
      <div class="skeleton-line full"></div>
      <div class="skeleton-line half"></div>
      <div class="skeleton-line full"></div>
      <div class="skeleton-line full"></div>
      <div class="skeleton-line three-quarter"></div>
      <div class="skeleton-line full"></div>
      <div class="skeleton-line half"></div>
    </div>
  </div>
</template>

<style scoped>
.skeleton-detail {
  padding: 20px 0;
}

.skeleton-back {
  width: 80px;
  height: 16px;
  border-radius: 4px;
  background: var(--skeleton-bg, #e5e7eb);
  margin-bottom: 24px;
  animation: skeleton-pulse 1.5s ease-in-out infinite;
}

.skeleton-title {
  width: 60%;
  height: 32px;
  border-radius: 4px;
  background: var(--skeleton-bg, #e5e7eb);
  margin-bottom: 16px;
  animation: skeleton-pulse 1.5s ease-in-out infinite 0.1s;
}

.skeleton-meta {
  display: flex;
  gap: 20px;
  margin-bottom: 32px;
}

.skeleton-meta-item {
  width: 100px;
  height: 14px;
  border-radius: 4px;
  background: var(--skeleton-bg, #e5e7eb);
  animation: skeleton-pulse 1.5s ease-in-out infinite 0.2s;
}

.skeleton-content {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.skeleton-line {
  height: 16px;
  border-radius: 4px;
  background: var(--skeleton-bg, #e5e7eb);
  animation: skeleton-pulse 1.5s ease-in-out infinite 0.3s;
}

.skeleton-line.full { width: 100%; }
.skeleton-line.three-quarter { width: 75%; }
.skeleton-line.half { width: 50%; }

@keyframes skeleton-pulse {
  0%, 100% { opacity: 1; }
  50% { opacity: 0.4; }
}
</style>
