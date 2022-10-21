package redemptions

type Core struct {
	ID            int
	VoucherID     int
	Voucher       Voucher
	TransactionID int
	Transaction   Transaction
	Quantity      int
}

type Transaction struct {
	ID   int
	Desc string
}

type Voucher struct {
	ID      int
	Name    string
	Price   int
	BrandID int
	Brand   Brand
}

type Brand struct {
	ID   int
	Name string
}

type Data interface {
	DataInsert(data Core) error
	DataGetByTransactionID(id string) ([]Core, error)
}

type Bussines interface {
	BussInsert(data Core) error
	BussGetByTransactionID(id string) (int, []Core, error)
}
