<head>
  <meta charset="UTF-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
  <title>PepkaStore</title>
  <link rel="stylesheet" href="./styles/main.css" />
  <link rel="stylesheet" href="./styles/header.css" />
  <link rel="stylesheet" href="./styles/body.css" />
  <link rel="stylesheet" href="./styles/buyhints.css">
</head>
<template>
  <body>
    <!-- <?php include "./components/header.php"; ?> -->
    <main>
      <div class="main-head-container">
        <img src="./images/logofull.svg" alt="fulllogo" class="fulllogo" />
        <label>MERCH STORE</label>
        <button>SHOP NOW</button>
        <div class="main-head-items">
          <img src="./images/tshirt.svg" alt="tshirt" />
          <img src="./images/pepe2.svg" alt="pepe2" />
          <img src="./images/cap.svg" alt="cap" />
        </div>
        <img src="./images/pepe1.svg" alt="pepe1" class="pepe1" />
      </div>
      <div class="container">
        <!-- <?php include "./components/buttons-filter.php"; ?> -->
        <div class="products-container">
          <!-- <?php include "./components/product-card.php"; ?>
        <?php include "./components/product-card.php"; ?>
        <?php include "./components/product-card.php"; ?>
        <?php include "./components/product-card.php"; ?>
        <?php include "./components/product-card.php"; ?>
        <?php include "./components/product-card.php"; ?>
        <?php include "./components/product-card.php"; ?>
        <?php include "./components/product-card.php"; ?>
        <?php include "./components/product-card.php"; ?>
        <?php include "./components/product-card.php"; ?>
        <?php include "./components/product-card.php"; ?>
        <?php include "./components/product-card.php"; ?>
        <?php include "./components/product-card.php"; ?>
        <?php include "./components/product-card.php"; ?>
        <?php include "./components/product-card.php"; ?> -->
        </div>
        <div class="about left" id="about">
          <div class="img">
            <img class="first_about" src="./images/logo.svg" alt="logo" />
            <div class="ellipse">
              <div class="ellipse">
                <div class="ellipse"></div>
              </div>
            </div>
          </div>
          <div class="info">
            <label
              ><strong style="font-size: 35px; text-wrap-mode: nowrap"
                >About
                <span
                  style="color: var(--background_color); font-size: 35px; text-wrap-mode: nowrap"
                  >CopeNothing</span
                ></strong
              ></label
            >
            <label class="info_text"
              >Are you ready for COPE? Our team has worked hard to create a unique merch for pepe &
              pndc holders. All our products are manufactured using the most modern technologies and
              high-quality materials. They are not only stylish and unique, but also comfortable to
              use. We strive to provide our customers with a high level of satisfaction and joy from
              shopping.</label
            >
            <button>
              <img src="./images/twiiter.svg" alt="twiter" />
              TWITTER
            </button>
          </div>
        </div>
        <!-- <?php include "./components/buttons.php"; ?>
      <?php include "./components/buy-hint.php"; ?> -->
        <div class="about right">
          <div class="img about2">
            <img src="./images/Pond.svg" alt="logo" />
            <img src="./images/pepe2.svg" alt="logo" />
          </div>
          <div class="info">
            <label
              ><strong style="font-size: 35px; text-wrap-mode: nowrap"
                >TOKEN<span
                  style="color: var(--background_color); font-size: 35px; text-wrap-mode: nowrap"
                  >$PNDC</span
                ></strong
              ></label
            >
            <label class="info_text"
              >PNDC it's a token drived by early adopters of PEPE which according to the Pauly0x
              (founder) will change the game once and for all!</label
            >
            <div class="custom-select-container">
              <button type="button" class="select-button">
                TOKEN CONTRACT
                <img src="./images/arrow-down.svg" class="arrow-icon" alt="arrow" />
              </button>

              <div class="select-dropdown">
                <a href="./" target="_blank" class="dropdown-link">LINK</a>
                <a href="./" target="_blank" class="dropdown-link">LINK</a>
                <a href="./" class="dropdown-link">LINK</a>
              </div>
            </div>
          </div>
        </div>
        <div class="about left">
          <div class="img about3">
            <img src="./images/peep.svg" alt="logo" />
            <img src="./images/pipipepe.svg" alt="pipipepe" />
            <img src="./images/pepe1.svg" alt="pepe1" />
          </div>
          <div class="info">
            <label
              ><strong style="font-size: 35px; text-wrap-mode: nowrap"
                >TOKEN<span
                  style="color: var(--background_color); font-size: 35px; text-wrap-mode: nowrap"
                  >$PEPE</span
                ></strong
              ></label
            >
            <label class="info_text"
              >PEPE is a deflationary memecoin launched on Ethereum. The cryptocurrency was created
              as a tribute to the Pepe the Frog internet meme, created by Matt Furie, which gained
              popularity in the early 2000s.</label
            >
            <div class="custom-select-container">
              <button type="button" class="select-button">
                TOKEN CONTRACT
                <img src="./images/arrow-down.svg" class="arrow-icon" alt="arrow" />
              </button>

              <div class="select-dropdown">
                <a href="./" target="_blank" class="dropdown-link">LINK</a>
                <a href="./" target="_blank" class="dropdown-link">LINK</a>
                <a href="./" class="dropdown-link">LINK</a>
              </div>
            </div>
          </div>
        </div>
        <div class="recomendation">
          <div class="card">
            <img src="./images/product.svg" alt="product" />
            <div class="card-bottom">
              <label>CAPS</label>
              <button>›</button>
            </div>
          </div>
          <div class="card">
            <img src="./images/cover.svg" alt="cover" />
            <div class="card-bottom">
              <label>T SHIRTS</label>
              <button>›</button>
            </div>
          </div>
        </div>
        <footer>
          <div class="footer-card">
            <div class="footer-left">
              <span class="footer-logo">COPE<br />NOTHING</span>
              <span class="footer-desc">Not only stylish and unique, but also comfortable</span>
              <div class="footer-chars">
                <img src="./images/tshirt.svg" alt="pndc" />
              </div>
            </div>
            <div class="footer-right">
              <nav class="footer-nav">
                <button>SHOP</button>
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
          <div class="footer-bottom">
            <span>COPYRIGHT 2023 COPENOTHING. ALL RIGHTS RESERVED</span>
            <div>
              <a href="#">TERMS OF SERVICE</a>
              <a href="#">PRIVACY POLICY</a>
            </div>
          </div>
        </footer>
      </div>
    </main>
  </body>
</template>
