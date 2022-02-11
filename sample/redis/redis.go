package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"sync"
	"time"
)


func NewClient() *redis.Client{
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

func Consumer(client *redis.Client, index int64, wg *sync.WaitGroup) {
	sub := client.Subscribe("chat:messaging")
	channel := sub.Channel()

	for message := range channel {
		fmt.Println("Consumer ", index, " read value: ", message.Payload)
	}

	wg.Done()
}

func Producer(client *redis.Client, wg *sync.WaitGroup) {
	var count = 0

	for {
		time.Sleep(time.Second*3)
		err := client.Publish("chat:messaging", fmt.Sprintf("Hello, this is the message number: %d", count)).Err()
		if err != nil {
			panic(err)
		}
		count++
	}
	wg.Done()
}

