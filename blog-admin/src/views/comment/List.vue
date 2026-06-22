<template>
  <div>
    <el-card>
      <template #header>
        <div style="display: flex; justify-content: space-between; align-items: center">
          <span>评论管理</span>
          <el-select v-model="statusFilter" placeholder="全部状态" clearable @change="loadComments">
            <el-option label="待审核" value="pending" />
            <el-option label="已通过" value="approved" />
            <el-option label="已拒绝" value="rejected" />
          </el-select>
        </div>
      </template>

      <el-table :data="comments" v-loading="loading" stripe>
        <el-table-column prop="nickname" label="用户" width="120" />
        <el-table-column prop="content" label="内容" min-width="200" />
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="statusType(row.status)">{{ statusText(row.status) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="created_at" label="时间" width="180">
          <template #default="{ row }">{{ formatDate(row.created_at) }}</template>
        </el-table-column>
        <el-table-column label="操作" width="200">
          <template #default="{ row }">
            <el-button v-if="row.status !== 'approved'" size="small" type="success" @click="handleStatus(row.id, 'approved')">通过</el-button>
            <el-button v-if="row.status !== 'rejected'" size="small" type="warning" @click="handleStatus(row.id, 'rejected')">拒绝</el-button>
            <el-popconfirm title="确定删除?" @confirm="handleDelete(row.id)">
              <template #reference>
                <el-button size="small" type="danger">删除</el-button>
              </template>
            </el-popconfirm>
          </template>
        </el-table-column>
      </el-table>
    </el-card>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { getCommentList, updateCommentStatus, deleteComment } from '../../api/comment'

const comments = ref([])
const loading = ref(false)
const statusFilter = ref('')

const formatDate = (d) => d ? d.split('T')[0] : ''
const statusType = (s) => ({ pending: 'warning', approved: 'success', rejected: 'danger' }[s] || 'info')
const statusText = (s) => ({ pending: '待审核', approved: '已通过', rejected: '已拒绝' }[s] || s)

const loadComments = async () => {
  loading.value = true
  try {
    const params = { page: 1, page_size: 50 }
    if (statusFilter.value) params.status = statusFilter.value
    const res = await getCommentList(params)
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
    await deleteComment(id)
    ElMessage.success('删除成功')
    loadComments()
  } catch (e) { console.error(e) }
}

onMounted(loadComments)
</script>
