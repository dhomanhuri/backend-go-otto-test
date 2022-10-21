package brands

type Core struct {
	ID   int
	Name string
}

type Data interface {
	DataInsert(data Core) error
	DataIsDuplicate(data Core) error
	DataFind(id string) (Core, error)
}

type Bussines interface {
	BussInsert(data Core) error
	BussGet(id string) (Core, error)
}
