<script setup lang="ts">
import { onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { useCartStore, type CartItem } from '@/stores/cart'

const router = useRouter()
const authStore = useAuthStore()
const cartStore = useCartStore()

onMounted(async () => {
  if (!authStore.isAuth) {
    await authStore.fetchMe()
  }
  if (!authStore.isAuth) {
    router.push('/auth')
    return
  }
  await cartStore.fetchCart()
})

// Удобный хелпер только для форматирования пути картинки
function getItemImage(item: CartItem): string {
  const img = item.product?.image_url
  if (!img) return '/images/product.svg'
  return img.startsWith('/') ? img : `/images/${img}`
}
</script>

<template>
  <main class="cart-container">
    <h1 class="cart-title">Корзина</h1>

    <div v-if="cartStore.validItems.length === 0" class="empty-cart">
      <p>Корзина пуста.</p>
      <router-link to="/shop">Перейти в магазин ←</router-link>
    </div>

    <div v-else>
      <div class="cart-items-list">
        <div v-for="item in cartStore.validItems" :key="item.id" class="cart-item">
          <div class="cart-item-image">
            <img :src="getItemImage(item)" :alt="item.product.title" />
          </div>

          <div class="cart-item-details">
            <h3 class="cart-item-title">{{ item.product.title }}</h3>
            <div class="cart-item-price">
              {{ item.product.cost.toLocaleString('ru-RU') }} ₽ / шт.
            </div>
          </div>

          <div class="cart-item-actions">
            <div class="quantity-form">
              <button
                type="button"
                @click="cartStore.updateQuantity(item.product_id, item.count - 1)"
              >
                −
              </button>
              <input type="text" :value="item.count" readonly />
              <button
                type="button"
                @click="cartStore.updateQuantity(item.product_id, item.count + 1)"
              >
                +
              </button>
            </div>
            <div class="remove-form">
              <button type="button" @click="cartStore.removeFromCart(item.product_id)">✕</button>
            </div>
          </div>

          <div class="cart-item-total">
            {{ (item.product.cost * item.count).toLocaleString('ru-RU') }} ₽
          </div>
        </div>
      </div>

      <div class="cart-summary">
        <div class="cart-total">
          Итого: <span>{{ cartStore.totalPrice.toLocaleString('ru-RU') }} ₽</span>
        </div>
        <router-link to="/checkout" class="btn-checkout"> Оформить заказ → </router-link>
      </div>
    </div>
  </main>
</template>
