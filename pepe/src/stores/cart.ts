import { ref, computed } from 'vue'
import { defineStore } from 'pinia'
import { api } from '@/api'

export interface ProductInCart {
  id: number
  title: string
  image_url?: string
  image?: string
  cost: number
}

export interface CartItem {
  id: number
  product_id: number
  count: number
  product?: ProductInCart
  title?: string
  cost?: number
  image_url?: string
}

export const useCartStore = defineStore('cart', () => {
  const items = ref<CartItem[]>([])

  const validItems = computed(() => {
    return items.value.filter((i) => Number(i.count) > 0)
  })

  const totalCount = computed(() => validItems.value.reduce((acc, i) => acc + Number(i.count), 0))

  const totalPrice = computed(() =>
    validItems.value.reduce((acc, i) => {
      const price = i.product?.cost ?? i.cost ?? 0
      return acc + price * Number(i.count)
    }, 0),
  )

  async function fetchCart(): Promise<void> {
    try {
      const res = await api<any>('/cart')

      let rawItems: CartItem[] = []
      if (Array.isArray(res)) {
        rawItems = res
      } else if (res && typeof res === 'object') {
        rawItems = res.cart || res.items || res.data || []
      }

      items.value = rawItems.filter((i) => Number(i.count) > 0)
    } catch (e: unknown) {
      console.error('Ошибка загрузки корзины:', e)
      items.value = []
    }
  }

  async function addToCart(productId: number, count: number = 1): Promise<void> {
    if (count <= 0) return

    try {
      await api('/cart', {
        method: 'POST',
        body: JSON.stringify({ product_id: productId, count }),
      })
    } catch (err: unknown) {
      const errorMessage = err instanceof Error ? err.message : ''
      if (errorMessage.includes('409') || errorMessage.includes('Conflict')) {
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

  async function updateQuantity(productId: number, count: number): Promise<void> {
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

  async function removeFromCart(productId: number): Promise<void> {
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
