package data

import (
	"errors"
	"otto/test/features/brands"

	"gorm.io/gorm"
)

type MysqlDB struct {
	DBConn *gorm.DB
}

func Repository(db *gorm.DB) brands.Data {
	return &MysqlDB{
		DBConn: db,
	}
}

func (db *MysqlDB) DataInsert(data brands.Core) error {
	brandModel := Brand{
		ID:   data.ID,
		Name: data.Name,
	}
	result := db.DBConn.Create(&brandModel)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (db *MysqlDB) DataIsDuplicate(data brands.Core) error {
	brandModel := Brand{
		ID:   data.ID,
		Name: data.Name,
	}
	result := db.DBConn.Where("name = ?", data.Name).Find(&brandModel)
	if result.Error != nil || result.RowsAffected > 0 {
		return errors.New("duplicate name")
	}
	return nil
}

func (repo MysqlDB) DataFind(id string) (dataCore brands.Core, err error) {
	brandModel := Brand{}

	result := repo.DBConn.Where("id = ?", id).Find(&brandModel)
	if result.RowsAffected == 0 {
		return dataCore, errors.New("user not found")
	}

	dataCore = brandModel.ToCore()
	return dataCore, nil
}
