<template>
  <div class="login-container">
    <!-- 动态背景光球 -->
    <div class="bg-orb orb-1"></div>
    <div class="bg-orb orb-2"></div>
    <div class="bg-orb orb-3"></div>
    
    <div class="login-card">
      <div class="login-header">
        <div class="logo-wrapper" v-if="siteLogo">
          <img :src="siteLogo" alt="logo" />
        </div>
        <h2>{{ siteName }}后台管理</h2>
        <p class="subtitle">欢迎回来，请登录您的账户</p>
      </div>

      <el-form :model="form" class="login-form" autocomplete="off" @submit.prevent>
        <el-form-item>
          <el-input 
            v-model="form.username" 
            placeholder="用户名" 
            :prefix-icon="User"
            autocomplete="off"
            size="large"
          />
        </el-form-item>
        
        <el-form-item>
          <el-input 
            v-model="form.password" 
            type="password" 
            placeholder="密码" 
            :prefix-icon="Lock"
            show-password 
            @keyup.enter="handleLogin"
            autocomplete="new-password"
            size="large"
          />
        </el-form-item>

        <div class="form-options">
          <el-checkbox v-model="form.remember" @change="handleRememberChange">记住密码</el-checkbox>
          <a href="javascript:void(0)" class="forgot-link">忘记密码？</a>
        </div>

        <el-button class="submit-btn" :loading="loading" @click="handleLogin">
          登 录
        </el-button>
      </el-form>
    </div>

    <!-- 版权信息 -->
    <div class="copyright">
      徐州邻兔跃IT提供技术支持 版权所有
    </div>
  </div>
</template>

<script setup lang="ts">
import { reactive, ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { User, Lock } from '@element-plus/icons-vue'
import { ElMessage } from 'element-plus'
import defaultLogo from '../assets/logo/logo.png'

const router = useRouter()
const loading = ref(false)
const siteLogo = ref(localStorage.getItem('admin_site_logo') || defaultLogo)
const siteName = ref(localStorage.getItem('admin_site_name') || 'Miren Al')
const form = reactive({
  username: '',
  password: '',
  remember: false
})

const fetchSettings = async () => {
  try {
    const res = await fetch(`${(window as any).APP_CONFIG?.API_BASE_URL || ''}/api/public/settings`)
    const json = await res.json()
    if (res.ok && json.code === 200 && json.data) {
      if (json.data.site_logo !== undefined) {
        if (json.data.site_logo) {
          siteLogo.value = json.data.site_logo
          localStorage.setItem('admin_site_logo', json.data.site_logo)
        } else {
          siteLogo.value = defaultLogo
          localStorage.removeItem('admin_site_logo')
        }
        window.dispatchEvent(new StorageEvent('storage', { key: 'admin_site_logo', newValue: siteLogo.value }))
      }
      if (json.data.site_name !== undefined) {
        if (json.data.site_name) {
          siteName.value = json.data.site_name
          localStorage.setItem('admin_site_name', json.data.site_name)
        } else {
          siteName.value = 'Miren Al'
          localStorage.removeItem('admin_site_name')
        }
        window.dispatchEvent(new StorageEvent('storage', { key: 'admin_site_name', newValue: siteName.value }))
      }
    }
  } catch (error) {
    console.error('Failed to fetch public settings', error)
  }
}

onMounted(() => {
  fetchSettings()
  const savedUsername = localStorage.getItem('admin_remembered_username')
  const savedPassword = localStorage.getItem('admin_remembered_password')
  if (savedUsername && savedPassword) {
    form.username = savedUsername
    form.password = savedPassword
    form.remember = true
  } else {
    form.username = ''
    form.password = ''
    form.remember = false
  }
})

const handleRememberChange = (val: boolean) => {
  if (!val) {
    localStorage.removeItem('admin_remembered_username')
    localStorage.removeItem('admin_remembered_password')
  }
}

const handleLogin = async () => {
  if (!form.username || !form.password) {
    ElMessage.warning('请输入用户名和密码')
    return
  }
  loading.value = true
  try {
    const res = await fetch(`${(window as any).APP_CONFIG?.API_BASE_URL || ''}/api/admin/login`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        username: form.username,
        password: form.password
      })
    })
    const data = await res.json()
    if (res.ok) {
      ElMessage.success('登录成功')
      localStorage.setItem('token', data.token)
      
      if (form.remember) {
        localStorage.setItem('admin_remembered_username', form.username)
        localStorage.setItem('admin_remembered_password', form.password)
      } else {
        localStorage.removeItem('admin_remembered_username')
        localStorage.removeItem('admin_remembered_password')
      }
      
      router.push('/dashboard')
    } else {
      ElMessage.error(data.error || '登录失败')
    }
  } catch (error) {
    ElMessage.error('网络错误，请稍后重试')
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
.login-container {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: #f0f2f5;
  position: relative;
  overflow: hidden;
  font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, sans-serif;
}

/* 动态背景光球 */
.bg-orb {
  position: absolute;
  border-radius: 50%;
  filter: blur(80px);
  z-index: 0;
  animation: float 15s infinite ease-in-out;
  opacity: 0.6;
}

.orb-1 {
  width: 500px;
  height: 500px;
  background: #409eff;
  top: -150px;
  left: -150px;
  animation-delay: 0s;
}

.orb-2 {
  width: 400px;
  height: 400px;
  background: #67c23a;
  bottom: -100px;
  right: -100px;
  animation-delay: -5s;
}

.orb-3 {
  width: 450px;
  height: 450px;
  background: #b37feb;
  top: 30%;
  left: 50%;
  transform: translate(-50%, -50%);
  animation-delay: -10s;
}

@keyframes float {
  0%, 100% { transform: translate(0, 0) scale(1); }
  33% { transform: translate(30px, -50px) scale(1.1); }
  66% { transform: translate(-20px, 20px) scale(0.9); }
}

/* 登录卡片 */
.login-card {
  width: 420px;
  background: rgba(255, 255, 255, 0.85);
  backdrop-filter: blur(20px);
  -webkit-backdrop-filter: blur(20px);
  border: 1px solid rgba(255, 255, 255, 0.6);
  border-radius: 24px;
  padding: 50px 40px;
  box-shadow: 0 10px 40px rgba(0, 0, 0, 0.08), 0 1px 3px rgba(0, 0, 0, 0.05);
  z-index: 1;
  position: relative;
  transition: all 0.3s ease;
}

.login-card:hover {
  box-shadow: 0 15px 50px rgba(0, 0, 0, 0.12), 0 2px 6px rgba(0, 0, 0, 0.04);
}

.login-header {
  text-align: center;
  margin-bottom: 40px;
}

.logo-wrapper {
  width: 64px;
  height: 64px;
  margin: 0 auto 20px;
  border-radius: 16px;
  background: #fff;
  box-shadow: 0 4px 12px rgba(0,0,0,0.05);
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 8px;
}

.logo-wrapper img {
  width: 100%;
  height: 100%;
  object-fit: contain;
}

.login-header h2 {
  font-size: 28px;
  color: #1d1d1f;
  margin: 0 0 8px;
  font-weight: 600;
  letter-spacing: 0.5px;
}

.login-header .subtitle {
  font-size: 14px;
  color: #86868b;
  margin: 0;
}

.login-form {
  margin-top: 10px;
}

:deep(.el-input__wrapper) {
  background-color: #f5f5f7 !important;
  box-shadow: none !important;
  border-radius: 12px;
  padding: 6px 16px;
  transition: all 0.3s ease;
}

:deep(.el-input__wrapper.is-focus) {
  background-color: #fff !important;
  box-shadow: 0 0 0 2px #409eff, 0 4px 12px rgba(64, 158, 255, 0.2) !important;
}

:deep(.el-input__inner) {
  height: 40px;
  font-size: 15px;
  color: #1d1d1f;
}

:deep(.el-input__prefix) {
  font-size: 18px;
  color: #86868b;
  margin-right: 8px;
}

.form-options {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin: 10px 0 30px;
}

:deep(.el-checkbox__input.is-checked .el-checkbox__inner) {
  background-color: #409eff;
  border-color: #409eff;
}
:deep(.el-checkbox__input.is-checked + .el-checkbox__label) {
  color: #409eff;
}

.forgot-link {
  font-size: 14px;
  color: #86868b;
  text-decoration: none;
  transition: color 0.3s;
}

.forgot-link:hover {
  color: #409eff;
}

.submit-btn {
  width: 100%;
  height: 48px;
  border-radius: 12px;
  font-size: 16px;
  font-weight: 600;
  background: linear-gradient(135deg, #409eff 0%, #2979ff 100%);
  border: none;
  color: #fff;
  transition: all 0.3s ease;
  box-shadow: 0 4px 12px rgba(64, 158, 255, 0.3);
}

.submit-btn:hover {
  background: linear-gradient(135deg, #66b1ff 0%, #409eff 100%);
  transform: translateY(-1px);
  box-shadow: 0 6px 16px rgba(64, 158, 255, 0.4);
}

.submit-btn:active {
  transform: translateY(1px);
  box-shadow: 0 2px 8px rgba(64, 158, 255, 0.3);
}

:deep(.el-form-item) {
  margin-bottom: 24px;
}

.copyright {
  position: absolute;
  bottom: 24px;
  left: 0;
  width: 100%;
  text-align: center;
  font-size: 14px;
  color: #86868b;
  letter-spacing: 1px;
  z-index: 10;
}
/* 响应式设计（手机端优化） */
@media screen and (max-width: 480px) {
  .login-card {
    width: 90%;
    padding: 40px 24px;
    border-radius: 20px;
    margin: 0 auto;
    box-sizing: border-box;
  }

  .login-header h2 {
    font-size: 24px;
  }

  .bg-orb {
    /* 手机端保留光球，但缩小尺寸和模糊度，避免影响性能并保持美观 */
    filter: blur(40px);
    opacity: 0.4;
  }
  
  .orb-1 {
    width: 250px;
    height: 250px;
    top: -50px;
    left: -50px;
  }

  .orb-2 {
    width: 200px;
    height: 200px;
    bottom: -50px;
    right: -50px;
  }

  .orb-3 {
    width: 220px;
    height: 220px;
  }

  .copyright {
    font-size: 12px;
    bottom: 16px;
  }
}
</style>