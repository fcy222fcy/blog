<template>
  <div ref="containerRef" class="lazy-image-container" :style="{ width, height }">
    <img
      v-if="loaded"
      :src="src"
      :alt="alt"
      class="lazy-image loaded"
      @load="onLoad"
      @error="onError"
    />
    <div v-else-if="error" class="lazy-image-error">
      <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round">
        <rect x="3" y="3" width="18" height="18" rx="2" ry="2"></rect>
        <circle cx="8.5" cy="8.5" r="1.5"></circle>
        <polyline points="21 15 16 10 5 21"></polyline>
      </svg>
    </div>
    <div v-else class="lazy-image-placeholder">
      <div class="lazy-image-spinner"></div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue'

const props = defineProps({
  src: {
    type: String,
    required: true
  },
  alt: {
    type: String,
    default: ''
  },
  width: {
    type: String,
    default: '100%'
  },
  height: {
    type: String,
    default: 'auto'
  }
})

const containerRef = ref(null)
const loaded = ref(false)
const error = ref(false)
let observer = null

const onLoad = () => {
  loaded.value = true
}

const onError = () => {
  error.value = true
}

onMounted(() => {
  if (!containerRef.value) return

  observer = new IntersectionObserver(
    (entries) => {
      if (entries[0].isIntersecting) {
        loaded.value = true
        observer.disconnect()
      }
    },
    { rootMargin: '200px' }
  )

  observer.observe(containerRef.value)
})

onUnmounted(() => {
  if (observer) {
    observer.disconnect()
  }
})
</script>

<style scoped>
.lazy-image-container {
  position: relative;
  overflow: hidden;
  background: var(--skeleton-bg, #f0f0f0);
  border-radius: inherit;
}

.lazy-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
  display: block;
  opacity: 0;
  transition: opacity 0.3s ease;
}

.lazy-image.loaded {
  opacity: 1;
}

.lazy-image-placeholder {
  position: absolute;
  inset: 0;
  display: flex;
  align-items: center;
  justify-content: center;
}

.lazy-image-spinner {
  width: 24px;
  height: 24px;
  border: 2px solid var(--border-color, #e5e7eb);
  border-top-color: var(--accent-color, #3b82f6);
  border-radius: 50%;
  animation: spin 1s linear infinite;
}

.lazy-image-error {
  position: absolute;
  inset: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  color: var(--text-color-tertiary, #9ca3af);
  background: var(--skeleton-bg, #f0f0f0);
}

@keyframes spin {
  to { transform: rotate(360deg); }
}
</style>
