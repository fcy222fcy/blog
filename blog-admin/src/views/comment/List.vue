<template>
  <div>
    <div class="stats-grid">
      <div class="stat-card">
        <div class="stat-label">全部评论</div>
        <div class="stat-value">{{ comments.length }}</div>
      </div>
      <div class="stat-card">
        <div class="stat-label">待审核</div>
        <div class="stat-value" style="color: var(--warning-color);">{{ comments.filter(c => c.status === 'pending').length }}</div>
      </div>
      <div class="stat-card">
        <div class="stat-label">已通过</div>
        <div class="stat-value" style="color: var(--success-color);">{{ comments.filter(c => c.status === 'approved').length }}</div>
      </div>
      <div class="stat-card">
        <div class="stat-label">已拒绝</div>
        <div class="stat-value" style="color: var(--danger-color);">{{ comments.filter(c => c.status === 'rejected').length }}</div>
      </div>
    </div>

    <div style="display: flex; gap: 12px; margin-bottom: 20px; align-items: center; justify-content: space-between;">
      <select class="form-select" style="width: auto; padding: 9px 32px 9px 12px;" v-model="statusFilter" @change="loadComments">
        <option value="">全部状态</option>
        <option value="pending">待审核</option>
        <option value="approved">已通过</option>
        <option value="rejected">已拒绝</option>
      </select>
      <div class="search-box">
        <span class="search-box-icon">⌕</span>
        <input type="text" v-model="keyword" placeholder="搜索评论...">
      </div>
    </div>

    <div class="comment-cards">
      <div v-for="comment in filteredComments" :key="comment.id" class="comment-card">
        <div class="comment-card-header">
          <div class="comment-card-user">
            <div class="comment-avatar" :style="getAvatarStyle(comment.status)">
              {{ (comment.nickname || '匿名').charAt(0) }}
            </div>
            <div class="comment-user-info">
              <div class="comment-user-name">{{ comment.nickname || '匿名用户' }}</div>
              <div class="comment-user-email">{{ comment.email || '' }}</div>
            </div>
          </div>
          <span class="status-badge" :class="'status-' + comment.status">{{ statusText(comment.status) }}</span>
        </div>
        <div class="comment-card-body">
          <p class="comment-content">{{ comment.content }}</p>
          <div class="comment-meta">
            <span>评论于 <a style="color: var(--accent-color);">{{ comment.article_title || '文章' }}</a></span>
            <span>·</span>
            <span>{{ formatDate(comment.created_at) }}</span>
          </div>
        </div>
        <div class="comment-card-footer">
          <button v-if="comment.status !== 'approved'" class="action-btn btn-edit btn-sm" @click="handleStatus(comment.id, 'approved')">通过</button>
          <button v-if="comment.status !== 'rejected'" class="action-btn btn-hide btn-sm" @click="handleStatus(comment.id, 'rejected')">拒绝</button>
          <button class="action-btn btn-delete btn-sm" @click="handleDelete(comment.id)">删除</button>
        </div>
      </div>
    </div>

    <div v-if="filteredComments.length === 0" class="card">
      <div class="card-body" style="text-align: center; color: var(--card-text-color-tertiary);">暂无评论</div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { getCommentList, updateCommentStatus, deleteComment } from '../../api/comment'

const comments = ref([])
const statusFilter = ref('')
const keyword = ref('')

const formatDate = (d) => d ? new Date(d).toLocaleDateString('zh-CN') : ''
const statusText = (s) => ({ pending: '待审核', approved: '已通过', rejected: '已拒绝' }[s] || s)

const getAvatarStyle = (status) => {
  const colors = {
    pending: 'background: rgba(var(--accent-color-rgb), 0.1); color: var(--accent-color);',
    approved: 'background: rgba(16, 185, 129, 0.1); color: var(--success-color);',
    rejected: 'background: rgba(239, 68, 68, 0.1); color: var(--danger-color);'
  }
  return colors[status] || colors.pending
}

const filteredComments = computed(() => {
  let list = comments.value
  if (statusFilter.value) list = list.filter(c => c.status === statusFilter.value)
  if (keyword.value) list = list.filter(c => (c.nickname || '').includes(keyword.value) || (c.content || '').includes(keyword.value))
  return list
})

const loadComments = async () => {
  try {
    const res = await getCommentList({ page: 1, page_size: 100 })
    comments.value = res.data?.list || []
  } catch (e) { console.error(e) }
}

const handleStatus = async (id, status) => {
  try {
    await updateCommentStatus(id, status)
    ElMessage.success('操作成功')
    loadComments()
  } catch (e) { console.error(e) }
}

const handleDelete = async (id) => {
  try {
    await ElMessageBox.confirm('确定要删除这条评论吗？', '确认删除', { confirmButtonText: '删除', cancelButtonText: '取消', type: 'warning' })
    await deleteComment(id)
    ElMessage.success('删除成功')
    loadComments()
  } catch (e) { if (e !== 'cancel') console.error(e) }
}

onMounted(loadComments)
</script>

<style scoped>
.comment-user-info { display: flex; flex-direction: column; }
.comment-user-name { font-size: 14px; font-weight: 600; color: var(--card-text-color-main); }
.comment-user-email { font-size: 12px; color: var(--card-text-color-tertiary); }
.comment-meta { font-size: 13px; color: var(--card-text-color-tertiary); display: flex; gap: 6px; }
.btn-hide { background: rgba(245, 158, 11, 0.08); color: var(--warning-color); }
.btn-hide:hover { background: var(--warning-color); color: white; }
</style>
