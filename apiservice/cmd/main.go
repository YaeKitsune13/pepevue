// @title           API Service
// @version         1.0
// @description     API для интернет-магазина: продукты, корзина, заказы, аккаунты.
// @termsOfService  http://swagger.io/terms/

// @host      localhost:8080
// @BasePath  /

// @securityDefinitions.apikey  CookieAuth
// @in                          cookie
// @name                        session_token
package main

import (
	"fmt"
	"log"
	"os"

	_ "apiservice/docs"
	"apiservice/internal/handler"
	"apiservice/internal/middleware"
	"apiservice/internal/model"
	"apiservice/internal/repository"
	"apiservice/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("файл .env не найден, использую переменные окружения из системы")
	}
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("не удалось подключиться к БД: ", err)
	}

	if err := db.AutoMigrate(
		&model.Account{},
		&model.Category{},
		&model.Product{},
		&model.Cart{},
		&model.Order{},
		&model.OrderItem{},
		&model.Log{},
	); err != nil {
		log.Fatal("не удалось выполнить миграции: ", err)
	}

	jwtSecret := os.Getenv("JWT_SECRET")

	// Репозитории
	accountRepo := repository.NewAccountRepository(db)
	productRepo := repository.NewProductRepository(db)
	cartRepo := repository.NewCartRepository(db)
	orderRepo := repository.NewOrderRepository(db)
	orderItemRepo := repository.NewOrderItemRepository(db)

	// Сервисы
	accountService := service.NewAccountService(accountRepo)
	productService := service.NewProductService(productRepo)
	cartService := service.NewCartService(cartRepo)
	orderService := service.NewOrderService(orderRepo, orderItemRepo, cartRepo, db)

	// Хендлеры
	accountHandler := handler.NewAccountHandler(accountService, jwtSecret)
	orderHandler := handler.NewOrderHandler(orderService)
	_ = productService // подключить к CartHandler/ProductHandler, когда появятся
	_ = cartService

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Публичные маршруты
	r.POST("/register", accountHandler.Register)
	r.POST("/login", accountHandler.Login)

	// Приватные маршруты (требуют cookie-сессию)
	auth := r.Group("/")
	auth.Use(middleware.AuthRequired(jwtSecret))
	{
		auth.POST("/logout", accountHandler.Logout)
		auth.POST("/order", orderHandler.PlaceOrder)
	}

	r.Run()
}
