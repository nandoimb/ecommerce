package models

type Product struct {
	ID          uint    `json:"id" gorm:"primaryKey"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
}

type User struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Order struct {
	ID       uint           `json:"id" gorm:"primaryKey"`
	UserID   uint           `json:"user_id"`
	Status   string         `json:"status"`
	Products []OrderProduct `json:"products" gorm:"foreignKey:OrderID"`
}

type OrderProduct struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	OrderID   uint `json:"order_id"`
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}

type CartItem struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	UserID    uint `json:"user_id"`
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}
