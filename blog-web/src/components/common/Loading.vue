<template>
  <div class="loading-container" :class="{ 'loading-inline': inline, 'loading-overlay': overlay }">
    <div class="spinner"></div>
    <p v-if="text" class="loading-text">{{ text }}</p>
  </div>
</template>

<script setup>
defineProps({
  text: {
    type: String,
    default: '加载中...'
  },
  inline: {
    type: Boolean,
    default: false
  },
  overlay: {
    type: Boolean,
    default: false
  }
})
</script>

<style scoped>
.loading-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  min-height: 200px;
  gap: 16px;
}

.loading-inline {
  min-height: auto;
  padding: 12px;
  gap: 8px;
}

.loading-inline .spinner {
  width: 20px;
  height: 20px;
  border-width: 2px;
}

.loading-inline .loading-text {
  font-size: 12px;
}

.loading-overlay {
  position: absolute;
  inset: 0;
  background: rgba(255, 255, 255, 0.8);
  z-index: 10;
  min-height: auto;
}

[data-scheme="dark"] .loading-overlay {
  background: rgba(0, 0, 0, 0.6);
}

.spinner {
  width: 40px;
  height: 40px;
  border: 4px solid var(--border-color, #e5e7eb);
  border-top: 4px solid var(--accent-color, #3b82f6);
  border-radius: 50%;
  animation: spin 1s linear infinite;
}

.loading-text {
  margin: 0;
  color: var(--text-color-secondary, #6b7280);
  font-size: 14px;
}

@keyframes spin {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}
</style>
