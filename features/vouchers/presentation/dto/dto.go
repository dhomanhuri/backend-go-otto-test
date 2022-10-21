package dto

import "otto/test/features/vouchers"

type VoucherRequest struct {
	Name    string `json:"name" form:"name" binding:"required"`
	Price   int    `json:"price" form:"price" binding:"required"`
	BrandID int    `json:"brand_id" form:"brand_id" binding:"required"`
}

type VoucherResponse struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Price   int    `json:"price"`
	BrandID int    `json:"brand_id"`
}

func FromCore(voucherCore vouchers.Core) VoucherResponse {
	voucherResponse := VoucherResponse{
		ID:      voucherCore.ID,
		Name:    voucherCore.Name,
		Price:   int(voucherCore.Price),
		BrandID: int(voucherCore.BrandID),
	}
	return voucherResponse
}

func FromCoreList(artCore []vouchers.Core) []VoucherResponse {
	var voucherResponse []VoucherResponse
	for _, val := range artCore {
		voucherResponse = append(voucherResponse, fromSliceCore(val))
	}
	return voucherResponse
}

func fromSliceCore(voucherCore vouchers.Core) VoucherResponse {
	return VoucherResponse{
		ID:      voucherCore.ID,
		Name:    voucherCore.Name,
		Price:   voucherCore.Price,
		BrandID: voucherCore.BrandID,
	}
}
