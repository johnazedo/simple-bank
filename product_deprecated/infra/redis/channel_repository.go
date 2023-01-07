package redis

import (
	"encoding/json"
	"github.com/JohnAzedo/eCommerce/product/domain"
	"github.com/go-redis/redis"
)

type ChannelRepositoryImpl struct {
	*redis.Client
}

func (r ChannelRepositoryImpl) PublishProduct(product *domain.Product) error {
	data, err := json.Marshal(&product)
	if err != nil {
		return err
	}
	err = r.Client.Publish("product_deprecated:messaging", data).Err()
	if err != nil {
		return err
	}
	return nil
}
