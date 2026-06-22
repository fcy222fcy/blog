<template>
  <div>
    <el-card>
      <template #header>
        <div style="display: flex; justify-content: space-between; align-items: center">
          <span>分类管理</span>
          <el-button type="primary" @click="showDialog()">新建分类</el-button>
        </div>
      </template>

      <el-table :data="categories" v-loading="loading" stripe>
        <el-table-column prop="name" label="名称" />
        <el-table-column prop="slug" label="别名" />
        <el-table-column prop="description" label="描述" />
        <el-table-column prop="article_count" label="文章数" width="100" />
        <el-table-column label="操作" width="180">
          <template #default="{ row }">
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

    <el-dialog v-model="dialogVisible" :title="editingId ? '编辑分类' : '新建分类'" width="500px">
      <el-form :model="form" label-width="80px">
        <el-form-item label="名称"><el-input v-model="form.name" /></el-form-item>
        <el-form-item label="别名"><el-input v-model="form.slug" /></el-form-item>
        <el-form-item label="描述"><el-input v-model="form.description" /></el-form-item>
        <el-form-item label="图标"><el-input v-model="form.icon" /></el-form-item>
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
import { getCategoryList, createCategory, updateCategory, deleteCategory } from '../../api/category'

const categories = ref([])
const loading = ref(false)
const dialogVisible = ref(false)
const editingId = ref(null)
const form = ref({ name: '', slug: '', description: '', icon: '' })

const loadCategories = async () => {
  loading.value = true
  try {
    const res = await getCategoryList()
    categories.value = res.data || []
  } catch (e) { console.error(e) }
  loading.value = false
}

const showDialog = (row) => {
  if (row) {
    editingId.value = row.id
    form.value = { name: row.name, slug: row.slug, description: row.description, icon: row.icon }
  } else {
    editingId.value = null
    form.value = { name: '', slug: '', description: '', icon: '' }
  }
  dialogVisible.value = true
}

const handleSave = async () => {
  try {
    if (editingId.value) {
      await updateCategory(editingId.value, form.value)
    } else {
      await createCategory(form.value)
    }
    ElMessage.success('保存成功')
    dialogVisible.value = false
    loadCategories()
  } catch (e) { console.error(e) }
}

const handleDelete = async (id) => {
  try {
    await deleteCategory(id)
    ElMessage.success('删除成功')
    loadCategories()
  } catch (e) { console.error(e) }
}

onMounted(loadCategories)
</script>
