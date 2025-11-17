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
  const value = data[path]
  if (value === undefined) {
    throw new Response('Resource not found', { status: 404 })
  }
  return value as T
}

async function get<T>(path: string, options?: { query?: Query }): Promise<T> {
  const finalPath = buildPathWithQuery(path, options?.query)
  return readFromMock<T>(finalPath)
}

// For now, POST/PUT/PATCH/DELETE read from mock as well to keep UI working
async function post<T>(path: string, _body?: unknown): Promise<T> {
  return readFromMock<T>(path)
}

async function put<T>(path: string, _body?: unknown): Promise<T> {
  return readFromMock<T>(path)
}

async function patch<T>(path: string, _body?: unknown): Promise<T> {
  return readFromMock<T>(path)
}

async function del<T>(path: string): Promise<T> {
  return readFromMock<T>(path)
}

export const $api = {
  get,
  post,
  put,
  patch,
  delete: del,
} as const
