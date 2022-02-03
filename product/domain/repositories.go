package domain

type ProductRepository interface {
	CreateProduct(product *Product) error
}

type UUIDRepository interface {
	GetNewUUID() (string, error)
}