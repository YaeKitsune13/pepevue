<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { api } from '@/api'
import { useAuthStore } from '@/stores/auth'
import { useCartStore } from '@/stores/cart'

const route = useRoute()
const router = useRouter()
const authStore = useAuthStore()
const cartStore = useCartStore()

const product = ref<any | null>(null)
const count = ref(1)
const loading = ref(true)
const adding = ref(false)
const error = ref('')

// Определяем остаток на складе с учетом регистра из Go
const maxStock = computed(() => {
  return Number(
    product.value?.count ??
      product.value?.Count ??
      product.value?.quantity ??
      product.value?.Quantity ??
      product.value?.stock ??
      product.value?.Stock ??
      99,
  )
})

function decr() {
  if (count.value > 1) {
    count.value--
  }
}

function incr() {
  if (count.value < maxStock.value) {
    count.value++
  }
}

async function handleAddToCart() {
  if (!product.value) return
  try {
    adding.value = true
    const productId = Number(product.value.id ?? product.value.ID)
    await cartStore.addToCart(productId, count.value)
    router.push('/cart')
  } catch (err: any) {
    alert(err.message || 'Не удалось добавить товар в корзину')
  } finally {
    adding.value = false
  }
}

onMounted(async () => {
  try {
    loading.value = true
    error.value = ''

    const id = route.params.id
    const res = await api<any>(`/products/${id}`)

    if (res && (res.id || res.ID)) {
      product.value = res
    } else if (res && res.product) {
      product.value = res.product
    } else if (res && res.data) {
      product.value = res.data
    } else {
      product.value = null
    }
  } catch (err: any) {
    error.value = err.message || 'Ошибка загрузки товара'
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
          :src="product.image?.startsWith('/') ? product.image : `/images/${product.image}`"
          :alt="product.title || product.Title"
        />
      </div>

      <div class="product-detail-info">
        <h1 class="product-detail-title">{{ product.title || product.Title }}</h1>
        <p class="product-detail-description">{{ product.description || product.Description }}</p>
        <div class="product-detail-price">
          {{ Number(product.cost || product.Cost || 0).toLocaleString('ru-RU') }} ₽
        </div>
        <!-- В блоке с информацией о товаре: -->
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
