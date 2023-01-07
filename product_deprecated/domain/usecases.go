package domain

type CreateProductUseCase interface {
	Execute(product *Product) error
}

type CreateProductUseCaseImpl struct {
	CreateProductUseCase
	ProductRepository
	ChannelRepository
	UUIDRepository
}

func NewCreateProductUseCase(pr ProductRepository, ur UUIDRepository) CreateProductUseCase{
	return &CreateProductUseCaseImpl{ProductRepository: pr, UUIDRepository: ur}
}

func (uc CreateProductUseCaseImpl) Execute(product *Product) error {
	uuid, err := uc.UUIDRepository.GetNewUUID()
	if err != nil { return err }

	err = product.CheckValidPrice()
	if err != nil { return err }

	product.CheckValidAmount()
	product.UUID = uuid

	err = uc.ProductRepository.CreateProduct(product)
	if err != nil { return err }

	err = uc.ChannelRepository.PublishProduct(product)
	if err != nil { return err }

	return nil
}

