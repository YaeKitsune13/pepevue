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

// Функция для переключения режима с очисткой полей и ошибок
function toggleMode(value: boolean) {
  isRegister.value = value
  error.value = ''
  login.value = ''
  password.value = ''
  name.value = ''
  surname.value = ''
  patronymic.value = ''
}

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
    error.value = err instanceof Error ? err.message : 'Ошибка входа'
  }
}
</script>

<template>
  <main class="form-page-container">
    <form @submit.prevent="handleSubmit">
      <h1>{{ isRegister ? 'Регистрация' : 'Авторизация' }}</h1>

      <p v-if="error" style="color: red; text-align: center">{{ error }}</p>

      <template v-if="isRegister">
        <input v-model="surname" type="text" placeholder="Фамилия*" required />
        <input v-model="name" type="text" placeholder="Имя*" required />
        <input v-model="patronymic" type="text" placeholder="Отчество" />
      </template>

      <input v-model="login" type="text" placeholder="Логин*" required />
      <input v-model="password" type="password" placeholder="Пароль*" required />

      <p v-if="isRegister">
        <label>
          <input type="checkbox" required />
          Согласие на обработку данных
        </label>
      </p>

      <button type="submit">{{ isRegister ? 'Зарегистрироваться' : 'Войти' }}</button>

      <p>
        <template v-if="!isRegister">
          Нет аккаунта? <a href="#" @click.prevent="toggleMode(true)">Зарегистрироваться</a>
        </template>
        <template v-else>
          Есть аккаунт? <a href="#" @click.prevent="toggleMode(false)">Войти</a>
        </template>
      </p>
    </form>
  </main>
</template>
