package business

import "otto/test/features/brands"

type BrandBuss struct {
	brandData brands.Data
}

func BrandBusiness(data brands.Data) brands.Bussines {
	return &BrandBuss{
		brandData: data,
	}
}

func (repo *BrandBuss) BussInsert(data brands.Core) error {
	dupErr := repo.brandData.DataIsDuplicate(data)
	if dupErr != nil {
		return dupErr
	}

	result := repo.brandData.DataInsert(data)
	if result != nil {
		return result
	}
	return nil
}
func (repo *BrandBuss) BussGet(id string) (brands.Core, error) {
	brandCore, nil := repo.brandData.DataFind(id)
	return brandCore, nil
}
