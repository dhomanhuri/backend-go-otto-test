package dto

import (
	"otto/test/features/transactions"
)

type TransactionRequest struct {
	Desc string `json:"desc" form:"desc" binding:"required"`
}

type TransactionResponse struct {
	ID   int    `json:"id"`
	Desc string `json:"desc"`
}

func FromCore(trancactionCore transactions.Core) TransactionResponse {
	trnasactionsResponse := TransactionResponse{
		ID:   trancactionCore.ID,
		Desc: trancactionCore.Desc,
	}
	return trnasactionsResponse
}
