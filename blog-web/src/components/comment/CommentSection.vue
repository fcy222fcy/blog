<template>
  <div class="comment-section">
    <!-- 评论表单 -->
    <div class="comment-form-wrapper">
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
            v-model="form.content"
            placeholder="欢迎评论，填写邮箱可收到回复提醒~"
            rows="4"
            required
          ></textarea>
        </div>
        <div class="form-footer">
          <div class="form-tools">
            <button type="button" class="tool-btn" title="Markdown">✍️</button>
            <button type="button" class="tool-btn" title="表情">😀</button>
            <button type="button" class="tool-btn" title="图片">🖼️</button>
            <button type="button" class="tool-btn" title="链接">🔗</button>
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
              <img v-if="comment.avatar" :src="comment.avatar" :alt="comment.nickname" />
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
                  <span class="icon">♡</span>
                </button>
                <button class="action-btn reply-btn" @click="startReply(comment)">
                  <span class="icon">💬</span>
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
                  <img v-if="reply.avatar" :src="reply.avatar" :alt="reply.nickname" />
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
                      <span class="icon">♡</span>
                    </button>
                    <button class="action-btn reply-btn" @click="startReply(comment, reply)">
                      <span class="icon">💬</span>
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
import { ref, computed, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { getCommentsByArticle, createComment } from '../../api/comment'

const route = useRoute()

const comments = ref([])
const loading = ref(false)
const submitting = ref(false)
const replyingTo = ref(null)
const replyContent = ref('')
const sortBy = ref('desc')
const isLoggedIn = ref(false)

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
    const articleId = route.params.id || route.params.slug
    await createComment({
      article_id: parseInt(articleId),
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
    const articleId = route.params.id || route.params.slug
    await createComment({
      article_id: parseInt(articleId),
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

onMounted(fetchComments)
</script>

<style scoped>
.comment-section {
  margin-top: 40px;
  padding-top: 40px;
  border-top: 1px solid var(--border-color);
}

/* 评论表单 */
.comment-form-wrapper {
  background: var(--card-bg);
  border-radius: 12px;
  margin-bottom: 32px;
  overflow: hidden;
}

.comment-form {
  display: flex;
  flex-direction: column;
}

.form-header {
  display: flex;
  gap: 16px;
  padding: 16px 20px;
  border-bottom: 1px dashed var(--border-color);
}

.form-field {
  flex: 1;
  display: flex;
  align-items: center;
  gap: 8px;
}

.form-field label {
  font-size: 0.9rem;
  color: var(--card-text-color-secondary);
  white-space: nowrap;
}

.form-field input {
  flex: 1;
  padding: 8px 12px;
  border: none;
  background: transparent;
  color: var(--card-text-color-main);
  font-size: 0.95rem;
}

.form-field input:focus {
  outline: none;
}

.form-field input::placeholder {
  color: var(--card-text-color-tertiary);
}

.form-body {
  padding: 16px 20px;
}

.form-body textarea {
  width: 100%;
  border: none;
  background: transparent;
  color: var(--card-text-color-main);
  font-size: 0.95rem;
  resize: none;
  min-height: 100px;
  font-family: inherit;
}

.form-body textarea:focus {
  outline: none;
}

.form-body textarea::placeholder {
  color: var(--card-text-color-tertiary);
}

.form-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px 20px;
  border-top: 1px solid var(--border-color);
}

.form-tools {
  display: flex;
  gap: 8px;
}

.tool-btn {
  width: 32px;
  height: 32px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: none;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  font-size: 1.1rem;
  transition: background 0.2s;
}

.tool-btn:hover {
  background: var(--hover-bg);
}

.form-actions {
  display: flex;
  align-items: center;
  gap: 12px;
}

.char-count {
  font-size: 0.85rem;
  color: var(--card-text-color-tertiary);
}

.login-btn {
  padding: 8px 16px;
  background: transparent;
  color: var(--card-text-color-secondary);
  border: 1px solid var(--border-color);
  border-radius: 6px;
  font-size: 0.9rem;
  cursor: pointer;
  transition: all 0.2s;
}

.login-btn:hover {
  background: var(--hover-bg);
}

.submit-btn {
  padding: 8px 20px;
  background: var(--accent-color);
  color: white;
  border: none;
  border-radius: 6px;
  font-size: 0.9rem;
  cursor: pointer;
  transition: opacity 0.2s;
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
  margin-top: 24px;
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
  color: var(--card-text-color-main);
  margin: 0;
}

.sort-options {
  display: flex;
  gap: 16px;
}

.sort-btn {
  background: none;
  border: none;
  color: var(--card-text-color-secondary);
  font-size: 0.9rem;
  cursor: pointer;
  padding: 4px 8px;
  border-radius: 4px;
  transition: all 0.2s;
}

.sort-btn:hover {
  color: var(--accent-color);
}

.sort-btn.active {
  color: var(--accent-color);
  font-weight: 500;
}

.loading,
.empty {
  text-align: center;
  padding: 40px;
  color: var(--card-text-color-secondary);
}

.comment-list {
  display: flex;
  flex-direction: column;
  gap: 24px;
}

.comment-item {
  background: var(--card-bg);
  padding: 20px;
  border-radius: 12px;
}

.comment-main {
  display: flex;
  gap: 16px;
}

.comment-main.small {
  gap: 12px;
}

.avatar {
  width: 48px;
  height: 48px;
  border-radius: 50%;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: 600;
  font-size: 1.1rem;
  flex-shrink: 0;
  overflow: hidden;
}

.avatar.small {
  width: 36px;
  height: 36px;
  font-size: 0.9rem;
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
  margin-bottom: 8px;
}

.nickname {
  font-weight: 600;
  color: var(--card-text-color-main);
  font-size: 0.95rem;
}

.admin-badge {
  display: inline-block;
  padding: 2px 8px;
  background: var(--accent-color);
  color: white;
  font-size: 0.75rem;
  border-radius: 4px;
  font-weight: 500;
}

.time {
  font-size: 0.8rem;
  color: var(--card-text-color-tertiary);
}

.comment-content {
  color: var(--card-text-color-main);
  line-height: 1.6;
  margin-bottom: 12px;
  word-wrap: break-word;
}

.reply-to {
  color: var(--accent-color);
  font-weight: 500;
}

.comment-actions {
  display: flex;
  gap: 16px;
}

.action-btn {
  display: flex;
  align-items: center;
  gap: 4px;
  background: none;
  border: none;
  color: var(--card-text-color-secondary);
  font-size: 0.85rem;
  cursor: pointer;
  padding: 4px 8px;
  border-radius: 4px;
  transition: all 0.2s;
}

.action-btn:hover {
  color: var(--accent-color);
  background: rgba(var(--accent-color-rgb), 0.1);
}

.action-btn .icon {
  font-size: 1rem;
}

.like-btn:hover {
  color: #e74c3c;
}

/* 回复表单 */
.reply-form {
  margin-top: 16px;
  margin-left: 64px;
  padding: 16px;
  background: var(--bg-color);
  border-radius: 8px;
}

.reply-form textarea {
  width: 100%;
  padding: 12px;
  border: 1px solid var(--border-color);
  border-radius: 8px;
  background: var(--card-bg);
  color: var(--card-text-color-main);
  font-size: 0.9rem;
  resize: vertical;
  min-height: 80px;
  box-sizing: border-box;
  font-family: inherit;
}

.reply-form textarea:focus {
  outline: none;
  border-color: var(--accent-color);
}

.reply-form textarea::placeholder {
  color: var(--card-text-color-tertiary);
}

.reply-actions {
  display: flex;
  justify-content: flex-end;
  gap: 12px;
  margin-top: 12px;
}

.cancel-btn {
  padding: 8px 16px;
  background: transparent;
  color: var(--card-text-color-secondary);
  border: 1px solid var(--border-color);
  border-radius: 6px;
  font-size: 0.9rem;
  cursor: pointer;
  transition: all 0.2s;
}

.cancel-btn:hover {
  background: var(--hover-bg);
}

.reply-form .submit-btn {
  padding: 8px 16px;
  font-size: 0.9rem;
}

/* 子评论列表 */
.reply-list {
  margin-top: 16px;
  margin-left: 64px;
}

.reply-item {
  padding: 16px 0;
}

.reply-item:not(:last-child) {
  border-bottom: 1px dashed var(--border-color);
}

@media (max-width: 768px) {
  .form-header {
    flex-direction: column;
    gap: 12px;
  }

  .comment-header-bar {
    flex-direction: column;
    align-items: flex-start;
    gap: 12px;
  }

  .reply-form,
  .reply-list {
    margin-left: 48px;
  }
}
</style>
