<template>
  <div>
    <el-row :gutter="16">
      <el-col :span="6">
        <el-card shadow="hover">
          <el-statistic title="文章总数" :value="stats.article_count || 0" />
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card shadow="hover">
          <el-statistic title="已发布" :value="stats.published_count || 0" />
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card shadow="hover">
          <el-statistic title="总浏览量" :value="stats.total_views || 0" />
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card shadow="hover">
          <el-statistic title="友情链接" :value="stats.link_count || 0" />
        </el-card>
      </el-col>
    </el-row>

    <el-row :gutter="16" style="margin-top: 16px">
      <el-col :span="6">
        <el-card shadow="hover">
          <el-statistic title="分类数" :value="stats.category_count || 0" />
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card shadow="hover">
          <el-statistic title="标签数" :value="stats.tag_count || 0" />
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card shadow="hover">
          <el-statistic title="待审核评论" :value="stats.pending_comment_count || 0" />
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card shadow="hover">
          <el-statistic title="每日一问" :value="stats.daily_question_count || 0" />
        </el-card>
      </el-col>
    </el-row>

    <el-card style="margin-top: 16px">
      <template #header>
        <span>快捷操作</span>
      </template>
      <el-space wrap>
        <el-button type="primary" @click="$router.push('/articles')">文章管理</el-button>
        <el-button type="success" @click="$router.push('/articles/edit')">新建文章</el-button>
        <el-button @click="$router.push('/categories')">分类管理</el-button>
        <el-button @click="$router.push('/comments')">评论管理</el-button>
        <el-button @click="$router.push('/daily-question')">每日一问</el-button>
      </el-space>
    </el-card>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import request from '../../api/request'

const stats = ref({})

onMounted(async () => {
  try {
    const res = await request.get('/admin/dashboard/stats')
    stats.value = res.data || {}
  } catch (e) { console.error(e) }
})
</script>
