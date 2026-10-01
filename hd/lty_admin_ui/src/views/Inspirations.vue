<template>
  <div class="inspirations-container">
    <div class="page-header">
      <h2>探索管理</h2>
      <el-button type="primary" @click="showAddDialog">
        <el-icon><Plus /></el-icon> 添加提示词
      </el-button>
    </div>

    <el-card class="box-card" shadow="never">
      <div class="toolbar">
        <div class="search-bar">
          <el-select 
            v-model="searchMainCategory" 
            placeholder="主分类" 
            clearable 
            @change="handleFilterMainCategoryChange" 
            class="filter-select"
          >
            <el-option v-for="item in mainCategories" :key="item" :label="item" :value="item" />
          </el-select>
          <el-select 
            v-model="searchSubCategory" 
            placeholder="副分类" 
            clearable 
            @change="handleFilterSubCategoryChange" 
            class="filter-select"
          >
            <el-option v-for="item in filterSubCategories" :key="item" :label="item" :value="item" />
          </el-select>
          <el-input
            v-model="searchQuery"
            placeholder="搜索提示词内容"
            clearable
            @clear="handleSearch"
            @keyup.enter="handleSearch"
            class="search-input"
          >
            <template #prefix>
              <el-icon><Search /></el-icon>
            </template>
          </el-input>
          
          <el-button type="primary" @click="handleSearch">搜索</el-button>
          <el-button type="danger" :disabled="selectedIds.length === 0" @click="handleBatchDelete" :loading="batchDeleting">批量删除</el-button>
        </div>
      </div>

      <el-table :data="inspirations" style="width: 100%" v-loading="loading" @selection-change="handleSelectionChange">
        <el-table-column type="selection" width="55" />
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column label="示例图片" width="120">
          <template #default="{ row }">
            <el-image 
              v-if="row.image_url"
              :src="row.image_url" 
              :preview-src-list="[row.image_url]"
              fit="cover" 
              style="width: 60px; height: 60px; border-radius: 4px;" 
              preview-teleported
            />
            <span v-else style="color: #999; font-size: 12px;">无图片</span>
          </template>
        </el-table-column>
        <el-table-column prop="main_category" label="主分类" width="120">
          <template #default="{ row }">
            {{ row.main_category || '推荐' }}
          </template>
        </el-table-column>
        <el-table-column prop="sub_category" label="副分类" width="120">
          <template #default="{ row }">
            {{ row.sub_category || '--' }}
          </template>
        </el-table-column>
        <el-table-column prop="need_reference_image" label="是否需参考图" width="120">
          <template #default="{ row }">
            <el-tag :type="row.need_reference_image ? 'success' : 'info'" size="small">
              {{ row.need_reference_image ? '是' : '否' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="is_active" label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="row.is_active ? 'success' : 'danger'" size="small">
              {{ row.is_active ? '启用' : '禁用' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="content" label="提示词内容" min-width="400">
          <template #default="{ row }">
            <div class="prompt-content-clamp">{{ row.content }}</div>
          </template>
        </el-table-column>
        <el-table-column prop="created_at" label="创建时间" width="180">
          <template #default="{ row }">
            {{ formatDate(row.created_at) }}
          </template>
        </el-table-column>
        <el-table-column label="操作" min-width="120" align="center" fixed="right">
          <template #default="{ row }">
            <div class="action-btns">
              <el-tooltip content="编辑" placement="top" :hide-after="0">
                <el-button link @click="handleEdit(row)" class="action-btn">
                  <el-icon :size="16"><EditPen /></el-icon>
                </el-button>
              </el-tooltip>
              <el-tooltip content="删除" placement="top" :hide-after="0">
                <div style="display: inline-block;">
                  <el-popconfirm title="确定删除此提示词吗？" @confirm="handleDelete(row.id)">
                    <template #reference>
                      <el-button link class="action-btn danger-btn">
                        <el-icon :size="16"><Delete /></el-icon>
                      </el-button>
                    </template>
                  </el-popconfirm>
                </div>
              </el-tooltip>
            </div>
          </template>
        </el-table-column>
      </el-table>

      <div class="pagination-container">
        <el-pagination
          v-model:current-page="currentPage"
          v-model:page-size="pageSize"
          :page-sizes="[10, 20, 50, 100]"
          layout="total, sizes, prev, pager, next, jumper"
          :total="total"
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
        />
      </div>
    </el-card>

    <!-- 添加/编辑对话框 -->
    <el-dialog append-to-body v-model="dialogVisible" :title="isEdit ? '编辑提示词' : '添加提示词'" width="500px">
      <el-form :model="form" :rules="rules" ref="formRef" label-width="100px">
        <el-form-item label="示例图片" prop="image_url">
          <el-input v-model="form.image_url" placeholder="请输入图片URL（选填）" clearable />
          <div style="margin-top: 10px;" v-if="form.image_url">
            <el-image :src="form.image_url" style="width: 100px; height: 100px; border-radius: 4px;" fit="cover" />
          </div>
        </el-form-item>
        <el-form-item label="主分类" prop="main_category">
          <el-select v-model="form.main_category" filterable allow-create clearable placeholder="请选择或输入主分类" @change="handleMainCategoryChange">
            <el-option v-for="item in mainCategories" :key="item" :label="item" :value="item" />
          </el-select>
        </el-form-item>
        <el-form-item label="副分类" prop="sub_category">
          <el-select v-model="form.sub_category" filterable allow-create clearable placeholder="请选择或输入副分类">
            <el-option v-for="item in subCategories" :key="item" :label="item" :value="item" />
          </el-select>
        </el-form-item>
        <el-form-item label="提示词内容" prop="content">
          <el-input 
            v-model="form.content" 
            type="textarea" 
            :rows="6"
            placeholder="请输入提示词内容，这将显示在前台迷你边栏的探索功能中" 
          />
        </el-form-item>
        <el-form-item label="需参考图" prop="need_reference_image">
          <el-switch v-model="form.need_reference_image" />
          <span style="margin-left: 10px; color: #999; font-size: 12px;">开启后，前台应用该提示词时将要求用户上传参考图</span>
        </el-form-item>
        <el-form-item label="是否启用" prop="is_active">
          <el-switch v-model="form.is_active" />
          <span style="margin-left: 10px; color: #999; font-size: 12px;">关闭后该提示词将不在前台展示</span>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="dialogVisible = false">取消</el-button>
          <el-button type="primary" @click="submitForm" :loading="submitting">
            确定
          </el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { Plus, Search, EditPen, Delete } from '@element-plus/icons-vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import type { FormInstance } from 'element-plus'
import dayjs from 'dayjs'
import axios from 'axios'

const getAuthHeaders = () => {
  const token = localStorage.getItem('token')
  return token ? { Authorization: `Bearer ${token}` } : {}
}

const getApiUrl = (path: string) => {
  return `${(window as any).APP_CONFIG?.API_BASE_URL || ''}/api${path}`
}

interface Inspiration {
  id: number
  content: string
  image_url: string
  main_category: string
  sub_category: string
  need_reference_image: boolean
  is_active: boolean
  created_at: string
  updated_at: string
}

const loading = ref(false)
const inspirations = ref<Inspiration[]>([])
const total = ref(0)
const currentPage = ref(1)
const pageSize = ref(10)
const searchQuery = ref('')
const searchMainCategory = ref('')
const searchSubCategory = ref('')
const filterSubCategories = ref<string[]>([])
const selectedIds = ref<number[]>([])
const batchDeleting = ref(false)

const mainCategories = ref<string[]>([])
const subCategories = ref<string[]>([])
const categoryTree = ref<any[]>([])

const loadCategories = async () => {
  try {
    const res = await axios.get(getApiUrl('/admin/inspiration-categories/list'), {
      headers: getAuthHeaders()
    })
    
    if (res.data?.data) {
      const parsed = res.data.data
      categoryTree.value = parsed
      mainCategories.value = parsed.map((c: any) => c.name)
    } else {
      // Default fallback
      const defaultMain = ['人物', '风景', '建筑', '科幻', '二次元', '其他']
      mainCategories.value = defaultMain
      categoryTree.value = defaultMain.map(name => ({ name, sub: ['其他'] }))
    }
  } catch (error: any) {
    console.error('获取分类失败', error)
  }
}

const handleMainCategoryChange = (val: string) => {
  form.value.sub_category = ''
  const target = categoryTree.value.find((c: any) => c.name === val)
  if (target && target.sub) {
    subCategories.value = target.sub
  } else {
    subCategories.value = []
  }
}

const handleFilterMainCategoryChange = (val: string) => {
  searchSubCategory.value = ''
  if (!val) {
    filterSubCategories.value = []
  } else {
    const target = categoryTree.value.find((c: any) => c.name === val)
    if (target && target.sub) {
      filterSubCategories.value = target.sub
    } else {
      filterSubCategories.value = []
    }
  }
  handleSearch()
}

const handleFilterSubCategoryChange = () => {
  handleSearch()
}

const dialogVisible = ref(false)
const submitting = ref(false)
const isEdit = ref(false)
const formRef = ref<FormInstance>()

const form = ref({
  id: 0,
  content: '',
  image_url: '',
  main_category: '',
  sub_category: '',
  need_reference_image: false,
  is_active: true
})

const rules = {
  content: [
    { required: true, message: '请输入提示词内容', trigger: 'blur' }
  ]
}

const formatDate = (date: string) => {
  if (!date) return '-'
  return dayjs(date).format('YYYY-MM-DD HH:mm:ss')
}

const loadData = async () => {
  loading.value = true
  try {
    const res = await axios.get(getApiUrl('/admin/inspirations/list'), {
      params: {
        page: currentPage.value,
        page_size: pageSize.value,
        search: searchQuery.value,
        main_category: searchMainCategory.value,
        sub_category: searchSubCategory.value
      },
      headers: getAuthHeaders()
    })
    inspirations.value = res.data.list || []
    total.value = res.data.total || 0
  } catch (error: any) {
    ElMessage.error(error.response?.data?.error || '获取数据失败')
  } finally {
    loading.value = false
  }
}

const handleSearch = () => {
  currentPage.value = 1
  loadData()
}

const handleSizeChange = (val: number) => {
  pageSize.value = val
  loadData()
}

const handleCurrentChange = (val: number) => {
  currentPage.value = val
  loadData()
}

const handleSelectionChange = (selection: Inspiration[]) => {
  selectedIds.value = selection.map(item => item.id)
}

const showAddDialog = () => {
  isEdit.value = false
  form.value = {
    id: 0,
    content: '',
    image_url: '',
    main_category: '',
    sub_category: '',
    need_reference_image: false,
    is_active: true
  }
  dialogVisible.value = true
  if (formRef.value) {
    formRef.value.clearValidate()
  }
}

const handleEdit = (row: Inspiration) => {
  isEdit.value = true
  form.value = {
    id: row.id,
    content: row.content,
    image_url: row.image_url || '',
    main_category: row.main_category || '',
    sub_category: row.sub_category || '',
    need_reference_image: !!row.need_reference_image,
    is_active: row.is_active === undefined ? true : !!row.is_active
  }
  
  if (row.main_category) {
    const target = categoryTree.value.find((c: any) => c.name === row.main_category)
    if (target && target.sub) {
      subCategories.value = target.sub
    } else {
      subCategories.value = []
    }
  } else {
    subCategories.value = []
  }
  
  dialogVisible.value = true
}

const submitForm = async () => {
  if (!formRef.value) return
  await formRef.value.validate(async (valid) => {
    if (valid) {
      submitting.value = true
      try {
        if (isEdit.value) {
          await axios.post(getApiUrl('/admin/inspirations/update'), {
            id: form.value.id,
            content: form.value.content,
            image_url: form.value.image_url,
            main_category: form.value.main_category,
            sub_category: form.value.sub_category,
            need_reference_image: form.value.need_reference_image,
            is_active: form.value.is_active
          }, { headers: getAuthHeaders() })
          ElMessage.success('更新成功')
        } else {
          await axios.post(getApiUrl('/admin/inspirations/create'), {
            content: form.value.content,
            image_url: form.value.image_url,
            main_category: form.value.main_category,
            sub_category: form.value.sub_category,
            need_reference_image: form.value.need_reference_image,
            is_active: form.value.is_active
          }, { headers: getAuthHeaders() })
          ElMessage.success('添加成功')
        }
        dialogVisible.value = false
        loadData()
        loadCategories()
      } catch (error: any) {
        ElMessage.error(error.response?.data?.error || (isEdit.value ? '更新失败' : '添加失败'))
      } finally {
        submitting.value = false
      }
    }
  })
}

const handleDelete = async (id: number) => {
  try {
    await axios.post(getApiUrl('/admin/inspirations/delete'), { id }, { headers: getAuthHeaders() })
    ElMessage.success('删除成功')
    if (inspirations.value.length === 1 && currentPage.value > 1) {
      currentPage.value--
    }
    loadData()
  } catch (error: any) {
    ElMessage.error(error.response?.data?.error || '删除失败')
  }
}

const handleBatchDelete = async () => {
  try {
    await ElMessageBox.confirm(`确定要删除选中的 ${selectedIds.value.length} 个提示词吗？`, '警告', {
      type: 'warning',
      confirmButtonText: '确定',
      cancelButtonText: '取消'
    })
    
    batchDeleting.value = true
    await axios.post(getApiUrl('/admin/inspirations/batch-delete'), { ids: selectedIds.value }, { headers: getAuthHeaders() })
    ElMessage.success('批量删除成功')
    
    // 如果当前页的数据全部被删除，且不是第一页，则跳到前一页
    if (selectedIds.value.length === inspirations.value.length && currentPage.value > 1) {
      currentPage.value--
    }
    loadData()
  } catch (error: any) {
    if (error !== 'cancel') {
      ElMessage.error(error.response?.data?.error || '批量删除失败')
    }
  } finally {
    batchDeleting.value = false
  }
}

onMounted(() => {
  loadData()
  loadCategories()
})
</script>

<style scoped>
.inspirations-container {
  padding-top: 8px;
}

.page-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24px;
}

.page-header h2 {
  margin: 0;
  font-size: 24px;
  font-weight: 700;
  color: #1d1d1f;
  letter-spacing: -0.5px;
}

.box-card {
  border-radius: 24px;
}

.toolbar {
  margin-bottom: 20px;
}

.search-bar {
  display: flex;
  gap: 12px;
  flex-wrap: wrap;
}

.search-input {
  width: 240px;
}

.filter-select {
  width: 140px;
}

.pagination-container {
  margin-top: 24px;
  display: flex;
  justify-content: flex-end;
}

.prompt-content-clamp {
  display: -webkit-box;
  -webkit-box-orient: vertical;
  -webkit-line-clamp: 3;
  overflow: hidden;
  white-space: pre-wrap;
  word-break: break-all;
  line-height: 1.6;
  cursor: pointer;
  color: #434344;
  font-size: 14px;
}

.action-btns {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 12px;
}

.action-btn {
  width: 32px;
  height: 32px;
  padding: 0 !important;
  margin: 0 !important;
  border-radius: 8px !important;
  background: rgba(245, 245, 247, 0.5) !important;
  border: 1px solid rgba(229, 229, 234, 0.5) !important;
  transition: all 0.2s;
  color: #86868b !important;
  outline: none !important;
  display: flex !important;
  align-items: center !important;
  justify-content: center !important;
}

.action-btn:focus,
.action-btn:active,
.action-btn:focus-visible {
  outline: none !important;
  box-shadow: none !important;
  background: rgba(245, 245, 247, 0.5) !important;
  color: #86868b !important;
}

.action-btn:hover {
  background: #ffffff !important;
  border-color: #e5e5ea !important;
  box-shadow: 0 4px 6px -1px rgba(0,0,0,0.05) !important;
  color: #409eff !important;
  transform: translateY(-2px);
}

.danger-btn:hover {
  color: #ff3b30 !important;
}

/* ================= 移动端响应式 ================= */
@media (max-width: 768px) {
  .page-header {
    flex-direction: column;
    align-items: flex-start;
    gap: 16px;
  }
  
  .page-header .el-button {
    width: 100%;
  }

  .search-bar {
    flex-direction: column;
    width: 100%;
  }

  .search-input, .filter-select {
    width: 100% !important;
  }
  
  .search-bar .el-button {
    width: 100%;
    margin-left: 0 !important;
  }

  .pagination-container {
    justify-content: center;
  }
}
</style>
