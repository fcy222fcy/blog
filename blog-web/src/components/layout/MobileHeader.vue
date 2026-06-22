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
      <button class="icon-button" @click="appStore.openSearch" title="search">⌕</button>
      <button class="icon-button" @click="$emit('toggle-menu')" title="menu">☰</button>
    </nav>
  </header>
</template>

<script setup>
import { useAppStore } from '../../stores/app'
defineEmits(['toggle-menu'])
const appStore = useAppStore()
</script>

<style scoped>
/* 移动端头部样式在全局main.css中定义 */
</style>
