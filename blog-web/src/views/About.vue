<template>
  <div class="page-view">
    <!-- 加载状态 -->
    <Loading v-if="loading" text="加载关于页面..." />

    <!-- 错误状态 -->
    <div v-else-if="error" class="error-container">
      <p class="error-message">{{ error }}</p>
      <button class="retry-btn" @click="fetchAbout">
        <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M21.5 2v6h-6M2.5 22v-6h6M2 11.5a10 10 0 0 1 18.8-4.3M22 12.5a10 10 0 0 1-18.8 4.2"/></svg>
        重试
      </button>
    </div>

    <!-- 内容 -->
    <template v-else>
      <div class="about-hero">
        <span class="about-wave">👋</span>
        <span>你好，我是</span>
        <h1>{{ aboutData.title || '关于我' }}</h1>
        <p class="about-role">{{ aboutData.subtitle }}</p>
        <p class="about-bio">{{ aboutData.bio }}</p>
        <div class="skill-tags">
          <span v-for="skill in skills" :key="skill" class="tag">{{ skill }}</span>
        </div>
      </div>

      <h2>🚀 项目 & 链接</h2>
      <div class="project-grid">
        <a v-for="project in projects" :key="project.name" :href="project.url" class="project-card" target="_blank">
          <div class="project-icon-box">{{ project.icon }}</div>
          <div><h3>{{ project.name }}</h3><p>{{ project.description }}</p></div>
        </a>
      </div>

      <div class="about-two-col">
        <div class="about-box">
          <h2>👨 关于我</h2>
          <dl class="about-dl">
            <div v-for="item in aboutMe" :key="item.label">
              <dt>{{ item.icon }} {{ item.label }}:</dt>
              <dd>{{ item.value }}</dd>
            </div>
          </dl>
        </div>
        <div class="about-box">
          <h2>🏠 关于网站</h2>
          <dl class="about-dl">
            <div v-for="item in aboutSite" :key="item.label">
              <dt>{{ item.icon }} {{ item.label }}:</dt>
              <dd>{{ item.value }}</dd>
            </div>
          </dl>
        </div>
      </div>
    </template>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { getAboutPage } from '../api/about'
import { handleError } from '../utils/errorHandler'
import Loading from '../components/common/Loading.vue'

const loading = ref(false)
const error = ref(null)
const aboutData = ref({})
const skills = ref([])
const projects = ref([])
const aboutMe = ref([])
const aboutSite = ref([])

const fetchAbout = async () => {
  loading.value = true
  error.value = null
  try {
    const res = await getAboutPage()
    if (res.data) {
      aboutData.value = res.data

      // 解析 JSON 字符串
      try {
        skills.value = JSON.parse(res.data.skills || '[]')
      } catch (e) { skills.value = [] }

      try {
        projects.value = JSON.parse(res.data.projects || '[]')
      } catch (e) { projects.value = [] }

      try {
        aboutMe.value = JSON.parse(res.data.about_me || '[]')
      } catch (e) { aboutMe.value = [] }

      try {
        aboutSite.value = JSON.parse(res.data.about_site || '[]')
      } catch (e) { aboutSite.value = [] }
    }
  } catch (e) {
    error.value = handleError(e, { showMessage: false })
    // 使用默认数据
    aboutData.value = {
      title: 'Liu Houliang',
      subtitle: 'Go 开发者 / 独立游戏开发者',
      bio: '来自中国的程序员，擅长游戏和后端开发，喜欢玩游戏和骑自行车。'
    }
    skills.value = ['Golang', 'Erlang', 'Unity', 'Docker']
    projects.value = [
      { name: 'GitHub', description: '我的 GitHub 主页，包含各种编程项目。', url: 'https://github.com/fcy222fcy?tab=repositories', icon: '⌘' },
      { name: 'DesktopSnap', description: 'Windows 桌面图标保存和恢复工具。', url: 'https://desktopsnap.liuhouliang.com/', icon: '🖥' }
    ]
    aboutMe.value = [
      { label: '职业', value: 'Go 开发者', icon: '💼' },
      { label: '爱好', value: '游戏 / 骑行', icon: '🎮' },
      { label: '技术栈', value: 'Go / Erlang / Unity', icon: '🛠' },
      { label: '邮箱', value: 'admin@liuhouliang.com', icon: '✉️' }
    ]
    aboutSite.value = [
      { label: '框架', value: 'Go + Gin', icon: '⚙️' },
      { label: '前端', value: 'Vue 3', icon: '🎨' },
      { label: '部署', value: 'Docker', icon: '☁️' }
    ]
  } finally {
    loading.value = false
  }
}

onMounted(fetchAbout)
</script>

<style scoped>
/* 关于我页面样式在全局main.css中定义 */
</style>
