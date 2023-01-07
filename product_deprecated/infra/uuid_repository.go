package infra

import (
	"github.com/JohnAzedo/eCommerce/product/domain"
	"github.com/google/uuid"
)

type UUIDRepositoryImpl struct {
	domain.UUIDRepository
}

func (repo UUIDRepositoryImpl) GetNewUUID() (string, error) {
	return uuid.New().String(), nil
}