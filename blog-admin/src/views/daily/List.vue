<template>
  <div>
    <el-card>
      <template #header>
        <div style="display: flex; justify-content: space-between; align-items: center">
          <span>每日一问管理</span>
          <el-button type="primary" @click="showDialog()">新建问题</el-button>
        </div>
      </template>

      <el-table :data="questions" v-loading="loading" stripe>
        <el-table-column prop="question" label="问题" min-width="200" show-overflow-tooltip />
        <el-table-column prop="date" label="日期" width="120" />
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="row.status === 1 ? 'success' : 'info'">
              {{ row.status === 1 ? '启用' : '禁用' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="view_count" label="浏览" width="80" />
        <el-table-column prop="like_count" label="点赞" width="80" />
        <el-table-column label="操作" width="200">
          <template #default="{ row }">
            <el-button size="small" @click="showDialog(row)">编辑</el-button>
            <el-button size="small" :type="row.status === 1 ? 'warning' : 'success'" @click="toggleStatus(row)">
              {{ row.status === 1 ? '禁用' : '启用' }}
            </el-button>
            <el-popconfirm title="确定删除?" @confirm="handleDelete(row.id)">
              <template #reference>
                <el-button size="small" type="danger">删除</el-button>
              </template>
            </el-popconfirm>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <el-dialog v-model="dialogVisible" :title="editingId ? '编辑问题' : '新建问题'" width="600px">
      <el-form :model="form" label-width="80px">
        <el-form-item label="日期">
          <el-date-picker v-model="form.date" type="date" value-format="YYYY-MM-DD" placeholder="选择日期" />
        </el-form-item>
        <el-form-item label="问题">
          <el-input v-model="form.question" type="textarea" :rows="3" />
        </el-form-item>
        <el-form-item label="答案">
          <el-input v-model="form.answer" type="textarea" :rows="6" />
        </el-form-item>
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
import { getDailyQuestionList, createDailyQuestion, updateDailyQuestion, deleteDailyQuestion, updateDailyQuestionStatus } from '../../api/daily'

const questions = ref([])
const loading = ref(false)
const dialogVisible = ref(false)
const editingId = ref(null)
const form = ref({ question: '', answer: '', date: '' })

const loadQuestions = async () => {
  loading.value = true
  try {
    const res = await getDailyQuestionList({ page: 1, page_size: 50 })
    questions.value = res.data?.list || []
  } catch (e) { console.error(e) }
  loading.value = false
}

const showDialog = (row) => {
  if (row) {
    editingId.value = row.id
    form.value = { question: row.question, answer: row.answer, date: row.date }
  } else {
    editingId.value = null
    form.value = { question: '', answer: '', date: '' }
  }
  dialogVisible.value = true
}

const handleSave = async () => {
  try {
    if (editingId.value) {
      await updateDailyQuestion(editingId.value, form.value)
    } else {
      await createDailyQuestion(form.value)
    }
    ElMessage.success('保存成功')
    dialogVisible.value = false
    loadQuestions()
  } catch (e) { console.error(e) }
}

const toggleStatus = async (row) => {
  try {
    await updateDailyQuestionStatus(row.id, row.status === 1 ? 0 : 1)
    ElMessage.success('操作成功')
    loadQuestions()
  } catch (e) { console.error(e) }
}

const handleDelete = async (id) => {
  try {
    await deleteDailyQuestion(id)
    ElMessage.success('删除成功')
    loadQuestions()
  } catch (e) { console.error(e) }
}

onMounted(loadQuestions)
</script>
