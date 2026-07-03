<template>
  <aside class="left-sidebar" :class="{ open: menuOpen }">
    <header>
      <figure class="site-avatar">
        <a href="#/">
          <img :src="userInfo.avatar ? userInfo.avatar + '?t=' + Date.now() : '/avatar.jpg'" class="site-logo" alt="Avatar">
        </a>
        <span class="avatar-badge">🤖</span>
      </figure>
      <div class="site-meta">
        <h1 class="site-name"><a href="#/">{{ userInfo.nickname || 'Liu Houliang' }}</a></h1>
        <h2 class="site-description">{{ userInfo.bio || '日常落灰的个人博客，擅长面向搜索引擎编程。分享 Golang 开发、AI 和 NAS 折腾经验' }}</h2>
      </div>
    </header>

    <ol class="menu-social">
      <li><a href="mailto:admin@liuhouliang.com" title="Email">
        <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><rect width="20" height="16" x="2" y="4" rx="2"></rect><path d="m22 7-8.97 5.7a1.94 1.94 0 0 1-2.06 0L2 7"></path></svg>
      </a></li>
      <li><a href="https://github.com/liu-houliang" title="GitHub" target="_blank">
        <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M15 22v-4a4.8 4.8 0 0 0-1-3.5c3 0 6-2 6-5.5.08-1.25-.27-2.48-1-3.5.28-1.15.28-2.35 0-3.5 0 0-1 0-3 1.5-2.64-.5-5.36-.5-8 0C6 2 5 2 5 2c-.3 1.15-.3 2.35 0 3.5A5.403 5.403 0 0 0 4 9c0 3.5 3 5.5 6 5.5-.39.49-.68 1.05-.85 1.65-.17.6-.22 1.23-.15 1.85v4"></path><path d="M9 18c-4.51 2-5-2-7-2"></path></svg>
      </a></li>
      <li><a href="https://space.bilibili.com/" title="Bilibili" target="_blank">
        <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="currentColor"><path d="M17.813 4.653h.854c1.51.054 2.769.578 3.773 1.574 1.004.995 1.524 2.249 1.56 3.76v7.36c-.036 1.51-.556 2.769-1.56 3.773s-2.262 1.524-3.773 1.56H5.333c-1.51-.036-2.769-.556-3.773-1.56S.036 18.858 0 17.347v-7.36c.036-1.511.556-2.765 1.56-3.76 1.004-.996 2.262-1.52 3.773-1.574h.774l-1.174-1.12a1.234 1.234 0 0 1-.373-.906c0-.356.124-.658.373-.907l.027-.027c.267-.249.573-.373.92-.373.347 0 .653.124.92.373L9.653 4.44c.071.071.134.142.187.213h4.267a.836.836 0 0 1 .16-.213l2.853-2.747c.267-.249.573-.373.92-.373.347 0 .662.124.929.373.249.249.373.551.373.907 0 .355-.124.657-.373.906zM5.333 7.24c-.746.018-1.373.276-1.88.773-.506.498-.769 1.13-.786 1.894v7.52c.017.764.28 1.395.786 1.893.507.498 1.134.756 1.88.773h13.334c.746-.017 1.373-.275 1.88-.773.506-.498.769-1.129.786-1.893v-7.52c-.017-.765-.28-1.396-.786-1.894-.507-.497-1.134-.755-1.88-.773zM8 11.107c.373 0 .684.124.933.373.25.249.383.569.4.96v1.173c-.017.391-.15.711-.4.96-.249.25-.56.374-.933.374s-.684-.125-.933-.374c-.25-.249-.383-.569-.4-.96V12.44c.017-.391.15-.711.4-.96.249-.249.56-.373.933-.373zm8 0c.373 0 .684.124.933.373.25.249.383.569.4.96v1.173c-.017.391-.15.711-.4.96-.249.25-.56.374-.933.374s-.684-.125-.933-.374c-.25-.249-.383-.569-.4-.96V12.44c.017-.391.15-.711.4-.96.249-.249.56-.373.933-.373z"/></svg>
      </a></li>
      <li><a href="#" title="WeChat">
        <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="currentColor"><path d="M8.691 2.188C3.891 2.188 0 5.476 0 9.53c0 2.212 1.17 4.203 3.002 5.55a.59.59 0 0 1 .213.665l-.39 1.48c-.019.07-.048.141-.048.213 0 .163.13.295.29.295a.326.326 0 0 0 .167-.054l1.903-1.114a.864.864 0 0 1 .717-.098 10.16 10.16 0 0 0 2.837.403c.276 0 .543-.027.811-.05-.857-2.578.157-4.972 1.932-6.446 1.703-1.415 3.882-1.98 5.853-1.838-.576-3.583-4.196-6.348-8.596-6.348zM5.785 5.991c.642 0 1.162.529 1.162 1.18a1.17 1.17 0 0 1-1.162 1.178A1.17 1.17 0 0 1 4.623 7.17c0-.651.52-1.18 1.162-1.18zm5.813 0c.642 0 1.162.529 1.162 1.18a1.17 1.17 0 0 1-1.162 1.178 1.17 1.17 0 0 1-1.162-1.178c0-.651.52-1.18 1.162-1.18zm3.636 2.885c-1.577-.072-3.182.395-4.44 1.44-1.47 1.22-2.174 3.038-2.013 4.865.148 1.67 1.07 3.164 2.558 4.09a8.8 8.8 0 0 0 2.37.963.69.69 0 0 1 .57.078l1.524.892a.262.262 0 0 0 .134.043.236.236 0 0 0 .233-.236c0-.058-.023-.114-.039-.17l-.313-1.185a.472.472 0 0 1 .17-.532c1.47-1.08 2.412-2.657 2.412-4.44 0-1.352-.535-2.628-1.512-3.634-.976-1.006-2.282-1.676-3.69-1.676zm-1.774 3.255c.515 0 .932.424.932.946a.94.94 0 0 1-.932.946.94.94 0 0 1-.931-.946c0-.522.416-.946.931-.946zm4.663 0c.515 0 .931.424.931.946a.94.94 0 0 1-.931.946.94.94 0 0 1-.932-.946c0-.522.417-.946.932-.946z"/></svg>
      </a></li>
      <li><a href="/api/v1/feed.xml" title="RSS 订阅" target="_blank">
        <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M4 11a9 9 0 0 1 9 9"></path><path d="M4 4a16 16 0 0 1 16 16"></path><circle cx="5" cy="19" r="1"></circle></svg>
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
import { ref, onMounted } from 'vue'
import { useAppStore } from '../../stores/app'

defineProps({ menuOpen: Boolean })
defineEmits(['close-menu'])
const appStore = useAppStore()

const userInfo = ref({
  nickname: '',
  bio: '',
  avatar: ''
})

const fetchUserInfo = async () => {
  try {
    const response = await fetch('/api/v1/user/info')
    const result = await response.json()
    if (response.ok && result.code === 0 && result.data) {
      userInfo.value = result.data
    }
  } catch (error) {
    console.log('使用默认用户信息')
  }
}

onMounted(() => {
  fetchUserInfo()
})
</script>
