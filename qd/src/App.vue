<script setup lang="ts">
import { onMounted } from 'vue'
import LoginModal from '@/components/LoginModal.vue'

onMounted(async () => {
  try {
    const apiUrl = (window as any).APP_CONFIG?.API_BASE_URL || ''
    const res = await fetch(`${apiUrl}/api/public/settings`)
    if (res.ok) {
      const json = await res.json()
      if (json.code === 200 && json.data) {
        localStorage.setItem('site_settings', JSON.stringify(json.data))
        // 触发自定义事件以便其他组件感知
        window.dispatchEvent(new StorageEvent('storage', {
          key: 'site_settings',
          newValue: JSON.stringify(json.data)
        }))
      }
    }
  } catch (error) {
    console.error('Failed to load public settings:', error)
  }
})
</script>

<template>
  <router-view />
  <LoginModal />
</template>

<style>
/* 去除默认边距 */
html, body {
  margin: 0 !important;
  padding: 0 !important;
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 'Helvetica Neue', Arial, sans-serif;
  overflow: hidden !important;
  width: 100vw !important;
  height: 100vh !important;
  background-color: #f9fafb; /* 兜底：任何情况下底部露出的是页面灰而不是 body 默认白 */
}

/* dvh 行单独用 html body 选择器： specificity 更高可覆盖上面的 100vh；
   旧内核（微信X5/旧WebView）不支持 dvh 会丢弃本行、回退 100vh，高度链不断裂。
   不要合并进上面的规则——构建时重复声明会被压缩工具删掉前面的兜底行 */
html body {
  height: 100dvh !important;
}
</style>
