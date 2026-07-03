<template>
  <div class="app-container" :data-scheme="scheme">
    <!-- 侧边栏 -->
    <aside class="sidebar" :class="{ open: menuOpen }">
      <div class="sidebar-header">
        <div class="sidebar-avatar" @click="handleAvatarClick">
          <span v-if="!userInfo.avatar">{{ userInfo.nickname?.charAt(0) || 'U' }}</span>
          <img v-else :src="userInfo.avatar" alt="头像">
          <div class="sidebar-avatar-edit">更换头像</div>
        </div>
        <input ref="avatarInput" type="file" accept="image/*" style="display: none;" @change="handleAvatarUpload">
        <div class="sidebar-username">{{ userInfo.nickname || '用户' }}</div>
        <div class="sidebar-desc">{{ userInfo.bio || '日常落灰的个人博客' }}</div>
        <button class="sidebar-edit-btn" @click="showProfileDialog = true">编辑资料</button>
        <div class="sidebar-social">
          <a v-for="link in userInfo.socialLinks" :key="link.name" :href="link.url" class="sidebar-social-item" target="_blank" :title="link.name">
            {{ link.icon || '🔗' }}
          </a>
        </div>
      </div>

      <nav class="sidebar-menu">
        <div v-for="item in menuItems" :key="item.path"
             class="menu-item" :class="{ active: currentPath === item.path }"
             @click="navigateTo(item.path)">
          <span class="menu-item-icon" v-html="item.icon"></span>
          <span class="menu-item-text">{{ item.title }}</span>
        </div>
      </nav>

      <div class="sidebar-features">
        <div class="feature-item">
          <span class="feature-item-icon"><svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"></circle><line x1="2" y1="12" x2="22" y2="12"></line><path d="M12 2a15.3 15.3 0 0 1 4 10 15.3 15.3 0 0 1-4 10 15.3 15.3 0 0 1-4-10 15.3 15.3 0 0 1 4-10z"></path></svg></span>
          <span>English</span>
        </div>
        <div class="feature-item" @click="toggleScheme">
          <span class="feature-item-icon"><svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M21 12.79A9 9 0 1 1 11.21 3 7 7 0 0 0 21 12.79z"></path></svg></span>
          <span>{{ scheme === 'dark' ? '亮色模式' : '暗色模式' }}</span>
        </div>
      </div>
    </aside>

    <!-- 主内容区 -->
    <div class="main-content">
      <!-- 顶部导航 -->
      <header class="header">
        <div class="header-left">
          <button class="mobile-menu-btn" @click="menuOpen = !menuOpen">☰</button>
          <h1 class="page-title">{{ currentPageTitle }}</h1>
        </div>
        <div class="header-right">
          <button class="header-btn" title="通知">
            <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M18 8A6 6 0 0 0 6 8c0 7-3 9-3 9h18s-3-2-3-9"></path><path d="M13.73 21a2 2 0 0 1-3.46 0"></path></svg>
          </button>
        </div>
      </header>

      <!-- 内容区域 -->
      <main class="content-area">
        <router-view />
      </main>
    </div>

    <!-- 回到顶部按钮 -->
    <BackToTop />

    <!-- 编辑个人资料弹窗 -->
    <div class="modal-overlay" :class="{ active: showProfileDialog }" @click.self="showProfileDialog = false">
      <div class="modal">
        <div class="modal-header">
          <h3 class="modal-title">编辑个人资料</h3>
          <button class="modal-close" @click="showProfileDialog = false">×</button>
        </div>
        <div class="modal-body">
          <div class="form-group">
            <label class="form-label">昵称</label>
            <input v-model="profileForm.nickname" class="form-input" placeholder="输入昵称">
          </div>
          <div class="form-group">
            <label class="form-label">个人简介</label>
            <textarea v-model="profileForm.bio" class="form-textarea" placeholder="输入个人简介..." rows="3"></textarea>
          </div>
        </div>
        <div class="modal-footer">
          <button class="btn btn-secondary" @click="showProfileDialog = false">取消</button>
          <button class="btn btn-primary" @click="saveProfile">保存</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { getUserInfo, updateUserInfo } from '../../api/auth'
import { ElMessage } from 'element-plus'
import BackToTop from '../common/BackToTop.vue'

const route = useRoute()
const router = useRouter()
const menuOpen = ref(false)
const showProfileDialog = ref(false)
const scheme = ref(localStorage.getItem('scheme') || 'light')
const avatarInput = ref(null)

const userInfo = ref({
  nickname: 'Liu Houliang',
  bio: '日常落灰的个人博客，分享 Golang、AI 和 NAS 折腾经验',
  avatar: '',
  socialLinks: []
})

const profileForm = ref({
  nickname: '',
  bio: ''
})

const menuItems = [
  { path: '/dashboard', title: '仪表盘', icon: '<svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M3 9l9-7 9 7v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"></path><polyline points="9 22 9 12 15 12 15 22"></polyline></svg>' },
  { path: '/articles', title: '文章', icon: '<svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"></path><polyline points="14 2 14 8 20 8"></polyline><line x1="16" y1="13" x2="8" y2="13"></line><line x1="16" y1="17" x2="8" y2="17"></line></svg>' },
  { path: '/categories', title: '分类', icon: '<svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M22 19a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h5l2 3h9a2 2 0 0 1 2 2z"></path></svg>' },
  { path: '/tags', title: '标签', icon: '<svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M20.59 13.41l-7.17 7.17a2 2 0 0 1-2.83 0L2 12V2h10l8.59 8.59a2 2 0 0 1 0 2.82z"></path><line x1="7" y1="7" x2="7.01" y2="7"></line></svg>' },
  { path: '/comments', title: '评论', icon: '<svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z"></path></svg>' },
  { path: '/links', title: '友链', icon: '<svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M10 13a5 5 0 0 0 7.54.54l3-3a5 5 0 0 0-7.07-7.07l-1.72 1.71"></path><path d="M14 11a5 5 0 0 0-7.54-.54l-3 3a5 5 0 0 0 7.07 7.07l1.71-1.71"></path></svg>' },
  { path: '/media', title: '媒体库', icon: '<svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect x="3" y="3" width="18" height="18" rx="2" ry="2"></rect><circle cx="8.5" cy="8.5" r="1.5"></circle><polyline points="21 15 16 10 5 21"></polyline></svg>' },
  { path: '/entertainment', title: '娱乐', icon: '<svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect x="2" y="2" width="20" height="20" rx="2.18" ry="2.18"></rect><line x1="7" y1="2" x2="7" y2="22"></line><line x1="17" y1="2" x2="17" y2="22"></line><line x1="2" y1="12" x2="22" y2="12"></line></svg>' },
  { path: '/daily-question', title: '每日一问', icon: '<svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"></circle><path d="M9.09 9a3 3 0 0 1 5.83 1c0 2-3 3-3 3"></path><line x1="12" y1="17" x2="12.01" y2="17"></line></svg>' },
  { path: '/about', title: '关于我', icon: '<svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"></path><circle cx="12" cy="7" r="4"></circle></svg>' },
]

const currentPath = computed(() => route.path)
const currentPageTitle = computed(() => {
  // 精确匹配
  const exactItem = menuItems.find(i => i.path === route.path)
  if (exactItem) return exactItem.title

  // 前缀匹配（处理子路由，如 /articles/123）
  const prefixItem = menuItems.find(i => route.path.startsWith(i.path + '/') || route.path.startsWith(i.path + '?'))
  if (prefixItem) return prefixItem.title

  return '仪表盘'
})

const navigateTo = (path) => {
  router.push(path)
  menuOpen.value = false
}

const toggleScheme = () => {
  scheme.value = scheme.value === 'light' ? 'dark' : 'light'
  localStorage.setItem('scheme', scheme.value)
}

const handleLogout = () => {
  localStorage.removeItem('token')
  router.push('/login')
}

const handleAvatarClick = () => {
  avatarInput.value?.click()
}

const handleAvatarUpload = async (event) => {
  const file = event.target.files?.[0]
  if (!file) return

  // 验证文件类型
  if (!file.type.startsWith('image/')) {
    ElMessage.error('请选择图片文件')
    return
  }

  // 验证文件大小 (2MB)
  if (file.size > 2 * 1024 * 1024) {
    ElMessage.error('图片大小不能超过 2MB')
    return
  }

  try {
    const formData = new FormData()
    formData.append('file', file)

    const token = localStorage.getItem('token')
    const response = await fetch('/api/v1/media/upload', {
      method: 'POST',
      headers: {
        'Authorization': `Bearer ${token}`
      },
      body: formData
    })

    const result = await response.json()

    if (response.ok && result.code === 0) {
      // 更新用户头像
      await updateUserInfo({ avatar: result.data.url })
      userInfo.value.avatar = result.data.url
      ElMessage.success('头像更新成功')
    } else {
      ElMessage.error(result.message || '上传失败')
    }
  } catch (error) {
    console.error('头像上传失败:', error)
    ElMessage.error('上传失败，请重试')
  }

  // 清空 input 的值
  event.target.value = ''
}

const saveProfile = async () => {
  try {
    await updateUserInfo(profileForm.value)
    userInfo.value.nickname = profileForm.value.nickname
    userInfo.value.bio = profileForm.value.bio
    showProfileDialog.value = false
    ElMessage.success('保存成功')
  } catch (e) {
    ElMessage.error('保存失败')
  }
}

onMounted(async () => {
  profileForm.value.nickname = userInfo.value.nickname
  profileForm.value.bio = userInfo.value.bio
  try {
    const res = await getUserInfo()
    if (res.data) {
      userInfo.value = { ...userInfo.value, ...res.data }
      profileForm.value.nickname = res.data.nickname || ''
      profileForm.value.bio = res.data.bio || ''
    }
  } catch (e) {}
})
</script>
