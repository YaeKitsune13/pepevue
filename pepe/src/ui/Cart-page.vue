<!-- <?php
require "./phpconnects/dbconnect.php";

if (!isset($_COOKIE["saveLogin"])) {
    echo "<script>window.location.href='./signIn.php';</script>";
    exit();
}

$login = $_COOKIE["saveLogin"];
$user_query = mysqli_query(
    $conn,
    "SELECT id FROM users WHERE login = '$login'",
);
$user = mysqli_fetch_assoc($user_query);
$user_id = $user["id"];

$cart_query = mysqli_query(
    $conn,
    "
    SELECT basket.id, basket.count, basket.product_id,
           products.title, products.image, products.cost
    FROM basket
    JOIN products ON basket.product_id = products.id
    WHERE basket.user_id = $user_id
",
);
$cart_items = mysqli_fetch_all($cart_query, MYSQLI_ASSOC);

$total = 0;
foreach ($cart_items as $item) {
    $total += $item["cost"] * $item["count"];
}
?> -->
<template>
  <body>
    <main class="cart-container">
      <h1 class="cart-title">Корзина</h1>

      <!-- <?php if (empty($cart_items)): ?> -->
      <!-- <div class="empty-cart">
                <p>Корзина пуста.</p>
                <a href="./shop.php">Перейти в магазин ←</a>
            </div>
        <?php else: ?>
            <div class="cart-items-list">
                <?php foreach ($cart_items as $item): ?>
                    <?php
                    $newCountMinus = $item["count"] - 1;
                    $newCountPlus = $item["count"] + 1;
                    ?>
                    <div class="cart-item">
                        <div class="cart-item-image">
                            <img src="./images/<?= htmlspecialchars(
                                $item["image"],
                            ) ?>" alt="<?= htmlspecialchars($item["title"]) ?>">
                        </div>
                        <div class="cart-item-details">
                            <h3 class="cart-item-title"><?= htmlspecialchars(
                                $item["title"],
                            ) ?></h3>
                            <div class="cart-item-price"><?= number_format(
                                $item["cost"],
                                2,
                                ".",
                                " ",
                            ) ?> ₽ / шт.</div>
                        </div>
                        <div class="cart-item-actions">
                            <div class="quantity-controls">
                                <?php if ($item["count"] > 1): ?>
                                    <form action="./phpconnects/cart_update.php" method="POST" style="display: inline;">
                                        <input type="hidden" name="basket_id" value="<?= $item[
                                            "id"
                                        ] ?>">
                                        <input type="hidden" name="count" value="<?= $newCountMinus ?>">
                                        <button type="submit" class="qty-btn">−</button>
                                    </form>
                                <?php else: ?>
                                    <button class="qty-btn disabled" disabled>−</button>
                                <?php endif; ?>
                                <span class="item-count"><?= $item[
                                    "count"
                                ] ?></span>
                                <form action="./phpconnects/cart_update.php" method="POST" style="display: inline;">
                                    <input type="hidden" name="basket_id" value="<?= $item[
                                        "id"
                                    ] ?>">
                                    <input type="hidden" name="count" value="<?= $newCountPlus ?>">
                                    <button type="submit" class="qty-btn">+</button>
                                </form>
                            </div>
                            <form action="./phpconnects/cart_remove.php" method="POST" style="display: inline;">
                                <input type="hidden" name="basket_id" value="<?= $item[
                                    "id"
                                ] ?>">
                                <button type="submit" class="remove-item">✕</button>
                            </form>
                        </div>
                        <div class="cart-item-total">
                            <?= number_format(
                                $item["cost"] * $item["count"],
                                2,
                                ".",
                                " ",
                            ) ?> ₽
                        </div>
                    </div>
                <?php endforeach; ?>
            </div>

            <div class="cart-summary">
                <div class="cart-total">
                    Итого: <span><?= number_format(
                        $total,
                        2,
                        ".",
                        " ",
                    ) ?> ₽</span>
                </div>
                <a href="./checkout.php" class="btn-checkout">
                    Оформить заказ →
                </a>
            </div>
        <?php endif; ?> -->
    </main>
  </body>
</template>
