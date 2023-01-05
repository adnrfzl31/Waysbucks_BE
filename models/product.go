package models

import "time"

type Product struct {
	ID          int    `json:"id" `
	Nameproduct string `json:"nameproduct" form:"nameproduct" gorm:"type: varchar(255)"`
	Price       int    `json:"price" form:"price" gorm:"type: int"`
	Image       string `json:"image" form:"image" gorm:"type: varchar(255)"`
	//Qty			 int					 `json:"qty" form:"qty" gorm:"type: varchar(255)"`
	CreateAt time.Time `json:"-"`
	UpdateAt time.Time `json:"-"`
}

type ProductResponse struct {
	ID          int    `json:"id"`
	Nameproduct string `json:"nameproduct"`
	Price       int    `json:"price"`
	Image       string `json:"image"`
	//Qty			 int					 			`json:"-"`
}

type ProductUserResponse struct {
	ID          int    `json:"id"`
	Nameproduct string `json:"nameproduct"`
	Price       int    `json:"price"`
	Image       string `json:"image"`
	//Qty    int    `json:"qty"`
}

func (ProductResponse) TableName() string {
	return "products"
}

func (ProductUserResponse) TableName() string {
	return "products"
}
