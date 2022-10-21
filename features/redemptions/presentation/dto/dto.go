package dto

import "otto/test/features/redemptions"

type RedemptionRequest struct {
	VoucherID     int `json:"voucher_id"`
	TransactionID int `json:"transaction_id"`
	Quantity      int `json:"quantity"`
}

type ParrentData struct {
	Vouchers []RedemptionResponse `json:"vouchers"`
	Total    int                  `json:"total"`
}

type RedemptionResponse struct {
	ID            int `json:"id" gorm:"primary_key:auto_increment"`
	VoucherID     int `json:"voucher_id"`
	TransactionID int `json:"transaction_id"`
	Quantity      int `json:"quantity"`
}

func FromCoreList(redCore []redemptions.Core) []RedemptionResponse {
	var redemptionResponse []RedemptionResponse
	for _, val := range redCore {
		redemptionResponse = append(redemptionResponse, fromSliceCore(val))
	}
	return redemptionResponse
}

func fromSliceCore(redemptionCore redemptions.Core) RedemptionResponse {
	return RedemptionResponse{
		ID:            redemptionCore.ID,
		VoucherID:     redemptionCore.VoucherID,
		TransactionID: redemptionCore.TransactionID,
		Quantity:      redemptionCore.Quantity,
	}
}
