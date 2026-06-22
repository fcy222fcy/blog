<template>
  <aside class="left-sidebar" :class="{ open: menuOpen }">
    <header>
      <figure class="site-avatar">
        <a href="#/">
          <img src="/avatar.jpg" class="site-logo" alt="Avatar">
        </a>
        <span class="avatar-badge">🤖</span>
      </figure>
      <div class="site-meta">
        <h1 class="site-name"><a href="#/">Liu Houliang</a></h1>
        <h2 class="site-description">日常落灰的个人博客，擅长面向搜索引擎编程。分享 Golang 开发、AI 和 NAS 折腾经验</h2>
      </div>
    </header>

    <ol class="menu-social">
      <li><a href="mailto:admin@liuhouliang.com" title="Email">
        <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><rect width="20" height="16" x="2" y="4" rx="2"></rect><path d="m22 7-8.97 5.7a1.94 1.94 0 0 1-2.06 0L2 7"></path></svg>
      </a></li>
      <li><a href="https://github.com/liu-houliang" title="GitHub" target="_blank">
        <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M15 22v-4a4.8 4.8 0 0 0-1-3.5c3 0 6-2 6-5.5.08-1.25-.27-2.48-1-3.5.28-1.15.28-2.35 0-3.5 0 0-1 0-3 1.5-2.64-.5-5.36-.5-8 0C6 2 5 2 5 2c-.3 1.15-.3 2.35 0 3.5A5.403 5.403 0 0 0 4 9c0 3.5 3 5.5 6 5.5-.39.49-.68 1.05-.85 1.65-.17.6-.22 1.23-.15 1.85v4"></path><path d="M9 18c-4.51 2-5-2-7-2"></path></svg>
      </a></li>
      <li><a href="https://x.com/liu_houliang" title="Twitter / X" target="_blank">
        <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M22 4s-.7 2.1-2 3.4c1.6 10-9.4 17.3-18 11.6 2.2.1 4.4-.6 6-2C3 15.5.5 9.6 3 5c2.2 2.6 5.6 4.1 9 4-.9-4.2 4-6.6 7-3.8 1.1 0 3-1.2 3-1.2z"></path></svg>
      </a></li>
      <li><a href="https://t.me/liu_houliang" title="Telegram" target="_blank">
        <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="m22 2-7 20-4-9-9-4Z"></path><path d="M22 2 11 13"></path></svg>
      </a></li>
    </ol>

    <div class="sidebar-search" @click="appStore.openSearch">
      <input type="text" placeholder="输入关键词..." readonly>
      <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="11" cy="11" r="8"></circle><line x1="21" y1="21" x2="16.65" y2="16.65"></line></svg>
    </div>

    <nav class="menu" id="main-menu">
      <li><router-link to="/" @click="$emit('close-menu')">
        <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="m3 9 9-7 9 7v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"></path><polyline points="9 22 9 12 15 12 15 22"></polyline></svg>
        <span>主页</span>
      </router-link></li>
      <li><router-link to="/archives" @click="$emit('close-menu')">
        <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H20v20H6.5a2.5 2.5 0 0 1 0-5H20"></path></svg>
        <span>文章</span>
      </router-link></li>
      <li><router-link to="/links" @click="$emit('close-menu')">
        <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M10 13a5 5 0 0 0 7.54.54l3-3a5 5 0 0 0-7.07-7.07l-1.72 1.71"></path><path d="M14 11a5 5 0 0 0-7.54-.54l-3 3a5 5 0 0 0 7.07 7.07l1.71-1.71"></path></svg>
        <span>友链</span>
      </router-link></li>
      <li><router-link to="/media" @click="$emit('close-menu')">
        <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="m16 13 5.223 3.482a.5.5 0 0 0 .777-.416V7.87a.5.5 0 0 0-.752-.432L16 10.5"></path><rect x="2" y="6" width="14" height="12" rx="2"></rect></svg>
        <span>娱乐</span>
      </router-link></li>
      <li><router-link to="/about" @click="$emit('close-menu')">
        <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M19 21v-2a4 4 0 0 0-4-4H9a4 4 0 0 0-4 4v2"></path><circle cx="12" cy="7" r="4"></circle></svg>
        <span>关于我</span>
      </router-link></li>
    </nav>

    <div class="sidebar-bottom">
      <button class="sidebar-bottom-item theme-toggle" type="button" @click="appStore.toggleScheme">
        <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M12 3a6 6 0 0 0 9 9 9 9 0 1 1-9-9Z"></path></svg>
        {{ appStore.scheme === 'dark' ? '亮色模式' : '暗色模式' }}
      </button>
    </div>
  </aside>
</template>

<script setup>
import { useAppStore } from '../../stores/app'

defineProps({ menuOpen: Boolean })
defineEmits(['close-menu'])
const appStore = useAppStore()
</script>
