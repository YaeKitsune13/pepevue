<script setup lang="ts">
import { onMounted } from 'vue'
import Headers from './ui/Headers.vue'
import { useAuthStore } from './stores/auth'
import { useCartStore } from './stores/cart'

const authStore = useAuthStore()
const cartStore = useCartStore()

onMounted(async () => {
  await authStore.fetchMe()
  if (authStore.isAuth) {
    await cartStore.fetchCart()
  }
})
</script>

<template>
  <Headers />
  <router-view />
</template>
