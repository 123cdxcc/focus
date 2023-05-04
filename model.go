package main

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id        int
	CreatedAt time.Time `gorm:"type:timestamp"`
	UpdatedAt time.Time `gorm:"type:timestamp"`
	DeletedAt time.Time `gorm:"type:timestamp"`
	Mobile    string
	Nickname  string
}

type Order struct {
	gorm.Model
	productName string
	totalPrice  string
	count       int
	unitPrice   string
	status      int8
	payType     int8
}

type RespPage[T any] struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Data    *RespData[T] `json:"data"`
}
type RespData[T any] struct {
	Total   int  `json:"total"`
	Data    []*T `json:"data"`
	Page    int  `json:"page"`
	Pers    int  `json:"pers"`
	HasNext bool `json:"hasNext"`
}
type Total struct {
	Total int
}
type UserOrder struct {
	Id          string `json:"id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	Nickname    string `json:"nickname"`
	Mobile      string `json:"mobile"`
	ProductName string `json:"product_name"`
	TotalPrice  string `json:"total_price"`
	Count       string `json:"count"`
	UnitPrice   string `json:"unit_price"`
	Status      int    `json:"status"`
	PayType     int    `json:"pay_type"`
}
