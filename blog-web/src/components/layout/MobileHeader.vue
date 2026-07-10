<template>
  <header class="mobile-header">
    <a class="mobile-brand" href="#/">
      <img src="/avatar.jpg" alt="Liu Houliang">
      <span>Liu Houliang</span>
    </a>
    <nav class="mobile-actions">
      <button class="icon-button" @click="appStore.toggleScheme" title="切换主题">
        {{ appStore.scheme === 'dark' ? '☀' : '☾' }}
      </button>
      <button class="icon-button" @click="goSearch" title="搜索">
        <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="11" cy="11" r="8"></circle><line x1="21" y1="21" x2="16.65" y2="16.65"></line></svg>
      </button>
      <button class="icon-button" @click="$emit('toggle-menu')" title="menu">☰</button>
    </nav>
  </header>
</template>

<script setup>
import { useRouter } from 'vue-router'
import { useAppStore } from '../../stores/app'
defineEmits(['toggle-menu'])
const appStore = useAppStore()
const router = useRouter()
const goSearch = () => {
  router.push({ name: 'Search' })
}
</script>

<style scoped>
/* 移动端头部样式在全局main.css中定义 */
</style>
