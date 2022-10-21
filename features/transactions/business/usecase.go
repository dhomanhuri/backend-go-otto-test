package business

import "otto/test/features/transactions"

type TransactionBuss struct {
	transactionData transactions.Data
}

func TransactionBusiness(data transactions.Data) transactions.Bussines {
	return &TransactionBuss{
		transactionData: data,
	}
}

func (repo *TransactionBuss) BussInsert(data transactions.Core) error {
	dupErr := repo.transactionData.DataIsDuplicate(data)
	if dupErr != nil {
		return dupErr
	}

	result := repo.transactionData.DataInsert(data)
	if result != nil {
		return result
	}
	return nil
}
