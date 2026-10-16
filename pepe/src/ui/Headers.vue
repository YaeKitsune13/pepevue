<script setup lang="ts">
import { useRoute } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { useCartStore } from '@/stores/cart'

const route = useRoute()
const authStore = useAuthStore()
const cartStore = useCartStore()
</script>

<template>
  <header>
    <router-link to="/" class="header-left">
      <img src="/images/logo.svg" alt="logo" />
    </router-link>

    <div class="nav-buttons-header header-center">
      <router-link to="/" class="header-button" :class="{ active: route.path === '/' }">
        HOME
      </router-link>
      <router-link to="/shop" class="header-button" :class="{ active: route.path === '/shop' }">
        SHOP
      </router-link>
      <a href="/#about" class="header-button">ABOUT</a>

      <router-link
        v-if="!authStore.isAuth"
        to="/auth"
        class="header-button"
        :class="{ active: route.path === '/auth' }"
      >
        SIGN IN
      </router-link>
      <router-link
        v-else
        to="/profile"
        class="header-button"
        :class="{ active: route.path === '/profile' }"
      >
        PROFILE
      </router-link>
    </div>

    <div class="additional-header-buttons header-right">
      <button name="search" type="button" class="header-button">
        <img src="/images/search.svg" alt="search" />
      </button>

      <router-link to="/cart" class="header-button" style="position: relative">
        <img src="/images/cart.svg" alt="cart" />
        <label
          v-if="cartStore.totalCount > 0"
          id="counter"
          style="
            position: absolute;
            top: 2px;
            right: 4px;
            background: #26d135;
            color: white;
            border-radius: 50%;
            padding: 2px 6px;
            font-size: 11px;
            font-weight: bold;
          "
        >
          {{ cartStore.totalCount }}
        </label>
      </router-link>

      <button type="button" class="header-button">CONNECT WALLET</button>
    </div>
  </header>
</template>

<style scoped>
.header-button.active {
  background-color: black !important;
  color: white !important;
}
</style>
