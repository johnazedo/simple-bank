package redis

import (
	"github.com/go-redis/redis"
	"time"
)

func MessageQueueServer() *redis.Client{
	client := redis.NewClient(&redis.Options{
		Addr: ":6379",
		Password: "",
		DB: 0,
	})

	err := client.Ping().Err()
	if err != nil {
		// Sleep for 3 seconds and wait for Redis to initialize
		time.Sleep(3 * time.Second)
		err := client.Ping().Err()
		if err != nil {
			panic(err)
		}
	}

	return client
}