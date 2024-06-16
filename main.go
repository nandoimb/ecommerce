package main

import (
	"ecommerce/models"
	"ecommerce/repository"
	"ecommerce/services"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var productService services.ProductService

func initDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("ecommerce.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&models.Product{}, &models.User{}, &models.Order{}, &models.OrderProduct{}, &models.CartItem{})
	return db
}

func main() {
	r := gin.Default()
	db := initDatabase()

	productRepo := repository.NewGormProductRepository(db)
	productService = services.NewProductService(productRepo)

	r.GET("/products", getProducts)
	r.GET("/products/:id", getProduct)

	r.POST("/products", createProduct)

	r.DELETE("/orders/:orderID/products/:productID", removeProductFromOrder)

	r.Run(":8080")
}

// createProduct creates a new product
func createProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := productService.CreateProduct(&product); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, product)
}

// getProducts retrieves all products
func getProducts(c *gin.Context) {
	products, err := productService.GetAllProducts()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, products)
}

// getProduct retrieves a product by ID
func getProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid product ID"})
		return
	}

	product, err := productService.GetProductByID(uint(id))
	if err != nil {
		c.JSON(404, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(200, product)
}

// removeProductFromOrder removes a product from an order
func removeProductFromOrder(c *gin.Context) {
	orderID := c.Param("orderID")
	productID := c.Param("productID")

	var orderProduct models.OrderProduct
	if err := db.Where("order_id = ? AND product_id = ?", orderID, productID).First(&orderProduct).Error; err != nil {
		c.JSON(404, gin.H{"error": "OrderProduct not found"})
		return
	}

	if err := db.Delete(&orderProduct).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(204, gin.H{})
}
