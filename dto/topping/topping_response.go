package toppingdto

type ToppingResponse struct {
	ID          int    `json:"id"`
	NameTopping string `json:"nameTopping" gorm:"type: varchar(255)"`
	Price       int    `json:"price" gorm:"type: int" `
	Image       string `json:"image" gorm:"type: varchar(255)"`
	//Qty   int    `json:"qty" form:"qty" gorm:"type: int"`
}
