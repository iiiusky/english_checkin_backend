<template>
  <div class="task-management">
    <a-card :bordered="false">
      <!-- Search Form -->
      <a-form layout="inline" :model="queryParams" class="search-form">
        <a-form-item label="英文内容">
          <a-input v-model:value="queryParams.englishText" placeholder="请输入英文内容" allowClear />
        </a-form-item>
        <a-form-item label="任务类型">
          <a-select v-model:value="queryParams.taskType" placeholder="请选择任务类型" allowClear style="width: 120px">
            <a-select-option value="reading">阅读</a-select-option>
            <a-select-option value="pronunciation">发音</a-select-option>
            <a-select-option value="writing">写作</a-select-option>
            <a-select-option value="listening">听力</a-select-option>
          </a-select>
        </a-form-item>
        <a-form-item label="难度等级">
          <a-select v-model:value="queryParams.difficultyLevel" placeholder="请选择难度等级" allowClear style="width: 120px">
            <a-select-option :value="1">1级</a-select-option>
            <a-select-option :value="2">2级</a-select-option>
            <a-select-option :value="3">3级</a-select-option>
            <a-select-option :value="4">4级</a-select-option>
            <a-select-option :value="5">5级</a-select-option>
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

      <!-- Task Table -->
      <a-table
        :columns="columns"
        :data-source="taskList"
        :loading="loading"
        :pagination="pagination"
        @change="handleTableChange"
        row-key="id"
      >
        <template #bodyCell="{ column, record }">
          <!-- Task Type Column -->
          <template v-if="column.dataIndex === 'taskType'">
            <a-tag :color="getTaskTypeColor(record.task_type)">
              {{ getTaskTypeText(record.task_type) }}
            </a-tag>
          </template>
          
          <!-- Difficulty Level Column -->
          <template v-if="column.dataIndex === 'difficultyLevel'">
            <a-rate :value="record.difficulty_level" disabled :count="5" />
          </template>
          
          <!-- Date Columns -->
          <template v-if="column.dataIndex === 'start_time' || column.dataIndex === 'end_time'">
            {{ formatDate(record[column.dataIndex]) }}
          </template>
          
          <!-- Actions Column -->
          <template v-if="column.dataIndex === 'action'">
            <a-space>
              <a @click="handleEdit(record)">编辑</a>
              <a-divider type="vertical" />
              <a @click="handleViewStats(record)">统计</a>
              <a-divider type="vertical" />
              <a-dropdown>
                <a class="ant-dropdown-link" @click.prevent>
                  更多 <down-outlined />
                </a>
                <template #overlay>
                  <a-menu>
                    <a-menu-item @click="handlePublish(record)" v-if="canPublish(record)">
                      <check-circle-outlined /> 发布
                    </a-menu-item>
                    <a-menu-item @click="handleCancel(record)" v-if="canCancel(record)">
                      <stop-outlined /> 取消
                    </a-menu-item>
                    <a-menu-item @click="handleComplete(record)" v-if="canComplete(record)">
                      <check-outlined /> 完成
                    </a-menu-item>
                    <a-menu-item @click="handleDelete(record)">
                      <delete-outlined /> 删除
                    </a-menu-item>
                  </a-menu>
                </template>
              </a-dropdown>
            </a-space>
          </template>
        </template>
      </a-table>
    </a-card>

    <!-- Task Form Drawer -->
    <a-drawer
      :title="formTitle"
      :visible="drawerVisible"
      :width="600"
      @close="closeDrawer"
      :body-style="{ paddingBottom: '80px' }"
    >
      <a-form
        :model="formData"
        :rules="formRules"
        ref="formRef"
        layout="vertical"
      >
        <a-form-item name="english_text" label="英文内容" required>
          <a-textarea 
            v-model:value="formData.english_text" 
            placeholder="请输入任务英文内容" 
            :auto-size="{ minRows: 3, maxRows: 6 }" 
          />
        </a-form-item>
        <a-form-item name="chinese_text" label="中文内容/翻译">
          <a-textarea 
            v-model:value="formData.chinese_text" 
            placeholder="请输入任务中文内容或翻译" 
            :auto-size="{ minRows: 3, maxRows: 6 }" 
          />
        </a-form-item>
        <a-row :gutter="16">
          <a-col :span="12">
            <a-form-item name="startTimeMoment" label="开始时间" required>
              <a-date-picker 
                v-model:value="formData.startTimeMoment" 
                show-time 
                format="YYYY-MM-DD HH:mm:ss" 
                style="width: 100%" 
              />
            </a-form-item>
          </a-col>
          <a-col :span="12">
            <a-form-item name="endTimeMoment" label="结束时间" required>
              <a-date-picker 
                v-model:value="formData.endTimeMoment" 
                show-time 
                format="YYYY-MM-DD HH:mm:ss" 
                style="width: 100%" 
                :disabled-date="disabledEndDate"
              />
            </a-form-item>
          </a-col>
        </a-row>

        <a-form-item name="students" label="分配学生">
          <a-select
            v-model:value="formData.students"
            mode="multiple"
            placeholder="请选择分配的学生"
            style="width: 100%"
            :options="studentOptions"
          ></a-select>
        </a-form-item>
        <a-row :gutter="16">
          <a-col :span="12">
            <a-form-item name="difficulty_level" label="难度等级" required>
              <a-rate v-model:value="formData.difficulty_level" :count="5" />
            </a-form-item>
          </a-col>
          <a-col :span="12">
            <a-form-item name="task_type" label="任务类型" required>
              <a-select v-model:value="formData.task_type" placeholder="请选择任务类型">
                <a-select-option value="reading">阅读</a-select-option>
                <a-select-option value="pronunciation">发音</a-select-option>
                <a-select-option value="writing">写作</a-select-option>
                <a-select-option value="listening">听力</a-select-option>
              </a-select>
            </a-form-item>
          </a-col>
        </a-row>
      </a-form>
      <div class="drawer-footer">
        <a-space>
          <a-button @click="closeDrawer">取消</a-button>
          <a-button type="primary" @click="submitForm" :loading="submitLoading">提交</a-button>
        </a-space>
      </div>
    </a-drawer>

    <!-- Task Stats Drawer -->
    <a-drawer
      title="任务完成统计"
      :visible="statsDrawerVisible"
      :width="600"
      @close="closeStatsDrawer"
    >
      <a-spin :spinning="statsLoading">
        <a-descriptions bordered :column="1" size="small">
          <a-descriptions-item label="任务英文内容">
            {{ currentTask.english_text }}
          </a-descriptions-item>
          <a-descriptions-item label="任务中文内容">
            {{ currentTask.chinese_text }}
          </a-descriptions-item>
          <a-descriptions-item label="任务类型">
            {{ getTaskTypeText(currentTask.task_type) }}
          </a-descriptions-item>
          <a-descriptions-item label="难度等级">
            <a-rate :value="currentTask.difficulty_level" disabled :count="5" />
          </a-descriptions-item>
          <a-descriptions-item label="打卡日期">
            {{ formatDate(currentTask.checkin) }}
          </a-descriptions-item>
        </a-descriptions>

        <a-divider>完成情况</a-divider>
        
        <a-progress 
          :percent="completionStats.completionRate" 
          :format="percent => `${percent}% 完成率`" 
          status="active" 
          style="margin-bottom: 20px" 
        />
        
        <a-row :gutter="16">
          <a-col :span="8">
            <a-statistic 
              title="总学生数" 
              :value="completionStats.totalStudents" 
              style="margin-bottom: 16px" 
            />
          </a-col>
          <a-col :span="8">
            <a-statistic 
              title="已完成" 
              :value="completionStats.completedCount" 
              style="margin-bottom: 16px" 
              :value-style="{ color: '#3f8600' }"
            />
          </a-col>
          <a-col :span="8">
            <a-statistic 
              title="未完成" 
              :value="completionStats.pendingCount" 
              style="margin-bottom: 16px"
              :value-style="{ color: '#cf1322' }"
            />
          </a-col>
        </a-row>

        <a-divider>学生完成列表</a-divider>
        
        <a-table
          :columns="studentCompletionColumns"
          :data-source="completionStats.studentList || []"
          size="small"
          :pagination="{ pageSize: 5 }"
        >
          <template #bodyCell="{ column, record }">
            <template v-if="column.dataIndex === 'status'">
              <a-tag :color="record.completed ? 'green' : 'red'">
                {{ record.completed ? '已完成' : '未完成' }}
              </a-tag>
            </template>
            <template v-if="column.dataIndex === 'completionTime'">
              {{ record.completion_time ? formatDate(record.completion_time) : '-' }}
            </template>
          </template>
        </a-table>
      </a-spin>
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
  CheckCircleOutlined,
  StopOutlined,
  CheckOutlined,
  DeleteOutlined,
  DownOutlined
} from '@ant-design/icons-vue'
import { 
  getTasksApi, 
  getTaskApi, 
  createTaskApi, 
  updateTaskApi, 
  deleteTaskApi,
  publishTaskApi,
  cancelTaskApi,
  completeTaskApi,
  getTaskCompletionStatsApi,
  exportTasksApi
} from '~@/api/english_check_in/task'
import { getStudentsApi } from '~@/api/english_check_in/student'
import dayjs from 'dayjs'

// Table columns
const columns = [
  {
    title: '英文内容',
    dataIndex: 'english_text',
    ellipsis: true,
    width: 250,
  },
  {
    title: '中文内容',
    dataIndex: 'chinese_text',
    ellipsis: true,
    width: 250,
  },
  {
    title: '任务类型',
    dataIndex: 'taskType',
    width: 100,
  },
  {
    title: '难度等级',
    dataIndex: 'difficultyLevel',
    width: 150,
  },
  {
    title: '开始时间',
    dataIndex: 'start_time',
    sorter: true,
    width: 150,
  },
  {
    title: '结束时间',
    dataIndex: 'end_time',
    sorter: true,
    width: 150,
  },

  {
    title: '操作',
    dataIndex: 'action',
    width: 180,
    fixed: 'right',
  },
]

// Student completion columns for stats drawer
const studentCompletionColumns = [
  {
    title: '学生姓名',
    dataIndex: 'name',
    ellipsis: true,
  },
  {
    title: '状态',
    dataIndex: 'status',
    width: 100,
  },
  {
    title: '完成时间',
    dataIndex: 'completionTime',
    width: 150,
  }
]

// State
const loading = ref(false)
const taskList = ref([])
const drawerVisible = ref(false)
const submitLoading = ref(false)
const formRef = ref(null)
const formTitle = ref('新增任务')
const isEdit = ref(false)
// 学生选项
const studentOptions = ref([])

// Stats drawer state
const statsDrawerVisible = ref(false)
const statsLoading = ref(false)
const currentTask = ref({})
const completionStats = ref({
  completionRate: 0,
  totalStudents: 0,
  completedCount: 0,
  pendingCount: 0,
  studentList: []
})

// Pagination
const pagination = reactive({
  current: 1,
  pageSize: 10,
  total: 0,
  showSizeChanger: true,
  showQuickJumper: true,
  showTotal: (total) => `共 ${total} 条记录`,
})

// Query parameters
const queryParams = reactive({
  englishText: '',
  taskType: undefined,
  difficultyLevel: undefined,
  pageNum: 1,
  pageSize: 10,
})

// Form data
const formData = reactive({
  id: undefined,
  english_text: '',
  chinese_text: '',
  startTimeMoment: null,
  endTimeMoment: null,

  start_time: 0,
  end_time: 0,

  students: [],
  difficulty_level: 3,
  task_type: 'reading',
})

// Form validation rules
const formRules = {
  english_text: [{ required: true, message: '请输入英文内容', trigger: 'blur' }],
  startTimeMoment: [{ required: true, message: '请选择开始时间', trigger: 'change' }],
  endTimeMoment: [{ required: true, message: '请选择结束时间', trigger: 'change' }],

  difficulty_level: [{ required: true, message: '请选择难度等级', trigger: 'change' }],
  task_type: [{ required: true, message: '请选择任务类型', trigger: 'change' }],
}

// Fetch task list
const fetchTaskList = async () => {
  loading.value = true
  try {
    const params = {
      ...queryParams,
      pageNum: pagination.current,
      pageSize: pagination.pageSize,
    }
    const res = await getTasksApi(params)
    if (res.code === 0) {
      taskList.value = res.data.list || []
      pagination.total = res.data.total || 0
    } else {
      message.error(res.msg || '获取任务列表失败')
    }
  } catch (error) {
    console.error('获取任务列表失败:', error)
    message.error('获取任务列表失败')
  } finally {
    loading.value = false
  }
}

// Fetch student options for select
const fetchStudentOptions = async () => {
  try {
    // 设置较大的pageSize确保获取所有学生
    const res = await getStudentsApi({ pageSize: 9999 })
    console.log('学生数据响应:', res)
    
    if (res && res.code === 0 && res.data && Array.isArray(res.data.list)) {
      console.log('学生列表原始数据:', res.data.list)
      
      // 确保学生列表不为空
      if (res.data.list.length > 0) {
        // 清空默认选项，使用API返回的学生数据
        studentOptions.value = res.data.list.map(student => {
          // 优先使用open_id，如果没有则尝试使用_openid
          const studentId = student.open_id || student._openid || ''
          return {
            label: student.name || `学生${student.id || studentId}`,
            value: studentId
          }
        }).filter(option => option.value) // 过滤掉无效的选项
        
        console.log('处理后的学生选项:', studentOptions.value)
        
        // 如果过滤后仍然为空，使用原始数据尝试提取id
        if (studentOptions.value.length === 0) {
          studentOptions.value = res.data.list.map(student => ({
            label: student.name || `学生${student.id}`,
            value: student.id ? student.id.toString() : ''
          })).filter(option => option.value)
          
          console.log('使用id作为备选的学生选项:', studentOptions.value)
        }
      }
    } else {
      console.error('获取学生列表失败: 响应格式不正确', res)
      message.error('获取学生列表失败')
    }
  } catch (error) {
    console.error('获取学生列表失败:', error)
    message.error('获取学生列表失败')
  }
}

// Handle table change (pagination, filters, sorter)
const handleTableChange = (pag, filters, sorter) => {
  pagination.current = pag.current
  pagination.pageSize = pag.pageSize
  
  // Handle sorting
  if (sorter.field && sorter.order) {
    queryParams.sortField = sorter.field
    queryParams.sortOrder = sorter.order === 'ascend' ? 'asc' : 'desc'
  } else {
    queryParams.sortField = undefined
    queryParams.sortOrder = undefined
  }
  
  fetchTaskList()
}

// Search
const handleSearch = () => {
  pagination.current = 1
  fetchTaskList()
}

// Reset search form
const resetQuery = () => {
  Object.keys(queryParams).forEach(key => {
    if (key !== 'pageNum' && key !== 'pageSize') {
      queryParams[key] = undefined
    }
  })
  pagination.current = 1
  fetchTaskList()
}

// Add new task
const handleAdd = () => {
  isEdit.value = false
  formTitle.value = '新增任务'
  resetForm()
  drawerVisible.value = true
}

// Edit task
const handleEdit = async (record) => {
  isEdit.value = true
  formTitle.value = '编辑任务'
  resetForm()
  
  try {
    const res = await getTaskApi({ id: record.id })
    if (res.code === 0 && res.data) {
      const taskData = res.data
      
      // Convert timestamps to moments for date pickers
      formData.id = taskData.id
      formData.english_text = taskData.english_text
      formData.chinese_text = taskData.chinese_text
      formData.startTimeMoment = taskData.start_time ? dayjs(taskData.start_time) : null
      formData.endTimeMoment = taskData.end_time ? dayjs(taskData.end_time) : null
      formData.start_time = taskData.start_time
      formData.end_time = taskData.end_time
      
      // Parse students JSON if it's a string
      if (typeof taskData.students === 'string') {
        try {
          formData.students = JSON.parse(taskData.students)
        } catch (e) {
          formData.students = []
        }
      } else {
        formData.students = taskData.students || []
      }
      
      formData.difficulty_level = taskData.difficulty_level
      formData.task_type = taskData.task_type
      
      drawerVisible.value = true
    } else {
      message.error(res.msg || '获取任务详情失败')
    }
  } catch (error) {
    console.error('获取任务详情失败:', error)
    message.error('获取任务详情失败')
  }
}

// Delete task
const handleDelete = async (record) => {
  try {
    const res = await deleteTaskApi({ id: record.id })
    if (res.code === 0) {
      message.success('删除成功')
      fetchTaskList()
    } else {
      message.error(res.msg || '删除失败')
    }
  } catch (error) {
    console.error('删除任务失败:', error)
    message.error('删除失败')
  }
}

// Export tasks
const handleExport = async () => {
  try {
    const res = await exportTasksApi(queryParams)
    if (res) {
      const blob = new Blob([res], { type: 'application/vnd.ms-excel' })
      const link = document.createElement('a')
      link.href = window.URL.createObjectURL(blob)
      link.download = `任务列表_${dayjs().format('YYYY-MM-DD')}.xlsx`
      link.click()
      window.URL.revokeObjectURL(link.href)
    }
  } catch (error) {
    console.error('导出任务失败:', error)
    message.error('导出失败')
  }
}

// Import tasks
const handleImport = () => {
  message.info('导入功能待实现')
}

// Publish task
const handlePublish = async (record) => {
  try {
    const res = await publishTaskApi({ id: record.id })
    if (res.code === 0) {
      message.success('发布成功')
      fetchTaskList()
    } else {
      message.error(res.msg || '发布失败')
    }
  } catch (error) {
    console.error('发布任务失败:', error)
    message.error('发布失败')
  }
}

// Cancel task
const handleCancel = async (record) => {
  try {
    const res = await cancelTaskApi({ id: record.id })
    if (res.code === 0) {
      message.success('取消成功')
      fetchTaskList()
    } else {
      message.error(res.msg || '取消失败')
    }
  } catch (error) {
    console.error('取消任务失败:', error)
    message.error('取消失败')
  }
}

// Complete task
const handleComplete = async (record) => {
  try {
    const res = await completeTaskApi({ id: record.id })
    if (res.code === 0) {
      message.success('标记完成成功')
      fetchTaskList()
    } else {
      message.error(res.msg || '标记完成失败')
    }
  } catch (error) {
    console.error('标记任务完成失败:', error)
    message.error('标记完成失败')
  }
}

// View task completion stats
const handleViewStats = async (record) => {
  statsLoading.value = true
  currentTask.value = record
  statsDrawerVisible.value = true
  
  try {
    const res = await getTaskCompletionStatsApi({ id: record.id })
    if (res.code === 0 && res.data) {
      completionStats.value = res.data
    } else {
      message.error(res.msg || '获取任务统计失败')
    }
  } catch (error) {
    console.error('获取任务统计失败:', error)
    message.error('获取任务统计失败')
  } finally {
    statsLoading.value = false
  }
}

// Submit form
const submitForm = async () => {
  if (!formRef.value) return
  
  try {
    await formRef.value.validate()
    
    // Convert moment objects to timestamps
    if (formData.startTimeMoment) {
      formData.start_time = formData.startTimeMoment.valueOf()
    }
    
    if (formData.endTimeMoment) {
      formData.end_time = formData.endTimeMoment.valueOf()
    }
    
    // Convert students array to JSON string
    const submitData = { ...formData }
    if (Array.isArray(submitData.students)) {
      submitData.students = JSON.stringify(submitData.students)
    }
    
    submitLoading.value = true
    
    // Create or update task
    const api = isEdit.value ? updateTaskApi : createTaskApi
    const res = await api(submitData)
    
    if (res.code === 0) {
      message.success(isEdit.value ? '更新成功' : '创建成功')
      closeDrawer()
      fetchTaskList()
    } else {
      message.error(res.msg || (isEdit.value ? '更新失败' : '创建失败'))
    }
  } catch (error) {
    console.error(isEdit.value ? '更新任务失败:' : '创建任务失败:', error)
    message.error(isEdit.value ? '更新失败' : '创建失败')
  } finally {
    submitLoading.value = false
  }
}

// Close drawer and reset form
const closeDrawer = () => {
  drawerVisible.value = false
  resetForm()
}

// Close stats drawer
const closeStatsDrawer = () => {
  statsDrawerVisible.value = false
  currentTask.value = {}
  completionStats.value = {
    completionRate: 0,
    totalStudents: 0,
    completedCount: 0,
    pendingCount: 0,
    studentList: []
  }
}

// Reset form
const resetForm = () => {
  if (formRef.value) {
    formRef.value.resetFields()
  }
  
  formData.id = undefined
  formData.english_text = ''
  formData.chinese_text = ''
  formData.startTimeMoment = null
  formData.endTimeMoment = null

  formData.start_time = 0
  formData.end_time = 0

  formData.students = []
  formData.difficulty_level = 3
  formData.task_type = 'reading'
}

// Format date
const formatDate = (timestamp) => {
  if (!timestamp) return '-'
  
  try {
    // 处理不同类型的时间戳
    let time = timestamp
    
    // 如果是字符串但可以转换为数字，则转换为数字
    if (typeof timestamp === 'string' && !isNaN(Number(timestamp))) {
      time = Number(timestamp)
    }
    
    // 如果是13位时间戳（毫秒），直接使用
    // 如果是10位时间戳（秒），转换为毫秒
    if (typeof time === 'number' && time.toString().length === 10) {
      time = time * 1000
    }
    
    return dayjs(time).format('YYYY-MM-DD HH:mm:ss')
  } catch (error) {
    console.error('格式化日期失败:', error, timestamp)
    return timestamp ? timestamp.toString() : '-'
  }
}

// Disable end date before start date
const disabledEndDate = (current) => {
  return formData.startTimeMoment && current && current < formData.startTimeMoment
}

// Get task type text
const getTaskTypeText = (type) => {
  const typeMap = {
    'reading': '阅读',
    'pronunciation': '发音',
    'writing': '写作',
    'listening': '听力'
  }
  return typeMap[type] || type
}

// Get task type color
const getTaskTypeColor = (type) => {
  const colorMap = {
    'reading': 'blue',
    'pronunciation': 'green',
    'writing': 'purple',
    'listening': 'orange'
  }
  return colorMap[type] || 'default'
}

// Check if task can be published
const canPublish = (record) => {
  // Add your logic here based on task status
  return true
}

// Check if task can be cancelled
const canCancel = (record) => {
  // Add your logic here based on task status
  return true
}

// Check if task can be completed
const canComplete = (record) => {
  // Add your logic here based on task status
  return true
}

// Lifecycle hooks
onMounted(() => {
  fetchTaskList()
  fetchStudentOptions()
})
</script>

<style scoped>
.task-management {
  padding: 24px;
}

.search-form {
  margin-bottom: 24px;
}

.table-operations {
  margin-bottom: 16px;
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
</style>
