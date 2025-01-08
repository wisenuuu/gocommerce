package main

import (
	"ecommerce/configs"
	"ecommerce/handlers"
	"ecommerce/middlewares"
	"ecommerce/migrations"
	"ecommerce/seeders"
	"net/http"
	_ "net/http/pprof"

	"github.com/gin-gonic/gin"
)

func main() {
	err := configs.ConnectDB()
	if err != nil {
		panic(err)
	}

	migrations.Migrate(configs.DB)

	seeders.DatabaseSeeder(configs.DB)

	route := gin.Default()

	//products
	route.GET("/product", middlewares.AuthMiddleware(), handlers.ListProduct(configs.DB))
	route.GET("/product/:id", middlewares.AuthMiddleware(), handlers.GetProduct(configs.DB))
	route.POST("/product", middlewares.AuthMiddleware(), handlers.CreateProduct(configs.DB))
	route.PUT("/product/:id", middlewares.AuthMiddleware(), handlers.UpdateProduct(configs.DB))
	route.DELETE("/product/:id", middlewares.AuthMiddleware(), handlers.DeleteProduct(configs.DB))

	//product categories
	route.GET("/product-categories", middlewares.AuthMiddleware(), handlers.GetProductCategories(configs.DB))
	route.GET("/product-categories/:id", middlewares.AuthMiddleware(), handlers.GetProductCategoriesById(configs.DB))
	route.POST("/product-categories", middlewares.AuthMiddleware(), handlers.CreateProductCategories(configs.DB))
	route.PUT("/product-categories/:id", middlewares.AuthMiddleware(), handlers.UpdateProductCategories(configs.DB))
	route.DELETE("/product-categories/:id", middlewares.AuthMiddleware(), handlers.DeleteProductCategories(configs.DB))

	//authentication
	route.POST("/login", handlers.Login(configs.DB))
	route.POST("/register", handlers.Register(configs.DB))

	route.POST("/transactions", handlers.CreateTransaction(configs.DB))
	route.GET("/transactions/:id", handlers.GetTransactionWithItems(configs.DB))

	route.GET("/debug/pprof/*pprof", gin.WrapH(http.DefaultServeMux))

	// route.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "WELCOME TO GOLANGG BROOO",
	// 	})
	// })

	// route.GET("/hello", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "Hello World",
	// 	})
	// })

	// route.POST("/login", func(c *gin.Context) {
	// 	var loginData struct {
	// 		Email    string `json:"email"`
	// 		Password string `json:"password"`
	// 	}

	// 	if err := c.ShouldBindJSON(&loginData); err != nil {
	// 		c.JSON(400, gin.H{
	// 			"success": false,
	// 			"message": "Invalid Request Body",
	// 		})
	// 		return
	// 	}

	// 	if loginData.Email == "admin@mail.com" && loginData.Password == "password" {
	// 		c.JSON(200, gin.H{
	// 			"success": true,
	// 			"message": "User logged in successfully",
	// 			"data": gin.H{
	// 				"email": loginData.Email,
	// 			},
	// 		})
	// 	} else {
	// 		c.JSON(401, gin.H{
	// 			"success": false,
	// 			"message": "Invalid Email or Password",
	// 		})
	// 	}
	// })

	route.Run(":8080")
}
