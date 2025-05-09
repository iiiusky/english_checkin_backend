export function getStudentProgressListApi(params) {
  return useGet('/v1/english_check_in/student_progress/list', params)
}

export function getStudentProgressApi(params) {
  return useGet('/v1/english_check_in/student_progress', params)
}

export function getStudentProgressByIdApi(id) {
  return useGet(`/v1/english_check_in/student_progress/${id}`)
}

export function createStudentProgressApi(params) {
  return usePost('/v1/english_check_in/student_progress', params)
}

export function updateStudentProgressApi(params) {
  return usePut('/v1/english_check_in/student_progress', params)
}

export function deleteStudentProgressApi(params) {
  return useDelete('/v1/english_check_in/student_progress', params)
}

// 获取学生进度统计摘要
export function getProgressSummaryApi(params) {
  return useGet('/v1/english_check_in/progress/summary', params)
}

// 获取学生进度趋势数据
export function getProgressTrendApi(params) {
  return useGet('/v1/english_check_in/progress/trend', params)
}

// 获取学生详细进度数据
export function getStudentDetailProgressApi(params) {
  return useGet('/v1/english_check_in/progress/student-detail', params)
}

// 教师评分和反馈
export function submitTeacherFeedbackApi(params) {
  return usePut('/v1/english_check_in/progress/teacher-feedback', params)
}

// 获取待教师点评的进度列表
export function getPendingFeedbackListApi(params) {
  return useGet('/v1/english_check_in/progress/pending-feedback', params)
}

// 导出学生进度数据
export function exportProgressDataApi(params) {
  return useGet('/v1/english_check_in/progress/export', params, {
    responseType: 'blob'
  })
}

// 获取班级进度统计
export function getClassProgressApi(params) {
  return useGet('/v1/english_check_in/progress/class', params)
}

// 获取单个任务的完成进度统计
export function getTaskProgressApi(params) {
  return useGet('/v1/english_check_in/progress/task', params)
}
