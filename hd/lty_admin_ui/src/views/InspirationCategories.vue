<template>
  <div class="settings-container fade-in">
    <div class="page-header">
      <div class="header-info">
        <div class="icon-wrapper">
          <el-icon :size="24" color="#409eff"><Folder /></el-icon>
        </div>
        <div class="header-title">
          <h2>探索分类设置</h2>
          <span class="subtitle">配置探索管理中的主分类（一级）和副分类（二级）。</span>
        </div>
      </div>
    </div>

    <el-card shadow="never" class="settings-card">
      <div class="card-header-actions">
        <el-button type="primary" @click="openAddDialog" class="add-main-btn">
          <el-icon><Plus /></el-icon> 添加主分类
        </el-button>
      </div>

      <div v-loading="loading" class="categories-editor">
        <el-empty v-if="inspirationCategories.length === 0" description="暂无分类数据，请点击上方按钮添加" />
        
        <draggable
          v-else
          v-model="inspirationCategories"
          class="categories-grid"
          item-key="name"
          handle=".drag-icon"
          @end="saveCategoriesOrder"
          :animation="200"
        >
          <template #item="{ element: cat, index }">
            <div class="category-item-card">
              <div class="category-header">
                <div class="header-left">
                  <el-icon class="drag-icon"><Grid /></el-icon>
                  <span class="category-name-display">{{ cat.name }}</span>
                </div>
                <div class="header-actions" style="display: flex; flex-wrap: wrap; gap: 8px;">
                  <el-button type="primary" link @click="openEditDialog(index)" class="edit-btn" style="margin-left: 0;">
                    <el-icon><Edit /></el-icon> 编辑
                  </el-button>
                  <el-popconfirm
                    title="确定要删除该主分类吗？"
                    @confirm="removeCategory(index)"
                    width="200"
                  >
                    <template #reference>
                      <el-button type="danger" link class="delete-btn" style="margin-left: 0;">
                        <el-icon><Delete /></el-icon> 删除
                      </el-button>
                    </template>
                  </el-popconfirm>
                </div>
              </div>
              
              <div class="category-divider"></div>
              
              <div class="category-body">
                <div class="sub-category-title">副分类标签 ({{ cat.sub.length }})</div>
                <draggable
                  v-model="cat.sub"
                  class="sub-category-list"
                  :item-key="getPrimitiveKey"
                  @end="saveCategoriesOrder"
                  :animation="200"
                >
                  <template #item="{ element: sub }">
                    <el-tag
                      effect="light"
                      class="sub-tag"
                      type="info"
                      style="cursor: grab;"
                    >
                      {{ sub }}
                    </el-tag>
                  </template>
                  <template #footer>
                    <span v-if="cat.sub.length === 0" class="no-sub-text">暂无副分类</span>
                  </template>
                </draggable>
              </div>
            </div>
          </template>
        </draggable>
      </div>
    </el-card>

    <!-- 添加/编辑分类弹窗 -->
    <el-dialog append-to-body 
      v-model="dialogVisible"
      :title="isEditing ? '编辑分类' : '添加主分类'"
      width="500px"
      destroy-on-close
      class="category-dialog"
      :before-close="handleDialogClose"
    >
      <el-form label-width="80px" @submit.prevent>
        <el-form-item label="主分类名" required>
          <el-input v-model="currentEditCategory.name" placeholder="请输入主分类名称" clearable />
        </el-form-item>
        <el-form-item label="副分类">
          <draggable
            v-model="currentEditCategory.sub"
            class="dialog-sub-list"
            :item-key="getPrimitiveKey"
            :animation="200"
          >
            <template #item="{ element: sub, index: subIndex }">
              <el-tag
                closable
                effect="light"
                @close="removeDialogSub(subIndex)"
                class="sub-tag"
                type="info"
                style="cursor: grab;"
              >
                {{ sub }}
              </el-tag>
            </template>
            <template #footer>
              <div class="input-wrapper" style="display: inline-block;">
                <el-input
                  v-if="inputVisible"
                  ref="InputRef"
                  v-model="inputValue"
                  class="sub-input"
                  size="small"
                  placeholder="输入后回车"
                  @keyup.enter="handleInputConfirm"
                  @blur="handleInputConfirm"
                />
                <el-button v-else class="button-new-tag" size="small" type="primary" plain @click="showInput">
                  <el-icon><Plus /></el-icon> 添加副分类
                </el-button>
              </div>
            </template>
          </draggable>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="handleCancel">取消</el-button>
          <el-button type="primary" @click="saveDialogCategory" :loading="saving">
            保存
          </el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onBeforeUnmount, nextTick } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { Folder, Delete, Plus, Edit, Grid } from '@element-plus/icons-vue'
import draggable from 'vuedraggable'

const loading = ref(false)
const saving = ref(false)

interface CategoryItem {
  name: string
  sub: string[]
}

const inspirationCategories = ref<CategoryItem[]>([])

// 弹窗状态
const dialogVisible = ref(false)
const isEditing = ref(false)
const editIndex = ref(-1)
const currentEditCategory = ref<CategoryItem>({ name: '', sub: [] })
const originalCategory = ref<CategoryItem>({ name: '', sub: [] })

// 弹窗内副分类输入状态
const inputVisible = ref(false)
const inputValue = ref('')
const InputRef = ref<any>(null)

const openAddDialog = () => {
  isEditing.value = false
  editIndex.value = -1
  currentEditCategory.value = { name: '', sub: [] }
  originalCategory.value = { name: '', sub: [] }
  dialogVisible.value = true
}

const openEditDialog = (index: number) => {
  isEditing.value = true
  editIndex.value = index
  const cat = JSON.parse(JSON.stringify(inspirationCategories.value[index]))
  currentEditCategory.value = cat
  originalCategory.value = JSON.parse(JSON.stringify(cat))
  dialogVisible.value = true
}

const checkUnsavedChanges = async () => {
  const currentStr = JSON.stringify(currentEditCategory.value)
  const originalStr = JSON.stringify(originalCategory.value)
  
  if (currentStr !== originalStr) {
    try {
      await ElMessageBox.confirm('您有未保存的修改，确定要关闭吗？', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
      })
      return true
    } catch {
      return false
    }
  }
  return true
}

const handleDialogClose = async (done: () => void) => {
  const canClose = await checkUnsavedChanges()
  if (canClose) {
    done()
  }
}

const handleCancel = async () => {
  const canClose = await checkUnsavedChanges()
  if (canClose) {
    dialogVisible.value = false
  }
}

const removeCategory = async (index: number) => {
  const newCategories = JSON.parse(JSON.stringify(inspirationCategories.value))
  newCategories.splice(index, 1)
  
  try {
    await saveCategories(newCategories)
    inspirationCategories.value = newCategories
    ElMessage.success('删除成功')
  } catch (error) {
    ElMessage.error('删除失败')
  }
}

const saveCategoriesOrder = async () => {
  try {
    await saveCategories(inspirationCategories.value)
    ElMessage.success('顺序已保存')
  } catch (error) {
    ElMessage.error('保存顺序失败')
  }
}

const removeDialogSub = (index: number) => {
  currentEditCategory.value.sub.splice(index, 1)
}

const getPrimitiveKey = (item: any) => item

const showInput = () => {
  inputVisible.value = true
  nextTick(() => {
    if (InputRef.value) {
      InputRef.value.focus()
    }
  })
}

const handleInputConfirm = () => {
  if (inputValue.value) {
    const val = inputValue.value.trim()
    if (val && !currentEditCategory.value.sub.includes(val)) {
      currentEditCategory.value.sub.push(val)
    }
  }
  inputVisible.value = false
  inputValue.value = ''
}

const getHeaders = () => {
  const token = localStorage.getItem('token')
  return {
    'Authorization': `Bearer ${token}`
  }
}

const loadCategoriesData = async () => {
  try {
    const res = await fetch(`${(window as any).APP_CONFIG?.API_BASE_URL || ''}/api/admin/inspiration-categories/list`, {
      headers: { ...getHeaders(), 'Content-Type': 'application/json' }
    })
    if (res.ok) {
      const json = await res.json()
      return json.data || []
    }
  } catch (error) {
    console.error(`Failed to fetch categories`, error)
  }
  return []
}

const saveCategories = async (categories: CategoryItem[]) => {
  try {
    const res = await fetch(`${(window as any).APP_CONFIG?.API_BASE_URL || ''}/api/admin/inspiration-categories/sync`, {
      method: 'POST',
      headers: { ...getHeaders(), 'Content-Type': 'application/json' },
      body: JSON.stringify(categories)
    })
    if (!res.ok) {
      throw new Error(`Failed to save categories`)
    }
  } catch (error) {
    throw error
  }
}

const saveDialogCategory = async () => {
  if (!currentEditCategory.value.name.trim()) {
    ElMessage.warning('主分类名称不能为空')
    return
  }

  saving.value = true
  try {
    const newCategories = JSON.parse(JSON.stringify(inspirationCategories.value))
    if (isEditing.value) {
      newCategories[editIndex.value] = {
        name: currentEditCategory.value.name.trim(),
        sub: currentEditCategory.value.sub
      }
    } else {
      newCategories.push({
        name: currentEditCategory.value.name.trim(),
        sub: currentEditCategory.value.sub
      })
    }
    
    await saveCategories(newCategories)
    inspirationCategories.value = newCategories
    dialogVisible.value = false
    ElMessage.success('保存成功')
  } catch (error) {
    ElMessage.error('保存失败')
  } finally {
    saving.value = false
  }
}

const loadSettings = async () => {
  loading.value = true
  try {
    const cats = await loadCategoriesData()
    if (cats && Array.isArray(cats)) {
      inspirationCategories.value = cats.map((c: any) => ({
        name: c.name,
        sub: c.sub || []
      }))
    } else {
      inspirationCategories.value = []
    }
  } catch (error) {
    ElMessage.error('获取配置失败')
  } finally {
    loading.value = false
  }
}

const handleBeforeUnload = (e: BeforeUnloadEvent) => {
  if (dialogVisible.value) {
    const currentStr = JSON.stringify(currentEditCategory.value)
    const originalStr = JSON.stringify(originalCategory.value)
    if (currentStr !== originalStr) {
      e.preventDefault()
      e.returnValue = ''
    }
  }
}

onMounted(() => {
  loadSettings()
  window.addEventListener('beforeunload', handleBeforeUnload)
})

onBeforeUnmount(() => {
  window.removeEventListener('beforeunload', handleBeforeUnload)
})
</script>

<style scoped>
.fade-in {
  animation: fadeIn 0.4s ease-in-out;
}

@keyframes fadeIn {
  from { opacity: 0; transform: translateY(10px); }
  to { opacity: 1; transform: translateY(0); }
}

.settings-container {
  padding-top: 8px;
}

.page-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24px;
}

.header-info {
  display: flex;
  align-items: center;
  gap: 16px;
}

.icon-wrapper {
  width: 48px;
  height: 48px;
  background: linear-gradient(135deg, rgba(64, 158, 255, 0.15) 0%, rgba(41, 121, 255, 0.05) 100%);
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #409eff;
}

.header-title h2 {
  margin: 0 0 4px 0;
  font-size: 24px;
  color: #1d1d1f;
  font-weight: 700;
  letter-spacing: -0.5px;
}

.subtitle {
  font-size: 14px;
  color: #86868b;
}

.settings-card {
  border: 1px solid rgba(255, 255, 255, 0.6) !important;
  box-shadow: 0 10px 40px rgba(0, 0, 0, 0.04), 0 1px 3px rgba(0, 0, 0, 0.02) !important;
  border-radius: 24px !important;
  background: rgba(255, 255, 255, 0.75) !important;
  backdrop-filter: blur(20px);
  -webkit-backdrop-filter: blur(20px);
  min-height: 500px;
  padding: 0;
}

.card-header-actions {
  display: flex;
  justify-content: flex-end;
  align-items: center;
  gap: 12px;
  padding: 20px 24px;
  border-bottom: 1px solid rgba(229, 229, 234, 0.5);
}

.categories-editor {
  padding: 24px;
}

.categories-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(320px, 1fr));
  gap: 20px;
}

.category-item-card {
  border: 1px solid rgba(229, 229, 234, 0.8);
  border-radius: 16px;
  background: rgba(255, 255, 255, 0.6);
  transition: all 0.3s ease;
  overflow: hidden;
  backdrop-filter: blur(10px);
}

.category-item-card:hover {
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.08);
  border-color: rgba(229, 229, 234, 1);
  transform: translateY(-2px);
}

.category-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 16px;
  background-color: rgba(245, 245, 247, 0.5);
}

.header-left {
  display: flex;
  align-items: center;
  gap: 8px;
  flex: 1;
}

.drag-icon {
  color: #86868b;
  cursor: grab;
  font-size: 18px;
}

.category-name-display {
  font-size: 16px;
  font-weight: 600;
  color: #1d1d1f;
}

.header-actions {
  display: flex;
  gap: 8px;
}

.edit-btn, .delete-btn {
  font-size: 14px;
  padding: 4px 8px;
}

.category-divider {
  height: 1px;
  background-color: rgba(229, 229, 234, 0.5);
}

.category-body {
  padding: 16px;
}

.sub-category-title {
  font-size: 13px;
  color: #86868b;
  margin-bottom: 12px;
  font-weight: 500;
}

.sub-category-list {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
  align-items: center;
  min-height: 32px;
}

.sub-tag {
  border-radius: 8px;
  padding: 0 12px;
  height: 30px;
  line-height: 28px;
  background: rgba(245, 245, 247, 0.8);
  border: 1px solid rgba(229, 229, 234, 0.8);
  color: #434344;
}

.no-sub-text {
  font-size: 13px;
  color: #86868b;
  font-style: italic;
}

.dialog-sub-list {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
  align-items: center;
  min-height: 32px;
}

.sub-input {
  width: 120px;
}

.button-new-tag {
  border-radius: 8px;
  border-style: dashed;
}

.save-btn, .add-main-btn {
  border-radius: 12px;
  font-weight: 600;
}

/* ================= 移动端响应式 ================= */
@media (max-width: 768px) {
  .page-header {
    flex-direction: column;
    align-items: flex-start;
    gap: 16px;
  }
  
  .card-header-actions {
    justify-content: flex-start;
  }
  
  .add-main-btn {
    width: 100%;
  }

  .categories-grid {
    grid-template-columns: 1fr;
  }
  
  .category-header {
    flex-direction: column;
    align-items: flex-start;
    gap: 12px;
  }
  
  .header-actions {
    width: 100%;
    justify-content: flex-end;
  }
}
</style>