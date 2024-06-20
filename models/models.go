package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID             uint          `gorm:"primaryKey;autoIncrement" json:"id"`
	FirstName      *string       `json:"first_name" validator:"required"`
	LastName       *string       `json:"last_name" validator:"required"`
	Email          *string       `json:"email" validate:"required"`
	Password       *string       `json:"password" validate:"required,min=6"`
	Phone          *string       `json:"phone" validate:"required"`
	Token          *string       `json:"token"`
	RefreshToken   *string       `json:"refresh_token"`
	CreatedAt      time.Time     `json:"created_at"`
	UpdatedAt      time.Time     `json:"updated_at"`
	UserID         uint          `json:"user_id"`
	AddressDetails []Address     `json:"address_details" gorm:"-"` // Define foreign key
	OrderStatus    []Order       `json:"order_status" gorm:"-"`    // Define foreign key
	UserCart       []ProductUser `json:"user_cart" gorm:"-"`       // Define foreign key
}

type Product struct {
	ProductID   uint      `gorm:"primaryKey;autoIncrement" json:"product_id"`
	ProductName string    `json:"product_name"`
	Price       uint64    `json:"price"`
	Rating      string    `json:"rating"`
	Image       string    `json:"image"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ProductUser struct {
	ProductID   uint      `gorm:"primaryKey" json:"product_id"`
	ProductName string    `json:"product_name"`
	Price       uint64    `json:"price"`
	Rating      string    `json:"rating"`
	Image       string    `json:"image"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Address struct {
	AddressID uint      `gorm:"primaryKey;autoIncrement" json:"address_id"`
	House     string    `json:"house"`
	Street    string    `json:"street"`
	City      string    `json:"city"`
	PinCode   string    `json:"pincode"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Order struct {
	OrderID       uint          `gorm:"primaryKey;autoIncrement" json:"order_id"`
	OrderedAt     time.Time     `json:"ordered_at"`
	Price         uint64        `json:"price"`
	Discount      int           `json:"discount"`
	PaymentMethod Payment       `json:"payment_method" gorm:"embedded"`
	OrderCart     []ProductUser `json:"order_cart" gorm:"-"` // Define foreign key
	CreatedAt     time.Time     `json:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at"`
}

type Payment struct {
	DigitalPayment bool      `json:"digital_payment"`
	COD            bool      `json:"cod"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
