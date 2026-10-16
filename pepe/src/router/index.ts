import { createRouter, createWebHistory } from 'vue-router'
import MainPage from '@/views/Main-page.vue'
import ShopPage from '@/views/Shop-page.vue'
import ProductInfo from '@/views/Product-info.vue'
import CartPage from '@/views/Cart-page.vue'
import CheckoutPage from '@/views/Checkout-page.vue'
import ProfilePage from '@/views/Profile-page.vue'
import AuthPage from '@/views/Autharization-page.vue'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: '/', name: 'home', component: MainPage },
    { path: '/shop', name: 'shop', component: ShopPage },
    { path: '/product/:id', name: 'product', component: ProductInfo },
    { path: '/cart', name: 'cart', component: CartPage },
    { path: '/checkout', name: 'checkout', component: CheckoutPage },
    { path: '/profile', name: 'profile', component: ProfilePage },
    { path: '/auth', name: 'auth', component: AuthPage },
  ],
})

export default router
