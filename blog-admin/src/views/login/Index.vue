<template>
  <div class="login-page">
    <!-- 顶部导航 -->
    <nav class="login-nav">
      <a href="#" class="login-logo" @click.prevent>
        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <path d="M12 2L2 7l10 5 10-5-10-5z"></path>
          <path d="M2 17l10 5 10-5"></path>
          <path d="M2 12l10 5 10-5"></path>
        </svg>
        <span>Blog Admin</span>
      </a>
    </nav>

    <!-- 主内容区 -->
    <main class="login-main">
      <div class="login-box">
        <div class="login-header">
          <h1 class="login-title">登录</h1>
          <p class="login-subtitle">输入您的账号信息继续</p>
        </div>

        <form @submit.prevent="handleLogin">
          <div class="login-field">
            <label class="login-label">用户名</label>
            <input
              type="text"
              class="login-input"
              v-model.trim="form.username"
              placeholder="请输入用户名"
              autocomplete="username"
              required
            >
          </div>

          <div class="login-field">
            <label class="login-label">密码</label>
            <input
              type="password"
              class="login-input"
              v-model="form.password"
              placeholder="请输入密码"
              autocomplete="current-password"
              required
            >
          </div>

          <div class="login-options">
            <div class="remember-me">
              <input type="checkbox" id="remember" v-model="rememberMe">
              <label for="remember">记住我</label>
            </div>
          </div>

          <button type="submit" class="login-submit" :disabled="loading">
            {{ loading ? '登录中...' : '登录' }}
          </button>
        </form>
      </div>
    </main>

    <!-- 底部 -->
    <footer class="login-footer">
      <p>© {{ year }} Blog Admin. All rights reserved.</p>
    </footer>

    <!-- 主题切换 -->
    <button class="login-theme-toggle" @click="toggleTheme" title="切换主题">
      <svg v-if="isDark" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
        <path d="M21 12.79A9 9 0 1 1 11.21 3 7 7 0 0 0 21 12.79z"></path>
      </svg>
      <svg v-else xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
        <circle cx="12" cy="12" r="5"></circle>
        <line x1="12" y1="1" x2="12" y2="3"></line>
        <line x1="12" y1="21" x2="12" y2="23"></line>
        <line x1="4.22" y1="4.22" x2="5.64" y2="5.64"></line>
        <line x1="18.36" y1="18.36" x2="19.78" y2="19.78"></line>
        <line x1="1" y1="12" x2="3" y2="12"></line>
        <line x1="21" y1="12" x2="23" y2="12"></line>
        <line x1="4.22" y1="19.78" x2="5.64" y2="18.36"></line>
        <line x1="18.36" y1="5.64" x2="19.78" y2="4.22"></line>
      </svg>
    </button>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { login } from '../../api/auth'

const router = useRouter()
const loading = ref(false)
const rememberMe = ref(false)
const form = ref({ username: '', password: '' })
const isDark = ref(false)
const year = new Date().getFullYear()

const initTheme = () => {
  const saved = localStorage.getItem('theme') || 'light'
  document.documentElement.setAttribute('data-scheme', saved)
  isDark.value = saved === 'dark'
}

const toggleTheme = () => {
  isDark.value = !isDark.value
  const scheme = isDark.value ? 'dark' : 'light'
  document.documentElement.setAttribute('data-scheme', scheme)
  localStorage.setItem('theme', scheme)
}

const initRemembered = () => {
  const remembered = localStorage.getItem('remembered_username')
  if (remembered) {
    form.value.username = remembered
    rememberMe.value = true
  }
}

const handleLogin = async () => {
  if (!form.value.username || !form.value.password) {
    ElMessage.warning('请输入用户名和密码')
    return
  }
  loading.value = true
  try {
    const res = await login(form.value)
    localStorage.setItem('token', res.data.token)
    if (rememberMe.value) {
      localStorage.setItem('remembered_username', form.value.username)
    } else {
      localStorage.removeItem('remembered_username')
    }
    ElMessage.success('登录成功')
    router.push('/dashboard')
  } catch (e) {
    // 错误已由 request 拦截器统一处理
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  initTheme()
  initRemembered()
})
</script>

<style scoped>
.login-page {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  background: var(--body-background);
  transition: background-color 0.3s ease;
}

/* 顶部导航 */
.login-nav {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 20px 40px;
}

.login-logo {
  display: flex;
  align-items: center;
  gap: 10px;
  font-size: 18px;
  font-weight: 700;
  color: var(--card-text-color-main);
  text-decoration: none;
}

.login-logo svg {
  width: 28px;
  height: 28px;
}

/* 主内容区 */
.login-main {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 40px 20px;
}

.login-box {
  width: 100%;
  max-width: 360px;
}

.login-header {
  text-align: center;
  margin-bottom: 48px;
}

.login-title {
  font-size: 32px;
  font-weight: 700;
  color: var(--card-text-color-main);
  margin-bottom: 12px;
  letter-spacing: -0.5px;
}

.login-subtitle {
  font-size: 15px;
  color: var(--card-text-color-secondary);
}

.login-field {
  margin-bottom: 24px;
}

.login-label {
  display: block;
  font-size: 13px;
  font-weight: 600;
  color: var(--card-text-color-main);
  margin-bottom: 8px;
  text-transform: uppercase;
  letter-spacing: 0.5px;
}

.login-input {
  width: 100%;
  height: 48px;
  padding: 0 16px;
  background: var(--card-background);
  border: 1px solid var(--card-separator-color);
  border-radius: var(--card-border-radius);
  font-size: 15px;
  color: var(--card-text-color-main);
  transition: all 0.2s ease;
}

.login-input:focus {
  outline: none;
  border-color: var(--accent-color);
}

.login-input::placeholder {
  color: var(--card-text-color-tertiary);
}

.login-options {
  display: flex;
  align-items: center;
  margin-bottom: 32px;
}

.remember-me {
  display: flex;
  align-items: center;
  gap: 8px;
  cursor: pointer;
}

.remember-me input[type="checkbox"] {
  width: 16px;
  height: 16px;
  accent-color: var(--accent-color);
  cursor: pointer;
}

.remember-me label {
  font-size: 14px;
  color: var(--card-text-color-secondary);
  cursor: pointer;
}

.login-submit {
  width: 100%;
  height: 48px;
  background: var(--accent-color);
  color: var(--accent-color-text);
  border: none;
  border-radius: var(--card-border-radius);
  font-size: 15px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s ease;
}

.login-submit:hover:not(:disabled) {
  background: var(--accent-color-darker);
}

.login-submit:active:not(:disabled) {
  transform: scale(0.98);
}

.login-submit:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

/* 底部 */
.login-footer {
  padding: 24px 40px;
  text-align: center;
}

.login-footer p {
  font-size: 13px;
  color: var(--card-text-color-tertiary);
}

/* 主题切换 */
.login-theme-toggle {
  position: fixed;
  top: 20px;
  right: 40px;
  width: 40px;
  height: 40px;
  background: transparent;
  border: 1px solid var(--card-separator-color);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: all 0.2s;
}

.login-theme-toggle:hover {
  border-color: var(--accent-color);
}

.login-theme-toggle svg {
  width: 18px;
  height: 18px;
  color: var(--card-text-color-main);
}

/* 响应式 */
@media (max-width: 480px) {
  .login-nav {
    padding: 16px 20px;
  }

  .login-header {
    margin-bottom: 36px;
  }

  .login-title {
    font-size: 28px;
  }

  .login-footer {
    padding: 20px;
  }

  .login-theme-toggle {
    right: 20px;
  }
}
</style>
