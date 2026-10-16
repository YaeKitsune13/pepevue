// src/stores/cart.ts
import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { api } from '@/api'

export interface CartItem {
  id: number
  product_id?: number
  ProductID?: number
  count?: number
  Count?: number
  quantity?: number
  product?: any
  Product?: any
  cost?: number
  Cost?: number
  title?: string
  Title?: string
}

export const useCartStore = defineStore('cart', () => {
  const items = ref<CartItem[]>([])

  // Отфильтрованные товары (только с количеством > 0)
  const validItems = computed(() => {
    return items.value.filter((i) => {
      const count = Number(i.count ?? i.Count ?? i.quantity ?? 0)
      return count > 0
    })
  })

  // Количество товаров для бейджика в шапке
  const totalCount = computed(() =>
    validItems.value.reduce((acc, i) => acc + Number(i.count ?? i.Count ?? i.quantity ?? 0), 0),
  )

  // Итоговая сумма корзины
  const totalPrice = computed(() =>
    validItems.value.reduce((acc, i) => {
      const cost = Number(i.product?.cost ?? i.Product?.Cost ?? i.cost ?? i.Cost ?? 0)
      const count = Number(i.count ?? i.Count ?? i.quantity ?? 0)
      return acc + cost * count
    }, 0),
  )

  async function fetchCart() {
    try {
      const res = await api<any>('/cart')

      let rawItems: CartItem[] = []
      if (Array.isArray(res)) {
        rawItems = res
      } else if (res && Array.isArray(res.cart)) {
        rawItems = res.cart
      } else if (res && Array.isArray(res.items)) {
        rawItems = res.items
      } else if (res && Array.isArray(res.data)) {
        rawItems = res.data
      }

      // Сразу убираем позиции с нулевым количеством
      items.value = rawItems.filter((i) => Number(i.count ?? i.Count ?? i.quantity ?? 0) > 0)
    } catch (e) {
      console.error('Ошибка загрузки корзины:', e)
      items.value = []
    }
  }

  async function addToCart(productId: number, count: number = 1) {
    if (count <= 0) return

    try {
      await api('/cart', {
        method: 'POST',
        body: JSON.stringify({ product_id: productId, count }),
      })
    } catch (err: any) {
      if (err.message?.includes('409') || err.message?.includes('Conflict')) {
        await api('/cart', {
          method: 'PUT',
          body: JSON.stringify({ product_id: productId, count }),
        })
      } else {
        throw err
      }
    }
    await fetchCart()
  }

  async function updateQuantity(productId: number, count: number) {
    // ЕСЛИ КОЛИЧЕСТВО СТАЛО 0 — УДАЛЯЕМ ТОВАР ИЗ КОРЗИНЫ
    if (count <= 0) {
      await removeFromCart(productId)
      return
    }

    await api('/cart', {
      method: 'PUT',
      body: JSON.stringify({ product_id: productId, count }),
    })
    await fetchCart()
  }

  async function removeFromCart(productId: number) {
    await api(`/cart/${productId}`, {
      method: 'DELETE',
    })
    await fetchCart()
  }

  return {
    items,
    validItems,
    totalCount,
    totalPrice,
    fetchCart,
    addToCart,
    updateQuantity,
    removeFromCart,
  }
})
