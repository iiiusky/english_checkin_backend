import { get } from 'lodash-es'
import router from '@/router'

export function getQueryParam(param, defaultVal = '') {
  const query = router.currentRoute.value?.query ?? {}
  const val = get(query, param) ?? defaultVal
  return decodeURIComponent(val)
}
