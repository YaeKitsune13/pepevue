<!-- <?php
include "./phpconnects/dbconnect.php";

$ItemStore_query = mysqli_query($conn, "select * from products");
$ItemStore_assoc = mysqli_fetch_all($ItemStore_query);
?> -->
<template>
  <body>
    <main>
      <div class="products-grid">
        <!-- <?php foreach ($ItemStore_assoc as $value):

                $id = $value[0];
                $title = htmlspecialchars($value[1]);
                $image = htmlspecialchars($value[2]);
                $description = htmlspecialchars($value[3]);
                $cost = number_format($value[4], 2, ".", " ");
                ?>
                <div class="product-card">
                    <div class="product-image">
                        <img src="./images/<?= $image ?>" alt="<?= $title ?>" loading="lazy">
                    </div>
                    <div class="product-title"><?= $title ?></div>
                    <div class="product-description"><?= $description ?></div>
                    <div class="product-price"><?= $cost ?> ₽</div>
                    <div class="product-actions">
                        <a href="./productinfo.php?itemid=<?= $id ?>" class="btn-buy">Купить</a>
                    </div>
                </div>
            <?php
            endforeach; ?> -->
      </div>
    </main>
  </body>
</template>
