// src/api.ts
// Теперь запросы идут через Vite Proxy на бэкенд
const BASE_URL = '/api'

export async function api<T>(endpoint: string, options: RequestInit = {}): Promise<T> {
  const response = await fetch(`${BASE_URL}${endpoint}`, {
    ...options,
    credentials: 'include',
    headers: {
      'Content-Type': 'application/json',
      ...(options.headers || {}),
    },
  })

  if (!response.ok) {
    const errorData = await response.json().catch(() => ({}))
    throw new Error(errorData.message || errorData.error || `Ошибка: ${response.status}`)
  }

  if (response.status === 204) {
    return {} as T
  }

  return response.json()
}
