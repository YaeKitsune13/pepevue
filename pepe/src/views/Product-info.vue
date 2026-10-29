<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { api } from '@/api'
import { useAuthStore } from '@/stores/auth'
import { useCartStore } from '@/stores/cart'

interface Product {
  id: number
  title: string
  image: string
  description: string
  cost: number
  count: number
}

interface ProductResponse {
  product?: Product
  data?: Product
  item?: Product
}

const route = useRoute()
const router = useRouter()
const authStore = useAuthStore()
const cartStore = useCartStore()

const product = ref<Product | null>(null)
const count = ref(1)
const loading = ref(true)
const adding = ref(false)
const error = ref('')

const maxStock = computed(() => product.value?.count ?? 99)

function decr(): void {
  if (count.value > 1) {
    count.value--
  }
}

function incr(): void {
  if (count.value < maxStock.value) {
    count.value++
  }
}

async function handleAddToCart(): Promise<void> {
  if (!product.value) return
  try {
    adding.value = true
    await cartStore.addToCart(product.value.id, count.value)
    router.push('/cart')
  } catch (err: unknown) {
    const msg = err instanceof Error ? err.message : 'Не удалось добавить товар в корзину'
    alert(msg)
  } finally {
    adding.value = false
  }
}

onMounted(async () => {
  try {
    loading.value = true
    error.value = ''

    const id = route.params.id as string
    const res = await api<Product | ProductResponse>(`/products/${id}`)

    if (res && typeof res === 'object') {
      if ('id' in res) {
        product.value = res as Product
      } else if ('product' in res && res.product) {
        product.value = res.product
      } else if ('data' in res && res.data) {
        product.value = res.data
      } else if ('item' in res && res.item) {
        product.value = res.item
      }
    }
  } catch (err: unknown) {
    console.error('Ошибка загрузки товара:', err)
    error.value = err instanceof Error ? err.message : 'Ошибка загрузки товара'
  } finally {
    loading.value = false
  }
})
</script>

<template>
  <main class="product-detail">
    <div v-if="loading" style="text-align: center; padding: 60px">Загрузка товара...</div>
    <div v-else-if="error" style="text-align: center; padding: 60px; color: red">{{ error }}</div>
    <div v-else-if="!product" style="text-align: center; padding: 60px">Товар не найден</div>
    <div v-else class="product-detail-card">
      <div class="product-detail-image">
        <img
          :src="
            product.image?.startsWith('/')
              ? product.image
              : `/images/${product.image || 'product.svg'}`
          "
          :alt="product.title"
        />
      </div>

      <div class="product-detail-info">
        <h1 class="product-detail-title">{{ product.title }}</h1>
        <p class="product-detail-description">{{ product.description }}</p>
        <div class="product-detail-price">{{ product.cost.toLocaleString('ru-RU') }} ₽</div>

        <div
          class="product-detail-stock"
          :style="maxStock === 0 ? 'background: #ffebee; color: #c62828;' : ''"
        >
          {{ maxStock > 0 ? `В наличии: ${maxStock} шт.` : 'Нет в наличии' }}
        </div>

        <div v-if="authStore.isAuth" class="product-detail-actions">
          <div v-if="maxStock > 0" class="quantity-selector">
            <button type="button" @click="decr">−</button>
            <input type="number" :value="count" readonly />
            <button type="button" @click="incr">+</button>
          </div>
          <button
            type="button"
            class="btn-add-to-cart"
            :disabled="adding || maxStock === 0"
            :style="maxStock === 0 ? 'background: #ccc; cursor: not-allowed;' : ''"
            @click="handleAddToCart"
          >
            <span v-if="adding">Добавление...</span>
            <span v-else-if="maxStock === 0">Закончился</span>
            <span v-else>Добавить в корзину</span>
          </button>
        </div>

        <div v-else class="product-detail-actions">
          <router-link to="/auth" class="btn-login-to-buy"> Войдите, чтобы купить </router-link>
        </div>
      </div>
    </div>
  </main>
</template>
