package business

import (
	"otto/test/features/vouchers"
)

type VoucherBuss struct {
	voucherData vouchers.Data
}

func VoucherBusiness(data vouchers.Data) vouchers.Bussines {
	return &VoucherBuss{
		voucherData: data,
	}
}

func (repo *VoucherBuss) BussInsert(data vouchers.Core) error {
	dupErr := repo.voucherData.DataIsDuplicate(data)
	if dupErr != nil {
		return dupErr
	}

	result := repo.voucherData.DataInsert(data)
	if result != nil {
		return result
	}
	return nil
}
func (repo *VoucherBuss) BussGet(id string) (vouchers.Core, error) {
	voucherCore, nil := repo.voucherData.DataFind(id)
	return voucherCore, nil
}

func (repo *VoucherBuss) BussGetByBrandID(id string) ([]vouchers.Core, error) {
	result, err := repo.voucherData.DataGetByBrandID(id)
	return result, err
}
