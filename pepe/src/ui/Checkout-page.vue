<!-- <?php
require "./phpconnects/dbconnect.php";

if (!isset($_COOKIE["saveLogin"])) {
    echo "<script>window.location.href='./signIn.php';</script>";
    exit();
}

$login = $_COOKIE["saveLogin"];
$user_query = mysqli_query($conn, "SELECT * FROM users WHERE login = '$login'");
$user = mysqli_fetch_assoc($user_query);
$user_id = $user["id"];

// Получаем корзину
$cart_query = mysqli_query(
    $conn,
    "
    SELECT basket.id, basket.count, basket.product_id,
           products.title, products.cost
    FROM basket
    JOIN products ON basket.product_id = products.id
    WHERE basket.user_id = $user_id
",
);
$cart_items = mysqli_fetch_all($cart_query, MYSQLI_ASSOC);

if (empty($cart_items)) {
    echo "<script>window.location.href='./cart.php';</script>";
    exit();
}

$total = 0;
foreach ($cart_items as $item) {
    $total += $item["cost"] * $item["count"];
}
?> -->
<!DOCTYPE html>
<template>
  <body>
    <main style="padding: 20px; max-width: 800px; margin: 0 auto">
      <div
        style="
          background-color: #fff;
          border-radius: 20px;
          padding: 30px;
          margin-bottom: 20px;
          margin-right: 10px;
        "
      >
        <h2>Ваш заказ:</h2>
        <!-- <?php foreach ($cart_items as $item): ?>
                <div style="display: flex; justify-content: space-between; padding: 8px 0; border-bottom: 1px solid #333;">
                    <span><?= htmlspecialchars($item["title"]) ?> × <?= $item[
     "count"
 ] ?></span>
                    <span>$<?= number_format(
                        $item["cost"] * $item["count"],
                        2,
                    ) ?></span>
                </div>
            <?php endforeach; ?> -->
        <div style="text-align: right; padding-top: 10px">
          <!-- <strong>Итого: $<?= number_format($total, 2) ?></strong> -->
        </div>
      </div>

      <!-- Форма оформления -->
      <form action="./phpconnects/order_create.php" method="POST">
        <h2>Данные доставки:</h2>
        <!-- <input type="text" name="surname" value="<?= htmlspecialchars(
                $user["surname"],
            ) ?>"
                placeholder="Фамилия*" required style="display:block; width:100%; margin: 8px 0; padding: 10px;">
            <input type="text" name="name" value="<?= htmlspecialchars(
                $user["name"],
            ) ?>"
                placeholder="Имя*" required style="display:block; width:100%; margin: 8px 0; padding: 10px;">
            <input type="text" name="address" placeholder="Адрес доставки*"
                required style="display:block; width:100%; margin: 8px 0; padding: 10px;">
            <input type="date" name="date_delivery"
                min="<?= date("Y-m-d", strtotime("+1 day")) ?>"
                required style="display:block; width:100%; margin: 8px 0; padding: 10px;"> -->
        <button
          type="submit"
          style="margin-top: 15px; padding: 12px 30px; font-size: 16px; width: 100%"
        >
          Подтвердить заказ ✓
        </button>
      </form>
    </main>
  </body>
</template>
