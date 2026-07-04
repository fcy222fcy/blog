<template>
  <div class="article-edit-container">
    <!-- 返回按钮 -->
    <button class="back-btn" @click="$router.push('/articles')">
      <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><line x1="19" y1="12" x2="5" y2="12"></line><polyline points="12 19 5 12 12 5"></polyline></svg>
      <span>返回列表</span>
    </button>

    <el-form
      ref="formRef"
      :model="form"
      :rules="rules"
      label-width="0"
      class="article-form"
    >
      <!-- 顶部标题栏 -->
      <div class="article-title-bar">
        <el-form-item prop="title" class="title-form-item">
          <input type="text" class="title-input" v-model="form.title" placeholder="输入文章标题...">
        </el-form-item>
      </div>

      <!-- 编辑器主体 -->
      <div class="editor-main">
        <MdEditor
          v-model="form.content"
          :theme="editorTheme"
          class="md-editor"
          :preview="true"
          previewTheme="github"
          @onUploadImg="onUploadImg"
        />
      </div>

      <!-- 底部操作栏 -->
      <div class="article-actions">
        <div class="action-left">
          <span class="word-count">
            <svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"></path><polyline points="14 2 14 8 20 8"></polyline><line x1="16" y1="13" x2="8" y2="13"></line><line x1="16" y1="17" x2="8" y2="17"></line></svg>
            {{ wordCount }} 字
          </span>
          <span v-if="autoSaveStatus" class="auto-save-status" :class="autoSaveStatus">
            <svg v-if="autoSaveStatus === 'saving'" xmlns="http://www.w3.org/2000/svg" width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"></circle></svg>
            <svg v-else-if="autoSaveStatus === 'saved'" xmlns="http://www.w3.org/2000/svg" width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polyline points="20 6 9 17 4 12"></polyline></svg>
            <svg v-else xmlns="http://www.w3.org/2000/svg" width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"></circle><line x1="15" y1="9" x2="9" y2="15"></line><line x1="9" y1="9" x2="15" y2="15"></line></svg>
            {{ autoSaveText }}
          </span>
        </div>
        <div class="action-right">
          <el-button @click="resetForm">
            <svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polyline points="1 4 1 10 7 10"></polyline><path d="M3.51 15a9 9 0 1 0 2.13-9.36L1 10"></path></svg>
            <span>重置</span>
          </el-button>
          <el-button type="default" @click="handleSaveDraft">
            <svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M19 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h11l5 5v11a2 2 0 0 1-2 2z"></path><polyline points="17 21 17 13 7 13 7 21"></polyline><polyline points="7 3 7 8 15 8"></polyline></svg>
            <span>保存草稿</span>
          </el-button>
          <el-button type="primary" @click="handlePublish">
            <svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><line x1="22" y1="2" x2="11" y2="13"></line><polygon points="22 2 15 22 11 13 2 9 22 2"></polygon></svg>
            <span>发布文章</span>
          </el-button>
        </div>
      </div>
    </el-form>

    <!-- 发布设置弹窗 -->
    <div class="publish-modal-overlay" :class="{ active: showPublishModal }" @click.self="showPublishModal = false">
      <div class="publish-modal">
        <div class="publish-modal-header">
          <h3>发布设置</h3>
          <button class="publish-modal-close" @click="showPublishModal = false">×</button>
        </div>
        <div class="publish-modal-body">
          <!-- 分类 -->
          <div class="publish-field">
            <label class="publish-label">分类 <span class="required">*</span></label>
            <select class="publish-select" v-model="form.category_id">
              <option value="">请选择分类</option>
              <option v-for="cat in categories" :key="cat.id" :value="cat.id">{{ cat.name }}</option>
            </select>
          </div>

          <!-- 标签 -->
          <div class="publish-field">
            <label class="publish-label">标签</label>
            <div class="publish-tags">
              <span v-for="tagId in form.tag_ids" :key="tagId" class="publish-tag-item">
                {{ getTagName(tagId) }}
                <button class="publish-tag-remove" @click="removeTag(tagId)">×</button>
              </span>
              <select class="publish-select publish-tag-select" v-model="selectedTag" @change="addTag">
                <option value="">添加标签...</option>
                <option v-for="tag in availableTags" :key="tag.id" :value="tag.id">{{ tag.name }}</option>
              </select>
            </div>
          </div>

          <!-- 封面图 -->
          <div class="publish-field">
            <label class="publish-label">封面图</label>
            <ImageUpload v-model="form.cover" />
          </div>

          <!-- 摘要 -->
          <div class="publish-field">
            <div class="publish-label-row">
              <label class="publish-label">摘要</label>
              <span class="publish-char-count" :class="{ 'exceeded': form.summary.length > 200 }">
                {{ form.summary.length }} / 200
              </span>
            </div>
            <textarea
              v-model="form.summary"
              class="publish-textarea"
              placeholder="请输入文章摘要（选填，最多200字）"
              rows="3"
              maxlength="200"
            ></textarea>
          </div>
        </div>
        <div class="publish-modal-footer">
          <button class="publish-btn-secondary" @click="showPublishModal = false">取消</button>
          <button class="publish-btn-primary" @click="confirmPublish">确认发布</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, watch, onMounted, onBeforeUnmount } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { ElMessage, ElMessageBox } from 'element-plus'
import { MdEditor } from 'md-editor-v3'
import 'md-editor-v3/lib/style.css'
import { getArticleDetail, createArticle, updateArticle } from '../../api/article'
import { getCategoryList } from '../../api/category'
import { getTagList } from '../../api/tag'
import { uploadFile } from '../../api/media'
import ImageUpload from '../../components/common/ImageUpload.vue'

const route = useRoute()
const router = useRouter()
const isEdit = computed(() => !!route.params.id)
const editorTheme = ref(localStorage.getItem('scheme') === 'dark' ? 'dark' : 'light')
const selectedTag = ref('')
const formRef = ref(null)
const showPublishModal = ref(false)

const form = ref({
  title: '',
  content: '',
  summary: '',
  cover: '',
  category_id: '',
  tag_ids: [],
  status: 'draft'
})

// 表单验证规则
const rules = {
  title: [
    { required: true, message: '请输入文章标题', trigger: 'blur' },
    { min: 2, max: 100, message: '标题长度在 2 到 100 个字符', trigger: 'blur' }
  ],
  category_id: [
    { required: true, message: '请选择文章分类', trigger: 'change' }
  ],
  summary: [
    { max: 200, message: '摘要不能超过 200 个字符', trigger: 'blur' }
  ]
}

const categories = ref([])
const tags = ref([])

const availableTags = computed(() => {
  return tags.value.filter(t => !form.value.tag_ids.includes(t.id))
})

const wordCount = computed(() => {
  const content = form.value.content || ''
  // 移除 Markdown 语法字符，统计实际内容字数
  const plainText = content
    .replace(/#{1,6}\s/g, '')
    .replace(/\*{1,3}(.*?)\*{1,3}/g, '$1')
    .replace(/`(.*?)`/g, '$1')
    .replace(/```[\s\S]*?```/g, '')
    .replace(/\[([^\]]+)\]\([^)]+\)/g, '$1')
    .replace(/!\[([^\]]*)\]\([^)]+\)/g, '')
    .replace(/[-*+]\s/g, '')
    .replace(/\d+\.\s/g, '')
    .replace(/>\s/g, '')
    .replace(/\|/g, '')
    .replace(/---+/g, '')
    .trim()
  return plainText.length
})

const getTagName = (id) => {
  const tag = tags.value.find(t => t.id === id)
  return tag?.name || id
}

const addTag = () => {
  if (selectedTag.value && !form.value.tag_ids.includes(selectedTag.value)) {
    form.value.tag_ids.push(selectedTag.value)
  }
  selectedTag.value = ''
}

const removeTag = (id) => {
  form.value.tag_ids = form.value.tag_ids.filter(t => t !== id)
}

// 图片上传处理
const onUploadImg = async (files, callback) => {
  const res = []
  for (const file of files) {
    try {
      const uploadRes = await uploadFile(file)
      if (uploadRes.code === 0) {
        res.push(uploadRes.data.url)
      } else {
        ElMessage.error('图片上传失败')
      }
    } catch (e) {
      console.error('上传失败:', e)
      ElMessage.error('图片上传失败')
    }
  }
  callback(res)
}

const loadData = async () => {
  try {
    const [catRes, tagRes] = await Promise.all([getCategoryList(), getTagList()])
    categories.value = catRes.data || []
    tags.value = tagRes.data || []
  } catch (e) { console.error(e) }

  // 新文章时尝试恢复 localStorage 草稿
  if (!isEdit.value) {
    await restoreLocalDraft()
    return
  }

  if (isEdit.value) {
    try {
      const res = await getArticleDetail(route.params.id)
      if (res.data) {
        form.value = {
          ...form.value,
          ...res.data,
          category_id: res.data.category_id || '',
          tag_ids: res.data.tag_ids || [],
          status: res.data.status || 'draft'
        }
      }
    } catch (e) { console.error(e) }
  }
}

// 自动保存草稿
const autoSaveStatus = ref('') // '' | 'saving' | 'saved' | 'error'
const autoSaveText = computed(() => {
  switch (autoSaveStatus.value) {
    case 'saving': return '正在保存...'
    case 'saved': return '草稿已自动保存'
    case 'error': return '自动保存失败'
    default: return ''
  }
})
let autoSaveTimer = null
let lastAutoSavedTitle = ''
let lastAutoSavedContent = ''

// localStorage 草稿相关
const DRAFT_KEY = 'article_draft_new'

const saveLocalDraft = () => {
  if (!form.value.title && !form.value.content) return
  const draft = { ...form.value, savedAt: new Date().toISOString() }
  localStorage.setItem(DRAFT_KEY, JSON.stringify(draft))
}

const loadLocalDraft = () => {
  try {
    const raw = localStorage.getItem(DRAFT_KEY)
    if (!raw) return null
    return JSON.parse(raw)
  } catch {
    return null
  }
}

const clearLocalDraft = () => {
  localStorage.removeItem(DRAFT_KEY)
}

const restoreLocalDraft = async () => {
  if (isEdit.value) return
  const draft = loadLocalDraft()
  if (!draft) return

  try {
    await ElMessageBox.confirm(
      `检测到未保存的草稿（${draft.title || '无标题'}），是否恢复？`,
      '恢复草稿',
      {
        confirmButtonText: '恢复',
        cancelButtonText: '丢弃',
        type: 'info'
      }
    )
    form.value = {
      ...form.value,
      title: draft.title || '',
      content: draft.content || '',
      summary: draft.summary || '',
      cover: draft.cover || '',
      category_id: draft.category_id || '',
      tag_ids: draft.tag_ids || [],
      status: draft.status || 'draft'
    }
    ElMessage.success('草稿已恢复')
  } catch {
    clearLocalDraft()
  }
}

const autoSaveDraft = async () => {
  // 新文章：保存到 localStorage
  if (!isEdit.value) {
    if (!form.value.title && !form.value.content) return
    saveLocalDraft()
    autoSaveStatus.value = 'saved'
    setTimeout(() => { autoSaveStatus.value = '' }, 3000)
    return
  }

  // 编辑模式：保存到服务端
  if (form.value.title === lastAutoSavedTitle && form.value.content === lastAutoSavedContent) return
  if (!form.value.title || form.value.title.trim().length < 2) return

  autoSaveStatus.value = 'saving'
  try {
    await updateArticle(route.params.id, { ...form.value, status: 'draft' })
    lastAutoSavedTitle = form.value.title
    lastAutoSavedContent = form.value.content
    autoSaveStatus.value = 'saved'
    setTimeout(() => { autoSaveStatus.value = '' }, 3000)
  } catch (e) {
    console.error('自动保存失败:', e)
    autoSaveStatus.value = 'error'
    setTimeout(() => { autoSaveStatus.value = '' }, 3000)
  }
}

const debouncedAutoSave = () => {
  if (autoSaveTimer) clearTimeout(autoSaveTimer)
  autoSaveTimer = setTimeout(autoSaveDraft, 3000)
}

// 监听标题和内容变化触发自动保存
watch(
  () => [form.value.title, form.value.content],
  () => { debouncedAutoSave() }
)

onBeforeUnmount(() => {
  if (autoSaveTimer) clearTimeout(autoSaveTimer)
})

const handleSaveDraft = async () => {
  form.value.status = 'draft'
  await handleSave(true)
}

const handlePublish = async () => {
  // 验证标题
  if (!form.value.title || form.value.title.trim().length < 2) {
    ElMessage.warning('请先输入文章标题（至少2个字符）')
    return
  }

  // 验证内容长度
  if (!form.value.content || form.value.content.trim().length < 10) {
    ElMessage.warning('文章内容至少需要 10 个字符')
    return
  }

  // 打开发布设置弹窗
  showPublishModal.value = true
}

const confirmPublish = async () => {
  // 验证分类
  if (!form.value.category_id) {
    ElMessage.warning('请选择文章分类')
    return
  }

  form.value.status = 'published'
  showPublishModal.value = false
  await handleSave()
}

const handleSave = async (isDraft = false) => {
  // 草稿保存只验证标题
  if (!form.value.title || form.value.title.trim().length < 2) {
    ElMessage.warning('请输入文章标题（至少2个字符）')
    return
  }

  try {
    if (isEdit.value) {
      await updateArticle(route.params.id, form.value)
    } else {
      await createArticle(form.value)
    }
    // 清除 localStorage 草稿
    clearLocalDraft()
    // 重置自动保存跟踪
    lastAutoSavedTitle = form.value.title
    lastAutoSavedContent = form.value.content
    if (isDraft) {
      autoSaveStatus.value = 'saved'
      setTimeout(() => { autoSaveStatus.value = '' }, 3000)
      ElMessage.success('草稿保存成功')
      return
    }
    ElMessage.success('保存成功')
    router.push('/articles')
  } catch (e) {
    console.error(e)
    ElMessage.error('保存失败，请重试')
  }
}

// 重置表单
const resetForm = async () => {
  try {
    await ElMessageBox.confirm(
      '确定要重置表单吗？所有未保存的更改将丢失。',
      '确认重置',
      {
        confirmButtonText: '确定重置',
        cancelButtonText: '取消',
        type: 'warning'
      }
    )

    // 重置表单验证状态
    if (formRef.value) {
      formRef.value.resetFields()
    }

    // 重置表单数据
    form.value = {
      title: '',
      content: '',
      summary: '',
      cover: '',
      category_id: '',
      tag_ids: [],
      status: 'draft'
    }

    // 清除 localStorage 草稿
    clearLocalDraft()

    // 重置自动保存跟踪
    lastAutoSavedTitle = ''
    lastAutoSavedContent = ''
    autoSaveStatus.value = ''

    // 如果是编辑模式，重新加载数据
    if (isEdit.value) {
      await loadData()
    }

    ElMessage.success('表单已重置')
  } catch (error) {
    // 用户取消操作
    if (error !== 'cancel') {
      console.error('重置失败:', error)
    }
  }
}

onMounted(loadData)
</script>

<style scoped>
.article-edit-container {
  display: flex;
  flex-direction: column;
  gap: 16px;
  padding-bottom: 80px;
}

/* 表单样式 */
.article-form {
  display: flex;
  flex-direction: column;
  gap: 0;
}

.article-form :deep(.el-form-item) {
  margin-bottom: 0;
}

.article-form :deep(.el-form-item__error) {
  padding-top: 4px;
  font-size: 12px;
}

/* 返回按钮 */
.back-btn {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  padding: 8px 16px;
  background: transparent;
  border: 1px solid var(--border-color);
  border-radius: var(--card-border-radius);
  color: var(--text-color-secondary);
  font-size: 14px;
  cursor: pointer;
  transition: all 0.2s ease;
  width: fit-content;
}

.back-btn:hover {
  background: var(--card-background);
  color: var(--text-color-main);
  border-color: var(--text-color-secondary);
}

/* 标题栏 */
.article-title-bar {
  background: var(--card-background);
  border-radius: var(--card-border-radius);
  padding: 16px 20px;
  box-shadow: var(--card-shadow);
}

.title-form-item {
  width: 100%;
}

.title-form-item :deep(.el-form-item__content) {
  line-height: normal;
}

.title-input {
  width: 100%;
  padding: 8px 0;
  border: none;
  font-size: 24px;
  font-weight: 700;
  background: transparent;
  color: var(--card-text-color-main);
}

.title-input::placeholder {
  color: var(--text-color-tertiary);
}

.title-input:focus {
  outline: none;
}

/* 编辑器主体 */
.editor-main {
  background: var(--card-background);
  border-radius: var(--card-border-radius);
  box-shadow: var(--card-shadow);
  overflow: hidden;
}

.md-editor {
  height: auto;
}

/* 底部操作栏 */
.article-actions {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 16px 24px;
  background: var(--card-background);
  border-radius: var(--card-border-radius);
  box-shadow: 0 -2px 10px rgba(0, 0, 0, 0.1);
  position: fixed;
  bottom: 0;
  left: var(--sidebar-width);
  right: 0;
  z-index: 100;
}

.action-left,
.action-right {
  display: flex;
  gap: 12px;
  align-items: center;
}

/* 字数统计 */
.word-count {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  font-size: 13px;
  color: var(--text-color-secondary);
  padding: 6px 12px;
  background: var(--input-background);
  border-radius: 6px;
}

.word-count svg {
  opacity: 0.6;
}

/* 自动保存状态 */
.auto-save-status {
  display: inline-flex;
  align-items: center;
  gap: 5px;
  font-size: 12px;
  padding: 4px 10px;
  border-radius: 12px;
  animation: fadeIn 0.3s ease;
}

.auto-save-status.saving {
  color: var(--text-color-secondary);
  background: var(--input-background);
}

.auto-save-status.saved {
  color: var(--success-color);
  background: rgba(40, 167, 69, 0.08);
}

.auto-save-status.error {
  color: var(--error-color);
  background: rgba(220, 53, 69, 0.08);
}

@keyframes fadeIn {
  from { opacity: 0; transform: translateY(4px); }
  to { opacity: 1; transform: translateY(0); }
}

/* 发布设置弹窗 */
.publish-modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.5);
  display: none;
  align-items: center;
  justify-content: center;
  z-index: 2000;
  backdrop-filter: blur(4px);
}

.publish-modal-overlay.active {
  display: flex;
}

.publish-modal {
  background: var(--card-background);
  border-radius: 12px;
  width: 90%;
  max-width: 560px;
  max-height: 85vh;
  display: flex;
  flex-direction: column;
  box-shadow: 0 20px 60px rgba(0, 0, 0, 0.3);
}

.publish-modal-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 20px 24px;
  border-bottom: 1px solid var(--card-separator-color);
}

.publish-modal-header h3 {
  font-size: 18px;
  font-weight: 600;
  color: var(--card-text-color-main);
  margin: 0;
}

.publish-modal-close {
  background: none;
  font-size: 24px;
  color: var(--card-text-color-tertiary);
  padding: 4px 8px;
  border-radius: 6px;
  transition: all 0.2s;
}

.publish-modal-close:hover {
  background: var(--body-background);
  color: var(--card-text-color-main);
}

.publish-modal-body {
  padding: 24px;
  overflow-y: auto;
  flex: 1;
}

.publish-field {
  margin-bottom: 24px;
}

.publish-field:last-child {
  margin-bottom: 0;
}

.publish-label {
  display: block;
  font-size: 14px;
  font-weight: 600;
  color: var(--card-text-color-main);
  margin-bottom: 10px;
}

.publish-label .required {
  color: var(--danger-color);
}

.publish-label-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 10px;
}

.publish-select {
  width: 100%;
  padding: 12px 16px;
  border: 1px solid var(--card-separator-color);
  border-radius: 8px;
  font-size: 14px;
  background: var(--card-background);
  color: var(--card-text-color-main);
  cursor: pointer;
  transition: all 0.2s;
}

.publish-select:focus {
  outline: none;
  border-color: var(--accent-color);
  box-shadow: 0 0 0 3px rgba(var(--accent-color-rgb), 0.1);
}

.publish-tags {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
  align-items: center;
}

.publish-tag-item {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  padding: 6px 12px;
  background: rgba(var(--accent-color-rgb), 0.1);
  color: var(--accent-color);
  border-radius: 6px;
  font-size: 13px;
  font-weight: 500;
}

.publish-tag-remove {
  background: none;
  color: var(--accent-color);
  font-size: 16px;
  padding: 0;
  line-height: 1;
  opacity: 0.7;
  transition: opacity 0.2s;
}

.publish-tag-remove:hover {
  opacity: 1;
}

.publish-tag-select {
  width: auto;
  min-width: 140px;
}

.publish-textarea {
  width: 100%;
  padding: 12px 16px;
  border: 1px solid var(--card-separator-color);
  border-radius: 8px;
  font-size: 14px;
  background: var(--card-background);
  color: var(--card-text-color-main);
  resize: vertical;
  min-height: 80px;
  font-family: inherit;
  transition: all 0.2s;
}

.publish-textarea:focus {
  outline: none;
  border-color: var(--accent-color);
  box-shadow: 0 0 0 3px rgba(var(--accent-color-rgb), 0.1);
}

.publish-char-count {
  font-size: 12px;
  color: var(--card-text-color-tertiary);
}

.publish-char-count.exceeded {
  color: var(--danger-color);
  font-weight: 600;
}

.publish-modal-footer {
  display: flex;
  justify-content: flex-end;
  gap: 12px;
  padding: 16px 24px;
  border-top: 1px solid var(--card-separator-color);
}

.publish-btn-secondary {
  padding: 10px 20px;
  border: 1px solid var(--card-separator-color);
  border-radius: 8px;
  background: var(--card-background);
  color: var(--card-text-color-main);
  font-size: 14px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s;
}

.publish-btn-secondary:hover {
  border-color: var(--card-text-color-tertiary);
}

.publish-btn-primary {
  padding: 10px 24px;
  border: none;
  border-radius: 8px;
  background: var(--accent-color);
  color: var(--accent-color-text);
  font-size: 14px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s;
}

.publish-btn-primary:hover {
  background: var(--accent-color-darker);
}

/* 响应式 */
@media (max-width: 768px) {
  .publish-modal {
    width: 95%;
    max-height: 90vh;
  }

  .publish-modal-body {
    padding: 20px;
  }
}
</style>
