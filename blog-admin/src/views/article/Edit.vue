<template>
  <div class="article-edit-container">
    <!-- 顶部标题栏 -->
    <div class="article-title-bar">
      <input type="text" class="title-input" v-model="form.title" placeholder="输入文章标题...">
    </div>

    <!-- 元信息栏 -->
    <div class="article-meta-bar">
      <div class="meta-item">
        <label class="form-label">分类</label>
        <select class="form-select" v-model="form.category_id">
          <option value="">选择分类</option>
          <option v-for="cat in categories" :key="cat.id" :value="cat.id">{{ cat.name }}</option>
        </select>
      </div>
      <div class="meta-item">
        <label class="form-label">状态</label>
        <select class="form-select" v-model="form.status">
          <option value="draft">草稿</option>
          <option value="published">发布</option>
        </select>
      </div>
      <div class="meta-item meta-item-grow">
        <label class="form-label">标签</label>
        <div class="tags-container">
          <span v-for="tagId in form.tag_ids" :key="tagId" class="tag-item">
            {{ getTagName(tagId) }}
            <button class="tag-remove" @click="removeTag(tagId)">×</button>
          </span>
          <select class="form-select tag-select" v-model="selectedTag" @change="addTag">
            <option value="">选择标签...</option>
            <option v-for="tag in availableTags" :key="tag.id" :value="tag.id">{{ tag.name }}</option>
          </select>
        </div>
      </div>
    </div>

    <!-- 编辑器主体 -->
    <div class="editor-main">
      <MdEditor v-model="form.content" :theme="editorTheme" class="md-editor" />
    </div>

    <!-- 底部操作栏 -->
    <div class="article-actions">
      <button class="btn btn-secondary" @click="$router.push('/articles')">
        <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><line x1="19" y1="12" x2="5" y2="12"></line><polyline points="12 19 5 12 12 5"></polyline></svg>
        <span>返回列表</span>
      </button>
      <div class="action-right">
        <button class="btn btn-secondary" @click="handleSaveDraft">
          <span>💾</span>
          <span>保存草稿</span>
        </button>
        <button class="btn btn-primary" @click="handlePublish">
          <span>🚀</span>
          <span>发布文章</span>
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { MdEditor } from 'md-editor-v3'
import 'md-editor-v3/lib/style.css'
import { getArticleDetail, createArticle, updateArticle } from '../../api/article'
import { getCategoryList } from '../../api/category'
import { getTagList } from '../../api/tag'

const route = useRoute()
const router = useRouter()
const isEdit = computed(() => !!route.params.id)
const editorTheme = ref(localStorage.getItem('scheme') === 'dark' ? 'dark' : 'light')
const selectedTag = ref('')

const form = ref({
  title: '',
  content: '',
  summary: '',
  category_id: '',
  tag_ids: [],
  status: 'draft'
})

const categories = ref([])
const tags = ref([])

const availableTags = computed(() => {
  return tags.value.filter(t => !form.value.tag_ids.includes(t.id))
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

const loadData = async () => {
  try {
    const [catRes, tagRes] = await Promise.all([getCategoryList(), getTagList()])
    categories.value = catRes.data || []
    tags.value = tagRes.data || []
  } catch (e) { console.error(e) }

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

const handleSaveDraft = async () => {
  form.value.status = 'draft'
  await handleSave()
}

const handlePublish = async () => {
  form.value.status = 'published'
  await handleSave()
}

const handleSave = async () => {
  if (!form.value.title) {
    ElMessage.warning('请输入文章标题')
    return
  }
  try {
    if (isEdit.value) {
      await updateArticle(route.params.id, form.value)
    } else {
      await createArticle(form.value)
    }
    ElMessage.success('保存成功')
    router.push('/articles')
  } catch (e) { console.error(e) }
}

onMounted(loadData)
</script>

<style scoped>
.article-edit-container {
  display: flex;
  flex-direction: column;
  height: calc(100vh - 120px);
  gap: 16px;
}

/* 标题栏 */
.article-title-bar {
  background: var(--card-background);
  border-radius: var(--card-border-radius);
  padding: 16px 20px;
  box-shadow: var(--card-shadow);
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

/* 元信息栏 */
.article-meta-bar {
  display: flex;
  gap: 16px;
  background: var(--card-background);
  border-radius: var(--card-border-radius);
  padding: 16px 20px;
  box-shadow: var(--card-shadow);
  flex-wrap: wrap;
}

.meta-item {
  display: flex;
  flex-direction: column;
  gap: 6px;
  min-width: 150px;
}

.meta-item-grow {
  flex: 1;
  min-width: 200px;
}

.form-label {
  font-size: 12px;
  font-weight: 600;
  color: var(--text-color-secondary);
  text-transform: uppercase;
  letter-spacing: 0.5px;
}

.tags-container {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
  align-items: center;
}

.tag-item {
  display: inline-flex;
  align-items: center;
  gap: 4px;
  padding: 4px 10px;
  background: rgba(var(--accent-color-rgb), 0.1);
  color: var(--accent-color);
  border-radius: var(--tag-border-radius);
  font-size: 13px;
}

.tag-remove {
  background: none;
  color: var(--accent-color);
  font-size: 16px;
  padding: 0;
  line-height: 1;
}

.tag-select {
  width: auto;
  min-width: 120px;
}

/* 编辑器主体 */
.editor-main {
  flex: 1;
  min-height: 0;
  background: var(--card-background);
  border-radius: var(--card-border-radius);
  box-shadow: var(--card-shadow);
  overflow: hidden;
}

.md-editor {
  height: 100%;
}

/* 底部操作栏 */
.article-actions {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 16px 0;
}

.action-right {
  display: flex;
  gap: 12px;
}

/* 响应式 */
@media (max-width: 1024px) {
  .article-meta-bar {
    flex-direction: column;
  }

  .meta-item {
    width: 100%;
  }
}
</style>
