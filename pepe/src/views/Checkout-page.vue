<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { useCartStore } from '@/stores/cart'
import { api } from '@/api'

const router = useRouter()
const authStore = useAuthStore()
const cartStore = useCartStore()

const address = ref('')
const deliveryDate = ref('')
const isSubmitting = ref(false)

onMounted(async () => {
  if (!authStore.isAuth) {
    router.push('/auth')
    return
  }
  await cartStore.fetchCart()
  if (cartStore.items.length === 0) {
    router.push('/cart')
  }
})

async function submitOrder() {
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
  } catch (err: any) {
    alert(err.message || 'Ошибка оформления заказа')
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
        v-for="item in cartStore.items"
        :key="item.id"
        style="
          display: flex;
          justify-content: space-between;
          padding: 8px 0;
          border-bottom: 1px solid #eee;
        "
      >
        <span>{{ item.product?.title }} × {{ item.count }}</span>
        <span>{{ Number((item.product?.cost || 0) * item.count).toLocaleString('ru-RU') }} ₽</span>
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
      <input v-model="deliveryDate" type="date" required />
      <button type="submit" :disabled="isSubmitting" style="width: 100%; margin-top: 15px">
        {{ isSubmitting ? 'Оформление...' : 'Подтвердить заказ ✓' }}
      </button>
    </form>
  </main>
</template>
