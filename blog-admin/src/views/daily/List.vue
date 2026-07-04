<template>
  <div>
    <div class="stats-grid">
      <div class="stat-card">
        <div class="stat-label">问题总数</div>
        <div class="stat-value">{{ questions.length }}</div>
      </div>
      <div class="stat-card">
        <div class="stat-label">显示问题</div>
        <div class="stat-value">{{ questions.filter(q => q.status === 1).length }}</div>
      </div>
      <div class="stat-card">
        <div class="stat-label">总浏览量</div>
        <div class="stat-value">{{ questions.reduce((s, q) => s + (q.view_count || 0), 0) }}</div>
      </div>
      <div class="stat-card">
        <div class="stat-label">总点赞数</div>
        <div class="stat-value">{{ questions.reduce((s, q) => s + (q.like_count || 0), 0) }}</div>
      </div>
    </div>

    <div class="card" style="margin-bottom: 20px;">
      <div class="card-header">
        <div class="card-title">问题管理</div>
        <div style="display: flex; gap: 12px; flex-wrap: wrap; align-items: center;">
          <div class="search-box">
            <span class="search-box-icon">⌕</span>
            <input type="text" v-model="keyword" placeholder="搜索问题...">
          </div>
          <select class="form-select" style="width: auto; padding: 9px 32px 9px 12px;" v-model="statusFilter">
            <option value="">全部状态</option>
            <option value="1">显示</option>
            <option value="0">隐藏</option>
          </select>
          <button class="btn btn-primary" @click="showModal = true; resetForm()">
            <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><line x1="12" y1="5" x2="12" y2="19"></line><line x1="5" y1="12" x2="19" y2="12"></line></svg>
            <span>新建问题</span>
          </button>
        </div>
      </div>
      <div class="card-body">
        <table class="table">
          <thead>
            <tr>
              <th>问题</th>
              <th>显示日期</th>
              <th>状态</th>
              <th>浏览/点赞</th>
              <th>操作</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="q in filteredQuestions" :key="q.id">
              <td style="max-width: 300px; overflow: hidden; text-overflow: ellipsis; white-space: nowrap;">{{ q.question }}</td>
              <td>{{ q.date }}</td>
              <td>
                <span class="status-badge" :class="q.status === 1 ? 'status-published' : 'status-draft'">
                  {{ q.status === 1 ? '显示' : '隐藏' }}
                </span>
              </td>
              <td>{{ q.view_count || 0 }} / {{ q.like_count || 0 }}</td>
              <td>
                <div style="display: flex; gap: 6px;">
                  <button class="action-btn btn-edit btn-sm" @click="editQuestion(q)">编辑</button>
                  <button class="action-btn btn-sm" :class="q.status === 1 ? 'btn-hide' : 'btn-edit'" @click="toggleStatus(q)">
                    {{ q.status === 1 ? '隐藏' : '显示' }}
                  </button>
                  <button class="action-btn btn-delete btn-sm" @click="handleDelete(q.id)">删除</button>
                </div>
              </td>
            </tr>
            <tr v-if="filteredQuestions.length === 0">
              <td colspan="5" style="text-align: center; color: var(--card-text-color-tertiary); padding: 40px;">暂无问题</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <div class="modal-overlay" :class="{ active: showModal }" @click.self="showModal = false">
      <div class="modal" style="max-width: 900px;">
        <div class="modal-header">
          <h3 class="modal-title">{{ editingId ? '编辑问题' : '新建问题' }}</h3>
          <button class="modal-close" @click="showModal = false">×</button>
        </div>
        <div class="modal-body">
          <div class="form-group">
            <label class="form-label">日期</label>
            <input type="date" class="form-input" v-model="form.date">
          </div>
          <div class="form-group">
            <label class="form-label">问题 <span class="required">*</span></label>
            <textarea class="form-textarea" v-model="form.question" placeholder="输入问题..." rows="3"></textarea>
          </div>
          <div class="form-group" style="margin-bottom: 0;">
            <label class="form-label">答案（支持 Markdown）</label>
            <MdEditor
              v-model="form.answer"
              :theme="editorTheme"
              class="md-editor"
              :preview="true"
              previewTheme="github"
              @onUploadImg="onUploadImg"
              style="height: 400px;"
            />
          </div>
        </div>
        <div class="modal-footer">
          <button class="btn btn-secondary" @click="showModal = false">取消</button>
          <button class="btn btn-primary" @click="handleSave">保存</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { MdEditor } from 'md-editor-v3'
import 'md-editor-v3/lib/style.css'
import { getDailyQuestionList, createDailyQuestion, updateDailyQuestion, deleteDailyQuestion, updateDailyQuestionStatus } from '../../api/daily'
import { uploadFile } from '../../api/media'

const questions = ref([])
const keyword = ref('')
const statusFilter = ref('')
const showModal = ref(false)
const editingId = ref(null)
const form = ref({ question: '', answer: '', date: '' })
const editorTheme = ref(localStorage.getItem('scheme') === 'dark' ? 'dark' : 'light')

const filteredQuestions = computed(() => {
  let list = questions.value
  if (statusFilter.value !== '') list = list.filter(q => String(q.status) === statusFilter.value)
  if (keyword.value) list = list.filter(q => (q.question || '').includes(keyword.value))
  return list
})

const resetForm = () => { editingId.value = null; form.value = { question: '', answer: '', date: '' } }

const editQuestion = (q) => {
  editingId.value = q.id
  form.value = { question: q.question, answer: q.answer, date: q.date }
  showModal.value = true
}

const loadQuestions = async () => {
  try {
    const res = await getDailyQuestionList({ page: 1, page_size: 100 })
    questions.value = res.data?.list || []
  } catch (e) { console.error(e) }
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

const handleSave = async () => {
  if (!form.value.question) { ElMessage.warning('请输入问题'); return }
  try {
    if (editingId.value) { await updateDailyQuestion(editingId.value, form.value) }
    else { await createDailyQuestion(form.value) }
    ElMessage.success('保存成功')
    showModal.value = false
    loadQuestions()
  } catch (e) { console.error(e) }
}

const toggleStatus = async (q) => {
  try {
    await updateDailyQuestionStatus(q.id, q.status === 1 ? 0 : 1)
    ElMessage.success('操作成功')
    loadQuestions()
  } catch (e) { console.error(e) }
}

const handleDelete = async (id) => {
  try {
    await ElMessageBox.confirm('确定要删除这个问题吗？', '确认删除', { confirmButtonText: '删除', cancelButtonText: '取消', type: 'warning' })
    await deleteDailyQuestion(id)
    ElMessage.success('删除成功')
    loadQuestions()
  } catch (e) { if (e !== 'cancel') console.error(e) }
}

onMounted(loadQuestions)
</script>

<style scoped>
.btn-hide { background: rgba(245, 158, 11, 0.08); color: var(--warning-color); }
.btn-hide:hover { background: var(--warning-color); color: white; }
.md-editor { border-radius: 8px; overflow: hidden; }
</style>
