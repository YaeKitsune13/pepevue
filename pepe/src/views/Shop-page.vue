<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { api } from '@/api'

interface ProductItem {
  id: number
  title: string
  image_url: string
  description: string
  cost: number
}

interface ProductsResponse {
  products?: ProductItem[]
  data?: ProductItem[]
}

const products = ref<ProductItem[]>([])
const loading = ref(true)
const error = ref('')

onMounted(async () => {
  try {
    loading.value = true
    error.value = ''

    const res = await api<ProductItem[] | ProductsResponse>('/products')

    if (Array.isArray(res)) {
      products.value = res
    } else if (res && typeof res === 'object') {
      products.value = res.products || res.data || []
    }
  } catch (err: unknown) {
    console.error('Ошибка загрузки каталога:', err)
    error.value = err instanceof Error ? err.message : 'Ошибка загрузки каталога'
  } finally {
    loading.value = false
  }
})
</script>

<template>
  <main>
    <div v-if="loading" style="text-align: center; padding: 60px; font-size: 18px">
      Загрузка товаров...
    </div>

    <div v-else-if="error" style="text-align: center; padding: 60px; color: red">
      {{ error }}
    </div>

    <div
      v-else-if="products.length === 0"
      style="text-align: center; padding: 60px; font-size: 18px"
    >
      В магазине пока нет товаров.
    </div>

    <div v-else class="products-grid">
      <div v-for="item in products" :key="item.id" class="product-card">
        <div class="product-image">
          <img
            :src="
              item.image_url?.startsWith('/')
                ? item.image_url
                : `/images/${item.image_url || 'product.svg'}`
            "
            :alt="item.title"
            loading="lazy"
          />
        </div>
        <div class="product-title">{{ item.title }}</div>
        <div class="product-description">{{ item.description }}</div>
        <div class="product-price">{{ item.cost.toLocaleString('ru-RU') }} ₽</div>
        <div class="product-actions">
          <router-link :to="`/product/${item.id}`" class="btn-buy">Купить</router-link>
        </div>
      </div>
    </div>
  </main>
</template>
