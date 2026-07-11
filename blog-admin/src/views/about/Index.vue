<template>
  <div>
    <div class="card">
      <div class="card-header">
        <div class="card-title">关于我</div>
        <div class="save-status">
          <span v-if="saveStatus === 'saving'" class="status-saving">
            <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="spin"><path d="M21 12a9 9 0 1 1-6.219-8.56"></path></svg>
            保存中...
          </span>
          <span v-else-if="saveStatus === 'saved'" class="status-saved">
            <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polyline points="20 6 9 17 4 12"></polyline></svg>
            已保存
          </span>
          <span v-else class="status-idle">
            <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"></circle><line x1="12" y1="8" x2="12" y2="12"></line><line x1="12" y1="16" x2="12.01" y2="16"></line></svg>
            编辑即保存
          </span>
        </div>
      </div>
      <div class="card-body">
        <!-- 基本信息 -->
        <div class="section">
          <h3 class="section-title">基本信息</h3>
          <div class="form-group">
            <label class="form-label">页面标题</label>
            <input type="text" class="form-input" v-model="form.title" placeholder="关于我" @blur="triggerSave">
          </div>
          <div class="form-group">
            <label class="form-label">副标题</label>
            <input type="text" class="form-input" v-model="form.subtitle" placeholder="Go 开发者 / 独立游戏开发者" @blur="triggerSave">
          </div>
          <div class="form-group">
            <label class="form-label">简介</label>
            <textarea class="form-textarea" v-model="form.bio" rows="3" placeholder="一句话介绍自己..." @blur="triggerSave"></textarea>
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
            <input type="text" class="form-input" v-model="newSkill" @keyup.enter="addSkill" @blur="triggerSave" placeholder="输入技能后按回车添加">
          </div>
        </div>

        <!-- 项目列表 -->
        <div class="section">
          <h3 class="section-title">项目列表</h3>
          <div class="item-grid">
            <div v-for="(project, index) in projects" :key="index" class="item-card">
              <div class="item-header">
                <span class="item-index">{{ index + 1 }}</span>
                <button class="btn-icon btn-danger" @click="removeProject(index)">
                  <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polyline points="3 6 5 6 21 6"></polyline><path d="M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6m3 0V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2"></path></svg>
                </button>
              </div>
              <div class="form-group">
                <label class="form-label">图标</label>
                <div class="icon-picker-wrapper">
                  <div class="icon-display" @click="openIconPicker('project', index)">
                    <component :is="renderIcon(project.icon)" class="icon-svg" v-if="project.icon" />
                    <span class="icon-placeholder" v-else>选择图标</span>
                    <svg width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="icon-caret"><polyline points="6 9 12 15 18 9"></polyline></svg>
                  </div>
                </div>
              </div>
              <div class="form-group">
                <label class="form-label">名称</label>
                <input type="text" class="form-input" v-model="project.name" placeholder="项目名称" @blur="triggerSave">
              </div>
              <div class="form-group">
                <label class="form-label">描述</label>
                <input type="text" class="form-input" v-model="project.description" placeholder="项目描述" @blur="triggerSave">
              </div>
              <div class="form-group">
                <label class="form-label">链接</label>
                <input type="text" class="form-input" v-model="project.url" placeholder="https://..." @blur="triggerSave">
              </div>
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
          <div class="item-grid">
            <div v-for="(item, index) in aboutMe" :key="index" class="item-card">
              <div class="item-header">
                <span class="item-index">{{ index + 1 }}</span>
                <button class="btn-icon btn-danger" @click="removeAboutMe(index)">
                  <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polyline points="3 6 5 6 21 6"></polyline><path d="M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6m3 0V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2"></path></svg>
                </button>
              </div>
              <div class="form-group">
                <label class="form-label">图标</label>
                <div class="icon-picker-wrapper">
                  <div class="icon-display" @click="openIconPicker('aboutMe', index)">
                    <component :is="renderIcon(item.icon)" class="icon-svg" v-if="item.icon" />
                    <span class="icon-placeholder" v-else>选择图标</span>
                    <svg width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="icon-caret"><polyline points="6 9 12 15 18 9"></polyline></svg>
                  </div>
                </div>
              </div>
              <div class="form-group">
                <label class="form-label">标签</label>
                <input type="text" class="form-input" v-model="item.label" placeholder="职业" @blur="triggerSave">
              </div>
              <div class="form-group">
                <label class="form-label">值</label>
                <input type="text" class="form-input" v-model="item.value" placeholder="Go 开发者" @blur="triggerSave">
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
          <div class="item-grid">
            <div v-for="(item, index) in aboutSite" :key="index" class="item-card">
              <div class="item-header">
                <span class="item-index">{{ index + 1 }}</span>
                <button class="btn-icon btn-danger" @click="removeAboutSite(index)">
                  <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polyline points="3 6 5 6 21 6"></polyline><path d="M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6m3 0V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2"></path></svg>
                </button>
              </div>
              <div class="form-group">
                <label class="form-label">图标</label>
                <div class="icon-picker-wrapper">
                  <div class="icon-display" @click="openIconPicker('aboutSite', index)">
                    <component :is="renderIcon(item.icon)" class="icon-svg" v-if="item.icon" />
                    <span class="icon-placeholder" v-else>选择图标</span>
                    <svg width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="icon-caret"><polyline points="6 9 12 15 18 9"></polyline></svg>
                  </div>
                </div>
              </div>
              <div class="form-group">
                <label class="form-label">标签</label>
                <input type="text" class="form-input" v-model="item.label" placeholder="框架" @blur="triggerSave">
              </div>
              <div class="form-group">
                <label class="form-label">值</label>
                <input type="text" class="form-input" v-model="item.value" placeholder="Go + Gin" @blur="triggerSave">
              </div>
            </div>
          </div>
          <button class="btn btn-secondary" @click="addAboutSite">
            <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><line x1="12" y1="5" x2="12" y2="19"></line><line x1="5" y1="12" x2="19" y2="12"></line></svg>
            <span>添加信息</span>
          </button>
        </div>

        <!-- 建站历程 -->
        <div class="section">
          <h3 class="section-title">建站历程</h3>
          <p class="section-desc">添加历程后将自动按日期从早到晚排序，前台以时间线形式展示。</p>
          <div class="timeline-list">
            <div v-for="(item, index) in siteHistorySorted" :key="index" class="timeline-item-card">
              <div class="item-header">
                <div class="timeline-item-date">{{ item.date || '未设置日期' }}</div>
                <button class="btn-icon btn-danger" @click="removeSiteHistory(getOriginalIndex(item))">
                  <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polyline points="3 6 5 6 21 6"></polyline><path d="M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6m3 0V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2"></path></svg>
                </button>
              </div>
              <div class="form-row">
                <div class="form-group" style="flex: 0 0 180px;">
                  <label class="form-label">日期</label>
                  <input type="date" class="form-input" v-model="siteHistory[getOriginalIndex(item)].date" @blur="triggerSave">
                </div>
                <div class="form-group" style="flex: 1;">
                  <label class="form-label">内容</label>
                  <input type="text" class="form-input" v-model="siteHistory[getOriginalIndex(item)].content" placeholder="例如：网站正式上线" @blur="triggerSave">
                </div>
              </div>
            </div>
            <div v-if="siteHistory.length === 0" class="empty-tip">
              暂无建站历程，点击下方按钮添加第一条吧。
            </div>
          </div>
          <button class="btn btn-secondary" @click="addSiteHistory" style="margin-top: 12px;">
            <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><line x1="12" y1="5" x2="12" y2="19"></line><line x1="5" y1="12" x2="19" y2="12"></line></svg>
            <span>添加历程</span>
          </button>
        </div>
      </div>
    </div>

    <!-- 图标选择弹层 -->
    <div v-if="iconPicker.visible" class="icon-picker-mask" @click.self="closeIconPicker">
      <div class="icon-picker-panel">
        <div class="icon-picker-header">
          <span>选择图标</span>
          <button class="btn-icon" @click="closeIconPicker">
            <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><line x1="18" y1="6" x2="6" y2="18"></line><line x1="6" y1="6" x2="18" y2="18"></line></svg>
          </button>
        </div>
        <div class="icon-picker-body">
          <div
            v-for="ic in iconList"
            :key="ic.key"
            class="icon-option"
            :class="{ active: isPickerTargetIcon(ic.key) }"
            @click="selectIcon(ic.key)"
            :title="ic.name"
          >
            <component :is="renderIcon(ic.key)" class="icon-svg-lg" />
            <span class="icon-name">{{ ic.name }}</span>
          </div>
        </div>
        <div class="icon-picker-footer">
          <button class="btn btn-secondary btn-sm" @click="clearPickerIcon">清除图标</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, h, onMounted, onBeforeUnmount, reactive, computed } from 'vue'
import { ElMessage } from 'element-plus'
import { getAboutPage, updateAboutPage } from '../../api/about'

const saveStatus = ref('idle')
const newSkill = ref('')
let saveTimer = null
let statusTimer = null

const form = ref({
  title: '关于我',
  subtitle: '',
  bio: ''
})

const skills = ref([])
const projects = ref([])
const aboutMe = ref([])
const aboutSite = ref([])
const siteHistory = ref([])

// 给每项生成唯一 id，用于定位原始索引
let uidCounter = 0
const genUid = () => `sh_${++uidCounter}_${Date.now()}`

const ensureUids = () => {
  siteHistory.value.forEach(item => {
    if (!item.__uid) item.__uid = genUid()
  })
}

// 按日期排序后的列表（用于展示）
const siteHistorySorted = computed(() => {
  ensureUids()
  return [...siteHistory.value].sort((a, b) => (a.date || '').localeCompare(b.date || ''))
})

// 根据排序后的项找到原始索引
const getOriginalIndex = (item) => siteHistory.value.findIndex(x => x.__uid === item.__uid)

// 图标选择器状态
const iconPicker = reactive({
  visible: false,
  targetType: '',
  targetIndex: -1
})

// 预设 SVG 图标库（使用 feather-icons 风格的 path）
const iconList = [
  { key: 'user', name: '用户', path: 'M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2;M12 11a4 4 0 1 0 0-8 4 4 0 0 0 0 8z' },
  { key: 'users', name: '多人', path: 'M17 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2;M9 11a4 4 0 1 0 0-8 4 4 0 0 0 0 8z;M23 21v-2a4 4 0 0 0-3-3.87;M16 3.13a4 4 0 0 1 0 7.75' },
  { key: 'briefcase', name: '工作', path: 'M20 7h-4V5a2 2 0 0 0-2-2h-4a2 2 0 0 0-2 2v2H4a2 2 0 0 0-2 2v11a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V9a2 2 0 0 0-2-2z;M16 7V5;M8 7V5' },
  { key: 'code', name: '代码', path: 'M16 18l6-6-6-6;M8 6l-6 6 6 6' },
  { key: 'terminal', name: '终端', path: 'M4 17l6-6-6-6;M12 19h8' },
  { key: 'file', name: '文件', path: 'M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z;M14 2v6h6;M16 13H8;M16 17H8;M10 9H8' },
  { key: 'folder', name: '文件夹', path: 'M22 19a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h5l2 3h9a2 2 0 0 1 2 2z' },
  { key: 'github', name: 'GitHub', path: 'M9 19c-5 1.5-5-2.5-7-3m14 6v-3.87a3.37 3.37 0 0 0-.94-2.61c3.14-.35 6.44-1.54 6.44-7A5.44 5.44 0 0 0 20 4.77 5.07 5.07 0 0 0 19.91 1S18.73.65 16 2.48a13.38 13.38 0 0 0-7 0C6.27.65 5.09 1 5.09 1A5.07 5.07 0 0 0 5 4.77a5.44 5.44 0 0 0-1.5 3.78c0 5.42 3.3 6.61 6.44 7A3.37 3.37 0 0 0 9 18.13V22' },
  { key: 'gitlab', name: 'GitLab', path: 'M22.65 14.39L12 22.13 1.35 14.39a.84.84 0 0 1-.3-.94l1.22-3.78 2.44-7.51A.42.42 0 0 1 4.82 2a.43.43 0 0 1 .78 0l2.44 7.51 8.92-.01 2.44-7.5A.43.43 0 0 1 19.45 2a.42.42 0 0 1 .77.22l2.44 7.51 1.22 3.78a.84.84 0 0 1-.3.94z;M12 12.71L5.82 2a.43.43 0 0 0-.78 0' },
  { key: 'twitter', name: 'Twitter', path: 'M23 3a10.9 10.9 0 0 1-3.14 1.53 4.48 4.48 0 0 0-7.86 3v1A10.66 10.66 0 0 1 3 4s-4 9 5 13a11.64 11.64 0 0 1-7 2c9 5 20 0 20-11.5a4.5 4.5 0 0 0-.08-.83A7.72 7.72 0 0 0 23 3z' },
  { key: 'facebook', name: 'Facebook', path: 'M18 2h-3a5 5 0 0 0-5 5v3H7v4h3v8h4v-8h3l1-4h-4V7a1 1 0 0 1 1-1h3z' },
  { key: 'instagram', name: 'Instagram', path: 'M16 11.37A4 4 0 1 1 12.63 8 4 4 0 0 1 16 11.37z;M17.5 6.5h.01;M3 8.5A5.5 5.5 0 0 1 8.5 3h7A5.5 5.5 0 0 1 21 8.5v7a5.5 5.5 0 0 1-5.5 5.5h-7A5.5 5.5 0 0 1 3 15.5z' },
  { key: 'youtube', name: 'YouTube', path: 'M22.54 6.42a2.78 2.78 0 0 0-1.94-2C18.88 4 12 4 12 4s-6.88 0-8.6.46a2.78 2.78 0 0 0-1.94 2A29 29 0 0 0 1 11.75a29 29 0 0 0 .46 5.33A2.78 2.78 0 0 0 3.4 19c1.72.46 8.6.46 8.6.46s6.88 0 8.6-.46a2.78 2.78 0 0 0 1.94-2 29 29 0 0 0 .46-5.25 29 29 0 0 0-.46-5.33z;M9.75 15.02l5.75-3.27-5.75-3.27z' },
  { key: 'linkedin', name: 'LinkedIn', path: 'M16 8a6 6 0 0 1 6 6v7h-4v-7a2 2 0 0 0-4 0v7h-4v-7a6 6 0 0 1 6-6z;M2 9h4v12H2zM4 2a2 2 0 1 1 0 4 2 2 0 0 1 0-4z' },
  { key: 'weibo', name: '微博', path: 'M10.5 18.5a7.47 7.47 0 0 1-7.26-5.3 3.2 3.2 0 0 1 .87-3 2.4 2.4 0 0 1 2.12-.53c.76.23 1.44.76 1.76 1.5a1.5 1.5 0 0 1-2.2 1.77c.13.85.55 1.65 1.2 2.25a5.3 5.3 0 0 0 3.51 1.81M17 4c2.5 1.5 4 3.8 4 6.5 0 3.5-3.3 6.5-7.5 6.5C8.5 17 5 14 5 10.5 5 8 6.5 5.5 10 3.5 12 2.5 14.5 2 17 4z;M18.5 6.5a2 2 0 0 0-1.5-1M19 10.5c1.5 0 2.5-.5 3-2M13 5.5c.5-.2 1-.4 2-.3' },
  { key: 'bilibili', name: 'B站', path: 'M17.8 3.3L15.8 5.3H21a1.5 1.5 0 0 1 1.5 1.5V18a1.5 1.5 0 0 1-1.5 1.5H3A1.5 1.5 0 0 1 1.5 18V6.8A1.5 1.5 0 0 1 3 5.3h5.2L6.2 3.3l1.3-1.3L10 4.5 12.5 2l1.3 1.3L11.8 5.3h2.7l1.7-2zm-14 4.8v9.4h18.4V8.1z;M8 12.5a1 1 0 0 1 1 1v1a1 1 0 1 1-2 0v-1a1 1 0 0 1 1-1zm8 0a1 1 0 0 1 1 1v1a1 1 0 1 1-2 0v-1a1 1 0 0 1 1-1z' },
  { key: 'zhihu', name: '知乎', path: 'M5.7 4.1h12.6c.9 0 1.6.7 1.6 1.6v12.6c0 .9-.7 1.6-1.6 1.6H5.7a1.6 1.6 0 0 1-1.6-1.6V5.7c0-.9.7-1.6 1.6-1.6zm2.4 3H6.4v1.7h2.3v3.5h1.7v-3.5h3V7.1h-3v-.8a1 1 0 0 1 1-1h2V3.5h-2a3 3 0 0 0-3 3v.6zm7.2 9.6l-2-3-2.3 1.3 1-3.3h2.7l.5-2h-3.5V9.3h4.4l-.7 2.6-2.3.5 1 2 1.6-.9 1.2 2.3z;M15 8.3h3V6.8h-3z' },
  { key: 'wechat', name: '微信', path: 'M8.7 13.9c3.1 0 5.6-2.3 5.6-5.2S11.8 3.5 8.7 3.5 3 5.8 3 8.7c0 1.2.4 2.2 1 3.1L3.3 14l2.4-1.2c1 .5 2 .8 3 1.1zM18 16.9c2.3 0 4.2-1.7 4.2-3.8S20.3 9.2 18 9.2s-4.2 1.7-4.2 3.8c0 .9.3 1.7.8 2.4l-.5 2.8 1.8-.9c.7.3 1.4.5 2.1.6zM9.2 8.2h.02;M6.8 8.2h.02;M15.8 13.8h.02;M19.8 13.8h.02' },
  { key: 'telegram', name: 'Telegram', path: 'M22 2L11 13;M22 2l-7 20-4-9-9-4 20-9z' },
  { key: 'rss', name: 'RSS', path: 'M4 11a9 9 0 0 1 9 9;M4 4a16 16 0 0 1 16 16;M5 19a1 1 0 1 0 0-2 1 1 0 0 0 0 2z' },
  { key: 'link', name: '链接', path: 'M10 13a5 5 0 0 0 7.54.54l3-3a5 5 0 0 0-7.07-7.07l-1.72 1.71;M14 11a5 5 0 0 0-7.54-.54l-3 3a5 5 0 0 0 7.07 7.07l1.71-1.71' },
  { key: 'globe', name: '地球', path: 'M12 2a10 10 0 1 0 0 20 10 10 0 0 0 0-20z;M2 12h20;M12 2a15.3 15.3 0 0 1 4 10 15.3 15.3 0 0 1-4 10 15.3 15.3 0 0 1-4-10 15.3 15.3 0 0 1 4-10z' },
  { key: 'rocket', name: '火箭', path: 'M4.5 16.5c-1.5 1.26-2 5-2 5s3.74-.5 5-2c.71-.84.7-2.13-.09-2.91a2.18 2.18 0 0 0-2.91-.09z;M12 15l-3-3a22 22 0 0 1 2-3.95A12.88 12.88 0 0 1 22 2c0 2.72-.78 7.5-6 11a22.35 22.35 0 0 1-4 2z;M9 12H4s.55-3.03 2-4c1.62-1.08 5 0 5 0;M12 15v5s3.03-.55 4-2c1.08-1.62 0-5 0-5' },
  { key: 'star', name: '星标', path: 'M12 2l3.09 6.26L22 9.27l-5 4.87 1.18 6.88L12 17.77l-6.18 3.25L7 14.14 2 9.27l6.91-1.01L12 2z' },
  { key: 'heart', name: '爱心', path: 'M20.84 4.61a5.5 5.5 0 0 0-7.78 0L12 5.67l-1.06-1.06a5.5 5.5 0 0 0-7.78 7.78l1.06 1.06L12 21.23l7.78-7.78 1.06-1.06a5.5 5.5 0 0 0 0-7.78z' },
  { key: 'thumbs-up', name: '点赞', path: 'M7 10v12;M15 5.88 14 10h5.83a2 2 0 0 1 1.92 2.56l-2.33 8A2 2 0 0 1 17.5 22H7V10l4.34-8.66A1.5 1.5 0 0 1 15 2.88z' },
  { key: 'book', name: '书籍', path: 'M4 19.5A2.5 2.5 0 0 1 6.5 17H20;M6.5 2H20v20H6.5A2.5 2.5 0 0 1 4 19.5v-15A2.5 2.5 0 0 1 6.5 2z' },
  { key: 'bookmark', name: '书签', path: 'M19 21l-7-5-7 5V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2z' },
  { key: 'feather', name: '写作', path: 'M20.24 12.24a6 6 0 0 0-8.49-8.49L5 10.5V19h8.5zM16 8 2 22;M17.5 15H9' },
  { key: 'monitor', name: '电脑', path: 'M20 3H4a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h7l-2 3h6l-2-3h7a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2z' },
  { key: 'laptop', name: '笔记本', path: 'M2 16.1A5 5 0 0 1 5.9 20H20a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2zM2 16.1l2 1.9m18-1.9-2 1.9' },
  { key: 'smartphone', name: '手机', path: 'M17 2H7a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2zM12 18h.01' },
  { key: 'server', name: '服务器', path: 'M2 4h20v8H2z;M2 12h20v8H2z;M6 8h.01;M6 16h.01' },
  { key: 'database', name: '数据库', path: 'M12 2C6.48 2 2 4.69 2 8v8c0 3.31 4.48 6 10 6s10-2.69 10-6V8c0-3.31-4.48-6-10-6z;M2 8c0 3.31 4.48 6 10 6s10-2.69 10-6;M2 16c0 3.31 4.48 6 10 6s10-2.69 10-6' },
  { key: 'cpu', name: '芯片', path: 'M9 3v2m6-2v2M9 19v2m6-2v2M5 9H3m2 6H3m18-6h-2m2 6h-2M7 19h10a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2H7a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2zM9 9h6v6H9z' },
  { key: 'hard-drive', name: '硬盘', path: 'M22 12H2;M5.45 5.11 2 12v6a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-6l-3.45-6.89A2 2 0 0 0 16.76 4H7.24a2 2 0 0 0-1.79 1.11zM6 16h.01;M10 16h.01' },
  { key: 'wifi', name: 'WiFi', path: 'M5 12.55a11 11 0 0 1 14.08 0;M1.42 9a16 16 0 0 1 21.16 0;M8.53 16.11a6 6 0 0 1 6.95 0;M12 20h.01' },
  { key: 'cloud', name: '云', path: 'M17.5 19a4.5 4.5 0 1 0-1.3-8.8A7 7 0 1 0 4 15.8' },
  { key: 'download', name: '下载', path: 'M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4;M7 10l5 5 5-5;M12 15V3' },
  { key: 'upload', name: '上传', path: 'M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4;M17 8l-5-5-5 5;M12 3v12' },
  { key: 'share', name: '分享', path: 'M4 12v8a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-8;M16 6l-4-4-4 4;M12 2v13' },
  { key: 'send', name: '发送', path: 'M22 2 11 13;M22 2l-7 20-4-9-9-4 20-7z' },
  { key: 'message', name: '消息', path: 'M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z' },
  { key: 'message-circle', name: '评论', path: 'M21 11.5a8.38 8.38 0 0 1-.9 3.8 8.5 8.5 0 0 1-7.6 4.7 8.38 8.38 0 0 1-3.8-.9L3 21l1.9-5.7a8.38 8.38 0 0 1-.9-3.8 8.5 8.5 0 0 1 4.7-7.6 8.38 8.38 0 0 1 3.8-.9h.5a8.48 8.48 0 0 1 8 8v.5z' },
  { key: 'bell', name: '通知', path: 'M18 8A6 6 0 0 0 6 8c0 7-3 9-3 9h18s-3-2-3-9;M13.73 21a2 2 0 0 1-3.46 0' },
  { key: 'coffee', name: '咖啡', path: 'M18 8h1a4 4 0 0 1 0 8h-1;M2 8h16v9a4 4 0 0 1-4 4H6a4 4 0 0 1-4-4V8z;M6 1v3M10 1v3M14 1v3' },
  { key: 'music', name: '音乐', path: 'M9 18V5l12-2v13;M9 18a3 3 0 1 1-6 0 3 3 0 0 1 6 0z;M21 16a3 3 0 1 1-6 0 3 3 0 0 1 6 0z' },
  { key: 'headphones', name: '耳机', path: 'M3 18v-6a9 9 0 0 1 18 0v6;M21 19a2 2 0 0 1-2 2h-1v-6h3zM3 19a2 2 0 0 0 2 2h1v-6H3z' },
  { key: 'video', name: '视频', path: 'M23 7l-7 5 7 5zM14 5H3a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h11a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2z' },
  { key: 'camera', name: '相机', path: 'M23 19a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h4l2-3h6l2 3h4a2 2 0 0 1 2 2z;M12 17a4 4 0 1 0 0-8 4 4 0 0 0 0 8z' },
  { key: 'image', name: '图片', path: 'M3 3h18v18H3zM3 15l5-5 4 4 3-3 6 6' },
  { key: 'gamepad', name: '游戏', path: 'M6 11h4m-2-2v4M15 12h.01M18 14h.01;M5.33 15H4a2 2 0 0 1-2-2V9a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2h-1.33a2 2 0 0 0-1.78 1.11l-.45.89a2 2 0 0 1-1.78 1H9.33a2 2 0 0 1-1.78-1l-.44-.89A2 2 0 0 0 5.33 15z' },
  { key: 'pen', name: '编辑', path: 'M12 20h9;M16.5 3.5a2.12 2.12 0 0 1 3 3L7 19l-4 1 1-4z' },
  { key: 'settings', name: '设置', path: 'M12 15a3 3 0 1 0 0-6 3 3 0 0 0 0 6z;M19.4 15a1.65 1.65 0 0 0 .33 1.82l.06.06a2 2 0 1 1-2.83 2.83l-.06-.06a1.65 1.65 0 0 0-1.82-.33 1.65 1.65 0 0 0-1 1.51V21a2 2 0 0 1-4 0v-.09A1.65 1.65 0 0 0 9 19.4a1.65 1.65 0 0 0-1.82.33l-.06.06a2 2 0 1 1-2.83-2.83l.06-.06a1.65 1.65 0 0 0 .33-1.82 1.65 1.65 0 0 0-1.51-1H3a2 2 0 0 1 0-4h.09A1.65 1.65 0 0 0 4.6 9a1.65 1.65 0 0 0-.33-1.82l-.06-.06a2 2 0 1 1 2.83-2.83l.06.06a1.65 1.65 0 0 0 1.82.33H9a1.65 1.65 0 0 0 1-1.51V3a2 2 0 0 1 4 0v.09a1.65 1.65 0 0 0 1 1.51 1.65 1.65 0 0 0 1.82-.33l.06-.06a2 2 0 1 1 2.83 2.83l-.06.06a1.65 1.65 0 0 0-.33 1.82V9a1.65 1.65 0 0 0 1.51 1H21a2 2 0 0 1 0 4h-.09a1.65 1.65 0 0 0-1.51 1z' },
  { key: 'home', name: '首页', path: 'M3 9l9-7 9 7v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z;M9 22V12h6v10' },
  { key: 'search', name: '搜索', path: 'M11 19a8 8 0 1 0 0-16 8 8 0 0 0 0 16zM21 21l-4.35-4.35' },
  { key: 'flag', name: '目标', path: 'M4 15s1-1 4-1 5 2 8 2 4-1 4-1V3s-1 1-4 1-5-2-8-2-4 1-4 1zM4 22V15' },
  { key: 'target', name: '靶心', path: 'M12 22A10 10 0 1 0 12 2a10 10 0 0 0 0 20z;M12 18a6 6 0 1 0 0-12 6 6 0 0 0 0 12z;M12 14a2 2 0 1 0 0-4 2 2 0 0 0 0 4z' },
  { key: 'compass', name: '指南', path: 'M12 2a10 10 0 1 0 0 20 10 10 0 0 0 0-20zM16.24 7.76l-2.12 6.36-6.36 2.12 2.12-6.36 6.36-2.12z' },
  { key: 'anchor', name: '锚点', path: 'M12 2a3 3 0 1 0 0 6 3 3 0 0 0 0-6zM12 8v14;M5 12H2a10 10 0 0 0 20 0h-3' },
  { key: 'sun', name: '太阳', path: 'M12 17a5 5 0 1 0 0-10 5 5 0 0 0 0 10zM12 1v2;M12 21v2;M4.22 4.22l1.42 1.42;M18.36 18.36l1.42 1.42;M1 12h2;M21 12h2;M4.22 19.78l1.42-1.42;M18.36 5.64l1.42-1.42' },
  { key: 'moon', name: '月亮', path: 'M21 12.79A9 9 0 1 1 11.21 3 7 7 0 0 0 21 12.79z' },
  { key: 'cloud-rain', name: '下雨', path: 'M16 13v8;M8 13v8;M12 15v8;M20 16.58A5 5 0 0 0 18 7h-1.26A8 8 0 1 0 4 15.25' },
  { key: 'umbrella', name: '雨伞', path: 'M23 12a11 11 0 0 0-22 0;M12 12v8a2 2 0 0 0 4 0;M12 2v1' },
  { key: 'bike', name: '骑行', path: 'M3.82 15.5a4 4 0 0 0-1.32 5.68 4.1 4.1 0 0 0 5.68-1.32L12 12l3 5 3-4 4 2;M5.5 9.5L8 12l2.5-2.5;M14.5 6.5L12 9l3 2' },
  { key: 'car', name: '汽车', path: 'M5 17h14;M5 17a2 2 0 0 1-2-2V9l1-4h16l1 4v6a2 2 0 0 1-2 2;M7 17v2a1 1 0 0 0 1 1h1a1 1 0 0 0 1-1v-2;M14 17v2a1 1 0 0 0 1 1h1a1 1 0 0 0 1-1v-2;M7 7l-2 4' },
  { key: 'plane', name: '飞机', path: 'M17.8 19.2 16 11l3.5-3.5C21 6 21.5 4 21 3c-1-.5-3 0-4.5 1.5L13 8 4.8 6.2c-.5-.1-.9.1-1.1.5l-.3.5c-.2.5-.1 1 .3 1.3L9 12l-2 3H4l-1 1 3 2 2 3 1-1v-3l3-2 3.5 5.3c.3.4.8.5 1.3.3l.5-.2c.4-.3.6-.7.5-1.2z' },
  { key: 'train', name: '火车', path: 'M4 15.5A2.5 2.5 0 0 1 6.5 13h11a2.5 2.5 0 0 1 2.5 2.5v5a2.5 2.5 0 0 1-2.5 2.5h-11A2.5 2.5 0 0 1 4 20.5zM4 6v7.5M20 6v7.5M6 2h12l-1 4H7zM6.5 19.5h.01;M17.5 19.5h.01;M9 17h6' },
  { key: 'shopping-cart', name: '购物车', path: 'M6 2L3 6v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V6l-3-4zM3 6h18;M16 10a4 4 0 0 1-8 0' },
  { key: 'credit-card', name: '信用卡', path: 'M21 4H3a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h18a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2zM1 10h22;M7 15h4' },
  { key: 'gift', name: '礼物', path: 'M20 12v10H4V12M2 7h20v5H2zM12 22V7M12 7H7.5a2.5 2.5 0 1 1 0-5C11 2 12 7 12 7zM12 7h4.5a2.5 2.5 0 0 0 0-5C13 2 12 7 12 7z' },
  { key: 'cake', name: '蛋糕', path: 'M12 2a3 3 0 0 0-3 3c0 .5.1 1 .3 1.4L7 9H4a2 2 0 0 0-2 2v9a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-9a2 2 0 0 0-2-2h-3l-2.3-2.6A3 3 0 0 0 15 5a3 3 0 0 0-3-3zM12 2v2;M8 15h.01;M16 15h.01;M2 17h20' },
  { key: 'calendar', name: '日历', path: 'M19 4H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2zM16 2v4M8 2v4M3 10h18' },
  { key: 'clock', name: '时钟', path: 'M12 22a10 10 0 1 0 0-20 10 10 0 0 0 0 20zM12 6v6l4 2' },
  { key: 'mail', name: '邮箱', path: 'M4 4h16c1.1 0 2 .9 2 2v12c0 1.1-.9 2-2 2H4c-1.1 0-2-.9-2-2V6c0-1.1.9-2 2-2zM22 6l-10 7L2 6' },
  { key: 'map', name: '地址', path: 'M12 22s-8-7.58-8-13a8 8 0 0 1 16 0c0 5.42-8 13-8 13zM12 11a2 2 0 1 0 0-4 2 2 0 0 0 0 4z' },
  { key: 'navigation', name: '导航', path: 'M3 11l19-9-9 19-2-8z' },
  { key: 'phone', name: '电话', path: 'M22 16.92v3a2 2 0 0 1-2.18 2 19.79 19.79 0 0 1-8.63-3.07 19.5 19.5 0 0 1-6-6 19.79 19.79 0 0 1-3.07-8.67A2 2 0 0 1 4.11 2h3a2 2 0 0 1 2 1.72 12.84 12.84 0 0 0 .7 2.81 2 2 0 0 1-.45 2.11L8.09 9.91a16 16 0 0 0 6 6l1.27-1.27a2 2 0 0 1 2.11-.45 12.84 12.84 0 0 0 2.81.7A2 2 0 0 1 22 16.92z' },
  { key: 'award', name: '成就', path: 'M12 15a5 5 0 1 0 0-10 5 5 0 0 0 0 10zM8.21 13.89L7 23l5-3 5 3-1.21-9.12' },
  { key: 'trophy', name: '奖杯', path: 'M6 9H4.5a2.5 2.5 0 0 1 0-5H6;M18 9h1.5a2.5 2.5 0 0 0 0-5H18;M4 22h16;M10 14.66V17c0 .55-.47.98-.97 1.21C7.85 18.75 7 20.24 7 22;M14 14.66V17c0 .55.47.98.97 1.21C16.15 18.75 17 20.24 17 22;M18 2H6v7a6 6 0 0 0 12 0zM8 2v5;M16 2v5' },
  { key: 'medal', name: '奖章', path: 'M7 21l5-10 5 10;M9 9l5-6 5 6zM9 9H5l4-6M15 9h4l-4-6' },
  { key: 'layers', name: '架构', path: 'M12 2L2 7l10 5 10-5-10-5zM2 17l10 5 10-5M2 12l10 5 10-5' },
  { key: 'box', name: '盒子', path: 'M21 16V8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73l7 4a2 2 0 0 0 2 0l7-4A2 2 0 0 0 21 16zM3.27 6.96 12 12.01l8.73-5.05;M12 22.08V12' },
  { key: 'package', name: '包裹', path: 'M16.5 9.4 7.55 4.24;M21 16V8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73l7 4a2 2 0 0 0 2 0l7-4A2 2 0 0 0 21 16zM3.27 6.96 12 12.01l8.73-5.05;M12 22.08V12;M11 7l9 5' },
  { key: 'briefcase-business', name: '商务', path: 'M20 7h-4V5a2 2 0 0 0-2-2h-4a2 2 0 0 0-2 2v2H4a2 2 0 0 0-2 2v11a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V9a2 2 0 0 0-2-2z' },
  { key: 'pie-chart', name: '饼图', path: 'M21.21 15.89A10 10 0 1 1 8 2.83;M22 12A10 10 0 0 0 12 2v10z' },
  { key: 'bar-chart', name: '柱状图', path: 'M12 20V10;M18 20V4;M6 20v-4' },
  { key: 'activity', name: '活动', path: 'M22 12h-4l-3 9L9 3l-3 9H2' },
  { key: 'zap', name: '速度', path: 'M13 2L3 14h9l-1 8 10-12h-9l1-8z' },
  { key: 'shield', name: '安全', path: 'M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10z' },
  { key: 'lock', name: '锁定', path: 'M5 11h14v10H5zM8 11V7a4 4 0 0 1 8 0v4' },
  { key: 'key', name: '密钥', path: 'M21 2l-2 2m-7.61 7.61a5.5 5.5 0 1 1-7.778 7.778 5.5 5.5 0 0 1 7.777-7.777zm0 0L15.5 7.5m0 0 3 3L22 7l-3-3m-3.5 3.5L19 4' },
  { key: 'unlock', name: '解锁', path: 'M19 11H5a2 2 0 0 0-2 2v7a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7a2 2 0 0 0-2-2zM7 11V7a5 5 0 0 1 9.9-1' }
]

// 渲染 SVG 组件：根据 key 返回一个带 path 的 SVG 虚拟组件
const renderIcon = (key) => {
  const ic = iconList.find(i => i.key === key)
  if (!ic) return null
  const paths = ic.path.split(';').map((d, idx) => h('path', {
    key: idx,
    d,
    fill: 'none',
    stroke: 'currentColor',
    'stroke-width': 2,
    'stroke-linecap': 'round',
    'stroke-linejoin': 'round'
  }))
  return {
    render() {
      return h('svg', { viewBox: '0 0 24 24', width: '100%', height: '100%' }, paths)
    }
  }
}

// 打开图标选择器
const openIconPicker = (type, index) => {
  iconPicker.targetType = type
  iconPicker.targetIndex = index
  iconPicker.visible = true
}

const closeIconPicker = () => {
  iconPicker.visible = false
  iconPicker.targetType = ''
  iconPicker.targetIndex = -1
}

// 检查当前选项是否为选中目标的图标
const isPickerTargetIcon = (key) => {
  const arr = getPickerTargetArray()
  if (!arr || iconPicker.targetIndex < 0 || !arr[iconPicker.targetIndex]) return false
  return arr[iconPicker.targetIndex].icon === key
}

const getPickerTargetArray = () => {
  switch (iconPicker.targetType) {
    case 'project': return projects.value
    case 'aboutMe': return aboutMe.value
    case 'aboutSite': return aboutSite.value
    default: return null
  }
}

const selectIcon = (key) => {
  const arr = getPickerTargetArray()
  if (arr && arr[iconPicker.targetIndex]) {
    arr[iconPicker.targetIndex].icon = key
    triggerSave()
  }
  closeIconPicker()
}

const clearPickerIcon = () => {
  const arr = getPickerTargetArray()
  if (arr && arr[iconPicker.targetIndex]) {
    arr[iconPicker.targetIndex].icon = ''
    triggerSave()
  }
  closeIconPicker()
}

const clearTimers = () => {
  if (saveTimer) {
    clearTimeout(saveTimer)
    saveTimer = null
  }
  if (statusTimer) {
    clearTimeout(statusTimer)
    statusTimer = null
  }
}

onBeforeUnmount(() => {
  clearTimers()
})

onMounted(async () => {
  try {
    const res = await getAboutPage()
    if (res.data) {
      form.value.title = res.data.title || '关于我'
      form.value.subtitle = res.data.subtitle || ''
      form.value.bio = res.data.bio || ''

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

      try {
        siteHistory.value = JSON.parse(res.data.site_history || '[]')
        ensureUids()
      } catch (e) { siteHistory.value = [] }
    }
  } catch (e) {
    console.error(e)
  }
})

const doSave = async () => {
  if (saveTimer) {
    clearTimeout(saveTimer)
    saveTimer = null
  }
  saveStatus.value = 'saving'
  try {
    const toSave = siteHistory.value.map(({ __uid, ...rest }) => rest)
    await updateAboutPage({
      title: form.value.title,
      subtitle: form.value.subtitle,
      bio: form.value.bio,
      skills: JSON.stringify(skills.value),
      projects: JSON.stringify(projects.value),
      about_me: JSON.stringify(aboutMe.value),
      about_site: JSON.stringify(aboutSite.value),
      site_history: JSON.stringify(toSave)
    })
    saveStatus.value = 'saved'
    if (statusTimer) clearTimeout(statusTimer)
    statusTimer = setTimeout(() => {
      saveStatus.value = 'idle'
    }, 2000)
  } catch (e) {
    saveStatus.value = 'idle'
    ElMessage.error('保存失败')
    console.error(e)
  }
}

const triggerSave = () => {
  if (saveTimer) clearTimeout(saveTimer)
  saveTimer = setTimeout(() => {
    doSave()
  }, 400)
}

const addSkill = () => {
  if (newSkill.value.trim() && !skills.value.includes(newSkill.value.trim())) {
    skills.value.push(newSkill.value.trim())
    newSkill.value = ''
    triggerSave()
  }
}

const removeSkill = (index) => {
  skills.value.splice(index, 1)
  triggerSave()
}

const addProject = () => {
  projects.value.push({ name: '', description: '', url: '', icon: '' })
  triggerSave()
}

const removeProject = (index) => {
  projects.value.splice(index, 1)
  triggerSave()
}

const addAboutMe = () => {
  aboutMe.value.push({ label: '', value: '', icon: '' })
  triggerSave()
}

const removeAboutMe = (index) => {
  aboutMe.value.splice(index, 1)
  triggerSave()
}

const addAboutSite = () => {
  aboutSite.value.push({ label: '', value: '', icon: '' })
  triggerSave()
}

const removeAboutSite = (index) => {
  aboutSite.value.splice(index, 1)
  triggerSave()
}

const addSiteHistory = () => {
  const today = new Date()
  const y = today.getFullYear()
  const m = String(today.getMonth() + 1).padStart(2, '0')
  const d = String(today.getDate()).padStart(2, '0')
  siteHistory.value.push({ date: `${y}-${m}-${d}`, content: '', __uid: genUid() })
  triggerSave()
}

const removeSiteHistory = (index) => {
  if (index < 0 || index >= siteHistory.value.length) return
  siteHistory.value.splice(index, 1)
  triggerSave()
}
</script>

<style scoped>
.save-status {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  font-size: 13px;
  padding: 6px 12px;
  border-radius: 6px;
}

.status-saving {
  color: var(--accent-color);
  background: rgba(var(--accent-color-rgb), 0.08);
}

.status-saved {
  color: var(--success-color, #10b981);
  background: rgba(16, 185, 129, 0.1);
}

.status-idle {
  color: var(--card-text-color-secondary, #6b7280);
  background: rgba(107, 114, 128, 0.08);
}

.save-status svg {
  flex-shrink: 0;
}

.spin {
  animation: spin 1s linear infinite;
}

@keyframes spin {
  from { transform: rotate(0deg); }
  to { transform: rotate(360deg); }
}

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

/* 卡片网格：每行两个 */
.item-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 16px;
  margin-bottom: 16px;
}

@media (max-width: 720px) {
  .item-grid {
    grid-template-columns: 1fr;
  }
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
  padding: 14px;
  display: flex;
  flex-direction: column;
  gap: 10px;
  margin-bottom: 0;
}

.item-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 2px;
}

.item-index {
  width: 22px;
  height: 22px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: var(--accent-color);
  color: white;
  border-radius: 50%;
  font-size: 11px;
  font-weight: 600;
}

.btn-icon {
  padding: 6px;
  border: none;
  background: none;
  cursor: pointer;
  border-radius: 4px;
  transition: background 0.2s;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  color: var(--card-text-color-secondary);
}

.btn-icon:hover {
  background: rgba(0, 0, 0, 0.05);
}

.btn-danger {
  color: var(--error-color);
}

/* 图标选择器 */
.icon-picker-wrapper {
  position: relative;
}

.icon-display {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 8px;
  padding: 8px 12px;
  border: 1px solid var(--card-border-color, #dcdfe6);
  border-radius: 6px;
  background: var(--card-background, #fff);
  cursor: pointer;
  height: 38px;
  transition: border-color 0.2s;
}

.icon-display:hover {
  border-color: var(--accent-color);
}

.icon-svg {
  width: 20px;
  height: 20px;
  color: var(--accent-color);
  flex-shrink: 0;
}

.icon-placeholder {
  font-size: 13px;
  color: var(--card-text-color-placeholder, #c0c4cc);
  flex: 1;
}

.icon-caret {
  flex-shrink: 0;
  opacity: 0.5;
}

/* 图标选择弹层 */
.icon-picker-mask {
  position: fixed;
  inset: 0;
  background: rgba(0, 0, 0, 0.4);
  z-index: 9999;
  display: flex;
  align-items: center;
  justify-content: center;
}

.icon-picker-panel {
  width: min(560px, 92vw);
  max-height: 80vh;
  background: var(--card-background, #fff);
  border-radius: 10px;
  box-shadow: 0 20px 60px rgba(0, 0, 0, 0.25);
  display: flex;
  flex-direction: column;
  overflow: hidden;
}

.icon-picker-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 14px 18px;
  border-bottom: 1px solid var(--card-separator-color, #eee);
  font-weight: 600;
  color: var(--card-text-color-main);
}

.icon-picker-body {
  flex: 1;
  overflow-y: auto;
  padding: 14px;
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(92px, 1fr));
  gap: 10px;
}

.icon-option {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 6px;
  padding: 12px 8px;
  border: 1.5px solid transparent;
  border-radius: 8px;
  cursor: pointer;
  background: var(--body-background, #f5f7fa);
  transition: all 0.15s;
}

.icon-option:hover {
  border-color: var(--accent-color);
  background: rgba(var(--accent-color-rgb), 0.06);
}

.icon-option.active {
  border-color: var(--accent-color);
  background: rgba(var(--accent-color-rgb), 0.12);
}

.icon-svg-lg {
  width: 26px;
  height: 26px;
  color: var(--card-text-color-main);
}

.icon-option.active .icon-svg-lg,
.icon-option:hover .icon-svg-lg {
  color: var(--accent-color);
}

.icon-name {
  font-size: 12px;
  color: var(--card-text-color-secondary);
  text-align: center;
  line-height: 1.2;
}

.icon-picker-footer {
  padding: 12px 18px;
  border-top: 1px solid var(--card-separator-color, #eee);
  display: flex;
  justify-content: flex-end;
}

.btn-sm {
  padding: 6px 14px;
  font-size: 13px;
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

/* 建站历程 */
.section-desc {
  font-size: 13px;
  color: var(--card-text-color-secondary, #6b7280);
  margin: -8px 0 16px 0;
  line-height: 1.5;
}

.timeline-list {
  display: flex;
  flex-direction: column;
  gap: 14px;
}

.timeline-item-card {
  background: var(--body-background);
  border-radius: 8px;
  padding: 12px 14px;
  display: flex;
  flex-direction: column;
  gap: 10px;
  border-left: 3px solid var(--accent-color);
}

.timeline-item-date {
  font-size: 13px;
  font-weight: 600;
  color: var(--accent-color);
  font-variant-numeric: tabular-nums;
  letter-spacing: 0.2px;
}

.empty-tip {
  padding: 28px 16px;
  text-align: center;
  color: var(--card-text-color-secondary, #9ca3af);
  font-size: 14px;
  border: 1.5px dashed var(--card-separator-color, #e5e7eb);
  border-radius: 8px;
}

@media (max-width: 640px) {
  .timeline-item-card .form-row > .form-group {
    flex: 1 1 100% !important;
  }
}
</style>
