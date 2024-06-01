package main

import (
	"ecommerce/models"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func initDatabase() {
	var err error
	db, err = gorm.Open(sqlite.Open("ecommerce.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&models.Product{}, &models.User{}, &models.Order{}, &models.OrderProduct{}, &models.CartItem{})
}

func main() {
	r := gin.Default()
	initDatabase()

	r.GET("/products", getProducts)
	r.GET("/products/:id", getProduct)

	r.POST("/products", createProduct)
	r.DELETE("/orders/:orderID/products/:productID", removeProductFromOrder)
	r.GET("/orders", getOrders)
	r.GET("/orders/:id", getOrder)
	r.POST("/orders", createOrder)
	r.PUT("/orders/:id", updateOrder)

	r.Run(":8080")
}

// curl -X POST http://localhost:8080/products \
//      -H "Content-Type: application/json" \
//      -d '{
//            "Name": "Sample Product",
//            "Description": "This is a sample product",
//            "Price": 19.99,
//            "Stock": 100
//          }'

func createProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := db.Create(&product).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, product)
}

// curl -X DELETE http://localhost:8080/orders/{orderID}/products/{productID}

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

// curl -X GET http://localhost:8080/products

func getProducts(c *gin.Context) {
	var products []models.Product
	if err := db.Find(&products).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, products)
}

// curl -X GET http://localhost:8080/products/{id}

func getProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := db.Where("id = ?", id).First(&product).Error; err != nil {
		c.JSON(404, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(200, product)
}

// curl -X POST http://localhost:8080/orders \
//      -H "Content-Type: application/json" \
//      -d '{
//            "UserID": 1,
//            "Status": "pending",
//            "Products": []
//          }'

func createOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := db.Create(&order).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, order)
}

// curl -X PUT http://localhost:8080/orders/{id} \
//      -H "Content-Type: application/json" \
//      -d '{
//            "UserID": 1,
//            "Status": "completed"
//          }'

func updateOrder(c *gin.Context) {
	id := c.Param("id")
	var order models.Order
	if err := db.Where("id = ?", id).First(&order).Error; err != nil {
		c.JSON(404, gin.H{"error": "Order not found"})
		return
	}

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := db.Save(&order).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, order)
}

// curl -X GET http://localhost:8080/orders

func getOrders(c *gin.Context) {
	var orders []models.Order
	if err := db.Find(&orders).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, orders)
}

// curl -X GET http://localhost:8080/orders/{id}

func getOrder(c *gin.Context) {
	id := c.Param("id")
	var order models.Order
	if err := db.Where("id = ?", id).First(&order).Error; err != nil {
		c.JSON(404, gin.H{"error": "Order not found"})
		return
	}

	c.JSON(200, order)
}
