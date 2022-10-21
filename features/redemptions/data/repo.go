package data

import (
	"errors"
	"fmt"
	"otto/test/features/redemptions"

	"gorm.io/gorm"
)

type MysqlDB struct {
	DBConn *gorm.DB
}

func Repository(db *gorm.DB) redemptions.Data {
	return &MysqlDB{
		DBConn: db,
	}
}

func (db *MysqlDB) DataInsert(data redemptions.Core) error {
	RedemptionModel := Redemptions{
		ID:            data.ID,
		VoucherID:     data.VoucherID,
		TransactionID: data.TransactionID,
		Quantity:      data.Quantity,
	}

	result := db.DBConn.Create(&RedemptionModel)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo MysqlDB) DataGetByTransactionID(id string) ([]redemptions.Core, error) {
	var redemptionModel []Redemptions
	result := repo.DBConn.Preload("Voucher").Where("transaction_id = ?", id).Find(&redemptionModel)
	if result.RowsAffected == 0 {
		return nil, errors.New("failed get data")
	}
	fmt.Println(redemptionModel)

	return toCoreList(redemptionModel), nil
}
