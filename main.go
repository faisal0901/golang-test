package main

import (
	config "go-test/Config"
	controller "go-test/Controller"
	model "go-test/Model"
	repository "go-test/Repository"
	services "go-test/Services"
	middleware "go-test/middleware"
	"log"

	"github.com/gin-gonic/gin"
)
func main(){
	db := config.SetupDatabase()
	db.AutoMigrate(&model.Token{},&model.Customer{}, &model.Merchant{},&model.UserLog{},&model.Product{},&model.Transaction{})

	repository := repository.NewRepository(db)
	logService:=services.NewLogService(repository)
	tokenService:=services.NewTokenService(repository)
	productService:=services.NewProductService(repository)
	merchantService:=services.NewMerchantService(repository)

	authService := services.NewAuthService(repository,logService,tokenService)
	authController := controller.NewAuthController(authService)

	transactionService:=services.NewTransactionService(repository,productService)
	transactionController:=controller.NewTransactionController(transactionService)
	productController:=controller.NewProductController(productService)
	merchantController:=controller.NewMerchantController(merchantService)
	r := gin.Default()
	public := r.Group("/api")
	
	public.POST("/register", authController.Register)
	public.POST("/login", authController.Login)
	protected := r.Group("/api/index")
	protected.Use(middleware.JwtAuthMiddleware())
	protected.POST("/transaction", transactionController.CreateNewTransaction)
	protected.POST("/product", productController.CreateNewProduct)
	protected.GET("/transaction", transactionController.GetAllTransaction)
	protected.GET("/merchant", merchantController.GetAllMerchant)
	protected.PUT("/logout", authController.Logout)
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
