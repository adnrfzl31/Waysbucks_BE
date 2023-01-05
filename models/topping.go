package models

import "time"

type Topping struct {
	ID          int    `json:"id"`
	Nametopping string `json:"nametopping" form:"nametopping" gorm:"type: varchar(255)"`
	Price       int    `json:"price" form:"price" gorm:"type: varchar(255)"`
	Image       string `json:"image" form:"image" gorm:"type: varchar(255)"`
	//Qty      int          `json:"qty" form:"qty"`
	CreateAt time.Time `json:"-"`
	UpdateAt time.Time `json:"-"`
}

type ToppingResponse struct {
	ID          int    `json:"id"`
	Nametopping string `json:"nametopping"`
	Price       int    `json:"price"`
	Image       string `json:"image"`
	//Qty    int    `json:"qty"`
}

type ToppingOrder struct {
	ID          int    `json:"id"`
	Nametopping string `json:"nametopping"`
}

func (ToppingResponse) TableName() string {
	return "topping"
}
