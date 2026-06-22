<template>
  <div>
    <el-card>
      <template #header>
        <div style="display: flex; justify-content: space-between; align-items: center">
          <span>文章列表</span>
          <el-button type="primary" @click="$router.push('/articles/edit')">新建文章</el-button>
        </div>
      </template>

      <el-table :data="articles" v-loading="loading" stripe>
        <el-table-column prop="title" label="标题" min-width="200" />
        <el-table-column prop="category.name" label="分类" width="120" />
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="row.status === 'published' ? 'success' : 'info'">
              {{ row.status === 'published' ? '已发布' : '草稿' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="view_count" label="浏览量" width="100" />
        <el-table-column prop="created_at" label="创建时间" width="180">
          <template #default="{ row }">{{ formatDate(row.created_at) }}</template>
        </el-table-column>
        <el-table-column label="操作" width="180" fixed="right">
          <template #default="{ row }">
            <el-button size="small" @click="$router.push('/articles/edit/' + row.id)">编辑</el-button>
            <el-popconfirm title="确定删除?" @confirm="handleDelete(row.id)">
              <template #reference>
                <el-button size="small" type="danger">删除</el-button>
              </template>
            </el-popconfirm>
          </template>
        </el-table-column>
      </el-table>

      <el-pagination
        v-model:current-page="page"
        :page-size="10"
        :total="total"
        layout="total, prev, pager, next"
        style="margin-top: 16px; justify-content: flex-end"
        @current-change="loadArticles"
      />
    </el-card>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { getArticleList, deleteArticle } from '../../api/article'
import { ElMessage } from 'element-plus'

const articles = ref([])
const loading = ref(false)
const page = ref(1)
const total = ref(0)

const formatDate = (d) => d ? d.split('T')[0] : ''

const loadArticles = async () => {
  loading.value = true
  try {
    const res = await getArticleList({ page: page.value, page_size: 10 })
    articles.value = res.data?.list || []
    total.value = res.data?.total || 0
  } catch (e) { console.error(e) }
  loading.value = false
}

const handleDelete = async (id) => {
  try {
    await deleteArticle(id)
    ElMessage.success('删除成功')
    loadArticles()
  } catch (e) { console.error(e) }
}

onMounted(loadArticles)
</script>
