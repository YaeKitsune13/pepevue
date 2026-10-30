<script setup lang="ts">
import { onMounted } from 'vue'
import { storeToRefs } from 'pinia'
import { useProductsStore } from '@/stores/products'

const productsStore = useProductsStore()
// Используем storeToRefs, чтобы сохранить реактивность реактивных полей стора
const { products, loading, error } = storeToRefs(productsStore)

onMounted(async () => {
  await productsStore.fetchProducts()
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
