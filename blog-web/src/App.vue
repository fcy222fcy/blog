<template>
  <div class="app" :data-scheme="scheme">
    <MobileHeader v-if="isMobile" @toggle-menu="toggleMenu" />
    <div class="site-shell" :class="{ 'menu-open': menuOpen }">
      <AppSidebar :menu-open="menuOpen" @close-menu="menuOpen = false" />
      <main class="content-area">
        <router-view />
      </main>
    </div>
    <BackToTop />
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue'
import { useAppStore } from './stores/app'
import MobileHeader from './components/layout/MobileHeader.vue'
import AppSidebar from './components/layout/AppSidebar.vue'
import BackToTop from './components/common/BackToTop.vue'

const appStore = useAppStore()
const scheme = ref(appStore.scheme)
const menuOpen = ref(false)
const isMobile = ref(false)

const toggleMenu = () => {
  menuOpen.value = !menuOpen.value
}

const checkMobile = () => {
  isMobile.value = window.innerWidth < 768
}

onMounted(() => {
  checkMobile()
  window.addEventListener('resize', checkMobile)
})

onUnmounted(() => {
  window.removeEventListener('resize', checkMobile)
})
</script>

<style>
@import './assets/styles/main.css';
</style>
