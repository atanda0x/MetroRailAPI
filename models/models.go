package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID              uuid.UUID
	First_Name      *string
	Last_Name       *string
	Password        *string
	Email           *string
	Phone           *string
	Token           *string
	Refresh_Token   *string
	User_ID         string
	UserCart        []ProductUser
	Address_Details []Address
	Order_Status    []Order
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type Product struct {
	Product_ID   uuid.UUID
	Product_Name *string
	Price        *uint64
	Rating       *uint8
	Image        *string
}

type ProductUser struct {
	Product_ID   uuid.UUID
	Product_Name *string
	Price        *uint64
	Rating       *uint8
	Image        *string
}

type Address struct {
	Address_ID uuid.UUID
	House      *string
	Street     *string
	City       *string
	Pincode    *string
}

type Order struct {
	Order_ID       uuid.UUID
	Order_Cart     []ProductUser
	Order_At       time.Time
	Price          int
	Discount       int
	Payment_Method Payment
}

type Payment struct {
	Digital bool
	COD     bool
}
