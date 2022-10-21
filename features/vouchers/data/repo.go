package data

import (
	"errors"
	"fmt"
	"otto/test/features/vouchers"

	"gorm.io/gorm"
)

type MysqlDB struct {
	DBConn *gorm.DB
}

func Repository(db *gorm.DB) vouchers.Data {
	return &MysqlDB{
		DBConn: db,
	}
}

func (db *MysqlDB) DataInsert(data vouchers.Core) error {
	voucherModel := Voucher{
		ID:      data.ID,
		Name:    data.Name,
		Price:   data.Price,
		BrandID: data.BrandID,
	}
	fmt.Println(voucherModel)
	result := db.DBConn.Create(&voucherModel)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (db *MysqlDB) DataIsDuplicate(data vouchers.Core) error {
	voucherModel := Voucher{
		ID:      data.ID,
		Name:    data.Name,
		Price:   data.Price,
		BrandID: int(data.BrandID),
		Brand: Brand{
			ID:   data.Brand.ID,
			Name: data.Brand.Name,
		},
	}
	result := db.DBConn.Where("name = ?", data.Name).Find(&voucherModel)
	if result.Error != nil || result.RowsAffected > 0 {
		return errors.New("duplicate name")
	}
	return nil
}

func (repo MysqlDB) DataFind(id string) (dataCore vouchers.Core, err error) {
	voucherModel := Voucher{}

	result := repo.DBConn.Where("id = ?", id).Find(&voucherModel)
	if result.RowsAffected == 0 {
		return dataCore, errors.New("voucher not found")
	}

	dataCore = voucherModel.ToCore()

	return dataCore, nil
}

func (repo MysqlDB) DataGetByBrandID(id string) ([]vouchers.Core, error) {
	var voucherModel []Voucher
	result := repo.DBConn.Where("brand_id = ?", id).Find(&voucherModel)
	if result.RowsAffected == 0 {
		return nil, errors.New("failed get data")
	}
	return toCoreList(voucherModel), nil
}
