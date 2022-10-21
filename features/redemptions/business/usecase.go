package business

import (
	"otto/test/features/redemptions"
)

type RedemptionBuss struct {
	redemptionData redemptions.Data
}

func RedemptionBusiness(data redemptions.Data) redemptions.Bussines {
	return &RedemptionBuss{
		redemptionData: data,
	}
}

func (repo *RedemptionBuss) BussInsert(data redemptions.Core) error {
	result := repo.redemptionData.DataInsert(data)
	if result != nil {
		return result
	}
	return nil
}

func (repo *RedemptionBuss) BussGetByTransactionID(id string) (int, []redemptions.Core, error) {
	result, err := repo.redemptionData.DataGetByTransactionID(id)
	var total int
	for i := 0; i < len(result); i++ {
		total += result[i].Voucher.Price
	}
	return total, result, err
}
