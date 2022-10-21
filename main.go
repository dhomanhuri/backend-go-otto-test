package main

import (
	"otto/test/config"
	_brands "otto/test/features/brands"
	_brandsBusiness "otto/test/features/brands/business"
	_brandsData "otto/test/features/brands/data"
	_brandsPresentation "otto/test/features/brands/presentation"
	_redemptions "otto/test/features/redemptions"
	_redemptionsBusiness "otto/test/features/redemptions/business"
	_redemptionsData "otto/test/features/redemptions/data"
	_redemptionsPresentation "otto/test/features/redemptions/presentation"
	_transactions "otto/test/features/transactions"
	_transactionsBusiness "otto/test/features/transactions/business"
	_transactionsData "otto/test/features/transactions/data"
	_transactionsPresentation "otto/test/features/transactions/presentation"
	_vouchers "otto/test/features/vouchers"
	_vouchersBusiness "otto/test/features/vouchers/business"
	_vouchersData "otto/test/features/vouchers/data"
	_vouchersPresentation "otto/test/features/vouchers/presentation"
	"otto/test/migrations"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db                 *gorm.DB                                   = config.InitDB()
	brandData          _brands.Data                               = _brandsData.Repository(db)
	brandBus           _brands.Bussines                           = _brandsBusiness.BrandBusiness(brandData)
	brandHandler       _brandsPresentation.BrandPres              = _brandsPresentation.BrandHandler(brandBus)
	voucherData        _vouchers.Data                             = _vouchersData.Repository(db)
	voucherBus         _vouchers.Bussines                         = _vouchersBusiness.VoucherBusiness(voucherData)
	voucherHandler     _vouchersPresentation.VoucherPres          = _vouchersPresentation.VoucherHandler(voucherBus)
	transactionData    _transactions.Data                         = _transactionsData.Repository(db)
	transactionBus     _transactions.Bussines                     = _transactionsBusiness.TransactionBusiness(transactionData)
	transactionHandler _transactionsPresentation.TransactionsPres = _transactionsPresentation.TransactionHandler(transactionBus)
	redemptionData     _redemptions.Data                          = _redemptionsData.Repository(db)
	redemptionBus      _redemptions.Bussines                      = _redemptionsBusiness.RedemptionBusiness(redemptionData)
	redemptionHandler  _redemptionsPresentation.RedemptionsPres   = _redemptionsPresentation.RedemptionHandler(redemptionBus)
)

func main() {
	db := config.InitDB()
	migrations.AutoMigrate(db)
	router := gin.Default()
	router.POST("/api/brand", brandHandler.BrandInsert)
	router.GET("/api/brand", brandHandler.GetBrand)
	router.POST("/api/voucher", voucherHandler.VoucherInsert)
	router.GET("/api/voucher", voucherHandler.GetVoucher)
	router.GET("/api/voucher/brand", voucherHandler.GetAllVouchersByBrandsId)
	router.POST("/api/transaction", transactionHandler.TransactionInsert)
	router.POST("/api/transaction/redemption", redemptionHandler.RedemptionInsert)
	router.GET("/api/transaction/redemption", redemptionHandler.GetAllRedemptionsByBrandsId)
	router.Run(":8080")
}
