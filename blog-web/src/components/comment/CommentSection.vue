<template>
  <div class="comment-section">
    <!-- 评论表单 -->
    <div class="comment-form-wrapper">
      <h3 class="section-title">发表评论</h3>
      <form class="comment-form" @submit.prevent="submitComment">
        <div class="form-row">
          <div class="form-group">
            <label for="nickname">昵称 *</label>
            <input
              id="nickname"
              v-model="form.nickname"
              type="text"
              placeholder="请输入昵称"
              required
            />
          </div>
          <div class="form-group">
            <label for="email">邮箱（选填，用于接收回复通知）</label>
            <input
              id="email"
              v-model="form.email"
              type="email"
              placeholder="请输入邮箱"
            />
          </div>
        </div>
        <div class="form-group">
          <label for="content">评论内容 *</label>
          <textarea
            id="content"
            v-model="form.content"
            placeholder="写下你的评论..."
            rows="4"
            required
          ></textarea>
        </div>
        <div class="form-actions">
          <span class="char-count">{{ form.content.length }} 字</span>
          <button type="submit" class="submit-btn" :disabled="submitting">
            {{ submitting ? '提交中...' : '提交评论' }}
          </button>
        </div>
      </form>
    </div>

    <!-- 评论列表 -->
    <div class="comment-list-wrapper">
      <h3 class="section-title">{{ comments.length }} 条评论</h3>

      <div v-if="loading" class="loading">加载中...</div>

      <div v-else-if="comments.length === 0" class="empty">暂无评论，快来抢沙发吧！</div>

      <div v-else class="comment-list">
        <div v-for="comment in comments" :key="comment.id" class="comment-item">
          <div class="comment-header">
            <div class="avatar">{{ comment.nickname.charAt(0).toUpperCase() }}</div>
            <div class="comment-info">
              <span class="nickname">{{ comment.nickname }}</span>
              <span class="time">{{ formatTime(comment.created_at) }}</span>
            </div>
          </div>
          <div class="comment-content">{{ comment.content }}</div>
          <div class="comment-actions">
            <button class="reply-btn" @click="startReply(comment)">回复</button>
          </div>

          <!-- 回复表单 -->
          <div v-if="replyingTo?.id === comment.id" class="reply-form">
            <textarea
              v-model="replyContent"
              placeholder="写下你的回复..."
              rows="3"
            ></textarea>
            <div class="reply-actions">
              <button class="cancel-btn" @click="cancelReply">取消</button>
              <button class="submit-btn" @click="submitReply" :disabled="submitting">
                {{ submitting ? '提交中...' : '提交回复' }}
              </button>
            </div>
          </div>

          <!-- 子评论（回复） -->
          <div v-if="comment.replies && comment.replies.length > 0" class="reply-list">
            <div v-for="reply in comment.replies" :key="reply.id" class="reply-item">
              <div class="comment-header">
                <div class="avatar small">{{ reply.nickname.charAt(0).toUpperCase() }}</div>
                <div class="comment-info">
                  <span class="nickname">{{ reply.nickname }}</span>
                  <span class="time">{{ formatTime(reply.created_at) }}</span>
                </div>
              </div>
              <div class="comment-content">{{ reply.content }}</div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { getCommentsByArticle, createComment } from '../../api/comment'

const route = useRoute()

const comments = ref([])
const loading = ref(false)
const submitting = ref(false)
const replyingTo = ref(null)
const replyContent = ref('')

const form = ref({
  nickname: '',
  email: '',
  content: ''
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
  if (!form.value.nickname || !form.value.content) return

  submitting.value = true
  try {
    const articleId = route.params.id || route.params.slug
    await createComment({
      article_id: parseInt(articleId),
      nickname: form.value.nickname,
      email: form.value.email,
      content: form.value.content
    })

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
const startReply = (comment) => {
  replyingTo.value = comment
  replyContent.value = ''
}

// 取消回复
const cancelReply = () => {
  replyingTo.value = null
  replyContent.value = ''
}

// 提交回复
const submitReply = async () => {
  if (!replyContent.value || !replyingTo.value) return

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

.section-title {
  font-size: 1.2rem;
  font-weight: 600;
  color: var(--card-text-color-main);
  margin-bottom: 20px;
}

/* 评论表单 */
.comment-form-wrapper {
  background: var(--card-bg);
  padding: 24px;
  border-radius: 12px;
  margin-bottom: 32px;
}

.comment-form .form-row {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 16px;
  margin-bottom: 16px;
}

@media (max-width: 768px) {
  .comment-form .form-row {
    grid-template-columns: 1fr;
  }
}

.form-group {
  margin-bottom: 16px;
}

.form-group label {
  display: block;
  font-size: 0.9rem;
  color: var(--card-text-color-secondary);
  margin-bottom: 6px;
}

.form-group input,
.form-group textarea {
  width: 100%;
  padding: 12px;
  border: 1px solid var(--border-color);
  border-radius: 8px;
  background: var(--bg-color);
  color: var(--card-text-color-main);
  font-size: 0.95rem;
  transition: border-color 0.2s;
  box-sizing: border-box;
}

.form-group input:focus,
.form-group textarea:focus {
  outline: none;
  border-color: var(--accent-color);
}

.form-group textarea {
  resize: vertical;
  min-height: 100px;
}

.form-actions {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.char-count {
  font-size: 0.85rem;
  color: var(--card-text-color-tertiary);
}

.submit-btn {
  padding: 10px 24px;
  background: var(--accent-color);
  color: white;
  border: none;
  border-radius: 8px;
  font-size: 0.95rem;
  cursor: pointer;
  transition: opacity 0.2s;
}

.submit-btn:hover:not(:disabled) {
  opacity: 0.9;
}

.submit-btn:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

/* 评论列表 */
.comment-list-wrapper {
  margin-top: 24px;
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

.comment-header {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-bottom: 12px;
}

.avatar {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  background: var(--accent-color);
  color: white;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: 600;
  font-size: 1rem;
}

.avatar.small {
  width: 32px;
  height: 32px;
  font-size: 0.85rem;
}

.comment-info {
  display: flex;
  flex-direction: column;
}

.nickname {
  font-weight: 600;
  color: var(--card-text-color-main);
  font-size: 0.95rem;
}

.time {
  font-size: 0.8rem;
  color: var(--card-text-color-tertiary);
}

.comment-content {
  color: var(--card-text-color-main);
  line-height: 1.6;
  margin-bottom: 12px;
}

.comment-actions {
  display: flex;
  gap: 12px;
}

.reply-btn {
  background: none;
  border: none;
  color: var(--card-text-color-secondary);
  font-size: 0.85rem;
  cursor: pointer;
  padding: 4px 8px;
  border-radius: 4px;
  transition: all 0.2s;
}

.reply-btn:hover {
  color: var(--accent-color);
  background: rgba(var(--accent-color-rgb), 0.1);
}

/* 回复表单 */
.reply-form {
  margin-top: 16px;
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
}

.reply-form textarea:focus {
  outline: none;
  border-color: var(--accent-color);
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
  background: var(--card-bg);
}

.reply-form .submit-btn {
  padding: 8px 16px;
  font-size: 0.9rem;
}

/* 子评论列表 */
.reply-list {
  margin-top: 16px;
  padding-left: 20px;
  border-left: 2px solid var(--border-color);
}

.reply-item {
  padding: 16px 0;
}

.reply-item:not(:last-child) {
  border-bottom: 1px dashed var(--border-color);
}
</style>
