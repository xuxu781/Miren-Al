<template>
  <main class="explore-main">
    <!-- 现代科技感光晕背景 -->
    <div class="absolute top-0 left-0 right-0 h-[600px] overflow-hidden pointer-events-none z-0 flex justify-center items-start opacity-70">
      <div class="absolute top-[-50px] left-[15%] w-[40vw] max-w-[500px] h-[300px] rounded-full mix-blend-multiply" style="background-color: rgba(99, 102, 241, 0.3); filter: blur(90px);"></div>
      <div class="absolute top-[20px] right-[15%] w-[35vw] max-w-[450px] h-[300px] rounded-full mix-blend-multiply" style="background-color: rgba(168, 85, 247, 0.2); filter: blur(90px);"></div>
      <div class="absolute top-[80px] left-[30%] w-[40vw] max-w-[500px] h-[250px] rounded-full mix-blend-multiply" style="background-color: rgba(56, 189, 248, 0.25); filter: blur(90px);"></div>
    </div>
    
    <header class="main-header">
      <div class="header-left">
        <div class="mobile-menu-btn" @click="$emit('toggle-menu')">
          <el-icon><Expand /></el-icon>
        </div>
        <div class="header-new-chat-icon desktop-only" @click="$emit('new-chat')" title="新建会话">
          <el-icon><Edit /></el-icon>
        </div>
        <div v-if="!isLoggedIn" class="header-login-btn mobile-only" @click="authStore.openLogin()">
          登录
        </div>
        <div v-else class="header-points-btn mobile-only" @click="$emit('points-click')" title="查看积分记录">
          <el-icon><Coin /></el-icon>
          <span>{{ parseFloat(Number(userPoints || 0).toFixed(2)) }}</span>
        </div>
      </div>
      <div class="header-actions">
        <div class="header-new-chat-icon mobile-only" @click="$emit('new-chat')" title="新建会话">
          <el-icon><Edit /></el-icon>
        </div>
        <div v-if="!isLoggedIn" class="header-login-btn desktop-only" @click="authStore.openLogin()">
          登录 / 注册
        </div>
        <div v-else class="header-points-btn desktop-only" @click="$emit('points-click')" title="查看积分记录">
          <el-icon><Coin /></el-icon>
          <span>{{ parseFloat(Number(userPoints || 0).toFixed(2)) }} 积分</span>
        </div>
      </div>
    </header>

    <div class="explore-content-area" @scroll="handleScroll">
      <div class="explore-container">
        
        <!-- Hero Section with Input Box -->
        <div class="relative z-50 w-full px-0 sm:px-6 md:px-8 mb-8 mt-0 trae-browser-inspect-draggable">
          <div class="mx-auto w-full max-w-[1454px]">
            <h1 class="relative z-[60] mb-6 md:mb-12 mt-4 md:mt-6 flex flex-col gap-1.5 md:gap-4 text-center font-sans items-center justify-center group">
              <span class="relative block text-[28px] sm:text-[44px] md:text-[64px] font-black leading-[1.2] tracking-tight px-2 flex flex-wrap justify-center items-center">
                <span class="text-[#0f172a]">探索视觉边界</span>
                <span class="text-[#94a3b8] font-light mx-2 md:mx-4 transform -translate-y-[2px]">·</span>
                <span class="text-[#4d6bfe]">Miren AI</span>
              </span>
            </h1>
            <div class="mx-auto w-full max-w-[1200px]">
              <div class="rounded-2xl transition-all duration-200 ease-in-out min-w-0 max-w-full p-0 flex min-h-0 flex-col">
                <div class="relative w-full flex min-h-0 flex-col">
                  <!-- Main Input Box -->
                  <div class="input-area-wrapper" :class="{ 'is-shrunk': isInputShrunk }" :style="{ '--keyboard-offset': `${keyboardOffset}px` }">
                    <!-- 嵌入与 Generate.vue 相同的输入框和工具栏 -->
                    <div class="input-container">
                      <div class="input-tools-container">
                        <div class="input-tools">
                          <el-dropdown class="model-dropdown-wrapper" trigger="click" placement="top-start" :teleported="true" @command="(val: string) => props.form.series_id = val">
                          <button class="combined-settings-btn model-settings-btn" type="button">
                            <div class="model-btn-content">
                              <el-icon class="cpu-icon"><Cpu /></el-icon>
                              <span class="btn-text model-name-text" :title="props.availableModels.find(m => m.series_id === props.form.series_id)?.name || '默认模型'">
                                {{ props.availableModels.find(m => m.series_id === props.form.series_id)?.name || '默认模型' }}
                              </span>
                              <span v-if="props.availableModels.find(m => m.series_id === props.form.series_id)?.activity_tag" class="shimmer-tag" :style="{ fontSize: '10px', color: '#fff', background: props.availableModels.find(m => m.series_id === props.form.series_id)?.activity_tag_color || '#10b981', padding: '2px 4px', borderRadius: '4px', marginLeft: '6px', whiteSpace: 'nowrap', flexShrink: 0 }">
                                {{ props.availableModels.find(m => m.series_id === props.form.series_id)?.activity_tag }}
                              </span>
                              <span v-else-if="props.availableModels.find(m => m.series_id === props.form.series_id) && props.hasFreeResolution(props.availableModels.find(m => m.series_id === props.form.series_id))" class="shimmer-tag" style="font-size: 10px; color: #fff; background: #10b981; padding: 2px 4px; border-radius: 4px; margin-left: 6px; white-space: nowrap; flex-shrink: 0;">
                                限时免费
                              </span>
                            </div>
                            <el-icon class="arrow-icon"><ArrowDown /></el-icon>
                          </button>
                          <template #dropdown>
                            <el-dropdown-menu class="model-dropdown-menu">
                              <el-dropdown-item 
                                v-for="model in props.availableModels" 
                                :key="model.series_id" 
                                :command="model.series_id"
                                :class="{ 'is-active-model': props.form.series_id === model.series_id }"
                              >
                                <div style="display: flex; align-items: center; justify-content: space-between; width: 100%; gap: 12px;">
                                    <div style="display: flex; align-items: center;">
                                      <el-icon v-if="props.form.series_id === model.series_id" style="color: #3b82f6; font-size: 14px; margin-right: 6px; font-weight: bold;"><Check /></el-icon>
                                      <span v-else style="width: 20px; display: inline-block;"></span>
                                      <span>{{ model.name }}</span>
                                    </div>
                                    <span v-if="model.activity_tag" class="shimmer-tag" :style="{ fontSize: '11px', color: '#fff', background: model.activity_tag_color || '#10b981', padding: '2px 6px', borderRadius: '4px', lineHeight: '1.2', whiteSpace: 'nowrap' }">{{ model.activity_tag }}</span>
                                    <span v-else-if="props.hasFreeResolution(model)" class="shimmer-tag" style="font-size: 11px; color: #fff; background: #10b981; padding: 2px 6px; border-radius: 4px; line-height: 1.2; white-space: nowrap;">限时免费</span>
                                  </div>
                              </el-dropdown-item>
                              <el-dropdown-item 
                                v-if="props.availableModels.length === 0" 
                                command="default"
                                :class="{ 'is-active-model': props.form.series_id === 'default' }"
                              >
                                <div style="display: flex; align-items: center;">
                                  <el-icon v-if="props.form.series_id === 'default'" style="color: #3b82f6; font-size: 14px; margin-right: 6px; font-weight: bold;"><Check /></el-icon>
                                  <span v-else style="width: 20px; display: inline-block;"></span>
                                  <span>默认模型</span>
                                </div>
                              </el-dropdown-item>
                            </el-dropdown-menu>
                          </template>
                        </el-dropdown>
                        
                        <el-popover placement="top-start" :width="props.windowWidth <= 768 ? props.windowWidth - 32 : 360" trigger="click" popper-class="settings-popover" :popper-options="{ modifiers: [{ name: 'preventOverflow', options: { padding: 16 } }] }">
                          <template #reference>
                            <button class="combined-settings-btn" type="button">
                              <el-icon><Crop /></el-icon>
                              <span class="btn-text">
                                <template v-if="props.currentModelRatios.length > 0">
                                  {{ props.form.aspect_ratio || 'auto' }}
                                  <div class="divider"></div>
                                </template>
                                {{ props.form.resolution || '1K' }}
                                <template v-if="props.currentModelImageCounts.length > 0">
                                  <div class="divider"></div>
                                  <span>{{ props.form.num_images || 1 }}<span class="unit-text">张</span></span>
                                </template>
                              </span>
                            </button>
                          </template>
                          <div class="settings-panel">
                            <div class="setting-item" v-if="props.currentModelRatios.length > 0">
                              <div class="setting-label">比例</div>
                              <div class="setting-options">
                                <div class="option-btn" v-for="ratio in props.currentModelRatios" :key="ratio" :class="{ active: props.form.aspect_ratio === ratio }" @click="props.form.aspect_ratio = ratio">
                                  {{ ratio }}
                                </div>
                              </div>
                            </div>
                            <div class="setting-item">
                              <div class="setting-label">分辨率</div>
                              <div class="setting-options">
                                <div class="option-btn" v-for="res in props.currentModelResolutions" :key="res" :class="{ active: props.form.resolution === res }" @click="props.form.resolution = res">
                                  <div style="display: flex; align-items: center; gap: 6px;">
                                    <span>{{ res }}</span>
                                    <span v-if="props.getTierConfigForRes(props.currentModelInfo, res, 'enabled', true) === 'maintenance'" class="shimmer-tag" style="font-size: 11px; color: #fff; background: #9ca3af; padding: 0 4px; border-radius: 4px; line-height: 1.4; white-space: nowrap;">维护中</span>
                                    <span v-else-if="props.getTierConfigForRes(props.currentModelInfo, res, 'tag', '')" class="shimmer-tag" :style="{ fontSize: '11px', color: '#fff', background: props.getTierConfigForRes(props.currentModelInfo, res, 'tag_color', '#10b981'), padding: '0 4px', borderRadius: '4px', lineHeight: '1.4', whiteSpace: 'nowrap' }">{{ props.getTierConfigForRes(props.currentModelInfo, res, 'tag', '') }}</span>
                                    <span v-else-if="Number(props.getTierConfigForRes(props.currentModelInfo, res, 'credits_per_image', 1)) === 0" class="shimmer-tag" style="font-size: 11px; color: #fff; background: #10b981; padding: 0 4px; border-radius: 4px; line-height: 1.4; white-space: nowrap;">限时免费</span>
                                    <span v-else style="font-size: 11px; opacity: 0.7;">{{ parseFloat(Number(props.getTierConfigForRes(props.currentModelInfo, res, 'credits_per_image', 1)).toFixed(2)) }}积分</span>
                                  </div>
                                </div>
                              </div>
                            </div>
                            <div class="setting-item" v-if="props.currentModelImageCounts.length > 0">
                              <div class="setting-label">生成数量</div>
                              <div class="setting-options">
                                <div class="option-btn" v-for="count in props.currentModelImageCounts" :key="count" :class="{ active: props.form.num_images === Number(count) }" @click="props.form.num_images = Number(count)">
                                  {{ count }}张
                                </div>
                              </div>
                            </div>
                          </div>
                        </el-popover>
                      </div>
                    </div>
                      
                    <div class="input-box" :class="{ 'has-refs': props.form.reference_image && props.form.reference_image.length > 0, 'is-focused': isInputFocused }" @click="isInputShrunk ? (isInputFocused = true, isInputShrunk = false) : null">
                        <div class="reference-upload-item initial-upload-item" @click.stop="isInputShrunk ? (isInputFocused = true, isInputShrunk = false) : props.triggerUpload()" v-if="(!props.form.reference_image || props.form.reference_image.length === 0) && props.currentModelMaxRefImages > 0">
                          <div class="reference-upload-content" style="transform: rotate(8deg);">
                            <svg width="1em" height="1em" viewBox="0 0 24 24" preserveAspectRatio="xMidYMid meet" fill="none" role="presentation" xmlns="http://www.w3.org/2000/svg" class="upload-icon-svg"><g><path data-follow-fill="currentColor" d="M10.8 20a1.2 1.2 0 0 0 2.4 0v-6.8H20a1.2 1.2 0 1 0 0-2.4h-6.8V4a1.2 1.2 0 0 0-2.4 0v6.8H4a1.2 1.2 0 0 0 0 2.4h6.8V20Z" clip-rule="evenodd" fill-rule="evenodd" fill="currentColor"></path></g></svg>
                          </div>
                        </div>
                        
                        <div class="reference-image-preview-inline" v-if="props.form.reference_image && props.form.reference_image.length > 0" :style="{ '--total-items': (props.form.reference_image.length + (props.form.reference_image.length < props.currentModelMaxRefImages ? 1 : 0)) } as any">
                          <div class="preview-list-inline" :class="{ 'is-expandable': props.form.reference_image.length > 1 || (props.form.reference_image.length === 1 && props.currentModelMaxRefImages > 1), 'is-mobile-expanded': isMobileRefExpanded }" @click="isInputShrunk ? (isInputFocused = true, isInputShrunk = false) : (props.windowWidth <= 768 ? isMobileRefExpanded = true : null)">
                              <div class="reference-group-hover-trigger"></div>
                              <div class="preview-container-inline" v-for="(img, index) in props.form.reference_image" :key="index" :style="{ '--index': index, '--rotate': `${index === 0 ? 0 : [6, -4, 2, -8, 8, -6, 4, -2, 10, -10, 5, -5, 7, -7, 3, -3][(Number(index) - 1) % 16]}deg`, zIndex: props.form.reference_image.length - Number(index) + 1 } as any">
                                <el-image 
                                  :src="img" 
                                  class="preview-img-inline" 
                                  :preview-src-list="props.windowWidth > 768 ? props.form.reference_image : []" 
                                  :initial-index="index"
                                  fit="cover"
                                  :preview-teleported="true"
                                  :hide-on-click-modal="true"
                                />
                                <div class="remove-btn-inline" @click.stop="props.clearReferenceImage(Number(index))">
                                  <el-icon><Close /></el-icon>
                                </div>
                                <div class="collapsed-upload-badge" @click.stop="props.triggerUpload" v-if="index === 0 && props.form.reference_image.length < props.currentModelMaxRefImages">
                                  <svg width="1em" height="1em" viewBox="0 0 24 24" preserveAspectRatio="xMidYMid meet" fill="none" role="presentation" xmlns="http://www.w3.org/2000/svg"><g><path data-follow-fill="currentColor" d="M10.8 20a1.2 1.2 0 0 0 2.4 0v-6.8H20a1.2 1.2 0 1 0 0-2.4h-6.8V4a1.2 1.2 0 0 0-2.4 0v6.8H4a1.2 1.2 0 0 0 0 2.4h6.8V20Z" clip-rule="evenodd" fill-rule="evenodd" fill="currentColor"></path></g></svg>
                                </div>
                              </div>
                              <div class="reference-upload-item-inline" @click.stop="props.triggerUpload" v-if="props.form.reference_image.length < props.currentModelMaxRefImages" :style="{ '--index': props.form.reference_image.length, '--rotate': `-12deg`, zIndex: 0 } as any">
                                <div class="reference-upload-content" style="transform: rotate(12deg);">
                                  <svg width="1em" height="1em" viewBox="0 0 24 24" preserveAspectRatio="xMidYMid meet" fill="none" role="presentation" xmlns="http://www.w3.org/2000/svg" class="upload-icon-svg"><g><path data-follow-fill="currentColor" d="M10.8 20a1.2 1.2 0 0 0 2.4 0v-6.8H20a1.2 1.2 0 1 0 0-2.4h-6.8V4a1.2 1.2 0 0 0-2.4 0v6.8H4a1.2 1.2 0 0 0 0 2.4h6.8V20Z" clip-rule="evenodd" fill-rule="evenodd" fill="currentColor"></path></g></svg>
                                </div>
                              </div>
                            </div>
                          </div>
                          
                          <input type="file" ref="fileInputRef" accept="image/*" style="display: none" @change="(e) => props.handleFileUpload(e)" multiple />
                          <el-input
                              ref="inputRef"
                              v-model="props.form.prompt"
                              type="textarea"
                              :rows="1"
                              :autosize="{ minRows: 1, maxRows: 6 }"
                              :placeholder="props.windowWidth <= 768 ? '请输入图片描述...' : '请输入你想生成的图片描述...'"
                              resize="none"
                              class="chat-input"
                              :class="{ 'has-references': props.form.reference_image && props.form.reference_image.length > 0, 'no-refs-allowed': props.currentModelMaxRefImages === 0 }"
                              @keydown.enter.exact.prevent="handleSend"
                              @keydown.enter.shift.exact.prevent="handleShiftEnter"
                              @focus="isInputFocused = true; isInputShrunk = false"
                              @blur="handleInputBlur"
                              @click="isInputFocused = true; isInputShrunk = false"
                              @input="handleInput"
                            />
                          <el-button
                            v-if="props.windowWidth <= 768 && isMobileRefExpanded"
                            type="info"
                            circle
                            class="send-btn"
                            @click.stop="isMobileRefExpanded = false"
                          >
                            <el-icon><Close /></el-icon>
                          </el-button>
                          <el-button
                            v-else
                            type="primary"
                            circle
                            class="send-btn"
                            :class="{ 'maintenance-btn': props.isCurrentResolutionMaintenance }"
                            :disabled="!props.form.prompt.trim() || props.isCurrentResolutionMaintenance"
                            @click="handleSend"
                            :title="props.isCurrentResolutionMaintenance ? '该模型正在维护中，暂时无法生成' : '开始生成'"
                          >
                            <el-icon v-if="!props.isCurrentResolutionMaintenance"><Position /></el-icon>
                            <span v-else style="font-size: 12px; transform: scale(0.9);">维护中</span>
                          </el-button>
                        </div>
                      </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- 极简风分类选择区 -->
        <div class="category-filter-section" v-if="categories.length > 0">
          <div class="main-categories-wrapper">
            <div class="main-categories">
              <div 
                class="category-tab" 
                :class="{ 'active': currentMainCategory === '' }"
                @click="selectMainCategory('')"
              >
                推荐
              </div>
              <div 
                v-for="cat in categories" 
                :key="cat.name"
                class="category-tab"
                :class="{ 'active': currentMainCategory === cat.name }"
                @click="selectMainCategory(cat.name)"
              >
                {{ cat.name }}
              </div>
            </div>
          </div>
          
          <!-- 优雅展开的副分类 -->
          <div 
            class="sub-categories-wrapper" 
            :class="{ 'is-expanded': currentMainCategory && currentSubCategories.length > 0 }"
          >
            <div class="sub-categories">
              <div class="sub-categories-inner">
                <div 
                  class="sub-category-tag"
                  :class="{ 'active': currentSubCategory === '' }"
                  @click="selectSubCategory('')"
                >
                  全部
                </div>
                <div 
                  v-for="sub in currentSubCategories" 
                  :key="sub"
                  class="sub-category-tag"
                  :class="{ 'active': currentSubCategory === sub }"
                  @click="selectSubCategory(sub)"
                >
                  {{ sub }}
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- 瀑布流布局区 (Flex Columns Masonry，彻底解决跳动问题) -->
        <div class="explore-masonry-flex" v-if="publicInspirations.length > 0 || loadingInspirations">
          
          <div 
            class="masonry-column" 
            v-for="(col, colIndex) in masonryColumnsData" 
            :key="'col-' + colIndex"
          >
            <!-- 真实数据 -->
            <div 
              v-for="item in col" 
              :key="'img-' + item.id" 
              class="prompt-card"
              @click="handleShowPrompt(item)"
            >
              <el-image 
                v-if="item.image_url" 
                :src="item.image_url" 
                fit="cover" 
                loading="lazy"
                class="card-image"
              >
                <template #placeholder>
                  <div class="image-placeholder shimmer-placeholder"></div>
                </template>
                <template #error>
                  <div class="image-error">
                    <el-icon><Picture /></el-icon>
                  </div>
                </template>
              </el-image>
              <div v-else class="explore-no-image">
                <el-icon><Picture /></el-icon>
              </div>
              
              <!-- 悬浮时的遮罩与内容 -->
              <div class="card-overlay">
                <div class="overlay-content">
                  <p class="card-prompt-text">{{ item.content }}</p>
                  <div class="overlay-actions">
                    <el-button type="primary" size="small" class="try-btn" round @click.stop="useSuggestion(item)">
                      <el-icon style="margin-right: 4px"><MagicStick /></el-icon> {{ item.need_reference_image ? '选择参考图' : '画同款' }}
                    </el-button>
                  </div>
                </div>
              </div>
            </div>

            <!-- 每列底部的骨架屏占位 -->
            <template v-if="loadingInspirations">
              <div 
                v-for="i in (publicInspirations.length === 0 ? 5 : 2)" 
                :key="'skeleton-' + colIndex + '-' + i" 
                class="prompt-card"
              >
                <div class="card-image shimmer-placeholder" :style="{ aspectRatio: (colIndex + i) % 3 === 0 ? '4/3' : ((colIndex + i) % 2 === 0 ? '1/1' : '3/4'), border: 'none' }"></div>
              </div>
            </template>
          </div>

        </div>

        <!-- 空状态 -->
        <el-empty 
          v-if="publicInspirations.length === 0 && !loadingInspirations" 
          description="暂无内容，敬请期待" 
          :image-size="100" 
        />
        
        <!-- 没有更多数据提示 -->
        <div class="no-more-data" v-if="noMoreData && publicInspirations.length > 0 && !loadingInspirations">
          <span class="divider-line"></span>
          <span>已经到底啦</span>
          <span class="divider-line"></span>
        </div>
      </div>
    </div>

    <!-- 灵感详情弹窗 (高级版分栏设计) -->
    <el-dialog 
      v-model="showPromptDialog" 
      :width="props.windowWidth <= 768 ? '90%' : '900px'"
      align-center 
      destroy-on-close
      :show-close="false"
      class="advanced-inspiration-dialog"
    >
      <div class="advanced-split-layout" :class="{ 'has-image': !!currentInspirationDetail?.image_url }">
        
        <!-- 左侧图片区 (有图片才显示) -->
        <div v-if="currentInspirationDetail?.image_url" class="advanced-left-image">
          <!-- 动态模糊背景 -->
          <div class="advanced-blur-bg" :style="{ backgroundImage: `url(${currentInspirationDetail.image_url})` }"></div>

          <!-- 移动端悬浮关闭按钮 -->
          <div class="mobile-close-btn" @click="showPromptDialog = false">
            <el-icon><Close /></el-icon>
          </div>
          
          <el-image 
            :src="currentInspirationDetail.image_url" 
            fit="contain"
            class="advanced-main-image"
            :preview-src-list="[currentInspirationDetail.image_url]"
            :preview-teleported="true"
          >
            <template #placeholder>
              <div class="advanced-image-skeleton">
                <div class="skeleton-pulse-ring"></div>
                <div class="skeleton-pulse-ring delay"></div>
                <el-icon class="skeleton-glow-icon"><Picture /></el-icon>
                <span class="skeleton-text">正在渲染高清画作...</span>
              </div>
            </template>
          </el-image>
        </div>

        <!-- 右侧内容区 -->
        <div class="advanced-right-content">
          <div class="advanced-right-header">
            <div class="advanced-title">
              <el-icon><MagicStick /></el-icon>
              <span>灵感详情</span>
            </div>
            <!-- PC端关闭按钮 -->
            <div class="desktop-close-btn" @click="showPromptDialog = false">
              <el-icon><Close /></el-icon>
            </div>
          </div>
          
          <div class="advanced-right-body">
            <div class="advanced-prompt-container">
              <div class="advanced-prompt-label">PROMPT 提示词</div>
              <div class="advanced-prompt-text">
                {{ currentInspirationDetail?.content }}
              </div>
            </div>
          </div>

          <div class="advanced-right-footer">
            <el-button class="advanced-btn use-btn" type="primary" @click="useSuggestionFromDialog">
              <el-icon><Position /></el-icon> {{ currentInspirationDetail?.need_reference_image ? '选择参考图' : '画同款' }}
            </el-button>
          </div>
        </div>

      </div>
    </el-dialog>
  </main>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, onUnmounted, watch, nextTick } from 'vue'
import { useRouter } from 'vue-router'
import { Expand, Picture, Coin, Edit, Position, Cpu, Crop, Close, MagicStick, Check } from '@element-plus/icons-vue'
import { ElMessage } from 'element-plus'
import request from '@/utils/request'
import { useAuthStore } from '@/store/auth'

const props = defineProps<{
  form: any;
  availableModels: any[];
  currentModelInfo: any;
  currentModelRatios: string[];
  currentModelResolutions: string[];
  currentModelImageCounts: (number | string)[];
  currentModelMaxRefImages: number;
  hasFreeResolution: Function;
  getTierConfigForRes: Function;
  isCurrentResolutionMaintenance: boolean;
  windowWidth: number;
  triggerUpload: () => void;
  clearReferenceImage: (index: number) => void;
  handleFileUpload: (e: Event) => void;
}>()

const emit = defineEmits(['toggle-menu', 'points-click', 'new-chat', 'generate'])

const router = useRouter()
const authStore = useAuthStore()

const isLoggedIn = computed(() => !!authStore.token)
const userPoints = computed(() => {
  const user = authStore.userInfo
  return user ? (user.points || 0) : 0
})

const isInputFocused = ref(false)
const isInputShrunk = ref(true)
const isMobileRefExpanded = ref(false)

// 键盘高度偏移
const keyboardOffset = ref(0)
const handleVisualViewport = () => {
  if (window.visualViewport) {
    // 当视觉视口高度小于窗口内部高度时，说明键盘弹起了
    const offset = window.innerHeight - window.visualViewport.height
    if (offset > 0 && isInputFocused.value) {
      keyboardOffset.value = offset
    } else {
      keyboardOffset.value = 0
    }
  }
}

// 处理失去焦点：如果输入框有内容，则不收缩
const handleInputBlur = () => {
  isInputFocused.value = false
  // 移除在 blur 时的收缩逻辑，收缩逻辑统一交给 handleClickOutside 处理，
  // 避免点击上传按钮打开文件选择器时输入框自动收缩。
}

// 处理输入事件：只要有输入，强制保持展开状态
const handleInput = () => {
  if (props.form.prompt.trim()) {
    isInputShrunk.value = false
  }
}

// Click outside to shrink input and remove focus
const handleClickOutside = (e: MouseEvent | TouchEvent) => {
  const target = e.target as HTMLElement
  
  // 排除 Element Plus 的弹出层，防止点击下拉菜单等 teleport 元素时收缩
  if (target.closest('.el-popper') || target.closest('.el-select-dropdown') || target.closest('.el-dialog') || target.closest('.el-overlay')) {
    return
  }

  if (!target.closest('.explore-input-wrapper') && !target.closest('.input-area-wrapper')) {
    isInputFocused.value = false
    isInputShrunk.value = true
  }
  // 点击外部时收起移动端的图片堆叠
  if (!target.closest('.reference-image-preview-inline')) {
    isMobileRefExpanded.value = false
  }
}

watch(() => props.form.reference_image?.length, (newLen, oldLen) => {
  if (newLen > (oldLen || 0) && props.windowWidth <= 768) {
    isMobileRefExpanded.value = true
    isInputShrunk.value = false
  }
})

const handleSend = (e?: KeyboardEvent) => {
  if (e && e.isComposing) return
  if (props.form.prompt.trim()) {
    // emit('generate') 
    // 如果想要点击后去生成页面，可以直接跳转
    router.push({ path: '/', query: { q: props.form.prompt } })
  }
}

const handleShiftEnter = (e: KeyboardEvent) => {
  if (e.isComposing) return
  e.preventDefault()
  const textarea = e.target as HTMLTextAreaElement
  const start = textarea.selectionStart
  const end = textarea.selectionEnd
  const value = props.form.prompt
  props.form.prompt = value.substring(0, start) + '\n' + value.substring(end)
  // 重新聚焦并设置光标位置
  nextTick(() => {
    textarea.focus()
    textarea.selectionStart = textarea.selectionEnd = start + 1
  })
}

// Inspirations logic
const publicInspirations = ref<{id: number, content: string, image_url: string}[]>([])
const loadingInspirations = ref(false)
const currentPage = ref(1)
const totalInspirations = ref(0)
const noMoreData = ref(false)

// 分列计算逻辑 (瀑布流完美布局)
const columnsCount = ref(5) // 默认5列

const updateColumnsCount = () => {
  const width = window.innerWidth
  if (width <= 800) columnsCount.value = 2
  else if (width <= 1200) columnsCount.value = 3
  else if (width <= 1600) columnsCount.value = 4
  else columnsCount.value = 5
}

// 将一维数据按照最短列算法分配到多列中
const masonryColumnsData = computed(() => {
  const cols = Array.from({ length: columnsCount.value }, () => [] as any[])
  // 简单轮询分配，为了性能和视觉稳定性，按顺序依次分配到各列
  publicInspirations.value.forEach((item, index) => {
    const colIndex = index % columnsCount.value
    cols[colIndex].push(item)
  })
  return cols
})

// Categories logic
const categories = ref<any[]>([])
const currentMainCategory = ref('')
const currentSubCategory = ref('')

const currentSubCategories = computed(() => {
  if (!currentMainCategory.value) return []
  const target = categories.value.find(c => c.name === currentMainCategory.value)
  return target ? (target.sub || []) : []
})

const fetchCategories = async () => {
  try {
    const res: any = await request.get('/api/public/inspirations/categories')
    if (res.data?.categories) {
      categories.value = res.data.categories
    } else if (res.data?.data?.categories) {
      categories.value = res.data.data.categories
    }
  } catch (error) {
    console.error('Failed to fetch categories', error)
  }
}

const selectMainCategory = (cat: string) => {
  currentMainCategory.value = cat
  currentSubCategory.value = '' // Reset sub category when main changes
  resetAndFetch()
}

const selectSubCategory = (sub: string) => {
  currentSubCategory.value = sub
  resetAndFetch()
}



const resetAndFetch = () => {
  currentPage.value = 1
  noMoreData.value = false
  
  // 仅在当前页面滚动距离较大时，才触发平滑滚动到顶部，避免轻微滚动导致的视差抖动
  const scrollContainer = document.querySelector('.explore-content-area')
  if (scrollContainer && scrollContainer.scrollTop > 50) {
    scrollContainer.scrollTo({ top: 0, behavior: 'smooth' })
  }
  
  fetchPublicInspirations()
}

const fetchPublicInspirations = async () => {
  if (loadingInspirations.value || noMoreData.value) return
  
  loadingInspirations.value = true
  try {
    const params: any = {
      page: currentPage.value,
      page_size: 20
    }
    if (currentMainCategory.value) params.main_category = currentMainCategory.value
    if (currentSubCategory.value) params.sub_category = currentSubCategory.value

    const res: any = await request.get('/api/public/inspirations/list', { params })
    
    let newList: any[] = []
    
    // 处理后端返回的不同数据格式
    if (res.data?.list) {
      newList = res.data.list
      if (res.data.total !== undefined) {
        totalInspirations.value = res.data.total
      }
    } else if (res.data?.data?.list) {
      newList = res.data.data.list
      if (res.data.data.total !== undefined) {
        totalInspirations.value = res.data.data.total
      }
    } else if (Array.isArray(res.data)) {
      newList = res.data
    } else if (Array.isArray(res.data?.data)) {
      newList = res.data.data
    }
    
    if (newList.length === 0) {
      noMoreData.value = true
      if (currentPage.value === 1) {
        publicInspirations.value = []
      }
    } else {
      if (currentPage.value === 1) {
        // 直接替换
        publicInspirations.value = newList
      } else {
        // 过滤掉可能重复的数据，再追加
        const existingIds = new Set(publicInspirations.value.map(item => item.id))
        const uniqueNewList = newList.filter(item => !existingIds.has(item.id))
        
        if (uniqueNewList.length === 0) {
           // 如果返回的数据全部是重复的，也认为到底了
           noMoreData.value = true
        } else {
           publicInspirations.value = [...publicInspirations.value, ...uniqueNewList]
        }
      }
      
      // Check if we loaded all data
      if (newList.length < 20) {
        noMoreData.value = true
      } else if (!noMoreData.value) {
        currentPage.value++
      }
    }
  } catch (error) {
    console.error('Failed to fetch inspirations', error)
  } finally {
    setTimeout(() => {
      loadingInspirations.value = false
    }, 500) // 延迟500ms关闭loading状态，让骨架屏能够展示一会
  }
}

// Scroll handler for infinite scrolling
const lastScrollTop = ref(0)
const handleScroll = () => {
  const scrollContainer = document.querySelector('.explore-content-area')
  if (!scrollContainer) return
  
  const scrollHeight = scrollContainer.scrollHeight
  const scrollTop = scrollContainer.scrollTop
  const clientHeight = scrollContainer.clientHeight
  
  // 忽略极小的滚动以防抖动
  if (Math.abs(scrollTop - lastScrollTop.value) > 10) {
    // 只要发生有效滑动，就关闭移动端参考图展开 (移除滑动收缩输入框的逻辑)
    isMobileRefExpanded.value = false
    lastScrollTop.value = scrollTop
  }
  
  // If user scrolls within 100px of bottom, fetch more
  if (scrollHeight - scrollTop - clientHeight < 100) {
    fetchPublicInspirations()
  }
}

const useSuggestion = (item: any) => {
  const isString = typeof item === 'string'
  const text = isString ? item : item.content
  props.form.prompt = text
  
  if (!isString && item.need_reference_image) {
    ElMessage({
      message: '该提示词建议您上传参考图，请在弹出的窗口中选择',
      type: 'warning',
      duration: 4000,
      showClose: true
    })
    if (typeof props.triggerUpload === 'function') {
      props.triggerUpload()
    } else {
      // 触发一个自定义事件或者通过总线通知
      window.dispatchEvent(new CustomEvent('trigger-upload-from-explore'))
    }
  }
  router.push({ path: '/', query: { q: text } })
}

// 提示词详情展示逻辑
const showPromptDialog = ref(false)
const currentInspirationDetail = ref<any>(null)

const handleShowPrompt = (item: any) => {
  currentInspirationDetail.value = item
  showPromptDialog.value = true
}

// const handleCopyPrompt = async () => {
//   if (!currentInspirationDetail.value?.content) return
//   try {
//     await navigator.clipboard.writeText(currentInspirationDetail.value.content)
//     ElMessage.success('提示词已复制')
//   } catch (err) {
//     ElMessage.error('复制失败，请手动复制')
//   }
// }

const useSuggestionFromDialog = () => {
  if (currentInspirationDetail.value?.content) {
    useSuggestion(currentInspirationDetail.value)
    showPromptDialog.value = false
  }
}

onMounted(() => {
  fetchCategories()
  fetchPublicInspirations()
  updateColumnsCount() // 初始化列数
  window.addEventListener('resize', updateColumnsCount) // 监听窗口大小变化
  document.addEventListener('mousedown', handleClickOutside)
  document.addEventListener('touchstart', handleClickOutside)
  
  if (window.visualViewport) {
    window.visualViewport.addEventListener('resize', handleVisualViewport)
    window.visualViewport.addEventListener('scroll', handleVisualViewport)
  }
})

onUnmounted(() => {
  window.removeEventListener('resize', updateColumnsCount)
  document.removeEventListener('mousedown', handleClickOutside)
  document.removeEventListener('touchstart', handleClickOutside)
  
  if (window.visualViewport) {
    window.visualViewport.removeEventListener('resize', handleVisualViewport)
    window.visualViewport.removeEventListener('scroll', handleVisualViewport)
  }
})
</script>

<style scoped>
.explore-main {
  flex: 1;
  display: flex;
  flex-direction: column;
  height: 100%;
  overflow: hidden;
  background-color: #f9fafb; /* 浅灰底色，贴近豆包 */
  position: relative;
}

/* 样式部分保持 Generate.vue 的 input 样式结构 */
.input-area-wrapper {
  position: relative;
  width: 100%;
  max-width: 1400px;
  margin: 0 auto;
  background-color: transparent;
  border-radius: 20px;
  transition: transform 0.1s ease-out, padding 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 10px 0; /* PC端去除左右内边距，使其与内容对齐 */
  pointer-events: auto;
}

.input-area-wrapper .input-container,
.input-area-wrapper .input-tools-container {
  pointer-events: auto;
}

.input-area-wrapper.is-shrunk {
  padding-bottom: 0px !important; /* 取消悬浮高度增加，让其更紧凑，稍微向上移一点 */
}

.input-area-wrapper.is-shrunk::before {
  opacity: 0;
}

.input-area-wrapper.is-shrunk .input-tools-container {
  height: 0;
  margin: 0;
  opacity: 0;
  overflow: hidden;
  pointer-events: none;
  transform: translateY(10px);
  position: absolute; /* 防止收缩时占据高度 */
}

.input-area-wrapper.is-shrunk .input-box {
  width: 100%;
  min-height: 40px !important;
  border-radius: 0;
  box-shadow: none;
  border-color: transparent;
  background-color: transparent;
  backdrop-filter: none;
  padding: 0 !important;
  gap: 8px !important; /* 收缩状态下减小图片与文字之间的间距 */
  margin-top: 0;
  align-items: center !important;
}

.input-area-wrapper.is-shrunk .initial-upload-item {
  width: 28px !important;
  height: 32px !important;
  margin-right: 0 !important; /* 取消额外边距，完全靠 gap 控制 */
  margin-top: 0 !important; /* 收缩状态下取消下移 */
  transform: rotate(-8deg) !important;
  border-radius: 6px;
  position: relative !important;
  pointer-events: auto; /* 允许交互 */
}

.input-area-wrapper.is-shrunk .initial-upload-item:hover,
.input-area-wrapper.is-shrunk .initial-upload-item:active {
  transform: translateY(-4px) scale(1.1) rotate(0deg) !important; /* 相对定位使用 px 向上浮动 */
  z-index: 30 !important;
  box-shadow: 0 8px 16px rgba(0, 0, 0, 0.15) !important;
}

.input-area-wrapper.is-shrunk .initial-upload-item .upload-icon-svg,
.input-area-wrapper.is-shrunk .reference-upload-item-inline .upload-icon-svg {
  font-size: 14px !important; /* 调整图标尺寸 */
}

.input-area-wrapper.is-shrunk .reference-image-preview-inline {
  margin-right: 0 !important; /* 取消额外边距，完全靠 gap 控制 */
  margin-top: 0 !important; /* 收缩状态下取消下移 */
  pointer-events: auto; /* 允许交互 */
  overflow: visible !important; /* 收缩状态下强制取消所有滚动条 */
}

.input-area-wrapper.is-shrunk .preview-list-inline {
  width: 28px !important;
  height: 32px !important; /* 恢复为收缩状态应有的高度，因为不展开了 */
  transition: all 0.3s ease;
  pointer-events: auto;
  cursor: pointer;
  margin-right: -4px !important; /* 让图片和输入框靠得更近一点 */
}

@media (max-width: 768px) {
  .input-area-wrapper.is-shrunk .preview-list-inline {
    transform: translateY(-10px); /* 仅在手机端向上修正位置 */
    margin-right: -6px !important; /* 手机端靠得更近一点 */
  }
}

.input-area-wrapper.is-shrunk .preview-container-inline {
  width: 28px !important;
  height: 32px !important;
  border-radius: 6px;
  transform: translateY(-50%) rotate(var(--rotate)) !important;
  pointer-events: none !important; /* 收缩状态下禁止图片本身的点击事件（防全屏预览） */
}

.input-area-wrapper.is-shrunk .reference-upload-item-inline {
  width: 28px !important;
  height: 32px !important;
  border-radius: 6px;
  transform: translateY(-50%) rotate(var(--rotate)) !important;
  pointer-events: none !important;
}

/* 强制收缩状态下的所有图片永远重叠，禁止任何形式的展开和单独悬浮 */
.input-area-wrapper.is-shrunk .preview-list-inline:hover,
.input-area-wrapper.is-shrunk .preview-list-inline.is-expandable:hover,
.input-area-wrapper.is-shrunk .preview-list-inline.is-mobile-expanded,
.input-area-wrapper.is-shrunk .preview-list-inline.is-expandable.is-mobile-expanded {
  width: 28px !important; /* 强制保持收缩状态宽度，禁止展开 */
  margin-right: 0 !important; /* 强制取消 hover/展开 状态下的右侧间距，防止文字位移 */
}

.input-area-wrapper.is-shrunk .preview-list-inline:hover .preview-container-inline,
.input-area-wrapper.is-shrunk .preview-list-inline.is-mobile-expanded .preview-container-inline,
.input-area-wrapper.is-shrunk .preview-list-inline.is-expandable:hover .preview-container-inline,
.input-area-wrapper.is-shrunk .preview-list-inline.is-expandable.is-mobile-expanded .preview-container-inline {
  left: 0 !important; /* 强制所有图片重叠在一起 */
  transform: translateY(-50%) rotate(var(--rotate)) !important; /* 强制取消所有放大和上浮，只保留堆叠 */
  z-index: calc(20 - var(--index)) !important;
}

.input-area-wrapper.is-shrunk .preview-list-inline:hover .reference-upload-item-inline,
.input-area-wrapper.is-shrunk .preview-list-inline.is-mobile-expanded .reference-upload-item-inline,
.input-area-wrapper.is-shrunk .preview-list-inline.is-expandable:hover .reference-upload-item-inline,
.input-area-wrapper.is-shrunk .preview-list-inline.is-expandable.is-mobile-expanded .reference-upload-item-inline {
  display: none !important; /* 收缩状态下隐藏额外上传按钮 */
}

/* 强制取消单张图片的悬浮/点击放大效果 */
.input-area-wrapper.is-shrunk .preview-container-inline:hover,
.input-area-wrapper.is-shrunk .preview-container-inline:active {
  transform: translateY(-50%) rotate(var(--rotate)) !important;
  box-shadow: none !important;
}

.input-area-wrapper.is-shrunk .collapsed-upload-badge,
.input-area-wrapper.is-shrunk .remove-btn-inline {
  display: none !important; /* 收缩状态下隐藏图片上的删除和添加小按钮，保持图标纯净 */
}

.input-area-wrapper.is-shrunk :deep(.chat-input) {
  min-height: 36px !important;
  height: 36px !important;
}

.input-area-wrapper.is-shrunk .input-box {
  min-height: 36px !important;
  height: 36px !important;
}

.input-area-wrapper.is-shrunk :deep(.chat-input.has-references) {
  min-height: 36px !important;
  height: 36px !important;
}

.input-area-wrapper.is-shrunk .input-box.has-refs {
  min-height: 36px !important;
  height: 36px !important;
}

.input-area-wrapper.is-shrunk :deep(.chat-input.has-references .el-textarea__inner) {
  min-height: 36px !important;
  height: 36px !important;
  padding-top: 9px !important;
  padding-bottom: 9px !important;
  max-height: 36px !important;
  margin-left: -8px !important;
}

.input-area-wrapper.is-shrunk :deep(.chat-input .el-textarea__inner) {
    min-height: 36px !important;
    height: 36px !important;
    padding-top: 9px !important;
    padding-bottom: 9px !important;
    font-size: 13px !important;
    line-height: 18px !important;
    display: block !important;
    max-height: 36px !important;
    overflow: hidden !important;
    white-space: nowrap !important;
    box-sizing: border-box !important;
    margin-left: -8px !important;
  }
  
  .input-area-wrapper.is-shrunk :deep(.chat-input .el-textarea__inner::placeholder) {
    white-space: nowrap !important;
    overflow: hidden !important;
    text-overflow: ellipsis !important;
    line-height: 18px !important;
  }

.input-area-wrapper.is-shrunk .send-btn {
  transform: scale(0.85);
  margin-bottom: 0;
  margin-right: -4px;
}

@media (max-width: 768px) {
  .input-area-wrapper {
    padding-bottom: calc(24px + env(safe-area-inset-bottom));
  }
  
  .input-area-wrapper {
    padding: 0 16px !important;
    width: 100%;
    margin-left: 0;
    transform: scale(1);
    transform-origin: center top;
    margin-bottom: 0;
  }
  
    .input-container {
    padding: 8px 6px 6px !important;
    border-radius: 12px !important;
    border: 1px solid rgba(229, 231, 235, 0.5) !important;
    box-shadow: 0 4px 20px rgba(0, 0, 0, 0.05), 0 2px 8px rgba(0, 0, 0, 0.02) !important;
  }
  
  .reference-upload-item.initial-upload-item {
    width: 40px;
    height: 48px;
    margin-right: 12px; /* 适当增加初始状态下旋转卡片与文字的间距，防止遮挡 */
    margin-top: -6px; /* 向上移一点 */
  }
  
  .preview-list-inline {
    width: 48px; /* 恢复为电脑端的宽度 */
    height: 72px; /* 恢复为电脑端的高度 */
  }
  
  .preview-container-inline {
    width: 48px; /* 恢复为电脑端的宽度 */
    height: 56px; /* 恢复为电脑端的高度 */
    transform: translateY(-50%) rotate(var(--rotate)) !important; /* 恢复原有的旋转效果 */
  }
  
  .input-box .preview-list-inline.is-expandable.is-mobile-expanded {
    width: calc(48px + (var(--total-items, 1) - 1) * 56px); /* 与电脑端间距保持完全一致 */
    margin-right: 16px; /* 恢复适当间距，防止旋转图片遮挡右侧输入框文字 */
  }
  
  .input-box .preview-list-inline.is-expandable.is-mobile-expanded .preview-container-inline {
    left: calc(var(--index) * 56px); /* 与电脑端间距保持完全一致 */
    transform: translateY(-50%) rotate(var(--rotate)) scale(1) !important;
    z-index: calc(20 - var(--index)) !important;
  }
  
  .input-box .preview-list-inline.is-expandable.is-mobile-expanded .preview-container-inline:active {
    transform: translateY(-65%) scale(1.1) rotate(var(--rotate)) !important;
    z-index: 30 !important;
    box-shadow: 0 8px 16px rgba(0, 0, 0, 0.15) !important; /* 增加阴影增强浮动感 */
  }
  
  .input-box .preview-list-inline.is-expandable.is-mobile-expanded .reference-upload-item-inline {
    left: calc(var(--index) * 56px); /* 与电脑端间距保持完全一致 */
    width: 48px; /* 恢复为电脑端的宽度 */
    height: 56px; /* 恢复为电脑端的高度 */
    transform: translateY(-50%) rotate(var(--rotate, 0deg)) !important;
    opacity: 1;
    z-index: 10 !important;
  }
  
  .input-box .preview-list-inline.is-expandable.is-mobile-expanded .reference-upload-item-inline:active {
    background-color: #f3f4f6;
    border-color: #d1d5db;
    color: #4b5563;
    transform: translateY(-65%) scale(1.1) rotate(var(--rotate, 0deg)) !important;
    z-index: 30 !important;
    box-shadow: 0 8px 16px rgba(0, 0, 0, 0.15) !important;
  }
  
  /* 移动端展开时显示删除按钮和隐藏小上传按钮 */
  .preview-list-inline.is-mobile-expanded .remove-btn-inline {
    opacity: 1;
  }
  
  .input-box .preview-list-inline.is-expandable.is-mobile-expanded .collapsed-upload-badge {
    opacity: 0;
    pointer-events: none;
  }

  .send-btn {
    margin-bottom: 0 !important; /* 手机端取消底部外边距 */
  }
}

.input-area-wrapper::before {
  display: none;
}

.reference-image-preview {
  padding: 12px 12px 0 12px;
}

.preview-list {
  position: relative;
  height: 72px;
  padding: 4px 12px;
  background-color: transparent;
  display: flex;
  align-items: center;
  pointer-events: none;
}
.preview-container {
  position: absolute;
  width: 48px;
  height: 64px;
  border-radius: 8px;
  overflow: hidden;
  border: 1px solid #e5e7eb;
  transition: transform 0.3s cubic-bezier(0.4, 0, 0.2, 1), box-shadow 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  transform-origin: center center;
  background-color: #fff;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.08);
  pointer-events: auto;
}
.preview-container:hover {
  transform: translateY(-8px) rotate(0deg) scale(1.05) !important;
  z-index: 20 !important;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
}

.preview-img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.remove-btn {
  position: absolute;
  top: 4px;
  right: 4px;
  width: 20px;
  height: 20px;
  background: rgba(0, 0, 0, 0.5);
  color: white;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  font-size: 12px;
  pointer-events: auto;
}

.remove-btn:hover {
  background: rgba(0, 0, 0, 0.7);
}

.reference-upload-item {
  position: absolute;
  width: 48px;
  height: 64px;
  border-radius: 8px;
  background-color: #f9fafb;
  border: 1px dashed #d1d5db;
  display: flex;
  justify-content: center;
  align-items: center;
  cursor: pointer;
  transition: transform 0.3s cubic-bezier(0.4, 0, 0.2, 1), box-shadow 0.3s cubic-bezier(0.4, 0, 0.2, 1), background-color 0.3s;
  transform-origin: center center;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.04);
  pointer-events: auto;
}
.reference-upload-item.initial-upload-item {
  position: relative;
  transform: rotate(-8deg);
  margin-right: 8px;
  margin-top: 8px; /* 向下移一点 */
  width: 48px;
  height: 56px;
  border-radius: 8px;
  border: 1px dashed #e5e7eb;
  background-color: #f9fafb; /* 恢复浅灰色背景 */
  color: #9ca3af;
  flex-shrink: 0;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.04); /* 恢复轻微阴影 */
  z-index: 10;
  pointer-events: auto;
  cursor: pointer;
  display: flex;
  justify-content: center;
  align-items: center;
  transform-origin: center center;
  transition: all 0.2s ease;
}
.reference-upload-item.initial-upload-item:hover {
  background-color: #f3f4f6; /* 悬浮时加深背景 */
  border-color: #d1d5db;
  color: #4b5563;
  transform: rotate(0deg) scale(1.02) !important;
  z-index: 20 !important;
}
.reference-upload-item.initial-upload-item .upload-icon-svg {
  font-size: 20px;
}

.reference-upload-item:hover {
  background-color: #f3f4f6;
  border-color: #9ca3af;
  transform: translateY(-8px) rotate(0deg) scale(1.05) !important;
  z-index: 20 !important;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
}

.reference-upload-content {
  display: flex;
  justify-content: center;
  align-items: center;
}
.upload-icon-svg {
  font-size: 20px;
  color: #6b7280;
}
.initial-upload-item .upload-icon-svg {
  font-size: 20px;
}

.reference-image-preview-inline {
  position: relative;
  z-index: 10;
  display: flex;
  align-items: center;
  margin-right: 4px; /* 电脑端默认减小右侧间距，拉近和文字距离 */
  margin-top: 8px; /* 向下移一点 */
  flex-shrink: 0;
}

.preview-list-inline {
  position: relative;
  width: 48px;
  height: 56px;
  background-color: transparent;
  display: flex;
  align-items: center;
  transition: width 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  overflow: visible;
}

.input-box .preview-list-inline.is-expandable.is-mobile-expanded {
  width: calc(48px + (var(--total-items, 1) - 1) * 56px);
}

.reference-group-hover-trigger {
  position: absolute;
  left: 0;
  top: -10px;
  bottom: -10px;
  width: 100%;
  z-index: -1;
  transition: width 0.3s;
}

.preview-list-inline.is-mobile-expanded .reference-group-hover-trigger {
  width: calc(48px + (var(--total-items, 1) - 1) * 56px);
}

.preview-container-inline {
  position: absolute;
  left: 0;
  top: 50%;
  width: 48px;
  height: 56px;
  border-radius: 8px;
  overflow: hidden;
  border: none;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  transform: translateY(-50%) rotate(var(--rotate));
  transform-origin: center center;
  background-color: transparent;
  box-shadow: none;
  pointer-events: auto;
}

.input-box .preview-list-inline.is-expandable.is-mobile-expanded .preview-container-inline {
  left: calc(var(--index) * 56px);
  z-index: calc(20 - var(--index)) !important;
}

.preview-img-inline {
  width: 100%;
  height: 100%;
  object-fit: cover;
  border-radius: 8px;
  box-shadow: 0 1px 4px rgba(0, 0, 0, 0.05);
  border: 1px solid rgba(0, 0, 0, 0.04);
}

.remove-btn-inline {
  position: absolute;
  top: 4px;
  right: 4px;
  width: 20px;
  height: 20px;
  background: rgba(0, 0, 0, 0.5);
  color: white;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  font-size: 12px;
  pointer-events: auto;
  opacity: 0;
  transition: opacity 0.2s, background 0.2s;
}

.preview-list-inline.is-mobile-expanded .remove-btn-inline {
  opacity: 1;
}

.collapsed-upload-badge {
  position: absolute;
  bottom: 2px;
  right: 2px;
  width: 16px;
  height: 16px;
  background-color: rgba(0, 0, 0, 0.6);
  color: white;
  border-radius: 4px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  font-size: 10px;
  pointer-events: auto;
  opacity: 1;
  transition: opacity 0.2s, background-color 0.2s;
  z-index: 10;
}

.input-box .preview-list-inline.is-expandable.is-mobile-expanded .collapsed-upload-badge {
  opacity: 0;
  pointer-events: none;
}

.reference-upload-item-inline {
  position: absolute;
  left: 0;
  top: 50%;
  width: 48px;
  height: 56px;
  border-radius: 8px;
  background-color: #f9fafb;
  border: 1px dashed #e5e7eb;
  display: flex;
  justify-content: center;
  align-items: center;
  cursor: pointer;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  transform-origin: center center;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.04);
  pointer-events: auto;
  opacity: 1;
  color: #9ca3af;
  transform: translateY(-50%) rotate(var(--rotate, 0deg));
}

.input-box .preview-list-inline.is-expandable.is-mobile-expanded .reference-upload-item-inline {
  left: calc(var(--index) * 56px);
  opacity: 1;
  z-index: 10 !important;
}

/* Desktop Hover Expansion Rules */
@media (hover: hover) and (pointer: fine) {
  .input-box .preview-list-inline.is-expandable:hover {
    width: calc(48px + (var(--total-items, 1) - 1) * 56px);
    margin-right: 4px; /* 电脑端悬浮展开时也减小右侧间距 */
  }
  
  .preview-list-inline:hover .reference-group-hover-trigger {
    width: calc(48px + (var(--total-items, 1) - 1) * 56px);
  }

  .input-box .preview-list-inline.is-expandable:hover .preview-container-inline {
    left: calc(var(--index) * 56px);
    z-index: calc(20 - var(--index)) !important;
  }

  .preview-container-inline:hover {
    transform: translateY(-65%) scale(1.1) rotate(var(--rotate)) !important;
    z-index: 30 !important;
    box-shadow: 0 8px 16px rgba(0, 0, 0, 0.15) !important;
  }

  .preview-container-inline:hover .remove-btn-inline {
    opacity: 1;
  }
  
  .remove-btn-inline:hover {
    background: rgba(0, 0, 0, 0.7);
  }

  .collapsed-upload-badge:hover {
    background-color: rgba(0, 0, 0, 0.8);
  }

  .input-box .preview-list-inline.is-expandable:hover .collapsed-upload-badge {
    opacity: 0;
    pointer-events: none;
  }

  .input-box .preview-list-inline.is-expandable:hover .reference-upload-item-inline {
    left: calc(var(--index) * 56px);
    opacity: 1;
    z-index: 10 !important;
  }

  .reference-upload-item-inline:hover {
    background-color: #f3f4f6;
    border-color: #d1d5db;
    color: #4b5563;
    transform: translateY(-65%) scale(1.1) rotate(var(--rotate, 0deg)) !important;
    z-index: 30 !important;
    box-shadow: 0 8px 16px rgba(0, 0, 0, 0.15) !important;
  }
}



.input-container {
  width: calc(100% - 40px);
  max-width: 860px;
  background: rgba(255, 255, 255, 0.92);
  backdrop-filter: blur(20px);
  -webkit-backdrop-filter: blur(20px);
  border: 1px solid rgba(229, 231, 235, 0.85);
  border-radius: 24px;
  box-sizing: border-box;
  box-shadow:
    0 2px 6px rgba(15, 23, 42, 0.04),
    0 12px 32px rgba(15, 23, 42, 0.06),
    inset 0 0 0 1px rgba(255, 255, 255, 0.65);
  padding: 12px 16px;
  display: flex;
  flex-direction: column;
  gap: 8px;
  margin: 0 auto;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  position: relative;
  overflow: visible;
}

.search-suggestions {
  width: 100%;
  max-width: 1400px;
  margin: 16px auto 0;
  padding: 0 16px;
  display: flex;
  align-items: center;
  justify-content: center;
  animation: fadeIn 0.5s ease-out;
}

.suggestion-tags {
  display: flex;
  flex-wrap: wrap;
  gap: 12px;
  align-items: center;
  justify-content: center;
}

.suggestion-label {
  display: inline-flex;
  align-items: center;
  gap: 4px;
  font-size: 14px;
  color: #6b7280;
  font-weight: 500;
  margin-right: 4px;
}

.suggestion-tag {
  display: inline-flex;
  align-items: center;
  font-size: 13px;
  color: #4b5563;
  background-color: rgba(255, 255, 255, 0.7);
  backdrop-filter: blur(8px);
  padding: 6px 14px;
  border-radius: 20px;
  cursor: pointer;
  transition: all 0.2s ease;
  border: 1px solid rgba(229, 231, 235, 0.8);
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.02);
}

.suggestion-tag:hover {
  background-color: #ffffff;
  color: #111827;
  border-color: #d1d5db;
  transform: translateY(-2px);
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.05);
}

@keyframes fadeIn {
  from { opacity: 0; transform: translateY(5px); }
  to { opacity: 1; transform: translateY(0); }
}

@media (max-width: 768px) {
  .search-suggestions {
    padding: 0 8px;
    margin-top: 24px;
  }
  .suggestion-tags {
    justify-content: flex-start;
    overflow-x: auto;
    flex-wrap: nowrap;
    padding-bottom: 8px;
    -ms-overflow-style: none;
    scrollbar-width: none;
  }
  .suggestion-tags::-webkit-scrollbar {
    display: none;
  }
  .suggestion-tag {
    white-space: nowrap;
    flex-shrink: 0;
  }
  .suggestion-label {
    white-space: nowrap;
    flex-shrink: 0;
  }
}

.input-area-wrapper.is-shrunk .input-container {
  width: 60%;
  min-width: 320px;
  max-width: 600px;
  padding: 8px 16px;
  border-radius: 36px;
  background-color: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(16px);
  -webkit-backdrop-filter: blur(16px);
  box-shadow: 0 8px 32px rgba(15, 23, 42, 0.12);
  border-color: transparent;
  gap: 0;
}

@media (max-width: 768px) {
  .input-area-wrapper.is-shrunk .input-container {
    width: 90%;
    min-width: unset;
  }
}

.input-container:focus-within {
  border-color: rgba(229, 231, 235, 0.85);
  box-shadow:
    0 2px 6px rgba(15, 23, 42, 0.04),
    0 12px 32px rgba(15, 23, 42, 0.06),
    inset 0 0 0 1px rgba(255, 255, 255, 0.65);
}

.model-dropdown-wrapper {
  max-width: 100%;
  display: inline-flex;
}

.model-settings-btn {
  width: 100%;
  min-width: 100px;
  max-width: 200px;
  justify-content: space-between;
}

.model-btn-content {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  flex: 1;
  min-width: 0;
}

.cpu-icon {
  flex-shrink: 0;
}

.model-name-text {
  display: block !important;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  flex: 1;
  min-width: 0;
  text-align: left;
  line-height: 1.5;
  max-width: 200px;
}

.arrow-icon {
  font-size: 12px !important;
  color: #9ca3af !important;
}

.desktop-only {
  display: flex !important;
}

.mobile-only {
  display: none !important;
}

.input-tools-container {
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  transform-origin: bottom center;
  width: 100%;
  max-width: 100%;
  margin-bottom: 0px;
  display: flex;
  justify-content: flex-start;
  position: relative;
  z-index: 100;
  min-width: 0;
}

.input-tools {
    display: flex;
    gap: 8px;
    padding: 0;
    align-items: center;
    flex-wrap: wrap; /* 允许在空间不足时换行 */
    overflow: visible; /* 移除任何截断 */
}
.input-tools::-webkit-scrollbar {
  display: none;
}

:deep(.input-tools .el-select__wrapper) {
  box-shadow: none !important;
  background-color: #f3f4f6;
  border-radius: 8px;
  transition: background-color 0.2s;
}

:deep(.input-tools .el-select__wrapper:hover) {
  background-color: #e5e7eb;
}

:deep(.input-tools .el-select__wrapper.is-focused) {
  background-color: #e5e7eb;
}

.ref-btn {
  display: inline-flex;
  align-items: center;
  gap: 4px;
  border: 1px solid transparent;
  background-color: #f4f4f5;
  border-radius: 20px;
  color: #3f3f46;
  font-weight: 500;
  padding: 4px 14px;
  height: 34px;
  transition: all 0.25s cubic-bezier(0.4, 0, 0.2, 1);
}

.ref-btn:hover {
  background-color: #e4e4e7;
  color: #18181b;
  transform: translateY(-1px);
}

.ref-btn:active {
  transform: translateY(0);
}

.ref-btn .el-icon {
  font-size: 15px;
  color: #71717a;
  transition: color 0.25s ease;
}

.ref-btn:hover .el-icon {
  color: #18181b;
}

.size-select {
  width: 120px;
}

.input-box {
  position: relative;
  display: flex;
  align-items: center;
  width: 100%;
  box-sizing: border-box;
  max-width: 100%;
  background: transparent;
  border: none;
  border-radius: 20px;
  padding: 0;
  gap: 12px;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  z-index: 1;
  min-height: 44px;
  box-shadow: none;
}

.input-box:focus-within {
  border-color: transparent;
  box-shadow: none;
}

.input-box.has-refs {
  min-height: 76px;
  padding-bottom: 8px;
}

.input-area-wrapper.is-shrunk .input-box.has-refs {
  min-height: 28px !important;
  padding-bottom: 0px !important;
}

:deep(.chat-input) {
  flex: 1;
  position: relative;
  z-index: 1;
  background-color: transparent !important;
  --el-input-focus-border-color: transparent;
  --el-input-hover-border-color: transparent;
  --el-input-border-color: transparent;
  min-height: 44px;
}

:deep(.chat-input .el-textarea__inner) {
  padding-left: 12px !important;
  padding-right: 12px !important;
  box-shadow: none !important;
  background-color: transparent !important;
  padding-top: 11px !important;
  padding-bottom: 11px !important;
  font-size: 15px !important;
  line-height: 22px !important;
  min-height: 44px !important;
  color: #111827;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  border: none !important;
  font-family: inherit;
  display: block !important;
  height: auto;
  max-height: 250px !important;
  overflow-y: auto !important;
}

/* 电脑端滚动条样式 */
:deep(.chat-input .el-textarea__inner::-webkit-scrollbar) {
  width: 6px;
}
:deep(.chat-input .el-textarea__inner::-webkit-scrollbar-thumb) {
  background-color: rgba(156, 163, 175, 0.5);
  border-radius: 6px;
}
:deep(.chat-input .el-textarea__inner::-webkit-scrollbar-track) {
  background: transparent;
}

/* 移除了其他的 padding-left 规则 */

:deep(.chat-input .el-textarea__inner::placeholder) {
  color: #9ca3af;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  line-height: inherit !important;
}

.send-btn {
  margin-bottom: 4px;
  background-color: #111827 !important;
  border-color: #111827 !important;
  transition: transform 0.2s ease, opacity 0.2s, width 0.3s ease, box-shadow 0.2s ease, background-color 0.2s ease;
  box-shadow: 0 2px 8px rgba(17, 24, 39, 0.15);
  width: 36px;
  height: 36px;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 50%;
  color: white;
}

.send-btn.maintenance-btn {
  width: auto;
  min-width: 52px;
  border-radius: 16px;
  background-color: #9ca3af !important;
  border-color: #9ca3af !important;
  color: #fff !important;
  opacity: 1 !important;
  box-shadow: none;
}

.send-btn:hover:not(.is-disabled) {
  transform: scale(1.05);
  opacity: 0.9;
  box-shadow: 0 6px 16px rgba(17, 24, 39, 0.24);
}

.send-btn.is-disabled {
  background-color: #e5e7eb !important;
  border-color: #e5e7eb !important;
  color: #9ca3af !important;
}

/* 提示框样式优化 */
:deep(.el-tooltip__popper) {
  padding: 8px 12px;
  border-radius: 8px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
}
/* Responsive */
@keyframes text-shimmer {
  0% { background-position: 0% 50%; }
  100% { background-position: 100% 50%; }
}

.animate-text-shimmer {
  animation: text-shimmer 3s ease-in-out infinite alternate;
}

.gradient-text {
  background: linear-gradient(110deg, #4F46E5 0%, #9333EA 45%, #EC4899 100%);
  background-size: 250% 100%;
  -webkit-background-clip: text;
  background-clip: text;
  -webkit-text-fill-color: transparent;
  color: transparent;
  filter: drop-shadow(0 1px 2px rgba(0, 0, 0, 0.1));
}

@media (min-width: 769px) {
  .mobile-menu-btn {
    display: none;
  }
  .mobile-only {
    display: none !important;
  }
}

@media (max-width: 768px) {
  .settings-popover {
    width: calc(100vw - 32px) !important;
  }
  .header-actions {
    display: flex;
  }

  .mini-sidebar {
      width: 60px;
      position: absolute;
      top: 0;
      bottom: 0;
      left: 0;
      border-radius: 0;
      transform: translateX(-100%);
      transition: transform 0.3s cubic-bezier(0.4, 0.0, 0.2, 1);
      padding-bottom: calc(48px + env(safe-area-inset-bottom));
      z-index: 1010;
      display: flex;
      flex-direction: column;
      justify-content: space-between;
      box-sizing: border-box;
    }
    
    .history-sidebar {
      width: 240px;
      position: absolute;
      top: 0;
      bottom: 0;
      left: 60px;
      border-radius: 0 24px 24px 0;
      transform: translateX(calc(-100% - 60px));
      transition: transform 0.3s cubic-bezier(0.4, 0.0, 0.2, 1);
      padding-bottom: calc(48px + env(safe-area-inset-bottom));
      box-shadow: 8px 0 24px rgba(0, 0, 0, 0.08);
      z-index: 1005;
      box-sizing: border-box;
    }
  
  .mini-sidebar.mobile-open,
  .history-sidebar.mobile-open {
    transform: translateX(0);
    transition: transform 0.3s cubic-bezier(0.4, 0.0, 0.2, 1);
  }
  
  .sidebar-overlay {
    position: fixed;
    inset: 0;
    background: rgba(0, 0, 0, 0.5);
    z-index: 1000;
    opacity: 0;
    animation: fadeInOverlay 0.25s ease forwards;
    -webkit-tap-highlight-color: transparent;
  }
  
  @keyframes fadeInOverlay {
  to {
    opacity: 1;
  }
}


.global-account-container {
    width: 300px;
    transform: translateX(-100%);
    transition: transform 0.3s cubic-bezier(0.4, 0.0, 0.2, 1);
    z-index: 2000; /* 确保移动端下也是最高层级 */
  }

  .global-account-container.explore-mode {
    width: 60px;
  }
  
  .global-account-container.mobile-open {
    transform: translateX(0);
  }
  
  .accountTriggerAvatar {
    width: 60px;
  }

  .main-header {
    padding: 0 16px;
  }

  .mobile-menu-btn {
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 20px;
    cursor: pointer;
    color: #4b5563;
    padding: 6px;
    border-radius: 8px;
    background-color: #ffffff;
    border: 1px solid #e5e7eb;
    box-shadow: 0 1px 2px 0 rgba(0, 0, 0, 0.05);
    transition: all 0.2s ease;
    -webkit-tap-highlight-color: transparent; /* 消除移动端点击时的默认高亮（蓝色背景） */
    outline: none; /* 消除可能存在的焦点轮廓 */
  }
  
  .mobile-menu-btn:hover {
    background-color: #f9fafb;
    border-color: #d1d5db;
    color: #111827;
  }
  
  .input-tools-container {
    max-width: 100%;
    width: auto; /* column 父容器 stretch 会约束到内容区宽度；width:100% 会解析为 padding box 导致溢出 */
    padding: 0;
    box-sizing: border-box;
    margin-bottom: 0;
    min-width: 0; /* flex item 默认 min-width:auto 会被内容撑开，导致设置按钮溢出容器右边界 */
  }

  .input-container {
    max-width: 100%;
    width: 100%;
    padding: 8px 6px 6px !important; /* 手机端减小顶部和底部 padding，特别是底部 */
    border-radius: 12px;
    border: 1px solid rgba(229, 231, 235, 0.5) !important;
    box-shadow: 0 4px 20px rgba(0, 0, 0, 0.05), 0 2px 8px rgba(0, 0, 0, 0.02) !important;
  }
  
  .input-area-wrapper.is-shrunk .input-container {
    padding: 8px 12px !important;
    border-radius: 20px !important;
  }

  .input-area-wrapper {
    padding: 6px 16px 6px 16px; /* 统一外边距 */
  }

  .message-inner {
    max-width: 100%;
  }

  .header-title {
    font-size: 14px;
    max-width: 160px;
  }

  .suggestion-cards {
    flex-direction: column;
    align-items: stretch;
  }

  .desktop-only {
    display: none !important;
  }

  .mobile-only {
    display: flex !important;
  }

  /* 移动端保留花哨效果，并在展开后允许横向滑动 */
  .reference-image-preview-inline {
    position: relative;
    z-index: 10;
    display: flex;
    align-items: center;
    margin-right: 0px; /* 移动端下移除右侧间距，让它更靠近文本 */
    margin-top: 0px; /* 移动端不需要额外下移，保持原有布局 */
    flex-shrink: 0;
    max-width: calc(100vw - 120px); /* 留出输入框和发送按钮的空间 */
    overflow-x: auto;
    overflow-y: visible; /* 允许上下内容溢出，防止截断 */
    scrollbar-width: none;
    -ms-overflow-style: none;
    padding: 24px 4px; /* 增加上下内边距，防止旋转或放大的图片被截断遮挡 */
    margin: -24px -4px; /* 负外边距抵消 padding 的影响，保持原本的布局位置 */
  }

  .reference-image-preview-inline::-webkit-scrollbar {
    display: none !important;
  }

  .reference-group-hover-trigger {
    display: none;
  }

  .input-tools {
    display: flex;
    gap: 8px;
    padding: 0; 
    align-items: center;
    flex-wrap: nowrap;
    overflow-x: auto;
    overflow-y: hidden;
    scrollbar-width: none !important; /* Firefox 强制隐藏滚动条 */
    -ms-overflow-style: none; /* IE and Edge 强制隐藏滚动条 */
    -webkit-overflow-scrolling: touch;
    margin-bottom: 8px;
  }
  
  /* Webkit 浏览器 (Chrome, Safari, iOS 等) 强制隐藏滚动条 */
  .input-tools::-webkit-scrollbar {
    display: none !important;
    width: 0 !important;
    height: 0 !important;
    background: transparent !important;
  }

  .model-dropdown-wrapper {
    max-width: 100%;
    flex: 0 1 auto; /* 允许压缩：模型名过长时自身省略，保证右侧设置按钮完整可见 */
    min-width: 0;
  }

  .combined-settings-btn {
    padding: 4px 12px;
    height: 32px;
    font-size: 13px;
    max-width: 100%;
    width: auto;
    display: inline-flex;
    align-items: center;
    justify-content: center;
    box-sizing: border-box;
    flex: 0 0 auto;
    min-width: 0; /* 覆盖桌面端 min-width: 100px，避免窄屏下挤压同行其他按钮 */
    border-radius: 16px;
    background-color: #ffffff; /* 增加背景色 */
    border: 1px solid #e5e7eb; /* 增加边框 */
    box-shadow: 0 1px 2px rgba(0, 0, 0, 0.05); /* 增加轻微阴影提升立体感 */
  }

  .model-settings-btn {
    justify-content: flex-start; /* 模型名字靠左，箭头靠右（如果有足够空间） */
    min-width: 0; /* 覆盖桌面端 100px 最小宽度，允许模型名过长时收缩省略 */
  }

  .model-btn-content {
    flex: 0 1 auto;
    min-width: 0; /* 允许文本截断 */
  }

  .model-name-text {
    flex: 0 1 auto;
    min-width: 0;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
    max-width: 140px; /* 收紧手机端模型名宽度，为右侧设置按钮留出空间 */
    line-height: 1.5;
  }

  .combined-settings-btn .divider {
    margin: 0 6px;
  }

  .upload-btn-wrapper {
    flex: 0 0 auto;
  }

  .ref-btn {
    height: 32px;
    font-size: 13px;
    padding: 0 12px;
    border-radius: 16px;
  }
}

/* Points Record Dialog Styles */
:deep(.el-overlay-dialog) {
  display: flex;
  align-items: center;
  justify-content: center;
}

/* 彻底重置 Element Plus 弹窗原生样式 */
.points-record-dialog.el-dialog,
:deep(.points-record-dialog.el-dialog) {
  background: transparent !important;
  box-shadow: none !important;
  border-radius: 0 !important;
  padding: 0 !important;
  margin: 0 !important;
  border: none !important;
  --el-dialog-bg-color: transparent !important;
  --el-dialog-box-shadow: none !important;
  --el-dialog-padding-primary: 0 !important;
}

:deep(.points-record-dialog .el-dialog__header) {
  display: none !important;
  padding: 0 !important;
  margin: 0 !important;
}

:deep(.points-record-dialog .el-dialog__body) {
  padding: 0 !important;
  margin: 0 !important;
  background: transparent !important;
}

.points-container {
  background: #ffffff;
  border-radius: 24px;
  overflow: hidden;
  display: flex;
  flex-direction: column;
  height: 600px;
  max-height: 80vh;
  box-shadow: 0 12px 32px rgba(0, 0, 0, 0.1);
}

.cdkey-redemption-card {
  margin: 0 32px 24px;
}
.cdkey-redemption-card :deep(.el-input-group__append) {
  background-color: var(--theme-color, #FE2C55);
  border-color: var(--theme-color, #FE2C55);
  color: white;
  border-radius: 0 12px 12px 0;
  padding: 0;
  overflow: hidden;
}
.cdkey-redemption-card :deep(.el-input__wrapper) {
  border-radius: 12px 0 0 12px;
  box-shadow: 0 0 0 1px rgba(0,0,0,0.05) inset;
  padding: 8px 16px;
}
.cdkey-redemption-card :deep(.el-input__wrapper.is-focus) {
  box-shadow: 0 0 0 1px var(--theme-color, #FE2C55) inset;
}
.redeem-btn {
  border: none;
  background: transparent;
  color: white;
  font-weight: 500;
  padding: 0 24px;
  height: 100%;
}
.redeem-btn:hover {
  background-color: rgba(0,0,0,0.1);
  color: white;
}

.points-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 24px 32px 16px;
  border-bottom: 1px solid transparent;
}

.points-title {
  font-size: 20px;
  font-weight: 600;
  color: #111827;
}

.points-close-btn {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 28px;
  height: 28px;
  border-radius: 4px;
  cursor: pointer;
  color: #909399;
  background-color: transparent;
  transition: all 0.3s;
}

.points-close-btn:hover {
  background-color: #f5f7fa;
  color: #303133;
}

.points-balance-card {
  margin: 0 32px 24px;
  padding: 24px;
  border-radius: 12px;
  background: #f3f4f6;
  color: #111827;
  display: flex;
  justify-content: space-between;
  align-items: center;
  box-shadow: none;
}

.balance-info {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.balance-label {
  font-size: 14px;
  color: #6b7280;
  font-weight: 500;
}

.balance-value {
  font-size: 32px;
  font-weight: 700;
  color: #111827;
  line-height: 1.1;
  font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, sans-serif;
}

.balance-actions {
  display: flex;
  gap: 12px;
  align-items: center;
  margin-top: 4px;
}

.balance-recharge-btn,
.balance-redeem-btn {
  height: 28px;
  padding: 0 14px;
  border-radius: 14px;
  font-size: 13px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s ease;
  display: flex;
  align-items: center;
  justify-content: center;
  outline: none;
  white-space: nowrap;
  flex-shrink: 0;
  box-sizing: border-box;
}

.balance-recharge-btn {
  background-color: #111827;
  color: #ffffff;
  border: none;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.balance-recharge-btn:hover {
  background-color: #1f2937;
  transform: translateY(-1px);
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.15);
}

.balance-recharge-btn:active {
  transform: translateY(1px);
  box-shadow: 0 1px 2px rgba(0, 0, 0, 0.1);
}

.balance-redeem-btn {
  background-color: #ffffff;
  color: #111827;
  border: 1px solid #d1d5db;
}

.balance-redeem-btn:hover {
  background-color: #f9fafb;
  border-color: #9ca3af;
  transform: translateY(-1px);
}

.balance-redeem-btn:active {
  transform: translateY(1px);
  background-color: #f3f4f6;
}



.premium-redeem-container {
  position: relative;
  padding: 40px 32px 32px;
  display: flex;
  flex-direction: column;
  align-items: center;
  text-align: center;
}

.redeem-close-btn {
  position: absolute;
  top: 16px;
  right: 16px;
  width: 32px;
  height: 32px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #9ca3af;
  cursor: pointer;
  border-radius: 50%;
  transition: all 0.2s;
}
.redeem-close-btn:hover {
  background-color: #f3f4f6;
  color: #111827;
}

.redeem-header {
  display: flex;
  flex-direction: column;
  align-items: center;
  width: 100%;
}

.redeem-icon-wrapper {
  width: 64px;
  height: 64px;
  background: linear-gradient(135deg, #111827 0%, #374151 100%);
  color: #ffffff;
  border-radius: 16px;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-bottom: 20px;
  box-shadow: 0 8px 16px rgba(17, 24, 39, 0.15);
  transform: rotate(-3deg);
}

.redeem-title {
  font-size: 22px;
  font-weight: 600;
  color: #111827;
  margin: 0 0 8px;
}

.redeem-desc {
  font-size: 14px;
  color: #6b7280;
  margin: 0 0 28px;
}

.redeem-body {
  width: 100%;
  margin-bottom: 32px;
}

.custom-redeem-input {
  width: 100%;
  height: 52px;
  background: #f9fafb;
  border: 2px solid #e5e7eb;
  border-radius: 12px;
  padding: 0 16px;
  font-size: 16px;
  color: #111827;
  text-align: center;
  letter-spacing: 1px;
  transition: all 0.3s;
  outline: none;
  box-sizing: border-box;
}
.custom-redeem-input:focus {
  border-color: #111827;
  background: #ffffff;
  box-shadow: 0 0 0 4px rgba(17, 24, 39, 0.05);
}
.custom-redeem-input::placeholder {
  color: #9ca3af;
  letter-spacing: normal;
}

.redeem-footer {
  width: 100%;
  display: flex;
  gap: 12px;
}

.redeem-action-btn {
  flex: 1;
  height: 44px;
  border-radius: 10px;
  font-size: 15px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s;
  display: flex;
  align-items: center;
  justify-content: center;
  outline: none;
}

.redeem-action-btn.cancel {
  background: #ffffff;
  color: #374151;
  border: 1px solid #d1d5db;
}
.redeem-action-btn.cancel:hover {
  background: #f9fafb;
  border-color: #9ca3af;
}

.redeem-action-btn.confirm {
  background: #111827;
  color: #ffffff;
  border: none;
  box-shadow: 0 4px 12px rgba(17, 24, 39, 0.15);
}
.redeem-action-btn.confirm:hover:not(:disabled) {
  background: #1f2937;
  transform: translateY(-1px);
  box-shadow: 0 6px 16px rgba(17, 24, 39, 0.2);
}
.redeem-action-btn.confirm:disabled {
  background: #9ca3af;
  cursor: not-allowed;
  box-shadow: none;
}

.points-list-section {
  flex: 1;
  display: flex;
  flex-direction: column;
  padding: 0 20px 20px;
  overflow: hidden;
}

.list-section-title {
  font-size: 14px;
  font-weight: 600;
  color: #374151;
  margin-bottom: 12px;
}

.points-list-content {
  flex: 1;
  overflow-y: auto;
  padding-right: 4px;
}

.points-list-content::-webkit-scrollbar {
  width: 4px;
}
.points-list-content::-webkit-scrollbar-thumb {
  background: #e5e7eb;
  border-radius: 2px;
}
.points-list-content::-webkit-scrollbar-track {
  background: transparent;
}

.points-list-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px 0;
  border-bottom: 1px solid #f3f4f6;
}

.points-list-item:last-child {
  border-bottom: none;
}

.item-left {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.item-reason {
  font-size: 14px;
  font-weight: 500;
  color: #111827;
}

.item-reason-detail {
  font-size: 12px;
  color: #6b7280;
  margin-top: -2px;
}

.item-time {
  font-size: 11px;
  color: #9ca3af;
}

.item-right {
  font-size: 15px;
  font-weight: 600;
  font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, sans-serif;
}

.item-right.is-positive {
  color: #10b981;
}

.item-right.is-negative {
  color: #111827;
}

.points-pagination-minimal {
  margin-top: 16px;
  display: flex;
  justify-content: center;
}

:deep(.points-pagination-minimal .el-pagination) {
  --el-pagination-bg-color: transparent;
  --el-pagination-hover-color: #111827;
}
.elegant-profile-dropdown {
  padding: 12px !important;
  min-width: 280px;
  border-radius: 16px;
  border: 1px solid rgba(255, 255, 255, 0.6);
  box-shadow: 0 10px 40px -10px rgba(0, 0, 0, 0.1), 0 1px 3px rgba(0, 0, 0, 0.05), inset 0 1px 0 rgba(255, 255, 255, 0.6);
  background: rgba(255, 255, 255, 0.85);
  backdrop-filter: blur(20px);
  -webkit-backdrop-filter: blur(20px);
  box-sizing: border-box;
  z-index: 2001 !important; /* 强制下拉菜单本身也具有超高层级 */
}

.elegant-header {
  display: flex;
  align-items: center;
  gap: 16px;
  padding: 8px 8px 16px;
}

.elegant-avatar-wrap {
  position: relative;
  box-shadow: 0 4px 10px rgba(0, 0, 0, 0.08);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
}

.elegant-info {
  display: flex;
  flex-direction: column;
  justify-content: center;
}

.elegant-name {
  font-size: 16px;
  font-weight: 600;
  color: #111827;
  line-height: 1.2;
}

.elegant-email {
  font-size: 12px;
  color: #6b7280;
  margin-top: 4px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  max-width: 140px;
}

.elegant-points-card {
  display: flex;
  justify-content: space-between;
  align-items: center;
  background: linear-gradient(135deg, #fffbeb 0%, #fef3c7 50%, #fce7f3 100%);
  border: 1px solid rgba(251, 191, 36, 0.3);
  box-shadow: 0 4px 12px rgba(251, 191, 36, 0.15), inset 0 1px 0 rgba(255, 255, 255, 0.8);
  border-radius: 12px;
  padding: 12px;
  margin: 0 4px 12px;
  position: relative;
  z-index: 1;
}

.elegant-points-card::before {
  content: '';
  position: absolute;
  top: 0; left: 0; right: 0; bottom: 0;
  background: radial-gradient(circle at top left, rgba(255,255,255,0.8) 0%, transparent 60%);
  border-radius: 12px;
  z-index: 0;
  pointer-events: none;
}

.elegant-points-left, .elegant-points-right {
  position: relative;
  z-index: 2;
}

.elegant-points-left {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.elegant-points-right {
  display: flex;
  align-items: center;
}

.elegant-points-label {
  font-size: 12px;
  color: #92400e;
  font-weight: 500;
}

.elegant-points-value {
  display: flex;
  align-items: center;
  gap: 6px;
  font-size: 20px;
  font-weight: 800;
  color: #92400e;
  letter-spacing: -0.5px;
}

.elegant-points-value .el-icon {
  color: #f59e0b;
  filter: drop-shadow(0 2px 4px rgba(245, 158, 11, 0.4));
}

.elegant-recharge-btn {
  background: linear-gradient(135deg, #f59e0b 0%, #d97706 100%) !important;
  color: white !important;
  border: none !important;
  font-weight: 600 !important;
  z-index: 2;
  white-space: nowrap;
  padding: 8px 16px;
  box-shadow: 0 2px 8px rgba(217, 119, 6, 0.3) !important;
  transition: all 0.3s ease !important;
}

.elegant-recharge-btn:hover {
  transform: translateY(-1px);
  box-shadow: 0 4px 12px rgba(217, 119, 6, 0.4) !important;
  background: linear-gradient(135deg, #fbbf24 0%, #f59e0b 100%) !important;
}

.elegant-divider {
  height: 1px;
  background: linear-gradient(to right, transparent, rgba(229, 231, 235, 0.8), transparent);
  margin: 8px 0;
}

.elegant-menu-group {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.elegant-menu-item {
  display: flex;
  align-items: center;
  padding: 10px 12px;
  border-radius: 10px;
  font-size: 14px;
  font-weight: 500;
  color: #4b5563;
  background: transparent !important;
  transition: all 0.25s cubic-bezier(0.4, 0, 0.2, 1);
}

.elegant-icon {
  font-size: 18px;
  margin-right: 12px;
  color: #9ca3af;
  transition: all 0.25s cubic-bezier(0.4, 0, 0.2, 1);
}

.elegant-arrow {
  margin-left: auto;
  font-size: 14px;
  color: #d1d5db;
  transition: all 0.25s cubic-bezier(0.4, 0, 0.2, 1);
}

.elegant-menu-item:hover {
  background: rgba(243, 244, 246, 0.8) !important;
  color: #111827;
  transform: translateX(4px);
}

.elegant-menu-item:hover .elegant-icon {
  color: #3b82f6;
  transform: scale(1.1);
}

.elegant-menu-item:hover .elegant-arrow {
  transform: translateX(2px);
  color: #9ca3af;
}

.elegant-logout:hover {
  background: rgba(254, 242, 242, 0.8) !important;
  color: #ef4444;
}

.elegant-logout:hover .elegant-icon {
  color: #ef4444;
}

/* 确保 teleport 后的下拉菜单在移动端有极高的层级 */
.el-dropdown__popper.el-popper {
  z-index: 3000 !important;
}
.combined-settings-btn {
  display: inline-flex;
  align-items: center;
  gap: 4px;
  background: rgba(244, 244, 245, 0.95);
  border: 1px solid rgba(228, 228, 231, 0.7);
  border-radius: 14px;
  padding: 2px 10px;
  height: 26px;
  font-size: 12px;
  color: #3f3f46;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.25s cubic-bezier(0.4, 0, 0.2, 1);
  box-shadow: 0 1px 2px rgba(15, 23, 42, 0.03);
}

.combined-settings-btn:hover {
  background: #e4e4e7;
  color: #18181b;
  transform: translateY(-1px);
}

.combined-settings-btn:active {
  transform: translateY(0);
}

.combined-settings-btn .el-icon {
  font-size: 13px;
  color: #71717a;
  transition: color 0.25s ease;
}

.combined-settings-btn:hover .el-icon {
  color: #18181b;
}

.combined-settings-btn .btn-text {
  display: inline-flex;
  align-items: center;
  line-height: normal;
}

.combined-settings-btn .divider {
  width: 1px;
  height: 10px; /* 减小分割线高度 */
  background-color: #e4e4e7; /* 稍微减淡颜色，使其不那么刺眼 */
  margin: 0 6px; /* 减小左右间距 */
}

.combined-settings-btn .unit-text {
  font-size: 12px;
  margin-left: 2px;
  color: #a1a1aa;
  font-weight: normal;
}

/* Settings Popover */
.settings-popover {
  padding: 16px !important;
  border-radius: 12px !important;
  box-shadow: 0 10px 25px -5px rgba(0, 0, 0, 0.1), 0 8px 10px -6px rgba(0, 0, 0, 0.1) !important;
  border: 1px solid #f3f4f6 !important;
  max-width: 90vw !important;
}

.settings-panel {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.setting-item {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.setting-label {
  font-size: 13px;
  font-weight: 500;
  color: #374151;
}

.setting-options {
  display: flex;
  flex-wrap: nowrap;
  gap: 8px;
  overflow-x: auto;
  padding-bottom: 4px;
  width: 100%;
  box-sizing: border-box;
}

.setting-options::-webkit-scrollbar {
  display: none;
}

.option-btn {
  flex-shrink: 0;
  padding: 6px 12px;
  border-radius: 6px;
  background: #f3f4f6;
  color: #4b5563;
  font-size: 13px;
  cursor: pointer;
  transition: all 0.2s;
  border: 1px solid transparent;
}

.option-btn.disabled {
  opacity: 0.6;
  cursor: not-allowed;
  background-color: #f9fafb;
}

.option-btn:not(.disabled):hover {
  background: #e5e7eb;
}

.option-btn.active {
  background: #eff6ff;
  color: #3b82f6;
  border-color: #bfdbfe;
  font-weight: 500;
}

.model-dropdown-menu .el-dropdown-item.is-active-model {
  color: #3b82f6;
  font-weight: 500;
  background-color: #eff6ff;
}

.shimmer-tag {
  position: relative;
  overflow: hidden;
}

.shimmer-tag::after {
  content: "";
  position: absolute;
  top: 0;
  left: -100%;
  width: 50%;
  height: 100%;
  background: linear-gradient(to right, rgba(255, 255, 255, 0) 0%, rgba(255, 255, 255, 0.4) 50%, rgba(255, 255, 255, 0) 100%);
  transform: skewX(-20deg);
  animation: sweep-light 2.5s infinite;
}

@keyframes sweep-light {
  0% {
    left: -100%;
  }
  20% {
    left: 200%;
  }
  100% {
    left: 200%;
  }
}

/* Header Styles */
.main-header {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 64px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0 24px;
  background: transparent;
  z-index: 150;
}

.header-left {
  display: flex;
  align-items: center;
  gap: 16px;
  flex: 1;
  min-width: 0;
}

.mobile-menu-btn {
  display: flex;
  align-items: center;
  justify-content: center;
  -webkit-app-region: no-drag;
  cursor: pointer;
  padding: 6px;
  border-radius: 8px;
  background-color: #ffffff;
  border: 1px solid #e5e7eb;
  box-shadow: 0 1px 2px 0 rgba(0, 0, 0, 0.05);
  color: #4b5563;
  transition: all 0.2s ease;
}

.mobile-menu-btn:hover {
  background-color: #f9fafb;
  border-color: #d1d5db;
  color: #111827;
}

.header-new-chat-icon {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 32px;
  height: 32px;
  background-color: #ffffff;
  border-radius: 8px;
  border: 1px solid #e5e7eb;
  box-shadow: 0 1px 2px 0 rgba(0, 0, 0, 0.05);
  color: #4b5563;
  cursor: pointer;
  transition: all 0.2s ease;
}

.header-new-chat-icon:hover {
  background-color: #f9fafb;
  border-color: #d1d5db;
  color: #111827;
}

.header-new-chat-icon .el-icon {
  font-size: 16px;
}

.header-actions {
  flex-shrink: 0;
  margin-left: 16px;
  display: flex;
  align-items: center;
  gap: 12px;
}

.header-login-btn {
  padding: 6px 16px;
  background-color: #111827;
  color: #ffffff;
  border-radius: 8px;
  font-size: 13px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s ease;
}

.header-login-btn:hover {
  background-color: #374151;
}

.header-points-btn {
  display: flex;
  align-items: center;
  gap: 6px;
  padding: 6px 12px;
  background-color: #ffffff;
  border: 1px solid #e5e7eb;
  color: #374151;
  border-radius: 8px;
  font-size: 13px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s ease;
  box-shadow: 0 1px 2px 0 rgba(0, 0, 0, 0.05);
}

.header-points-btn:hover {
  background-color: #f9fafb;
  border-color: #d1d5db;
  color: #111827;
}

.header-points-btn .el-icon {
  color: #f59e0b;
  font-size: 15px;
}

.explore-content-area {
  flex: 1;
  overflow-y: auto;
  padding: 84px 32px 32px 32px; /* 顶部留出 header 的空间 */
}

.explore-container {
  max-width: 1400px;
  margin: 0 auto;
}

/* Loading and No More Data styles */
.loading-more, .no-more-data {
  width: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 30px 0;
  color: #6b7280;
  font-size: 14px;
  gap: 12px;
}

.loading-more .el-icon {
  font-size: 18px;
  color: #3b82f6;
  animation: loading-rotate 2s linear infinite;
}

@keyframes loading-rotate {
  100% {
    transform: rotate(360deg);
  }
}

.no-more-data .divider-line {
  height: 1px;
  width: 40px;
  background-color: #e5e7eb;
}

/* Category Filter Styles */
.category-filter-section {
  width: 100%;
  max-width: 1400px;
  margin: 0 auto 24px;
  display: flex;
  flex-direction: column;
  gap: 8px; /* Tighter gap */
  padding: 0 16px;
  box-sizing: border-box;
}

.main-categories-wrapper {
  width: 100%;
  overflow-x: auto;
  -webkit-overflow-scrolling: touch;
  scrollbar-width: none;
  padding-bottom: 2px;
}

.main-categories-wrapper::-webkit-scrollbar {
  display: none;
}

.main-categories {
  display: flex;
  gap: 24px;
  align-items: center;
  padding: 0;
  width: max-content;
  min-width: 100%;
}

.category-tab {
  position: relative;
  padding: 12px 4px;
  font-size: 16px;
  font-weight: 400;
  color: #4e5969; /* 经典字节灰 */
  background: transparent;
  cursor: pointer;
  transition: color 0.2s ease;
  white-space: nowrap;
  user-select: none;
  border-radius: 0;
}

.category-tab:hover:not(.active) {
  color: #1d2129; /* 经典字节深黑 */
  background: transparent;
}

.category-tab.active {
  color: #1d2129;
  font-weight: 600;
  background: transparent;
}

/* 经典底部蓝色指示线 */
.category-tab.active::after {
  content: '';
  position: absolute;
  bottom: 0;
  left: 50%;
  transform: translateX(-50%);
  width: 16px;
  height: 3px;
  border-radius: 2px;
  background-color: #165dff; /* 经典字节蓝 */
}

.category-tab:active {
  transform: none;
  opacity: 0.8;
}

.sub-categories-wrapper {
  display: grid;
  grid-template-rows: 0fr;
  opacity: 0;
  margin-top: 0;
  pointer-events: none;
  transition: grid-template-rows 0.3s cubic-bezier(0.4, 0, 0.2, 1), opacity 0.2s ease, margin-top 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.sub-categories-wrapper.is-expanded {
  grid-template-rows: 1fr;
  opacity: 1;
  margin-top: 8px;
  pointer-events: auto;
}

.sub-categories {
  width: 100%;
  overflow-x: auto;
  -webkit-overflow-scrolling: touch;
  scrollbar-width: none;
  padding: 0; /* padding 放在 inner 避免影响 grid 动画 */
  min-height: 0; /* grid transition 必备 */
  display: flex;
  align-items: center;
}

.sub-categories::-webkit-scrollbar {
  display: none;
}

.sub-categories-inner {
  display: inline-flex;
  gap: 12px;
  align-items: center;
  padding: 8px 8px 4px;
}

.sub-category-tag {
  padding: 6px 16px;
  border-radius: 999px;
  font-size: 14px;
  font-weight: 400;
  color: #4e5969;
  background: #f2f3f5;
  border: 1px solid transparent;
  cursor: pointer;
  transition: all 0.2s ease;
  white-space: nowrap;
  user-select: none;
}

.sub-category-tag:hover {
  color: #1d2129;
  background: #e5e6eb;
}

.sub-category-tag.active {
  color: #165dff;
  background: #e8f3ff;
  border-color: transparent;
  font-weight: 500;
}

.sub-category-tag:active {
  transform: scale(0.96);
}



/* Prompts Gallery 风格瀑布流/网格 */
.explore-masonry-flex {
  display: flex;
  flex-direction: row;
  width: 100%;
  gap: 20px;
  min-height: 50vh;
}

.masonry-column {
  display: flex;
  flex-direction: column;
  flex: 1;
  gap: 20px;
  min-width: 0; /* 防止内容撑破 flex 容器 */
}

@media (max-width: 800px) {
  .explore-masonry-flex {
    gap: 8px;
    padding: 0;
  }
  .masonry-column {
    gap: 8px;
  }
}

/* 提示词卡片样式 (类似 lexica.art / promptsref) */
.prompt-card {
  break-inside: avoid;
  margin-bottom: 0; /* 改用 gap 控制间距 */
  position: relative;
  border-radius: 12px;
  overflow: hidden;
  cursor: pointer;
  background-color: #ffffff;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.04);
  border: 1px solid rgba(0, 0, 0, 0.05);
  transition: transform 0.2s ease, box-shadow 0.2s ease;
  display: block;
  transform: translateZ(0); /* 开启硬件加速，防止卡片出现闪烁 */
}

.prompt-card:hover {
  transform: translateY(-4px);
  box-shadow: 0 12px 24px rgba(0, 0, 0, 0.08);
}

.card-image {
  width: 100%;
  display: block; /* 移除底部空白 */
  height: auto; /* 允许高度自适应，形成真实的瀑布流 */
  transition: transform 0.5s cubic-bezier(0.4, 0, 0.2, 1);
}

/* 让占位符的高度自适应撑开（使用一个大概的比例），当真实图片加载后占位符消失，图片自然撑开真实高度 */
.card-image :deep(.el-image__placeholder),
.card-image :deep(.el-image__error) {
  position: relative;
  width: 100%;
  height: auto !important;
  aspect-ratio: 3 / 4; /* 默认使用 3:4 比例撑开 */
}

.prompt-card:hover .card-image {
  transform: scale(1.08);
}

.explore-no-image, .image-placeholder, .image-error {
  width: 100%;
  aspect-ratio: 3 / 4; /* 为没有加载出图片的状态保留一个基础高度 */
  display: flex;
  align-items: center;
  justify-content: center;
  color: #9ca3af;
  font-size: 32px;
  background-color: #f3f4f6; /* 添加背景色，让占位符更清晰 */
}

/* 使用骨架屏闪烁效果替代原本的 Loading 图标，避免图标被遮挡切断 */
.shimmer-placeholder {
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, #f3f4f6 25%, #e5e7eb 37%, #f3f4f6 63%);
  background-size: 400% 100%;
  animation: el-skeleton-loading 1.4s ease infinite;
}

@keyframes el-skeleton-loading {
  0% {
    background-position: 100% 50%;
  }
  100% {
    background-position: 0 50%;
  }
}

/* 如果 .is-loading 在占位符内部没有正常旋转，确保其样式正确 */
.image-placeholder .is-loading {
  animation: loading-rotate 2s linear infinite;
  color: #3b82f6;
}

/* 悬浮遮罩和内容 */
.card-overlay {
  position: absolute;
  inset: 0;
  background: linear-gradient(to top, rgba(17, 24, 39, 0.9) 0%, rgba(17, 24, 39, 0.4) 40%, transparent 100%);
  opacity: 0;
  transition: opacity 0.3s ease;
  display: flex;
  flex-direction: column;
  justify-content: flex-end;
  padding: 20px;
  /* 确保覆盖层不会因为内部元素的动画而出现裁剪或闪烁 */
  will-change: opacity;
}

.prompt-card:hover .card-overlay {
  opacity: 1;
}

.overlay-content {
  display: flex;
  flex-direction: column;
  gap: 12px;
  /* 防止内部绝对定位元素或者动画导致高度坍塌或溢出 */
  position: relative;
  z-index: 2;
}

.card-prompt-text {
  color: #ffffff;
  font-size: 14px;
  line-height: 1.5;
  margin: 0;
  display: -webkit-box;
  -webkit-line-clamp: 3;
  -webkit-box-orient: vertical;
  overflow: hidden;
  text-overflow: ellipsis;
  text-shadow: 0 2px 4px rgba(0,0,0,0.5);
  transform: translateY(10px);
  transition: transform 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  /* 添加此属性解决部分浏览器动画渲染bug */
  will-change: transform;
  /* 确保文本在动画过程中不会改变其布局属性 */
  position: relative;
}

.prompt-card:hover .card-prompt-text {
  transform: translateY(0);
}

.overlay-actions {
  display: flex;
  justify-content: flex-end;
  transform: translateY(10px);
  transition: transform 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  transition-delay: 0.05s;
  /* 添加此属性解决部分浏览器动画渲染bug */
  will-change: transform;
  /* 确保按钮在动画过程中不会改变其布局属性 */
  position: relative;
  /* 防止与文本重叠，确保至少有间距 */
  margin-top: auto;
}

.prompt-card:hover .overlay-actions {
  transform: translateY(0);
}

.try-btn {
  background: rgba(255, 255, 255, 0.2) !important;
  color: #ffffff !important;
  border: 1px solid rgba(255, 255, 255, 0.4) !important;
  backdrop-filter: blur(8px);
  font-weight: 500;
  padding: 8px 20px;
  transition: all 0.2s ease;
}

.try-btn:hover {
  background: #ffffff !important;
  color: #111827 !important;
  transform: scale(1.05);
}

.explore-custom-input :deep(.el-textarea__inner) {
  background: transparent !important;
  border: none !important;
  box-shadow: none !important;
  padding: 0 !important;
  font-size: 16px;
  line-height: 1.6;
  color: #374151;
}

.explore-custom-input :deep(.el-textarea__inner:focus) {
  box-shadow: none !important;
}

@media (min-width: 769px) {
  .mobile-menu-btn {
    display: none;
  }
  .mobile-only {
    display: none !important;
  }
}

@media (max-width: 768px) {
  .desktop-only {
    display: none !important;
  }
  .main-header {
    padding: 0 16px;
  }
  .explore-content-area {
    padding: 84px 0 16px 0;
    overflow-x: hidden;
  }
  
  .category-filter-section {
    padding: 0 12px;
    margin-bottom: 12px;
    gap: 4px;
  }
  
  .main-categories-wrapper {
    padding-bottom: 0;
  }
  
  .main-categories {
    gap: 16px;
    background: transparent;
    border-radius: 0;
    backdrop-filter: none;
    border: none;
    box-shadow: none;
    padding: 0 8px;
  }
  
  .category-tab {
    padding: 10px 2px;
    font-size: 15px;
    border-radius: 0;
    background: transparent;
    color: #4e5969;
  }
  
  .category-tab.active {
    color: #1d2129;
    background: transparent;
    font-weight: 500;
  }
  
  .category-tab .active-bg {
    display: none;
  }

  .sub-categories-inner {
    padding: 4px 8px;
    background: transparent;
    border-radius: 0;
    backdrop-filter: none;
    border: none;
    box-shadow: none;
    gap: 8px;
  }
  
  .sub-category-tag {
    padding: 5px 12px;
    font-size: 13px;
    border-radius: 999px;
    background: #f2f3f5;
    border: 1px solid transparent;
  }
  
  .sub-category-tag.active {
    background: #e8f3ff;
    border-color: transparent;
    color: #165dff;
  }
  
  .explore-masonry-flex {
    padding: 0; /* 给瀑布流整体加一点边距，避免贴边太死 */
  }
  
  .prompt-card {
    margin-bottom: 0; /* flex gap 控制间距，取消 margin */
    border-radius: 10px; /* 移动端圆角再小一点 */
  }
  
  .card-overlay {
    padding: 10px;
  }
  
  .card-prompt-text {
    font-size: 11px;
    -webkit-line-clamp: 2; /* 手机端显示两行即可 */
  }
  
  .try-btn {
    padding: 4px 10px;
    font-size: 11px;
  }
  
  .mobile-menu-btn {
    font-size: 20px;
    -webkit-tap-highlight-color: transparent;
    outline: none;
  }
}

/* 高级版灵感详情弹窗 CSS */
:deep(.advanced-inspiration-dialog) {
  border-radius: 24px !important;
  overflow: hidden;
  box-shadow: 0 25px 50px -12px rgba(0, 0, 0, 0.25) !important;
  background-color: #ffffff !important;
  padding: 0 !important;
  display: flex;
  flex-direction: column;
}

:deep(.advanced-inspiration-dialog .el-dialog__header),
:deep(.advanced-inspiration-dialog .el-dialog__body),
:deep(.advanced-inspiration-dialog .el-dialog__footer) {
  padding: 0 !important;
  margin: 0 !important;
  display: none; /* 隐藏默认结构 */
}

/* 强制显示自定义body，因为它被包裹在 .el-dialog__body 里 */
:deep(.advanced-inspiration-dialog .el-dialog__body) {
  display: block !important;
}

.advanced-split-layout {
  display: flex;
  flex-direction: row;
  height: 600px;
  width: 100%;
}

.advanced-left-image {
  flex: 1.5; /* 左侧占更大比例 */
  background-color: #0f172a; /* 深夜蓝/黑色，让图片显得更高级 */
  position: relative;
  display: flex;
  justify-content: center;
  align-items: center;
  overflow: hidden;
}

.advanced-blur-bg {
  position: absolute;
  inset: -40px; /* 扩展边缘防止模糊露底 */
  background-size: cover;
  background-position: center;
  filter: blur(30px) brightness(0.4) saturate(1.2);
  z-index: 1;
  transform: scale(1.1); /* 放大一点避免边缘发白 */
  transition: all 0.3s ease;
}

.advanced-main-image {
  position: relative;
  z-index: 2;
  width: 100%;
  height: 100%;
  padding: 32px;
  box-sizing: border-box;
}

:deep(.advanced-main-image .el-image__inner) {
  object-fit: contain !important;
  border-radius: 8px;
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.4);
  transition: transform 0.4s cubic-bezier(0.4, 0, 0.2, 1);
}

:deep(.advanced-main-image:hover .el-image__inner) {
  transform: scale(1.03);
}

.advanced-image-skeleton {
  position: relative;
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  gap: 16px;
  background: linear-gradient(135deg, rgba(15, 23, 42, 0.8) 0%, rgba(30, 41, 59, 0.8) 100%);
  border-radius: 12px;
  overflow: hidden;
}

.skeleton-pulse-ring {
  position: absolute;
  width: 80px;
  height: 80px;
  border-radius: 50%;
  background: radial-gradient(circle, rgba(59, 130, 246, 0.4) 0%, rgba(59, 130, 246, 0) 70%);
  animation: pulse-ring 2s cubic-bezier(0.4, 0, 0.6, 1) infinite;
}

.skeleton-pulse-ring.delay {
  animation-delay: 1s;
}

.skeleton-glow-icon {
  font-size: 48px;
  color: rgba(255, 255, 255, 0.9);
  filter: drop-shadow(0 0 12px rgba(59, 130, 246, 0.8));
  z-index: 2;
  animation: float-icon 3s ease-in-out infinite;
}

.skeleton-text {
  z-index: 2;
  font-size: 14px;
  font-weight: 500;
  letter-spacing: 2px;
  background: linear-gradient(90deg, rgba(255,255,255,0.4) 0%, rgba(255,255,255,1) 50%, rgba(255,255,255,0.4) 100%);
  background-size: 200% auto;
  color: transparent;
  -webkit-background-clip: text;
  background-clip: text;
  animation: text-shine 2s linear infinite;
}

@keyframes pulse-ring {
  0% {
    transform: scale(0.5);
    opacity: 1;
  }
  100% {
    transform: scale(2.5);
    opacity: 0;
  }
}

@keyframes float-icon {
  0%, 100% { transform: translateY(0); }
  50% { transform: translateY(-10px); }
}

@keyframes text-shine {
  to {
    background-position: 200% center;
  }
}

.advanced-right-content {
  flex: 1;
  display: flex;
  flex-direction: column;
  background-color: #ffffff;
  min-width: 340px; /* 保证右侧不要太窄 */
}

.advanced-split-layout:not(.has-image) .advanced-right-content {
  flex: 1;
}

.advanced-right-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 24px 24px 16px;
}

.advanced-title {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 18px;
  font-weight: 700;
  color: #111827;
}

.advanced-title .el-icon {
  color: #111827;
  font-size: 20px;
}

.desktop-close-btn {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 32px;
  height: 32px;
  border-radius: 50%;
  background-color: #f3f4f6;
  color: #6b7280;
  cursor: pointer;
  transition: all 0.2s;
}

.desktop-close-btn:hover {
  background-color: #e5e7eb;
  color: #111827;
  transform: rotate(90deg);
}

.mobile-close-btn {
  display: none; /* PC端隐藏 */
  position: absolute;
  top: 16px;
  right: 16px;
  z-index: 10;
  width: 32px;
  height: 32px;
  border-radius: 50%;
  background-color: rgba(255, 255, 255, 0.2);
  color: #ffffff;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  backdrop-filter: blur(4px);
  transition: all 0.2s;
}

.mobile-close-btn:hover {
  background-color: rgba(255, 255, 255, 0.4);
}

.advanced-right-body {
  flex: 1;
  padding: 0 24px;
  overflow-y: auto;
}

/* 优化 PC 端详情页内容的滚动条 */
.advanced-right-body::-webkit-scrollbar {
  width: 6px;
}

.advanced-right-body::-webkit-scrollbar-thumb {
  background-color: #d1d5db;
  border-radius: 3px;
}

.advanced-right-body::-webkit-scrollbar-thumb:hover {
  background-color: #9ca3af;
}

.advanced-right-body::-webkit-scrollbar-track {
  background: transparent;
}

.advanced-prompt-container {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.advanced-prompt-label {
  font-size: 12px;
  font-weight: 700;
  color: #9ca3af;
  letter-spacing: 1px;
}

.advanced-prompt-text {
  font-size: 15px;
  line-height: 1.7;
  color: #1f2937;
  background-color: #f9fafb;
  padding: 16px;
  border-radius: 12px;
  white-space: pre-wrap;
  word-break: break-word;
  border: 1px solid #f3f4f6;
}

.advanced-right-footer {
  padding: 20px 24px 24px;
  display: flex;
  gap: 12px;
}

.advanced-btn {
  height: 48px !important;
  border-radius: 24px !important;
  font-size: 15px !important;
  font-weight: 600 !important;
  flex: 1;
  display: flex !important;
  justify-content: center !important;
  align-items: center !important;
  gap: 6px !important;
  border: none !important;
  transition: all 0.2s ease !important;
}

.advanced-btn.copy-btn {
  background-color: #f3f4f6 !important;
  color: #374151 !important;
}

.advanced-btn.copy-btn:hover {
  background-color: #e5e7eb !important;
  color: #111827 !important;
}

.advanced-btn.use-btn {
  background-color: #111827 !important;
  color: #ffffff !important;
}

.advanced-btn.use-btn:hover {
  background-color: #374151 !important;
  transform: translateY(-2px);
  box-shadow: 0 8px 16px rgba(0, 0, 0, 0.15) !important;
}

.advanced-btn.use-btn:active {
  transform: translateY(0);
}

@media (max-width: 768px) {
  .advanced-split-layout {
    flex-direction: column;
    height: auto;
    max-height: 85vh;
  }
  
  .advanced-left-image {
    flex: none;
    height: 320px; /* 移动端给一个固定高度 */
    border-right: none;
  }
  
  .advanced-main-image {
    padding: 16px;
  }
  
  .desktop-close-btn {
    display: none;
  }
  
  .mobile-close-btn {
    display: flex;
  }
  
  .advanced-right-content {
    flex: none;
    min-width: auto;
    display: flex;
    flex-direction: column;
    height: calc(85vh - 320px); /* 限制右侧内容区的总高度为弹窗剩余高度 */
  }
  
  .advanced-right-body {
    padding-top: 16px;
    flex: 1; /* 让主体部分占据剩余空间 */
    overflow-y: auto; /* 允许滚动 */
    max-height: none; /* 移除之前的 max-height 限制，改用 flex: 1 控制 */
  }
  
  .advanced-right-footer {
    padding: 16px 24px 24px;
    flex-shrink: 0; /* 防止底部按钮区被压缩 */
    background-color: #ffffff; /* 确保背景色为白色，防止内容透出 */
    position: sticky; /* 固定在底部 */
    bottom: 0;
    z-index: 10;
  }
}
</style>
<style scoped>
@media (max-width: 768px) {
  :deep(.chat-input) {
    min-height: 36px !important;
  }
  
  :deep(.chat-input .el-textarea__inner) {
    min-height: 36px !important;
    padding-top: 8px !important;
    padding-bottom: 8px !important;
    font-size: 14px !important;
    line-height: 20px !important;
    display: block !important;
    max-height: 200px !important;
    overflow-y: auto !important;
    box-sizing: border-box !important;
  }
  
  /* 移动端滚动条样式 */
  :deep(.chat-input .el-textarea__inner::-webkit-scrollbar) {
    width: 4px;
  }
  :deep(.chat-input .el-textarea__inner::-webkit-scrollbar-thumb) {
    background-color: rgba(156, 163, 175, 0.5);
    border-radius: 4px;
  }
  :deep(.chat-input .el-textarea__inner::-webkit-scrollbar-track) {
    background: transparent;
  }
  
  .input-area-wrapper.is-shrunk :deep(.chat-input) {
    min-height: 36px !important;
    height: 36px !important;
  }
  
  .input-area-wrapper.is-shrunk .input-box {
    min-height: 36px !important;
    height: 36px !important;
  }
  
  .input-area-wrapper.is-shrunk :deep(.chat-input.has-references) {
    min-height: 36px !important;
    height: 36px !important;
  }
  
  .input-area-wrapper.is-shrunk .input-box.has-refs {
    min-height: 36px !important;
    height: 36px !important;
  }
  
  .input-area-wrapper.is-shrunk :deep(.chat-input.has-references .el-textarea__inner) {
    min-height: 36px !important;
    height: 36px !important;
    padding-top: 8px !important;
    padding-bottom: 8px !important;
    max-height: 36px !important;
    margin-left: -8px !important;
  }
  
  .input-area-wrapper.is-shrunk :deep(.chat-input .el-textarea__inner) {
    min-height: 36px !important;
    height: 36px !important;
    padding-top: 8px !important;
    padding-bottom: 8px !important;
    font-size: 13px !important;
    line-height: 20px !important;
    display: block !important;
    max-height: 36px !important;
    overflow: hidden !important;
    white-space: nowrap !important;
    box-sizing: border-box !important;
    margin-left: -8px !important; /* 手机端输入框向左移动更多，靠近图片 */
  }
  
  .input-area-wrapper.is-shrunk :deep(.chat-input .el-textarea__inner::placeholder) {
    white-space: nowrap !important;
    overflow: hidden !important;
    text-overflow: ellipsis !important;
    line-height: 20px !important;
  }
  
  :deep(.chat-input .el-textarea__inner::placeholder) {
    line-height: 20px !important;
  }

  .input-box {
    min-height: 36px !important;
    padding: 0px 0px !important;
  }
  
  .input-area-wrapper.is-shrunk .input-box {
    min-height: 28px !important;
    padding: 0px 0px !important;
  }
}
</style>
