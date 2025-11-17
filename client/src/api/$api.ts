export type Query = Record<string, string | number | boolean | undefined | null>

const DATA_SOURCE_PATH = '/data.json' as const

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

async function readDataJson(): Promise<Record<string, unknown>> {
  const response = await fetch(DATA_SOURCE_PATH)
  if (!response.ok) {
    throw new Response('Failed to load API data', { status: 500 })
  }
  return response.json()
}

async function readFromMock<T>(path: string): Promise<T> {
  const data = await readDataJson()
  const value = (data as Record<string, unknown>)[path]
  if (value === undefined) {
    throw new Response('Resource not found', { status: 404 })
  }
  return value as T
}

async function get<T>(path: string, options?: { query?: Query }): Promise<T> {
  const finalPath = buildPathWithQuery(path, options?.query)
  return readFromMock<T>(finalPath)
}

async function requestJson<T>(
  method: 'POST' | 'PUT' | 'PATCH' | 'DELETE',
  path: string,
  body?: unknown
): Promise<T> {
  const response = await fetch(path, {
    method,
    headers: { 'Content-Type': 'application/json' },
    credentials: 'include',
    body: body === undefined ? undefined : JSON.stringify(body),
  })

  if (!response.ok) {
    const errorData = await response.json().catch(() => ({}))
    const message =
      (errorData as { message?: string }).message || 'Request failed'
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
