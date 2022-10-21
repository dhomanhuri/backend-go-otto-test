package data

type Transaction struct {
	ID   int    `json:"id" gorm:"primary_key:auto_increment"`
	Desc string `json:"desc" binding:"require"`
}
