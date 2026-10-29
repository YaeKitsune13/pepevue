<script setup lang="ts">
import ButtonFilter from '@/ui/Button-Filter.vue'
import ButtonBase from '@/ui/Button-Base.vue'
import ButHint from '@/ui/But-Hint.vue'
import ProductCard from '@/ui/Product-Card.vue'
</script>

<template>
  <main>
    <div class="main-head-container">
      <img src="/images/logofull.svg" alt="fulllogo" class="fulllogo" />
      <label>MERCH STORE</label>
      <button @click="$router.push('/shop')">SHOP NOW</button>

      <div class="main-head-items">
        <img src="/images/tshirt.svg" alt="tshirt" />
        <img src="/images/pepe2.svg" alt="pepe2" />
        <img src="/images/cap.svg" alt="cap" />
      </div>
      <img src="/images/pepe1.svg" alt="pepe1" class="pepe1" />
    </div>

    <div class="container">
      <ButtonFilter />

      <div class="products-container">
        <ProductCard />
        <ProductCard />
        <ProductCard />
      </div>

      <div class="about left" id="about">
        <div class="img">
          <img class="first_about" src="/images/logo.svg" alt="logo" />
          <div class="ellipse">
            <div class="ellipse">
              <div class="ellipse"></div>
            </div>
          </div>
        </div>

        <div class="info">
          <label>
            <strong style="font-size: 35px; text-wrap-mode: nowrap">
              About
              <span style="color: var(--background_color); font-size: 35px; text-wrap-mode: nowrap">
                CopeNothing
              </span>
            </strong>
          </label>

          <label class="info_text">
            Are you ready for COPE? Our team has worked hard to create a unique merch for pepe &
            pndc holders.
          </label>

          <button>
            <img src="/images/twiiter.svg" alt="twiter" />
            TWITTER
          </button>
        </div>
      </div>

      <ButtonBase />
      <ButHint />

      <div class="about right">
        <div class="img about2">
          <img src="/images/Pond.svg" alt="logo" />
          <img src="/images/pepe2.svg" alt="logo" />
        </div>

        <div class="info">
          <label>
            <strong style="font-size: 35px; text-wrap-mode: nowrap">
              TOKEN
              <span style="color: var(--background_color); font-size: 35px; text-wrap-mode: nowrap">
                $PNDC
              </span>
            </strong>
          </label>

          <label class="info_text"> PNDC it's a token drived by early adopters of PEPE... </label>
        </div>
      </div>

      <footer>
        <div class="footer-card">
          <div class="footer-left">
            <span class="footer-logo">COPE<br />NOTHING</span>
            <span class="footer-desc">Not only stylish and unique, but also comfortable</span>
            <div class="footer-chars">
              <img src="/images/tshirt.svg" alt="pndc" />
            </div>
          </div>

          <div class="footer-right">
            <nav class="footer-nav">
              <button @click="$router.push('/shop')">SHOP</button>
              <button>COPENOTHING</button>
              <button>HOW TO BUY</button>
              <button>TOKENS</button>
            </nav>

            <div class="footer-support">
              <span>ARE THERE ANY ISSUES WITH THE DELIVERY OR PAYMENT CONFIRMATION?</span>
              <button>CONTACT SUPPORT</button>
            </div>
          </div>
        </div>
      </footer>
    </div>
  </main>
</template>
