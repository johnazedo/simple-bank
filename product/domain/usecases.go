package domain

type CreateProductUseCase struct {
	ProductRepository
	UUIDRepository
}

func NewCreateProductUseCase(pr ProductRepository, ur UUIDRepository) *CreateProductUseCase{
	return &CreateProductUseCase{pr, ur}
}

func (uc CreateProductUseCase) Execute(product *Product) error {
	uuid, err := uc.UUIDRepository.GetNewUUID()
	if err != nil { return err }

	err = product.CheckValidPrice()
	if err != nil { return err }

	product.CheckValidAmount()
	product.UUID = uuid

	err = uc.ProductRepository.CreateProduct(product)
	if err != nil { return err }

	return nil
}

