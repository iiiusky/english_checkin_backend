export function getStudentsApi(params) {
  return useGet('/v1/english_check_in/students', params)
}

export function getStudentApi(params) {
  return useGet('/v1/english_check_in/student', params)
}

export function createStudentApi(params) {
  return usePost('/v1/english_check_in/student', params)
}

export function updateStudentApi(params) {
  return usePut('/v1/english_check_in/student', params)
}

export function deleteStudentApi(params) {
  return useDelete('/v1/english_check_in/student', params)
}

// 批量导入学生
export function importStudentsApi(params) {
  return usePost('/v1/english_check_in/students/import', params)
}

// 导出学生数据
export function exportStudentsApi(params) {
  return useGet('/v1/english_check_in/students/export', params, {
    responseType: 'blob'
  })
}
