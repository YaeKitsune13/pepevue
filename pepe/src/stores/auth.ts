import { defineStore } from 'pinia'
import { ref } from 'vue'
import { api } from '@/api'

export interface User {
  id: number
  login: string
  name?: string
  surname?: string
  patronymic?: string
}

export const useAuthStore = defineStore('auth', () => {
  const user = ref<User | null>(null)
  const isAuth = ref<boolean>(false)
  const isLoading = ref<boolean>(false)

  // Проверка сессии при загрузке страницы
  async function fetchMe() {
    try {
      isLoading.value = true
      const data = await api<User>('/me')
      user.value = data
      isAuth.value = true
    } catch {
      user.value = null
      isAuth.value = false
    } finally {
      isLoading.value = false
    }
  }

  async function login(credentials: Record<string, any>) {
    await api('/login', {
      method: 'POST',
      body: JSON.stringify(credentials),
    })
    await fetchMe()
  }

  async function register(data: Record<string, any>) {
    await api('/register', {
      method: 'POST',
      body: JSON.stringify(data),
    })
    await fetchMe()
  }

  async function logout() {
    try {
      await api('/logout', { method: 'POST' })
    } finally {
      user.value = null
      isAuth.value = false
    }
  }

  async function updateProfile(data: Partial<User>) {
    const updated = await api<User>('/me', {
      method: 'PUT',
      body: JSON.stringify(data),
    })
    user.value = updated
  }

  async function deleteAccount() {
    await api('/me', { method: 'DELETE' })
    user.value = null
    isAuth.value = false
  }

  return {
    user,
    isAuth,
    isLoading,
    fetchMe,
    login,
    register,
    logout,
    updateProfile,
    deleteAccount,
  }
})
