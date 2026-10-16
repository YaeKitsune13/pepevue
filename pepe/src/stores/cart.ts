import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { api } from '@/api'

export interface ProductInCart {
  id?: number
  ID?: number
  title?: string
  Title?: string
  image?: string
  Image?: string
  cost?: number
  Cost?: number
}

export interface CartItem {
  id: number
  product_id?: number
  ProductID?: number
  count?: number
  Count?: number
  quantity?: number
  product?: ProductInCart
  Product?: ProductInCart
  cost?: number
  Cost?: number
  title?: string
  Title?: string
}

interface CartResponse {
  cart?: CartItem[]
  items?: CartItem[]
  data?: CartItem[]
}

export const useCartStore = defineStore('cart', () => {
  const items = ref<CartItem[]>([])

  const validItems = computed(() => {
    return items.value.filter((i) => {
      const count = Number(i.count ?? i.Count ?? i.quantity ?? 0)
      return count > 0
    })
  })

  const totalCount = computed(() =>
    validItems.value.reduce((acc, i) => acc + Number(i.count ?? i.Count ?? i.quantity ?? 0), 0),
  )

  const totalPrice = computed(() =>
    validItems.value.reduce((acc, i) => {
      const cost = Number(i.product?.cost ?? i.Product?.Cost ?? i.cost ?? i.Cost ?? 0)
      const count = Number(i.count ?? i.Count ?? i.quantity ?? 0)
      return acc + cost * count
    }, 0),
  )

  async function fetchCart(): Promise<void> {
    try {
      const res = await api<CartItem[] | CartResponse>('/cart')

      let rawItems: CartItem[] = []
      if (Array.isArray(res)) {
        rawItems = res
      } else if (res && typeof res === 'object') {
        rawItems = res.cart || res.items || res.data || []
      }

      items.value = rawItems.filter((i) => Number(i.count ?? i.Count ?? i.quantity ?? 0) > 0)
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
