// src/stores/products.ts
import { ref, computed } from 'vue'
import { defineStore } from 'pinia'
import { api } from '@/api'

export interface ProductItem {
  id: number
  title: string
  image_url: string
  description: string
  cost: number
  count?: number // добавлено на случай, если приходит с бэка
}

interface ProductsResponse {
  products?: ProductItem[]
  data?: ProductItem[]
}

interface ProductResponse {
  product?: ProductItem
  data?: ProductItem
  item?: ProductItem
}

export const useProductsStore = defineStore('products', () => {
  // Состояние
  const products = ref<ProductItem[]>([])
  const product = ref<ProductItem | null>(null)
  const loading = ref(false)
  const error = ref('')

  // Геттеры
  const maxStock = computed(() => product.value?.count ?? 99)

  // Действия (Actions)
  async function fetchProducts(): Promise<void> {
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
  }

  async function fetchProductById(id: string | number): Promise<void> {
    try {
      loading.value = true
      error.value = ''
      product.value = null

      const res = await api<ProductItem | ProductResponse>(`/products/${id}`)

      if (res && typeof res === 'object') {
        if ('id' in res) {
          product.value = res as ProductItem
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
  }

  return {
    products,
    product,
    loading,
    error,
    maxStock,
    fetchProducts,
    fetchProductById,
  }
})
