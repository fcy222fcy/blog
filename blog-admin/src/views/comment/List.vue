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

    <div class="filter-bar">
      <div class="custom-select-wrapper" v-click-outside="closeStatusDropdown">
        <div class="custom-select" @click="toggleStatusDropdown">
          <span class="select-value">{{ selectedStatusLabel }}</span>
          <span class="select-arrow">
            <svg xmlns="http://www.w3.org/2000/svg" width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polyline points="6 9 12 15 18 9"></polyline></svg>
          </span>
        </div>
        <div class="custom-dropdown" v-if="isStatusOpen">
          <div class="dropdown-item" :class="{ active: statusFilter === '' }" @click="selectStatus('')">全部状态</div>
          <div class="dropdown-item" :class="{ active: statusFilter === 'pending' }" @click="selectStatus('pending')">待审核</div>
          <div class="dropdown-item" :class="{ active: statusFilter === 'approved' }" @click="selectStatus('approved')">已通过</div>
          <div class="dropdown-item" :class="{ active: statusFilter === 'rejected' }" @click="selectStatus('rejected')">已拒绝</div>
        </div>
      </div>
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
      </template>
    </div>

    <div v-if="filteredComments.length === 0 && !loading" class="card">
      <div class="card-body" style="text-align: center; color: var(--card-text-color-tertiary);">暂无评论</div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { getCommentList, updateCommentStatus, deleteComment } from '../../api/comment'
import SkeletonLoader from '../../components/common/SkeletonLoader.vue'

const comments = ref([])
const statusFilter = ref('')
const keyword = ref('')
const loading = ref(true)

// 自定义下拉菜单逻辑
const isStatusOpen = ref(false)

const selectedStatusLabel = computed(() => {
  const map = { '': '全部状态', pending: '待审核', approved: '已通过', rejected: '已拒绝' }
  return map[statusFilter.value] || '全部状态'
})

const toggleStatusDropdown = () => {
  isStatusOpen.value = !isStatusOpen.value
}

const closeStatusDropdown = () => {
  isStatusOpen.value = false
}

const selectStatus = (value) => {
  statusFilter.value = value
  isStatusOpen.value = false
}

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

// 自定义指令：点击外部关闭下拉菜单
const vClickOutside = {
  mounted(el, binding) {
    el._clickOutside = (event) => {
      if (!el.contains(event.target)) {
        binding.value()
      }
    }
    document.addEventListener('click', el._clickOutside)
  },
  unmounted(el) {
    document.removeEventListener('click', el._clickOutside)
  }
}
</script>

<style scoped>
.filter-bar {
  display: flex;
  gap: 12px;
  margin-bottom: 20px;
  align-items: center;
  justify-content: space-between;
}

.custom-select-wrapper {
  position: relative;
  display: inline-flex;
}

.custom-select {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 8px 32px 8px 12px;
  border: 1px solid var(--card-separator-color);
  border-radius: var(--card-border-radius);
  font-size: 13px;
  font-weight: 500;
  color: var(--card-text-color-main);
  background: var(--card-background);
  cursor: pointer;
  transition: all 0.15s ease;
  min-width: 100px;
}

.custom-select:hover {
  border-color: var(--accent-color);
}

.select-value {
  flex: 1;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.select-arrow {
  position: absolute;
  right: 10px;
  top: 50%;
  transform: translateY(-50%);
  pointer-events: none;
  color: var(--card-text-color-tertiary);
  display: flex;
  align-items: center;
}

.custom-dropdown {
  position: absolute;
  top: calc(100% + 4px);
  left: 0;
  right: 0;
  min-width: 200px;
  background: var(--card-background);
  border: 1px solid var(--card-separator-color);
  border-radius: var(--card-border-radius);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  z-index: 100;
}

.dropdown-item {
  padding: 10px 12px;
  font-size: 13px;
  color: var(--card-text-color-main);
  cursor: pointer;
  transition: all 0.15s ease;
}

.dropdown-item:hover {
  background: rgba(var(--accent-color-rgb), 0.06);
  color: var(--accent-color);
}

.dropdown-item.active {
  background: rgba(var(--accent-color-rgb), 0.1);
  color: var(--accent-color);
  font-weight: 600;
}

.comment-user-info { display: flex; flex-direction: column; }
.comment-user-name { font-size: 14px; font-weight: 600; color: var(--card-text-color-main); }
.comment-user-email { font-size: 12px; color: var(--card-text-color-tertiary); }
.comment-meta { font-size: 13px; color: var(--card-text-color-tertiary); display: flex; gap: 6px; }
.btn-hide { background: rgba(245, 158, 11, 0.08); color: var(--warning-color); }
.btn-hide:hover { background: var(--warning-color); color: white; }
</style>
