<template>
  <div class="elegant-install-page">
    <!-- 动态背景光球 -->
    <div class="bg-orb orb-1"></div>
    <div class="bg-orb orb-2"></div>
    <div class="bg-orb orb-3"></div>
    
    <div class="install-card">
      <!-- Logo & Header -->
      <div class="card-header">
        <div class="logo-box">
          <img src="../assets/logo/logo.png" alt="Logo" class="logo-img" />
        </div>
        <h1 class="page-title">系统安装向导</h1>
        <p class="page-subtitle">欢迎使用 Miren AI，只需两步即可完成专属配置</p>
      </div>

      <!-- 自定义高级进度条 -->
      <div class="progress-section">
        <div class="progress-text">
          <span class="step-label">{{ activeStep === 0 ? '第一步：配置数据库' : '第二步：设置管理员' }}</span>
          <span class="step-count">{{ activeStep + 1 }} / 2</span>
        </div>
        <div class="progress-bar-bg">
          <div class="progress-bar-fill" :style="{ width: activeStep === 0 ? '50%' : '100%' }"></div>
        </div>
      </div>

      <!-- 表单区域 -->
      <el-form :model="form" :rules="rules" ref="formRef" label-position="top" class="elegant-form" @submit.prevent>
        <transition name="slide-fade" mode="out-in">
          <div v-if="activeStep === 0" key="step0" class="form-step">
            <el-row :gutter="16">
              <el-col :span="16">
                <el-form-item label="数据库地址" prop="db_host">
                  <el-input v-model="form.db_host" placeholder="例如: 127.0.0.1">
                    <template #prefix><el-icon><Platform /></el-icon></template>
                  </el-input>
                </el-form-item>
              </el-col>
              <el-col :span="8">
                <el-form-item label="端口" prop="db_port">
                  <el-input v-model="form.db_port" placeholder="3306" />
                </el-form-item>
              </el-col>
            </el-row>

            <el-form-item label="数据库用户" prop="db_user">
              <el-input v-model="form.db_user" placeholder="例如: root">
                <template #prefix><el-icon><User /></el-icon></template>
              </el-input>
            </el-form-item>

            <el-form-item label="数据库密码" prop="db_pass">
              <el-input v-model="form.db_pass" type="password" show-password placeholder="请输入数据库密码">
                <template #prefix><el-icon><Lock /></el-icon></template>
              </el-input>
            </el-form-item>

            <el-form-item label="数据库名称" prop="db_name">
              <el-input v-model="form.db_name" placeholder="例如: lty_db">
                <template #prefix><el-icon><Coin /></el-icon></template>
              </el-input>
            </el-form-item>

            <el-button type="primary" class="submit-btn" @click="nextStep" :loading="loading">
              下一步
            </el-button>
          </div>

          <div v-else key="step1" class="form-step">
            <el-form-item label="管理员账号" prop="admin_user">
              <el-input v-model="form.admin_user" placeholder="例如: admin">
                <template #prefix><el-icon><UserFilled /></el-icon></template>
              </el-input>
            </el-form-item>

            <el-form-item label="管理员密码" prop="admin_pass">
              <el-input v-model="form.admin_pass" type="password" show-password placeholder="请设置高强度密码">
                <template #prefix><el-icon><Key /></el-icon></template>
              </el-input>
            </el-form-item>

            <div class="btn-row">
              <el-button class="back-btn" @click="prevStep" :disabled="loading">
                返回
              </el-button>
              <el-button type="primary" class="submit-btn" @click="handleInstall" :loading="loading">
                完成安装
              </el-button>
            </div>
          </div>
        </transition>
      </el-form>
    </div>

    <!-- 版权信息 -->
    <div class="copyright">
      徐州邻兔跃IT提供技术支持 版权所有
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'
import { ElMessage } from 'element-plus'
import type { FormInstance, FormRules } from 'element-plus'
import { User, Lock, Platform, Coin, UserFilled, Key } from '@element-plus/icons-vue'
import axios from 'axios'

const formRef = ref<FormInstance>()
const loading = ref(false)
const activeStep = ref(0)

const form = reactive({
  db_host: '127.0.0.1',
  db_port: '3306',
  db_user: 'root',
  db_pass: '123456',
  db_name: 'lty_db',
  admin_user: 'admin',
  admin_pass: '123456'
})

const rules = reactive<FormRules>({
  db_host: [{ required: true, message: '请输入数据库地址', trigger: 'blur' }],
  db_port: [{ required: true, message: '请输入数据库端口', trigger: 'blur' }],
  db_user: [{ required: true, message: '请输入数据库用户', trigger: 'blur' }],
  db_name: [{ required: true, message: '请输入数据库名称', trigger: 'blur' }],
  admin_user: [{ required: true, message: '请输入管理员账号', trigger: 'blur' }],
  admin_pass: [{ required: true, message: '请输入管理员密码', trigger: 'blur' }]
})

const nextStep = async () => {
  if (!formRef.value) return
  await formRef.value.validateField(['db_host', 'db_port', 'db_user', 'db_name'], async (valid) => {
    if (valid) {
      loading.value = true
      try {
        const res = await axios.post(`${(window as any).APP_CONFIG?.API_BASE_URL || ''}/api/check-db`, {
          db_host: form.db_host,
          db_port: form.db_port,
          db_user: form.db_user,
          db_pass: form.db_pass,
          db_name: form.db_name
        })
        if (res.status === 200) {
          ElMessage.success('数据库连接成功')
          activeStep.value = 1
        }
      } catch (error: any) {
        ElMessage.error(error.response?.data?.error || '数据库连接失败，请检查配置')
      } finally {
        loading.value = false
      }
    }
  })
}

const prevStep = () => {
  activeStep.value = 0
}

const handleInstall = async () => {
  if (!formRef.value) return
  await formRef.value.validate(async (valid) => {
    if (valid) {
      loading.value = true
      try {
        const res = await axios.post(`${(window as any).APP_CONFIG?.API_BASE_URL || ''}/api/install`, form)
        if (res.status === 200) {
          ElMessage.success('系统安装成功！')
          window.location.href = '/login'
        }
      } catch (error: any) {
        ElMessage.error(error.response?.data?.error || '安装失败，请检查数据库配置是否正确')
      } finally {
        loading.value = false
      }
    }
  })
}
</script>

<style scoped>
/* 全局美化 */
.elegant-install-page {
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

/* 毛玻璃悬浮卡片 */
.install-card {
  position: relative;
  z-index: 1;
  width: 100%;
  max-width: 480px;
  background: rgba(255, 255, 255, 0.85);
  backdrop-filter: blur(20px);
  -webkit-backdrop-filter: blur(20px);
  border: 1px solid rgba(255, 255, 255, 0.6);
  border-radius: 24px;
  padding: 50px 40px;
  box-shadow: 0 10px 40px rgba(0, 0, 0, 0.08), 0 1px 3px rgba(0, 0, 0, 0.05);
  box-sizing: border-box;
  transition: all 0.3s ease;
}

.install-card:hover {
  box-shadow: 0 15px 50px rgba(0, 0, 0, 0.12), 0 2px 6px rgba(0, 0, 0, 0.04);
}

/* 头部设计 */
.logo-box {
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

.logo-img {
  width: 100%;
  height: 100%;
  object-fit: contain;
}

.page-title {
  text-align: center;
  font-size: 28px;
  color: #1d1d1f;
  margin: 0 0 8px;
  font-weight: 600;
  letter-spacing: 0.5px;
}

.page-subtitle {
  text-align: center;
  font-size: 14px;
  color: #86868b;
  margin: 0 0 36px;
}

/* 自定义进度条 */
.progress-section {
  margin-bottom: 32px;
}

.progress-text {
  display: flex;
  justify-content: space-between;
  font-size: 13px;
  font-weight: 600;
  color: #86868b;
  margin-bottom: 10px;
}

.progress-bar-bg {
  height: 6px;
  background: #f5f5f7;
  border-radius: 12px;
  overflow: hidden;
}

.progress-bar-fill {
  height: 100%;
  background: linear-gradient(135deg, #409eff 0%, #2979ff 100%);
  border-radius: 12px;
  transition: width 0.5s cubic-bezier(0.4, 0, 0.2, 1);
}

/* 表单精细化美化 */
.elegant-form :deep(.el-form-item) {
  margin-bottom: 24px;
}

.elegant-form :deep(.el-form-item__label) {
  font-weight: 600;
  color: #1d1d1f;
  padding-bottom: 8px;
  line-height: 1;
  font-size: 14px;
}

.elegant-form :deep(.el-input__wrapper) {
  background-color: #f5f5f7 !important;
  box-shadow: none !important;
  border-radius: 12px;
  padding: 6px 16px;
  transition: all 0.3s ease;
}

.elegant-form :deep(.el-input__wrapper.is-focus) {
  background-color: #fff !important;
  box-shadow: 0 0 0 2px #409eff, 0 4px 12px rgba(64, 158, 255, 0.2) !important;
}

.elegant-form :deep(.el-input__inner) {
  height: 28px;
  font-size: 15px;
  color: #1d1d1f;
}

.elegant-form :deep(.el-input__prefix) {
  font-size: 18px;
  color: #86868b;
  margin-right: 8px;
}

/* 按钮设计 */
.submit-btn {
  width: 100%;
  height: 48px;
  border-radius: 12px;
  font-size: 16px;
  font-weight: 600;
  background: linear-gradient(135deg, #409eff 0%, #2979ff 100%);
  border: none;
  color: #fff;
  margin-top: 12px;
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

.btn-row {
  display: flex;
  gap: 16px;
  margin-top: 12px;
}

.btn-row .submit-btn {
  margin-top: 0;
  flex: 2;
}

.btn-row .back-btn {
  flex: 1;
  height: 48px;
  border-radius: 12px;
  background: #ffffff;
  color: #86868b;
  border: 1px solid #dcdfe6;
  font-size: 15px;
  font-weight: 600;
  transition: all 0.3s ease;
}

.btn-row .back-btn:hover {
  color: #409eff;
  border-color: #409eff;
  background: #ecf5ff;
}

/* 版权信息 */
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

/* 丝滑动画 */
.slide-fade-enter-active,
.slide-fade-leave-active {
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1);
}

.slide-fade-enter-from {
  opacity: 0;
  transform: translateX(20px);
}

.slide-fade-leave-to {
  opacity: 0;
  transform: translateX(-20px);
}
/* 响应式设计（手机端优化） */
@media screen and (max-width: 480px) {
  .install-card {
    width: 90%;
    padding: 30px 20px;
    border-radius: 20px;
    margin: 0 auto;
    box-sizing: border-box;
  }

  .page-title {
    font-size: 24px;
  }

  .bg-orb {
    /* 手机端保留光球，缩小尺寸和透明度 */
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
    position: relative; /* 手机端避免绝对定位遮挡输入内容 */
    margin-top: 20px;
  }
}
</style>