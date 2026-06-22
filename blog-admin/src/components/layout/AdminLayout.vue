<template>
  <el-container class="admin-layout">
    <el-aside width="240px" class="admin-sidebar">
      <div class="sidebar-header">
        <h2>博客管理</h2>
      </div>
      <el-menu :default-active="activeMenu" router>
        <el-menu-item index="/dashboard">
          <el-icon><Monitor /></el-icon>
          <span>仪表盘</span>
        </el-menu-item>
        <el-menu-item index="/articles">
          <el-icon><Document /></el-icon>
          <span>文章管理</span>
        </el-menu-item>
        <el-menu-item index="/categories">
          <el-icon><Folder /></el-icon>
          <span>分类管理</span>
        </el-menu-item>
        <el-menu-item index="/tags">
          <el-icon><PriceTag /></el-icon>
          <span>标签管理</span>
        </el-menu-item>
        <el-menu-item index="/comments">
          <el-icon><ChatDotRound /></el-icon>
          <span>评论管理</span>
        </el-menu-item>
        <el-menu-item index="/links">
          <el-icon><Link /></el-icon>
          <span>友链管理</span>
        </el-menu-item>
        <el-menu-item index="/daily-question">
          <el-icon><QuestionFilled /></el-icon>
          <span>每日一问</span>
        </el-menu-item>
      </el-menu>
    </el-aside>

    <el-container>
      <el-header class="admin-header">
        <div class="header-left">
          <span class="page-title">{{ $route.meta.title }}</span>
        </div>
        <div class="header-right">
          <el-button type="danger" @click="handleLogout">退出登录</el-button>
        </div>
      </el-header>

      <el-main class="admin-main">
        <router-view />
      </el-main>
    </el-container>
  </el-container>
</template>

<script setup>
import { computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'

const route = useRoute()
const router = useRouter()
const activeMenu = computed(() => route.path)

const handleLogout = () => {
  localStorage.removeItem('token')
  router.push('/login')
}
</script>

<style scoped>
.admin-layout { height: 100vh; }
.admin-sidebar { background: #001529; }
.sidebar-header { padding: 20px; color: #fff; text-align: center; }
.sidebar-header h2 { font-size: 18px; font-weight: 600; }
.admin-header { display: flex; align-items: center; justify-content: space-between; background: #fff; border-bottom: 1px solid #eee; padding: 0 20px; }
.page-title { font-size: 16px; font-weight: 600; }
.admin-main { background: #f5f5f5; }
</style>
