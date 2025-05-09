<template>
  <div class="student-management">
    <a-card :bordered="false">
      <!-- Search Form -->
      <a-form layout="inline" :model="queryParams" class="search-form">
        <a-form-item label="姓名">
          <a-input v-model:value="queryParams.name" placeholder="请输入学生姓名" allowClear />
        </a-form-item>
        <a-form-item label="手机号">
          <a-input v-model:value="queryParams.phoneNumber" placeholder="请输入手机号" allowClear />
        </a-form-item>
        <a-form-item label="状态">
          <a-select v-model:value="queryParams.status" placeholder="请选择状态" allowClear style="width: 120px">
            <a-select-option value="active">活跃</a-select-option>
            <a-select-option value="expired">已过期</a-select-option>
          </a-select>
        </a-form-item>
        <a-form-item>
          <a-space>
            <a-button type="primary" @click="handleSearch">
              <template #icon><SearchOutlined /></template>
              查询
            </a-button>
            <a-button @click="resetQuery">
              <template #icon><ReloadOutlined /></template>
              重置
            </a-button>
          </a-space>
        </a-form-item>
      </a-form>

      <!-- Table Toolbar -->
      <div class="table-operations">
        <a-space>
          <a-button type="primary" @click="handleAdd">
            <template #icon><PlusOutlined /></template>
            新增
          </a-button>
          <a-button @click="handleExport">
            <template #icon><ExportOutlined /></template>
            导出
          </a-button>
          <a-button @click="handleImport">
            <template #icon><ImportOutlined /></template>
            导入
          </a-button>
        </a-space>
      </div>

      <!-- Student Table -->
      <a-table
        :columns="columns"
        :data-source="studentList"
        :loading="loading"
        :pagination="pagination"
        @change="handleTableChange"
        row-key="id"
      >
        <template #bodyCell="{ column, record }">
          <!-- Avatar Column -->
          <template v-if="column.dataIndex === 'avatarUrl'">
            <a-avatar :src="record.avatar_url" :size="40">
              <template #icon><UserOutlined /></template>
            </a-avatar>
          </template>
          
          <!-- Status Column -->
          <template v-if="column.dataIndex === 'status'">
            <a-tag :color="record.status === 'active' ? 'green' : 'red'">
              {{ record.status === 'active' ? '活跃' : '已过期' }}
            </a-tag>
          </template>
          
          <!-- Expiry Date Column -->
          <template v-if="column.dataIndex === 'expiryDate'">
            {{ formatDate(record.expiry_date) }}
          </template>
          
          <!-- Actions Column -->
          <template v-if="column.dataIndex === 'action'">
            <a-space>
              <a @click="handleEdit(record)">编辑</a>
              <a-divider type="vertical" />
              <a-popconfirm
                title="确定要删除该学生吗？"
                ok-text="确定"
                cancel-text="取消"
                @confirm="handleDelete(record)"
              >
                <a class="text-danger">删除</a>
              </a-popconfirm>
            </a-space>
          </template>
        </template>
      </a-table>
    </a-card>

    <!-- Student Form Drawer -->
    <a-drawer
      :title="formTitle"
      :visible="drawerVisible"
      :width="500"
      @close="closeDrawer"
      :body-style="{ paddingBottom: '80px' }"
    >
      <a-form
        :model="formData"
        :rules="formRules"
        ref="formRef"
        layout="vertical"
      >
        <a-form-item name="name" label="姓名" required>
          <a-input v-model:value="formData.name" placeholder="请输入学生姓名" />
        </a-form-item>
        <a-form-item name="phoneNumber" label="手机号">
          <a-input v-model:value="formData.phone_number" placeholder="请输入手机号" />
        </a-form-item>
        <a-form-item name="expiryDate" label="会员到期日期">
          <a-date-picker 
            v-model:value="formData.expiryDateMoment" 
            style="width: 100%" 
            :disabledDate="disabledDate"
            format="YYYY-MM-DD"
          />
        </a-form-item>
        <a-form-item name="status" label="状态">
          <a-select v-model:value="formData.status" placeholder="请选择状态">
            <a-select-option value="active">活跃</a-select-option>
            <a-select-option value="expired">已过期</a-select-option>
          </a-select>
        </a-form-item>
      </a-form>
      <div class="drawer-footer">
        <a-space>
          <a-button @click="closeDrawer">取消</a-button>
          <a-button type="primary" @click="submitForm" :loading="submitLoading">提交</a-button>
        </a-space>
      </div>
    </a-drawer>
  </div>
</template>

<script setup>
import { ref, reactive, computed, onMounted } from 'vue'
import { message } from 'ant-design-vue'
import { 
  SearchOutlined, 
  ReloadOutlined, 
  PlusOutlined, 
  ExportOutlined, 
  ImportOutlined,
  UserOutlined
} from '@ant-design/icons-vue'
import { 
  getStudentsApi, 
  getStudentApi, 
  createStudentApi, 
  updateStudentApi, 
  deleteStudentApi,
  exportStudentsApi
} from '~@/api/english_check_in/student'
import dayjs from 'dayjs'

// Table columns
const columns = [
  {
    title: '头像',
    dataIndex: 'avatarUrl',
    width: 80,
  },
  {
    title: '姓名',
    dataIndex: 'name',
    ellipsis: true,
  },
  {
    title: '手机号',
    dataIndex: 'phone_number',
    ellipsis: true,
  },
  {
    title: '会员到期日期',
    dataIndex: 'expiryDate',
    sorter: true,
  },
  {
    title: '状态',
    dataIndex: 'status',
    width: 100,
  },
  {
    title: '操作',
    dataIndex: 'action',
    width: 150,
    fixed: 'right',
  },
]

// State
const loading = ref(false)
const studentList = ref([])
const pagination = reactive({
  current: 1,
  pageSize: 10,
  total: 0,
  showSizeChanger: true,
  showTotal: (total) => `共 ${total} 条`,
})

// Query parameters
const queryParams = reactive({
  name: '',
  phoneNumber: '',
  status: undefined,
  pageNum: 1,
  pageSize: 10,
})

// Form state
const formRef = ref(null)
const drawerVisible = ref(false)
const submitLoading = ref(false)
const formMode = ref('add') // 'add' or 'edit'
const formTitle = computed(() => formMode.value === 'add' ? '新增学生' : '编辑学生')
const formData = reactive({
  id: undefined,
  name: '',
  phone_number: '',
  expiryDateMoment: null,
  status: 'active',
})

// Form validation rules
const formRules = {
  name: [{ required: true, message: '请输入学生姓名', trigger: 'blur' }],
}

// Fetch student list
const fetchStudentList = async () => {
  loading.value = true
  try {
    const params = {
      pageNum: pagination.current,
      pageSize: pagination.pageSize,
      ...queryParams,
    }
    const res = await getStudentsApi(params)
    studentList.value = res.data.list || []
    pagination.total = res.data.total || 0
  } catch (error) {
    console.error('获取学生列表失败:', error)
    message.error('获取学生列表失败')
  } finally {
    loading.value = false
  }
}

// Handle table change (pagination, filters, sorter)
const handleTableChange = (pag) => {
  pagination.current = pag.current
  pagination.pageSize = pag.pageSize
  fetchStudentList()
}

// Search
const handleSearch = () => {
  pagination.current = 1
  fetchStudentList()
}

// Reset search form
const resetQuery = () => {
  Object.assign(queryParams, {
    name: '',
    phoneNumber: '',
    status: undefined,
  })
  pagination.current = 1
  fetchStudentList()
}

// Add new student
const handleAdd = () => {
  formMode.value = 'add'
  resetForm()
  drawerVisible.value = true
}

// Edit student
const handleEdit = async (record) => {
  formMode.value = 'edit'
  resetForm()
  
  try {
    const res = await getStudentApi({ id: record.id })
    const studentData = res.data
    
    Object.assign(formData, {
      id: studentData.id,
      name: studentData.name,
      phone_number: studentData.phone_number,
      expiryDateMoment: studentData.expiry_date ? dayjs(studentData.expiry_date) : null,
      status: studentData.status,
    })
    
    drawerVisible.value = true
  } catch (error) {
    console.error('获取学生详情失败:', error)
    message.error('获取学生详情失败')
  }
}

// Delete student
const handleDelete = async (record) => {
  try {
    await deleteStudentApi({ id: record.id })
    message.success('删除成功')
    fetchStudentList()
  } catch (error) {
    console.error('删除学生失败:', error)
    message.error('删除学生失败')
  }
}

// Export students
const handleExport = async () => {
  try {
    const res = await exportStudentsApi(queryParams)
    const blob = new Blob([res], { type: 'application/vnd.ms-excel' })
    const url = window.URL.createObjectURL(blob)
    const link = document.createElement('a')
    link.href = url
    link.download = `学生列表_${new Date().getTime()}.xlsx`
    link.click()
    window.URL.revokeObjectURL(url)
    message.success('导出成功')
  } catch (error) {
    console.error('导出学生失败:', error)
    message.error('导出学生失败')
  }
}

// Import students
const handleImport = () => {
  message.info('导入功能待实现')
}

// Submit form
const submitForm = async () => {
  try {
    await formRef.value.validate()
    
    submitLoading.value = true
    
    // Convert form data to API format
    const apiData = {
      id: formData.id,
      name: formData.name,
      phone_number: formData.phone_number,
      expiry_date: formData.expiryDateMoment ? formData.expiryDateMoment.valueOf() : null,
      status: formData.status,
    }
    
    if (formMode.value === 'add') {
      await createStudentApi(apiData)
      message.success('添加成功')
    } else {
      await updateStudentApi(apiData)
      message.success('更新成功')
    }
    
    closeDrawer()
    fetchStudentList()
  } catch (error) {
    console.error('提交表单失败:', error)
    message.error('提交失败，请检查表单')
  } finally {
    submitLoading.value = false
  }
}

// Close drawer and reset form
const closeDrawer = () => {
  drawerVisible.value = false
}

// Reset form
const resetForm = () => {
  if (formRef.value) {
    formRef.value.resetFields()
  }
  
  Object.assign(formData, {
    id: undefined,
    name: '',
    phone_number: '',
    expiryDateMoment: null,
    status: 'active',
  })
}

// Format date
const formatDate = (timestamp) => {
  if (!timestamp) return '--'
  return dayjs(timestamp).format('YYYY-MM-DD')
}

// Disable past dates
const disabledDate = (current) => {
  return current && current < dayjs().startOf('day')
}

// Lifecycle hooks
onMounted(() => {
  fetchStudentList()
})
</script>

<style scoped>
.student-management {
  padding: 24px;
}

.search-form {
  margin-bottom: 24px;
}

.table-operations {
  margin-bottom: 16px;
  display: flex;
  justify-content: flex-end;
}

.drawer-footer {
  position: absolute;
  bottom: 0;
  width: 100%;
  border-top: 1px solid #e8e8e8;
  padding: 10px 16px;
  text-align: right;
  left: 0;
  background: #fff;
}

.text-danger {
  color: #ff4d4f;
}
</style>
