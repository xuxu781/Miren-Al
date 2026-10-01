<template>
  <div class="admin-container fade-in">
    <div class="page-header-modern">
      <div class="header-info">
        <div class="icon-wrapper">
          <el-icon :size="24"><Avatar /></el-icon>
        </div>
        <div class="header-title">
          <h2>系统管理员</h2>
          <span class="subtitle">管理系统所有后台账号，设置权限及修改密码</span>
        </div>
      </div>
      <div class="header-actions">
        <el-input
          v-model="searchQuery"
          placeholder="搜索管理员账号..."
          :prefix-icon="Search"
          clearable
          class="modern-search"
        />
        <el-button type="primary" class="add-btn" @click="showAddDialog = true">
          <el-icon class="el-icon--left"><Plus /></el-icon>
          添加管理员
        </el-button>
      </div>
    </div>

    <el-card class="modern-card">
      <el-table 
        :data="pagedList" 
        v-loading="loading" 
        style="width: 100%" 
        class="modern-table"
      >
        <el-table-column prop="id" label="ID" width="100" align="center">
          <template #default="scope">
            <div class="id-badge">#{{ scope.row.id }}</div>
          </template>
        </el-table-column>
        
        <el-table-column prop="username" label="管理员账号" min-width="150">
          <template #default="scope">
            <div class="user-info">
              <el-avatar :size="36" class="user-avatar" :style="{ background: getAvatarColor(scope.row.username) }">
                {{ scope.row.username.charAt(0).toUpperCase() }}
              </el-avatar>
              <div class="user-details">
                <span class="username">{{ scope.row.username }}</span>
                <el-tag v-if="scope.row.id === currentAdminId" size="small" type="success" effect="light" class="current-tag">当前账号</el-tag>
              </div>
            </div>
          </template>
        </el-table-column>

        <el-table-column prop="created_at" label="创建时间" min-width="180" align="center">
          <template #default="scope">
            <div class="time-cell">
              <el-icon><Calendar /></el-icon>
              <span>{{ new Date(scope.row.created_at).toLocaleString() }}</span>
            </div>
          </template>
        </el-table-column>
        
        <el-table-column prop="updated_at" label="更新时间" min-width="180" align="center">
          <template #default="scope">
            <div class="time-cell">
              <el-icon><Clock /></el-icon>
              <span>{{ new Date(scope.row.updated_at).toLocaleString() }}</span>
            </div>
          </template>
        </el-table-column>

        <el-table-column label="操作" min-width="120" align="center" fixed="right">
          <template #default="scope">
            <div class="action-btns">
              <el-tooltip content="修改密码" placement="top">
                <el-button type="primary" link @click="openPwdDialog(scope.row)" class="action-btn">
                  <el-icon :size="16"><EditPen /></el-icon>
                </el-button>
              </el-tooltip>
              
              <el-tooltip content="删除账号" placement="top" :disabled="scope.row.id === currentAdminId">
                <div style="display: inline-block;">
                  <el-popconfirm 
                    title="确定要删除该管理员吗？此操作不可恢复！" 
                    confirm-button-text="确定删除"
                    cancel-button-text="取消"
                    confirm-button-type="danger"
                    :icon="Warning"
                    icon-color="#f56c6c"
                    @confirm="handleDelete(scope.row)"
                  >
                    <template #reference>
                      <el-button type="danger" link :disabled="scope.row.id === currentAdminId" class="action-btn danger-btn">
                        <el-icon :size="16"><Delete /></el-icon>
                      </el-button>
                    </template>
                  </el-popconfirm>
                </div>
              </el-tooltip>
            </div>
          </template>
        </el-table-column>
        
        <template #empty>
          <el-empty description="暂无管理员数据" :image-size="120" />
        </template>
      </el-table>

      <div class="pagination-container">
        <el-pagination
          v-model:current-page="currentPage"
          v-model:page-size="pageSize"
          :page-sizes="[10, 20, 50, 100]"
          :background="true"
          layout="total, sizes, prev, pager, next, jumper"
          :total="filteredList.length"
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
        />
      </div>
    </el-card>

    <!-- 修改密码对话框 -->
    <el-dialog append-to-body v-model="showPwdDialog" title="修改管理员密码" width="420px" destroy-on-close class="custom-dialog">
      <div class="dialog-desc">请输入新的登录密码，修改后下次登录生效。</div>
      <el-form ref="pwdFormRef" :model="pwdForm" :rules="pwdRules" label-width="0" size="large">
        <el-form-item prop="new_password">
          <el-input v-model="pwdForm.new_password" type="password" placeholder="请输入新密码 (至少6位)" show-password :prefix-icon="Lock" />
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="showPwdDialog = false">取 消</el-button>
          <el-button type="primary" @click="handleUpdatePassword" :loading="pwdLoading">确认修改</el-button>
        </span>
      </template>
    </el-dialog>

    <!-- 添加管理员对话框 -->
    <el-dialog append-to-body v-model="showAddDialog" title="添加新管理员" width="420px" destroy-on-close class="custom-dialog">
      <div class="dialog-desc">添加新的系统管理员账号，默认拥有所有后台权限。</div>
      <el-form ref="formRef" :model="form" :rules="rules" label-width="0" size="large">
        <el-form-item prop="username">
          <el-input v-model="form.username" placeholder="请输入管理员用户名" :prefix-icon="User" />
        </el-form-item>
        <el-form-item prop="password">
          <el-input v-model="form.password" type="password" placeholder="请输入初始密码 (至少6位)" show-password :prefix-icon="Lock" />
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="showAddDialog = false">取 消</el-button>
          <el-button type="primary" @click="handleAdd" :loading="submitLoading">确认添加</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { Plus, EditPen, Delete, Calendar, Clock, User, Lock, Warning, Search, Avatar } from '@element-plus/icons-vue'

const adminList = ref<any[]>([])
const loading = ref(false)
const showAddDialog = ref(false)
const submitLoading = ref(false)
const formRef = ref()

// 搜索和分页
const searchQuery = ref('')
const currentPage = ref(1)
const pageSize = ref(10)

const filteredList = computed(() => {
  if (!searchQuery.value) return adminList.value
  return adminList.value.filter(item => 
    item.username.toLowerCase().includes(searchQuery.value.toLowerCase())
  )
})

const pagedList = computed(() => {
  const start = (currentPage.value - 1) * pageSize.value
  const end = start + pageSize.value
  return filteredList.value.slice(start, end)
})

const handleSizeChange = (val: number) => {
  pageSize.value = val
  currentPage.value = 1
}

const handleCurrentChange = (val: number) => {
  currentPage.value = val
}

const getAvatarColor = (name: string) => {
  const colors = ['#409eff', '#67c23a', '#e6a23c', '#f56c6c', '#8e44ad', '#16a085', '#d35400'];
  let hash = 0;
  for (let i = 0; i < name.length; i++) {
    hash = name.charCodeAt(i) + ((hash << 5) - hash);
  }
  return colors[Math.abs(hash) % colors.length];
}

const form = ref({
  username: '',
  password: ''
})

const showPwdDialog = ref(false)
const pwdLoading = ref(false)
const pwdFormRef = ref()
const pwdForm = ref({
  id: null as number | null,
  new_password: ''
})

const pwdRules = {
  new_password: [
    { required: true, message: '请输入新密码', trigger: 'blur' },
    { min: 6, message: '密码长度不能少于6位', trigger: 'blur' }
  ]
}

const rules = {
  username: [{ required: true, message: '请输入用户名', trigger: 'blur' }],
  password: [
    { required: true, message: '请输入密码', trigger: 'blur' },
    { min: 6, message: '密码长度不能少于6位', trigger: 'blur' }
  ]
}

const currentAdminId = ref<number | null>(null)

const getHeaders = () => {
  const token = localStorage.getItem('token')
  return {
    'Content-Type': 'application/json',
    'Authorization': `Bearer ${token}`
  }
}

const fetchCurrentAdmin = async () => {
  try {
    const res = await fetch(`${(window as any).APP_CONFIG?.API_BASE_URL || ''}/api/admin/info`, { headers: getHeaders() })
    if (res.ok) {
      const data = await res.json()
      currentAdminId.value = data.id
    }
  } catch (error) {
    console.error('Failed to fetch current admin info', error)
  }
}

const fetchAdminList = async () => {
  loading.value = true
  try {
    const res = await fetch(`${(window as any).APP_CONFIG?.API_BASE_URL || ''}/api/admin/list`, { headers: getHeaders() })
    const data = await res.json()
    if (res.ok) {
      adminList.value = data.list || []
    } else {
      ElMessage.error(data.error || '获取管理员列表失败')
    }
  } catch (error) {
    ElMessage.error('网络错误，请稍后重试')
  } finally {
    loading.value = false
  }
}

const handleAdd = async () => {
  if (!formRef.value) return
  await formRef.value.validate(async (valid: boolean) => {
    if (valid) {
      submitLoading.value = true
      try {
        const res = await fetch(`${(window as any).APP_CONFIG?.API_BASE_URL || ''}/api/admin/create`, {
          method: 'POST',
          headers: getHeaders(),
          body: JSON.stringify(form.value)
        })
        const data = await res.json()
        if (res.ok) {
          ElMessage.success('添加成功')
          showAddDialog.value = false
          form.value = { username: '', password: '' }
          fetchAdminList()
        } else {
          ElMessage.error(data.error || '添加失败')
        }
      } catch (error) {
        ElMessage.error('网络错误，请稍后重试')
      } finally {
        submitLoading.value = false
      }
    }
  })
}

const openPwdDialog = (row: any) => {
  pwdForm.value.id = row.id
  pwdForm.value.new_password = ''
  showPwdDialog.value = true
}

const handleUpdatePassword = async () => {
  if (!pwdFormRef.value) return
  await pwdFormRef.value.validate(async (valid: boolean) => {
    if (valid) {
      pwdLoading.value = true
      try {
        const res = await fetch(`${(window as any).APP_CONFIG?.API_BASE_URL || ''}/api/admin/update-password`, {
          method: 'POST',
          headers: getHeaders(),
          body: JSON.stringify({
            id: pwdForm.value.id,
            new_password: pwdForm.value.new_password
          })
        })
        const data = await res.json()
        if (res.ok) {
          ElMessage.success('密码修改成功')
          showPwdDialog.value = false
        } else {
          ElMessage.error(data.error || '密码修改失败')
        }
      } catch (error) {
        ElMessage.error('网络错误，请稍后重试')
      } finally {
        pwdLoading.value = false
      }
    }
  })
}

const handleDelete = async (row: any) => {
  try {
    const res = await fetch(`${(window as any).APP_CONFIG?.API_BASE_URL || ''}/api/admin/delete`, {
      method: 'POST',
      headers: getHeaders(),
      body: JSON.stringify({ id: row.id })
    })
    const data = await res.json()
    if (res.ok) {
      ElMessage.success('删除成功')
      fetchAdminList()
    } else {
      ElMessage.error(data.error || '删除失败')
    }
  } catch (error) {
    ElMessage.error('网络错误，请稍后重试')
  }
}

onMounted(() => {
  fetchCurrentAdmin()
  fetchAdminList()
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

.admin-container {
  padding-top: 8px;
}

.page-header-modern {
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

.header-actions {
  display: flex;
  align-items: center;
  gap: 16px;
}

.modern-search {
  width: 280px;
}

.id-badge {
  display: inline-block;
  padding: 4px 10px;
  background: rgba(245, 245, 247, 0.8);
  border-radius: 6px;
  color: #86868b;
  font-family: -apple-system, BlinkMacSystemFont, monospace;
  font-size: 13px;
  font-weight: 600;
}

.user-info {
  display: flex;
  align-items: center;
  gap: 12px;
}

.user-avatar {
  color: #fff;
  font-weight: 600;
  font-size: 16px;
  border: 2px solid #ffffff;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
}

.user-details {
  display: flex;
  flex-direction: column;
  gap: 4px;
  align-items: flex-start;
}

.username {
  font-weight: 600;
  color: #1d1d1f;
  font-size: 14px;
}

.current-tag {
  background: rgba(103, 194, 58, 0.1);
  color: #67c23a;
  border: none;
  font-weight: 600;
}

.time-cell {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  color: #86868b;
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
  padding: 8px;
  border-radius: 8px;
  background: rgba(245, 245, 247, 0.5);
  border: 1px solid rgba(229, 229, 234, 0.5);
  transition: all 0.2s;
  color: #86868b;
}

.action-btn:hover {
  background: #ffffff;
  border-color: #e5e5ea;
  box-shadow: 0 4px 6px -1px rgba(0,0,0,0.05);
  color: #409eff;
  transform: translateY(-2px);
}

.danger-btn:hover {
  color: #ff3b30;
}

.pagination-container {
  display: flex;
  justify-content: flex-end;
  padding: 24px 0 0;
  margin-top: 12px;
}

.dialog-desc {
  color: #64748b;
  font-size: 14px;
  margin-bottom: 24px;
  margin-top: -10px;
}

/* ================= 移动端响应式 ================= */
@media (max-width: 768px) {
  .page-header-modern {
    flex-direction: column;
    align-items: flex-start;
    gap: 16px;
  }
  
  .header-actions {
    width: 100%;
    flex-direction: column;
    align-items: stretch;
  }

  .modern-search {
    width: 100%;
  }

  .add-btn {
    width: 100%;
    margin-left: 0 !important;
  }

  .pagination-container {
    justify-content: center;
  }
}
</style>
