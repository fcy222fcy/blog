import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useAppStore = defineStore('app', () => {
  const scheme = ref(localStorage.getItem('scheme') || 'light')
  const searchOpen = ref(false)

  const toggleScheme = () => {
    scheme.value = scheme.value === 'light' ? 'dark' : 'light'
    localStorage.setItem('scheme', scheme.value)
    document.documentElement.setAttribute('data-scheme', scheme.value)
  }

  const openSearch = () => {
    searchOpen.value = true
  }

  const closeSearch = () => {
    searchOpen.value = false
  }

  return {
    scheme,
    searchOpen,
    toggleScheme,
    openSearch,
    closeSearch
  }
})
