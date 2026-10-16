<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const router = useRouter()
const authStore = useAuthStore()

const isRegister = ref(false)
const login = ref('')
const password = ref('')
const name = ref('')
const surname = ref('')
const patronymic = ref('')
const error = ref('')

async function handleSubmit() {
  error.value = ''
  try {
    if (isRegister.value) {
      await authStore.register({
        login: login.value,
        password: password.value,
        name: name.value,
        surname: surname.value,
        patronymic: patronymic.value,
      })
    } else {
      await authStore.login({
        login: login.value,
        password: password.value,
      })
    }
    router.push('/')
  } catch (err: unknown) {
    const message = err instanceof Error ? err.message : 'Ошибка входа'
    error.value = message
  }
}
</script>

<template>
  <main class="form-page-container">
    <form @submit.prevent="handleSubmit">
      <h1>{{ isRegister ? 'Регистрация' : 'Авторизация' }}</h1>

      <p v-if="error" style="color: red; justify-content: center">{{ error }}</p>

      <template v-if="isRegister">
        <input v-model="surname" type="text" placeholder="Фамилия*" required />
        <input v-model="name" type="text" placeholder="Имя*" required />
        <input v-model="patronymic" type="text" placeholder="Отчество" />
      </template>

      <input v-model="login" type="text" placeholder="Логин*" required />
      <input v-model="password" type="password" placeholder="Пароль*" required />

      <p v-if="isRegister">
        Согласие на обработку данных
        <input type="checkbox" required />
      </p>

      <button type="submit">{{ isRegister ? 'Зарегистрироваться' : 'Войти' }}</button>

      <p v-if="!isRegister">
        Нет аккаунта?
        <a href="#" @click.prevent="isRegister = true">Зарегистрироваться</a>
      </p>
      <p v-else>
        Есть аккаунт?
        <a href="#" @click.prevent="isRegister = false">Войти</a>
      </p>
    </form>
  </main>
</template>
