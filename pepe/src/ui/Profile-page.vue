<!-- <?php
require "phpconnects/dbconnect.php";

// Проверяем наличие куки
if (isset($_COOKIE["saveLogin"])) {
    $saveLogin = $_COOKIE["saveLogin"]; // Запомнили имя переменной

    $query_text = "SELECT * FROM users WHERE login = '$saveLogin'";
    $query_user = mysqli_query($conn, $query_text);

    if ($query_user && mysqli_num_rows($query_user) > 0) {
        $user = mysqli_fetch_array($query_user);
    }
} else {
    exit();
}
?> -->
<template>
  <body>
    <main>
      <form action="phpconnects/user_change.php" method="POST">
        <!-- <h1>Профиль</h1>
            <input value="<?php echo $user["surname"] ??
                ""; ?>" type="text" name="surnameUser" placeholder="Фамилия*"
                required>
            <input value="<?php echo $user["name"] ??
                ""; ?>" type="text" name="nameUser" placeholder="Имя*" required>
            <input value="<?php echo $user["patronymic"] ??
                ""; ?>" type="text" name="patronymicUser"
                placeholder="Отчество">
            <input value="<?php echo $user["login"] ??
                ""; ?>" type="text" name="loginUser" placeholder="Логин*"
                required>
            <input value="<?php echo $user["password"] ??
                ""; ?>" type="password" name="passwordUser"
                placeholder="Пароль" required>
            <button type="submit" name="action" value="update">Изменить</button>
            <button style="background-color: red;" type="submit" name="action" value="delete"
                style="color: red;">Удалить аккаунт</button>
            <button style="background-color: orange;" type="submit" name="action" value="quit" style="color: red;">Выйти
                из аккаунта</button> -->
      </form>
    </main>
  </body>
</template>
