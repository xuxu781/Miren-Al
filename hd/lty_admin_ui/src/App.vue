<script setup lang="ts">
import { onMounted } from 'vue'

onMounted(() => {
  // 动态更新页面标题和Favicon图标
  const updateSiteMeta = () => {
    const siteName = localStorage.getItem('admin_site_name')
    const siteLogo = localStorage.getItem('admin_site_logo')

    if (siteName) {
      document.title = `${siteName}后台管理`
    }
    
    if (siteLogo) {
      let link: HTMLLinkElement | null = document.querySelector("link[rel*='icon']")
      if (!link) {
        link = document.createElement('link')
        link.rel = 'icon'
        document.getElementsByTagName('head')[0].appendChild(link)
      }
      link.href = siteLogo
    }
  }

  // 初始调用
  updateSiteMeta()

  // 监听其他组件（如Settings.vue）通过 localStorage 派发的 storage 事件
  window.addEventListener('storage', (e) => {
    if (e.key === 'admin_site_name' || e.key === 'admin_site_logo') {
      updateSiteMeta()
    }
  })
})
</script>

<template>
  <router-view />
</template>

<style>
/* 去除默认边距 */
body {
  margin: 0;
  padding: 0;
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 'Helvetica Neue', Arial, sans-serif;
}
</style>
