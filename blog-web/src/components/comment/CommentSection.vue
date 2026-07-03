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
          <div
            ref="editorRef"
            class="editor-input"
            contenteditable="true"
            data-placeholder="欢迎评论，填写邮箱可收到回复提醒~"
            @input="onEditorInput"
            @keydown="onEditorKeydown"
          ></div>
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
              <div class="comment-content" v-html="renderCommentContent(comment.content)"></div>
              <div class="comment-actions">
                <button class="action-btn like-btn" :class="{ 'liked': likedComments.has(comment.id) }" @click="handleLikeComment(comment)">
                  <span class="icon">
                    <svg v-if="likedComments.has(comment.id)" xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="#e74c3c" stroke="#e74c3c" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M20.84 4.61a5.5 5.5 0 0 0-7.78 0L12 5.67l-1.06-1.06a5.5 5.5 0 0 0-7.78 7.78l1.06 1.06L12 21.23l7.78-7.78 1.06-1.06a5.5 5.5 0 0 0 0-7.78z"></path></svg>
                    <svg v-else xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M20.84 4.61a5.5 5.5 0 0 0-7.78 0L12 5.67l-1.06-1.06a5.5 5.5 0 0 0-7.78 7.78l1.06 1.06L12 21.23l7.78-7.78 1.06-1.06a5.5 5.5 0 0 0 0-7.78z"></path></svg>
                  </span>
                  <span v-if="comment.like_count > 0" class="like-count">{{ comment.like_count }}</span>
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
                    <span v-html="renderCommentContent(reply.content)"></span>
                  </div>
                  <div class="comment-actions">
                    <button class="action-btn like-btn" :class="{ 'liked': likedComments.has(reply.id) }" @click="handleLikeComment(reply)">
                      <span class="icon">
                        <svg v-if="likedComments.has(reply.id)" xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="#e74c3c" stroke="#e74c3c" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M20.84 4.61a5.5 5.5 0 0 0-7.78 0L12 5.67l-1.06-1.06a5.5 5.5 0 0 0-7.78 7.78l1.06 1.06L12 21.23l7.78-7.78 1.06-1.06a5.5 5.5 0 0 0 0-7.78z"></path></svg>
                        <svg v-else xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M20.84 4.61a5.5 5.5 0 0 0-7.78 0L12 5.67l-1.06-1.06a5.5 5.5 0 0 0-7.78 7.78l1.06 1.06L12 21.23l7.78-7.78 1.06-1.06a5.5 5.5 0 0 0 0-7.78z"></path></svg>
                      </span>
                      <span v-if="reply.like_count > 0" class="like-count">{{ reply.like_count }}</span>
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
import { getCommentsByArticle, createComment, likeComment } from '../../api/comment'
import EmojiPicker from 'vue3-emoji-picker'
import 'vue3-emoji-picker/css'
import { getAvatarUrl } from '../../utils/avatar'
import { marked } from 'marked'

const route = useRoute()
const editorRef = ref(null)

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

// 表情图片 CDN 地址
const EMOJI_CDN = 'https://cdn.jsdelivr.net/npm/emoji-datasource-apple@6.0.1/img/apple/64'

// unicode 字符转 emoji 图片文件名（如 😊 → 1f600）
const unicodeToEmojiCode = (str) => {
  const codes = []
  for (const char of str) {
    const code = char.codePointAt(0)
    if (code > 0x1f000) {
      codes.push(code.toString(16))
    }
  }
  return codes.join('-')
}

// 排序后的评论
const sortedComments = computed(() => {
  const list = [...comments.value]
  if (sortBy.value === 'asc') {
    return list.sort((a, b) => new Date(a.created_at) - new Date(b.created_at))
  } else if (sortBy.value === 'desc') {
    return list.sort((a, b) => new Date(b.created_at) - new Date(a.created_at))
  } else {
    return list.sort((a, b) => (b.like_count || 0) - (a.like_count || 0))
  }
})

// 切换表情面板
const toggleEmojiPanel = () => {
  showEmojiPanel.value = !showEmojiPanel.value
}

// 编辑器输入同步（用 innerHTML 保留 emoji 图片标签）
const onEditorInput = () => {
  if (editorRef.value) {
    form.value.content = editorRef.value.innerHTML || ''
  }
}

// 编辑器键盘事件
const onEditorKeydown = (e) => {
  // Enter 提交（Shift+Enter 换行）
  if (e.key === 'Enter' && !e.shiftKey) {
    e.preventDefault()
    if (form.value.content.trim()) {
      submitComment()
    }
  }
}

// 插入 emoji 图片到 contenteditable
const insertEmoji = (emojiStr) => {
  const editor = editorRef.value
  if (!editor) return

  const emojiCode = unicodeToEmojiCode(emojiStr)
  if (!emojiCode) return

  const img = document.createElement('img')
  img.src = `${EMOJI_CDN}/${emojiCode}.png`
  img.alt = emojiStr
  img.className = 'inline-emoji'
  img.style.cssText = 'width:1.2em;height:1.2em;vertical-align:middle;margin:0 2px;'

  // 插入到光标位置
  const selection = window.getSelection()
  if (selection.rangeCount > 0) {
    const range = selection.getRangeAt(0)
    // 确保光标在编辑器内
    if (editor.contains(range.commonAncestorContainer)) {
      range.deleteContents()
      range.insertNode(img)
      range.collapse(false)
      selection.removeAllRanges()
      selection.addRange(range)
    } else {
      editor.appendChild(img)
    }
  } else {
    editor.appendChild(img)
  }

  // 在图片后插入空格，方便继续输入
  const space = document.createTextNode(' ')
  img.parentNode.insertBefore(space, img.nextSibling)

  // 同步内容
  form.value.content = editor.innerHTML || ''
  showEmojiPanel.value = false
}

// EmojiPicker 选择回调
const onEmojiSelect = (emoji) => {
  const emojiStr = typeof emoji === 'string' ? emoji : (emoji.i || '')
  if (emojiStr) {
    insertEmoji(emojiStr)
  }
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

// 净化 HTML：移除危险属性和标签
const sanitizeCommentHtml = (html) => {
  let clean = html.replace(/\s+on\w+\s*=\s*(?:"[^"]*"|'[^']*'|[^\s>]+)/gi, '')
  clean = clean.replace(/href\s*=\s*(?:"javascript:[^"]*"|'javascript:[^']*')/gi, '')
  clean = clean.replace(/src\s*=\s*(?:"javascript:[^"]*"|'javascript:[^']*')/gi, '')
  clean = clean.replace(/<script[\s\S]*?<\/script>/gi, '')
  clean = clean.replace(/<iframe[\s\S]*?<\/iframe>/gi, '')
  return clean
}

// 渲染评论内容（将 emoji unicode 转为图片）
const renderCommentContent = (text) => {
  if (!text) return ''
  // 如果内容已包含 <img> 标签，净化后返回
  if (text.includes('<img')) {
    return sanitizeCommentHtml(text)
  }
  // 纯文本内容：转义 HTML，然后将 emoji unicode 转为 <img> 标签
  const escaped = text.replace(/&/g, '&amp;').replace(/</g, '&lt;').replace(/>/g, '&gt;')
  return escaped.replace(/[\u{1F600}-\u{1F64F}\u{1F300}-\u{1F5FF}\u{1F680}-\u{1F6FF}\u{1F1E0}-\u{1F1FF}\u{2600}-\u{26FF}\u{2700}-\u{27BF}\u{FE00}-\u{FE0F}\u{1F900}-\u{1F9FF}\u{1FA00}-\u{1FA6F}\u{1FA70}-\u{1FAFF}]/gu, (char) => {
    const code = unicodeToEmojiCode(char)
    if (code) {
      return `<img src="${EMOJI_CDN}/${code}.png" alt="${char}" class="inline-emoji" style="width:1.2em;height:1.2em;vertical-align:middle;margin:0 2px;">`
    }
    return char
  })
}

// 插入 Markdown 格式
const insertMarkdown = () => {
  const editor = editorRef.value
  if (!editor) return

  const selection = window.getSelection()
  const selectedText = selection.toString() || '粗体文本'
  const markdownText = `**${selectedText}**`

  document.execCommand('insertText', false, markdownText)
  form.value.content = editor.innerHTML || ''
}

// 插入图片
const insertImage = () => {
  const url = prompt('请输入图片URL:')
  if (url) {
    const editor = editorRef.value
    if (!editor) return
    const imageText = `![图片](${url})`
    document.execCommand('insertText', false, imageText)
    form.value.content = editor.innerHTML || ''
  }
}

// 插入链接
const insertLink = () => {
  const url = prompt('请输入链接URL:')
  if (url) {
    const editor = editorRef.value
    if (!editor) return
    const selection = window.getSelection()
    const selectedText = selection.toString() || '链接文本'
    const linkText = `[${selectedText}](${url})`
    document.execCommand('insertText', false, linkText)
    form.value.content = editor.innerHTML || ''
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
    // 从 localStorage 恢复已点赞状态
    comments.value.forEach(comment => {
      if (localStorage.getItem(`liked_comment_${comment.id}`) === 'true') {
        likedComments.value.add(comment.id)
      }
      if (comment.replies) {
        comment.replies.forEach(reply => {
          if (localStorage.getItem(`liked_comment_${reply.id}`) === 'true') {
            likedComments.value.add(reply.id)
          }
        })
      }
    })
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
    if (editorRef.value) editorRef.value.innerHTML = ''
    // 重新获取评论
    await fetchComments()
  } catch (error) {
    console.error('提交评论失败:', error)
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
  } finally {
    submitting.value = false
  }
}

// 检查是否已点赞
const likedComments = ref(new Set())

// 点赞评论
const handleLikeComment = async (comment) => {
  // 检查本地是否已点赞
  if (likedComments.value.has(comment.id)) {
    return
  }

  try {
    await likeComment(comment.id)
    comment.like_count = (comment.like_count || 0) + 1
    likedComments.value.add(comment.id)
    // 持久化到 localStorage
    localStorage.setItem(`liked_comment_${comment.id}`, 'true')
  } catch (error) {
    console.error('点赞失败:', error)
    // 如果是已点赞的业务错误，也标记为已点赞
    if (error.response?.data?.code === 4006) {
      likedComments.value.add(comment.id)
      localStorage.setItem(`liked_comment_${comment.id}`, 'true')
    }
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
  margin-bottom: 24px;
  position: relative;
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

.editor-input {
  width: 100%;
  min-height: 100px;
  max-height: 300px;
  overflow-y: auto;
  border: none;
  background: transparent;
  color: #333;
  font-size: 0.95rem;
  font-family: inherit;
  line-height: 1.6;
  word-wrap: break-word;
  outline: none;
}

.editor-input:empty::before {
  content: attr(data-placeholder);
  color: #999;
  pointer-events: none;
}

.editor-input .inline-emoji {
  width: 1.2em;
  height: 1.2em;
  vertical-align: middle;
  margin: 0 2px;
}

/* 评论内容中的 emoji 图片 */
.comment-content .inline-emoji {
  width: 1.2em;
  height: 1.2em;
  vertical-align: middle;
  margin: 0 2px;
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
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.12);
  z-index: 100;
  animation: emoji-fade-in 0.2s ease;
}

@keyframes emoji-fade-in {
  from {
    opacity: 0;
    transform: translateY(8px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
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

.like-btn.liked {
  color: #e74c3c;
}

.like-count {
  font-size: 0.8rem;
  margin-left: 2px;
  min-width: 12px;
  display: inline-block;
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
