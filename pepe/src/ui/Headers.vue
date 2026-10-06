<script lang="ts">
import { useCookies } from '@vueuse/integrations/useCookies'

// let cartCount: number = 0;
// let cartError: string;

const cookies = useCookies(['token'])
const token = cookies.get('token')

if (token === null) {
  // добавить авто выход из аккаунта
  console.log('babaika')
}

// if (isset($_COOKIE["saveLogin"])) {
//     $login = $_COOKIE["saveLogin"];
//     $user_query = mysqli_query(
//         $conn,
//         "SELECT id FROM users WHERE login = '" .
//             mysqli_real_escape_string($conn, $login) .
//             "'",
//     );
//     $user = mysqli_fetch_assoc($user_query);
//     if ($user) {
//         $userId = $user["id"];
//         $cart_query = mysqli_query(
//             $conn,
//             "SELECT SUM(count) AS total FROM basket WHERE user_id = $userId",
//         );
//         $cart = mysqli_fetch_assoc($cart_query);
//         $cartCount = (int) $cart["total"];
//     }
// }
</script>
<template>
  <header>
    <img class="header-left" src="/images/logo.svg" alt="logo" />
    <div class="nav-buttons-header header-center">
      <span class="nav-indicator"></span>
      <!-- <a href="./index.php" class="header-button <?php echo $current_page ==
    "index.php"
        ? "active"
        : ""; ?>">HOME</a>
    <a href="./shop.php" class="header-button <?php echo $current_page ==
    "shop.php"
        ? "active"
        : ""; ?>">SHOP</a>
    <a href="./index.php#about" class="header-button">ABOUT</a>
    <?php $isLoggedIn = isset($_COOKIE["saveLogin"]); ?>
    <?php if (!$isLoggedIn): ?>
      <a href="./signIn.php"
        class="header-button <?php echo $current_page == "signIn.php" ||
        $current_page == "signUp.php"
            ? "active"
            : ""; ?>">
        SIGN IN
      </a>
    <?php else: ?>
      <a href="./profile.php" class="header-button <?php echo $current_page ==
      "profile.php"
          ? "active"
          : ""; ?>">
        PROFILE
      </a>
    <?php endif; ?> -->
    </div>
    <div class="additional-header-buttons header-right">
      <button name="search" type="button" class="header-button">
        <img src="/images/search.svg" alt="search" />
      </button>
      <a href="./cart.php" class="header-button" style="position: relative">
        <img src="/images/cart.svg" alt="cart" />
        <!-- <?php if ($cartCount > 0): ?>
          <label id="counter"><?php echo $cartCount; ?></label>
      <?php endif; ?> -->
      </a>
      <button type="button" class="header-button">CONNECT WALLET</button>
    </div>
  </header>
</template>
