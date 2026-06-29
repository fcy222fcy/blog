<template>
  <div class="comment-section">
    <!-- 评论表单卡片 -->
    <div class="comment-form-card">
      <form class="comment-form" @submit.prevent="submitComment">
        <div class="form-header">
          <div class="form-field">
            <label>昵称</label>
            <input
              v-model="form.nickname"
              type="text"
              placeholder="昵称"
              required
            />
          </div>
          <div class="form-field">
            <label>邮箱(可选)</label>
            <input
              v-model="form.email"
              type="email"
              placeholder="邮箱"
            />
          </div>
        </div>
        <div class="form-body">
          <textarea
            ref="textareaRef"
            v-model="form.content"
            placeholder="欢迎评论，填写邮箱可收到回复提醒~"
            rows="4"
            required
          ></textarea>
          <div v-if="form.content" class="form-preview">
            <div class="preview-label">预览:</div>
            <div class="preview-content" v-html="renderMarkdown(form.content)"></div>
          </div>
        </div>
        <div class="form-footer">
          <div class="form-tools">
            <button type="button" class="tool-btn md-btn" title="Markdown" @click="insertMarkdown">
              <svg width="24" height="24" viewBox="0 0 16 16">
                <path d="M14.85 3H1.15C.52 3 0 3.52 0 4.15v7.69C0 12.48.52 13 1.15 13h13.69c.64 0 1.15-.52 1.15-1.15v-7.7C16 3.52 15.48 3 14.85 3zM9 11H7V8L5.5 9.92 4 8v3H2V5h2l1.5 2L7 5h2v6zm2.99.5L9.5 8H11V5h2v3h1.5l-2.51 3.5z" fill="currentColor"/>
              </svg>
            </button>
            <div class="emoji-wrapper">
              <button type="button" class="tool-btn emoji-btn" title="表情" @click="toggleEmojiPanel">
                <svg width="24" height="24" viewBox="0 0 1024 1024">
                  <path d="M563.2 463.3 677 540c1.7 1.2 3.7 1.8 5.8 1.8.7 0 1.4-.1 2-.2 2.7-.5 5.1-2.1 6.6-4.4l25.3-37.8c1.5-2.3 2.1-5.1 1.6-7.8s-2.1-5.1-4.4-6.6l-73.6-49.1 73.6-49.1c2.3-1.5 3.9-3.9 4.4-6.6.5-2.7 0-5.5-1.6-7.8l-25.3-37.8a10.1 10.1 0 0 0-6.6-4.4c-.7-.1-1.3-.2-2-.2-2.1 0-4.1.6-5.8 1.8l-113.8 76.6c-9.2 6.2-14.7 16.4-14.7 27.5.1 11 5.5 21.3 14.7 27.4zM387 348.8h-45.5c-5.7 0-10.4 4.7-10.4 10.4v153.3c0 5.7 4.7 10.4 10.4 10.4H387c5.7 0 10.4-4.7 10.4-10.4V359.2c0-5.7-4.7-10.4-10.4-10.4zm333.8 241.3-41-20a10.3 10.3 0 0 0-8.1-.5c-2.6.9-4.8 2.9-5.9 5.4-30.1 64.9-93.1 109.1-164.4 115.2-5.7.5-9.9 5.5-9.5 11.2l3.9 45.5c.5 5.3 5 9.5 10.3 9.5h.9c94.8-8 178.5-66.5 218.6-152.7 2.4-5 .3-11.2-4.8-13.6zm186-186.1c-11.9-42-30.5-81.4-55.2-117.1-24.1-34.9-53.5-65.6-87.5-91.2-33.9-25.6-71.5-45.5-111.6-59.2-41.2-14-84.1-21.1-127.8-21.1h-1.2c-75.4 0-148.8 21.4-212.5 61.7-63.7 40.3-114.3 97.6-146.5 165.8-32.2 68.1-44.3 143.6-35.1 218.4 9.3 74.8 39.4 145 87.3 203.3.1.2.3.3.4.5l36.2 38.4c1.1 1.2 2.5 2.1 3.9 2.6 73.3 66.7 168.2 103.5 267.5 103.5 73.3 0 145.2-20.3 207.7-58.7 37.3-22.9 70.3-51.5 98.1-85 27.1-32.7 48.7-69.5 64.2-109.1 15.5-39.7 24.4-81.3 26.6-123.8 2.4-43.6-2.5-87-14.5-129zm-60.5 181.1c-8.3 37-22.8 72-43 104-19.7 31.1-44.3 58.6-73.1 81.7-28.8 23.1-61 41-95.7 53.4-35.6 12.7-72.9 19.1-110.9 19.1-82.6 0-161.7-30.6-222.8-86.2l-34.1-35.8c-23.9-29.3-42.4-62.2-55.1-97.7-12.4-34.7-18.8-71-19.2-107.9-.4-36.9 5.4-73.3 17.1-108.2 12-35.8 30-69.2 53.4-99.1 31.7-40.4 71.1-72 117.2-94.1 44.5-21.3 94-32.6 143.4-32.6 49.3 0 97 10.8 141.8 32 34.3 16.3 65.3 38.1 92 64.8 26.1 26 47.5 56 63.6 89.2 16.2 33.2 26.6 68.5 31 105.1 4.6 37.5 2.7 75.3-5.6 112.3z" fill="currentColor"/>
                </svg>
              </button>
              <!-- 表情选择面板 -->
              <div v-if="showEmojiPanel" class="emoji-panel">
                <EmojiPicker @select="onEmojiSelect" />
              </div>
            </div>
            <label class="tool-btn img-btn" title="图片">
              <svg width="24" height="24" viewBox="0 0 1024 1024">
                <path d="M784 112H240c-88 0-160 72-160 160v480c0 88 72 160 160 160h544c88 0 160-72 160-160V272c0-88-72-160-160-160zm96 640c0 52.8-43.2 96-96 96H240c-52.8 0-96-43.2-96-96V272c0-52.8 43.2-96 96-96h544c52.8 0 96 43.2 96 96v480z" fill="currentColor"/>
                <path d="M352 480c52.8 0 96-43.2 96-96s-43.2-96-96-96-96 43.2-96 96 43.2 96 96 96zm0-128c17.6 0 32 14.4 32 32s-14.4 32-32 32-32-14.4-32-32 14.4-32 32-32zm462.4 379.2-3.2-3.2-177.6-177.6c-25.6-25.6-65.6-25.6-91.2 0l-80 80-36.8-36.8c-25.6-25.6-65.6-25.6-91.2 0L200 728c-4.8 6.4-8 14.4-8 24 0 17.6 14.4 32 32 32 9.6 0 16-3.2 22.4-9.6L380.8 640l134.4 134.4c6.4 6.4 14.4 9.6 24 9.6 17.6 0 32-14.4 32-32 0-9.6-4.8-17.6-9.6-24l-52.8-52.8 80-80L769.6 776c6.4 4.8 12.8 8 20.8 8 17.6 0 32-14.4 32-32 0-8-3.2-16-8-20.8z" fill="currentColor"/>
              </svg>
            </label>
            <button type="button" class="tool-btn link-btn" title="链接" @click="insertLink">
              <svg width="24" height="24" viewBox="0 0 1024 1024">
                <path d="M710.816 654.301c70.323-96.639 61.084-230.578-23.705-314.843-46.098-46.098-107.183-71.109-172.28-71.109-65.008 0-126.092 25.444-172.28 71.109-45.227 46.098-70.756 107.183-70.756 172.106 0 64.923 25.444 126.007 71.194 172.106 46.099 46.098 107.184 71.109 172.28 71.109 51.414 0 100.648-16.212 142.824-47.404l126.53 126.006c7.058 7.06 16.297 10.979 26.406 10.979 10.105 0 19.343-3.919 26.402-10.979 14.467-14.467 14.467-38.172 0-52.723L710.816 654.301zm-315.107-23.265c-65.88-65.88-65.88-172.54 0-238.42 32.069-32.07 74.245-49.149 119.471-49.149 45.227 0 87.407 17.603 119.472 49.149 65.88 65.879 65.88 172.539 0 238.42-63.612 63.178-175.242 63.178-238.943 0z" fill="currentColor"/>
              </svg>
            </button>
          </div>
          <div class="form-actions">
            <span class="char-count">{{ form.content.length }} 字</span>
            <button type="button" class="login-btn" v-if="!isLoggedIn">登录</button>
            <button type="submit" class="submit-btn" :disabled="submitting || !form.content.trim()">
              提交
            </button>
          </div>
        </div>
      </form>
    </div>

    <!-- 评论列表 -->
    <div class="comment-list-wrapper">
      <div class="comment-header-bar">
        <h3 class="comment-count">{{ comments.length }} 评论</h3>
        <div class="sort-options">
          <button
            :class="['sort-btn', { active: sortBy === 'asc' }]"
            @click="sortBy = 'asc'"
          >按正序</button>
          <button
            :class="['sort-btn', { active: sortBy === 'desc' }]"
            @click="sortBy = 'desc'"
          >按倒序</button>
          <button
            :class="['sort-btn', { active: sortBy === 'hot' }]"
            @click="sortBy = 'hot'"
          >按热度</button>
        </div>
      </div>

      <div v-if="loading" class="loading">加载中...</div>

      <div v-else-if="comments.length === 0" class="empty">暂无评论，快来抢沙发吧！</div>

      <div v-else class="comment-list">
        <div v-for="comment in sortedComments" :key="comment.id" class="comment-item">
          <div class="comment-main">
            <div class="avatar">
              <img v-if="comment.email" :src="getAvatar(comment.email)" :alt="comment.nickname" />
              <span v-else>{{ comment.nickname.charAt(0).toUpperCase() }}</span>
            </div>
            <div class="comment-body">
              <div class="comment-meta">
                <span class="nickname">{{ comment.nickname }}</span>
                <span class="time">{{ formatTime(comment.created_at) }}</span>
              </div>
              <div class="comment-content">{{ comment.content }}</div>
              <div class="comment-actions">
                <button class="action-btn like-btn" @click="likeComment(comment)">
                  <span class="icon"><svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M20.84 4.61a5.5 5.5 0 0 0-7.78 0L12 5.67l-1.06-1.06a5.5 5.5 0 0 0-7.78 7.78l1.06 1.06L12 21.23l7.78-7.78 1.06-1.06a5.5 5.5 0 0 0 0-7.78z"></path></svg></span>
                </button>
                <button class="action-btn reply-btn" @click="startReply(comment)">
                  <span class="icon"><svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z"></path></svg></span>
                </button>
              </div>
            </div>
          </div>

          <!-- 回复表单 -->
          <div v-if="replyingTo?.id === comment.id" class="reply-form">
            <textarea
              v-model="replyContent"
              :placeholder="`@${comment.nickname}: `"
              rows="3"
            ></textarea>
            <div class="reply-actions">
              <button class="cancel-btn" @click="cancelReply">取消</button>
              <button class="submit-btn" @click="submitReply" :disabled="submitting || !replyContent.trim()">
                提交
              </button>
            </div>
          </div>

          <!-- 子评论（回复） -->
          <div v-if="comment.replies && comment.replies.length > 0" class="reply-list">
            <div v-for="reply in comment.replies" :key="reply.id" class="reply-item">
              <div class="comment-main small">
                <div class="avatar small">
                  <img v-if="reply.email" :src="getAvatar(reply.email)" :alt="reply.nickname" />
                  <span v-else>{{ reply.nickname.charAt(0).toUpperCase() }}</span>
                </div>
                <div class="comment-body">
                  <div class="comment-meta">
                    <span class="nickname">{{ reply.nickname }}</span>
                    <span v-if="reply.is_admin" class="admin-badge">博主</span>
                    <span class="time">{{ formatTime(reply.created_at) }}</span>
                  </div>
                  <div class="comment-content">
                    <span v-if="reply.reply_to" class="reply-to">@{{ reply.reply_to }}: </span>
                    {{ reply.content }}
                  </div>
                  <div class="comment-actions">
                    <button class="action-btn like-btn" @click="likeComment(reply)">
                      <span class="icon"><svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M20.84 4.61a5.5 5.5 0 0 0-7.78 0L12 5.67l-1.06-1.06a5.5 5.5 0 0 0-7.78 7.78l1.06 1.06L12 21.23l7.78-7.78 1.06-1.06a5.5 5.5 0 0 0 0-7.78z"></path></svg></span>
                    </button>
                    <button class="action-btn reply-btn" @click="startReply(comment, reply)">
                      <span class="icon"><svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z"></path></svg></span>
                    </button>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { useRoute } from 'vue-router'
import { getCommentsByArticle, createComment } from '../../api/comment'
import EmojiPicker from 'vue3-emoji-picker'
import 'vue3-emoji-picker/css'
import { getAvatarUrl } from '../../utils/avatar'
import { marked } from 'marked'

const route = useRoute()
const textareaRef = ref(null)

const comments = ref([])
const loading = ref(false)
const submitting = ref(false)
const replyingTo = ref(null)
const replyContent = ref('')
const sortBy = ref('desc')
const isLoggedIn = ref(false)
const showEmojiPanel = ref(false)

const form = ref({
  nickname: localStorage.getItem('comment_nickname') || '',
  email: localStorage.getItem('comment_email') || '',
  content: ''
})

// 排序后的评论
const sortedComments = computed(() => {
  const list = [...comments.value]
  if (sortBy.value === 'asc') {
    return list.sort((a, b) => new Date(a.created_at) - new Date(b.created_at))
  } else if (sortBy.value === 'desc') {
    return list.sort((a, b) => new Date(b.created_at) - new Date(a.created_at))
  } else {
    // 按热度（点赞数）
    return list.sort((a, b) => (b.like_count || 0) - (a.like_count || 0))
  }
})

// 切换表情面板
const toggleEmojiPanel = () => {
  showEmojiPanel.value = !showEmojiPanel.value
}

// 插入表情
const insertEmoji = (emoji) => {
  const textarea = textareaRef.value
  if (textarea) {
    const start = textarea.selectionStart
    const end = textarea.selectionEnd
    const content = form.value.content
    form.value.content = content.substring(0, start) + emoji + content.substring(end)
    // 光标移到插入位置后面
    setTimeout(() => {
      textarea.selectionStart = textarea.selectionEnd = start + emoji.length
      textarea.focus()
    }, 0)
  } else {
    form.value.content += emoji
  }
  showEmojiPanel.value = false
}

// EmojiPicker 选择回调
const onEmojiSelect = (emoji) => {
  insertEmoji(emoji)
}

// 根据邮箱获取头像
const getAvatar = (email) => {
  return getAvatarUrl(email, 80)
}

// Markdown 渲染
const renderMarkdown = (text) => {
  if (!text) return ''
  return marked.parse(text)
}

// 插入 Markdown 格式
const insertMarkdown = () => {
  const textarea = textareaRef.value
  if (textarea) {
    const start = textarea.selectionStart
    const end = textarea.selectionEnd
    const selectedText = form.value.content.substring(start, end)
    const markdownText = selectedText ? `**${selectedText}**` : '**粗体文本**'
    form.value.content = form.value.content.substring(0, start) + markdownText + form.value.content.substring(end)
    setTimeout(() => {
      if (selectedText) {
        textarea.selectionStart = start
        textarea.selectionEnd = start + markdownText.length
      } else {
        textarea.selectionStart = start + 2
        textarea.selectionEnd = start + 6
      }
      textarea.focus()
    }, 0)
  }
}

// 插入图片
const insertImage = () => {
  const url = prompt('请输入图片URL:')
  if (url) {
    const textarea = textareaRef.value
    if (textarea) {
      const start = textarea.selectionStart
      const imageText = `![图片](${url})`
      form.value.content = form.value.content.substring(0, start) + imageText + form.value.content.substring(start)
      setTimeout(() => {
        textarea.selectionStart = textarea.selectionEnd = start + imageText.length
        textarea.focus()
      }, 0)
    } else {
      form.value.content += `![图片](${url})`
    }
  }
}

// 插入链接
const insertLink = () => {
  const url = prompt('请输入链接URL:')
  if (url) {
    const textarea = textareaRef.value
    if (textarea) {
      const start = textarea.selectionStart
      const end = textarea.selectionEnd
      const selectedText = form.value.content.substring(start, end) || '链接文本'
      const linkText = `[${selectedText}](${url})`
      form.value.content = form.value.content.substring(0, start) + linkText + form.value.content.substring(end)
      setTimeout(() => {
        textarea.selectionStart = start + 1
        textarea.selectionEnd = start + 1 + selectedText.length
        textarea.focus()
      }, 0)
    } else {
      form.value.content += `[链接文本](${url})`
    }
  }
}

// 点击外部关闭表情面板
const handleClickOutside = (e) => {
  if (!e.target.closest('.emoji-wrapper')) {
    showEmojiPanel.value = false
  }
}

// 获取文章评论
const fetchComments = async () => {
  loading.value = true
  try {
    const articleId = route.params.id || route.params.slug
    const res = await getCommentsByArticle(articleId, { page: 1, page_size: 100 })
    comments.value = res.data.list || []
  } catch (error) {
    console.error('获取评论失败:', error)
  } finally {
    loading.value = false
  }
}

// 提交评论
const submitComment = async () => {
  if (!form.value.nickname || !form.value.content.trim()) return

  submitting.value = true
  try {
    const articleSlug = route.params.slug || route.params.id
    await createComment({
      article_slug: articleSlug,
      nickname: form.value.nickname,
      email: form.value.email,
      content: form.value.content
    })

    // 保存昵称和邮箱到本地存储
    localStorage.setItem('comment_nickname', form.value.nickname)
    if (form.value.email) {
      localStorage.setItem('comment_email', form.value.email)
    }

    // 清空表单
    form.value.content = ''
    // 重新获取评论
    await fetchComments()
  } catch (error) {
    console.error('提交评论失败:', error)
    alert('评论提交失败，请稍后重试')
  } finally {
    submitting.value = false
  }
}

// 开始回复
const startReply = (comment, reply = null) => {
  replyingTo.value = comment
  replyContent.value = reply ? `@${reply.nickname} ` : ''
}

// 取消回复
const cancelReply = () => {
  replyingTo.value = null
  replyContent.value = ''
}

// 提交回复
const submitReply = async () => {
  if (!replyContent.value.trim() || !replyingTo.value) return

  submitting.value = true
  try {
    const articleSlug = route.params.slug || route.params.id
    await createComment({
      article_slug: articleSlug,
      nickname: form.value.nickname || '匿名用户',
      email: form.value.email,
      content: replyContent.value,
      parent_id: replyingTo.value.id
    })

    // 清空回复
    cancelReply()
    // 重新获取评论
    await fetchComments()
  } catch (error) {
    console.error('提交回复失败:', error)
    alert('回复提交失败，请稍后重试')
  } finally {
    submitting.value = false
  }
}

// 点赞评论
const likeComment = (comment) => {
  // TODO: 实现点赞功能
  comment.like_count = (comment.like_count || 0) + 1
}

// 格式化时间
const formatTime = (dateStr) => {
  if (!dateStr) return ''
  const date = new Date(dateStr)
  const now = new Date()
  const diff = now - date

  // 小于1分钟
  if (diff < 60 * 1000) return '刚刚'
  // 小于1小时
  if (diff < 60 * 60 * 1000) return Math.floor(diff / (60 * 1000)) + ' 分钟前'
  // 小于24小时
  if (diff < 24 * 60 * 60 * 1000) return Math.floor(diff / (60 * 60 * 1000)) + ' 小时前'
  // 小于30天
  if (diff < 30 * 24 * 60 * 60 * 1000) return Math.floor(diff / (24 * 60 * 60 * 1000)) + ' 天前'
  // 其他
  return dateStr.split('T')[0]
}

onMounted(() => {
  fetchComments()
  document.addEventListener('click', handleClickOutside)
})

onUnmounted(() => {
  document.removeEventListener('click', handleClickOutside)
})
</script>

<style scoped>
.comment-section {
  margin-top: 32px;
  background: #fdfdfb;
  border: 1px solid transparent;
  border-radius: 10px;
  box-shadow: 0 1px 3px rgba(0,0,0,0.12), 0 1px 2px rgba(0,0,0,0.24);
  padding: 24px;
}

/* 评论表单内嵌卡片 */
.comment-form-card {
  background: #fff;
  border: 1px solid rgba(0, 0, 0, 0.08);
  border-radius: 8px;
  overflow: hidden;
  margin-bottom: 24px;
}

.comment-form {
  display: flex;
  flex-direction: column;
}

.form-header {
  display: flex;
  padding: 16px 20px;
  border-bottom: 1px dashed #e8e8e8;
}

.form-field {
  flex: 1;
  display: flex;
  align-items: center;
  gap: 8px;
}

.form-field:first-child {
  border-right: 1px dashed #e8e8e8;
  padding-right: 16px;
  margin-right: 16px;
}

.form-field label {
  font-size: 0.9rem;
  color: #666;
  white-space: nowrap;
}

.form-field input {
  flex: 1;
  padding: 8px 12px;
  border: none;
  background: transparent;
  color: #333;
  font-size: 0.95rem;
}

.form-field input:focus {
  outline: none;
}

.form-field input::placeholder {
  color: #999;
}

.form-body {
  padding: 16px 20px;
}

.form-body textarea {
  width: 100%;
  border: none;
  background: transparent;
  color: #333;
  font-size: 0.95rem;
  resize: none;
  min-height: 100px;
  font-family: inherit;
}

.form-body textarea:focus {
  outline: none;
}

.form-body textarea::placeholder {
  color: #999;
}

/* 预览区 */
.form-preview {
  margin-top: 12px;
  padding-top: 12px;
  border-top: 1px dashed #e8e8e8;
}

.preview-label {
  font-size: 0.85rem;
  color: #999;
  margin-bottom: 8px;
  font-weight: 500;
}

.preview-content {
  font-size: 0.9rem;
  color: #333;
  line-height: 1.6;
}

.preview-content p {
  margin: 0 0 8px;
}

.preview-content p:last-child {
  margin-bottom: 0;
}

.form-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px 20px;
  border-top: 1px solid rgba(0, 0, 0, 0.08);
}

.form-tools {
  display: flex;
  gap: 4px;
}

.tool-btn {
  width: 36px;
  height: 36px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: none;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  color: #666;
  transition: all 0.2s;
}

.tool-btn:hover {
  color: #10b981;
}

.tool-btn:hover {
  background: rgba(16, 185, 129, 0.1);
}

/* 表情面板 */
.emoji-wrapper {
  position: relative;
}

.emoji-panel {
  position: absolute;
  bottom: 100%;
  left: 0;
  margin-bottom: 8px;
  background: #ffffff;
  border: 1px solid #e8e8e8;
  border-radius: 8px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  z-index: 100;
}

:deep(.em-emoji-picker) {
  --border-radius: 8px;
  height: 350px !important;
  width: 350px !important;
}

.form-actions {
  display: flex;
  align-items: center;
  gap: 12px;
}

.char-count {
  font-size: 0.85rem;
  color: #999;
}

.login-btn {
  padding: 8px 16px;
  background: transparent;
  color: #666;
  border: 1px solid #e8e8e8;
  border-radius: 6px;
  font-size: 0.9rem;
  cursor: pointer;
  transition: all 0.2s;
}

.login-btn:hover {
  background: #f5f5f5;
}

.submit-btn {
  padding: 8px 20px;
  background: #10b981;
  color: white;
  border: none;
  border-radius: 6px;
  font-size: 0.9rem;
  cursor: pointer;
  transition: opacity 0.2s;
  font-weight: 500;
}

.submit-btn:hover:not(:disabled) {
  opacity: 0.9;
}

.submit-btn:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

/* 评论列表 */
.comment-list-wrapper {
  padding: 0;
}

.comment-list-wrapper .comment-header-bar {
  padding: 0;
}

.comment-header-bar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24px;
}

.comment-count {
  font-size: 1.2rem;
  font-weight: 600;
  color: #333;
  margin: 0;
}

.sort-options {
  display: flex;
  gap: 16px;
}

.sort-btn {
  background: none;
  border: none;
  color: #999;
  font-size: 0.9rem;
  cursor: pointer;
  padding: 4px 8px;
  border-radius: 4px;
  transition: all 0.2s;
}

.sort-btn:hover {
  color: #10b981;
}

.sort-btn.active {
  color: #10b981;
  font-weight: 500;
}

.loading,
.empty {
  text-align: center;
  padding: 40px;
  color: #999;
}

.comment-list {
  display: flex;
  flex-direction: column;
}

.comment-item {
  padding: 16px 0;
  border-bottom: 1px solid rgba(0, 0, 0, 0.08);
}

.comment-item:last-child {
  border-bottom: none;
}

.comment-main {
  display: flex;
  gap: 12px;
}

.comment-main.small {
  gap: 10px;
}

.avatar {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  background: linear-gradient(135deg, #10b981 0%, #059669 100%);
  color: white;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: 600;
  font-size: 1rem;
  flex-shrink: 0;
  overflow: hidden;
}

.avatar.small {
  width: 32px;
  height: 32px;
  font-size: 0.8rem;
}

.avatar img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.comment-body {
  flex: 1;
  min-width: 0;
}

.comment-meta {
  display: flex;
  align-items: center;
  gap: 8px;
  margin-bottom: 6px;
}

.nickname {
  font-weight: 600;
  color: #333;
  font-size: 0.9rem;
}

.admin-badge {
  display: inline-block;
  padding: 1px 6px;
  background: #10b981;
  color: white;
  font-size: 0.7rem;
  border-radius: 3px;
  font-weight: 500;
}

.time {
  font-size: 0.8rem;
  color: #999;
}

.comment-content {
  color: #333;
  line-height: 1.6;
  margin-bottom: 8px;
  word-wrap: break-word;
  font-size: 0.95rem;
}

.reply-to {
  color: #10b981;
  font-weight: 500;
}

.comment-actions {
  display: flex;
  gap: 12px;
}

.action-btn {
  display: flex;
  align-items: center;
  gap: 4px;
  background: none;
  border: none;
  color: #999;
  font-size: 0.8rem;
  cursor: pointer;
  padding: 2px 4px;
  border-radius: 4px;
  transition: all 0.2s;
}

.action-btn:hover {
  color: #10b981;
}

.action-btn .icon {
  display: flex;
  align-items: center;
}

.action-btn .icon svg {
  width: 16px;
  height: 16px;
}

.like-btn:hover {
  color: #e74c3c;
}

/* 回复表单 */
.reply-form {
  margin-top: 12px;
  margin-left: 52px;
}

.reply-form textarea {
  width: 100%;
  padding: 10px 12px;
  border: 1px solid #e8e8e8;
  border-radius: 6px;
  background: #ffffff;
  color: #333;
  font-size: 0.9rem;
  resize: vertical;
  min-height: 80px;
  box-sizing: border-box;
  font-family: inherit;
}

.reply-form textarea:focus {
  outline: none;
  border-color: #10b981;
}

.reply-form textarea::placeholder {
  color: #999;
}

.reply-actions {
  display: flex;
  justify-content: flex-end;
  gap: 10px;
  margin-top: 10px;
}

.cancel-btn {
  padding: 6px 14px;
  background: transparent;
  color: #666;
  border: 1px solid #e8e8e8;
  border-radius: 4px;
  font-size: 0.85rem;
  cursor: pointer;
  transition: all 0.2s;
}

.cancel-btn:hover {
  background: #f5f5f5;
}

.reply-form .submit-btn {
  padding: 6px 14px;
  font-size: 0.85rem;
  background: #10b981;
}

/* 子评论列表 */
.reply-list {
  margin-top: 12px;
  margin-left: 52px;
  padding-left: 12px;
  border-left: 2px solid #f0f0f0;
}

.reply-item {
  padding: 10px 0;
}

.reply-item:not(:last-child) {
  border-bottom: 1px solid #f5f5f5;
}

@media (max-width: 768px) {
  .form-header {
    flex-direction: column;
    gap: 12px;
  }

  .form-field:first-child {
    border-right: none;
    padding-right: 0;
    margin-right: 0;
    padding-bottom: 12px;
    border-bottom: 1px dashed #e8e8e8;
  }

  .comment-header-bar {
    flex-direction: column;
    align-items: flex-start;
    gap: 12px;
  }

  .reply-form,
  .reply-list {
    margin-left: 44px;
  }

  .emoji-panel {
    width: 280px;
  }

  .emoji-list {
    grid-template-columns: repeat(7, 1fr);
  }
}
</style>
