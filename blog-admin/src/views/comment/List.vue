<template>
  <div>
    <!-- 统计卡片（加载态 + 真实态） -->
    <SkeletonLoader v-if="loading" type="stats" :count="4" />
    <div v-else class="stats-grid">
      <div class="stat-card">
        <div class="stat-label">全部评论</div>
        <div class="stat-value">{{ comments.length }}</div>
      </div>
      <div class="stat-card">
        <div class="stat-label">待审核</div>
        <div class="stat-value text-warning">{{ comments.filter(c => c.status === 'pending').length }}</div>
      </div>
      <div class="stat-card">
        <div class="stat-label">已通过</div>
        <div class="stat-value text-success">{{ comments.filter(c => c.status === 'approved').length }}</div>
      </div>
      <div class="stat-card">
        <div class="stat-label">已拒绝</div>
        <div class="stat-value text-danger">{{ comments.filter(c => c.status === 'rejected').length }}</div>
      </div>
    </div>

    <div class="filter-bar">
      <CustomSelect v-model="statusFilter" :options="statusOptions" />
      <div class="search-box">
        <span class="search-box-icon">⌕</span>
        <input type="text" v-model="keyword" placeholder="搜索评论...">
      </div>
    </div>

    <div class="comment-cards">
      <!-- 骨架屏加载态 -->
      <SkeletonLoader v-if="loading" type="comment" :count="4" />
      <!-- 真实列表 -->
      <template v-else>
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
            <span>评论于 <a class="text-accent">{{ comment.article_title || '文章' }}</a></span>
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
      </template>
    </div>

    <div v-if="filteredComments.length === 0 && !loading" class="card">
      <div class="card-body empty-state-sm">暂无评论</div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { getCommentList, updateCommentStatus, deleteComment } from '../../api/comment'
import SkeletonLoader from '../../components/common/SkeletonLoader.vue'
import CustomSelect from '../../components/common/CustomSelect.vue'

const comments = ref([])
const statusFilter = ref('')
const keyword = ref('')
const loading = ref(true)

const statusOptions = [
  { value: '', label: '全部状态' },
  { value: 'pending', label: '待审核' },
  { value: 'approved', label: '已通过' },
  { value: 'rejected', label: '已拒绝' }
]

const formatDate = (d) => d ? new Date(d).toLocaleDateString('zh-CN') : ''
const statusText = (s) => ({ pending: '待审核', approved: '已通过', rejected: '已拒绝' }[s] || s)

const getAvatarStyle = (status) => {
  const colors = {
    pending: 'background: rgba(var(--accent-color-rgb), 0.1); color: var(--accent-color);',
    approved: 'background: rgba(var(--success-color-rgb), 0.1); color: var(--success-color);',
    rejected: 'background: rgba(var(--danger-color-rgb), 0.1); color: var(--danger-color);'
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
  loading.value = false
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
.filter-bar {
  display: flex;
  gap: 12px;
  margin-bottom: 20px;
  align-items: center;
  justify-content: space-between;
}

.comment-user-info { display: flex; flex-direction: column; }
.comment-user-name { font-size: 14px; font-weight: 600; color: var(--card-text-color-main); }
.comment-user-email { font-size: 12px; color: var(--card-text-color-tertiary); }
.comment-meta { font-size: 13px; color: var(--card-text-color-tertiary); display: flex; gap: 6px; }
.btn-hide { background: rgba(var(--warning-color-rgb), 0.08); color: var(--warning-color); }
.btn-hide:hover { background: var(--warning-color); color: white; }
</style>
