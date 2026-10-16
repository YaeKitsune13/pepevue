<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { api } from '@/api'

interface Product {
  id: number
  title: string
  image: string
  description: string
  cost: number
}

const products = ref<Product[]>([])
const loading = ref(true)
const error = ref('')

onMounted(async () => {
  try {
    loading.value = true
    error.value = ''

    // Получаем ответ от API
    const res = await api<any>('/products')
    console.log('Ответ от сервера /products:', res)

    // Проверяем формат данных, который прислал Go:
    if (Array.isArray(res)) {
      products.value = res
    } else if (res && Array.isArray(res.products)) {
      products.value = res.products
    } else if (res && Array.isArray(res.data)) {
      products.value = res.data
    } else {
      products.value = []
    }
  } catch (err: any) {
    console.error('Ошибка загрузки каталога:', err)
    error.value = err.message || 'Ошибка загрузки каталога'
  } finally {
    // Гарантированно выключаем статус загрузки
    loading.value = false
  }
})
</script>

<template>
  <main>
    <!-- 1. Идет загрузка -->
    <div v-if="loading" style="text-align: center; padding: 60px; font-size: 18px">
      Загрузка товаров...
    </div>

    <!-- 2. Произошла ошибка -->
    <div v-else-if="error" style="text-align: center; padding: 60px; color: red">
      {{ error }}
    </div>

    <!-- 3. Загрузка завершилась, но товаров в БД нет -->
    <div
      v-else-if="products.length === 0"
      style="text-align: center; padding: 60px; font-size: 18px"
    >
      В магазине пока нет товаров.
    </div>

    <!-- 4. Сетка товаров -->
    <div v-else class="products-grid">
      <div v-for="item in products" :key="item.id" class="product-card">
        <div class="product-image">
          <img
            :src="item.image?.startsWith('/') ? item.image : `/images/${item.image}`"
            :alt="item.title"
            loading="lazy"
          />
        </div>
        <div class="product-title">{{ item.title }}</div>
        <div class="product-description">{{ item.description }}</div>
        <div class="product-price">{{ Number(item.cost || 0).toLocaleString('ru-RU') }} ₽</div>
        <div class="product-actions">
          <router-link :to="`/product/${item.id}`" class="btn-buy">Купить</router-link>
        </div>
      </div>
    </div>
  </main>
</template>
