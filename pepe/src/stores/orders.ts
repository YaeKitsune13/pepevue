// src/stores/orders.ts
import { ref } from 'vue'
import { defineStore } from 'pinia'
import { api } from '@/api'
import type { ProductItem } from './products'

export interface OrderItem {
  ID?: number
  id?: number
  order_id?: number
  product_id: number
  count: number
  cost: number
  Product?: ProductItem
  product?: ProductItem
}

export interface Order {
  ID?: number
  id?: number
  price_all: number
  user_id: number
  status: string
  address: string
  date_delivery: string
  CreatedAt?: string
  created_at?: string
  items?: OrderItem[]
}

export const useOrdersStore = defineStore('orders', () => {
  const orders = ref<Order[]>([])
  const loading = ref(false)
  const error = ref('')

  // Получение заказов пользователя по user_id
  async function fetchOrdersByUserId(userId: number): Promise<void> {
    try {
      loading.value = true
      error.value = ''

      const res = await api<Order[] | { orders?: Order[]; data?: Order[] }>(`/orders/get/${userId}`)

      if (Array.isArray(res)) {
        orders.value = res
      } else if (res && typeof res === 'object') {
        orders.value = res.orders || res.data || []
      }
    } catch (err: unknown) {
      console.error('Ошибка загрузки заказов:', err)
      error.value = err instanceof Error ? err.message : 'Ошибка загрузки заказов'
      orders.value = []
    } finally {
      loading.value = false
    }
  }

  // Оформление заказа из корзины
  async function placeOrder(address: string, dateDelivery?: string): Promise<Order> {
    const payload: { address: string; date_delivery?: string } = { address }
    if (dateDelivery) {
      payload.date_delivery = dateDelivery
    }

    return await api<Order>('/order', {
      method: 'POST',
      body: JSON.stringify(payload),
    })
  }

  return {
    orders,
    loading,
    error,
    fetchOrdersByUserId,
    placeOrder,
  }
})
