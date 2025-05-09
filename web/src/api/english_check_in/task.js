export function getTasksApi(params) {
  return useGet('/v1/english_check_in/tasks', params)
}

export function getTaskApi(params) {
  return useGet('/v1/english_check_in/task', params)
}

export function createTaskApi(params) {
  return usePost('/v1/english_check_in/task', params)
}

export function updateTaskApi(params) {
  return usePut('/v1/english_check_in/task', params)
}

export function deleteTaskApi(params) {
  return useDelete('/v1/english_check_in/task', params)
}

// 发布任务
export function publishTaskApi(params) {
  return usePut('/v1/english_check_in/task/publish', params)
}

// 取消任务
export function cancelTaskApi(params) {
  return usePut('/v1/english_check_in/task/cancel', params)
}

// 完成任务
export function completeTaskApi(params) {
  return usePut('/v1/english_check_in/task/complete', params)
}

// 获取任务完成情况统计
export function getTaskCompletionStatsApi(params) {
  return useGet('/v1/english_check_in/task/completion-stats', params)
}

// 批量导入任务
export function importTasksApi(params) {
  return usePost('/v1/english_check_in/tasks/import', params)
}

// 导出任务数据
export function exportTasksApi(params) {
  return useGet('/v1/english_check_in/tasks/export', params, {
    responseType: 'blob'
  })
}
