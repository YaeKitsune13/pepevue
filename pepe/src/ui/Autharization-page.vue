<template>
  <body>
    <main>
      <form action="phpconnects/user_auth.php" method="POST">
        <h1>Авторизация</h1>
        <input type="text" name="loginUser" placeholder="Логин*" required />
        <input type="password" name="passwordUser" placeholder="Пароль" required />
        <button>Войти</button>
        <p>Нет аккаунта?<a href="signUp.php">Зарегистрироваться</a></p>
      </form>
    </main>
  </body>
</template>
