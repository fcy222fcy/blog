<template>
  <button class="back-to-top" :class="{ visible: show }" @click="scrollToTop" type="button">
    <svg class="progress-ring" viewBox="0 0 38 38">
      <circle class="progress-bg" cx="19" cy="19" r="18"></circle>
      <circle class="progress-bar" :style="{ strokeDashoffset: progressOffset }" cx="19" cy="19" r="18"></circle>
    </svg>
    <svg class="arrow" xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5">
      <polyline points="18 15 12 9 6 15"></polyline>
    </svg>
  </button>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue'

const show = ref(false)
const progressOffset = ref(113)

const handleScroll = () => {
  const scrollTop = window.scrollY
  const docHeight = document.documentElement.scrollHeight - window.innerHeight
  show.value = scrollTop > 300
  const progress = docHeight > 0 ? scrollTop / docHeight : 0
  progressOffset.value = 113 - (progress * 113)
}

const scrollToTop = () => window.scrollTo({ top: 0, behavior: 'smooth' })

onMounted(() => window.addEventListener('scroll', handleScroll))
onUnmounted(() => window.removeEventListener('scroll', handleScroll))
</script>

<style scoped>
/* 回到顶部按钮样式在全局main.css中定义 */
</style>
