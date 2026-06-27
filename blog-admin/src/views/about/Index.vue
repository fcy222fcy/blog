<template>
  <div>
    <div class="card">
      <div class="card-header">
        <div class="card-title">关于我</div>
        <button class="btn btn-primary" @click="showProfileDialog = true" style="margin-left: auto;">
          <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M11 4H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7"></path><path d="M18.5 2.5a2.121 2.121 0 0 1 3 3L12 15l-4 1 1-4 9.5-9.5z"></path></svg>
          编辑资料
        </button>
      </div>
      <div class="card-body">
        <div style="text-align: center; padding: 40px 20px;">
          <div class="sidebar-avatar" style="margin: 0 auto 20px; width: 120px; height: 120px; font-size: 48px;">
            <span>{{ userInfo.nickname?.charAt(0) || 'U' }}</span>
          </div>
          <h2 style="font-size: 24px; font-weight: 700; margin-bottom: 12px; color: var(--card-text-color-main);">{{ userInfo.nickname || '用户' }}</h2>
          <p style="color: var(--card-text-color-secondary); max-width: 500px; margin: 0 auto; line-height: 1.8;">
            {{ userInfo.bio || '暂无简介' }}
          </p>
          <div class="sidebar-social" style="justify-content: center; margin-top: 24px;">
            <a v-for="link in userInfo.socialLinks" :key="link.name" :href="link.url" class="sidebar-social-item" target="_blank" :title="link.name">
              {{ link.icon || '🔗' }}
            </a>
          </div>
        </div>
      </div>
    </div>

    <div class="modal-overlay" :class="{ active: showProfileDialog }" @click.self="showProfileDialog = false">
      <div class="modal" style="max-width: 560px;">
        <div class="modal-header">
          <h3 class="modal-title">编辑个人资料</h3>
          <button class="modal-close" @click="showProfileDialog = false">×</button>
        </div>
        <div class="modal-body" style="max-height: 70vh; overflow-y: auto;">
          <div class="form-group">
            <label class="form-label">昵称</label>
            <input type="text" class="form-input" v-model="profileForm.nickname" placeholder="输入昵称">
          </div>
          <div class="form-group">
            <label class="form-label">个人简介</label>
            <textarea class="form-textarea" v-model="profileForm.bio" placeholder="输入个人简介..." rows="3"></textarea>
          </div>
        </div>
        <div class="modal-footer">
          <button class="btn btn-secondary" @click="showProfileDialog = false">取消</button>
          <button class="btn btn-primary" @click="saveProfile">
            <span>💾</span>
            <span>保存</span>
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { getUserInfo, updateUserInfo } from '../../api/auth'

const showProfileDialog = ref(false)
const userInfo = ref({ nickname: '', bio: '', socialLinks: [] })
const profileForm = ref({ nickname: '', bio: '' })

onMounted(async () => {
  try {
    const res = await getUserInfo()
    if (res.data) {
      userInfo.value = { ...userInfo.value, ...res.data }
      profileForm.value.nickname = res.data.nickname || ''
      profileForm.value.bio = res.data.bio || ''
    }
  } catch (e) { console.error(e) }
})

const saveProfile = async () => {
  try {
    await updateUserInfo(profileForm.value)
    userInfo.value.nickname = profileForm.value.nickname
    userInfo.value.bio = profileForm.value.bio
    showProfileDialog.value = false
    ElMessage.success('保存成功')
  } catch (e) {
    ElMessage.error('保存失败')
  }
}
</script>
