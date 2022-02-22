package domain

type ProductRepository interface {
	CreateProduct(product *Product) error
}

type UUIDRepository interface {
	GetNewUUID() (string, error)
}

type ChannelRepository interface {
	PublishProduct(product *Product) error
}