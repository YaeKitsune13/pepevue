<!-- <!DOCTYPE html>
<?php
include "./phpconnects/dbconnect.php";

$id = intval($_GET["itemid"]);
$ItemStore_query = mysqli_query($conn, "SELECT * FROM products WHERE id = $id");
$item = mysqli_fetch_assoc($ItemStore_query);

if (!$item) {
    echo "<h1>Товар не найден</h1>";
    exit();
}

$isLoggedIn = isset($_COOKIE["saveLogin"]);
?> -->
<template>
  <body>
    <main class="product-detail">
      <div class="product-detail-card">
        <div class="product-detail-image">
          <!-- <img src="./images/<?= htmlspecialchars(
                    $item["image"],
                 ) ?>" alt="<?= htmlspecialchars($item["title"]) ?>">  -->
        </div>
        <div class="product-detail-info">
          <!-- <h1 class="product-detail-title"><?= htmlspecialchars(
                    $item["title"],
                ) ?></h1>
                <p class="product-detail-description"><?= htmlspecialchars(
                    $item["description"],
                ) ?></p>
                <div class="product-detail-price"><?= number_format(
                    $item["cost"],
                    2,
                    ".",
                    " ",
                ) ?> ₽</div>
                <div class="product-detail-stock">В наличии: <?= $item[
                    "count"
                ] ?> шт.</div>

                <?php if ($isLoggedIn): ?>
                    <form action="./phpconnects/cart_add.php" method="POST" id="addToCartForm">
                        <input type="hidden" name="product_id" value="<?= $item[
                            "id"
                        ] ?>">
                        <div class="product-detail-actions">
                            <div class="quantity-selector">
                                <button type="button" id="decr">−</button>
                                <input type="number" name="count" id="countInput" value="1" min="1" max="<?= $item[
                                    "count"
                                ] ?>">
                                <button type="button" id="incr">+</button>
                            </div>
                            <button type="submit" class="btn-add-to-cart">Добавить в корзину</button>
                        </div>
                    </form>
                <?php else: ?>
                    <div class="product-detail-actions">
                        <a href="./signIn.php" class="btn-login-to-buy">Войдите, чтобы купить</a>
                    </div>
                <?php endif; ?> -->
        </div>
      </div>
    </main>

    <script>
      // Управление количеством товара
      const decrBtn = document.getElementById('decr')
      const incrBtn = document.getElementById('incr')
      const countInput = document.getElementById('countInput')

      if (decrBtn && incrBtn && countInput) {
        decrBtn.addEventListener('click', () => {
          let val = parseInt(countInput.value)
          if (val > 1) {
            countInput.value = val - 1
          }
        })
        incrBtn.addEventListener('click', () => {
          let val = parseInt(countInput.value)
          let max = parseInt(countInput.getAttribute('max'))
          if (val < max) {
            countInput.value = val + 1
          }
        })
        countInput.addEventListener('change', () => {
          let val = parseInt(countInput.value)
          let max = parseInt(countInput.getAttribute('max'))
          if (isNaN(val)) val = 1
          if (val < 1) val = 1
          if (val > max) val = max
          countInput.value = val
        })
      }
    </script>
  </body>
</template>
