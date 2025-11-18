import { API_BASE_URL } from '../utils/constants'

export type Query = Record<string, string | number | boolean | undefined | null>

const buildPathWithQuery = (path: string, query?: Query): string => {
  if (!query) return path
  const params = new URLSearchParams()
  Object.entries(query).forEach(([key, value]) => {
    if (value === undefined || value === null) return
    params.append(key, String(value))
  })
  const qs = params.toString()
  return qs ? `${path}?${qs}` : path
}

const buildUrl = (path: string): string => {
  const baseUrl = API_BASE_URL.replace(/\/$/, '')
  const cleanPath = path.startsWith('/') ? path : `/${path}`
  return `${baseUrl}${cleanPath}`
}

async function get<T>(path: string, options?: { query?: Query }): Promise<T> {
  const finalPath = buildPathWithQuery(path, options?.query)
  const url = buildUrl(finalPath)

  const response = await fetch(url, {
    method: 'GET',
    credentials: 'include',
    headers: {
      'Content-Type': 'application/json',
    },
  })

  if (!response.ok) {
    const errorData = await response.json().catch(() => ({}))
    const message =
      (errorData as { error?: string; message?: string }).error ||
      (errorData as { message?: string }).message ||
      'Request failed'
    throw new Error(message)
  }

  if (response.status === 204) return undefined as unknown as T
  return (await response.json()) as T
}

async function requestJson<T>(
  method: 'POST' | 'PUT' | 'PATCH' | 'DELETE',
  path: string,
  body?: unknown
): Promise<T> {
  const url = buildUrl(path)

  const response = await fetch(url, {
    method,
    headers: { 'Content-Type': 'application/json' },
    credentials: 'include',
    body: body === undefined ? undefined : JSON.stringify(body),
  })

  if (!response.ok) {
    const errorData = await response.json().catch(() => ({}))
    const message =
      (errorData as { error?: string; message?: string }).error ||
      (errorData as { message?: string }).message ||
      'Request failed'
    throw new Error(message)
  }

  // allow 204
  if (response.status === 204) return undefined as unknown as T
  return (await response.json()) as T
}

async function post<T>(path: string, body?: unknown): Promise<T> {
  return requestJson<T>('POST', path, body)
}

async function put<T>(path: string, body?: unknown): Promise<T> {
  return requestJson<T>('PUT', path, body)
}

async function patch<T>(path: string, body?: unknown): Promise<T> {
  return requestJson<T>('PATCH', path, body)
}

async function del<T>(path: string): Promise<T> {
  return requestJson<T>('DELETE', path)
}

export const $api = {
  get,
  post,
  put,
  patch,
  delete: del,
} as const

export type ApiClient = typeof $api
