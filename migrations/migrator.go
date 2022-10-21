package migrations

import (
	_dataBrand "otto/test/features/brands/data"
	_dataRedemption "otto/test/features/redemptions/data"
	_dataTransaction "otto/test/features/transactions/data"
	_dataVoucher "otto/test/features/vouchers/data"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(_dataBrand.Brand{})
	db.AutoMigrate(_dataVoucher.Voucher{})
	db.AutoMigrate(_dataTransaction.Transaction{})
	db.AutoMigrate(_dataRedemption.Redemptions{})
}
