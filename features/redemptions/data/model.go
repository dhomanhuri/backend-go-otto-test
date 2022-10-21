package data

import "otto/test/features/redemptions"

type Redemptions struct {
	ID            int `json:"id" gorm:"primary_key:auto_increment" binding:"require"`
	VoucherID     int `json:"voucher_id" binding:"require"`
	Voucher       Voucher
	TransactionID int `json:"transaction_id" binding:"require"`
	Transaction   Transaction
	Quantity      int `json:"quantity" binding:"require"`
}

type Transaction struct {
	ID   int    `json:"id" gorm:"primary_key:auto_increment"`
	Desc string `json:"desc"`
}

type Voucher struct {
	ID      int    `json:"id" gorm:"primary_key:auto_increment"`
	Name    string `json:"name"`
	Price   int    `json:"price"`
	BrandID int    `json:"brand_id"`
	Brand   Brand
}

type Brand struct {
	ID   int    `json:"id" gorm:"primary_key:auto_increment"`
	Name string `json:"brand_name"`
}

func (redemption Redemptions) ToCore() redemptions.Core {
	brandCore := redemptions.Core{
		ID:        redemption.ID,
		VoucherID: redemption.VoucherID,
		Voucher: redemptions.Voucher{
			Price: redemption.Voucher.Price,
		},
		TransactionID: redemption.TransactionID,
		Transaction:   redemptions.Transaction{},
		Quantity:      redemption.Quantity,
	}
	return brandCore
}

func toCoreList(redemption []Redemptions) []redemptions.Core {
	var coreList []redemptions.Core
	for _, val := range redemption {
		coreList = append(coreList, val.ToCore())
	}
	return coreList
}
