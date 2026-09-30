<template>
  <div class="dashboard-container">
    <!-- 数据概览 -->
    <el-row :gutter="24" class="mb-6">
      <el-col :span="6" v-for="(stat, index) in statistics" :key="index">
        <el-card class="modern-stat-card">
          <div class="stat-content">
            <div class="stat-info">
              <div class="stat-title">{{ stat.title }}</div>
              <div class="stat-value">{{ stat.value }}</div>
            </div>
            <div class="stat-icon" :style="{ background: stat.bgGradient }">
              <el-icon :style="{ color: stat.color }"><component :is="stat.icon" /></el-icon>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <el-row :gutter="24">
      <!-- 快捷操作 -->
      <el-col :span="16">
        <el-card class="modern-box-card">
          <template #header>
            <div class="card-header">
              <span>系统概览</span>
              <el-tag type="success" effect="light" class="modern-tag">运行正常</el-tag>
            </div>
          </template>
          <div class="welcome-section">
            <div class="welcome-text">
              <h3>欢迎使用 Miren Al 现代后台系统</h3>
              <p>这是一个基于 Vue 3 + Element Plus 构建的顶级 SaaS 级后台管理模板，提供了开箱即用的功能，以及极致的交互体验。</p>
              <div class="action-buttons">
                <el-button type="primary">系统设置</el-button>
                <el-button plain class="modern-plain-btn">查看文档</el-button>
              </div>
            </div>
            <div class="welcome-img-wrapper">
              <img src="../assets/vue.svg" class="welcome-img" alt="vue" />
            </div>
          </div>
        </el-card>
      </el-col>

      <!-- 最近动态 -->
      <el-col :span="8">
        <el-card class="modern-box-card">
          <template #header>
            <div class="card-header">
              <span>最近动态</span>
              <el-button link type="primary" class="more-btn">查看更多</el-button>
            </div>
          </template>
          <el-timeline class="modern-timeline">
            <el-timeline-item
              v-for="(activity, index) in activities"
              :key="index"
              :type="activity.type"
              :color="activity.color"
              :size="activity.size"
              :timestamp="activity.timestamp"
            >
              {{ activity.content }}
            </el-timeline-item>
          </el-timeline>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script setup lang="ts">
import { User, DataLine, ShoppingCart, ChatDotRound } from '@element-plus/icons-vue'
import { markRaw } from 'vue'

const statistics = [
  {
    title: '总访问量',
    value: '1,234,567',
    icon: markRaw(DataLine),
    color: '#409eff',
    bgGradient: 'linear-gradient(135deg, rgba(64, 158, 255, 0.15) 0%, rgba(41, 121, 255, 0.05) 100%)'
  },
  {
    title: '新增用户',
    value: '8,432',
    icon: markRaw(User),
    color: '#3b82f6',
    bgGradient: 'linear-gradient(135deg, rgba(59, 130, 246, 0.15) 0%, rgba(37, 99, 235, 0.05) 100%)'
  },
  {
    title: '本月订单',
    value: '6,231',
    icon: markRaw(ShoppingCart),
    color: '#ec4899',
    bgGradient: 'linear-gradient(135deg, rgba(236, 72, 153, 0.15) 0%, rgba(219, 39, 119, 0.05) 100%)'
  },
  {
    title: '未读消息',
    value: '95',
    icon: markRaw(ChatDotRound),
    color: '#10b981',
    bgGradient: 'linear-gradient(135deg, rgba(16, 185, 129, 0.15) 0%, rgba(5, 150, 105, 0.05) 100%)'
  }
]

const activities = [
  {
    content: '管理员 admin 登录系统',
    timestamp: '2026-08-17 10:30',
    type: 'primary',
    size: 'large',
    color: '#409eff'
  },
  {
    content: '用户管理模块更新完成',
    timestamp: '2026-08-16 15:20',
    color: '#10b981'
  },
  {
    content: '系统初始化配置完成',
    timestamp: '2026-08-15 09:00',
    color: '#3b82f6'
  },
  {
    content: 'Miren Al 项目立项',
    timestamp: '2026-08-10 14:00',
    color: '#86868b'
  }
]
</script>

<style scoped>
.dashboard-container {
  animation: fadeIn 0.5s cubic-bezier(0.4, 0, 0.2, 1);
}

@keyframes fadeIn {
  from { opacity: 0; transform: translateY(15px); }
  to { opacity: 1; transform: translateY(0); }
}

.mb-6 {
  margin-bottom: 24px;
}

.modern-stat-card {
  height: 120px;
  display: flex;
  flex-direction: column;
  justify-content: center;
}

.stat-content {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.stat-icon {
  width: 56px;
  height: 56px;
  border-radius: 16px;
  display: flex;
  justify-content: center;
  align-items: center;
  transition: transform 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.modern-stat-card:hover .stat-icon {
  transform: scale(1.1) rotate(5deg);
}

.stat-icon .el-icon {
  font-size: 26px;
}

.stat-info {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.stat-title {
  font-size: 14px;
  color: #64748b;
  font-weight: 500;
}

.stat-value {
  font-size: 32px;
  font-weight: 700;
  color: #0f172a;
  line-height: 1;
  font-family: 'Inter', -apple-system, sans-serif;
  letter-spacing: -1px;
}

.modern-box-card {
  min-height: 420px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  font-weight: 600;
  color: #0f172a;
  font-size: 16px;
}

.modern-tag {
  border-radius: 6px;
  border: none;
  font-weight: 600;
  background: rgba(16, 185, 129, 0.1);
  color: #10b981;
}

.welcome-section {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 24px 12px;
}

.welcome-text h3 {
  margin: 0 0 16px 0;
  font-size: 28px;
  color: #0f172a;
  font-weight: 700;
  letter-spacing: -0.5px;
}

.welcome-text p {
  color: #64748b;
  line-height: 1.8;
  margin-bottom: 32px;
  font-size: 15px;
  max-width: 480px;
}

.welcome-img-wrapper {
  width: 200px;
  height: 200px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: radial-gradient(circle, rgba(99, 102, 241, 0.1) 0%, transparent 70%);
  border-radius: 50%;
}

.welcome-img {
  width: 120px;
  height: 120px;
  animation: float 6s ease-in-out infinite;
  filter: drop-shadow(0 10px 15px rgba(65, 184, 131, 0.3));
}

@keyframes float {
  0% { transform: translateY(0px); }
  50% { transform: translateY(-12px); }
  100% { transform: translateY(0px); }
}

.action-buttons {
  display: flex;
  gap: 16px;
}

.action-buttons .el-button {
  padding: 12px 28px;
}

.modern-plain-btn {
  background: #f8fafc !important;
  border: 1px solid #e2e8f0 !important;
  color: #0f172a !important;
}

.modern-plain-btn:hover {
  background: #f1f5f9 !important;
  border-color: #cbd5e1 !important;
}

.more-btn {
  color: #409eff;
  font-weight: 500;
}

.modern-timeline {
  padding-top: 16px;
  padding-left: 8px;
}

:deep(.el-timeline-item__content) {
  color: #1d1d1f;
  font-size: 14px;
  font-weight: 500;
}

:deep(.el-timeline-item__timestamp) {
  color: #86868b;
  font-size: 13px;
  margin-top: 6px;
}
</style>
