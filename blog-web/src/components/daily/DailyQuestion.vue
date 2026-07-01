<template>
  <article class="article-card dq-daily-card">
    <a class="card-link dq-daily-link">
      <div class="dq-daily-header">
        <span class="category-pill category-life">每日一问</span>
        <div class="dq-daily-nav">
          <button class="dq-daily-nav-btn" @click="prevDay" :disabled="!hasPrev">◀ 前一天</button>
          <div class="dq-daily-date-wrapper">
            <button class="dq-daily-date-btn" @click.stop="toggleCalendar">
              <svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><rect width="18" height="18" x="3" y="4" rx="2" ry="2"></rect><line x1="16" x2="16" y1="2" y2="6"></line><line x1="8" x2="8" y1="2" y2="6"></line><line x1="3" x2="21" y1="10" y2="10"></line></svg>
              <span>{{ question?.date || '加载中...' }}</span>
            </button>
            <!-- 日历选择器 -->
            <div v-if="showCalendar" class="dq-calendar-popup" @click.stop>
              <div class="dq-calendar-header">
                <button class="dq-calendar-nav" @click="prevMonth">◀</button>
                <span class="dq-calendar-title">{{ calendarYear }}年{{ calendarMonth + 1 }}月</span>
                <button class="dq-calendar-nav" @click="nextMonth">▶</button>
              </div>
              <div class="dq-calendar-weekdays">
                <span v-for="day in ['日', '一', '二', '三', '四', '五', '六']" :key="day">{{ day }}</span>
              </div>
              <div class="dq-calendar-days">
                <button
                  v-for="(day, index) in calendarDays"
                  :key="index"
                  class="dq-calendar-day"
                  :class="{
                    'other-month': day.otherMonth,
                    'is-today': day.isToday,
                    'is-selected': day.date === question?.date,
                    'has-question': day.hasQuestion
                  }"
                  :disabled="day.otherMonth"
                  @click="selectDate(day.date)"
                >
                  {{ day.day }}
                </button>
              </div>
              <button class="dq-calendar-close" @click="showCalendar = false">关闭</button>
            </div>
          </div>
          <button class="dq-daily-nav-btn" @click="nextDay" :disabled="!hasNext">后一天 ▶</button>
        </div>
      </div>

      <h2 class="dq-daily-title">{{ question?.question || '暂无问题' }}</h2>

      <div class="dq-daily-answer-wrapper" v-if="question">
        <p class="dq-daily-preview">{{ question.answer }}</p>
        <div v-if="!answerVisible" class="dq-daily-mask" @click="answerVisible = true">
          <span class="dq-daily-mask-text">👉 点击查看答案</span>
        </div>
        <div v-if="answerVisible" class="dq-daily-hide-btn" @click="answerVisible = false">
          <span class="dq-daily-hide-text">👈 点击收起答案</span>
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
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { getLatestQuestion, getQuestionByDate, getPreviousQuestion, getNextQuestion, likeQuestion } from '../../api/daily'

const question = ref(null)
const answerVisible = ref(false)
const isLiked = ref(false)
const hasPrev = ref(false)
const hasNext = ref(false)
const showCalendar = ref(false)
const calendarYear = ref(new Date().getFullYear())
const calendarMonth = ref(new Date().getMonth())

// 计算日历天数
const calendarDays = computed(() => {
  const year = calendarYear.value
  const month = calendarMonth.value
  const firstDay = new Date(year, month, 1)
  const lastDay = new Date(year, month + 1, 0)
  const startDayOfWeek = firstDay.getDay()
  const daysInMonth = lastDay.getDate()

  const days = []
  const today = new Date()
  const todayStr = `${today.getFullYear()}-${String(today.getMonth() + 1).padStart(2, '0')}-${String(today.getDate()).padStart(2, '0')}`

  // 填充上个月的日期
  const prevMonthLastDay = new Date(year, month, 0).getDate()
  for (let i = startDayOfWeek - 1; i >= 0; i--) {
    const day = prevMonthLastDay - i
    const m = month === 0 ? 12 : month
    const y = month === 0 ? year - 1 : year
    days.push({
      day,
      date: `${y}-${String(m).padStart(2, '0')}-${String(day).padStart(2, '0')}`,
      otherMonth: true,
      isToday: false,
      hasQuestion: false
    })
  }

  // 填充当月日期
  for (let day = 1; day <= daysInMonth; day++) {
    const dateStr = `${year}-${String(month + 1).padStart(2, '0')}-${String(day).padStart(2, '0')}`
    days.push({
      day,
      date: dateStr,
      otherMonth: false,
      isToday: dateStr === todayStr,
      hasQuestion: false
    })
  }

  // 填充下个月的日期
  const remaining = 42 - days.length
  for (let day = 1; day <= remaining; day++) {
    const m = month + 2 > 12 ? 1 : month + 2
    const y = month + 2 > 12 ? year + 1 : year
    days.push({
      day,
      date: `${y}-${String(m).padStart(2, '0')}-${String(day).padStart(2, '0')}`,
      otherMonth: true,
      isToday: false,
      hasQuestion: false
    })
  }

  return days
})

const loadQuestion = async (date) => {
  try {
    const res = date ? await getQuestionByDate(date).catch(() => getLatestQuestion()) : await getLatestQuestion()
    if (res.data) {
      question.value = res.data
      answerVisible.value = false
      isLiked.value = localStorage.getItem('dq_liked_' + res.data?.id) === 'true'
      // 更新日历月份到问题所在月份
      if (res.data.date) {
        const [y, m] = res.data.date.split('-').map(Number)
        calendarYear.value = y
        calendarMonth.value = m - 1
        await checkNavigation(res.data.date)
      }
    }
  } catch (e) {
    console.error(e)
  }
}

const checkNavigation = async (date) => {
  try {
    await getPreviousQuestion(date)
    hasPrev.value = true
  } catch {
    hasPrev.value = false
  }
  try {
    await getNextQuestion(date)
    hasNext.value = true
  } catch {
    hasNext.value = false
  }
}

const toggleCalendar = () => {
  showCalendar.value = !showCalendar.value
}

const prevMonth = () => {
  if (calendarMonth.value === 0) {
    calendarMonth.value = 11
    calendarYear.value--
  } else {
    calendarMonth.value--
  }
}

const nextMonth = () => {
  if (calendarMonth.value === 11) {
    calendarMonth.value = 0
    calendarYear.value++
  } else {
    calendarMonth.value++
  }
}

const selectDate = async (date) => {
  try {
    const res = await getQuestionByDate(date)
    if (res.data) {
      question.value = res.data
      answerVisible.value = false
      isLiked.value = localStorage.getItem('dq_liked_' + res.data?.id) === 'true'
      await checkNavigation(res.data.date)
    }
    showCalendar.value = false
  } catch {
    alert('该日期没有问题')
  }
}

const prevDay = async () => {
  if (question.value?.date) {
    try {
      const res = await getPreviousQuestion(question.value.date)
      if (res.data) {
        question.value = res.data
        answerVisible.value = false
        isLiked.value = localStorage.getItem('dq_liked_' + res.data?.id) === 'true'
        const [y, m] = res.data.date.split('-').map(Number)
        calendarYear.value = y
        calendarMonth.value = m - 1
        await checkNavigation(res.data.date)
      }
    } catch (e) {
      console.error('没有前一天的问题了')
    }
  }
}

const nextDay = async () => {
  if (question.value?.date) {
    try {
      const res = await getNextQuestion(question.value.date)
      if (res.data) {
        question.value = res.data
        answerVisible.value = false
        isLiked.value = localStorage.getItem('dq_liked_' + res.data?.id) === 'true'
        const [y, m] = res.data.date.split('-').map(Number)
        calendarYear.value = y
        calendarMonth.value = m - 1
        await checkNavigation(res.data.date)
      }
    } catch (e) {
      console.error('没有后一天的问题了')
    }
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

// 点击外部关闭日历
const handleClickOutside = (e) => {
  if (showCalendar.value && !e.target.closest('.dq-calendar-popup') && !e.target.closest('.dq-daily-date-btn')) {
    showCalendar.value = false
  }
}

onMounted(() => {
  loadQuestion()
  document.addEventListener('click', handleClickOutside)
})

onUnmounted(() => {
  document.removeEventListener('click', handleClickOutside)
})
</script>
