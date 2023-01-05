package productdto

type ProductRespone struct {
	ID          int    `json:"id"`
	NameProduct string `json:"nameProduct" form:"nameProduct" gorm:"type : varchar(255)"`
	Price       int    `json:"price" form:"price" gorm:"type: int"`
	Image       string `json:"image" form:"image" gorm:"type: varchar(255)"`
}

type DeleteResponse struct {
	ID int `json:"id"`
}
