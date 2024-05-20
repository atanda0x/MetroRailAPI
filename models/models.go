package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID              uuid.UUID     `json:"id"`
	First_Name      *string       `json:"first_name" validate:"required,min=2,max=30"`
	Last_Name       *string       `json:"last_name" validate:"required,min=2,max=30"`
	Password        *string       `json:"passwrd" validate:"required,min=6"`
	Email           *string       `json:"emil" validate:"email, required"`
	Phone           *string       `json:"phone" validate:"required"`
	Token           *string       `json:"token"`
	Refresh_Token   *string       `json:"refresh_token"`
	User_ID         string        `json:"user_id"`
	UserCart        []ProductUser `json:"user_cart"`
	Address_Details []Address     `json:"address_details"`
	Order_Status    []Order       `json:"order_status"`
	CreatedAt       time.Time     `json:"created_at"`
	UpdatedAt       time.Time     `json:"update_at"`
}

type Product struct {
	Product_ID   uuid.UUID `json:"id"`
	Product_Name *string   `json:"product_name"`
	Price        *uint64   `json:"price"`
	Rating       *uint8    `json:"rating"`
	Image        *string   `json:"image"`
}

type ProductUser struct {
	Product_ID   uuid.UUID `json:"id"`
	Product_Name *string   `json:"product_name"`
	Price        *uint64   `json:"price"`
	Rating       *uint8    `json:"rating"`
	Image        *string   `json:"image"`
}

type Address struct {
	Address_ID uuid.UUID `json:"id"`
	House      *string   `json:"house"`
	Street     *string   `json:"street"`
	City       *string   `json:"city"`
	Pincode    *string   `json:"pin_code"`
}

type Order struct {
	Order_ID       uuid.UUID     `json:"id"`
	Order_Cart     []ProductUser `json:"order_cart"`
	Order_At       time.Time     `json:"order_at"`
	Price          int           `json:"price"`
	Discount       int           `json:"discount"`
	Payment_Method Payment       `json:"payment_method"`
}

type Payment struct {
	Digital bool `json:"digital"`
	COD     bool `json:"cod"`
}
