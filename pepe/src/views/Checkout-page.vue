<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { useCartStore, type CartItem } from '@/stores/cart'
import { api } from '@/api'

const router = useRouter()
const authStore = useAuthStore()
const cartStore = useCartStore()

const address = ref('')
const deliveryDate = ref('')
const isSubmitting = ref(false)

const minDeliveryDate = computed(() => {
  const date = new Date()
  date.setDate(date.getDate() + 1)

  const year = date.getFullYear()
  const month = String(date.getMonth() + 1).padStart(2, '0')
  const day = String(date.getDate()).padStart(2, '0')

  return `${year}-${month}-${day}`
})

onMounted(async () => {
  if (!authStore.isAuth) {
    await authStore.fetchMe()
  }
  if (!authStore.isAuth) {
    router.push('/auth')
    return
  }
  await cartStore.fetchCart()
  if (cartStore.validItems.length === 0) {
    router.push('/cart')
  }
  deliveryDate.value = minDeliveryDate.value
})

function getItemTitle(item: CartItem): string {
  return item.product?.title || item.title || 'Товар'
}

function getItemCost(item: CartItem): number {
  return Number(item.product?.cost ?? item.cost ?? 0)
}

async function submitOrder() {
  if (deliveryDate.value < minDeliveryDate.value) {
    alert('Дата доставки не может быть раньше завтрашнего дня')
    return
  }

  try {
    isSubmitting.value = true
    await api('/order', {
      method: 'POST',
      body: JSON.stringify({
        address: address.value,
        delivery_date: deliveryDate.value,
      }),
    })
    alert('Заказ успешно оформлен!')
    await cartStore.fetchCart()
    router.push('/')
  } catch (err: unknown) {
    const message = err instanceof Error ? err.message : 'Ошибка оформления заказа'
    alert(message)
  } finally {
    isSubmitting.value = false
  }
}
</script>

<template>
  <main style="padding: 20px; max-width: 800px; margin: 0 auto">
    <div style="background-color: #fff; border-radius: 20px; padding: 30px; margin-bottom: 20px">
      <h2>Ваш заказ:</h2>
      <div
        v-for="item in cartStore.validItems"
        :key="item.id"
        style="
          display: flex;
          justify-content: space-between;
          padding: 8px 0;
          border-bottom: 1px solid #eee;
        "
      >
        <span>{{ getItemTitle(item) }} × {{ item.count }}</span>
        <span>{{ (getItemCost(item) * item.count).toLocaleString('ru-RU') }} ₽</span>
      </div>
      <div style="text-align: right; padding-top: 15px">
        <strong>Итого: {{ cartStore.totalPrice.toLocaleString('ru-RU') }} ₽</strong>
      </div>
    </div>

    <form @submit.prevent="submitOrder" style="max-width: 100%">
      <h2>Данные доставки:</h2>
      <input type="text" :value="authStore.user?.surname" placeholder="Фамилия*" disabled />
      <input type="text" :value="authStore.user?.name" placeholder="Имя*" disabled />
      <input v-model="address" type="text" placeholder="Адрес доставки*" required />

      <!-- Атрибут :min блокирует выбор более ранних дат в календаре -->
      <input v-model="deliveryDate" type="date" :min="minDeliveryDate" required />

      <button type="submit" :disabled="isSubmitting" style="width: 100%; margin-top: 15px">
        {{ isSubmitting ? 'Оформление...' : 'Подтвердить заказ ✓' }}
      </button>
    </form>
  </main>
</template>
