package domain

type CreateProductUseCase struct {
	ProductRepository
	UUIDRepository
}

func (uc CreateProductUseCase) Execute(product *Product) State{
	uuid, err := uc.UUIDRepository.GetNewUUID()
	if err != nil { return ProductCreateError }

	product.CheckValidPrice()
	product.CheckValidAmount()
	product.UUID = uuid

	err = uc.ProductRepository.CreateProduct(product)
	if err != nil { return ProductCreateError }
	return ProductCreated
}
