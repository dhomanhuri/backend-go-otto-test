package transactions

type Core struct {
	ID   int
	Desc string
}

type Data interface {
	DataInsert(data Core) error
	DataIsDuplicate(data Core) error
}

type Bussines interface {
	BussInsert(data Core) error
}
