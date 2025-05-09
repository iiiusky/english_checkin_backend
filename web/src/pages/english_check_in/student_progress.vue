<template>
  <div class="student-progress-management">
    <a-card :bordered="false">
      <!-- Search Form -->
      <a-form layout="inline" :model="queryParams" class="search-form">
        <a-form-item label="学生">
          <a-select
            v-model:value="queryParams.studentId"
            placeholder="请选择学生"
            allowClear
            style="width: 180px"
            :options="studentOptions"
          ></a-select>
        </a-form-item>
        <a-form-item label="任务">
          <a-select
            v-model:value="queryParams.taskId"
            placeholder="请选择任务"
            allowClear
            style="width: 180px"
            :options="taskOptions"
          ></a-select>
        </a-form-item>
        <a-form-item label="任务类型">
          <a-select v-model:value="queryParams.taskType" placeholder="请选择任务类型" allowClear style="width: 120px">
            <a-select-option value="reading">阅读</a-select-option>
            <a-select-option value="pronunciation">发音</a-select-option>
            <a-select-option value="writing">写作</a-select-option>
            <a-select-option value="listening">听力</a-select-option>
          </a-select>
        </a-form-item>
        <a-form-item label="教师点评">
          <a-select v-model:value="queryParams.hasFeedback" placeholder="请选择" allowClear style="width: 120px">
            <a-select-option :value="true">已点评</a-select-option>
            <a-select-option :value="false">未点评</a-select-option>
          </a-select>
        </a-form-item>
        <a-form-item label="时间范围">
          <a-range-picker
            v-model:value="dateRange"
            :show-time="{ format: 'HH:mm' }"
            format="YYYY-MM-DD HH:mm"
            style="width: 380px"
          />
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
          <a-button @click="handleExport">
            <template #icon><ExportOutlined /></template>
            导出
          </a-button>
        </a-space>
      </div>

      <!-- Progress Table -->
      <a-table
        :columns="columns"
        :data-source="progressList"
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
          
          <!-- AI Score Column -->
          <template v-if="column.dataIndex === 'aiScore'">
            <a-progress
              :percent="record.ai_score"
              :stroke-color="getScoreColor(record.ai_score)"
              size="small"
              :format="(percent) => `${percent.toFixed(2)}`"
            />
          </template>
          
          <!-- Teacher Score Column -->
          <template v-if="column.dataIndex === 'teacherScore'">
            <template v-if="record.teacher_score > 0">
              <a-progress
                :percent="record.teacher_score"
                :stroke-color="getScoreColor(record.teacher_score)"
                size="small"
                :format="(percent) => `${percent.toFixed(2)}`"
              />
            </template>
            <template v-else>
              <a-tag color="orange">未评分</a-tag>
            </template>
          </template>
          
          <!-- Check-in Time Column -->
          <template v-if="column.dataIndex === 'checkInAt'">
            {{ formatDate(record.checkin_at) }}
          </template>
          
          <!-- Actions Column -->
          <template v-if="column.dataIndex === 'action'">
            <a-space>
              <a @click="handleViewDetails(record)">详情</a>
              <a-divider type="vertical" />
              <a @click="handleAddFeedback(record)" v-if="!hasFeedback(record)">点评</a>
              <a @click="handleEditFeedback(record)" v-else>修改点评</a>
            </a-space>
          </template>
        </template>
      </a-table>
    </a-card>

    <!-- Progress Details Drawer -->
    <a-drawer
      title="进度详情"
      :visible="detailsDrawerVisible"
      :width="600"
      @close="closeDetailsDrawer"
    >
      <a-spin :spinning="detailsLoading">
        <a-descriptions bordered :column="1" size="small">
          <a-descriptions-item label="学生姓名">
            {{ currentProgress.student_name || '未知' }}
          </a-descriptions-item>
          <a-descriptions-item label="任务内容">
            {{ currentProgress.task_english_text || '未知' }}
          </a-descriptions-item>
          <a-descriptions-item label="任务类型">
            <a-tag :color="getTaskTypeColor(currentProgress.task_type)">
              {{ getTaskTypeText(currentProgress.task_type) }}
            </a-tag>
          </a-descriptions-item>
          <a-descriptions-item label="打卡时间">
            {{ formatDate(currentProgress.checkin_at) }}
          </a-descriptions-item>
          <a-descriptions-item label="音频">
            <audio v-if="currentProgress.audio_url" controls :src="currentProgress.audio_url"></audio>
            <span v-else>无音频</span>
          </a-descriptions-item>
          <a-descriptions-item label="AI评分">
            <a-row>
              <a-col :span="12">
                <a-statistic 
                  title="总分" 
                  :value="currentProgress.ai_score" 
                  :precision="2" 
                  suffix="分"
                />
              </a-col>
              <a-col :span="12">
                <a-statistic 
                  title="准确度" 
                  :value="getAiSuggestionScore('pronunciation')" 
                  :precision="2" 
                  suffix="分"
                />
              </a-col>
            </a-row>
            <a-row style="margin-top: 12px;">
              <a-col :span="12">
                <a-statistic 
                  title="流利度" 
                  :value="getAiSuggestionScore('fluency')" 
                  :precision="2" 
                  suffix="分"
                />
              </a-col>
              <a-col :span="12">
                <a-statistic 
                  title="完整度" 
                  :value="getAiSuggestionScore('completion')" 
                  :precision="2" 
                  suffix="分"
                />
              </a-col>
            </a-row>
          </a-descriptions-item>
          <a-descriptions-item label="AI建议">
            <div v-for="(suggestion, index) in parsedAiSuggestions" :key="index" style="margin-bottom: 8px;">
              <div><strong>{{ getSuggestionTypeText(suggestion.type) }}:</strong> {{ suggestion.score.toFixed(2) }} 分</div>
              <div>{{ suggestion.suggestion }}</div>
            </div>
            <div v-if="!parsedAiSuggestions.length">无建议</div>
          </a-descriptions-item>
          <a-descriptions-item label="教师评分">
            <template v-if="currentProgress.teacher_score > 0">
              <a-statistic 
                :value="currentProgress.teacher_score" 
                :precision="2" 
                suffix="分"
              />
            </template>
            <template v-else>
              <span>未评分</span>
            </template>
          </a-descriptions-item>
          <a-descriptions-item label="教师反馈">
            {{ currentProgress.teacher_feedback || '无反馈' }}
          </a-descriptions-item>
        </a-descriptions>
      </a-spin>
    </a-drawer>

    <!-- Feedback Form Drawer -->
    <a-drawer
      :title="feedbackFormTitle"
      :visible="feedbackDrawerVisible"
      :width="700"
      @close="closeFeedbackDrawer"
      :body-style="{ paddingBottom: '80px' }"
    >
      <!-- Student Progress Information -->
      <a-card title="学生进度信息" :bordered="false" style="margin-bottom: 24px;">
        <a-descriptions :column="1" size="small">
          <a-descriptions-item label="学生姓名">
            {{ feedbackProgress.student_name || '未知' }}
          </a-descriptions-item>
          <a-descriptions-item label="任务内容">
            {{ feedbackProgress.task_english_text || '未知' }}
          </a-descriptions-item>
          <a-descriptions-item label="音频">
            <audio v-if="feedbackProgress.audio_url" controls :src="feedbackProgress.audio_url" style="width: 100%"></audio>
            <span v-else>无音频</span>
          </a-descriptions-item>
          <a-descriptions-item label="AI评分">
            <a-progress
              :percent="feedbackProgress.ai_score"
              :stroke-color="getScoreColor(feedbackProgress.ai_score)"
              :format="(percent) => `${percent.toFixed(2)}`"
            />
          </a-descriptions-item>
          <a-descriptions-item label="AI建议">
            <div v-for="(suggestion, index) in parsedFeedbackSuggestions" :key="index" style="margin-bottom: 8px;">
              <div><strong>{{ getSuggestionTypeText(suggestion.type) }}:</strong> {{ suggestion.score.toFixed(2) }} 分</div>
              <div>{{ suggestion.suggestion }}</div>
            </div>
            <div v-if="!parsedFeedbackSuggestions.length">无建议</div>
          </a-descriptions-item>
        </a-descriptions>
      </a-card>
      
      <!-- Teacher Feedback Form -->
      <a-card title="教师点评" :bordered="false">
        <a-form
          :model="feedbackForm"
          :rules="feedbackRules"
          ref="feedbackFormRef"
          layout="vertical"
        >
          <a-form-item name="teacher_score" label="教师评分" required>
            <a-slider
              v-model:value="feedbackForm.teacher_score"
              :min="0"
              :max="100"
              :step="1"
              :tooltip-visible="true"
            />
          </a-form-item>
          <a-form-item name="teacher_feedback" label="教师反馈" required>
            <a-textarea 
              v-model:value="feedbackForm.teacher_feedback" 
              placeholder="请输入教师反馈" 
              :auto-size="{ minRows: 3, maxRows: 6 }" 
            />
          </a-form-item>
        </a-form>
      </a-card>
      
      <div class="drawer-footer">
        <a-space>
          <a-button @click="closeFeedbackDrawer">取消</a-button>
          <a-button type="primary" @click="submitFeedbackForm" :loading="submitLoading">提交</a-button>
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
  ExportOutlined
} from '@ant-design/icons-vue'
import { 
  getStudentProgressListApi, 
  getStudentProgressApi,
  getStudentProgressByIdApi,
  submitTeacherFeedbackApi,
  exportProgressDataApi
} from '~@/api/english_check_in/progress'
import { getStudentsApi } from '~@/api/english_check_in/student'
import { getTasksApi } from '~@/api/english_check_in/task'
import dayjs from 'dayjs'

// Table columns
const columns = [
  {
    title: '学生',
    dataIndex: 'student_name',
    width: 120,
    ellipsis: true,
  },
  {
    title: 'AI评分',
    dataIndex: 'aiScore',
    width: 150,
    sorter: true,
  },
  {
    title: '教师评分',
    dataIndex: 'teacherScore',
    width: 150,
    sorter: true,
  },
  {
    title: '打卡时间',
    dataIndex: 'checkInAt',
    width: 160,
    sorter: true,
  },
  {
    title: '操作',
    dataIndex: 'action',
    width: 120,
    fixed: 'right',
  },
]

// State
const loading = ref(false)
const progressList = ref([])
const studentOptions = ref([])
const taskOptions = ref([])
const dateRange = ref(null)
const detailsDrawerVisible = ref(false)
const detailsLoading = ref(false)
const currentProgress = ref({})
const feedbackDrawerVisible = ref(false)
const feedbackProgress = ref({})
const feedbackForm = reactive({
  id: null,
  teacher_score: 0,
  teacher_feedback: '',
  teacher_id: 'admin' // 默认使用admin作为教师ID，实际应该从登录信息获取
})
const feedbackFormRef = ref(null)
const submitLoading = ref(false)

// Computed
const feedbackFormTitle = computed(() => {
  return feedbackForm.id ? '修改教师点评' : '添加教师点评'
})

// Pagination
const pagination = reactive({
  current: 1,
  pageSize: 10,
  total: 0,
  showSizeChanger: true,
  showTotal: (total) => `共 ${total} 条记录`
})

// Query parameters
const queryParams = reactive({
  studentId: undefined,
  taskId: undefined,
  taskType: undefined,
  hasFeedback: undefined,
  startDate: undefined,
  endDate: undefined
})

// Form rules
const feedbackRules = {
  teacher_score: [{ required: true, message: '请输入教师评分' }],
  teacher_feedback: [{ required: true, message: '请输入教师反馈' }]
}

// Fetch progress list
const fetchProgressList = async () => {
  loading.value = true
  try {
    // 处理日期范围
    if (dateRange.value && dateRange.value.length === 2) {
      queryParams.startDate = dateRange.value[0].valueOf()
      queryParams.endDate = dateRange.value[1].valueOf()
    } else {
      queryParams.startDate = undefined
      queryParams.endDate = undefined
    }

    const response = await getStudentProgressListApi({
      page: pagination.current,
      pageSize: pagination.pageSize,
      ...queryParams
    })
    
    progressList.value = response.data.list
    pagination.total = response.data.total
  } catch (error) {
    console.error('获取进度列表失败:', error)
    message.error('获取进度列表失败')
  } finally {
    loading.value = false
  }
}

// Fetch student options
const fetchStudentOptions = async () => {
  try {
    const response = await getStudentsApi({ page: 1, pageSize: 100 })
    studentOptions.value = response.data.list.map(student => ({
      label: student.name,
      value: student.open_id
    }))
  } catch (error) {
    console.error('获取学生列表失败:', error)
    message.error('获取学生列表失败')
  }
}

// Fetch task options
const fetchTaskOptions = async () => {
  try {
    const response = await getTasksApi({ page: 1, pageSize: 100 })
    taskOptions.value = response.data.list.map(task => ({
      label: task.english_text.substring(0, 20) + (task.english_text.length > 20 ? '...' : ''),
      value: task.id
    }))
  } catch (error) {
    console.error('获取任务列表失败:', error)
    message.error('获取任务列表失败')
  }
}

// Handle table change (pagination, filters, sorter)
const handleTableChange = (pag, filters, sorter) => {
  pagination.current = pag.current
  pagination.pageSize = pag.pageSize
  
  // 处理排序
  if (sorter && sorter.field) {
    if (sorter.field === 'aiScore') {
      queryParams.sortField = 'ai_score'
    } else if (sorter.field === 'teacherScore') {
      queryParams.sortField = 'teacher_score'
    } else if (sorter.field === 'checkInAt') {
      queryParams.sortField = 'checkin_at'
    }
    queryParams.sortOrder = sorter.order === 'ascend' ? 'asc' : 'desc'
  } else {
    queryParams.sortField = undefined
    queryParams.sortOrder = undefined
  }
  
  fetchProgressList()
}

// Search
const handleSearch = () => {
  pagination.current = 1
  fetchProgressList()
}

// Reset search form
const resetQuery = () => {
  Object.keys(queryParams).forEach(key => {
    queryParams[key] = undefined
  })
  dateRange.value = null
  pagination.current = 1
  fetchProgressList()
}

// Export progress data
const handleExport = async () => {
  try {
    message.loading('正在导出数据...', 0)
    const response = await exportProgressDataApi({
      ...queryParams,
      startDate: dateRange.value ? dateRange.value[0].valueOf() : undefined,
      endDate: dateRange.value ? dateRange.value[1].valueOf() : undefined
    })
    
    // 创建下载链接
    const url = window.URL.createObjectURL(new Blob([response]))
    const link = document.createElement('a')
    link.href = url
    link.setAttribute('download', `学生进度数据_${new Date().getTime()}.xlsx`)
    document.body.appendChild(link)
    link.click()
    document.body.removeChild(link)
    
    message.destroy()
    message.success('导出成功')
  } catch (error) {
    message.destroy()
    console.error('导出失败:', error)
    message.error('导出失败')
  }
}

// View progress details
const handleViewDetails = async (record) => {
  detailsDrawerVisible.value = true
  detailsLoading.value = true
  currentProgress.value = { ...record }
  
  try {
    // 使用基于ID的路由获取进度详情
    const response = await getStudentProgressByIdApi(record.id)
    currentProgress.value = response.data
  } catch (error) {
    console.error('获取进度详情失败:', error)
    message.error('获取进度详情失败')
  } finally {
    detailsLoading.value = false
  }
}

// Close details drawer
const closeDetailsDrawer = () => {
  detailsDrawerVisible.value = false
  currentProgress.value = {}
}

// Add teacher feedback
const handleAddFeedback = async (record) => {
  feedbackDrawerVisible.value = true
  feedbackProgress.value = { ...record }
  feedbackForm.id = record.id
  feedbackForm.teacher_score = 0
  feedbackForm.teacher_feedback = ''
  
  try {
    // 获取完整的进度详情
    const response = await getStudentProgressByIdApi(record.id)
    feedbackProgress.value = response.data
  } catch (error) {
    console.error('获取进度详情失败:', error)
    message.error('获取进度详情失败')
  }
}

// Edit teacher feedback
const handleEditFeedback = async (record) => {
  feedbackDrawerVisible.value = true
  feedbackProgress.value = { ...record }
  feedbackForm.id = record.id
  feedbackForm.teacher_score = record.teacher_score || 0
  feedbackForm.teacher_feedback = record.teacher_feedback || ''
  
  try {
    // 获取完整的进度详情
    const response = await getStudentProgressByIdApi(record.id)
    feedbackProgress.value = response.data
  } catch (error) {
    console.error('获取进度详情失败:', error)
    message.error('获取进度详情失败')
  }
}

// Submit feedback form
const submitFeedbackForm = async () => {
  try {
    await feedbackFormRef.value.validate()
    
    submitLoading.value = true
    await submitTeacherFeedbackApi({
      id: feedbackForm.id,
      teacher_score: feedbackForm.teacher_score,
      teacher_feedback: feedbackForm.teacher_feedback,
      teacher_id: feedbackForm.teacher_id
    })
    
    message.success('提交成功')
    closeFeedbackDrawer()
    fetchProgressList()
  } catch (error) {
    console.error('提交失败:', error)
    message.error('提交失败')
  } finally {
    submitLoading.value = false
  }
}

// Close feedback drawer
const closeFeedbackDrawer = () => {
  feedbackDrawerVisible.value = false
  feedbackFormRef.value?.resetFields()
}

// Format date
const formatDate = (timestamp) => {
  if (!timestamp) return '未知'
  
  // 处理不同类型的时间戳
  let date
  if (typeof timestamp === 'string' && !isNaN(Number(timestamp))) {
    date = dayjs(Number(timestamp))
  } else if (typeof timestamp === 'number') {
    date = dayjs(timestamp)
  } else {
    date = dayjs(timestamp)
  }
  
  if (!date.isValid()) return '未知'
  return date.format('YYYY-MM-DD HH:mm:ss')
}

// Get task type text
const getTaskTypeText = (type) => {
  const typeMap = {
    'reading': '阅读',
    'pronunciation': '发音',
    'writing': '写作',
    'listening': '听力'
  }
  return typeMap[type] || '未知'
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

// Get score color
const getScoreColor = (score) => {
  if (score >= 90) return '#52c41a'
  if (score >= 70) return '#1890ff'
  if (score >= 60) return '#faad14'
  return '#f5222d'
}

// Check if record has feedback
const hasFeedback = (record) => {
  return record.teacher_score > 0 && record.teacher_feedback
}

// Parse AI suggestions from JSON string
const parsedAiSuggestions = computed(() => {
  if (!currentProgress.value.ai_suggestions) return []
  try {
    return JSON.parse(currentProgress.value.ai_suggestions)
  } catch (error) {
    console.error('解析AI建议失败:', error)
    return []
  }
})

// Parse AI suggestions for feedback drawer
const parsedFeedbackSuggestions = computed(() => {
  if (!feedbackProgress.value.ai_suggestions) return []
  try {
    return JSON.parse(feedbackProgress.value.ai_suggestions)
  } catch (error) {
    console.error('解析反馈界面AI建议失败:', error)
    return []
  }
})

// Get AI suggestion score by type
const getAiSuggestionScore = (type) => {
  if (!parsedAiSuggestions.value.length) return 0
  
  const suggestion = parsedAiSuggestions.value.find(item => item.type === type)
  if (suggestion) {
    return suggestion.score
  }
  return 0
}

// Get suggestion type text
const getSuggestionTypeText = (type) => {
  const typeMap = {
    'overall': '总体评价',
    'pronunciation': '发音准确度',
    'fluency': '流利度',
    'completion': '完整度'
  }
  return typeMap[type] || type
}

// Lifecycle hooks
onMounted(() => {
  fetchProgressList()
  fetchStudentOptions()
  fetchTaskOptions()
})
</script>

<style scoped>
.search-form {
  margin-bottom: 24px;
}

.table-operations {
  margin-bottom: 16px;
}

.drawer-footer {
  position: absolute;
  right: 0;
  bottom: 0;
  width: 100%;
  border-top: 1px solid #e9e9e9;
  padding: 10px 16px;
  background: #fff;
  text-align: right;
}
</style>
