<template>
  <div>
    <div class="card">
      <div class="card-header">
        <div class="card-title">操作日志</div>
        <div class="filter-group">
          <div class="custom-select-wrapper" v-click-outside="closeActionDropdown">
            <div class="custom-select" @click="toggleActionDropdown">
              <span class="select-value">{{ selectedActionLabel }}</span>
              <span class="select-arrow">
                <svg xmlns="http://www.w3.org/2000/svg" width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polyline points="6 9 12 15 18 9"></polyline></svg>
              </span>
            </div>
            <div class="custom-dropdown" v-if="isActionOpen">
              <div class="dropdown-item" :class="{ active: actionFilter === '' }" @click="selectAction('')">全部操作</div>
              <div class="dropdown-item" :class="{ active: actionFilter === 'create' }" @click="selectAction('create')">创建</div>
              <div class="dropdown-item" :class="{ active: actionFilter === 'update' }" @click="selectAction('update')">更新</div>
              <div class="dropdown-item" :class="{ active: actionFilter === 'delete' }" @click="selectAction('delete')">删除</div>
              <div class="dropdown-item" :class="{ active: actionFilter === 'approve' }" @click="selectAction('approve')">审核通过</div>
              <div class="dropdown-item" :class="{ active: actionFilter === 'reject' }" @click="selectAction('reject')">审核拒绝</div>
            </div>
          </div>
          <div class="custom-select-wrapper" v-click-outside="closeTypeDropdown">
            <div class="custom-select" @click="toggleTypeDropdown">
              <span class="select-value">{{ selectedTypeLabel }}</span>
              <span class="select-arrow">
                <svg xmlns="http://www.w3.org/2000/svg" width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polyline points="6 9 12 15 18 9"></polyline></svg>
              </span>
            </div>
            <div class="custom-dropdown" v-if="isTypeOpen">
              <div class="dropdown-item" :class="{ active: typeFilter === '' }" @click="selectType('')">全部类型</div>
              <div class="dropdown-item" :class="{ active: typeFilter === 'article' }" @click="selectType('article')">文章</div>
              <div class="dropdown-item" :class="{ active: typeFilter === 'comment' }" @click="selectType('comment')">评论</div>
              <div class="dropdown-item" :class="{ active: typeFilter === 'category' }" @click="selectType('category')">分类</div>
              <div class="dropdown-item" :class="{ active: typeFilter === 'tag' }" @click="selectType('tag')">标签</div>
              <div class="dropdown-item" :class="{ active: typeFilter === 'link' }" @click="selectType('link')">友链</div>
            </div>
          </div>
        </div>
      </div>

      <div class="card-body" style="padding: 0;">
        <div v-if="loading" style="padding: 60px; text-align: center; color: var(--card-text-color-tertiary);">
          加载中...
        </div>
        <div v-else-if="logs.length === 0" style="padding: 60px; text-align: center; color: var(--card-text-color-tertiary);">
          暂无操作日志
        </div>
        <table v-else class="data-table">
          <thead>
            <tr>
              <th style="width: 60px;">ID</th>
              <th style="width: 80px;">操作人</th>
              <th style="width: 70px;">操作</th>
              <th style="width: 60px;">类型</th>
              <th style="min-width: 120px; max-width: 200px;">目标</th>
              <th style="width: 150px;">时间</th>
              <th style="width: 90px;">IP</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="log in logs" :key="log.id" class="log-row">
              <td class="cell-id"><span class="log-id">#{{ log.id }}</span></td>
              <td class="cell-operator">{{ log.operator_name || '-' }}</td>
              <td class="cell-action">
                <span class="action-badge" :class="'action-' + log.action">{{ actionText(log.action) }}</span>
              </td>
              <td class="cell-type">{{ targetTypeText(log.target_type) }}</td>
              <td class="cell-target">
                <span class="log-target-title">{{ log.target_title || '-' }}</span>
              </td>
              <td class="cell-time">{{ formatTime(log.created_at) }}</td>
              <td class="cell-ip">{{ log.ip }}</td>
            </tr>
          </tbody>
        </table>
      </div>

      <div class="card-body" v-if="total > pageSize">
        <div class="pagination">
          <div class="pagination-info">
            共 <span>{{ total }}</span> 条记录
          </div>
          <div class="pagination-buttons">
            <button class="pagination-btn" :disabled="page <= 1" @click="page--; loadLogs()">上一页</button>
            <button class="pagination-btn active">{{ page }}</button>
            <button class="pagination-btn" :disabled="page * pageSize >= total" @click="page++; loadLogs()">下一页</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { getAuditLogList } from '../../api/audit'

const logs = ref([])
const loading = ref(false)
const page = ref(1)
const pageSize = 20
const total = ref(0)
const actionFilter = ref('')
const typeFilter = ref('')

// 自定义下拉菜单逻辑
const isActionOpen = ref(false)
const isTypeOpen = ref(false)

const selectedActionLabel = computed(() => {
  const map = { '': '全部操作', create: '创建', update: '更新', delete: '删除', approve: '审核通过', reject: '审核拒绝' }
  return map[actionFilter.value] || '全部操作'
})

const selectedTypeLabel = computed(() => {
  const map = { '': '全部类型', article: '文章', comment: '评论', category: '分类', tag: '标签', link: '友链' }
  return map[typeFilter.value] || '全部类型'
})

const toggleActionDropdown = () => {
  isActionOpen.value = !isActionOpen.value
  isTypeOpen.value = false
}

const closeActionDropdown = () => {
  isActionOpen.value = false
}

const selectAction = (value) => {
  actionFilter.value = value
  isActionOpen.value = false
  loadLogs()
}

const toggleTypeDropdown = () => {
  isTypeOpen.value = !isTypeOpen.value
  isActionOpen.value = false
}

const closeTypeDropdown = () => {
  isTypeOpen.value = false
}

const selectType = (value) => {
  typeFilter.value = value
  isTypeOpen.value = false
  loadLogs()
}

const actionText = (a) => {
  const map = { create: '创建', update: '更新', delete: '删除', approve: '通过', reject: '拒绝' }
  return map[a] || a
}

const targetTypeText = (t) => {
  const map = { article: '文章', comment: '评论', category: '分类', tag: '标签', link: '友链' }
  return map[t] || t
}

const formatTime = (d) => {
  if (!d) return ''
  return new Date(d).toLocaleString('zh-CN')
}

const loadLogs = async () => {
  loading.value = true
  try {
    const params = { page: page.value, page_size: pageSize }
    if (actionFilter.value) params.action = actionFilter.value
    if (typeFilter.value) params.target_type = typeFilter.value
    const res = await getAuditLogList(params)
    logs.value = res.data?.list || []
    total.value = res.data?.total || 0
  } catch (e) { console.error(e) }
  loading.value = false
}

onMounted(loadLogs)

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
.filter-group {
  display: flex;
  gap: 12px;
  flex-wrap: wrap;
  align-items: center;
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

.action-badge {
  display: inline-flex;
  align-items: center;
  padding: 4px 10px;
  border-radius: 4px;
  font-size: 12px;
  font-weight: 600;
  white-space: nowrap;
}
.action-create { background: rgba(16, 185, 129, 0.1); color: #10b981; }
.action-update { background: rgba(59, 130, 246, 0.1); color: #3b82f6; }
.action-delete { background: rgba(239, 68, 68, 0.1); color: #ef4444; }
.action-approve { background: rgba(16, 185, 129, 0.1); color: #10b981; }
.action-reject { background: rgba(245, 158, 11, 0.1); color: #f59e0b; }

.data-table {
  width: 100%;
  border-collapse: collapse;
  table-layout: fixed;
}
.data-table th {
  padding: 12px 12px;
  text-align: left;
  border-bottom: 2px solid var(--card-separator-color);
  font-size: 13px;
  font-weight: 600;
  color: var(--card-text-color-secondary);
  background: var(--body-background);
  white-space: nowrap;
}
.data-table td {
  padding: 10px 12px;
  border-bottom: 1px solid var(--card-separator-color);
  font-size: 13px;
  color: var(--card-text-color-main);
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}
.data-table tbody tr {
  height: 48px;
}
.data-table tbody tr:hover {
  background: rgba(var(--accent-color-rgb), 0.02);
}

.cell-id {
  color: var(--card-text-color-tertiary);
  font-size: 12px;
}
.cell-operator {
  font-weight: 500;
}
.cell-action {
  text-align: center;
}
.cell-type {
  color: var(--card-text-color-secondary);
}
.cell-target {
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}
.cell-time {
  font-size: 12px;
  color: var(--card-text-color-secondary);
}
.cell-ip {
  font-size: 12px;
  color: var(--card-text-color-tertiary);
  font-family: monospace;
}

.log-id {
  font-size: 12px;
  color: var(--card-text-color-tertiary);
}
.log-target-title {
  font-size: 13px;
  color: var(--card-text-color-main);
}
</style>
