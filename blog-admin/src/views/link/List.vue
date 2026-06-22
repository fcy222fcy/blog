<template>
  <div>
    <el-card>
      <template #header>
        <div style="display: flex; justify-content: space-between; align-items: center">
          <span>友链管理</span>
          <el-button type="primary" @click="showDialog()">添加链接</el-button>
        </div>
      </template>

      <el-table :data="links" v-loading="loading" stripe>
        <el-table-column prop="name" label="名称" width="150" />
        <el-table-column prop="url" label="链接" min-width="200">
          <template #default="{ row }">
            <a :href="row.url" target="_blank" style="color: var(--el-color-primary)">{{ row.url }}</a>
          </template>
        </el-table-column>
        <el-table-column prop="description" label="描述" />
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="row.status === 'approved' ? 'success' : row.status === 'pending' ? 'warning' : 'danger'">
              {{ { approved: '已通过', pending: '待审核', rejected: '已拒绝' }[row.status] }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="250">
          <template #default="{ row }">
            <el-button v-if="row.status !== 'approved'" size="small" type="success" @click="handleStatus(row.id, 'approved')">通过</el-button>
            <el-button size="small" @click="showDialog(row)">编辑</el-button>
            <el-popconfirm title="确定删除?" @confirm="handleDelete(row.id)">
              <template #reference>
                <el-button size="small" type="danger">删除</el-button>
              </template>
            </el-popconfirm>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <el-dialog v-model="dialogVisible" :title="editingId ? '编辑友链' : '添加友链'" width="500px">
      <el-form :model="form" label-width="80px">
        <el-form-item label="名称"><el-input v-model="form.name" /></el-form-item>
        <el-form-item label="链接"><el-input v-model="form.url" /></el-form-item>
        <el-form-item label="描述"><el-input v-model="form.description" /></el-form-item>
        <el-form-item label="头像"><el-input v-model="form.avatar" /></el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="handleSave">保存</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { getLinkList, createLink, updateLink, deleteLink, updateLinkStatus } from '../../api/link'

const links = ref([])
const loading = ref(false)
const dialogVisible = ref(false)
const editingId = ref(null)
const form = ref({ name: '', url: '', description: '', avatar: '' })

const loadLinks = async () => {
  loading.value = true
  try {
    const res = await getLinkList({ page: 1, page_size: 50 })
    links.value = res.data?.list || []
  } catch (e) { console.error(e) }
  loading.value = false
}

const showDialog = (row) => {
  if (row) {
    editingId.value = row.id
    form.value = { name: row.name, url: row.url, description: row.description, avatar: row.avatar }
  } else {
    editingId.value = null
    form.value = { name: '', url: '', description: '', avatar: '' }
  }
  dialogVisible.value = true
}

const handleSave = async () => {
  try {
    if (editingId.value) {
      await updateLink(editingId.value, form.value)
    } else {
      await createLink(form.value)
    }
    ElMessage.success('保存成功')
    dialogVisible.value = false
    loadLinks()
  } catch (e) { console.error(e) }
}

const handleStatus = async (id, status) => {
  try {
    await updateLinkStatus(id, status)
    ElMessage.success('操作成功')
    loadLinks()
  } catch (e) { console.error(e) }
}

const handleDelete = async (id) => {
  try {
    await deleteLink(id)
    ElMessage.success('删除成功')
    loadLinks()
  } catch (e) { console.error(e) }
}

onMounted(loadLinks)
</script>
