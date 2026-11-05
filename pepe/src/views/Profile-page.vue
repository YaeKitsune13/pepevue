<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { useOrdersStore, type OrderItem } from '@/stores/orders'

const router = useRouter()
const authStore = useAuthStore()
const ordersStore = useOrdersStore()

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

    await ordersStore.fetchOrdersByUserId(authStore.user.id)
  }
})

async function handleUpdate() {
  try {
    const updateData: {
      surname: string
      name: string
      patronymic: string
      login: string
      password?: string
    } = {
      surname: surname.value,
      name: name.value,
      patronymic: patronymic.value,
      login: login.value,
    }

    if (password.value.trim() !== '') {
      updateData.password = password.value
    }

    await authStore.updateProfile(updateData)

    password.value = ''
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

function formatDate(dateStr?: string) {
  if (!dateStr) return '—'
  return new Date(dateStr).toLocaleDateString('ru-RU')
}

function getProductTitle(item: OrderItem) {
  return item.product?.title || item.Product?.title || `Товар #${item.product_id}`
}

function getProductImage(item: OrderItem) {
  return item.product?.image_url || item.Product?.image_url || ''
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

    <section class="orders-section">
      <h2>Мои заказы</h2>

      <div v-if="ordersStore.loading">Загрузка заказов...</div>
      <div v-else-if="ordersStore.error" class="error">{{ ordersStore.error }}</div>
      <div v-else-if="ordersStore.orders.length === 0">У вас пока нет заказов</div>

      <div v-else class="orders-list">
        <div v-for="order in ordersStore.orders" :key="order.id ?? order.ID" class="order-card">
          <div class="order-header">
            <strong>Заказ №{{ order.id ?? order.ID }}</strong>
            <span class="order-status">{{ order.status }}</span>
          </div>

          <div class="order-info">
            <p><strong>Адрес доставки:</strong> {{ order.address }}</p>
            <p><strong>Дата доставки:</strong> {{ formatDate(order.date_delivery) }}</p>
            <p><strong>Сумма заказа:</strong> {{ order.price_all }} ₽</p>
          </div>

          <div v-if="order.items && order.items.length > 0" class="order-items">
            <h4>Товары в заказе:</h4>
            <div
              v-for="item in order.items"
              :key="item.id ?? item.ID ?? item.product_id"
              class="order-item"
            >
              <img
                v-if="getProductImage(item)"
                :src="getProductImage(item)"
                :alt="getProductTitle(item)"
                class="order-item-img"
              />
              <div class="order-item-details">
                <span class="order-item-title">{{ getProductTitle(item) }}</span>
                <span class="order-item-meta">
                  {{ item.count }} шт. × {{ item.cost }} ₽ =
                  <strong>{{ item.count * item.cost }} ₽</strong>
                </span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>
  </main>
</template>

<style scoped>
.orders-section {
  margin-top: 32px;
}

.orders-list {
  display: flex;
  flex-direction: column;
  gap: 16px;
  margin-top: 12px;
}

.order-card {
  border: 1px solid #ddd;
  border-radius: 8px;
  padding: 16px;
  background: #fafafa;
}

.order-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 12px;
  padding-bottom: 8px;
  border-bottom: 1px solid #eee;
}

.order-status {
  padding: 4px 10px;
  border-radius: 4px;
  background: #e0f2fe;
  color: #0369a1;
  font-size: 14px;
}

.order-info p {
  margin: 4px 0;
}

.order-items {
  margin-top: 12px;
  padding-top: 12px;
  border-top: 1px solid #eee;
}

.order-items h4 {
  margin: 0 0 8px 0;
}

.order-item {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 6px 0;
}

.order-item-img {
  width: 48px;
  height: 48px;
  object-fit: cover;
  border-radius: 4px;
}

.order-item-details {
  display: flex;
  flex-direction: column;
}

.order-item-meta {
  font-size: 14px;
  color: #555;
}

.error {
  color: red;
}
</style>
