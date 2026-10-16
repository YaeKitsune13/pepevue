<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const router = useRouter()
const authStore = useAuthStore()

const surname = ref('')
const name = ref('')
const patronymic = ref('')
const login = ref('')
const password = ref('')

onMounted(async () => {
  if (!authStore.isAuth) {
    await authStore.fetchMe()
  }
  if (!authStore.isAuth) {
    router.push('/auth')
    return
  }
  if (authStore.user) {
    surname.value = authStore.user.surname || ''
    name.value = authStore.user.name || ''
    patronymic.value = authStore.user.patronymic || ''
    login.value = authStore.user.login || ''
  }
})

async function handleUpdate() {
  try {
    await authStore.updateProfile({
      surname: surname.value,
      name: name.value,
      patronymic: patronymic.value,
    })
    alert('Данные обновлены')
  } catch (err: unknown) {
    const message = err instanceof Error ? err.message : 'Ошибка обновления'
    alert(message)
  }
}

async function handleDelete() {
  if (confirm('Вы уверены, что хотите удалить аккаунт?')) {
    await authStore.deleteAccount()
    router.push('/')
  }
}

async function handleLogout() {
  await authStore.logout()
  router.push('/auth')
}
</script>

<template>
  <main class="profile-container">
    <form @submit.prevent="handleUpdate">
      <h1>Профиль</h1>
      <input v-model="surname" type="text" placeholder="Фамилия*" required />
      <input v-model="name" type="text" placeholder="Имя*" required />
      <input v-model="patronymic" type="text" placeholder="Отчество" />
      <input v-model="login" type="text" placeholder="Логин*" required />
      <input v-model="password" type="password" placeholder="Новый пароль" />

      <button type="submit">Изменить</button>
      <button type="button" style="background-color: orange; color: white" @click="handleLogout">
        Выйти из аккаунта
      </button>
      <button type="button" style="background-color: red; color: white" @click="handleDelete">
        Удалить аккаунт
      </button>
    </form>
  </main>
</template>
