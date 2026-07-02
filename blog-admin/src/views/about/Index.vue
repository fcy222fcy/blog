<template>
  <div>
    <div class="card">
      <div class="card-header">
        <div class="card-title">关于我</div>
        <button class="btn btn-primary" @click="handleSave" :disabled="saving">
          <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M19 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h11l5 5v11a2 2 0 0 1-2 2z"></path><polyline points="17 21 17 13 7 13 7 21"></polyline><polyline points="7 3 7 8 15 8"></polyline></svg>
          <span>{{ saving ? '保存中...' : '保存' }}</span>
        </button>
      </div>
      <div class="card-body">
        <!-- 基本信息 -->
        <div class="section">
          <h3 class="section-title">基本信息</h3>
          <div class="form-group">
            <label class="form-label">页面标题</label>
            <input type="text" class="form-input" v-model="form.title" placeholder="关于我">
          </div>
          <div class="form-group">
            <label class="form-label">副标题</label>
            <input type="text" class="form-input" v-model="form.subtitle" placeholder="Go 开发者 / 独立游戏开发者">
          </div>
          <div class="form-group">
            <label class="form-label">简介</label>
            <textarea class="form-textarea" v-model="form.bio" rows="3" placeholder="一句话介绍自己..."></textarea>
          </div>
        </div>

        <!-- 技能标签 -->
        <div class="section">
          <h3 class="section-title">技能标签</h3>
          <div class="tags-input">
            <div class="tags-list">
              <span v-for="(tag, index) in skills" :key="index" class="tag-item">
                {{ tag }}
                <button class="tag-remove" @click="removeSkill(index)">×</button>
              </span>
            </div>
            <input type="text" class="form-input" v-model="newSkill" @keyup.enter="addSkill" placeholder="输入技能后按回车添加">
          </div>
        </div>

        <!-- 项目列表 -->
        <div class="section">
          <h3 class="section-title">项目列表</h3>
          <div v-for="(project, index) in projects" :key="index" class="item-card">
            <div class="item-header">
              <span class="item-index">{{ index + 1 }}</span>
              <button class="btn-icon btn-danger" @click="removeProject(index)">
                <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polyline points="3 6 5 6 21 6"></polyline><path d="M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6m3 0V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2"></path></svg>
              </button>
            </div>
            <div class="form-row">
              <div class="form-group" style="flex: 1;">
                <label class="form-label">名称</label>
                <input type="text" class="form-input" v-model="project.name" placeholder="项目名称">
              </div>
              <div class="form-group" style="flex: 1;">
                <label class="form-label">图标</label>
                <input type="text" class="form-input" v-model="project.icon" placeholder="图标 emoji 或文字">
              </div>
            </div>
            <div class="form-group">
              <label class="form-label">描述</label>
              <input type="text" class="form-input" v-model="project.description" placeholder="项目描述">
            </div>
            <div class="form-group">
              <label class="form-label">链接</label>
              <input type="text" class="form-input" v-model="project.url" placeholder="https://...">
            </div>
          </div>
          <button class="btn btn-secondary" @click="addProject">
            <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><line x1="12" y1="5" x2="12" y2="19"></line><line x1="5" y1="12" x2="19" y2="12"></line></svg>
            <span>添加项目</span>
          </button>
        </div>

        <!-- 关于我详情 -->
        <div class="section">
          <h3 class="section-title">关于我详情</h3>
          <div v-for="(item, index) in aboutMe" :key="index" class="item-card">
            <div class="item-header">
              <span class="item-index">{{ index + 1 }}</span>
              <button class="btn-icon btn-danger" @click="removeAboutMe(index)">
                <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polyline points="3 6 5 6 21 6"></polyline><path d="M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6m3 0V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2"></path></svg>
              </button>
            </div>
            <div class="form-row">
              <div class="form-group" style="flex: 1;">
                <label class="form-label">图标</label>
                <input type="text" class="form-input" v-model="item.icon" placeholder="图标 emoji">
              </div>
              <div class="form-group" style="flex: 1;">
                <label class="form-label">标签</label>
                <input type="text" class="form-input" v-model="item.label" placeholder="职业">
              </div>
              <div class="form-group" style="flex: 1;">
                <label class="form-label">值</label>
                <input type="text" class="form-input" v-model="item.value" placeholder="Go 开发者">
              </div>
            </div>
          </div>
          <button class="btn btn-secondary" @click="addAboutMe">
            <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><line x1="12" y1="5" x2="12" y2="19"></line><line x1="5" y1="12" x2="19" y2="12"></line></svg>
            <span>添加信息</span>
          </button>
        </div>

        <!-- 关于网站 -->
        <div class="section">
          <h3 class="section-title">关于网站</h3>
          <div v-for="(item, index) in aboutSite" :key="index" class="item-card">
            <div class="item-header">
              <span class="item-index">{{ index + 1 }}</span>
              <button class="btn-icon btn-danger" @click="removeAboutSite(index)">
                <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polyline points="3 6 5 6 21 6"></polyline><path d="M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6m3 0V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2"></path></svg>
              </button>
            </div>
            <div class="form-row">
              <div class="form-group" style="flex: 1;">
                <label class="form-label">图标</label>
                <input type="text" class="form-input" v-model="item.icon" placeholder="图标 emoji">
              </div>
              <div class="form-group" style="flex: 1;">
                <label class="form-label">标签</label>
                <input type="text" class="form-input" v-model="item.label" placeholder="框架">
              </div>
              <div class="form-group" style="flex: 1;">
                <label class="form-label">值</label>
                <input type="text" class="form-input" v-model="item.value" placeholder="Go + Gin">
              </div>
            </div>
          </div>
          <button class="btn btn-secondary" @click="addAboutSite">
            <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><line x1="12" y1="5" x2="12" y2="19"></line><line x1="5" y1="12" x2="19" y2="12"></line></svg>
            <span>添加信息</span>
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { getAboutPage, updateAboutPage } from '../../api/about'

const saving = ref(false)
const newSkill = ref('')

const form = ref({
  title: '关于我',
  subtitle: '',
  bio: ''
})

const skills = ref([])
const projects = ref([])
const aboutMe = ref([])
const aboutSite = ref([])

// 加载数据
onMounted(async () => {
  try {
    const res = await getAboutPage()
    if (res.data) {
      form.value.title = res.data.title || '关于我'
      form.value.subtitle = res.data.subtitle || ''
      form.value.bio = res.data.bio || ''

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
  }
})

// 技能标签
const addSkill = () => {
  if (newSkill.value.trim() && !skills.value.includes(newSkill.value.trim())) {
    skills.value.push(newSkill.value.trim())
    newSkill.value = ''
  }
}

const removeSkill = (index) => {
  skills.value.splice(index, 1)
}

// 项目
const addProject = () => {
  projects.value.push({ name: '', description: '', url: '', icon: '' })
}

const removeProject = (index) => {
  projects.value.splice(index, 1)
}

// 关于我
const addAboutMe = () => {
  aboutMe.value.push({ label: '', value: '', icon: '' })
}

const removeAboutMe = (index) => {
  aboutMe.value.splice(index, 1)
}

// 关于网站
const addAboutSite = () => {
  aboutSite.value.push({ label: '', value: '', icon: '' })
}

const removeAboutSite = (index) => {
  aboutSite.value.splice(index, 1)
}

// 保存
const handleSave = async () => {
  saving.value = true
  try {
    await updateAboutPage({
      title: form.value.title,
      subtitle: form.value.subtitle,
      bio: form.value.bio,
      skills: JSON.stringify(skills.value),
      projects: JSON.stringify(projects.value),
      about_me: JSON.stringify(aboutMe.value),
      about_site: JSON.stringify(aboutSite.value)
    })
    ElMessage.success('保存成功')
  } catch (e) {
    ElMessage.error('保存失败')
    console.error(e)
  } finally {
    saving.value = false
  }
}
</script>

<style scoped>
.section {
  margin-bottom: 32px;
  padding-bottom: 24px;
  border-bottom: 1px solid var(--card-separator-color);
}

.section:last-child {
  border-bottom: none;
  margin-bottom: 0;
  padding-bottom: 0;
}

.section-title {
  font-size: 16px;
  font-weight: 600;
  color: var(--card-text-color-main);
  margin-bottom: 16px;
}

.tags-input {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.tags-list {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
}

.tag-item {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  padding: 6px 12px;
  background: rgba(var(--accent-color-rgb), 0.1);
  color: var(--accent-color);
  border-radius: 6px;
  font-size: 14px;
}

.tag-remove {
  background: none;
  color: var(--accent-color);
  font-size: 18px;
  padding: 0;
  line-height: 1;
  opacity: 0.7;
}

.tag-remove:hover {
  opacity: 1;
}

.item-card {
  background: var(--body-background);
  border-radius: 8px;
  padding: 16px;
  margin-bottom: 12px;
}

.item-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 12px;
}

.item-index {
  width: 24px;
  height: 24px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: var(--accent-color);
  color: white;
  border-radius: 50%;
  font-size: 12px;
  font-weight: 600;
}

.btn-icon {
  padding: 6px;
  border: none;
  background: none;
  cursor: pointer;
  border-radius: 4px;
  transition: background 0.2s;
}

.btn-icon:hover {
  background: rgba(0, 0, 0, 0.05);
}

.btn-danger {
  color: var(--error-color);
}

.form-row {
  display: flex;
  gap: 16px;
}

@media (max-width: 768px) {
  .form-row {
    flex-direction: column;
  }
}
</style>
