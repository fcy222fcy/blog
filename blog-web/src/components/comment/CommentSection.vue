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
              :disabled="isLoggedIn"
              required
            />
          </div>
          <div class="form-field">
            <label>邮箱(可选)</label>
            <input
              v-model="form.email"
              type="email"
              placeholder="邮箱"
              :disabled="isLoggedIn"
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
                  <path d="M563.2 463.3 677 540c1.7 1.2 3.7 1.8 5.8 1.8.7 0 1.4-.1 2-.2 2.7-.5 5.1-2.1 6.6-4.4l25.3-37.8c1.5-2.3 2.1-5.1 1.6-7.8s-2.1-5.1-4.4-6.6l-73.6-49.1 73.6-49.1c2.3-1.5 3.9-3.9 4.4-6.6.5-2.7 0-5.5-1.6-7.8l-25.3-37.8a10.1 10.1 0 0 0-6.6-4.4c-.7-.1-1.3-.2-2-.2-2.1 0-4.1.6-5.8 1.8l-113.8 76.6c-9.2 6.2-14.7 16.4-14.7  27.5.1 11  5.5  21.3  14.7  27.4zM387 348.8h-45.5c-5.7  0-10.4  4.7-10.4  10.4v153.3c0  5.7  4.7  10.4  10.4  10.4H387c5.7  0  10.4-4.7  10.4-10.4V359.2c0-5.7-4.7-10.4-10.4-10.4zm333.8 241.3-41-20a10.3  10.3  0  0  0-8.1-.5c-2.6.9-4.8  2.9-5.9  5.4-30.1  64.9-93.1  109.1-164.4  115.2-5.7.5-9.9  5.5-9.5  11.2l3.9  45.5c.5  5.3  5  9.5  10.3  9.5h.9c94.8-8  178.5-66.5  218.6-152.7  2.4-5  .3-11.2-4.8-13.6zm186-186.1c-11.9-42-30.5-81.4-55.2-117.1-24.1-34.9-53.5-65.6-87.5-91.2-33.9-25.6-71.5-45.5-111.6-59.2-41.2-14-84.1-21.1-127.8-21.1h-1.2c-75.4  0-148.8  21.4-212.5  61.7-63.7  40.3-114.3  97.6-146.5  165.8-32.2  68.1-44.3  143.6-35.1  218.4  9.3  74.8  39.4  145  87.3  203.3.1.2.3.3.4.5l36.2  38.4c1.1  1.2  2.5  2.1  3.9  2.6  73.3  66.7  168.2  103.5  267.5  103.5  73.3  0  145.2-20.3  207.7-58.7  37.3-22.9  70.3-51.5  98.1-85  27.1-32.7  48.7-69.5  64.2-109.1  15.5-39.7  24.4-81.3  26.6-123.8  2.4-43.6-2.5-87-14.5-129zm-60.5  181.1c-8.3  37-22.8  72-43  104-19.7  31.1-44.3  58.6-73.1  81.7-28.8  23.1-61  41-95.7  53.4-35.6  12.7-72.9  19.1-110.9  19.1-82.6  0-161.7-30.6-222.8-86.2l-34.1-35.8c-23.9-29.3-42.4-62.2-55.1-97.7-12.4-34.7-18.8-71-19.2-107.9-.4-36.9  5.4-73.3  17.1-108.2  12-35.8  30-69.2  53.4-99.1  31.7-40.4  71.1-72  117.2-94.1  44.5-21.3  94-32.6  143.4-32.6  49.3  0  97  10.8  141.8  32  34.3  16.3  65.3  38.1  92  64.8  26.1  26  47.5  56  63.6  89.2  16.2  33.2  26.6  68.5  31  105.1  4.6  37.5  2.7  75.3-5.6  112.3z" fill="currentColor"/>
                </svg>
              </button>
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
            <template v-if="isLoggedIn">
              <button type="button" class="profile-btn" @click="openProfileModal" title="个人设置">
                <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"></path><circle cx="12" cy="7" r="4"></circle></svg>
              </button>
              <button type="button" class="logout-btn" @click="handleLogout">
                <span class="user-info" v-if="currentUser">
                  <span class="user-avatar" v-if="currentUser.avatar">
                    <img :src="currentUser.avatar" />
                  </span>
                  <span class="user-name">{{ currentUser.nickname || currentUser.username }}</span>
                </span>
                退出
              </button>
            </template>
            <button v-else type="button" class="login-btn" @click="authTab='login'; showLoginModal = true">登录 / 注册</button>
            <button type="submit" class="submit-btn" :disabled="submitting || !form.content.trim()">
              提交
            </button>
          </div>
        </div>
      </form>
    </div>

    <!-- 登录/注册弹窗 -->
    <div v-if="showLoginModal" class="modal-overlay" @click.self="showLoginModal = false">
      <div class="login-modal auth-modal">
        <div class="modal-header">
          <div class="auth-tabs">
            <button :class="['auth-tab', { active: authTab === 'login' }]" @click="authTab='login'; loginError=''; registerError=''; registerSuccess=''">登录</button>
            <button :class="['auth-tab', { active: authTab === 'register' }]" @click="authTab='register'; loginError=''; registerError=''; registerSuccess=''">注册</button>
          </div>
          <button type="button" class="close-btn" @click="showLoginModal = false">×</button>
        </div>
        <div class="modal-body">
          <!-- 登录 Tab -->
          <div v-show="authTab === 'login'">
            <p class="login-tip">普通用户也可登录发表评论；博主登录会显示「博主」标识，防止他人冒充。</p>
            <div class="login-field">
              <label>邮箱</label>
              <input v-model="loginForm.email" type="email" placeholder="请输入邮箱（推荐）" @keyup.enter="handleLogin" />
            </div>
            <div class="login-or-divider"><span>或</span></div>
            <div class="login-field">
              <label>用户名</label>
              <input v-model="loginForm.username" type="text" placeholder="请输入用户名（可选）" @keyup.enter="handleLogin" />
            </div>
            <div class="login-field">
              <label>密码</label>
              <input v-model="loginForm.password" type="password" placeholder="请输入密码" @keyup.enter="handleLogin" />
            </div>
            <div v-if="loginError" class="login-error">{{ loginError }}</div>
            <div class="modal-footer no-border">
              <button type="button" class="cancel-btn" @click="showLoginModal = false">取消</button>
              <button type="button" class="submit-btn" @click="handleLogin" :disabled="loginLoading">
                {{ loginLoading ? '登录中...' : '登录' }}
              </button>
            </div>
          </div>
          <!-- 注册 Tab -->
          <div v-show="authTab === 'register'">
            <p class="login-tip">使用邮箱+密码注册账号，登录后可自定义昵称和头像。</p>
            <div class="login-field">
              <label>邮箱</label>
              <input v-model="registerForm.email" type="email" placeholder="请输入邮箱" @keyup.enter="handleRegister" />
            </div>
            <div class="login-field">
              <label>昵称</label>
              <input v-model="registerForm.nickname" type="text" placeholder="请输入昵称" @keyup.enter="handleRegister" />
            </div>
            <div class="login-field">
              <label>密码</label>
              <input v-model="registerForm.password" type="password" placeholder="至少 6 位" @keyup.enter="handleRegister" />
            </div>
            <div class="login-field">
              <label>确认密码</label>
              <input v-model="registerForm.confirmPassword" type="password" placeholder="再次输入密码" @keyup.enter="handleRegister" />
            </div>
            <div v-if="registerError" class="login-error">{{ registerError }}</div>
            <div v-if="registerSuccess" class="register-success">{{ registerSuccess }}</div>
            <div class="modal-footer no-border">
              <button type="button" class="cancel-btn" @click="showLoginModal = false">取消</button>
              <button type="button" class="submit-btn" @click="handleRegister" :disabled="registerLoading">
                {{ registerLoading ? '注册中...' : '注册' }}
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- 个人设置弹窗 -->
    <div v-if="showProfileModal" class="modal-overlay" @click.self="showProfileModal = false">
      <div class="login-modal profile-modal">
        <div class="modal-header">
          <h3>个人设置</h3>
          <button type="button" class="close-btn" @click="showProfileModal = false">×</button>
        </div>
        <div class="modal-body">
          <div v-if="currentUser" class="profile-info-row">
            <span class="profile-label">邮箱：</span>
            <span class="profile-value">{{ currentUser.email }}</span>
          </div>
          <div v-if="currentUser && currentUser.username" class="profile-info-row">
            <span class="profile-label">账号：</span>
            <span class="profile-value">{{ currentUser.username }}</span>
          </div>
          <div class="login-field">
            <label>昵称</label>
            <input v-model="profileForm.nickname" type="text" placeholder="请输入昵称" maxlength="50" />
          </div>
          <div class="login-field">
            <label>头像 URL</label>
            <input v-model="profileForm.avatar" type="text" placeholder="填写图片链接（可选）" />
          </div>
          <div v-if="profileForm.avatar" class="avatar-preview-row">
            <span class="avatar-preview-label">预览：</span>
            <img class="avatar-preview" :src="profileForm.avatar" @error="profileForm.avatar=''" />
          </div>
          <div class="login-field">
            <label>快速选择头像</label>
            <div class="avatar-grid">
              <div
                v-for="(av, idx) in defaultAvatars"
                :key="idx"
                :class="['avatar-grid-item', { selected: profileForm.avatar === av }]"
                @click="profileForm.avatar = av"
              >
                <img :src="av" />
              </div>
            </div>
          </div>
          <div v-if="profileError" class="login-error">{{ profileError }}</div>
          <div v-if="profileSuccess" class="register-success">{{ profileSuccess }}</div>
        </div>
        <div class="modal-footer">
          <button type="button" class="cancel-btn" @click="showProfileModal = false">取消</button>
          <button type="button" class="submit-btn" @click="handleUpdateProfile" :disabled="profileLoading">
            {{ profileLoading ? '保存中...' : '保存' }}
          </button>
        </div>
      </div>
    </div>

    <!-- 评论列表 -->
    <div class="comment-list-wrapper">
      <div class="comment-header-bar">
        <h3 class="comment-count" @click="focusEditor" title="点击跳转到输入框">
          <span class="comment-count-icon">
            <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" style="vertical-align:middle;margin-right:6px;"><path d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z"></path></svg>
          </span>
          {{ comments.length }} 评论
        </h3>
        <div class="sort-options">
          <button
            :class="['sort-btn', { active: sortBy === 'asc' }]"
            @click="sortBy = 'asc'"
            :disabled="refreshing"
          >按正序</button>
          <button
            :class="['sort-btn', { active: sortBy === 'desc' }]"
            @click="sortBy = 'desc'"
            :disabled="refreshing"
          >按倒序</button>
          <button
            :class="['sort-btn', { active: sortBy === 'hot' }]"
            @click="sortBy = 'hot'"
            :disabled="refreshing"
          >按热度</button>
          <span v-if="refreshing" class="refreshing-tip">刷新中...</span>
        </div>
      </div>

      <div v-if="loading" class="loading">加载中...</div>

      <div v-else-if="comments.length === 0" class="empty">暂无评论，快来抢沙发吧！</div>

      <div v-else class="comment-list" :class="{ refreshing: refreshing }">
        <div v-for="comment in sortedComments" :key="comment.id" class="comment-item">
          <div class="comment-main">
            <div class="avatar-col">
              <div class="avatar">
                <img v-if="comment.avatar || comment.email" :src="comment.avatar || getAvatar(comment.email)" :alt="comment.nickname" />
                <span v-else>{{ comment.nickname.charAt(0).toUpperCase() }}</span>
              </div>
            </div>
            <div class="comment-body">
              <div class="comment-meta">
                <span class="nickname">{{ comment.nickname }}</span>
                <span v-if="comment.is_admin" class="admin-badge">博主</span>
                <span class="time">{{ formatTime(comment.created_at) }}</span>
              </div>
              <div class="ua-info">
                <span v-if="formatOS(comment)" class="ua-text" :title="formatOS(comment)">{{ formatOS(comment) }}</span>
                <span v-if="formatBrowser(comment)" class="ua-text" :title="formatBrowser(comment)">{{ formatBrowser(comment) }}</span>
              </div>
              <div class="comment-content" v-html="renderCommentContent(comment.content)"></div>
              <div class="comment-actions">
                <button class="action-btn like-btn" :class="{ 'liked': likedComments.has(comment.id) }" @click="handleLikeComment(comment)">
                  <span class="icon">
                    <svg v-if="likedComments.has(comment.id)" xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="var(--danger-color)" stroke="var(--danger-color)" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M20.84 4.61a5.5 5.5 0 0 0-7.78 0L12 5.67l-1.06-1.06a5.5 5.5 0 0 0-7.78 7.78l1.06 1.06L12 21.23l7.78-7.78 1.06-1.06a5.5 5.5 0 0 0 0-7.78z"></path></svg>
                    <svg v-else xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M20.84 4.61a5.5 5.5 0 0 0-7.78 0L12 5.67l-1.06-1.06a5.5 5.5 0 0 0-7.78 7.78l1.06 1.06L12 21.23l7.78-7.78 1.06-1.06a5.5 5.5 0 0 0 0-7.78z"></path></svg>
                  </span>
                  <span v-if="(comment.like_count || 0) > 0" class="like-count">{{ comment.like_count }}</span>
                </button>
                <button class="action-btn reply-btn" @click="startReply(comment)">
                  <span class="icon"><svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z"></path></svg></span>
                </button>
              </div>
            </div>
          </div>

          <!-- 回复表单（回复根评论时显示在这里） -->
          <div v-if="replyingTargetId === comment.id" class="reply-form">
            <textarea
              ref="replyInputRef"
              v-model="replyContent"
              :placeholder="`@${replyingTargetNickname}: `"
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
                <div class="avatar-col small">
                  <div class="avatar small">
                    <img v-if="reply.avatar || reply.email" :src="reply.avatar || getAvatar(reply.email)" :alt="reply.nickname" />
                    <span v-else>{{ reply.nickname.charAt(0).toUpperCase() }}</span>
                  </div>
                </div>
                <div class="comment-body">
                  <div class="comment-meta">
                    <span class="nickname">{{ reply.nickname }}</span>
                    <span v-if="reply.is_admin" class="admin-badge">博主</span>
                    <span class="time">{{ formatTime(reply.created_at) }}</span>
                  </div>
                  <div class="ua-info small">
                    <span v-if="formatOS(reply)" class="ua-text small" :title="formatOS(reply)">{{ formatOS(reply) }}</span>
                    <span v-if="formatBrowser(reply)" class="ua-text small" :title="formatBrowser(reply)">{{ formatBrowser(reply) }}</span>
                  </div>
                  <div class="comment-content">
                    <span v-if="reply.reply_to" class="reply-to">@{{ reply.reply_to }}: </span>
                    <span v-html="renderCommentContent(reply.content)"></span>
                  </div>
                  <div class="comment-actions">
                    <button class="action-btn like-btn" :class="{ 'liked': likedComments.has(reply.id) }" @click="handleLikeComment(reply)">
                      <span class="icon">
                        <svg v-if="likedComments.has(reply.id)" xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="var(--danger-color)" stroke="var(--danger-color)" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M20.84 4.61a5.5 5.5 0 0 0-7.78 0L12 5.67l-1.06-1.06a5.5 5.5 0 0 0-7.78 7.78l1.06 1.06L12 21.23l7.78-7.78 1.06-1.06a5.5 5.5 0 0 0 0-7.78z"></path></svg>
                        <svg v-else xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M20.84 4.61a5.5 5.5 0 0 0-7.78 0L12 5.67l-1.06-1.06a5.5 5.5 0 0 0-7.78 7.78l1.06 1.06L12 21.23l7.78-7.78 1.06-1.06a5.5 5.5 0 0 0 0-7.78z"></path></svg>
                      </span>
                      <span v-if="(reply.like_count || 0) > 0" class="like-count">{{ reply.like_count }}</span>
                    </button>
                    <button class="action-btn reply-btn" @click="startReply(comment, reply)">
                      <span class="icon"><svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z"></path></svg></span>
                    </button>
                  </div>
                </div>
              </div>
              <!-- 回复表单（回复此子评论时显示在这里） -->
              <div v-if="replyingTargetId === reply.id" class="reply-form nested">
                <textarea
                  ref="replyInputRef"
                  v-model="replyContent"
                  :placeholder="`@${replyingTargetNickname}: `"
                  rows="3"
                ></textarea>
                <div class="reply-actions">
                  <button class="cancel-btn" @click="cancelReply">取消</button>
                  <button class="submit-btn" @click="submitReply" :disabled="submitting || !replyContent.trim()">
                    提交
                  </button>
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
import { ref, computed, onMounted, onUnmounted, watch, nextTick } from 'vue'
import { useRoute } from 'vue-router'
import { getCommentsByArticle, createComment, likeComment } from '../../api/comment'
import { login, register } from '../../api/auth'
import { getUserProfile, updateUserProfile } from '../../api/user'
import EmojiPicker from 'vue3-emoji-picker'
import 'vue3-emoji-picker/css'
import { getAvatarUrl } from '../../utils/avatar'
import { marked } from 'marked'

const route = useRoute()
const editorRef = ref(null)
const replyInputRef = ref(null)

const comments = ref([])
const loading = ref(false) // 首次加载用的大 loading
const refreshing = ref(false) // 切换排序/刷新时不隐藏列表
const submitting = ref(false)
const replyingTo = ref(null)
const replyingTargetId = ref(null)
const replyingTargetNickname = ref('')
const replyingReplyToId = ref(null)
const replyContent = ref('')
const sortBy = ref('desc')
const isLoggedIn = ref(false)
const currentUser = ref(null)
const showEmojiPanel = ref(false)
const showLoginModal = ref(false)
const loginLoading = ref(false)
const loginError = ref('')
const registerLoading = ref(false)
const registerError = ref('')
const registerSuccess = ref('')
const authTab = ref('login') // 'login' | 'register'

// 个人设置弹窗相关
const showProfileModal = ref(false)
const profileLoading = ref(false)
const profileError = ref('')
const profileSuccess = ref('')
const profileForm = ref({
  nickname: '',
  avatar: ''
})

// 默认头像选项（供用户快速选择）
const defaultAvatars = [
  'https://api.dicebear.com/7.x/avataaars/svg?seed=Felix',
  'https://api.dicebear.com/7.x/avataaars/svg?seed=Aneka',
  'https://api.dicebear.com/7.x/avataaars/svg?seed=Luna',
  'https://api.dicebear.com/7.x/avataaars/svg?seed=Max',
  'https://api.dicebear.com/7.x/avataaars/svg?seed=Zoe',
  'https://api.dicebear.com/7.x/avataaars/svg?seed=Jack',
  'https://api.dicebear.com/7.x/avataaars/svg?seed=Mia',
  'https://api.dicebear.com/7.x/avataaars/svg?seed=Leo'
]

const loginForm = ref({
  email: '',
  username: '',
  password: ''
})

const registerForm = ref({
  email: '',
  nickname: '',
  password: '',
  confirmPassword: ''
})

const form = ref({
  nickname: localStorage.getItem('comment_nickname') || '',
  email: localStorage.getItem('comment_email') || '',
  content: ''
})

// 刷新用户信息（从后端拉取最新资料，保证与数据库一致）
const refreshCurrentUser = async () => {
  try {
    const res = await getUserProfile()
    const u = res.data
    if (u) {
      currentUser.value = u
      localStorage.setItem('comment_user', JSON.stringify(u))
      form.value.nickname = u.nickname || u.username || form.value.nickname
      form.value.email = u.email || form.value.email
    }
  } catch (e) {
    // 静默失败：token失效等情况下交由全局处理
  }
}

const checkLoginStatus = () => {
  const token = localStorage.getItem('token')
  const userStr = localStorage.getItem('comment_user')
  if (token && userStr) {
    try {
      currentUser.value = JSON.parse(userStr)
      isLoggedIn.value = true
      if (currentUser.value) {
        form.value.nickname = currentUser.value.nickname || currentUser.value.username || form.value.nickname
        form.value.email = currentUser.value.email || form.value.email
      }
      // 后台静默刷新用户资料
      refreshCurrentUser()
    } catch (e) {
      isLoggedIn.value = false
      currentUser.value = null
    }
  }
}

const handleLogin = async () => {
  const hasEmail = loginForm.value.email && loginForm.value.email.trim()
  const hasUsername = loginForm.value.username && loginForm.value.username.trim()
  if (!hasEmail && !hasUsername) {
    loginError.value = '请输入邮箱或用户名'
    return
  }
  if (!loginForm.value.password) {
    loginError.value = '请输入密码'
    return
  }
  loginLoading.value = true
  loginError.value = ''
  try {
    const payload = { password: loginForm.value.password }
    if (hasEmail) payload.email = loginForm.value.email.trim()
    if (hasUsername) payload.username = loginForm.value.username.trim()
    const res = await login(payload)
    const { token, user } = res.data || {}
    if (token && user) {
      localStorage.setItem('token', token)
      localStorage.setItem('comment_user', JSON.stringify(user))
      currentUser.value = user
      isLoggedIn.value = true
      form.value.nickname = user.nickname || user.username || form.value.nickname
      form.value.email = user.email || form.value.email
      showLoginModal.value = false
      loginForm.value = { email: '', username: '', password: '' }
    }
  } catch (err) {
    const msg = err.response?.data?.message || err.response?.data?.msg || '登录失败，请检查邮箱和密码'
    loginError.value = msg
  } finally {
    loginLoading.value = false
  }
}

const handleRegister = async () => {
  if (!registerForm.value.email || !registerForm.value.email.trim()) {
    registerError.value = '请输入邮箱'
    return
  }
  if (!registerForm.value.nickname || !registerForm.value.nickname.trim()) {
    registerError.value = '请输入昵称'
    return
  }
  if (!registerForm.value.password || registerForm.value.password.length < 6) {
    registerError.value = '密码至少 6 位'
    return
  }
  if (registerForm.value.password !== registerForm.value.confirmPassword) {
    registerError.value = '两次密码输入不一致'
    return
  }
  registerLoading.value = true
  registerError.value = ''
  registerSuccess.value = ''
  try {
    await register({
      email: registerForm.value.email.trim(),
      nickname: registerForm.value.nickname.trim(),
      password: registerForm.value.password
    })
    registerSuccess.value = '注册成功！正在自动登录…'
    // 注册成功后直接用邮箱+密码登录
    setTimeout(async () => {
      try {
        const res = await login({
          email: registerForm.value.email.trim(),
          password: registerForm.value.password
        })
        const { token, user } = res.data || {}
        if (token && user) {
          localStorage.setItem('token', token)
          localStorage.setItem('comment_user', JSON.stringify(user))
          currentUser.value = user
          isLoggedIn.value = true
          form.value.nickname = user.nickname || user.username || form.value.nickname
          form.value.email = user.email || form.value.email
          showLoginModal.value = false
          registerForm.value = { email: '', nickname: '', password: '', confirmPassword: '' }
          registerSuccess.value = ''
        }
      } catch (e) {
        // 自动登录失败，切到登录Tab让用户手动登录
        authTab.value = 'login'
        loginForm.value.email = registerForm.value.email
        loginError.value = '注册成功，请手动登录'
      } finally {
        registerLoading.value = false
      }
    }, 600)
  } catch (err) {
    registerLoading.value = false
    const msg = err.response?.data?.message || err.response?.data?.msg || '注册失败，请稍后重试'
    registerError.value = msg
  }
}

// 打开个人设置弹窗
const openProfileModal = () => {
  if (!currentUser.value) return
  profileForm.value.nickname = currentUser.value.nickname || ''
  profileForm.value.avatar = currentUser.value.avatar || ''
  profileError.value = ''
  profileSuccess.value = ''
  showProfileModal.value = true
}

// 提交个人资料修改
const handleUpdateProfile = async () => {
  if (!profileForm.value.nickname || !profileForm.value.nickname.trim()) {
    profileError.value = '昵称不能为空'
    return
  }
  profileLoading.value = true
  profileError.value = ''
  profileSuccess.value = ''
  try {
    const payload = { nickname: profileForm.value.nickname.trim() }
    if (profileForm.value.avatar && profileForm.value.avatar.trim()) {
      payload.avatar = profileForm.value.avatar.trim()
    }
    await updateUserProfile(payload)
    profileSuccess.value = '修改成功！'
    // 刷新用户信息
    await refreshCurrentUser()
    setTimeout(() => {
      showProfileModal.value = false
    }, 600)
  } catch (err) {
    const msg = err.response?.data?.message || err.response?.data?.msg || '修改失败，请稍后重试'
    profileError.value = msg
  } finally {
    profileLoading.value = false
  }
}

const handleLogout = () => {
  localStorage.removeItem('token')
  localStorage.removeItem('comment_user')
  isLoggedIn.value = false
  currentUser.value = null
  showProfileModal.value = false
  form.value.nickname = localStorage.getItem('comment_nickname') || ''
  form.value.email = localStorage.getItem('comment_email') || ''
}

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

// 排序后的评论：后端已按 sort_by 排序完成，前端直接使用，保证分页正确
// 仅在数据异常时做兜底返回
const sortedComments = computed(() => {
  if (!Array.isArray(comments.value)) return []
  return comments.value
})

// 排序方式变化时，重新向后端请求排好序的数据
watch(sortBy, () => {
  fetchComments()
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
const isFirstLoad = ref(true)
const fetchComments = async () => {
  if (isFirstLoad.value) {
    loading.value = true
  } else {
    refreshing.value = true
  }
  try {
    const articleId = route.params.id || route.params.slug
    const res = await getCommentsByArticle(articleId, { page: 1, page_size: 100, sort_by: sortBy.value })
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
    isFirstLoad.value = false
  } catch (error) {
    console.error('获取评论失败:', error)
  } finally {
    loading.value = false
    refreshing.value = false
  }
}

// 提交评论
const submitComment = async () => {
  if (!form.value.nickname || !form.value.content.trim()) return

  submitting.value = true
  try {
    const articleSlug = route.params.slug || route.params.id
    // 精确检测当前客户端（解决 Win11 识别不准等问题）
    const client = await detectClientInfo()
    await createComment({
      article_slug: articleSlug,
      nickname: form.value.nickname,
      email: form.value.email,
      content: form.value.content,
      os: client.os,
      os_version: client.os_version,
      browser: client.browser,
      browser_version: client.browser_version
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
  if (reply) {
    replyingTargetId.value = reply.id
    replyingTargetNickname.value = reply.nickname
    replyingReplyToId.value = reply.id
  } else {
    replyingTargetId.value = comment.id
    replyingTargetNickname.value = comment.nickname
    replyingReplyToId.value = comment.id
  }
  replyContent.value = ''
  nextTick(() => {
    if (replyInputRef.value) {
      const inputs = Array.isArray(replyInputRef.value) ? replyInputRef.value : [replyInputRef.value]
      const target = inputs.find(el => el && el.offsetParent !== null) || inputs[0]
      if (target) {
        target.focus()
      }
    }
  })
}

// 取消回复
const cancelReply = () => {
  replyingTo.value = null
  replyingTargetId.value = null
  replyingTargetNickname.value = ''
  replyingReplyToId.value = null
  replyContent.value = ''
}

// 提交回复
const submitReply = async () => {
  if (!replyContent.value.trim() || !replyingTo.value) return

  submitting.value = true
  try {
    const articleSlug = route.params.slug || route.params.id
    // 精确检测当前客户端（解决 Win11 识别不准等问题）
    const client = await detectClientInfo()
    await createComment({
      article_slug: articleSlug,
      nickname: form.value.nickname || '匿名用户',
      email: form.value.email,
      content: replyContent.value,
      parent_id: replyingTo.value.id,
      reply_to_id: replyingReplyToId.value,
      os: client.os,
      os_version: client.os_version,
      browser: client.browser,
      browser_version: client.browser_version
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

// 格式化时间：<=3天显示相对时间，>3天显示具体日期 YYYY-MM-DD（不显示时分）
const pad = (n) => String(n).padStart(2, '0')
const formatFullDate = (date) => {
  const y = date.getFullYear()
  const m = pad(date.getMonth() + 1)
  const d = pad(date.getDate())
  return `${y}-${m}-${d}`
}

const formatTime = (dateStr) => {
  if (!dateStr) return ''
  const date = new Date(dateStr)
  if (isNaN(date.getTime())) return dateStr
  const now = new Date()
  const diff = now - date

  const THREE_DAYS = 3 * 24 * 60 * 60 * 1000
  if (diff >= THREE_DAYS) return formatFullDate(date)

  if (diff < 60 * 1000) return '刚刚'
  if (diff < 60 * 60 * 1000) return Math.floor(diff / (60 * 1000)) + ' 分钟前'
  if (diff < 24 * 60 * 60 * 1000) return Math.floor(diff / (60 * 60 * 1000)) + ' 小时前'
  return Math.floor(diff / (24 * 60 * 60 * 1000)) + ' 天前'
}

// 格式化操作系统 + 版本（独立代码块用）
const formatOS = (item) => {
  if (!item.os || item.os === '未知') return ''
  const name = item.os
  const ver = (item.os_version || '').trim()
  return ver ? `${name} ${ver}` : name
}

// 格式化浏览器 + 版本（独立代码块用）
const formatBrowser = (item) => {
  if (!item.browser || item.browser === '未知') return ''
  const name = item.browser
  const ver = (item.browser_version || '').trim()
  return ver ? `${name} ${ver}` : name
}

// —— 版本号工具：强制「主版本.次版本」两位（例如 11→11.0，126.0.6478→126.0，17.5→17.5）——
const toTwoPartVersion = (raw) => {
  if (!raw) return ''
  const s = String(raw).trim()
  if (!s) return ''
  if (s.includes('.')) {
    const parts = s.split('.').filter(p => p !== '')
    if (parts.length === 1) return parts[0] + '.0'
    const major = parts[0] || '0'
    const minor = /^\d+$/.test(parts[1]) ? parts[1] : '0'
    return `${major}.${minor}`
  }
  return /^\d+$/.test(s) ? `${s}.0` : s
}

// —— 精确检测客户端信息（区分 Win10/Win11、浏览器版本）——
// 优先用 navigator.userAgentData.getHighEntropyValues() 新 API，
// 兜底用 UA 字符串解析。
const detectClientInfo = async () => {
  const info = {
    os: '', os_version: '',
    browser: '', browser_version: ''
  }
  const ua = navigator.userAgent || ''

  // 第 1 步：新 API（Chromium / Edge / 新版 Safari 支持）—— 精确区分 Win11
  try {
    if (navigator.userAgentData && typeof navigator.userAgentData.getHighEntropyValues === 'function') {
      const hint = await navigator.userAgentData.getHighEntropyValues([
        'platform', 'platformVersion', 'fullVersionList'
      ])
      const platform = (hint.platform || '').toLowerCase()
      const platformVersion = hint.platformVersion || ''

      // —— 操作系统精确识别 ——
      if (platform === 'windows') {
        info.os = 'Windows'
        // platformVersion: Win10 = "13.0.xxxxx"; Win11 = "14.0.22621+" / "15.x.xxxxx"
        const parts = platformVersion.split('.').map(s => Number.isNaN(parseInt(s, 10)) ? 0 : parseInt(s, 10))
        const major = parts[0] || 0
        const build = parts[2] || 0
        let rawVer = ''
        if (major >= 14 || build >= 22000) {
          rawVer = '11'
        } else if (major === 13 || (build > 0 && build < 22000)) {
          rawVer = '10'
        } else if (major > 0) {
          rawVer = String(major)
        }
        info.os_version = toTwoPartVersion(rawVer)
      } else if (platform === 'macos') {
        info.os = 'macOS'
        const v = platformVersion.replace(/_/g, '.')
        const major = v.split('.')[0]
        const twoPart = toTwoPartVersion(v)
        const named = {
          '11': '(Big Sur)', '12': '(Monterey)',
          '13': '(Ventura)', '14': '(Sonoma)', '15': '(Sequoia)'
        }
        const name = named[major] || ''
        info.os_version = name ? `${twoPart} ${name}` : twoPart
      } else if (platform === 'ios' || platform === 'iphone' || platform === 'ipod') {
        info.os = 'iOS'
        info.os_version = toTwoPartVersion(platformVersion.replace(/_/g, '.'))
      } else if (platform === 'ipados' || platform === 'ipad') {
        info.os = 'iPadOS'
        info.os_version = toTwoPartVersion(platformVersion.replace(/_/g, '.'))
      } else if (['linux', 'ubuntu', 'fedora', 'debian'].includes(platform)) {
        const map = { linux: 'Linux', ubuntu: 'Ubuntu', fedora: 'Fedora', debian: 'Debian' }
        info.os = map[platform]
        info.os_version = toTwoPartVersion(platformVersion.replace(/_/g, '.'))
      } else if (platform === 'android') {
        info.os = 'Android'
        info.os_version = toTwoPartVersion(platformVersion)
      } else if (platform.toLowerCase().includes('harmony')) {
        info.os = '鸿蒙'
        info.os_version = toTwoPartVersion(platformVersion)
      }

      // —— 浏览器精确识别（fullVersionList 优先级：Edge > Opera > Vivaldi > Firefox > Chrome > Safari）——
      const list = Array.isArray(hint.fullVersionList) ? hint.fullVersionList : []
      const order = [
        { k: 'Microsoft Edge', o: 'Edge' },
        { k: 'Edg', o: 'Edge' },
        { k: 'Opera', o: 'Opera' },
        { k: 'OPR', o: 'Opera' },
        { k: 'Vivaldi', o: 'Vivaldi' },
        { k: 'Firefox', o: 'Firefox' },
        { k: 'Brave', o: 'Brave' },
        { k: 'Chrome', o: 'Chrome' },
        { k: 'Safari', o: 'Safari' },
      ]
      for (const rule of order) {
        const found = list.find(b => b && b.brand === rule.k)
        if (found && found.version) {
          info.browser = rule.o
          info.browser_version = toTwoPartVersion(found.version)
          break
        }
      }
    }
  } catch (e) {
    // 静默失败，走兜底解析
  }

  // —— 内嵌 App（微信/钉钉/支付宝）需要 UA 补充 ——
  if (ua.includes('MicroMessenger')) {
    info.browser = '微信'
    const m = ua.match(/MicroMessenger\/([\d.]+)/)
    info.browser_version = m ? toTwoPartVersion(m[1]) : ''
  } else if (ua.includes('DingTalk')) {
    info.browser = '钉钉'
    const m = ua.match(/DingTalk\/([\d.]+)/)
    info.browser_version = m ? toTwoPartVersion(m[1]) : ''
  } else if (ua.includes('AlipayClient')) {
    info.browser = '支付宝'
    const m = ua.match(/AlipayClient\/([\d.]+)/)
    info.browser_version = m ? toTwoPartVersion(m[1]) : ''
  }

  // 第 2 步：兜底（没拿到才用 UA 解析）
  if (!info.os) {
    const o = parseOSFromUA(ua)
    info.os = o.os
    info.os_version = o.version
  }
  if (!info.browser) {
    const b = parseBrowserFromUA(ua)
    info.browser = b.browser
    info.browser_version = b.version
  }

  return info
}

// 从 UA 兜底解析 OS（老浏览器）
const parseOSFromUA = (ua) => {
  let os = '', version = ''
  switch (true) {
    case ua.includes('Windows'):
      os = 'Windows'
      if (ua.includes('Windows NT 10.0')) version = toTwoPartVersion('10')
      else if (ua.includes('Windows NT 6.3')) version = '8.1'
      else if (ua.includes('Windows NT 6.2')) version = toTwoPartVersion('8')
      else if (ua.includes('Windows NT 6.1')) version = toTwoPartVersion('7')
      else if (ua.includes('Windows NT 5.1')) version = 'XP'
      break
    case ua.includes('Mac OS X') && !ua.includes('iPhone') && !ua.includes('iPad'):
      os = 'macOS'
      {
        const m = ua.match(/Mac OS X ([\d_.]+)/)
        if (m) {
          const v = m[1].replace(/_/g, '.')
          const major = v.split('.')[0]
          const twoPart = toTwoPartVersion(v)
          const named = {
            '11': '(Big Sur)', '12': '(Monterey)',
            '13': '(Ventura)', '14': '(Sonoma)', '15': '(Sequoia)'
          }
          version = named[major] ? `${twoPart} ${named[major]}` : twoPart
        }
      }
      break
    case ua.includes('iPhone') || ua.includes('iPod'):
      os = 'iOS'
      {
        const m = ua.match(/(?:iPhone OS|OS) ([\d_]+)/)
        if (m) version = toTwoPartVersion(m[1].replace(/_/g, '.'))
      }
      break
    case ua.includes('iPad'):
      os = 'iPadOS'
      {
        const m = ua.match(/OS ([\d_]+)/)
        if (m) version = toTwoPartVersion(m[1].replace(/_/g, '.'))
      }
      break
    case ua.includes('HarmonyOS') || ua.includes('Harmony'):
      os = '鸿蒙'
      {
        const m = ua.match(/Harmony(?:OS)?\/([\d.]+)/)
        if (m) version = toTwoPartVersion(m[1])
      }
      break
    case ua.includes('Android'):
      os = 'Android'
      {
        const m = ua.match(/Android ([\d.]+)/)
        if (m) version = toTwoPartVersion(m[1])
      }
      break
    case ua.includes('Ubuntu'):
      os = 'Ubuntu'
      {
        const m = ua.match(/Ubuntu\/([\d.]+)/)
        if (m) version = toTwoPartVersion(m[1])
      }
      break
    case ua.includes('Fedora'): os = 'Fedora'; break
    case ua.includes('Debian'): os = 'Debian'; break
    case ua.includes('Linux'): os = 'Linux'; break
  }
  return { os, version }
}

// 从 UA 兜底解析 Browser（老浏览器）
const parseBrowserFromUA = (ua) => {
  let browser = '', version = ''
  const extractVer = (prefix) => {
    const idx = ua.indexOf(prefix)
    if (idx < 0) return ''
    let rest = ua.slice(idx + prefix.length)
    let end = rest.length
    for (let i = 0; i < rest.length; i++) {
      const ch = rest[i]
      if (ch === ' ' || ch === ';' || ch === ')') { end = i; break }
    }
    const v = rest.slice(0, end).trim()
    return toTwoPartVersion(v)
  }
  switch (true) {
    case ua.includes('MicroMessenger'):
      browser = '微信'
      version = extractVer('MicroMessenger/')
      break
    case ua.includes('QQBrowser'):
      browser = 'QQ浏览器'
      version = extractVer('QQBrowser/')
      break
    case ua.includes('UCBrowser'):
      browser = 'UC浏览器'
      version = extractVer('UCBrowser/')
      break
    case ua.includes('DingTalk'):
      browser = '钉钉'
      version = extractVer('DingTalk/')
      break
    case ua.includes('AlipayClient'):
      browser = '支付宝'
      version = extractVer('AlipayClient/')
      break
    case ua.includes('Edg/') || ua.includes('Edge/'):
      browser = 'Edge'
      version = extractVer('Edg/') || extractVer('Edge/')
      break
    case ua.includes('OPR/') || ua.includes('Opera/'):
      browser = 'Opera'
      version = extractVer('OPR/') || extractVer('Opera/')
      break
    case ua.includes('Vivaldi/'):
      browser = 'Vivaldi'
      version = extractVer('Vivaldi/')
      break
    case ua.includes('Firefox/') || ua.includes('FxiOS/'):
      browser = 'Firefox'
      version = extractVer('Firefox/') || extractVer('FxiOS/')
      break
    case ua.includes('CriOS/'):
      browser = 'Chrome'
      version = extractVer('CriOS/')
      break
    case ua.includes('MSIE '):
      browser = 'IE'
      version = extractVer('MSIE ')
      break
    case ua.includes('Trident/'):
      browser = 'IE'
      {
        const m = ua.match(/rv:([\d.]+)/)
        if (m) version = toTwoPartVersion(m[1])
      }
      break
    case ua.includes('Chrome/') && !ua.includes('Edg/'):
      browser = 'Chrome'
      version = extractVer('Chrome/')
      break
    case ua.includes('Safari/') && !ua.includes('Chrome/') && !ua.includes('Edg/'):
      browser = 'Safari'
      version = extractVer('Version/')
      break
  }
  return { browser, version }
}

onMounted(() => {
  checkLoginStatus()
  fetchComments()
  document.addEventListener('click', handleClickOutside)
})

onUnmounted(() => {
  document.removeEventListener('click', handleClickOutside)
})

// 暴露给父组件的方法：聚焦到主编辑器
const focusEditor = () => {
  if (editorRef.value) {
    editorRef.value.focus()
  }
}

defineExpose({
  focusEditor
})
</script>

<style scoped>
.comment-section {
  margin-top: 32px;
  background: var(--card-background);
  border: 1px solid transparent;
  border-radius: 10px;
  box-shadow: 0 1px 3px rgba(0,0,0,0.12), 0 1px 2px rgba(0,0,0,0.24);
  padding: 24px;
}

/* 登录弹窗 */
.modal-overlay {
  position: fixed;
  inset: 0;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
  animation: fade-in 0.2s ease;
}

@keyframes fade-in {
  from { opacity: 0; }
  to { opacity: 1; }
}

.login-modal {
  background: var(--card-background);
  border-radius: 12px;
  width: 90%;
  max-width: 400px;
  box-shadow: 0 20px 60px rgba(0, 0, 0, 0.3);
  overflow: hidden;
  animation: slide-up 0.25s ease;
}

@keyframes slide-up {
  from { transform: translateY(20px); opacity: 0; }
  to { transform: translateY(0); opacity: 1; }
}

.modal-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 16px 20px;
  border-bottom: 1px solid var(--card-separator-color);
}

.modal-header h3 {
  margin: 0;
  font-size: 1.1rem;
  color: var(--card-text-color-main);
}

.close-btn {
  background: none;
  border: none;
  font-size: 1.5rem;
  color: var(--card-text-color-tertiary);
  cursor: pointer;
  padding: 0 4px;
  line-height: 1;
}

.close-btn:hover {
  color: var(--card-text-color-main);
}

.modal-body {
  padding: 20px;
}

.login-tip {
  background: rgba(var(--success-color-rgb), 0.1);
  color: var(--success-color-darker);
  padding: 10px 12px;
  border-radius: 6px;
  font-size: 0.85rem;
  margin: 0 0 16px 0;
  border: 1px solid rgba(var(--success-color-rgb), 0.3);
}

.login-field {
  margin-bottom: 14px;
}

.login-field label {
  display: block;
  font-size: 0.9rem;
  color: var(--card-text-color-secondary);
  margin-bottom: 6px;
}

.login-field input {
  width: 100%;
  padding: 10px 12px;
  border: 1px solid var(--card-separator-color);
  border-radius: 6px;
  font-size: 0.95rem;
  box-sizing: border-box;
  transition: border-color 0.2s;
}

.login-field input:focus {
  outline: none;
  border-color: var(--success-color);
}

.login-error {
  background: rgba(var(--danger-color-rgb), 0.08);
  color: var(--danger-color);
  padding: 8px 12px;
  border-radius: 6px;
  font-size: 0.85rem;
  margin-top: 8px;
  border: 1px solid rgba(var(--danger-color-rgb), 0.2);
}

.register-success {
  background: rgba(var(--success-color-rgb), 0.1);
  color: var(--success-color-darker);
  padding: 8px 12px;
  border-radius: 6px;
  font-size: 0.85rem;
  margin-top: 8px;
  border: 1px solid rgba(var(--success-color-rgb), 0.3);
}

/* auth 弹窗 Tab */
.auth-modal .modal-header {
  flex-direction: row;
}
.auth-tabs {
  display: flex;
  gap: 20px;
}
.auth-tab {
  background: none;
  border: none;
  padding: 6px 2px;
  font-size: 1rem;
  color: var(--card-text-color-tertiary);
  cursor: pointer;
  font-weight: 500;
  border-bottom: 2px solid transparent;
  transition: all 0.2s;
}
.auth-tab:hover {
  color: var(--card-text-color-main);
}
.auth-tab.active {
  color: var(--success-color);
  border-bottom-color: var(--success-color);
}

/* 登录中的「或」分隔线 */
.login-or-divider {
  position: relative;
  margin: 2px 0 12px 0;
  text-align: center;
}
.login-or-divider::before {
  content: '';
  position: absolute;
  top: 50%;
  left: 0;
  right: 0;
  height: 1px;
  background: var(--card-separator-color);
}
.login-or-divider span {
  position: relative;
  background: var(--card-background);
  padding: 0 10px;
  font-size: 0.8rem;
  color: var(--card-text-color-tertiary);
}

/* Tab 内嵌的 footer，去掉上下边框和背景 */
.modal-footer.no-border {
  border: none;
  background: none;
  padding: 4px 0 0 0;
  margin-top: 6px;
}

/* 个人设置弹窗样式 */
.profile-modal {
  max-width: 460px;
}
.profile-info-row {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 8px 12px;
  background: rgba(var(--accent-color-rgb), 0.04);
  border-radius: 6px;
  margin-bottom: 14px;
  font-size: 0.9rem;
}
.profile-label {
  color: var(--card-text-color-tertiary);
  white-space: nowrap;
}
.profile-value {
  color: var(--card-text-color-main);
  font-weight: 500;
  word-break: break-all;
}
.avatar-preview-row {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 8px 0 14px 0;
}
.avatar-preview-label {
  color: var(--card-text-color-tertiary);
  font-size: 0.85rem;
  white-space: nowrap;
}
.avatar-preview {
  width: 52px;
  height: 52px;
  border-radius: 50%;
  object-fit: cover;
  border: 2px solid var(--card-separator-color);
  background: var(--card-background);
}
.avatar-grid {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 10px;
  margin-top: 6px;
}
.avatar-grid-item {
  width: 100%;
  aspect-ratio: 1 / 1;
  border-radius: 10px;
  overflow: hidden;
  border: 2px solid transparent;
  cursor: pointer;
  background: var(--card-background);
  transition: all 0.2s;
  padding: 4px;
  box-sizing: border-box;
}
.avatar-grid-item:hover {
  border-color: var(--success-color);
  transform: translateY(-1px);
}
.avatar-grid-item.selected {
  border-color: var(--success-color);
  box-shadow: 0 0 0 3px rgba(var(--success-color-rgb), 0.2);
}
.avatar-grid-item img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  border-radius: 6px;
}

/* 个人设置按钮 */
.profile-btn {
  width: 32px;
  height: 32px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: transparent;
  color: var(--card-text-color-secondary);
  border: 1px solid var(--card-separator-color);
  border-radius: 6px;
  font-size: 0.85rem;
  cursor: pointer;
  transition: all 0.2s;
}
.profile-btn:hover {
  background: rgba(var(--success-color-rgb), 0.08);
  color: var(--success-color);
  border-color: rgba(var(--success-color-rgb), 0.3);
}

.modal-footer {
  display: flex;
  justify-content: flex-end;
  gap: 10px;
  padding: 14px 20px;
  border-top: 1px solid var(--card-separator-color);
  background: rgba(var(--accent-color-rgb), 0.02);
}

.modal-footer .cancel-btn {
  padding: 8px 18px;
  background: var(--card-background);
  color: var(--card-text-color-secondary);
  border: 1px solid var(--card-separator-color);
  border-radius: 6px;
  font-size: 0.9rem;
  cursor: pointer;
  transition: all 0.2s;
}

.modal-footer .cancel-btn:hover {
  background: rgba(var(--accent-color-rgb), 0.04);
}

.modal-footer .submit-btn {
  padding: 8px 20px;
  background: var(--success-color);
  color: white;
  border: none;
  border-radius: 6px;
  font-size: 0.9rem;
  cursor: pointer;
  transition: opacity 0.2s;
  font-weight: 500;
}

.modal-footer .submit-btn:hover:not(:disabled) {
  opacity: 0.9;
}

.modal-footer .submit-btn:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

/* 退出按钮样式 */
.logout-btn {
  padding: 6px 12px;
  background: transparent;
  color: var(--card-text-color-secondary);
  border: 1px solid var(--card-separator-color);
  border-radius: 6px;
  font-size: 0.85rem;
  cursor: pointer;
  transition: all 0.2s;
  display: flex;
  align-items: center;
  gap: 8px;
}

.logout-btn:hover {
  background: rgba(var(--danger-color-rgb), 0.08);
  color: var(--danger-color);
  border-color: rgba(var(--danger-color-rgb), 0.2);
}

.user-info {
  display: flex;
  align-items: center;
  gap: 6px;
  padding-right: 6px;
  border-right: 1px solid var(--card-separator-color);
  margin-right: 2px;
}

.user-avatar {
  width: 24px;
  height: 24px;
  border-radius: 50%;
  overflow: hidden;
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg, var(--success-color) 0%, var(--success-color-darker) 100%);
}

.user-avatar img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.user-name {
  font-weight: 500;
  font-size: 0.85rem;
  color: var(--card-text-color-main);
}

.form-field input:disabled {
  background: rgba(var(--accent-color-rgb), 0.04);
  color: var(--card-text-color-tertiary);
  cursor: not-allowed;
}

/* 评论表单内嵌卡片 */
.comment-form-card {
  background: var(--card-background);
  border: 1px solid var(--card-separator-color);
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
  border-bottom: 1px dashed var(--card-separator-color);
}

.form-field {
  flex: 1;
  display: flex;
  align-items: center;
  gap: 8px;
}

.form-field:first-child {
  border-right: 1px dashed var(--card-separator-color);
  padding-right: 16px;
  margin-right: 16px;
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

.editor-input {
  width: 100%;
  min-height: 100px;
  max-height: 300px;
  overflow-y: auto;
  border: none;
  background: transparent;
  color: var(--card-text-color-main);
  font-size: 0.95rem;
  font-family: inherit;
  line-height: 1.6;
  word-wrap: break-word;
  outline: none;
}

.editor-input:empty::before {
  content: attr(data-placeholder);
  color: var(--card-text-color-tertiary);
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
  border-top: 1px solid var(--card-separator-color);
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
  color: var(--card-text-color-secondary);
  transition: all 0.2s;
}

.tool-btn:hover {
  color: var(--success-color);
}

.tool-btn:hover {
  background: rgba(var(--success-color-rgb), 0.1);
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
  background: var(--card-background);
  border: 1px solid var(--card-separator-color);
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
  color: var(--card-text-color-tertiary);
}

.login-btn {
  padding: 8px 16px;
  background: transparent;
  color: var(--card-text-color-secondary);
  border: 1px solid var(--card-separator-color);
  border-radius: 6px;
  font-size: 0.9rem;
  cursor: pointer;
  transition: all 0.2s;
}

.login-btn:hover {
  background: rgba(var(--accent-color-rgb), 0.04);
}

.submit-btn {
  padding: 8px 20px;
  background: var(--success-color);
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
  color: var(--card-text-color-main);
  margin: 0;
  cursor: pointer;
  transition: color 0.2s;
}

.comment-count:hover {
  color: var(--success-color);
}

.sort-options {
  display: flex;
  gap: 16px;
}

.sort-btn {
  background: none;
  border: none;
  color: var(--card-text-color-tertiary);
  font-size: 0.9rem;
  cursor: pointer;
  padding: 4px 8px;
  border-radius: 4px;
  transition: all 0.2s;
}

.sort-btn:hover {
  color: var(--success-color);
}

.sort-btn.active {
  color: var(--success-color);
  font-weight: 500;
}

.sort-btn:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.refreshing-tip {
  font-size: 0.8rem;
  color: var(--success-color);
  margin-left: 8px;
  animation: pulse 1.2s ease-in-out infinite;
}

@keyframes pulse {
  0%, 100% { opacity: 0.5; }
  50% { opacity: 1; }
}

.loading,
.empty {
  text-align: center;
  padding: 40px;
  color: var(--card-text-color-tertiary);
}

.comment-list {
  display: flex;
  flex-direction: column;
  transition: opacity 0.25s ease;
}

.comment-list.refreshing {
  opacity: 0.6;
  pointer-events: none;
}

.comment-item {
  padding: 16px 0;
  border-bottom: 1px solid var(--card-separator-color);
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
  background: linear-gradient(135deg, var(--success-color) 0%, var(--success-color-darker) 100%);
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

.avatar-col {
  display: flex;
  flex-direction: column;
  align-items: center;
  flex-shrink: 0;
}

.ua-info {
  display: inline-flex;
  align-items: center;
  gap: 6px;                    /* 两个代码块之间的间距 */
  flex-wrap: wrap;             /* 窄屏自动换行 */
  margin-top: 4px;
  margin-bottom: 6px;
}

.ua-info.small {
  gap: 4px;                    /* 回复的代码块间距更小 */
  margin-top: 2px;
  margin-bottom: 4px;
}

.ua-text {
  display: inline-block;
  /* Markdown 行内代码块风格 */
  font-family: ui-monospace, SFMono-Regular, Menlo, Monaco, Consolas,
               "Liberation Mono", "Courier New", "PingFang SC", monospace;
  font-size: 0.75rem;
  font-weight: 500;
  line-height: 1.4;
  padding: 2px 8px;
  background: rgba(var(--accent-color-rgb), 0.04);              /* 浅色背景 */
  color: var(--card-text-color-secondary);
  border: 1px solid var(--card-separator-color);
  border-radius: 6px;
  white-space: nowrap;
  /* 代码块那股味儿的细节：轻微内阴影 + 过渡 */
  box-shadow: inset 0 1px 2px rgba(17, 24, 39, 0.03);
  transition: background 0.15s ease, border-color 0.15s ease;
}

.ua-text:hover {
  background: rgba(var(--accent-color-rgb), 0.06);
  border-color: var(--accent-color);
}

.ua-text.small {
  font-size: 0.68rem;
  padding: 1.5px 6px;
  border-radius: 5px;
}

/* 深色模式适配（如果项目启用了暗色主题自动切换） */
@media (prefers-color-scheme: dark) {
  .ua-text {
    background: #1f2937;
    color: #e5e7eb;
    border-color: #374151;
    box-shadow: inset 0 1px 2px rgba(0, 0, 0, 0.3);
  }
  .ua-text:hover {
    background: #273244;
    border-color: #4b5563;
  }
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
  color: var(--card-text-color-main);
  font-size: 0.9rem;
}

.admin-badge {
  display: inline-block;
  padding: 1px 6px;
  background: var(--success-color);
  color: white;
  font-size: 0.7rem;
  border-radius: 3px;
  font-weight: 500;
}

.time {
  font-size: 0.8rem;
  color: var(--card-text-color-tertiary);
}

.comment-content {
  color: var(--card-text-color-main);
  line-height: 1.6;
  margin-bottom: 8px;
  word-wrap: break-word;
  font-size: 0.95rem;
}

.reply-to {
  color: var(--success-color);
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
  color: var(--card-text-color-tertiary);
  font-size: 0.8rem;
  cursor: pointer;
  padding: 2px 4px;
  border-radius: 4px;
  transition: all 0.2s;
}

.action-btn:hover {
  color: var(--success-color);
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
  color: var(--danger-color);
}

.like-btn.liked {
  color: var(--danger-color);
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
  border: 1px solid var(--card-separator-color);
  border-radius: 6px;
  background: var(--card-background);
  color: var(--card-text-color-main);
  font-size: 0.9rem;
  resize: vertical;
  min-height: 80px;
  box-sizing: border-box;
  font-family: inherit;
}

.reply-form textarea:focus {
  outline: none;
  border-color: var(--success-color);
}

.reply-form textarea::placeholder {
  color: var(--card-text-color-tertiary);
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
  color: var(--card-text-color-secondary);
  border: 1px solid var(--card-separator-color);
  border-radius: 4px;
  font-size: 0.85rem;
  cursor: pointer;
  transition: all 0.2s;
}

.cancel-btn:hover {
  background: rgba(var(--accent-color-rgb), 0.04);
}

.reply-form .submit-btn {
  padding: 6px 14px;
  font-size: 0.85rem;
  background: var(--success-color);
}

/* 嵌套在子回复列表中的回复表单，不需要左侧缩进（reply-list 已经有 margin-left + padding-left + border-left）*/
.reply-form.nested {
  margin-left: 0;
  margin-top: 10px;
}

/* 子评论列表 */
.reply-list {
  margin-top: 12px;
  margin-left: 52px;
  padding-left: 12px;
  border-left: 2px solid var(--card-separator-color);
}

.reply-item {
  padding: 10px 0;
}

.reply-item:not(:last-child) {
  border-bottom: 1px solid var(--card-separator-color);
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
    border-bottom: 1px dashed var(--card-separator-color);
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

/* ===== 暗色模式适配（手动主题切换）===== */
[data-scheme="dark"] .comment-section {
  background: #1c2128;
  border: 1px solid #30363d;
  box-shadow: 0 1px 3px rgba(0,0,0,0.3), 0 1px 2px rgba(0,0,0,0.5);
}

[data-scheme="dark"] .comment-form-card {
  background: #22272e;
  border: 1px solid #30363d;
}

[data-scheme="dark"] .form-header {
  border-bottom: 1px dashed #383838;
}

[data-scheme="dark"] .form-field:first-child {
  border-right: 1px dashed #383838;
}

[data-scheme="dark"] .form-field label {
  color: #939ca6;
}

[data-scheme="dark"] .form-field input {
  color: #dadee3;
}

[data-scheme="dark"] .form-field input::placeholder {
  color: #6e7681;
}

[data-scheme="dark"] .form-field input:disabled {
  background: #1c2128;
  color: #939ca6;
}

[data-scheme="dark"] .editor-input {
  color: #dadee3;
}

[data-scheme="dark"] .editor-input:empty::before {
  color: #6e7681;
}

[data-scheme="dark"] .form-footer {
  border-top: 1px solid #30363d;
}

[data-scheme="dark"] .tool-btn {
  color: #939ca6;
}

[data-scheme="dark"] .tool-btn:hover {
  color: #58a6ff;
  background: rgba(88, 166, 255, 0.1);
}

[data-scheme="dark"] .char-count {
  color: #939ca6;
}

[data-scheme="dark"] .login-btn {
  color: #b7c0ca;
  border-color: #30363d;
}

[data-scheme="dark"] .login-btn:hover {
  background: #2d333b;
}

[data-scheme="dark"] .logout-btn {
  color: #b7c0ca;
  border-color: #30363d;
}

[data-scheme="dark"] .logout-btn:hover {
  background: #2d1f1f;
  color: #f85149;
  border-color: #67211d;
}

[data-scheme="dark"] .user-info {
  border-right: 1px solid #30363d;
}

[data-scheme="dark"] .user-name {
  color: #dadee3;
}

[data-scheme="dark"] .submit-btn {
  background: #238636;
}

[data-scheme="dark"] .submit-btn:hover:not(:disabled) {
  background: #2ea043;
}

[data-scheme="dark"] .emoji-panel {
  background: #22272e;
  border: 1px solid #30363d;
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.4);
}

[data-scheme="dark"] .comment-count {
  color: #dadee3;
}

[data-scheme="dark"] .sort-btn {
  color: #939ca6;
}

[data-scheme="dark"] .sort-btn:hover {
  color: #58a6ff;
}

[data-scheme="dark"] .sort-btn.active {
  color: #58a6ff;
}

[data-scheme="dark"] .refreshing-tip {
  color: #58a6ff;
}

[data-scheme="dark"] .loading,
[data-scheme="dark"] .empty {
  color: #939ca6;
}

[data-scheme="dark"] .comment-item {
  border-bottom: 1px solid #30363d;
}

[data-scheme="dark"] .nickname {
  color: #dadee3;
}

[data-scheme="dark"] .time {
  color: #939ca6;
}

[data-scheme="dark"] .comment-content {
  color: #dadee3;
}

[data-scheme="dark"] .reply-to {
  color: #58a6ff;
}

[data-scheme="dark"] .action-btn {
  color: #939ca6;
}

[data-scheme="dark"] .action-btn:hover {
  color: #58a6ff;
}

[data-scheme="dark"] .like-btn:hover {
  color: #f85149;
}

[data-scheme="dark"] .like-btn.liked {
  color: #f85149;
}

[data-scheme="dark"] .reply-form textarea {
  border: 1px solid #30363d;
  background: #22272e;
  color: #dadee3;
}

[data-scheme="dark"] .reply-form textarea:focus {
  border-color: #238636;
}

[data-scheme="dark"] .reply-form textarea::placeholder {
  color: #6e7681;
}

[data-scheme="dark"] .reply-form .submit-btn {
  background: #238636;
}

[data-scheme="dark"] .reply-form .cancel-btn {
  color: #b7c0ca;
  border-color: #30363d;
}

[data-scheme="dark"] .reply-form .cancel-btn:hover {
  background: #2d333b;
}

[data-scheme="dark"] .reply-list {
  border-left: 2px solid #30363d;
}

[data-scheme="dark"] .reply-item:not(:last-child) {
  border-bottom: 1px solid #30363d;
}

[data-scheme="dark"] .ua-text {
  background: #2d333b;
  color: #adbac7;
  border-color: #444c56;
  box-shadow: inset 0 1px 2px rgba(0, 0, 0, 0.3);
}

[data-scheme="dark"] .ua-text:hover {
  background: #373e47;
  border-color: #545d68;
}

/* ===== 登录弹窗暗色模式 ===== */
[data-scheme="dark"] .login-modal {
  background: #22272e;
  box-shadow: 0 20px 60px rgba(0, 0, 0, 0.5);
}

[data-scheme="dark"] .modal-header {
  border-bottom: 1px solid #30363d;
}

[data-scheme="dark"] .modal-header h3 {
  color: #dadee3;
}

[data-scheme="dark"] .close-btn {
  color: #939ca6;
}

[data-scheme="dark"] .close-btn:hover {
  color: #dadee3;
}

[data-scheme="dark"] .login-tip {
  background: #0f2a1e;
  color: #3fb950;
  border: 1px solid #23633e;
}

[data-scheme="dark"] .login-field label {
  color: #b7c0ca;
}

[data-scheme="dark"] .login-field input {
  border: 1px solid #30363d;
  background: #1c2128;
  color: #dadee3;
}

[data-scheme="dark"] .login-field input:focus {
  border-color: #238636;
}

[data-scheme="dark"] .login-error {
  background: #2d1f1f;
  color: #f85149;
  border: 1px solid #67211d;
}

[data-scheme="dark"] .modal-footer {
  border-top: 1px solid #30363d;
  background: #1c2128;
}

[data-scheme="dark"] .modal-footer .cancel-btn {
  background: #22272e;
  color: #b7c0ca;
  border-color: #30363d;
}

[data-scheme="dark"] .modal-footer .cancel-btn:hover {
  background: #2d333b;
}

[data-scheme="dark"] .modal-footer .submit-btn {
  background: #238636;
}

[data-scheme="dark"] .modal-footer .submit-btn:hover:not(:disabled) {
  background: #2ea043;
}

/* ===== 移动端响应式暗色模式 ===== */
@media (max-width: 768px) {
  [data-scheme="dark"] .form-field:first-child {
    border-bottom: 1px dashed #383838;
  }
}
</style>
