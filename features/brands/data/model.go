package data

import "otto/test/features/brands"

type Brand struct {
	ID   int    `json:"id" gorm:"primary_key:auto_increment"`
	Name string `json:"name"`
}

func (brand Brand) ToCore() brands.Core {
	brandCore := brands.Core{
		ID:   int(brand.ID),
		Name: brand.Name,
	}
	return brandCore
}
