<template>
  <div class="pro-layout" :class="{ 'sidebar-collapsed': isCollapse }">
    <!-- 侧边栏 (玻璃拟态 + 悬浮感) -->
    <aside class="pro-sidebar">
      <div class="pro-logo">
        <img v-if="siteLogo" :src="siteLogo" class="logo-img" alt="logo" />
        <transition name="fade">
          <h1 v-show="!isCollapse" class="logo-text">{{ siteName }}</h1>
        </transition>
      </div>
      <el-scrollbar>
        <el-menu
          :default-active="$route.path"
          class="pro-menu"
          :collapse="isCollapse"
          :collapse-transition="false"
          router
        >
          <el-menu-item index="/dashboard">
            <el-icon><DataLine /></el-icon>
            <template #title>数据概览</template>
          </el-menu-item>
          
          <el-sub-menu index="/user-operations">
            <template #title>
              <el-icon><User /></el-icon>
              <span>用户运营</span>
            </template>
            <el-menu-item index="/users">用户管理</el-menu-item>
            <el-menu-item index="/tasks">任务列表</el-menu-item>
          </el-sub-menu>

          <el-menu-item index="/admin">
            <el-icon><Avatar /></el-icon>
            <template #title>管理员管理</template>
          </el-menu-item>
          <el-menu-item index="/media">
            <el-icon><Picture /></el-icon>
            <template #title>媒体管理</template>
          </el-menu-item>

          <el-sub-menu index="/content">
            <template #title>
              <el-icon><Document /></el-icon>
              <span>内容管理</span>
            </template>
            <el-menu-item index="/content/inspirations">探索管理</el-menu-item>
            <el-menu-item index="/content/inspiration-categories">探索分类设置</el-menu-item>
          </el-sub-menu>

          <el-sub-menu index="/business">
            <template #title>
              <el-icon><Goods /></el-icon>
              <span>商业中心</span>
            </template>
            <el-menu-item index="/business/cdkeys">卡密管理</el-menu-item>
          </el-sub-menu>
          
          <el-sub-menu index="/system">
            <template #title>
              <el-icon><Setting /></el-icon>
              <span>系统配置</span>
            </template>
            <el-menu-item index="/settings">系统设置</el-menu-item>
            <el-menu-item index="/models">模型上游</el-menu-item>
          </el-sub-menu>
        </el-menu>
      </el-scrollbar>
    </aside>

    <!-- 右侧主区域 -->
    <div class="pro-main">
      <!-- 顶部导航栏 -->
      <header class="pro-header">
        <div class="header-left">
          <div class="collapse-trigger" @click="toggleCollapse">
            <el-icon :size="20"><Fold v-if="!isCollapse"/><Expand v-else/></el-icon>
          </div>
          <el-breadcrumb separator="/" class="pro-breadcrumb">
            <el-breadcrumb-item :to="{ path: '/' }">首页</el-breadcrumb-item>
            <el-breadcrumb-item>{{ currentRouteName }}</el-breadcrumb-item>
          </el-breadcrumb>
        </div>

        <div class="header-right">
          <div class="action-item">
            <el-icon :size="18"><Search /></el-icon>
          </div>
          <div class="action-item">
            <el-badge is-dot class="badge-item">
              <el-icon :size="18"><Bell /></el-icon>
            </el-badge>
          </div>
          <el-dropdown trigger="click" class="user-dropdown">
            <div class="user-info">
              <img :src="siteLogo || defaultAvatar" class="avatar" />
              <span class="name">Admin</span>
              <el-icon class="el-icon--right"><CaretBottom /></el-icon>
            </div>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item>
                  <el-icon><User /></el-icon> 个人中心
                </el-dropdown-item>
                <el-dropdown-item>
                  <el-icon><Setting /></el-icon> 系统设置
                </el-dropdown-item>
                <el-dropdown-item divided @click="handleLogout">
                  <span style="color: #f56c6c; display: flex; align-items: center; gap: 4px;">
                    <el-icon><SwitchButton /></el-icon> 退出登录
                  </span>
                </el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </div>
      </header>

      <!-- 页面内容 -->
      <main class="pro-content">
        <div class="page-wrapper">
          <router-view v-slot="{ Component }">
            <transition name="fade-slide" mode="out-in">
              <component :is="Component" />
            </transition>
          </router-view>
        </div>
      </main>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { 
  DataLine, User, Setting, Picture, Document,
  Search, Bell, SwitchButton, Fold, Expand, Avatar, Goods, CaretBottom
} from '@element-plus/icons-vue'
import defaultLogo from '../assets/logo/logo.png'
import defaultAvatar from '../assets/logo/logo.png'

const router = useRouter()
const route = useRoute()
const isCollapse = ref(false)

const siteLogo = ref(localStorage.getItem('admin_site_logo') || defaultLogo)
const siteName = ref(localStorage.getItem('admin_site_name') || 'Miren Al后台管理')

const handleStorageChange = (e: StorageEvent) => {
  if (e.key === 'admin_site_logo') {
    siteLogo.value = e.newValue || defaultLogo
  }
  if (e.key === 'admin_site_name') {
    siteName.value = e.newValue || 'Miren Al后台管理'
  }
}

onMounted(() => {
  window.addEventListener('storage', handleStorageChange)
})

onUnmounted(() => {
  window.removeEventListener('storage', handleStorageChange)
})

const toggleCollapse = () => {
  isCollapse.value = !isCollapse.value
}

const currentRouteName = computed(() => {
  const map: Record<string, string> = {
    '/dashboard': '数据概览',
    '/admin': '管理员管理',
    '/users': '用户管理',
    '/tasks': '任务列表',
    '/media': '媒体管理',
    '/settings': '系统设置',
    '/models': '模型上游',
    '/business/cdkeys': '卡密管理',
    '/content/inspirations': '探索管理',
    '/content/inspiration-categories': '探索分类设置'
  }
  return map[route.path] || '页面'
})

const handleLogout = () => {
  router.push('/login')
}
</script>

<style scoped>
.pro-layout {
  display: flex;
  height: 100vh;
  width: 100vw;
  background-color: #f0f2f5; /* 灰色背景 */
  overflow: hidden;
  position: relative;
}

/* ================= 侧边栏 (玻璃拟态 + 悬浮感) ================= */
.pro-sidebar {
  width: 260px;
  background: rgba(255, 255, 255, 0.75);
  backdrop-filter: blur(20px);
  -webkit-backdrop-filter: blur(20px);
  display: flex;
  flex-direction: column;
  transition: width 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  box-shadow: 0 10px 40px rgba(0, 0, 0, 0.04), 0 1px 3px rgba(0, 0, 0, 0.02);
  border: 1px solid rgba(255, 255, 255, 0.6);
  z-index: 20;
  margin: 12px 0 12px 12px;
  border-radius: 24px;
  height: calc(100vh - 24px);
  overflow: hidden;
}

.sidebar-collapsed .pro-sidebar {
  width: 80px;
}

.pro-logo {
  height: 72px;
  padding: 0 20px;
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: transparent;
  overflow: hidden;
  white-space: nowrap;
  border-bottom: 1px solid rgba(226, 232, 240, 0.6);
}

.logo-img {
  width: 36px;
  height: 36px;
  flex-shrink: 0;
  border-radius: 10px;
  box-shadow: 0 4px 10px rgba(0,0,0,0.05);
}

.logo-text {
  color: #0f172a;
  font-size: 18px;
  font-weight: 700;
  letter-spacing: -0.5px;
  margin: 0 0 0 12px;
}

:deep(.pro-menu) {
  border-right: none;
  background: transparent;
  padding: 12px 16px; /* 增加左右 padding，让内容向中间靠拢 */
}

:deep(.pro-menu:not(.el-menu--collapse)) {
  width: 100%;
}

:deep(.pro-menu .el-menu-item),
:deep(.pro-menu .el-sub-menu__title) {
  height: 48px;
  line-height: 48px;
  margin: 4px 0;
  border-radius: 12px;
  color: #86868b;
  font-weight: 500;
  transition: all 0.2s ease;
  padding: 0 16px !important; /* 强制覆盖原生 padding，保证左侧对齐 */
}

/* 统一图标和文字的间距 */
:deep(.pro-menu .el-menu-item .el-icon),
:deep(.pro-menu .el-sub-menu__title .el-icon) {
  margin-right: 12px;
  font-size: 18px;
  text-align: center;
  width: 24px;
}

/* 修复子菜单展开箭头 ">" 的位置 */
:deep(.pro-menu .el-sub-menu__icon-arrow) {
  position: absolute !important;
  right: 16px !important;
  top: 50% !important;
  transform: translateY(-50%) !important;
  margin: 0 !important;
  width: auto !important;
  font-size: 12px !important;
  color: #86868b !important;
}

/* 修复折叠状态下，隐藏箭头 */
:deep(.pro-menu.el-menu--collapse .el-sub-menu__icon-arrow) {
  display: none !important;
}

/* ================= 折叠状态完美对齐方案 ================= */
/* 1. 整体容器垂直居中 */
:deep(.pro-menu.el-menu--collapse) {
  padding: 12px 0 !important;
  width: 100% !important;
  display: flex !important;
  flex-direction: column !important;
  align-items: center !important;
}

/* 2. 统一所有的列表项(li)外壳为 48x48，确保所有图标按钮在同一条垂直中线上 */
:deep(.pro-menu.el-menu--collapse .el-menu-item),
:deep(.pro-menu.el-menu--collapse .el-sub-menu) {
  width: 48px !important;
  height: 48px !important;
  padding: 0 !important;
  margin: 4px 0 !important;
  display: flex !important;
  justify-content: center !important;
  align-items: center !important;
  border-radius: 12px !important;
}

/* 3. 内部的点击区域完全撑满外壳 */
:deep(.pro-menu.el-menu--collapse .el-menu-tooltip__trigger),
:deep(.pro-menu.el-menu--collapse .el-sub-menu__title) {
  width: 100% !important;
  height: 100% !important;
  padding: 0 !important;
  margin: 0 !important;
  display: flex !important;
  justify-content: center !important;
  align-items: center !important;
  border-radius: 12px !important;
}

/* 4. 彻底隐藏文字内容，防止其隐形占据 flex 空间导致图标向左/向右偏离中心 */
:deep(.pro-menu.el-menu--collapse .el-menu-item span),
:deep(.pro-menu.el-menu--collapse .el-sub-menu__title span) {
  display: none !important;
}

/* 5. 统一图标尺寸，确保正中心对齐 */
:deep(.pro-menu.el-menu--collapse .el-icon) {
  margin: 0 !important;
  width: 20px !important;
  height: 20px !important;
  font-size: 20px !important;
}

/* 6. 高亮背景：只加在外部容器，内部强制透明，防止双重背景颜色变深 */
:deep(.pro-menu.el-menu--collapse .el-menu-item.is-active),
:deep(.pro-menu.el-menu--collapse .el-sub-menu.is-active) {
  background: linear-gradient(135deg, rgba(64, 158, 255, 0.15) 0%, rgba(41, 121, 255, 0.05) 100%) !important;
}
:deep(.pro-menu.el-menu--collapse .el-menu-item.is-active .el-menu-tooltip__trigger),
:deep(.pro-menu.el-menu--collapse .el-sub-menu.is-active .el-sub-menu__title) {
  background: transparent !important;
}

:deep(.pro-menu .el-sub-menu__title:hover),
:deep(.pro-menu .el-menu-item:not(.is-active):hover) {
  background-color: rgba(255, 255, 255, 0.5) !important;
  color: #1d1d1f !important;
}

:deep(.pro-menu .el-menu-item.is-active) {
  background: linear-gradient(135deg, rgba(64, 158, 255, 0.15) 0%, rgba(41, 121, 255, 0.05) 100%) !important;
  color: #409eff !important;
  font-weight: 600;
}

/* 修复子菜单展开时的背景透明问题，为其添加淡淡的背景色以便在侧边栏展开时看得更清楚 */
:deep(.pro-menu .el-menu) {
  background-color: rgba(64, 158, 255, 0.04) !important;
  border-radius: 12px;
  margin: 4px 12px;
  padding: 6px;
  border: 1px solid rgba(64, 158, 255, 0.1);
}

:deep(.pro-menu .el-menu .el-menu-item) {
  border-radius: 8px;
  margin: 2px 0;
  height: 40px;
  line-height: 40px;
}

/* 修复父菜单处于 active 状态时图标和文字的高亮 */
:deep(.pro-menu .el-sub-menu.is-active > .el-sub-menu__title) {
  color: #409eff !important;
  font-weight: 600;
}

/* 修复侧边栏折叠时的弹出菜单背景 */
:deep(.el-menu--popup) {
  background: rgba(255, 255, 255, 0.85) !important;
  backdrop-filter: blur(25px) saturate(180%);
  -webkit-backdrop-filter: blur(25px) saturate(180%);
  border-radius: 16px !important;
  border: 1px solid rgba(255, 255, 255, 0.5) !important;
  box-shadow: 0 12px 32px rgba(0, 0, 0, 0.1) !important;
  padding: 8px !important;
}
:deep(.el-menu--popup .el-menu-item) {
  height: 40px !important;
  line-height: 40px !important;
  border-radius: 8px !important;
  margin: 4px 0 !important;
  color: #434344 !important;
}
:deep(.el-menu--popup .el-menu-item:hover) {
  background-color: rgba(64, 158, 255, 0.08) !important;
  color: #409eff !important;
}
:deep(.el-menu--popup .el-menu-item.is-active) {
  background: linear-gradient(135deg, rgba(64, 158, 255, 0.15) 0%, rgba(41, 121, 255, 0.05) 100%) !important;
  color: #409eff !important;
  font-weight: 600 !important;
}

/* ================= 主区域 ================= */
.pro-main {
  flex: 1;
  display: flex;
  flex-direction: column;
  overflow: hidden;
  position: relative;
  z-index: 10; /* 保证在光球上方 */
}

/* 毛玻璃顶部导航 */
.pro-header {
  height: 72px;
  background: transparent;
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0 32px;
  z-index: 9;
  position: sticky;
  top: 0;
}

.header-left {
  display: flex;
  align-items: center;
  height: 100%;
  gap: 16px;
}

.collapse-trigger {
  width: 40px;
  height: 40px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  font-size: 20px;
  border-radius: 12px;
  background: rgba(255, 255, 255, 0.7);
  backdrop-filter: blur(10px);
  border: 1px solid rgba(255, 255, 255, 0.6);
  color: #434344;
  box-shadow: 0 4px 12px rgba(0,0,0,0.02);
  transition: all 0.3s;
}

.collapse-trigger:hover {
  box-shadow: 0 6px 16px rgba(0,0,0,0.05);
  color: #1d1d1f;
  background: rgba(255, 255, 255, 0.9);
}

.pro-breadcrumb {
  margin-left: 8px;
}

:deep(.pro-breadcrumb .el-breadcrumb__inner) {
  font-weight: 500;
  color: #86868b;
}

:deep(.pro-breadcrumb .el-breadcrumb__item:last-child .el-breadcrumb__inner) {
  color: #1d1d1f;
  font-weight: 600;
}

.header-right {
  display: flex;
  align-items: center;
  height: 100%;
  gap: 16px;
}

.action-item {
  width: 40px;
  height: 40px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  border-radius: 50%;
  transition: all 0.3s;
  color: #434344;
  background: rgba(255, 255, 255, 0.7);
  backdrop-filter: blur(10px);
  border: 1px solid rgba(255, 255, 255, 0.6);
  box-shadow: 0 4px 12px rgba(0,0,0,0.02);
}

.action-item:hover {
  color: #409eff;
  background: rgba(255, 255, 255, 0.9);
  box-shadow: 0 6px 16px rgba(0,0,0,0.05);
  transform: translateY(-1px);
}

.badge-item :deep(.el-badge__content.is-fixed.is-dot) {
  right: 4px;
  top: 6px;
  background-color: #ff3b30;
}

.user-dropdown {
  height: 40px;
  display: flex;
  align-items: center;
  padding: 0 16px 0 8px;
  cursor: pointer;
  border-radius: 20px;
  background: rgba(255, 255, 255, 0.7);
  backdrop-filter: blur(10px);
  border: 1px solid rgba(255, 255, 255, 0.6);
  box-shadow: 0 4px 12px rgba(0,0,0,0.02);
  transition: all 0.3s;
  outline: none !important;
}

.user-dropdown:hover {
  background: rgba(255, 255, 255, 0.9);
  box-shadow: 0 6px 16px rgba(0,0,0,0.05);
}

.user-info {
  display: flex;
  align-items: center;
  gap: 10px;
}

.avatar {
  width: 32px;
  height: 32px;
  border-radius: 50%;
  border: 2px solid rgba(255, 255, 255, 0.8);
  box-shadow: 0 2px 8px rgba(0,0,0,0.1);
}

.name {
  font-size: 14px;
  color: #1d1d1f;
  font-weight: 600;
}

/* ================= 内容区 ================= */
.pro-content {
  flex: 1;
  overflow-y: auto;
  overflow-x: hidden;
  padding: 0 32px 32px 32px;
}

.page-wrapper {
  max-width: 1600px;
  margin: 0 auto;
  min-height: calc(100vh - 104px);
  position: relative;
}

/* 顺滑路由动画 */
.fade-slide-enter-active,
.fade-slide-leave-active {
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1);
}
.fade-slide-enter-from {
  opacity: 0;
  transform: translateY(10px) scale(0.99);
}
.fade-slide-leave-to {
  opacity: 0;
  transform: translateY(-10px) scale(0.99);
}
</style>
