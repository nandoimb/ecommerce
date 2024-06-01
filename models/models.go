package models

// Product model
type Product struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"unique"`
	Description string
	Price       float64
	Stock       int
}

// User model
type User struct {
	ID         uint `gorm:"primaryKey"`
	Nombre     string
	Email      string `gorm:"unique"`
	Contraseña string
	Cart       []CartItem
}

// Order model
type Order struct {
	ID       uint `gorm:"primaryKey"`
	UserID   uint
	Products []OrderProduct
	Status   string
}

// OrderProduct model
type OrderProduct struct {
	ID        uint `gorm:"primaryKey"`
	OrderID   uint
	Product   Product `gorm:"foreignKey:ProductID"`
	ProductID uint
	Cantidad  int
}

// CartItem model
type CartItem struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	Product   Product `gorm:"foreignKey:ProductID"`
	ProductID uint
	Cantidad  int
}
