<template>
  <article class="article-card dq-daily-card">
    <a class="card-link dq-daily-link">
      <div class="dq-daily-header">
        <span class="category-pill category-life">每日一问</span>
        <div class="dq-daily-nav">
          <button class="dq-daily-nav-btn" @click="prevDay" :disabled="!prevDate">◀ 前一天</button>
          <button class="dq-daily-date-btn" @click="showCalendar = !showCalendar">
            <svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><rect width="18" height="18" x="3" y="4" rx="2" ry="2"></rect><line x1="16" x2="16" y1="2" y2="6"></line><line x1="8" x2="8" y1="2" y2="6"></line><line x1="3" x2="21" y1="10" y2="10"></line></svg>
            <span>{{ question?.date || '加载中...' }}</span>
          </button>
          <button class="dq-daily-nav-btn" @click="nextDay" :disabled="!nextDate">后一天 ▶</button>
        </div>
      </div>

      <h2 class="dq-daily-title">{{ question?.question || '暂无问题' }}</h2>

      <div class="dq-daily-answer-wrapper" v-if="question" :class="{ expanded: answerVisible }">
        <p class="dq-daily-preview">{{ question.answer }}</p>
        <div v-if="!answerVisible" class="dq-daily-mask" @click="answerVisible = true">
          <span class="dq-daily-mask-text">点击查看答案</span>
        </div>
        <div v-if="answerVisible" class="dq-daily-hide-btn" @click="answerVisible = false">
          <span class="dq-daily-hide-text">点击收起答案</span>
        </div>
      </div>

      <div class="dq-daily-footer">
        <button class="dq-daily-like-btn" :class="{ liked: isLiked }" @click="handleLike">
          <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M20.84 4.61a5.5 5.5 0 0 0-7.78 0L12 5.67l-1.06-1.06a5.5 5.5 0 0 0-7.78 7.78l1.06 1.06L12 21.23l7.78-7.78 1.06-1.06a5.5 5.5 0 0 0 0-7.78z"></path></svg>
          <span>{{ question?.like_count || 0 }}</span>
        </button>
        <button class="dq-daily-comment-btn">
          <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z"></path></svg>
          <span>{{ question?.comment_count || 0 }}</span>
        </button>
      </div>
    </a>
  </article>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { getLatestQuestion, getQuestionByDate, likeQuestion } from '../../api/daily'

const question = ref(null)
const answerVisible = ref(false)
const isLiked = ref(false)
const prevDate = ref('')
const nextDate = ref('')
const showCalendar = ref(false)

const loadQuestion = async (date) => {
  try {
    const res = date ? await getQuestionByDate(date) : await getLatestQuestion()
    question.value = res.data
    answerVisible.value = false
    isLiked.value = localStorage.getItem('dq_liked_' + res.data?.id) === 'true'
    if (res.data?.date) {
      prevDate.value = res.data.date
      nextDate.value = res.data.date
    }
  } catch (e) {
    console.error(e)
  }
}

const prevDay = async () => {
  if (prevDate.value) {
    const d = new Date(prevDate.value)
    d.setDate(d.getDate() - 1)
    await loadQuestion(d.toISOString().split('T')[0])
  }
}

const nextDay = async () => {
  if (nextDate.value) {
    const d = new Date(nextDate.value)
    d.setDate(d.getDate() + 1)
    await loadQuestion(d.toISOString().split('T')[0])
  }
}

const handleLike = async () => {
  if (isLiked.value || !question.value) return
  try {
    const res = await likeQuestion(question.value.id)
    question.value.like_count = res.data?.like_count || question.value.like_count + 1
    localStorage.setItem('dq_liked_' + question.value.id, 'true')
    isLiked.value = true
  } catch (e) { console.error(e) }
}

onMounted(() => loadQuestion())
</script>

<style scoped>
.dq-daily-card {
  border-left: 4px solid var(--accent-color, #667eea);
  display: flex;
  flex-direction: column;
}

.dq-daily-link {
  display: flex;
  flex-direction: column;
  height: 100%;
  padding: 24px 24px 18px;
  text-decoration: none;
  color: inherit;
}

.dq-daily-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 8px;
  flex-wrap: wrap;
  gap: 8px;
}

.dq-daily-nav {
  display: flex;
  align-items: center;
  gap: 8px;
}

.dq-daily-nav-btn,
.dq-daily-date-btn {
  padding: 6px 12px;
  background: var(--card-background, #fff);
  border: 1px solid var(--card-separator-color, #e5e7eb);
  border-radius: 6px;
  font-size: 12px;
  cursor: pointer;
  transition: all 0.2s;
  color: var(--card-text-color-secondary, #6b7280);
  display: flex;
  align-items: center;
  gap: 6px;
}

.dq-daily-nav-btn:hover:not(:disabled),
.dq-daily-date-btn:hover {
  background: var(--accent-color, #667eea);
  color: #fff;
  border-color: var(--accent-color, #667eea);
}

.dq-daily-nav-btn:disabled {
  opacity: 0.4;
  cursor: not-allowed;
}

.dq-daily-title {
  font-size: 1.35rem;
  font-weight: 700;
  margin-bottom: 8px;
  line-height: 1.4;
  color: var(--card-text-color-main, #111827);
}

.dq-daily-answer-wrapper {
  position: relative;
  margin-bottom: 12px;
  max-height: 140px;
  overflow: hidden;
  transition: max-height 0.35s ease;
}

.dq-daily-answer-wrapper.expanded {
  max-height: 2000px;
  overflow: visible;
}

.dq-daily-preview {
  color: var(--card-text-color-secondary, #6b7280);
  line-height: 1.6;
  margin-bottom: 0;
}

.dq-daily-mask {
  position: absolute;
  bottom: 0;
  left: 0;
  right: 0;
  height: 100%;
  background: linear-gradient(
    to bottom,
    transparent 0%,
    transparent 15%,
    rgba(255, 255, 255, 0.6) 35%,
    var(--card-background, #fff) 60%
  );
  display: flex;
  align-items: flex-end;
  justify-content: center;
  padding-bottom: 12px;
  cursor: pointer;
}

.dq-daily-mask-text {
  font-size: 13px;
  color: var(--accent-color, #667eea);
  font-weight: 600;
  background: var(--card-background, #fff);
  padding: 6px 18px;
  border-radius: 20px;
  border: 1px solid var(--accent-color, #667eea);
  transition: all 0.2s;
}

.dq-daily-mask-text:hover {
  background: var(--accent-color, #667eea);
  color: #fff;
}

.dq-daily-hide-btn {
  text-align: center;
  margin-top: 8px;
}

.dq-daily-hide-text {
  font-size: 13px;
  color: var(--accent-color, #667eea);
  font-weight: 500;
  cursor: pointer;
}

.dq-daily-footer {
  display: flex;
  align-items: center;
  gap: 16px;
  margin-top: auto;
  padding-top: 8px;
}

.dq-daily-like-btn,
.dq-daily-comment-btn {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  padding: 6px 12px;
  background: transparent;
  border: 1px solid var(--card-separator-color, #e5e7eb);
  border-radius: 6px;
  font-size: 13px;
  cursor: pointer;
  color: var(--card-text-color-secondary, #6b7280);
  transition: all 0.2s;
}

.dq-daily-like-btn:hover:not(.liked),
.dq-daily-comment-btn:hover {
  border-color: var(--accent-color, #667eea);
  color: var(--accent-color, #667eea);
}

.dq-daily-like-btn.liked {
  background: rgba(239, 68, 68, 0.1);
  border-color: #ef4444;
  color: #ef4444;
}

.category-life {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
}
</style>
