<template>
  <div>
    <div style="margin-bottom: 16px;">
      <button class="btn btn-secondary" @click="$router.push('/articles')">
        <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><line x1="19" y1="12" x2="5" y2="12"></line><polyline points="12 19 5 12 12 5"></polyline></svg>
        <span>返回列表</span>
      </button>
    </div>

    <div class="edit-layout">
      <div class="edit-panel">
        <div class="card">
          <div class="card-header">
            <div class="card-title">Markdown 编辑</div>
          </div>
          <div class="card-body">
            <div class="form-group" style="margin-bottom: 12px;">
              <input type="text" class="title-input" v-model="form.title" placeholder="输入文章标题...">
            </div>
            <div class="form-row">
              <div class="form-group" style="flex: 1;">
                <label class="form-label">分类</label>
                <select class="form-select" v-model="form.category_id">
                  <option value="">选择分类</option>
                  <option v-for="cat in categories" :key="cat.id" :value="cat.id">{{ cat.name }}</option>
                </select>
              </div>
              <div class="form-group" style="flex: 1;">
                <label class="form-label">状态</label>
                <select class="form-select" v-model="form.status">
                  <option value="draft">草稿</option>
                  <option value="published">发布</option>
                </select>
              </div>
            </div>
            <div class="form-group">
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
            <MdEditor v-model="form.content" :theme="editorTheme" style="height: 500px;" />
          </div>
        </div>
      </div>

      <div class="edit-panel">
        <div class="card">
          <div class="card-header">
            <div class="card-title">实时预览</div>
          </div>
          <div class="card-body">
            <div class="preview-content" v-html="previewHtml"></div>
          </div>
        </div>
      </div>
    </div>

    <div style="margin-top: 16px; display: flex; gap: 12px; justify-content: flex-end;">
      <button class="btn btn-secondary" @click="$router.push('/articles')">
        <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><line x1="19" y1="12" x2="5" y2="12"></line><polyline points="12 19 5 12 12 5"></polyline></svg>
        <span>返回列表</span>
      </button>
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
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { marked } from 'marked'
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

const previewHtml = computed(() => {
  try {
    return marked(form.value.content || '')
  } catch (e) {
    return '<p>预览加载中...</p>'
  }
})

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
      form.value = { ...form.value, ...res.data }
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
.edit-layout {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 16px;
}
.edit-panel {
  min-width: 0;
}
.title-input {
  width: 100%;
  padding: 12px 16px;
  border: 1px solid var(--card-separator-color);
  border-radius: var(--card-border-radius);
  font-size: 18px;
  font-weight: 600;
  background: var(--card-background);
  color: var(--card-text-color-main);
}
.title-input:focus {
  outline: none;
  border-color: var(--accent-color);
  box-shadow: 0 0 0 3px rgba(var(--accent-color-rgb), 0.1);
}
.form-row {
  display: flex;
  gap: 16px;
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
.preview-content {
  min-height: 500px;
  line-height: 1.8;
  color: var(--card-text-color-main);
}
.preview-content :deep(h1) { font-size: 24px; font-weight: 700; margin: 20px 0 12px; }
.preview-content :deep(h2) { font-size: 20px; font-weight: 600; margin: 18px 0 10px; }
.preview-content :deep(h3) { font-size: 17px; font-weight: 600; margin: 16px 0 8px; }
.preview-content :deep(p) { margin: 10px 0; }
.preview-content :deep(code) {
  background: var(--body-background);
  padding: 2px 6px;
  border-radius: 4px;
  font-size: 14px;
}
.preview-content :deep(pre) {
  background: var(--body-background);
  padding: 16px;
  border-radius: var(--card-border-radius);
  overflow-x: auto;
}
.preview-content :deep(blockquote) {
  border-left: 4px solid var(--accent-color);
  padding-left: 16px;
  color: var(--card-text-color-secondary);
  margin: 12px 0;
}
.preview-content :deep(ul), .preview-content :deep(ol) {
  padding-left: 24px;
  margin: 10px 0;
}

@media (max-width: 1024px) {
  .edit-layout {
    grid-template-columns: 1fr;
  }
}
</style>
