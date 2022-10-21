package data

import (
	"otto/test/features/vouchers"
)

type Voucher struct {
	ID      int    `json:"id" gorm:"primary_key:auto_increment" binding:"require"`
	Name    string `json:"name" binding:"require"`
	Price   int    `json:"price" binding:"require"`
	BrandID int    `json:"brand_id" binding:"require"`
	Brand   Brand
}

type Brand struct {
	ID   int    `json:"id" gorm:"primary_key:auto_increment"`
	Name string `json:"brand_name"`
}

func (voucher Voucher) ToCore() vouchers.Core {
	brandCore := vouchers.Core{
		ID:      0,
		Name:    voucher.Name,
		Price:   voucher.Price,
		BrandID: voucher.BrandID,
	}
	return brandCore
}

func toCoreList(voucher []Voucher) []vouchers.Core {
	var coreList []vouchers.Core
	for _, val := range voucher {
		coreList = append(coreList, val.ToCore())
	}
	return coreList
}
