<template>
  <div class="page-view">
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
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { getAboutPage } from '../api/about'

const aboutData = ref({})
const skills = ref([])
const projects = ref([])
const aboutMe = ref([])
const aboutSite = ref([])

onMounted(async () => {
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
    console.error(e)
    // 使用默认数据
    aboutData.value = {
      title: 'Liu Houliang',
      subtitle: 'Go 开发者 / 独立游戏开发者',
      bio: '来自中国的程序员，擅长游戏和后端开发，喜欢玩游戏和骑自行车。'
    }
    skills.value = ['Golang', 'Erlang', 'Unity', 'Docker']
    projects.value = [
      { name: 'GitHub', description: '我的 GitHub 主页，包含各种编程项目。', url: 'https://github.com/liu-houliang', icon: '⌘' },
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
  }
})
</script>

<style scoped>
/* 关于我页面样式在全局main.css中定义 */
</style>
