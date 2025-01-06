package main

import (
	"ecommerce/configs"
	"ecommerce/handlers"
	"ecommerce/migrations"
	"ecommerce/seeders"

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
	route.GET("/product", handlers.ListProduct(configs.DB))
	route.GET("/product/:id", handlers.GetProduct(configs.DB))
	route.POST("/product", handlers.CreateProduct(configs.DB))
	route.PUT("/product/:id", handlers.UpdateProduct(configs.DB))
	route.DELETE("/product/:id", handlers.DeleteProduct(configs.DB))

	//product categories
	route.GET("/product-categories", handlers.GetProductCategories(configs.DB))
	route.GET("/product-categories/:id", handlers.GetProductCategoriesById(configs.DB))
	route.POST("/product-categories", handlers.CreateProductCategories(configs.DB))
	route.PUT("/product-categories/:id", handlers.UpdateProductCategories(configs.DB))
	route.DELETE("/product-categories/:id", handlers.DeleteProductCategories(configs.DB))

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
