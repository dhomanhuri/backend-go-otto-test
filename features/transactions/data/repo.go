package data

import (
	"errors"
	"fmt"
	"otto/test/features/transactions"

	"gorm.io/gorm"
)

type MysqlDB struct {
	DBConn *gorm.DB
}

func Repository(db *gorm.DB) transactions.Data {
	return &MysqlDB{
		DBConn: db,
	}
}

func (db *MysqlDB) DataInsert(data transactions.Core) error {
	TransactionModel := Transaction{
		ID:   data.ID,
		Desc: data.Desc,
	}
	result := db.DBConn.Create(&TransactionModel)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (db *MysqlDB) DataIsDuplicate(data transactions.Core) error {
	TransactionModel := Transaction{}

	result := db.DBConn.Where("transactions.desc = ?", data.Desc).Find(&TransactionModel)

	fmt.Println(result)
	if result.Error != nil || result.RowsAffected > 0 {
		return errors.New("duplicate name")
	}
	return nil
}
