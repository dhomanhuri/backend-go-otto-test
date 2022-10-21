package vouchers

type Core struct {
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
	DataIsDuplicate(data Core) error
	DataFind(id string) (Core, error)
	DataGetByBrandID(id string) ([]Core, error)
}

type Bussines interface {
	BussInsert(data Core) error
	BussGet(id string) (Core, error)
	BussGetByBrandID(id string) ([]Core, error)
}
